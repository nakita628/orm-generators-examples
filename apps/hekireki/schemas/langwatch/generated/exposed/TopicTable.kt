package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * PROJECTION of the topic-clustering event stream (ADR-051): topics are
 * `topics_recorded` events; this table is written ONLY by
 * TopicModelFoldProjection (cursor: TopicModelProjection) and is
 * rebuildable by replay. Kept named `Topic` because it is the domain
 * entity every surface reads — the `*Projection` suffix is reserved for
 * cursor/read-model tables.
 */
object TopicTable : IdTable<String>("\"Topic\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")
    val name = text("name")
    val parentId = text("parentId").nullable()
    val embeddingsModel = text("embeddings_model")
    val centroid = jsonb<String>("centroid", { it }, { it })
    val p95Distance = double("p95Distance")
    val automaticallyGenerated = bool("automaticallyGenerated").default(true)

    /**
     * The `topics_recorded` event that recorded this topic (its id in the
     * event log). Null means the row predates event ownership and has not
     * been seeded onto the stream yet; the projection sets it on every row
     * it writes.
     */
    val lastEventId = text("lastEventId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Topic_pkey\"")

    init {
        index("\"Topic_parentId_idx\"", false, parentId)
        index("\"Topic_projectId_idx\"", false, projectId)
    }
}
