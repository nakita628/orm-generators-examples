package shelf.models

enum class InviteStatuses(
    val dbName: String,
) {
    PENDING("PENDING"),
    ACCEPTED("ACCEPTED"),
    REJECTED("REJECTED"),
    INVALIDATED("INVALIDATED"),
}
