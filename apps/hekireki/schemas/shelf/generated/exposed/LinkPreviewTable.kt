package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * What Shelf knows about one web link used in a workspace's custom fields,
 * keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
 * the link preview job (fetch columns) and people (customTitle).
 */
object LinkPreviewTable : IdTable<String>("\"LinkPreview\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /**
     * `LinkKey` from @shelf/links, never a raw field value. At most 2048
     * characters (CHECK constraint in the migration).
     */
    val url = text("url")
    val pageTitle = text("pageTitle").nullable()

    /** Null until the first fetch finishes. */
    val outcome = pgEnum("outcome", "\"LinkPreviewOutcome\"", LinkPreviewOutcome.entries, LinkPreviewOutcome::dbName).nullable()
    val attempts = integer("attempts").default(0)
    val fetchedAt = pgTimestamp("fetchedAt", 3).nullable()

    /**
     * When the link may be fetched next; null means never again. A claim moves
     * it a few minutes ahead, which is also the lease on a fetch in flight.
     */
    val nextFetchAt =
        pgTimestamp("nextFetchAt", 3).nullable().databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** The name a person gave the link in this workspace; wins over pageTitle. */
    val customTitle = text("customTitle").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LinkPreview_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"LinkPreview_pkey\"")

    init {
        uniqueIndex("\"LinkPreview_organizationId_url_key\"", organizationId, url)
    }
}
