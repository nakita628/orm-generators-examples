package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Stores a user's response to a survey, including their answers and metadata.
 * Each response is linked to a specific survey and optionally to a contact.
 *
 * @property id - Unique identifier for the response
 * @property finished - Whether the survey was completed
 * @property survey - The associated survey
 * @property contact - The optional contact who provided the response
 * @property data - JSON object containing the actual response data
 * @property variables - Custom variables used in the response
 * @property ttc - Time to completion metrics
 * @property meta - Additional metadata about the response
 */
object ResponseTable : IdTable<String>("\"Response\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val finished = bool("finished").default(false)
    val surveyId =
        reference(
            "surveyId",
            SurveyTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Response_surveyId_fkey\"",
        )
    val contactId =
        optReference(
            "contactId",
            ContactTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Response_contactId_fkey\"",
        )
    val endingId = text("endingId").nullable()

    /** [ResponseData] */
    val data = jsonb<String>("data", { it }, { it }).default("{}")

    /** [ResponseVariables] */
    val variables = jsonb<String>("variables", { it }, { it }).default("{}")

    /** [ResponseTtc] */
    val ttc = jsonb<String>("ttc", { it }, { it }).default("{}")

    /** [ResponseMeta] */
    val meta = jsonb<String>("meta", { it }, { it }).default("{}")

    /**
     * Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
     * or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
     * from the client. Null means no ingest boundary has written it, which is not the same claim as
     * an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
     * and surfacing that in the UI (v2) would have nothing to read.
     * [ResponseIngestFlags]
     */
    val ingestFlags = jsonb<String>("ingest_flags", { it }, { it }).nullable()

    /**
     * A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
     * export by default and from Feedback Records, and nothing else: it is never deleted and every
     * count, list, webhook and integration still includes it (ENG-6780).
     */
    val outcome = pgEnum("outcome", "\"ResponseOutcome\"", ResponseOutcome.entries, ResponseOutcome::dbName).nullable()

    /**
     * What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
     * from them, and per element its options, matrix rows and columns, by stable ids. Written by the
     * survey while they answer; null means unknown (older responses and bundles), read as before.
     * Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
     * webhook or integration.
     * [ResponseNotShown]
     */
    val notShown = jsonb<String>("not_shown", { it }, { it }).nullable()

    /** [ResponseContactAttributes] */
    val contactAttributes = jsonb<String>("contactAttributes", { it }, { it }).nullable()
    val singleUseId = text("singleUseId").nullable()
    val language = text("language").nullable()
    val displayId =
        optReference(
            "displayId",
            DisplayTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Response_displayId_fkey\"",
        ).uniqueIndex("\"Response_displayId_key\"")
    override val primaryKey = PrimaryKey(id, name = "\"Response_pkey\"")

    init {
        index("\"Response_created_at_idx\"", false, createdAt)
        index("\"Response_created_at_id_idx\"", false, createdAt, id)
        index("\"Response_surveyId_created_at_idx\"", false, surveyId, createdAt)
        index("\"Response_surveyId_created_at_id_idx\"", false, surveyId, createdAt, id)
        index("\"Response_contactId_created_at_idx\"", false, contactId, createdAt)
        index("\"Response_contactId_created_at_id_idx\"", false, contactId, createdAt, id)
        index("\"Response_surveyId_finished_idx\"", false, surveyId, finished)
        uniqueIndex("\"Response_surveyId_singleUseId_key\"", surveyId, singleUseId)
    }
}
