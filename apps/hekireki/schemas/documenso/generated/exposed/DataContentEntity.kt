package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Generic stored data (e.g. an uploaded image) which other models reference.
 * The kind of data and its kind specific details live in `metadata`.
 * @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
 */
class DataContentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DataContentEntity>(DataContentTable)

    /** How `data` is stored, e.g. an S3 path or inline bytes. */
    var type by DataContentTable.type
    var data by DataContentTable.data

    /** [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema) */
    var metadata by DataContentTable.metadata
    val envelopeContents by EnvelopeContentEntity optionalReferrersOn EnvelopeContentTable.dataContentId
}
