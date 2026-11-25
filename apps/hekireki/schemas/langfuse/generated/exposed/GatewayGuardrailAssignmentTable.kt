package langfuse.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object GatewayGuardrailAssignmentTable : IdTable<String>("gateway_guardrail_assignments") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val guardrailId =
        reference(
            "guardrail_id",
            GatewayGuardrailTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrail_assignments_guardrail_id_fkey",
        )
    val organizationId =
        reference(
            "organization_id",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrail_assignments_organization_id_fkey",
        )
    val scope =
        pgEnum(
            "scope",
            "\"GatewayGuardrailAssignmentScope\"",
            GatewayGuardrailAssignmentScope.entries,
            GatewayGuardrailAssignmentScope::dbName,
        )
    val apiKeyId =
        optReference(
            "api_key_id",
            ApiKeyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrail_assignments_api_key_id_fkey",
        )
    val userId =
        optReference(
            "user_id",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_guardrail_assignments_user_id_fkey",
        )
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "gateway_guardrail_assignments_pkey")

    init {
        index("gateway_guardrail_assignments_organization_id_scope_idx", false, organizationId, scope)
        index("gateway_guardrail_assignments_api_key_id_idx", false, apiKeyId)
        index("gateway_guardrail_assignments_user_id_organization_id_idx", false, userId, organizationId)
    }
}
