package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/** A project's ISO week, folded from its owners' events. The id is the fold's aggregate key. */
object DigestProjectWeekTable : IdTable<String>("\"DigestProjectWeek\"") {
    override val id = text("id").entityId()
    val projectId = text("projectId")
    val organizationId = text("organizationId")
    val isoWeek = text("isoWeek")

    /** The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps. */
    val state = jsonb<String>("state", { it }, { it })
    val occurredAt = pgTimestamp("occurredAt", 3)
    val lastEventId = text("lastEventId")
    val acceptedAt = pgTimestamp("acceptedAt", 3)
    val projectionVersion = text("projectionVersion")
    val createdAt = pgTimestamp("createdAt", 3)
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"DigestProjectWeek_pkey\"")

    init {
        index("\"DigestProjectWeek_organizationId_isoWeek_idx\"", false, organizationId, isoWeek)
        index("\"DigestProjectWeek_isoWeek_idx\"", false, isoWeek)
    }
}
