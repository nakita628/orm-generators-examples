// Runs the plan through the tables and entities hekireki-exposed wrote, against the database
// Prisma created.
//
//   compat crud <plan.json> <PrismaSchema class> <url>   insert, read, update, relations, delete
//   compat ddl <plan.json> <PrismaSchema class> <url>    SchemaUtils.create(tables) into an empty database
//
// Columns are found by Column.name, so the plan's column lines up with whatever the property is
// called; values go in through the column type's own valueFromDB, so each lands in the column's
// Kotlin type.
package compat

import com.fasterxml.jackson.databind.JsonNode
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.node.ArrayNode
import com.fasterxml.jackson.databind.node.JsonNodeFactory
import com.fasterxml.jackson.databind.node.ObjectNode
import org.jetbrains.exposed.v1.core.Column
import org.jetbrains.exposed.v1.core.Op
import org.jetbrains.exposed.v1.core.Table
import org.jetbrains.exposed.v1.core.and
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.core.eq
import org.jetbrains.exposed.v1.core.isNull
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass
import org.jetbrains.exposed.v1.jdbc.Database
import org.jetbrains.exposed.v1.jdbc.SchemaUtils
import org.jetbrains.exposed.v1.jdbc.deleteWhere
import org.jetbrains.exposed.v1.jdbc.insert
import org.jetbrains.exposed.v1.jdbc.selectAll
import org.jetbrains.exposed.v1.jdbc.transactions.transaction
import org.jetbrains.exposed.v1.jdbc.update
import java.io.File
import java.math.BigDecimal
import java.net.URI
import java.time.Instant
import java.time.LocalDate
import java.time.LocalDateTime
import java.time.LocalTime
import java.time.OffsetDateTime
import java.time.ZoneOffset
import java.time.format.DateTimeFormatter
import java.util.Base64
import java.util.UUID
import kotlin.reflect.full.companionObjectInstance
import kotlin.reflect.full.memberProperties

val mapper = ObjectMapper()
val nodes: JsonNodeFactory = JsonNodeFactory.instance
val iso: DateTimeFormatter = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ss.SSS'Z'")

@Suppress("UNCHECKED_CAST")
fun main(args: Array<String>) {
    val (mode, planPath, schemaClass, url) = args
    val plan = mapper.readTree(File(planPath))
    val u = URI(url)
    val (user, pass) = u.userInfo.split(":")
    Database.connect("jdbc:postgresql://${u.host}:${u.port}${u.path}", driver = "org.postgresql.Driver", user = user, password = pass)
    val schema = Class.forName(schemaClass).getField("INSTANCE").get(null)
    val tables = schema.javaClass.getMethod("getTables").invoke(schema) as List<Table>
    val enumTypes = schema.javaClass.getMethod("getEnumTypes").invoke(schema) as List<String>
    if (mode == "ddl") {
        transaction {
            enumTypes.forEach { exec(it) }
            SchemaUtils.create(*tables.toTypedArray())
        }
        return
    }
    print(mapper.writeValueAsString(crud(plan, tables, schemaClass.substringBeforeLast('.'))))
}

class Inserted(val table: Table, val key: Map<Column<Any?>, Any?>)

