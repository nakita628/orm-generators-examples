package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserUpdateReadTable : IdTable<String>("\"UserUpdateRead\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserUpdateRead_userId_fkey\"",
        )
    val updateId =
        reference(
            "updateId",
            UpdateTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserUpdateRead_updateId_fkey\"",
        )
    val readAt = pgTimestamp("readAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"UserUpdateRead_pkey\"")

    init {
        index("\"UserUpdateRead_userId_idx\"", false, userId)
        index("\"UserUpdateRead_updateId_idx\"", false, updateId)
        index("\"UserUpdateRead_readAt_idx\"", false, readAt)
        uniqueIndex("\"UserUpdateRead_userId_updateId_key\"", userId, updateId)
    }
}
