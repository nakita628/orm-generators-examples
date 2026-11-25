package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/** A person's ISO week of coding-agent activity in one organization, folded from the agents' facts. */
object DigestUserWeekTable : IdTable<String>("\"DigestUserWeek\"") {
    override val id = text("id").entityId()
    val userId = text("userId")
    val organizationId = text("organizationId")
    val isoWeek = text("isoWeek")

    /** Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins. */
    val state = jsonb<String>("state", { it }, { it })
    val occurredAt = pgTimestamp("occurredAt", 3)
    val lastEventId = text("lastEventId")
    val acceptedAt = pgTimestamp("acceptedAt", 3)
    val projectionVersion = text("projectionVersion")
    val createdAt = pgTimestamp("createdAt", 3)
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"DigestUserWeek_pkey\"")

    init {
        index("\"DigestUserWeek_organizationId_isoWeek_idx\"", false, organizationId, isoWeek)
        index("\"DigestUserWeek_isoWeek_idx\"", false, isoWeek)
        index("\"DigestUserWeek_userId_idx\"", false, userId)
    }
}
