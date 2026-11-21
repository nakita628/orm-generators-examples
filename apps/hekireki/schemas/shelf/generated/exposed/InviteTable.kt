package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object InviteTable : IdTable<String>("\"Invite\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val inviterId =
        reference(
            "inviterId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Invite_inviterId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Invite_organizationId_fkey\"",
        )
    val inviteeUserId =
        optReference(
            "inviteeUserId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Invite_inviteeUserId_fkey\"",
        )
    val teamMemberId =
        reference(
            "teamMemberId",
            TeamMemberTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Invite_teamMemberId_fkey\"",
        )
    val inviteeEmail = text("inviteeEmail")
    val status = pgEnum("status", "\"InviteStatuses\"", InviteStatuses.entries, InviteStatuses::dbName).default(InviteStatuses.PENDING)
    val inviteCode = text("inviteCode")
    val roles = array("roles", PgEnumColumnType("\"OrganizationRoles\"", OrganizationRoles.entries, OrganizationRoles::dbName)).nullable()
    val inviteMessage = varchar("inviteMessage", 1000).nullable()
    val expiresAt = pgTimestamp("expiresAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Invite_pkey\"")

    init {
        index("\"Invite_inviteeUserId_idx\"", false, inviteeUserId)
        index("\"Invite_inviterId_idx\"", false, inviterId)
        index("\"Invite_organizationId_idx\"", false, organizationId)
        index("\"Invite_teamMemberId_idx\"", false, teamMemberId)
    }
}
