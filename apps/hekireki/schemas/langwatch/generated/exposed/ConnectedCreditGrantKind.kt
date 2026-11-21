package langwatch.models

/** Why a credit was granted. The amount is always prepaid, never promotional. */
enum class ConnectedCreditGrantKind(
    val dbName: String,
) {
    COMMIT("commit"),
    ADDED("added"),
    RENEWAL("renewal"),
}
