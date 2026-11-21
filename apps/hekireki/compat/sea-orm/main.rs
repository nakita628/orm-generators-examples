// Runs the plan through the entities hekireki-sea-orm wrote, against the database Prisma created.
//
//   compat crud <plan.json> <url>   insert, read, update, relations, delete
//   compat ddl <plan.json> <url>    Schema::create_table_from_entity for every entity, into an empty database
//
// run.mjs copies the generated entities to src/models and writes src/registry.rs, which calls the
// generic functions below for each entity by its table name. Columns are found by
// ColumnTrait::as_str (the database name). A value goes in through ActiveModel::set as the first
// sea_orm::Value variant the field's Rust type takes; it reads back through ModelTrait::get.
mod models;
mod registry;

use base64::Engine;
use sea_orm::sea_query::{ArrayType, Value};
use sea_orm::*;
use serde_json::{json, Map, Value as Json};
use std::collections::HashMap;
use std::panic::{catch_unwind, AssertUnwindSafe};

pub struct State {
    pub rows: HashMap<String, Map<String, Json>>,
    pub keys: HashMap<String, Vec<(String, Value)>>,
}

#[tokio::main]
async fn main() {
    let args: Vec<String> = std::env::args().collect();
    let plan: Json = serde_json::from_str(&std::fs::read_to_string(&args[2]).unwrap()).unwrap();
    let url = connection_url(&args[3]);
    let db = Database::connect(url).await.expect("connect");
    std::panic::set_hook(Box::new(|_| {}));
    if args[1] == "ddl" {
        // Parents first, in the plan's order, as each CREATE TABLE carries its foreign keys.
        for m in plan["models"].as_array().unwrap() {
            if let Some(Err(e)) = registry::create(&db, m["table"].as_str().unwrap()).await {
                eprintln!("{}: {e}", m["table"]);
                std::process::exit(1);
            }
        }
        return;
    }
    let mut st = State { rows: HashMap::new(), keys: HashMap::new() };
    let models = plan["models"].as_array().unwrap();
    let mut results: Vec<Json> = Vec::new();
    for m in models {
        let r = registry::insert(&db, m, &mut st).await.unwrap_or_else(|| json!({ "model": m["name"], "missing": true }));
        results.push(r);
    }
    for (i, m) in models.iter().enumerate() {
        if let Some(rels) = registry::relations(&db, m, &st).await {
            results[i]["relations"] = rels;
        }
    }
    for (i, m) in models.iter().enumerate().rev() {
        if let Some(del) = registry::delete(&db, m, &st).await {
            results[i]["delete"] = del;
        }
    }
    print!("{}", json!({ "orm": "sea-orm", "models": results }));
}

fn connection_url(url: &str) -> String {
    if let Some(path) = url.strip_prefix("file:") {
        return format!("sqlite://{path}?mode=rwc");
    }
    url.replace("postgresql://", "postgres://")
}

fn squash(s: &str) -> String {
    s.to_lowercase().replace('_', "")
}

pub async fn create_table<E: EntityTrait>(db: &DatabaseConnection, entity: E) -> Result<(), DbErr> {
    let backend = db.get_database_backend();
    let schema = Schema::new(backend);
    let stmt = backend.build(&schema.create_table_from_entity(entity));
    db.execute(stmt).await.map(|_| ())
}

