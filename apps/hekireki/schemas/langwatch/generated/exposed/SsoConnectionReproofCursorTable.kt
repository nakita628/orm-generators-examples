package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Operational round-robin position for ADR-123's domain re-proof sweep.
 *
 * This is deliberately not a column on `SsoConnection`: that row is a
 * whole-row event projection and its fold is its only writer, while a
 * healthy re-read states no fact. Keeping the scheduling cursor beside the
 * projection preserves both rules and still lets a fixed-size sweep revisit
 * every connection, including one whose proof has started wavering.
 */
object SsoConnectionReproofCursorTable : IdTable<String>("\"SsoConnectionReproofCursor\"") {
    override val id = text("connectionId").entityId()
    val lastReproofAt = pgTimestamp("lastReproofAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"SsoConnectionReproofCursor_pkey\"")

    init {
        index("\"SsoConnectionReproofCursor_lastReproofAt_idx\"", false, lastReproofAt)
    }
}