@Suppress("UNCHECKED_CAST")
fun crud(plan: JsonNode, tables: List<Table>, pkg: String): ObjectNode {
    val byName = tables.associateBy { it.tableName.trim('"') }
    val results = nodes.arrayNode()
    val rows = mutableMapOf<String, ObjectNode>()
    val inserted = mutableMapOf<String, Inserted>()
    val models = plan["models"].toList()

    for (m in models) {
        val r = nodes.objectNode().put("model", m["name"].asText())
        results.add(r)
        val table = byName[m["table"].asText()]
        if (table == null) {
            r.put("missing", true)
            continue
        }
        val cols = table.columns.associateBy { it.name.trim('"') } as Map<String, Column<Any?>>
        val unmapped = r.putArray("unmapped")
        m["fields"].filter { !cols.containsKey(it["column"].asText()) }.forEach { unmapped.add(it["column"].asText()) }
        if (!m["insertable"].asBoolean()) {
            r.put("skipped", "a required key closes a cycle")
            continue
        }
        val key: Map<Column<Any?>, Any?>
        try {
            key =
                transaction {
                    val statement =
                        table.insert { st ->
                            for (f in m["fields"]) {
                                val c = cols[f["column"].asText()] ?: continue
                                val value =
                                    when (f["write"].asText()) {
                                        "value" -> f["value"]
                                        "parent" -> {
                                            val fk = f["fk"]
                                            if (fk["available"].asBoolean()) rows[fk["model"].asText()]?.get(fk["column"].asText()) else null
                                        }
                                        else -> continue
                                    }
                                st[c] = toColumn(c, f, value)
                            }
                        }
                    val row = statement.resultedValues!!.single()
                    m["pk"].associate { pk -> cols.getValue(pk.asText()) to row[cols.getValue(pk.asText())] }
                }
            r.putObject("insert").put("ok", true)
        } catch (e: Exception) {
            r.set<ObjectNode>("insert", error(e))
            continue
        }
        val pkJson = r.putObject("pk")
        key.forEach { (c, v) -> pkJson.set<JsonNode>(c.name.trim('"'), normal(v)) }
        try {
            val values =
                transaction {
                    val row = table.selectAll().where(where(key)).singleOrNull() ?: error("no row came back for the key the insert returned")
                    val values = nodes.objectNode()
                    cols.forEach { (name, c) -> values.set<JsonNode>(name, normal(row[c])) }
                    values
                }
            r.putObject("read").put("ok", true).set<ObjectNode>("values", values)
            rows[m["name"].asText()] = values
        } catch (e: Exception) {
            r.set<ObjectNode>("read", error(e))
            continue
        }
        val upd = m["update"]
        if (upd != null && !upd.isNull) {
            try {
                val c = cols.getValue(upd["column"].asText())
                val f = m["fields"].first { it["column"].asText() == upd["column"].asText() }
                val value =
                    transaction {
                        table.update(where = { where(key) }) { it[c] = toColumn(c, f, upd["value"]) }
                        table.selectAll().where(where(key)).single()[c]
                    }
                r.putObject("update").put("ok", true).set<JsonNode>("value", normal(value))
            } catch (e: Exception) {
                r.set<ObjectNode>("update", error(e))
            }
        }
        inserted[m["name"].asText()] = Inserted(table, key)
    }

    // Relations through the DAO: the entity whose EntityClass is over the table, found by its id,
    // and the property named after the Prisma relation.
    val entityClasses = entityClasses(pkg)
    models.forEachIndexed { i, m ->
        val ins = inserted[m["name"].asText()] ?: return@forEachIndexed
        val rels = (results[i] as ObjectNode).putArray("relations")
        for (rel in m["relations"]) {
            val res = rels.addObject().put("name", rel["name"].asText()).put("kind", rel["kind"].asText())
            if (!inserted.containsKey(rel["model"].asText())) {
                res.put("skipped", "no row on the other side")
                continue
            }
            if (ins.key.size != 1) {
                // An entity is found by its id here (findById); a composite key's is not.
                res.put("skipped", "a composite key: the harness finds entities by a single id")
                continue
            }
            val pair = entityClasses.entries.firstOrNull { it.value.table == ins.table }
            val prop = pair?.key?.memberProperties?.firstOrNull { squash(it.name) == squash(rel["name"].asText()) }
            if (pair == null || prop == null) {
                res.put("found", false)
                continue
            }
            res.put("found", true)
            try {
                val count =
                    transaction {
                        val id = ins.key.values.single()
                        val entity = (pair.value as EntityClass<Any, Entity<Any>>).findById(if (id is EntityID<*>) id.value!! else id!!) ?: error("findById found nothing")
                        when (val v = prop.getter.call(entity)) {
                            null -> 0
                            is Iterable<*> -> v.count()
                            else -> 1
                        }
                    }
                res.put("loaded", count)
            } catch (e: Exception) {
                res.put("error", message(e))
            }
        }
    }

    for (i in models.indices.reversed()) {
        val ins = inserted[models[i]["name"].asText()] ?: continue
        val r = results[i] as ObjectNode
        try {
            val left =
                transaction {
                    ins.table.deleteWhere { where(ins.key) }
                    ins.table.selectAll().where(where(ins.key)).count()
                }
            if (left == 0L) r.putObject("delete").put("ok", true) else r.putObject("delete").put("ok", false).put("error", "the row is still there")
        } catch (e: Exception) {
            r.set<ObjectNode>("delete", error(e))
        }
    }
    return nodes.objectNode().put("orm", "exposed").set("models", results)
}