pub async fn insert<E>(db: &DatabaseConnection, m: &Json, st: &mut State) -> Json
where
    E: EntityTrait,
    E::Model: Sync + IntoActiveModel<E::ActiveModel>,
    E::ActiveModel: ActiveModelTrait<Entity = E> + ActiveModelBehavior + Send,
    E::Column: Iterable,
{
    let name = m["name"].as_str().unwrap().to_string();
    let mut r = json!({ "model": name });
    let cols: HashMap<String, E::Column> = E::Column::iter().map(|c| (c.as_str().to_string(), c)).collect();
    let unmapped: Vec<Json> = m["fields"].as_array().unwrap().iter().filter(|f| !cols.contains_key(f["column"].as_str().unwrap())).map(|f| f["column"].clone()).collect();
    r["unmapped"] = Json::Array(unmapped);
    if !m["insertable"].as_bool().unwrap() {
        r["skipped"] = json!("a required key closes a cycle");
        return r;
    }
    // ActiveModelBehavior::new() is where the generated client-side defaults (uuid(), cuid()) are.
    let mut am = <E::ActiveModel as ActiveModelBehavior>::new();
    for f in m["fields"].as_array().unwrap() {
        let Some(col) = cols.get(f["column"].as_str().unwrap()) else { continue };
        let value = match f["write"].as_str().unwrap() {
            "value" => f["value"].clone(),
            "parent" => {
                let fk = &f["fk"];
                if fk["available"].as_bool().unwrap() {
                    st.rows.get(fk["model"].as_str().unwrap()).and_then(|row| row.get(fk["column"].as_str().unwrap()).cloned()).unwrap_or(Json::Null)
                } else {
                    Json::Null
                }
            }
            _ => continue,
        };
        if let Err(e) = set(&mut am, *col, f, &value) {
            r["insert"] = json!({ "ok": false, "error": e });
            return r;
        }
    }
    let model = match am.insert(db).await {
        Ok(model) => model,
        Err(e) => {
            r["insert"] = json!({ "ok": false, "error": message(&e) });
            return r;
        }
    };
    r["insert"] = json!({ "ok": true });
    let pk: Vec<(String, Value)> = m["pk"].as_array().unwrap().iter().filter_map(|c| cols.get(c.as_str().unwrap()).map(|col| (c.as_str().unwrap().to_string(), model.get(*col)))).collect();
    r["pk"] = Json::Object(pk.iter().map(|(c, v)| (c.clone(), normal(v))).collect());
    let filter = |pk: &Vec<(String, Value)>| {
        let mut cond = Condition::all();
        for (c, v) in pk {
            cond = cond.add(cols[c].eq(v.clone()));
        }
        cond
    };
    let loaded = match E::find().filter(filter(&pk)).one(db).await {
        Ok(Some(row)) => row,
        Ok(None) => {
            r["read"] = json!({ "ok": false, "error": "no row came back for the key the insert returned" });
            return r;
        }
        Err(e) => {
            r["read"] = json!({ "ok": false, "error": message(&e) });
            return r;
        }
    };
    let values: Map<String, Json> = cols.iter().map(|(c, col)| (c.clone(), normal(&loaded.get(*col)))).collect();
    r["read"] = json!({ "ok": true, "values": values.clone() });
    st.rows.insert(name.clone(), values);
    if let Some(upd) = m.get("update").filter(|u| !u.is_null()) {
        let c = upd["column"].as_str().unwrap();
        let f = m["fields"].as_array().unwrap().iter().find(|f| f["column"].as_str() == Some(c)).unwrap();
        let col = cols[c];
        let result = async {
            let row = E::find().filter(filter(&pk)).one(db).await.map_err(|e| message(&e))?.ok_or("the row is gone")?;
            let mut am: E::ActiveModel = row.into_active_model();
            set(&mut am, col, f, &upd["value"])?;
            am.update(db).await.map_err(|e| message(&e))?;
            let again = E::find().filter(filter(&pk)).one(db).await.map_err(|e| message(&e))?.ok_or("the row is gone")?;
            Ok::<Json, String>(normal(&again.get(col)))
        }
        .await;
        r["update"] = match result {
            Ok(v) => json!({ "ok": true, "value": v }),
            Err(e) => json!({ "ok": false, "error": e }),
        };
    }
    st.keys.insert(name, pk);
    r
}

pub async fn relations<E>(db: &DatabaseConnection, m: &Json, st: &State) -> Option<Json>
where
    E: EntityTrait,
    E::Model: Sync,
    E::Column: Iterable,
    E::Relation: Iterable + std::fmt::Debug,
{
    let pk = st.keys.get(m["name"].as_str().unwrap())?;
    let cols: HashMap<String, E::Column> = E::Column::iter().map(|c| (c.as_str().to_string(), c)).collect();
    let mut out = Vec::new();
    for rel in m["relations"].as_array().unwrap() {
        let mut res = json!({ "name": rel["name"], "kind": rel["kind"] });
        if !st.keys.contains_key(rel["model"].as_str().unwrap()) {
            res["skipped"] = json!("no row on the other side");
            out.push(res);
            continue;
        }
        let Some(variant) = E::Relation::iter().find(|v| squash(&format!("{v:?}")) == squash(rel["name"].as_str().unwrap())) else {
            res["found"] = json!(false);
            out.push(res);
            continue;
        };
        res["found"] = json!(true);
        if rel["model"] == m["name"] {
            // A join of the table to itself needs an alias, which RelationDef does not carry.
            res["skipped"] = json!("a self relation: the harness joins through RelationDef, which has no alias");
            out.push(res);
            continue;
        }
        let mut cond = Condition::all();
        for (c, v) in pk {
            cond = cond.add(cols[c].eq(v.clone()));
        }
        match E::find().filter(cond).join(JoinType::InnerJoin, variant.def()).count(db).await {
            Ok(n) => res["loaded"] = json!(n),
            Err(e) => res["error"] = json!(message(&e)),
        }
        out.push(res);
    }
    Some(Json::Array(out))
}

