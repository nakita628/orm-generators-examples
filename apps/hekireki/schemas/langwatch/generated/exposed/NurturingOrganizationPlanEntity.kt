package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
 * copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
 * fact only when it changes this row.
 */
class NurturingOrganizationPlanEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingOrganizationPlanEntity>(NurturingOrganizationPlanTable)

    var plan by NurturingOrganizationPlanTable.plan
    var occurredAt by NurturingOrganizationPlanTable.occurredAt
    var fromCatchUp by NurturingOrganizationPlanTable.fromCatchUp

    /** The event ids folded in, so a redelivery folds once. */
    var appliedEventIds by NurturingOrganizationPlanTable.appliedEventIds

    /** The fact that last changed the row; only it may tell Customer.io. */
    var winnerEventId by NurturingOrganizationPlanTable.winnerEventId
}
