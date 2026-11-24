package litellm.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit
import java.util.UUID

object LiteLLMAppTableTable : IdTable<String>("\"LiteLLM_AppTable\"") {
    override val id = text("app_id").clientDefault { UUID.randomUUID().toString() }.entityId()
    val appGroupId =
        optReference(
            "app_group_id",
            LiteLLMAppGroupTableTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppTable_app_group_id_fkey\"",
        )
    val resourceType = text("resource_type").nullable()
    val resourceId = text("resource_id").nullable()
    val alias = text("alias").uniqueIndex("\"LiteLLM_AppTable_alias_key\"")
    val spend = double("spend").default(0.0)
    val spendResetWindow = text("spend_reset_window").nullable()
    val budgetId =
        optReference(
            "budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppTable_budget_id_fkey\"",
        )
    val defaultTierBudgetId =
        optReference(
            "default_tier_budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppTable_default_tier_budget_id_fkey\"",
        )
    val defaultCustomerMemberBudgetId =
        optReference(
            "default_customer_member_budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppTable_default_customer_member_budget_id_fkey\"",
        )
    val autoCreateCustomers = bool("auto_create_customers").default(true)
    val customerMargin = double("customer_margin").nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdBy = text("created_by")
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedBy = text("updated_by")
    override val primaryKey = PrimaryKey(id, name = "\"LiteLLM_AppTable_pkey\"")

    init {
        index("\"LiteLLM_AppTable_app_group_id_idx\"", false, appGroupId)
        uniqueIndex("\"LiteLLM_AppTable_resource_type_resource_id_key\"", resourceType, resourceId)
    }
}
