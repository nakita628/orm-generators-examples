package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZFieldMetaNotOptionalSchema } from '@documenso/lib/types/field-meta';"]) */
class FieldEntity(
    id: EntityID<Int>,
) : Entity<Int>(id) {
    companion object : EntityClass<Int, FieldEntity>(FieldTable)

    var secondaryId by FieldTable.secondaryId
    var envelopeId by FieldTable.envelopeId
    var envelopeItemId by FieldTable.envelopeItemId
    var recipientId by FieldTable.recipientId
    var type by FieldTable.type

    /** @zod.number.describe("The page number of the field on the document. Starts from 1.") */
    var page by FieldTable.page
    var positionX by FieldTable.positionX
    var positionY by FieldTable.positionY
    var width by FieldTable.width
    var height by FieldTable.height
    var customText by FieldTable.customText
    var inserted by FieldTable.inserted
    var envelopeItem by EnvelopeItemEntity referencedOn FieldTable.envelopeItemId
    var envelope by EnvelopeEntity referencedOn FieldTable.envelopeId
    var recipient by RecipientEntity referencedOn FieldTable.recipientId
    val signature by SignatureEntity optionalBackReferencedOn SignatureTable.fieldId

    /** [FieldMeta] @zod.custom.use(ZFieldMetaNotOptionalSchema) */
    var fieldMeta by FieldTable.fieldMeta
}
