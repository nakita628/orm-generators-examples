package shelf.models

/** * Controls what happens after checkout for quantity-tracked assets. */
enum class ConsumptionType(
    val dbName: String,
) {
    ONE_WAY("ONE_WAY"),
    TWO_WAY("TWO_WAY"),
}
