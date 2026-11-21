package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** deprecated, we hadn't included the project id in the unique constraint */
object TaskRunCounterTable : IdTable<String>("\"TaskRunCounter\"") {
    override val id = text("taskIdentifier").entityId()
    val lastNumber = integer("lastNumber").default(0)
    override val primaryKey = PrimaryKey(id, name = "\"TaskRunCounter_pkey\"")
}
