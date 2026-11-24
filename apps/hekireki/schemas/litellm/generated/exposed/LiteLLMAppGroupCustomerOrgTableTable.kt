package litellm.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

object LiteLLMAppGroupCustomerOrgTableTable : CompositeIdTable("\"LiteLLM_AppGroupCustomerOrgTable\"") {
    val appGroupId =
        reference(
            "app_group_id",
            LiteLLMAppGroupTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppGroupCustomerOrgTable_app_group_id_fkey\"",
        )
    val customerOrgId =
        reference(
            "customer_org_id",
            LiteLLMCustomerOrganizationTableTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppGroupCustomerOrgTable_customer_org_id_fkey\"",
        )
    val budgetId =
        optReference(
            "budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppGroupCustomerOrgTable_budget_id_fkey\"",
        )
    val spend = double("spend").default(0.0)
    val spendResetWindow = text("spend_reset_window").nullable()
    override val primaryKey = PrimaryKey(appGroupId, customerOrgId, name = "\"LiteLLM_AppGroupCustomerOrgTable_pkey\"")

    init {
        addIdColumn(appGroupId)
        addIdColumn(customerOrgId)
        index("\"LiteLLM_AppGroupCustomerOrgTable_customer_org_id_idx\"", false, customerOrgId)
    }
}
