package shelf.models

enum class OrganizationRoles(
    val dbName: String,
) {
    ADMIN("ADMIN"),
    BASE("BASE"),
    OWNER("OWNER"),
    SELF_SERVICE("SELF_SERVICE"),
    MANAGER("MANAGER"),
}
