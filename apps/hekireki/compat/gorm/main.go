// Runs the plan through the GORM structs hekireki wrote, against the database Prisma created.
//
//	harness crud <plan.json> <url>   insert, read, update, relations, delete
//	harness ddl <plan.json> <url>    db.AutoMigrate(model.All...) into an empty database
//
// run.mjs copies the generated package to ./model beside this file, with model.All listing every
// struct. Columns are found through GORM's own parsed schema (schema.Field.DBName), so the plan's
// column lines up with whatever the Go field is called; a column no field maps is a finding.
package main

import (
	"compat/model"
	"context"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Plan struct {
	Dialect string  `json:"dialect"`
	Models  []Model `json:"models"`
}

type Model struct {
	Name       string     `json:"name"`
	Table      string     `json:"table"`
	PK         []string   `json:"pk"`
	Fields     []Field    `json:"fields"`
	Relations  []Relation `json:"relations"`
	Insertable bool       `json:"insertable"`
	Update     *struct {
		Column string `json:"column"`
		Value  any    `json:"value"`
	} `json:"update"`
}

type Field struct {
	Column   string `json:"column"`
	Type     string `json:"type"`
	List     bool   `json:"list"`
	Required bool   `json:"required"`
	Write    string `json:"write"`
	Value    any    `json:"value"`
	Native   []any  `json:"native"`
	FK       *struct {
		Model     string `json:"model"`
		Column    string `json:"column"`
		Available bool   `json:"available"`
	} `json:"fk"`
}

type Relation struct {
	Name    string   `json:"name"`
	Kind    string   `json:"kind"`
	Model   string   `json:"model"`
	Columns []string `json:"columns"`
}

type entry struct {
	typ    reflect.Type
	schema *schema.Schema
}

var ctx = context.Background()

func main() {
	mode, planPath, url := os.Args[1], os.Args[2], os.Args[3]
	var plan Plan
	raw, err := os.ReadFile(planPath)
	must(err)
	must(json.Unmarshal(raw, &plan))
	db, err := gorm.Open(dialector(plan.Dialect, url), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent), TranslateError: false})
	must(err)
	if mode == "ddl" {
		must(db.AutoMigrate(model.All...))
		return
	}

	cache := &sync.Map{}
	entries := map[string]entry{}
	for _, m := range model.All {
		s, err := schema.Parse(m, cache, db.NamingStrategy)
		must(err)
		entries[s.Table] = entry{reflect.TypeOf(m).Elem(), s}
	}

	type result = map[string]any
	out := map[string]any{"orm": "gorm"}
	results := []result{}
	rows := map[string]map[string]any{}
	keys := map[string]map[string]any{}

	for _, m := range plan.Models {
		r := result{"model": m.Name}
		results = append(results, r)
		e, ok := entries[m.Table]
		if !ok {
			r["missing"] = true
			continue
		}
		unmapped := []string{}
		for _, f := range m.Fields {
			if e.schema.LookUpField(f.Column) == nil {
				unmapped = append(unmapped, f.Column)
			}
		}
		r["unmapped"] = unmapped
		if !m.Insertable {
			r["skipped"] = "a required key closes a cycle"
			continue
		}
		obj := reflect.New(e.typ)
		err := func() error {
			for _, f := range m.Fields {
				sf := fieldOf(e.schema, f.Column)
				if sf == nil {
					continue
				}
				var v any
				switch f.Write {
				case "value":
					v = f.Value
				case "parent":
					if parent, ok := rows[f.FK.Model]; ok && f.FK.Available {
						v = parent[f.FK.Column]
					}
				default:
					continue
				}
				if err := set(sf, obj, f, v); err != nil {
					return fmt.Errorf("set %s: %w", f.Column, err)
				}
			}
			return db.Create(obj.Interface()).Error
		}()
		if err != nil {
			r["insert"] = result{"ok": false, "error": message(err)}
			continue
		}
		r["insert"] = result{"ok": true}
		pk := map[string]any{}
		for _, col := range m.PK {
			sf := fieldOf(e.schema, col)
			if sf == nil {
				continue
			}
			v, _ := sf.ValueOf(ctx, obj.Elem())
			pk[col] = v
		}
		normalPK := map[string]any{}
		for k, v := range pk {
			normalPK[k] = normal(v)
		}
		r["pk"] = normalPK
		read := reflect.New(e.typ)
		if err := db.Where(pkWhere(pk)).First(read.Interface()).Error; err != nil {
			r["read"] = result{"ok": false, "error": message(err)}
			continue
		}
		values := map[string]any{}
		for _, sf := range e.schema.Fields {
			if sf.DBName == "" {
				continue
			}
			v, _ := sf.ValueOf(ctx, read.Elem())
			values[sf.DBName] = normal(v)
		}
		r["read"] = result{"ok": true, "values": values}
		rows[m.Name] = values
		keys[m.Name] = pk
		if m.Update != nil {
			err := db.Model(reflect.New(e.typ).Interface()).Where(pkWhere(pk)).Update(m.Update.Column, m.Update.Value).Error
			if err != nil {
				r["update"] = result{"ok": false, "error": message(err)}
			} else {
				again := reflect.New(e.typ)
				if err := db.Where(pkWhere(pk)).First(again.Interface()).Error; err != nil {
					r["update"] = result{"ok": false, "error": message(err)}
				} else {
					v, _ := fieldOf(e.schema, m.Update.Column).ValueOf(ctx, again.Elem())
					r["update"] = result{"ok": true, "value": normal(v)}
				}
			}
		}
	}

	byModel := map[string]result{}
	for _, r := range results {
		byModel[r["model"].(string)] = r
	}
	for _, m := range plan.Models {
		pk, ok := keys[m.Name]
		if !ok {
			continue
		}
		e := entries[m.Table]
		rels := []result{}
		for _, rel := range m.Relations {
			res := result{"name": rel.Name, "kind": rel.Kind}
			rels = append(rels, res)
			if _, ok := keys[rel.Model]; !ok {
				res["skipped"] = "no row on the other side"
				continue
			}
			name := findRelation(e.schema, rel.Name)
			if name == "" {
				res["found"] = false
				continue
			}
			res["found"] = true
			loaded := reflect.New(e.typ)
			if err := db.Preload(name).Where(pkWhere(pk)).First(loaded.Interface()).Error; err != nil {
				res["ok"] = false
				res["error"] = message(err)
				continue
			}
			count := countLoaded(loaded.Elem().FieldByName(e.schema.Relationships.Relations[name].Field.Name))
			expected := true
			if rel.Kind == "toOne" {
				for _, c := range rel.Columns {
					if rows[m.Name][c] == nil {
						expected = false
					}
				}
			}
			res["ok"] = (expected && count > 0) || (!expected && count == 0)
			res["loaded"] = count
		}
		byModel[m.Name]["relations"] = rels
	}

	for i := len(plan.Models) - 1; i >= 0; i-- {
		m := plan.Models[i]
		pk, ok := keys[m.Name]
		if !ok {
			continue
		}
		e := entries[m.Table]
		r := byModel[m.Name]
		if err := db.Where(pkWhere(pk)).Delete(reflect.New(e.typ).Interface()).Error; err != nil {
			r["delete"] = result{"ok": false, "error": message(err)}
			continue
		}
		var n int64
		db.Model(reflect.New(e.typ).Interface()).Where(pkWhere(pk)).Count(&n)
		if n == 0 {
			r["delete"] = result{"ok": true}
		} else {
			r["delete"] = result{"ok": false, "error": "the row is still there"}
		}
	}

	out["models"] = results
	must(json.NewEncoder(os.Stdout).Encode(out))
}