pub async fn delete<E>(db: &DatabaseConnection, m: &Json, st: &State) -> Option<Json>
where
    E: EntityTrait,
    E::Model: Sync,
    E::Column: Iterable,
{
    let pk = st.keys.get(m["name"].as_str().unwrap())?;
    let cols: HashMap<String, E::Column> = E::Column::iter().map(|c| (c.as_str().to_string(), c)).collect();
    let cond = || {
        let mut cond = Condition::all();
        for (c, v) in pk {
            cond = cond.add(cols[c].eq(v.clone()));
        }
        cond
    };
    let result = async {
        E::delete_many().filter(cond()).exec(db).await.map_err(|e| message(&e))?;
        let left = E::find().filter(cond()).count(db).await.map_err(|e| message(&e))?;
        Ok::<u64, String>(left)
    }
    .await;
    Some(match result {
        Ok(0) => json!({ "ok": true }),
        Ok(_) => json!({ "ok": false, "error": "the row is still there" }),
        Err(e) => json!({ "ok": false, "error": e }),
    })
}

// Sets the column to the first Value variant the field's type takes: ActiveModel::set panics on
// a variant it cannot convert, so each candidate is tried in turn.
fn set<A: ActiveModelTrait>(am: &mut A, col: <A::Entity as EntityTrait>::Column, f: &Json, value: &Json) -> Result<(), String> {
    let candidates = candidates(f, value);
    for v in candidates.iter() {
        let attempt = catch_unwind(AssertUnwindSafe(|| {
            let mut copy = am.clone();
            copy.set(col, v.clone());
            copy
        }));
        if let Ok(done) = attempt {
            *am = done;
            return Ok(());
        }
    }
    Err(format!("no sea_orm::Value variant fits {} ({})", f["column"], f["type"]))
}

fn candidates(f: &Json, value: &Json) -> Vec<Value> {
    let ty = f["type"].as_str().unwrap();
    if f["list"].as_bool().unwrap_or(false) {
        let Some(items) = value.as_array() else { return nulls() };
        let per: Vec<Vec<Value>> = items.iter().map(|v| scalar(ty, f, v)).collect();
        let kinds: [(ArrayType, usize); 6] = [(ArrayType::String, 0), (ArrayType::Int, 0), (ArrayType::BigInt, 0), (ArrayType::Double, 0), (ArrayType::Bool, 0), (ArrayType::ChronoDateTime, 0)];
        let mut out = Vec::new();
        for (kind, _) in kinds {
            let picked: Option<Vec<Value>> = per.iter().map(|c| c.iter().find(|v| array_type(v) == Some(kind.clone())).cloned()).collect();
            if let Some(vals) = picked {
                out.push(Value::Array(kind, Some(Box::new(vals))));
            }
        }
        return out;
    }
    if value.is_null() {
        return nulls();
    }
    scalar(ty, f, value)
}

fn array_type(v: &Value) -> Option<ArrayType> {
    Some(match v {
        Value::String(_) => ArrayType::String,
        Value::Int(_) => ArrayType::Int,
        Value::BigInt(_) => ArrayType::BigInt,
        Value::Double(_) => ArrayType::Double,
        Value::Bool(_) => ArrayType::Bool,
        Value::ChronoDateTime(_) => ArrayType::ChronoDateTime,
        _ => return None,
    })
}

