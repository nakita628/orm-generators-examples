package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * The one open stop an agent may have: the api takes this row before it sends a stop's request,
 * so two requests at once cannot both open one. A later request takes it over once the stop it
 * names has closed.
 * `sentAt` is set once the request command was sent: from then on the claim passes on only after
 * its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
 * `decisionTakenAt` is set while that decision is in flight and cleared once its command was
 * sent; a decision left in flight is taken over once abandoned.
 */
class AgentStopClaimEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<AgentStopClaimEntity>(AgentStopClaimTable)

    val organizationId by AgentStopClaimTable.organizationId
    val agentId by AgentStopClaimTable.agentId
    var stopId by AgentStopClaimTable.stopId
    var claimedAt by AgentStopClaimTable.claimedAt
    var sentAt by AgentStopClaimTable.sentAt
    var decision by AgentStopClaimTable.decision
    var decisionTakenAt by AgentStopClaimTable.decisionTakenAt
}
