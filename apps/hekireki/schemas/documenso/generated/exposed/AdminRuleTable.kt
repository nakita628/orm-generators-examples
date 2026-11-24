package documenso.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"]) */
object AdminRuleTable : IdTable<String>("\"AdminRule\"") {
    override val id = text("id").entityId()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val description = text("description").nullable()
    val mode = pgEnum("mode", "\"AdminRuleMode\"", AdminRuleMode.entries, AdminRuleMode::dbName)

    /** [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema) */
    val trigger = text("trigger")

    /** [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema) */
    val definition = jsonb<String>("definition", { it }, { it })
    val priority = integer("priority")
    val stopOnMatch = bool("stopOnMatch")

    /**
     * Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
     * organisation
     */
    val exemptOrganisationIds = array("exemptOrganisationIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    override val primaryKey = PrimaryKey(id, name = "\"AdminRule_pkey\"")

    init {
        index("\"AdminRule_trigger_mode_priority_idx\"", false, trigger, mode, priority)
    }
}
