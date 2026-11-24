package langfuse.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.IntegerColumnType
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object GatewaySpendLimitTable : IdTable<String>("gateway_spend_limits") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organization_id",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_spend_limits_organization_id_fkey",
        )
    val scope = pgEnum("scope", "\"GatewaySpendLimitScope\"", GatewaySpendLimitScope.entries, GatewaySpendLimitScope::dbName)
    val apiKeyId =
        optReference(
            "api_key_id",
            ApiKeyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_spend_limits_api_key_id_fkey",
        )
    val userId =
        optReference(
            "user_id",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_spend_limits_user_id_fkey",
        )
    val interval = pgEnum("interval", "\"GatewaySpendLimitInterval\"", GatewaySpendLimitInterval.entries, GatewaySpendLimitInterval::dbName)
    val limitUsd = decimal("limit_usd", 65, 30)
    val alertThresholds = array("alert_thresholds", IntegerColumnType()).nullable().databaseDefault("'{\"100\"}'") { listOf(100) }
    val countingStartedAt = pgTimestamp("counting_started_at", 3)
    val ownerUserId =
        optReference(
            "owner_user_id",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_spend_limits_owner_user_id_fkey",
        )
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "gateway_spend_limits_pkey")

    init {
        index("gateway_spend_limits_organization_id_idx", false, organizationId)
        index("gateway_spend_limits_user_id_idx", false, userId)
        uniqueIndex("gateway_spend_limits_api_key_id_interval_key", apiKeyId, interval)
        uniqueIndex("gateway_spend_limits_organization_id_user_id_interval_key", organizationId, userId, interval)
    }
}