fn scalar(ty: &str, f: &Json, v: &Json) -> Vec<Value> {
    let text = v.as_str().map(|s| s.to_string()).unwrap_or_else(|| v.to_string());
    match ty {
        "String" | "Enum" => {
            let mut out = vec![Value::String(Some(Box::new(text.clone())))];
            if let Ok(u) = uuid::Uuid::parse_str(&text) {
                out.push(Value::Uuid(Some(Box::new(u))));
            }
            out
        }
        "Int" => {
            let n = v.as_i64().unwrap_or_else(|| text.parse().unwrap_or(0));
            vec![Value::Int(Some(n as i32)), Value::BigInt(Some(n)), Value::SmallInt(Some(n as i16)), Value::TinyInt(Some(n as i8)), Value::Unsigned(Some(n as u32))]
        }
        "BigInt" => {
            let n: i64 = text.parse().unwrap_or(0);
            vec![Value::BigInt(Some(n)), Value::Int(Some(n as i32))]
        }
        "Float" => {
            let x = v.as_f64().unwrap_or(0.0);
            vec![Value::Double(Some(x)), Value::Float(Some(x as f32))]
        }
        "Decimal" => {
            let mut out = Vec::new();
            if let Ok(d) = text.parse::<rust_decimal::Decimal>() {
                out.push(Value::Decimal(Some(Box::new(d))));
            }
            if let Ok(d) = text.parse::<sea_orm::prelude::BigDecimal>() {
                out.push(Value::BigDecimal(Some(Box::new(d))));
            }
            out.push(Value::Double(Some(text.parse().unwrap_or(0.0))));
            out
        }
        "Boolean" => vec![Value::Bool(v.as_bool())],
        "DateTime" => {
            let Ok(dt) = chrono::DateTime::parse_from_rfc3339(&text) else { return vec![] };
            let utc = dt.with_timezone(&chrono::Utc);
            vec![
                Value::ChronoDateTime(Some(Box::new(utc.naive_utc()))),
                Value::ChronoDateTimeUtc(Some(Box::new(utc))),
                Value::ChronoDateTimeWithTimeZone(Some(Box::new(dt))),
                Value::ChronoDate(Some(Box::new(utc.date_naive()))),
                Value::ChronoTime(Some(Box::new(utc.time()))),
                // A newtype over text (hekireki's PrismaDateTime for SQLite) parses the ISO text.
                Value::String(Some(Box::new(text.clone()))),
            ]
        }
        "Json" => vec![Value::Json(Some(Box::new(v.clone())))],
        "Bytes" => vec![Value::Bytes(Some(Box::new(base64::engine::general_purpose::STANDARD.decode(&text).unwrap_or_default())))],
        _ => vec![Value::String(Some(Box::new(text)))],
    }
}

fn nulls() -> Vec<Value> {
    vec![
        Value::String(None),
        Value::Int(None),
        Value::BigInt(None),
        Value::SmallInt(None),
        Value::Double(None),
        Value::Float(None),
        Value::Bool(None),
        Value::ChronoDateTime(None),
        Value::ChronoDateTimeUtc(None),
        Value::ChronoDateTimeWithTimeZone(None),
        Value::ChronoDate(None),
        Value::ChronoTime(None),
        Value::Decimal(None),
        Value::BigDecimal(None),
        Value::Json(None),
        Value::Bytes(None),
        Value::Uuid(None),
    ]
}

// A Value in the plan's JSON form.
fn normal(v: &Value) -> Json {
    let iso = |dt: chrono::DateTime<chrono::Utc>| dt.format("%Y-%m-%dT%H:%M:%S%.3fZ").to_string();
    match v {
        Value::Bool(Some(b)) => json!(b),
        Value::TinyInt(Some(n)) => json!(n),
        Value::SmallInt(Some(n)) => json!(n),
        Value::Int(Some(n)) => json!(n),
        Value::BigInt(Some(n)) => if n.abs() > 9_007_199_254_740_992 { json!(n.to_string()) } else { json!(n) },
        Value::Unsigned(Some(n)) => json!(n),
        Value::BigUnsigned(Some(n)) => json!(n),
        Value::Float(Some(x)) => json!(x),
        Value::Double(Some(x)) => json!(x),
        Value::String(Some(s)) => json!(s.as_str()),
        Value::Char(Some(c)) => json!(c.to_string()),
        Value::Bytes(Some(b)) => json!(base64::engine::general_purpose::STANDARD.encode(b.as_slice())),
        Value::Json(Some(j)) => (**j).clone(),
        Value::ChronoDate(Some(d)) => json!(format!("{}T00:00:00.000Z", d.format("%Y-%m-%d"))),
        Value::ChronoTime(Some(t)) => json!(format!("1970-01-01T{}Z", t.format("%H:%M:%S%.3f"))),
        Value::ChronoDateTime(Some(dt)) => json!(iso(dt.and_utc())),
        Value::ChronoDateTimeUtc(Some(dt)) => json!(iso(**dt)),
        Value::ChronoDateTimeLocal(Some(dt)) => json!(iso(dt.with_timezone(&chrono::Utc))),
        Value::ChronoDateTimeWithTimeZone(Some(dt)) => json!(iso(dt.with_timezone(&chrono::Utc))),
        Value::Uuid(Some(u)) => json!(u.to_string()),
        Value::Decimal(Some(d)) => json!(d.normalize().to_string()),
        Value::BigDecimal(Some(d)) => json!(d.normalized().to_string()),
        Value::Array(_, Some(items)) => Json::Array(items.iter().map(normal).collect()),
        _ => Json::Null,
    }
}

fn message(e: &DbErr) -> String {
    let text = e.to_string();
    text.chars().take(600).collect()
}
