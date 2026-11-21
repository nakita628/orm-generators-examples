package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetIndexSettingsTable : IdTable<String>("\"AssetIndexSettings\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetIndexSettings_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetIndexSettings_organizationId_fkey\"",
        )
    val mode = pgEnum("mode", "\"AssetIndexMode\"", AssetIndexMode.entries, AssetIndexMode::dbName).default(AssetIndexMode.SIMPLE)
    val columns1 =
        jsonb<String>("columns", { it }, { it })
            .default(
                "[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]",
            )
    val freezeColumn = bool("freezeColumn").default(true)
    val showAssetImage = bool("showAssetImage").default(true)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetIndexSettings_pkey\"")

    init {
        index("\"AssetIndexSettings_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"AssetIndexSettings_userId_organizationId_key\"", userId, organizationId)
    }
}
