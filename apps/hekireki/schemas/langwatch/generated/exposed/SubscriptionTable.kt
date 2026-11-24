package langwatch.models

import org.jetbrains.exposed.v1.core.IntegerColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object SubscriptionTable : IdTable<String>("\"Subscription\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val plan = pgEnum("plan", "\"PlanTypes\"", PlanTypes.entries, PlanTypes::dbName)
    val stripeSubscriptionId = text("stripeSubscriptionId").nullable().uniqueIndex("\"Subscription_stripeSubscriptionId_key\"")
    val status = pgEnum("status", "\"SubscriptionStatus\"", SubscriptionStatus.entries, SubscriptionStatus::dbName)
    val startDate = pgTimestamp("startDate", 3).nullable()
    val endDate = pgTimestamp("endDate", 3).nullable()
    val lastPaymentFailedDate = pgTimestamp("lastPaymentFailedDate", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val maxMembers = integer("maxMembers").nullable()
    val maxProjects = integer("maxProjects").nullable()
    val maxMessagesPerMonth = integer("maxMessagesPerMonth").nullable()
    val maxWorkflows = integer("maxWorkflows").nullable()
    val maxRetentionDays = integer("maxRetentionDays").nullable()
    val maxMembersLite = integer("maxMembersLite").nullable()

    /** Developer seats the organization pays for; null falls to the plan's own cap. */
    val maxMembersDeveloper = integer("maxMembersDeveloper").nullable()
    val maxTeams = integer("maxTeams").nullable()
    val maxPrompts = integer("maxPrompts").nullable()
    val maxEvaluators = integer("maxEvaluators").nullable()
    val maxScenarios = integer("maxScenarios").nullable()
    val maxAgents = integer("maxAgents").nullable()
    val maxExperiments = integer("maxExperiments").nullable()
    val maxOnlineEvaluations = integer("maxOnlineEvaluations").nullable()
    val maxDatasets = integer("maxDatasets").nullable()
    val maxDashboards = integer("maxDashboards").nullable()
    val maxCustomGraphs = integer("maxCustomGraphs").nullable()
    val maxAutomations = integer("maxAutomations").nullable()

    /** Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil. */
    val seatsKept = integer("seatsKept").nullable()
    val seatsKeptUntil = pgTimestamp("seatsKeptUntil", 3).nullable()

    /** The seat quantity the last finished seat change set the seat line to; provider events never write it. */
    val seatLineQuantity = integer("seatLineQuantity").nullable()
    val seatLineQuantityAt = pgTimestamp("seatLineQuantityAt", 3).nullable()

    /**
     * A Stripe seat line write recorded before it is sent and cleared once it is known to have
     * finished under the lease; one left behind is settled by the seat line repair.
     * Every seat quantity written while this record was pending.
     */
    val seatLineWriteQuantities = array("seatLineWriteQuantities", IntegerColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** The first write still pending; a later holder's write keeps it. */
    val seatLineWriteSince = pgTimestamp("seatLineWriteSince", 3).nullable()

    /** The latest write sent; the repair waits a grace window after it. */
    val seatLineWriteAt = pgTimestamp("seatLineWriteAt", 3).nullable()
    val seatLineWriteAttempts = integer("seatLineWriteAttempts").nullable()
    val seatLineWriteFlaggedAt = pgTimestamp("seatLineWriteFlaggedAt", 3).nullable()

    /** The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled. */
    val seatRepairQuantity = integer("seatRepairQuantity").nullable()

    /** The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it. */
    val seatRepairLiftedQuantity = integer("seatRepairLiftedQuantity").nullable()
    val seatRepairSince = pgTimestamp("seatRepairSince", 3).nullable()
    val seatRepairAttempts = integer("seatRepairAttempts").nullable()

    /** Set once the seat line repair has failed often enough that Ops should look at it. */
    val seatRepairFlaggedAt = pgTimestamp("seatRepairFlaggedAt", 3).nullable()

    /** The lease a seat change or seat line repair holds while it reads and moves the seat line. */
    val seatLineLeaseToken = text("seatLineLeaseToken").nullable()
    val seatLineLeaseUntil = pgTimestamp("seatLineLeaseUntil", 3).nullable()

    /**
     * An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
     * charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
     */
    val paymentRecoveryInvoiceId = text("paymentRecoveryInvoiceId").nullable()
    val paymentRecoveryStartedAt = pgTimestamp("paymentRecoveryStartedAt", 3).nullable()
    val paymentRecoveryNoticeDay = integer("paymentRecoveryNoticeDay").nullable()

    /**
     * The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
     * organization is on Free with more members than Free allows; ended once that no longer holds.
     */
    val overPlanSince = pgTimestamp("overPlanSince", 3).nullable()
    val overPlanEndedAt = pgTimestamp("overPlanEndedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Subscription_pkey\"")

    init {
        index("\"Subscription_organizationId_idx\"", false, organizationId)
    }
}
