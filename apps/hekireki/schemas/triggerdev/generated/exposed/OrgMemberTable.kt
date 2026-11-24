package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object OrgMemberTable : IdTable<String>("\"OrgMember\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"OrgMember_organizationId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"OrgMember_userId_fkey\"",
        )
    val role = pgEnum("role", "\"OrgMemberRole\"", OrgMemberRole.entries, OrgMemberRole::dbName).default(OrgMemberRole.MEMBER)

    /** Who invited them (a User id, kept if that user is deleted). Null if not invited or joined before this was recorded. */
    val invitedById = text("invitedById").nullable()

    /** How they joined: created, invite, sso_jit, directory_sync or manual. Null if joined before this was recorded. */
    val joinMethod = text("joinMethod").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"OrgMember_pkey\"")

    init {
        index("\"OrgMember_userId_idx\"", false, userId)
        uniqueIndex("\"OrgMember_organizationId_userId_key\"", organizationId, userId)
    }
}
