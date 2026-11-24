package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"]) */
class AdminRuleMatchEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AdminRuleMatchEntity>(AdminRuleMatchTable)

    var createdAt by AdminRuleMatchTable.createdAt
    var mode by AdminRuleMatchTable.mode

    /** [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema) */
    var trigger by AdminRuleMatchTable.trigger
    var organisationId by AdminRuleMatchTable.organisationId
    var envelopeId by AdminRuleMatchTable.envelopeId

    /** [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema) */
    var eventSnapshot by AdminRuleMatchTable.eventSnapshot

    /** [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema) */
    var actionsRun by AdminRuleMatchTable.actionsRun
    var ruleId by AdminRuleMatchTable.ruleId
    var rule by AdminRuleEntity referencedOn AdminRuleMatchTable.ruleId
    var actorUserId by AdminRuleMatchTable.actorUserId
    var actorUser by UserEntity optionalReferencedOn AdminRuleMatchTable.actorUserId
}
