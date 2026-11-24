package litellm.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

object LiteLLMAppCustomerMemberTableTable : CompositeIdTable("\"LiteLLM_AppCustomerMemberTable\"") {
    val appId =
        reference(
            "app_id",
            LiteLLMAppTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerMemberTable_app_id_fkey\"",
        )
    val customerMemberId =
        reference(
            "customer_member_id",
            LiteLLMCustomerMemberTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerMemberTable_customer_member_id_fkey\"",
        )
    val spend = double("spend").default(0.0)
    val spendResetWindow = text("spend_reset_window").nullable()
    val budgetOverrideId =
        optReference(
            "budget_override_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerMemberTable_budget_override_id_fkey\"",
        )
    val budgetId =
        optReference(
            "budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerMemberTable_budget_id_fkey\"",
        )
    val budgetOverrideExpiresAt = pgTimestamp("budget_override_expires_at", 3).nullable()
    override val primaryKey = PrimaryKey(appId, customerMemberId, name = "\"LiteLLM_AppCustomerMemberTable_pkey\"")

    init {
        addIdColumn(appId)
        addIdColumn(customerMemberId)
    }
}
