package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object WorkingHoursOverrideTable : IdTable<String>("\"WorkingHoursOverride\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val date = pgDate("date")
    val isOpen = bool("isOpen").default(false)
    val openTime = text("openTime").nullable()
    val closeTime = text("closeTime").nullable()
    val reason = text("reason").nullable()
    val workingHoursId =
        reference(
            "workingHoursId",
            WorkingHoursTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"WorkingHoursOverride_workingHoursId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"WorkingHoursOverride_pkey\"")

    init {
        index("\"WorkingHoursOverride_workingHoursId_date_idx\"", false, workingHoursId, date)
        index("\"WorkingHoursOverride_date_isOpen_idx\"", false, date, isOpen)
        uniqueIndex("\"WorkingHoursOverride_workingHoursId_date_key\"", workingHoursId, date)
    }
}
