package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Junction table linking responses to quotas.
 * Tracks which responses counted towards specific quotas and their status.
 *
 * @property response - The response that counted towards the quota
 * @property quota - The quota this response counted towards
 * @property status - Whether the response was screened in or out
 */
class ResponseQuotaLinkEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<ResponseQuotaLinkEntity>(ResponseQuotaLinkTable)

    val response by ResponseEntity referencedOn ResponseQuotaLinkTable.responseId
    val responseId by ResponseQuotaLinkTable.responseId
    val quota by SurveyQuotaEntity referencedOn ResponseQuotaLinkTable.quotaId
    val quotaId by ResponseQuotaLinkTable.quotaId
    var status by ResponseQuotaLinkTable.status
}
