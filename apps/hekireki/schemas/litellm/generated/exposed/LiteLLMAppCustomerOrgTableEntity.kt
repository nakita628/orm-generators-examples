package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class LiteLLMAppCustomerOrgTableEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<LiteLLMAppCustomerOrgTableEntity>(LiteLLMAppCustomerOrgTableTable)

    val appId by LiteLLMAppCustomerOrgTableTable.appId
    val customerOrgId by LiteLLMAppCustomerOrgTableTable.customerOrgId
    var spend by LiteLLMAppCustomerOrgTableTable.spend
    var spendResetWindow by LiteLLMAppCustomerOrgTableTable.spendResetWindow
    var tierBudgetId by LiteLLMAppCustomerOrgTableTable.tierBudgetId
    var budgetId by LiteLLMAppCustomerOrgTableTable.budgetId
    var budget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerOrgTableTable.budgetId
    var budgetOverrideId by LiteLLMAppCustomerOrgTableTable.budgetOverrideId
    var budgetOverrideExpiresAt by LiteLLMAppCustomerOrgTableTable.budgetOverrideExpiresAt
    var customerMemberBudgetId by LiteLLMAppCustomerOrgTableTable.customerMemberBudgetId
    val app by LiteLLMAppTableEntity referencedOn LiteLLMAppCustomerOrgTableTable.appId
    val customerOrg by LiteLLMCustomerOrganizationTableEntity referencedOn LiteLLMAppCustomerOrgTableTable.customerOrgId
    var tierBudget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerOrgTableTable.tierBudgetId
    var budgetOverride by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerOrgTableTable.budgetOverrideId
    var customerMemberBudget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerOrgTableTable.customerMemberBudgetId
}
