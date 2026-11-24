package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

class LiteLLMAppCustomerMemberTableEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<LiteLLMAppCustomerMemberTableEntity>(LiteLLMAppCustomerMemberTableTable)

    val appId by LiteLLMAppCustomerMemberTableTable.appId
    val customerMemberId by LiteLLMAppCustomerMemberTableTable.customerMemberId
    var spend by LiteLLMAppCustomerMemberTableTable.spend
    var spendResetWindow by LiteLLMAppCustomerMemberTableTable.spendResetWindow
    var budgetOverrideId by LiteLLMAppCustomerMemberTableTable.budgetOverrideId
    var budgetId by LiteLLMAppCustomerMemberTableTable.budgetId
    var budget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerMemberTableTable.budgetId
    var budgetOverrideExpiresAt by LiteLLMAppCustomerMemberTableTable.budgetOverrideExpiresAt
    val app by LiteLLMAppTableEntity referencedOn LiteLLMAppCustomerMemberTableTable.appId
    val customerMember by LiteLLMCustomerMemberTableEntity referencedOn LiteLLMAppCustomerMemberTableTable.customerMemberId
    var budgetOverride by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppCustomerMemberTableTable.budgetOverrideId
}
