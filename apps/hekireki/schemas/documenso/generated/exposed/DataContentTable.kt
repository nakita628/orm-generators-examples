package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * Generic stored data (e.g. an uploaded image) which other models reference.
 * The kind of data and its kind specific details live in `metadata`.
 * @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
 */
object DataContentTable : IdTable<String>("\"DataContent\"") {
    override val id = text("id").entityId()

    /** How `data` is stored, e.g. an S3 path or inline bytes. */
    val type = pgEnum("type", "\"DocumentDataType\"", DocumentDataType.entries, DocumentDataType::dbName)
    val data = text("data")

    /** [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema) */
    val metadata = jsonb<String>("metadata", { it }, { it })
    override val primaryKey = PrimaryKey(id, name = "\"DataContent_pkey\"")
}
