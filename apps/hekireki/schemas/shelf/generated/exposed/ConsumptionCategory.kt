package shelf.models

/** * Classifies the type of quantity change in a ConsumptionLog entry. */
enum class ConsumptionCategory(
    val dbName: String,
) {
    CHECKOUT("CHECKOUT"),
    RETURN("RETURN"),
    RESTOCK("RESTOCK"),
    ADJUSTMENT("ADJUSTMENT"),
    LOSS("LOSS"),
    CONSUME("CONSUME"),
    DAMAGE("DAMAGE"),
    MOVE("MOVE"),
    INITIAL("INITIAL"),
    OPENING_BALANCE("OPENING_BALANCE"),
}
