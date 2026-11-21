package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The notice for one target under one policy, so each notice goes out once: a survey or a member, or
 * the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
 * claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
 * members) before the target's clock: then a later cycle claims it again and sends a new one.
 *
 * A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
 * `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
 * hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
 * 'surveys' AND "clockAt" IS NOT NULL`) in
 * `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
 * that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
 * "Indexes Prisma cannot express" in the README).
 */
object RetentionNoticeTable : IdTable<String>("\"RetentionNotice\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionNotice_organizationId_fkey\"",
        )
    val entity = pgEnum("entity", "\"RetentionEntity\"", RetentionEntity.entries, RetentionEntity::dbName)
    val surveyId =
        optReference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionNotice_surveyId_fkey\"",
        )
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RetentionNotice_userId_fkey\"",
        )

    /** When the notice was claimed for sending. A claim that never got delivered can be taken over. */
    val sentAt = pgTimestamp("sentAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * When the notice was handed to the mail transport, or recorded without one (no SMTP). The
     * warning only counts from a delivered notice, so a crash between claim and send can't let the
     * action run without it.
     */
    val deliveredAt = pgTimestamp("deliveredAt", 3).nullable()

    /** Whether an email actually went out. False when SMTP isn't configured. */
    val emailSent = bool("emailSent").default(false)

    /**
     * Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
     * can't mark a newer claim delivered.
     */
    val claimToken = text("claimToken").nullable()

    /**
     * The target's clock the notice was computed for (surveys and members). A notice counts only while
     * the target's clock is still this one, so activity between the sweep's read and its claim voids
     * it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
     * for a survey notice voided for good: a person deleted or moved one of its responses, or the
     * responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
     */
    val clockAt = pgTimestamp("clockAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"RetentionNotice_pkey\"")

    init {
        index("\"RetentionNotice_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"RetentionNotice_surveyId_entity_key\"", surveyId, entity)
        uniqueIndex("\"RetentionNotice_userId_organizationId_entity_key\"", userId, organizationId, entity)
    }
}
