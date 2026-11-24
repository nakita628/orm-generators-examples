package documenso.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"]) */
object AdminRuleMatchTable : IdTable<String>("\"AdminRuleMatch\"") {
    override val id = text("id").entityId()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val mode = pgEnum("mode", "\"AdminRuleMode\"", AdminRuleMode.entries, AdminRuleMode::dbName)

    /** [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema) */
    val trigger = text("trigger")
    val organisationId = text("organisationId").nullable()
    val envelopeId = text("envelopeId").nullable()

    /** [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema) */
    val eventSnapshot = jsonb<String>("eventSnapshot", { it }, { it })

    /** [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema) */
    val actionsRun = jsonb<String>("actionsRun", { it }, { it })
    val ruleId =
        reference(
            "ruleId",
            AdminRuleTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminRuleMatch_ruleId_fkey\"",
        )
    val actorUserId =
        optReference(
            "actorUserId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AdminRuleMatch_actorUserId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"AdminRuleMatch_pkey\"")

    init {
        index("\"AdminRuleMatch_ruleId_createdAt_idx\"", false, ruleId, createdAt)
        index("\"AdminRuleMatch_createdAt_idx\"", false, createdAt)
    }
}
