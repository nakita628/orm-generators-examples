package langfuse.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times. */
object GatewaySpendLimitAlertTable : IdTable<String>("gateway_spend_limit_alerts") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val spendLimitId =
        reference(
            "spend_limit_id",
            GatewaySpendLimitTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "gateway_spend_limit_alerts_spend_limit_id_fkey",
        )
    val subjectId = text("subject_id").default("")
    val type = pgEnum("type", "\"GatewaySpendLimitAlertType\"", GatewaySpendLimitAlertType.entries, GatewaySpendLimitAlertType::dbName)
    val thresholdPercent = integer("threshold_percent").default(100)
    val countingStartedAt = pgTimestamp("counting_started_at", 3)
    val windowStart = pgTimestamp("window_start", 3)
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "gateway_spend_limit_alerts_pkey")

    init {
        uniqueIndex(
            "gateway_spend_limit_alerts_limit_type_counting_window_key",
            spendLimitId,
            subjectId,
            type,
            thresholdPercent,
            countingStartedAt,
            windowStart,
        )
    }
}
