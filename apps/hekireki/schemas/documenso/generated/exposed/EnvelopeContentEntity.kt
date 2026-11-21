package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZEnvelopeContentMetaSchema } from '@documenso/lib/types/envelope-content-meta';"]) */
class EnvelopeContentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, EnvelopeContentEntity>(EnvelopeContentTable)

    var envelopeId by EnvelopeContentTable.envelopeId
    var envelopeItemId by EnvelopeContentTable.envelopeItemId
    var envelopeItem by EnvelopeItemEntity referencedOn EnvelopeContentTable.envelopeItemId
    var envelope by EnvelopeEntity referencedOn EnvelopeContentTable.envelopeId

    /** [EnvelopeContentMeta] @zod.custom.use(ZEnvelopeContentMetaSchema) */
    var contentMeta by EnvelopeContentTable.contentMeta

    /**
     * Any number of contents may point at the same data content, e.g. a
     * document created from a template shares the template's images.
     */
    var dataContentId by EnvelopeContentTable.dataContentId
    var dataContent by DataContentEntity optionalReferencedOn EnvelopeContentTable.dataContentId
}