fun where(key: Map<Column<Any?>, Any?>): Op<Boolean> =
    key.map { (c, v) -> if (v == null) c.isNull() else c eq v }.reduce { a, b -> a and b }

fun squash(s: String) = s.lowercase().replace("_", "")

@Suppress("UNCHECKED_CAST")
fun entityClasses(pkg: String): Map<kotlin.reflect.KClass<*>, EntityClass<*, *>> {
    val dir = Thread.currentThread().contextClassLoader.getResources(pkg.replace('.', '/')).toList()
    val out = mutableMapOf<kotlin.reflect.KClass<*>, EntityClass<*, *>>()
    for (url in dir) {
        val names =
            if (url.protocol == "jar") {
                val path = url.path.substringAfter("file:").substringBefore("!")
                java.util.jar.JarFile(path).entries().toList().map { it.name }.filter { it.startsWith(pkg.replace('.', '/') + "/") && it.endsWith(".class") }
            } else {
                File(url.toURI()).listFiles()!!.map { pkg.replace('.', '/') + "/" + it.name }
            }
        for (n in names) {
            val cls = runCatching { Class.forName(n.removeSuffix(".class").replace('/', '.')).kotlin }.getOrNull() ?: continue
            val companion = runCatching { cls.companionObjectInstance }.getOrNull()
            if (companion is EntityClass<*, *>) out[cls] = companion
        }
    }
    return out
}

// The plan's JSON value in the column's Kotlin type, through the column type itself.
fun toColumn(c: Column<Any?>, f: JsonNode, value: JsonNode?): Any? {
    if (value == null || value.isNull) return null
    val type = c.columnType
    if (f["list"].asBoolean()) {
        // An array column takes a List of its element type: each element through the element's
        // own column type (ArrayColumnType.delegate, hekireki's PgListColumnType.element).
        val inner = listOf("delegate", "element").firstNotNullOfOrNull { name -> runCatching { type.javaClass.getMethod("get" + name.replaceFirstChar { it.uppercase() }).invoke(type) as org.jetbrains.exposed.v1.core.IColumnType<*> }.getOrNull() }
        return value.map { v -> fromDb(inner ?: type, f, v) }
    }
    return fromDb(type, f, value)
}

// The column type's valueFromDB on the first JDBC-side shape it accepts: a timestamp column type
// may want a LocalDateTime, an OffsetDateTime or an Instant.
fun fromDb(type: org.jetbrains.exposed.v1.core.IColumnType<*>, f: JsonNode, v: JsonNode): Any? {
    var last: Throwable? = null
    for (candidate in jdbcCandidates(f, v)) {
        try {
            return type.valueFromDB(candidate)
        } catch (e: Throwable) {
            last = e
        }
    }
    throw last ?: IllegalStateException("no value for ${f["column"]}")
}

fun jdbcCandidates(f: JsonNode, v: JsonNode): List<Any> {
    val first = jdbcValue(f, v)
    if (f["type"].asText() != "DateTime") return listOf(first)
    val instant = Instant.parse(v.asText())
    return listOf(first, instant.atOffset(ZoneOffset.UTC), instant, java.sql.Timestamp.from(instant), LocalDateTime.ofInstant(instant, ZoneOffset.UTC))
}

