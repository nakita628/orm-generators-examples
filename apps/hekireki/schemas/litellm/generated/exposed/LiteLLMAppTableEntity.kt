package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class LiteLLMAppTableEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LiteLLMAppTableEntity>(LiteLLMAppTableTable)

    var appGroupId by LiteLLMAppTableTable.appGroupId
    var appGroup by LiteLLMAppGroupTableEntity optionalReferencedOn LiteLLMAppTableTable.appGroupId
    var resourceType by LiteLLMAppTableTable.resourceType
    var resourceId by LiteLLMAppTableTable.resourceId
    var alias by LiteLLMAppTableTable.alias
    var spend by LiteLLMAppTableTable.spend
    var spendResetWindow by LiteLLMAppTableTable.spendResetWindow
    var budgetId by LiteLLMAppTableTable.budgetId
    var budget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppTableTable.budgetId
    var defaultTierBudgetId by LiteLLMAppTableTable.defaultTierBudgetId
    var defaultTierBudget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppTableTable.defaultTierBudgetId
    var defaultCustomerMemberBudgetId by LiteLLMAppTableTable.defaultCustomerMemberBudgetId
    var defaultCustomerMemberBudget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppTableTable.defaultCustomerMemberBudgetId
    var autoCreateCustomers by LiteLLMAppTableTable.autoCreateCustomers
    var customerMargin by LiteLLMAppTableTable.customerMargin
    val customerOrganizations by LiteLLMAppCustomerOrgTableEntity referrersOn LiteLLMAppCustomerOrgTableTable.appId
    val customerMembers by LiteLLMAppCustomerMemberTableEntity referrersOn LiteLLMAppCustomerMemberTableTable.appId
    var createdAt by LiteLLMAppTableTable.createdAt
    var createdBy by LiteLLMAppTableTable.createdBy
    var updatedAt by LiteLLMAppTableTable.updatedAt
    var updatedBy by LiteLLMAppTableTable.updatedBy

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(LiteLLMAppTableTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
