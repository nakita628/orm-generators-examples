package litellm.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class LiteLLMAppGroupTableEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LiteLLMAppGroupTableEntity>(LiteLLMAppGroupTableTable)

    var name by LiteLLMAppGroupTableTable.name
    var description by LiteLLMAppGroupTableTable.description
    var metadata by LiteLLMAppGroupTableTable.metadata
    var budgetId by LiteLLMAppGroupTableTable.budgetId
    var budget by LiteLLMBudgetTableEntity optionalReferencedOn LiteLLMAppGroupTableTable.budgetId
    var spend by LiteLLMAppGroupTableTable.spend
    var spendResetWindow by LiteLLMAppGroupTableTable.spendResetWindow
    val apps by LiteLLMAppTableEntity optionalReferrersOn LiteLLMAppTableTable.appGroupId
    val customerOrganizations by LiteLLMAppGroupCustomerOrgTableEntity referrersOn LiteLLMAppGroupCustomerOrgTableTable.appGroupId
    var createdAt by LiteLLMAppGroupTableTable.createdAt
    var createdBy by LiteLLMAppGroupTableTable.createdBy
    var updatedAt by LiteLLMAppGroupTableTable.updatedAt
    var updatedBy by LiteLLMAppGroupTableTable.updatedBy

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(LiteLLMAppGroupTableTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
