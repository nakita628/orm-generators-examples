package langfuse.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object GatewayGuardrailTable : IdTable<String>("gateway_guardrails") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organization_id",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrails_organization_id_fkey",
        )
    val name = text("name")
    val type =
        pgEnum("type", "\"GatewayGuardrailType\"", GatewayGuardrailType.entries, GatewayGuardrailType::dbName).default(
            GatewayGuardrailType.REGEX,
        )
    val action = pgEnum("action", "\"GatewayGuardrailAction\"", GatewayGuardrailAction.entries, GatewayGuardrailAction::dbName)
    val config = jsonb<String>("config", { it }, { it }).default("{}")
    val encryptedCredential = text("encrypted_credentials").nullable()
    val createdBy =
        optReference(
            "created_by",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrails_created_by_fkey",
        )
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "gateway_guardrails_pkey")

    init {
        uniqueIndex("gateway_guardrails_organization_id_name_key", organizationId, name)
    }
}
