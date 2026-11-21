package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserContactTable : IdTable<String>("\"UserContact\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val phone = text("phone").nullable()
    val street = text("street").nullable()
    val city = text("city").nullable()
    val stateProvince = text("stateProvince").nullable()
    val zipPostalCode = text("zipPostalCode").nullable()
    val countryRegion = text("countryRegion").nullable()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserContact_userId_fkey\"",
        ).uniqueIndex("\"UserContact_userId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"UserContact_pkey\"")

    init {
        index("\"UserContact_userId_idx\"", false, userId)
        index("\"UserContact_phone_idx\"", false, phone)
        index("\"UserContact_city_stateProvince_idx\"", false, city, stateProvince)
        index("\"UserContact_countryRegion_idx\"", false, countryRegion)
        index("\"UserContact_zipPostalCode_idx\"", false, zipPostalCode)
        index("\"UserContact_city_countryRegion_idx\"", false, city, countryRegion)
    }
}
