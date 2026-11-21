package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object WorkerInstanceTable : IdTable<String>("\"WorkerInstance\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /** For example "worker-1" */
    val name = text("name")

    /**
     * If managed, it will default to the name, e.g. "worker-1"
     * If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1"
     */
    val resourceIdentifier = text("resourceIdentifier")
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()
    val workerGroupId =
        reference(
            "workerGroupId",
            WorkerInstanceGroupTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkerInstance_workerGroupId_fkey\"",
        )
    val organizationId =
        optReference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkerInstance_organizationId_fkey\"",
        )
    val projectId =
        optReference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkerInstance_projectId_fkey\"",
        )
    val environmentId =
        optReference(
            "environmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkerInstance_environmentId_fkey\"",
        )
    val deploymentId =
        optReference(
            "deploymentId",
            WorkerDeploymentTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkerInstance_deploymentId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastDequeueAt = pgTimestamp("lastDequeueAt", 3).nullable()
    val lastHeartbeatAt = pgTimestamp("lastHeartbeatAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"WorkerInstance_pkey\"")

    init {
        uniqueIndex("\"WorkerInstance_workerGroupId_resourceIdentifier_key\"", workerGroupId, resourceIdentifier)
    }
}
