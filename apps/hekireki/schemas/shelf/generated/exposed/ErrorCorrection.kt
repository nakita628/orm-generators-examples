package shelf.models

enum class ErrorCorrection(
    val dbName: String,
) {
    L("L"),
    M("M"),
    Q("Q"),
    H("H"),
}
