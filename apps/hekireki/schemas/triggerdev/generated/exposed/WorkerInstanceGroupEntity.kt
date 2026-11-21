package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class WorkerInstanceGroupEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkerInstanceGroupEntity>(WorkerInstanceGroupTable)

    var type by WorkerInstanceGroupTable.type

    /** For example "us-east-1" */
    var name by WorkerInstanceGroupTable.name

    /**
     * If managed, it will default to the name, e.g. "us-east-1"
     * If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
     */
    var masterQueue by WorkerInstanceGroupTable.masterQueue

    /** "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes */
    var description by WorkerInstanceGroupTable.description
    var hidden by WorkerInstanceGroupTable.hidden
    var token by WorkerGroupTokenEntity referencedOn WorkerInstanceGroupTable.tokenId
    var tokenId by WorkerInstanceGroupTable.tokenId
    val workers by WorkerInstanceEntity referrersOn WorkerInstanceTable.workerGroupId
    val backgroundWorkers by BackgroundWorkerEntity optionalReferrersOn BackgroundWorkerTable.workerGroupId
    val defaultForProjects by ProjectEntity optionalReferrersOn ProjectTable.defaultWorkerGroupId
    var organization by OrganizationEntity optionalReferencedOn WorkerInstanceGroupTable.organizationId
    var organizationId by WorkerInstanceGroupTable.organizationId
    var project by ProjectEntity optionalReferencedOn WorkerInstanceGroupTable.projectId
    var projectId by WorkerInstanceGroupTable.projectId
    var cloudProvider by WorkerInstanceGroupTable.cloudProvider

    /** "usa", "europe", etc. Used like a pseudo enum for things like flags */
    var location by WorkerInstanceGroupTable.location
    var staticIPs by WorkerInstanceGroupTable.staticIPs
    var workloadType by WorkerInstanceGroupTable.workloadType

    /** Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs. */
    var region by WorkerInstanceGroupTable.region

    /**
     * When true, runs enqueued to this worker queue may skip the intermediate queue
     * and be pushed directly to the worker queue when concurrency is available.
     */
    var enableFastPath by WorkerInstanceGroupTable.enableFastPath
    var createdAt by WorkerInstanceGroupTable.createdAt
    var updatedAt by WorkerInstanceGroupTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkerInstanceGroupTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
