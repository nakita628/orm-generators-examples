package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A saved layout for a generated PDF, kept per team workspace under a name:
 * which columns, rows and header lines print, and in what order the rows
 * come. A row is a layout, not a generated file; the person generating a PDF
 * picks a template by name.
 */
class PdfTemplateEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PdfTemplateEntity>(PdfTemplateTable)

    var name by PdfTemplateTable.name
    var kind by PdfTemplateTable.kind

    /**
     * The layout's switches. Read through `parseBookingTemplateConfig`,
     * which fills any key a stored row lacks from the Checklist template's
     * values.
     */
    var config by PdfTemplateTable.config

    /**
     * The template the PDF dialog offers first. One per workspace and kind,
     * enforced by a partial unique index in raw SQL (migration
     * 20261103120000_add_pdf_templates). Names are also unique per workspace
     * and kind ignoring case, through an expression index in the same
     * migration.
     */
    var isDefault by PdfTemplateTable.isDefault
    var organization by OrganizationEntity referencedOn PdfTemplateTable.organizationId
    var organizationId by PdfTemplateTable.organizationId
    var createdAt by PdfTemplateTable.createdAt
    var updatedAt by PdfTemplateTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PdfTemplateTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
