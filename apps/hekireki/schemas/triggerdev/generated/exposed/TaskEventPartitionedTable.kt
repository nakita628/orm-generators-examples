package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
object TaskEventPartitionedTable : CompositeIdTable("\"TaskEventPartitioned\"") {
    val id1 = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /** This matches the span name for a trace event, or the log body for a log event */
    val message = text("message")
    val traceId = text("traceId")
    val spanId = text("spanId")
    val parentId = text("parentId").nullable()
    val tracestate = text("tracestate").nullable()
    val isError = bool("isError").default(false)
    val isPartial = bool("isPartial").default(false)
    val isCancelled = bool("isCancelled").default(false)
    val serviceName = text("serviceName")
    val serviceNamespace = text("serviceNamespace")
    val level = pgEnum("level", "\"TaskEventLevel\"", TaskEventLevel.entries, TaskEventLevel::dbName).default(TaskEventLevel.TRACE)
    val kind = pgEnum("kind", "\"TaskEventKind\"", TaskEventKind.entries, TaskEventKind::dbName).default(TaskEventKind.INTERNAL)
    val status = pgEnum("status", "\"TaskEventStatus\"", TaskEventStatus.entries, TaskEventStatus::dbName).default(TaskEventStatus.UNSET)
    val links = jsonb<String>("links", { it }, { it }).nullable()
    val events = jsonb<String>("events", { it }, { it }).nullable()

    /** This is the time the event started in nanoseconds since the epoch */
    val startTime = long("startTime")

    /** This is the duration of the event in nanoseconds */
    val duration = long("duration").default(0L)
    val attemptId = text("attemptId").nullable()
    val attemptNumber = integer("attemptNumber").nullable()
    val environmentId = text("environmentId")
    val environmentType =
        pgEnum("environmentType", "\"RuntimeEnvironmentType\"", RuntimeEnvironmentType.entries, RuntimeEnvironmentType::dbName)
    val organizationId = text("organizationId")
    val projectId = text("projectId")
    val projectRef = text("projectRef")
    val runId = text("runId")
    val runIsTest = bool("runIsTest").default(false)
    val idempotencyKey = text("idempotencyKey").nullable()
    val taskSlug = text("taskSlug")
    val taskPath = text("taskPath").nullable()
    val taskExportName = text("taskExportName").nullable()
    val workerId = text("workerId").nullable()
    val workerVersion = text("workerVersion").nullable()
    val queueId = text("queueId").nullable()
    val queueName = text("queueName").nullable()
    val batchId = text("batchId").nullable()

    /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
    val properties = jsonb<String>("properties", { it }, { it })

    /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()

    /** This represents all span attributes in the $style namespace, like $style */
    val style = jsonb<String>("style", { it }, { it }).nullable()

    /** This represents all span attributes in the $output namespace, like $output */
    val output = jsonb<String>("output", { it }, { it }).nullable()

    /** This represents the mimetype of the output, such as application/json or application/super+json */
    val outputType = text("outputType").nullable()
    val payload = jsonb<String>("payload", { it }, { it }).nullable()
    val payloadType = text("payloadType").nullable()
    val createdAt =
        pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }.entityId()
    val usageDurationMs = integer("usageDurationMs").default(0)
    val usageCostInCents = double("usageCostInCents").default(0.0)
    val machinePreset = text("machinePreset").nullable()
    val machinePresetCpu = double("machinePresetCpu").nullable()
    val machinePresetMemory = double("machinePresetMemory").nullable()
    val machinePresetCentsPerMs = double("machinePresetCentsPerMs").nullable()
    override val primaryKey = PrimaryKey(id1, createdAt, name = "\"TaskEventPartitioned_pkey\"")

    init {
        index("\"TaskEventPartitioned_traceId_idx\"", false, traceId)
        index("\"TaskEventPartitioned_spanId_idx\"", false, spanId)
        index("\"TaskEventPartitioned_runId_idx\"", false, runId)
    }
}
