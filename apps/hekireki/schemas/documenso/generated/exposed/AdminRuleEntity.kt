package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"]) */
class AdminRuleEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AdminRuleEntity>(AdminRuleTable)

    var createdAt by AdminRuleTable.createdAt
    var updatedAt by AdminRuleTable.updatedAt
    var name by AdminRuleTable.name
    var description by AdminRuleTable.description
    var mode by AdminRuleTable.mode

    /** [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema) */
    var trigger by AdminRuleTable.trigger

    /** [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema) */
    var definition by AdminRuleTable.definition
    var priority by AdminRuleTable.priority
    var stopOnMatch by AdminRuleTable.stopOnMatch

    /**
     * Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
     * organisation
     */
    var exemptOrganisationIds by AdminRuleTable.exemptOrganisationIds
    val matches by AdminRuleMatchEntity referrersOn AdminRuleMatchTable.ruleId
    val adminTags by AdminTagEntity optionalReferrersOn AdminTagTable.adminRuleId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AdminRuleTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