func fieldOf(s *schema.Schema, column string) *schema.Field {
	if f, ok := s.FieldsByDBName[column]; ok {
		return f
	}
	return nil
}

func findRelation(s *schema.Schema, prismaName string) string {
	squash := func(x string) string { return strings.ToLower(strings.ReplaceAll(x, "_", "")) }
	for name := range s.Relationships.Relations {
		if squash(name) == squash(prismaName) {
			return name
		}
	}
	return ""
}

func countLoaded(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Slice:
		return v.Len()
	case reflect.Pointer:
		if v.IsNil() {
			return 0
		}
		return 1
	case reflect.Struct:
		if v.IsZero() {
			return 0
		}
		return 1
	}
	return 0
}

func pkWhere(pk map[string]any) map[string]any {
	where := map[string]any{}
	for k, v := range pk {
		where[k] = v
	}
	return where
}

// set writes the plan's value into the struct field, in whatever Go type the field has: GORM's own
// conversion first, then the field's sql.Scanner, then a JSON round trip.
func set(sf *schema.Field, obj reflect.Value, f Field, v any) error {
	if v == nil {
		return sf.Set(ctx, obj.Elem(), nil)
	}
	goValue := toGo(f, v)
	if err := sf.Set(ctx, obj.Elem(), goValue); err == nil {
		return nil
	}
	target := reflect.New(sf.FieldType)
	if scanner, ok := target.Interface().(interface{ Scan(any) error }); ok {
		for _, candidate := range []any{goValue, scanValue(f, v)} {
			if err := scanner.Scan(candidate); err == nil {
				return sf.Set(ctx, obj.Elem(), target.Elem().Interface())
			}
		}
	}
	if sf.FieldType.Kind() == reflect.Pointer {
		inner := reflect.New(sf.FieldType.Elem())
		if scanner, ok := inner.Interface().(interface{ Scan(any) error }); ok {
			if err := scanner.Scan(scanValue(f, v)); err == nil {
				return sf.Set(ctx, obj.Elem(), inner.Interface())
			}
		}
	}
	payload := v
	if f.List {
		payload = goValue
	}
	data, _ := json.Marshal(payload)
	if err := json.Unmarshal(data, target.Interface()); err == nil {
		return sf.Set(ctx, obj.Elem(), target.Elem().Interface())
	}
	return fmt.Errorf("cannot put %v (%s) into %s", v, f.Type, sf.FieldType)
}

