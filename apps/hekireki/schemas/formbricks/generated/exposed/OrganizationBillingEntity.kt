package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Stores billing and Stripe synchronization data for an organization.
 *
 * @property organizationId - Reference to the organization
 * @property stripeCustomerId - Stripe customer ID for cloud billing
 * @property limits - Plan limits used by feature and quota checks
 * @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
 * @property stripe - Stripe synchronization snapshot data
 */
class OrganizationBillingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationBillingEntity>(OrganizationBillingTable)

    val organization by OrganizationEntity referencedOn OrganizationBillingTable.id
    var stripeCustomerId by OrganizationBillingTable.stripeCustomerId

    /** [OrganizationBillingPlanLimits] */
    var limits by OrganizationBillingTable.limits
    var usageCycleAnchor by OrganizationBillingTable.usageCycleAnchor

    /** [OrganizationStripeBilling] */
    var stripe by OrganizationBillingTable.stripe
    var createdAt by OrganizationBillingTable.createdAt
    var updatedAt by OrganizationBillingTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrganizationBillingTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
