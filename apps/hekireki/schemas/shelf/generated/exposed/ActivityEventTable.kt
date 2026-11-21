package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object ActivityEventTable : IdTable<String>("\"ActivityEvent\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ActivityEvent_organizationId_fkey\"",
        )
    val occurredAt = pgTimestamptz("occurredAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val actorUserId = text("actorUserId").nullable()
    val actorSnapshot = jsonb<String>("actorSnapshot", { it }, { it }).nullable()
    val action = pgEnum("action", "\"ActivityAction\"", ActivityAction.entries, ActivityAction::dbName)
    val entityType = pgEnum("entityType", "\"ActivityEntity\"", ActivityEntity.entries, ActivityEntity::dbName)
    val entityId = text("entityId")
    val assetId = text("assetId").nullable()
    val bookingId = text("bookingId").nullable()
    val auditSessionId = text("auditSessionId").nullable()
    val auditAssetId = text("auditAssetId").nullable()
    val kitId = text("kitId").nullable()
    val locationId = text("locationId").nullable()
    val teamMemberId = text("teamMemberId").nullable()
    val targetUserId = text("targetUserId").nullable()
    val field = text("field").nullable()
    val fromValue = jsonb<String>("fromValue", { it }, { it }).nullable()
    val toValue = jsonb<String>("toValue", { it }, { it }).nullable()
    val meta = jsonb<String>("meta", { it }, { it }).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"ActivityEvent_pkey\"")

    init {
        index("\"ActivityEvent_organizationId_occurredAt_idx\"", false, organizationId, occurredAt)
        index("\"ActivityEvent_organizationId_action_occurredAt_idx\"", false, organizationId, action, occurredAt)
        index("ActivityEvent_organizationId_entityType_entityId_occurredAt_idx", false, organizationId, entityType, entityId, occurredAt)
        index("\"ActivityEvent_actorUserId_occurredAt_idx\"", false, actorUserId, occurredAt)
        index("\"ActivityEvent_assetId_occurredAt_idx\"", false, assetId, occurredAt)
        index("\"ActivityEvent_bookingId_occurredAt_idx\"", false, bookingId, occurredAt)
        index("\"ActivityEvent_auditSessionId_occurredAt_idx\"", false, auditSessionId, occurredAt)
    }
}
