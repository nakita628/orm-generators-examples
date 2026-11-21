package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class SubscriptionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SubscriptionEntity>(SubscriptionTable)

    var organizationId by SubscriptionTable.organizationId
    var plan by SubscriptionTable.plan
    var stripeSubscriptionId by SubscriptionTable.stripeSubscriptionId
    var status by SubscriptionTable.status
    var startDate by SubscriptionTable.startDate
    var endDate by SubscriptionTable.endDate
    var lastPaymentFailedDate by SubscriptionTable.lastPaymentFailedDate
    var createdAt by SubscriptionTable.createdAt
    var updatedAt by SubscriptionTable.updatedAt
    var maxMembers by SubscriptionTable.maxMembers
    var maxProjects by SubscriptionTable.maxProjects
    var maxMessagesPerMonth by SubscriptionTable.maxMessagesPerMonth
    var maxWorkflows by SubscriptionTable.maxWorkflows
    var maxRetentionDays by SubscriptionTable.maxRetentionDays
    var maxMembersLite by SubscriptionTable.maxMembersLite

    /** Developer seats the organization pays for; null falls to the plan's own cap. */
    var maxMembersDeveloper by SubscriptionTable.maxMembersDeveloper
    var maxTeams by SubscriptionTable.maxTeams
    var maxPrompts by SubscriptionTable.maxPrompts
    var maxEvaluators by SubscriptionTable.maxEvaluators
    var maxScenarios by SubscriptionTable.maxScenarios
    var maxAgents by SubscriptionTable.maxAgents
    var maxExperiments by SubscriptionTable.maxExperiments
    var maxOnlineEvaluations by SubscriptionTable.maxOnlineEvaluations
    var maxDatasets by SubscriptionTable.maxDatasets
    var maxDashboards by SubscriptionTable.maxDashboards
    var maxCustomGraphs by SubscriptionTable.maxCustomGraphs
    var maxAutomations by SubscriptionTable.maxAutomations

    /** Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil. */
    var seatsKept by SubscriptionTable.seatsKept
    var seatsKeptUntil by SubscriptionTable.seatsKeptUntil

    /** The seat quantity the last finished seat change set the seat line to; provider events never write it. */
    var seatLineQuantity by SubscriptionTable.seatLineQuantity
    var seatLineQuantityAt by SubscriptionTable.seatLineQuantityAt

    /**
     * A Stripe seat line write recorded before it is sent and cleared once it is known to have
     * finished under the lease; one left behind is settled by the seat line repair.
     * Every seat quantity written while this record was pending.
     */
    var seatLineWriteQuantities by SubscriptionTable.seatLineWriteQuantities

    /** The first write still pending; a later holder's write keeps it. */
    var seatLineWriteSince by SubscriptionTable.seatLineWriteSince

    /** The latest write sent; the repair waits a grace window after it. */
    var seatLineWriteAt by SubscriptionTable.seatLineWriteAt
    var seatLineWriteAttempts by SubscriptionTable.seatLineWriteAttempts
    var seatLineWriteFlaggedAt by SubscriptionTable.seatLineWriteFlaggedAt

    /** The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled. */
    var seatRepairQuantity by SubscriptionTable.seatRepairQuantity

    /** The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it. */
    var seatRepairLiftedQuantity by SubscriptionTable.seatRepairLiftedQuantity
    var seatRepairSince by SubscriptionTable.seatRepairSince
    var seatRepairAttempts by SubscriptionTable.seatRepairAttempts

    /** Set once the seat line repair has failed often enough that Ops should look at it. */
    var seatRepairFlaggedAt by SubscriptionTable.seatRepairFlaggedAt

    /** The lease a seat change or seat line repair holds while it reads and moves the seat line. */
    var seatLineLeaseToken by SubscriptionTable.seatLineLeaseToken
    var seatLineLeaseUntil by SubscriptionTable.seatLineLeaseUntil

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SubscriptionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
