package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.IntegerColumnType
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A progressive checkout SESSION: one scan batch, covering N assets via the
 * positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
 *
 * Records sessions and the units they claimed — NOT whether an asset is out.
 * The all-at-once checkout (`checkoutBooking`, the Check out button) flips
 * statuses and writes NO row here, so **absence proves nothing** and must never
 * be read as "was not checked out". Reading it that way is what let a single
 * later scan mark every button-checked-out asset on a booking as
 * "never checked out". `BookingAsset.checkedOutAt` answers that question.
 */
object PartialBookingCheckoutTable : IdTable<String>("\"PartialBookingCheckout\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val assetIds = array("assetIds", TextColumnType()).nullable()
    val quantities = array("quantities", IntegerColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val bookingAssetIds = array("bookingAssetIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val checkoutCount = integer("checkoutCount")
    val checkoutTimestamp =
        pgTimestamptz("checkoutTimestamp", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val bookingId =
        reference(
            "bookingId",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PartialBookingCheckout_bookingId_fkey\"",
        )
    val checkedOutById =
        reference(
            "checkedOutById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PartialBookingCheckout_checkedOutById_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PartialBookingCheckout_pkey\"")

    init {
        index("\"PartialBookingCheckout_bookingId_idx\"", false, bookingId)
        index("\"PartialBookingCheckout_checkedOutById_idx\"", false, checkedOutById)
        index("\"PartialBookingCheckout_checkoutTimestamp_idx\"", false, checkoutTimestamp)
        index("\"PartialBookingCheckout_bookingId_checkoutTimestamp_idx\"", false, bookingId, checkoutTimestamp)
    }
}
