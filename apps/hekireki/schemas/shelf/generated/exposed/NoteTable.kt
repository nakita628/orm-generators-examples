package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object NoteTable : IdTable<String>("\"Note\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val content = text("content")
    val type = pgEnum("type", "\"NoteType\"", NoteType.entries, NoteType::dbName).default(NoteType.COMMENT)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Note_userId_fkey\"",
        )
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Note_assetId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Note_pkey\"")

    init {
        index("\"Note_assetId_idx\"", false, assetId)
        index("\"Note_userId_idx\"", false, userId)
    }
}
