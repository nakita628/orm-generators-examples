package litellm.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

object LiteLLMAppCustomerOrgTableTable : CompositeIdTable("\"LiteLLM_AppCustomerOrgTable\"") {
    val appId =
        reference(
            "app_id",
            LiteLLMAppTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_app_id_fkey\"",
        )
    val customerOrgId =
        reference(
            "customer_org_id",
            LiteLLMCustomerOrganizationTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_customer_org_id_fkey\"",
        )
    val spend = double("spend").default(0.0)
    val spendResetWindow = text("spend_reset_window").nullable()
    val tierBudgetId =
        optReference(
            "tier_budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_tier_budget_id_fkey\"",
        )
    val budgetId =
        optReference(
            "budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_budget_id_fkey\"",
        )
    val budgetOverrideId =
        optReference(
            "budget_override_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_budget_override_id_fkey\"",
        )
    val budgetOverrideExpiresAt = pgTimestamp("budget_override_expires_at", 3).nullable()
    val customerMemberBudgetId =
        optReference(
            "customer_member_budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppCustomerOrgTable_customer_member_budget_id_fkey\"",
        )
    override val primaryKey = PrimaryKey(appId, customerOrgId, name = "\"LiteLLM_AppCustomerOrgTable_pkey\"")

    init {
        addIdColumn(appId)
        addIdColumn(customerOrgId)
    }
}
