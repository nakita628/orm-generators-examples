package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class LiteLLMAppGroupCustomerOrgTableEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<LiteLLMAppGroupCustomerOrgTableEntity>(LiteLLMAppGroupCustomerOrgTableTable)

    val appGroupId by LiteLLMAppGroupCustomerOrgTableTable.appGroupId
    val customerOrgId by LiteLLMAppGroupCustomerOrgTableTable.customerOrgId
    var budgetId by LiteLLMAppGroupCustomerOrgTableTable.budgetId
    var budget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppGroupCustomerOrgTableTable.budgetId
    var spend by LiteLLMAppGroupCustomerOrgTableTable.spend
    var spendResetWindow by LiteLLMAppGroupCustomerOrgTableTable.spendResetWindow
    val appGroup by LiteLLMAppGroupTableEntity referencedOn LiteLLMAppGroupCustomerOrgTableTable.appGroupId
    val customerOrg by LiteLLMCustomerOrganizationTableEntity referencedOn LiteLLMAppGroupCustomerOrgTableTable.customerOrgId
}
