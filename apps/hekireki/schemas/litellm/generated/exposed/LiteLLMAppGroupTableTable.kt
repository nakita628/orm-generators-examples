package litellm.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit
import java.util.UUID

object LiteLLMAppGroupTableTable : IdTable<String>("\"LiteLLM_AppGroupTable\"") {
    override val id = text("app_group_id").clientDefault { UUID.randomUUID().toString() }.entityId()
    val name = text("name").uniqueIndex("\"LiteLLM_AppGroupTable_name_key\"")
    val description = text("description").nullable()
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()
    val budgetId =
        optReference(
            "budget_id",
            LiteLLMBudgetTableTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"LiteLLM_AppGroupTable_budget_id_fkey\"",
        )
    val spend = double("spend").default(0.0)
    val spendResetWindow = text("spend_reset_window").nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdBy = text("created_by")
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedBy = text("updated_by")
    override val primaryKey = PrimaryKey(id, name = "\"LiteLLM_AppGroupTable_pkey\"")
}