fun jdbcValue(f: JsonNode, v: JsonNode): Any =
    when (f["type"].asText()) {
        "DateTime" -> {
            val instant = Instant.parse(v.asText())
            val native = f["native"]?.takeIf { !it.isNull }?.get(0)?.asText()
            when (native) {
                "Date" -> LocalDate.ofInstant(instant, ZoneOffset.UTC)
                "Time", "Timetz" -> LocalTime.ofInstant(instant, ZoneOffset.UTC)
                else -> LocalDateTime.ofInstant(instant, ZoneOffset.UTC)
            }
        }
        "Int" -> v.asInt()
        "BigInt" -> v.asText().toLong()
        "Float" -> v.asDouble()
        "Decimal" -> BigDecimal(v.asText())
        "Boolean" -> v.asBoolean()
        "Bytes" -> Base64.getDecoder().decode(v.asText())
        "Json" -> mapper.writeValueAsString(v)
        else -> if (v.isTextual) v.asText() else mapper.writeValueAsString(v)
    }

fun normal(v: Any?): JsonNode =
    when (v) {
        null -> nodes.nullNode()
        is EntityID<*> -> normal(v.value)
        is Instant -> nodes.textNode(iso.format(v.atOffset(ZoneOffset.UTC)))
        is LocalDateTime -> nodes.textNode(iso.format(v.atOffset(ZoneOffset.UTC)))
        is OffsetDateTime -> nodes.textNode(iso.format(v.withOffsetSameInstant(ZoneOffset.UTC)))
        is LocalDate -> nodes.textNode("${v}T00:00:00.000Z")
        is LocalTime -> nodes.textNode("1970-01-01T${v.withNano(v.nano / 1_000_000 * 1_000_000).format(DateTimeFormatter.ofPattern("HH:mm:ss.SSS"))}Z")
        is BigDecimal -> nodes.textNode(v.stripTrailingZeros().toPlainString())
        is Long -> if (kotlin.math.abs(v) > 9007199254740992L) nodes.textNode(v.toString()) else nodes.numberNode(v)
        is Int -> nodes.numberNode(v)
        is Short -> nodes.numberNode(v)
        is Double -> nodes.numberNode(v)
        is Float -> nodes.numberNode(v)
        is Boolean -> nodes.booleanNode(v)
        is ByteArray -> nodes.textNode(Base64.getEncoder().encodeToString(v))
        is UUID -> nodes.textNode(v.toString())
        is Enum<*> -> nodes.textNode(enumDbName(v))
        is JsonNode -> v
        is String -> if ((v.startsWith("{") || v.startsWith("[")) && runCatching { mapper.readTree(v) }.isSuccess) mapper.readTree(v) else nodes.textNode(v)
        is Iterable<*> -> nodes.arrayNode().also { a -> v.forEach { a.add(normal(it)) } }
        is Array<*> -> nodes.arrayNode().also { a -> v.forEach { a.add(normal(it)) } }
        is Map<*, *> -> mapper.valueToTree(v)
        else -> runCatching { mapper.valueToTree<JsonNode>(v) }.getOrElse { nodes.textNode(v.toString()) }
    }

// The database value of an enum: a `dbName`/`value` property when the generated enum has one.
fun enumDbName(v: Enum<*>): String {
    for (name in listOf("dbName", "value", "sqlName", "pgName")) {
        val p = v::class.memberProperties.firstOrNull { it.name == name } ?: continue
        val got = runCatching { p.getter.call(v) }.getOrNull()
        if (got is String) return got
    }
    return v.name
}

fun message(e: Throwable): String {
    var t = e
    while (t.cause != null && t.cause !== t) t = t.cause!!
    val text = if (t === e) "${e::class.simpleName}: ${e.message}" else "${t::class.simpleName}: ${t.message}"
    return text.take(600)
}

fun error(e: Throwable): ObjectNode = nodes.objectNode().put("ok", false).put("error", message(e))
