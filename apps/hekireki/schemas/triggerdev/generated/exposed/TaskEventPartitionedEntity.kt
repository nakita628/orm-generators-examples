package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
class TaskEventPartitionedEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<TaskEventPartitionedEntity>(TaskEventPartitionedTable)

    val id1 by TaskEventPartitionedTable.id1

    /** This matches the span name for a trace event, or the log body for a log event */
    var message by TaskEventPartitionedTable.message
    var traceId by TaskEventPartitionedTable.traceId
    var spanId by TaskEventPartitionedTable.spanId
    var parentId by TaskEventPartitionedTable.parentId
    var tracestate by TaskEventPartitionedTable.tracestate
    var isError by TaskEventPartitionedTable.isError
    var isPartial by TaskEventPartitionedTable.isPartial
    var isCancelled by TaskEventPartitionedTable.isCancelled
    var serviceName by TaskEventPartitionedTable.serviceName
    var serviceNamespace by TaskEventPartitionedTable.serviceNamespace
    var level by TaskEventPartitionedTable.level
    var kind by TaskEventPartitionedTable.kind
    var status by TaskEventPartitionedTable.status
    var links by TaskEventPartitionedTable.links
    var events by TaskEventPartitionedTable.events

    /** This is the time the event started in nanoseconds since the epoch */
    var startTime by TaskEventPartitionedTable.startTime

    /** This is the duration of the event in nanoseconds */
    var duration by TaskEventPartitionedTable.duration
    var attemptId by TaskEventPartitionedTable.attemptId
    var attemptNumber by TaskEventPartitionedTable.attemptNumber
    var environmentId by TaskEventPartitionedTable.environmentId
    var environmentType by TaskEventPartitionedTable.environmentType
    var organizationId by TaskEventPartitionedTable.organizationId
    var projectId by TaskEventPartitionedTable.projectId
    var projectRef by TaskEventPartitionedTable.projectRef
    var runId by TaskEventPartitionedTable.runId
    var runIsTest by TaskEventPartitionedTable.runIsTest
    var idempotencyKey by TaskEventPartitionedTable.idempotencyKey
    var taskSlug by TaskEventPartitionedTable.taskSlug
    var taskPath by TaskEventPartitionedTable.taskPath
    var taskExportName by TaskEventPartitionedTable.taskExportName
    var workerId by TaskEventPartitionedTable.workerId
    var workerVersion by TaskEventPartitionedTable.workerVersion
    var queueId by TaskEventPartitionedTable.queueId
    var queueName by TaskEventPartitionedTable.queueName
    var batchId by TaskEventPartitionedTable.batchId

    /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
    var properties by TaskEventPartitionedTable.properties

    /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
    var metadata by TaskEventPartitionedTable.metadata

    /** This represents all span attributes in the $style namespace, like $style */
    var style by TaskEventPartitionedTable.style

    /** This represents all span attributes in the $output namespace, like $output */
    var output by TaskEventPartitionedTable.output

    /** This represents the mimetype of the output, such as application/json or application/super+json */
    var outputType by TaskEventPartitionedTable.outputType
    var payload by TaskEventPartitionedTable.payload
    var payloadType by TaskEventPartitionedTable.payloadType
    val createdAt by TaskEventPartitionedTable.createdAt
    var usageDurationMs by TaskEventPartitionedTable.usageDurationMs
    var usageCostInCents by TaskEventPartitionedTable.usageCostInCents
    var machinePreset by TaskEventPartitionedTable.machinePreset
    var machinePresetCpu by TaskEventPartitionedTable.machinePresetCpu
    var machinePresetMemory by TaskEventPartitionedTable.machinePresetMemory
    var machinePresetCentsPerMs by TaskEventPartitionedTable.machinePresetCentsPerMs
}
