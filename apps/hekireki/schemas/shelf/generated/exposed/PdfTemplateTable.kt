package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A saved layout for a generated PDF, kept per team workspace under a name:
 * which columns, rows and header lines print, and in what order the rows
 * come. A row is a layout, not a generated file; the person generating a PDF
 * picks a template by name.
 */
object PdfTemplateTable : IdTable<String>("\"PdfTemplate\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val kind = pgEnum("kind", "\"PdfTemplateKind\"", PdfTemplateKind.entries, PdfTemplateKind::dbName).default(PdfTemplateKind.BOOKING)

    /**
     * The layout's switches. Read through `parseBookingTemplateConfig`,
     * which fills any key a stored row lacks from the Checklist template's
     * values.
     */
    val config = jsonb<String>("config", { it }, { it }).default("{}")

    /**
     * The template the PDF dialog offers first. One per workspace and kind,
     * enforced by a partial unique index in raw SQL (migration
     * 20261103120000_add_pdf_templates). Names are also unique per workspace
     * and kind ignoring case, through an expression index in the same
     * migration.
     */
    val isDefault = bool("isDefault").default(false)
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PdfTemplate_organizationId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PdfTemplate_pkey\"")

    init {
        uniqueIndex("\"PdfTemplate_organizationId_kind_name_key\"", organizationId, kind, name)
    }
}
