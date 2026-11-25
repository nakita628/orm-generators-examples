package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
 * copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
 * fact only when it changes this row.
 */
object NurturingOrganizationPlanTable : IdTable<String>("\"NurturingOrganizationPlan\"") {
    override val id = text("organizationId").entityId()
    val plan = text("plan").nullable()
    val occurredAt = pgTimestamp("occurredAt", 3)
    val fromCatchUp = bool("fromCatchUp")

    /** The event ids folded in, so a redelivery folds once. */
    val appliedEventIds = array("appliedEventIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** The fact that last changed the row; only it may tell Customer.io. */
    val winnerEventId = text("winnerEventId").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"NurturingOrganizationPlan_pkey\"")
}
