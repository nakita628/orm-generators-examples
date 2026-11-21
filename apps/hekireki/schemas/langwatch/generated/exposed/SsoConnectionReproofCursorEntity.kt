package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Operational round-robin position for ADR-123's domain re-proof sweep.
 *
 * This is deliberately not a column on `SsoConnection`: that row is a
 * whole-row event projection and its fold is its only writer, while a
 * healthy re-read states no fact. Keeping the scheduling cursor beside the
 * projection preserves both rules and still lets a fixed-size sweep revisit
 * every connection, including one whose proof has started wavering.
 */
class SsoConnectionReproofCursorEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SsoConnectionReproofCursorEntity>(SsoConnectionReproofCursorTable)

    var lastReproofAt by SsoConnectionReproofCursorTable.lastReproofAt
}
