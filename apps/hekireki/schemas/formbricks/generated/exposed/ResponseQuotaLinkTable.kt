package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Junction table linking responses to quotas.
 * Tracks which responses counted towards specific quotas and their status.
 *
 * @property response - The response that counted towards the quota
 * @property quota - The quota this response counted towards
 * @property status - Whether the response was screened in or out
 */
object ResponseQuotaLinkTable : CompositeIdTable("\"ResponseQuotaLink\"") {
    val responseId =
        reference(
            "responseId",
            ResponseTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ResponseQuotaLink_responseId_fkey\"",
        )
    val quotaId =
        reference(
            "quotaId",
            SurveyQuotaTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ResponseQuotaLink_quotaId_fkey\"",
        )
    val status = pgEnum("status", "\"ResponseQuotaLinkStatus\"", ResponseQuotaLinkStatus.entries, ResponseQuotaLinkStatus::dbName)
    override val primaryKey = PrimaryKey(responseId, quotaId, name = "\"ResponseQuotaLink_pkey\"")

    init {
        addIdColumn(responseId)
        addIdColumn(quotaId)
        index("\"ResponseQuotaLink_quotaId_status_idx\"", false, quotaId, status)
    }
}
