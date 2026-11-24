package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * The one open stop an agent may have: the api takes this row before it sends a stop's request,
 * so two requests at once cannot both open one. A later request takes it over once the stop it
 * names has closed.
 * `sentAt` is set once the request command was sent: from then on the claim passes on only after
 * its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
 * `decisionTakenAt` is set while that decision is in flight and cleared once its command was
 * sent; a decision left in flight is taken over once abandoned.
 */
object AgentStopClaimTable : CompositeIdTable("\"AgentStopClaim\"") {
    val organizationId = text("organizationId").entityId()
    val agentId = text("agentId").entityId()
    val stopId = text("stopId")
    val claimedAt = pgTimestamp("claimedAt", 3)
    val sentAt = pgTimestamp("sentAt", 3).nullable()
    val decision = text("decision").nullable()
    val decisionTakenAt = pgTimestamp("decisionTakenAt", 3).nullable()
    override val primaryKey = PrimaryKey(organizationId, agentId, name = "\"AgentStopClaim_pkey\"")
}
