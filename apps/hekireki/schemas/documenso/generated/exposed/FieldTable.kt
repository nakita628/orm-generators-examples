package documenso.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.math.BigDecimal

/** @zod.import(["import { ZFieldMetaNotOptionalSchema } from '@documenso/lib/types/field-meta';"]) */
object FieldTable : IdTable<Int>("\"Field\"") {
    override val id = integer("id").autoIncrement().entityId()
    val secondaryId = text("secondaryId").clientDefault { CUID.randomCUID1().toString() }.uniqueIndex("\"Field_secondaryId_key\"")
    val envelopeId =
        reference(
            "envelopeId",
            EnvelopeTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Field_envelopeId_fkey\"",
        )
    val envelopeItemId =
        reference(
            "envelopeItemId",
            EnvelopeItemTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Field_envelopeItemId_fkey\"",
        )
    val recipientId =
        reference(
            "recipientId",
            RecipientTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Field_recipientId_fkey\"",
        )
    val type = pgEnum("type", "\"FieldType\"", FieldType.entries, FieldType::dbName)

    /** @zod.number.describe("The page number of the field on the document. Starts from 1.") */
    val page = integer("page")
    val positionX = decimal("positionX", 65, 30).default(BigDecimal("0"))
    val positionY = decimal("positionY", 65, 30).default(BigDecimal("0"))
    val width = decimal("width", 65, 30).default(BigDecimal("-1"))
    val height = decimal("height", 65, 30).default(BigDecimal("-1"))
    val customText = text("customText")
    val inserted = bool("inserted")

    /** [FieldMeta] @zod.custom.use(ZFieldMetaNotOptionalSchema) */
    val fieldMeta = jsonb<String>("fieldMeta", { it }, { it }).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Field_pkey\"")

    init {
        index("\"Field_envelopeId_idx\"", false, envelopeId)
        index("\"Field_envelopeItemId_idx\"", false, envelopeItemId)
        index("\"Field_recipientId_idx\"", false, recipientId)
    }
}