func toGo(f Field, v any) any {
	if f.List {
		// A BigInt list holds strings in the plan; a Go list of int64 takes numbers.
		if list, ok := v.([]any); ok && f.Type == "BigInt" {
			out := make([]any, len(list))
			for i, x := range list {
				if n, ok := new(big.Int).SetString(fmt.Sprint(x), 10); ok {
					out[i] = n.Int64()
				} else {
					out[i] = x
				}
			}
			return out
		}
		return v
	}
	switch f.Type {
	case "DateTime":
		if t, err := time.Parse(time.RFC3339Nano, fmt.Sprint(v)); err == nil {
			return t
		}
	case "Int":
		return int(number(v))
	case "BigInt":
		n, ok := new(big.Int).SetString(fmt.Sprint(v), 10)
		if ok {
			return n.Int64()
		}
	case "Float":
		return number(v)
	case "Decimal":
		return number(v)
	case "Bytes":
		b, _ := base64.StdEncoding.DecodeString(fmt.Sprint(v))
		return b
	case "Json":
		b, _ := json.Marshal(v)
		return b
	}
	return v
}

// number reads a JSON number from the plan (float64) or a value read back from a parent row
// (int, int64, ...).
func number(v any) float64 {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	}
	var f float64
	fmt.Sscan(fmt.Sprint(v), &f)
	return f
}

func scanValue(f Field, v any) any {
	switch f.Type {
	case "Json":
		b, _ := json.Marshal(v)
		return b
	case "Decimal", "BigInt", "String", "Enum":
		return fmt.Sprint(v)
	case "DateTime":
		return toGo(f, v)
	}
	if f.List {
		b, _ := json.Marshal(v)
		return string(b)
	}
	return v
}

// normal turns a field's value into the plan's JSON form.
func normal(v any) any {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	v = rv.Interface()
	switch x := v.(type) {
	case time.Time:
		return x.UTC().Format("2006-01-02T15:04:05.000Z")
	case []byte:
		if json.Valid(x) && len(x) > 0 && (x[0] == '{' || x[0] == '[') {
			var parsed any
			_ = json.Unmarshal(x, &parsed)
			return parsed
		}
		return base64.StdEncoding.EncodeToString(x)
	case json.RawMessage:
		var parsed any
		_ = json.Unmarshal(x, &parsed)
		return parsed
	case int64:
		if x > 1<<53 || x < -(1<<53) {
			return fmt.Sprint(x)
		}
		return x
	case uint64:
		return x
	}
	// A list (hekireki-gorm's ScalarList, a []T with a Value() that writes PostgreSQL's array
	// text): its elements, not what Value() would write to the database.
	if rv.Kind() == reflect.Slice && rv.Type().Elem().Kind() != reflect.Uint8 {
		list := make([]any, rv.Len())
		for i := range list {
			list[i] = normal(rv.Index(i).Interface())
		}
		return list
	}
	if valuer, ok := v.(driver.Valuer); ok {
		dv, err := valuer.Value()
		if err != nil {
			return fmt.Sprintf("Value() error: %v", err)
		}
		if s, ok := dv.(string); ok && len(s) > 0 && (s[0] == '{' || s[0] == '[') && json.Valid([]byte(s)) {
			var parsed any
			_ = json.Unmarshal([]byte(s), &parsed)
			if _, isMap := parsed.(map[string]any); isMap || isJSONArrayOfNonString(parsed) {
				return parsed
			}
		}
		return normal(dv)
	}
	switch rv.Kind() {
	case reflect.Slice:
		list := make([]any, rv.Len())
		for i := range list {
			list[i] = normal(rv.Index(i).Interface())
		}
		return list
	case reflect.String:
		return rv.String()
	case reflect.Struct, reflect.Map:
		data, _ := json.Marshal(v)
		var parsed any
		_ = json.Unmarshal(data, &parsed)
		return parsed
	}
	return v
}

func isJSONArrayOfNonString(v any) bool {
	list, ok := v.([]any)
	return ok && len(list) > 0
}

func dialector(dialect, url string) gorm.Dialector {
	switch dialect {
	case "postgresql":
		return postgres.Open(url)
	case "mysql":
		// mysql://user:pass@host:port/db -> user:pass@tcp(host:port)/db
		rest := strings.TrimPrefix(url, "mysql://")
		at := strings.LastIndex(rest, "@")
		slash := strings.Index(rest[at:], "/") + at
		return mysql.Open(fmt.Sprintf("%s@tcp(%s)%s?parseTime=true&loc=UTC", rest[:at], rest[at+1:slash], rest[slash:]))
	default:
		return sqlite.Open(strings.TrimPrefix(url, "file:") + "?_foreign_keys=on")
	}
}

func message(err error) string {
	text := err.Error()
	var target interface{ Unwrap() error }
	if errors.As(err, &target) && target.Unwrap() != nil {
		text = target.Unwrap().Error()
	}
	if len(text) > 600 {
		text = text[:600]
	}
	return text
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
