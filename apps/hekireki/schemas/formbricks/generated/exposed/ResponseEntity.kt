package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class ResponseEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ResponseEntity>(ResponseTable)

    var createdAt by ResponseTable.createdAt
    var updatedAt by ResponseTable.updatedAt
    var finished by ResponseTable.finished
    var survey by SurveyEntity referencedOn ResponseTable.surveyId
    var surveyId by ResponseTable.surveyId
    var contact by ContactEntity optionalReferencedOn ResponseTable.contactId
    var contactId by ResponseTable.contactId
    var endingId by ResponseTable.endingId

    /** [ResponseData] */
    var data by ResponseTable.data

    /** [ResponseVariables] */
    var variables by ResponseTable.variables

    /** [ResponseTtc] */
    var ttc by ResponseTable.ttc

    /** [ResponseMeta] */
    var meta by ResponseTable.meta

    /**
     * Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
     * or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
     * from the client. Null means no ingest boundary has written it, which is not the same claim as
     * an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
     * and surfacing that in the UI (v2) would have nothing to read.
     * [ResponseIngestFlags]
     */
    var ingestFlags by ResponseTable.ingestFlags

    /**
     * A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
     * export by default and from Feedback Records, and nothing else: it is never deleted and every
     * count, list, webhook and integration still includes it (ENG-6780).
     */
    var outcome by ResponseTable.outcome

    /**
     * What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
     * from them, and per element its options, matrix rows and columns, by stable ids. Written by the
     * survey while they answer; null means unknown (older responses and bundles), read as before.
     * Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
     * webhook or integration.
     * [ResponseNotShown]
     */
    var notShown by ResponseTable.notShown
    val tags by TagsOnResponsesEntity referrersOn TagsOnResponsesTable.responseId
    val quotaLinks by ResponseQuotaLinkEntity referrersOn ResponseQuotaLinkTable.responseId

    /** [ResponseContactAttributes] */
    var contactAttributes by ResponseTable.contactAttributes
    var singleUseId by ResponseTable.singleUseId
    var language by ResponseTable.language
    var displayId by ResponseTable.displayId
    var display by DisplayEntity optionalReferencedOn ResponseTable.displayId
    val workflowRuns by WorkflowRunEntity optionalReferrersOn WorkflowRunTable.responseId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ResponseTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
