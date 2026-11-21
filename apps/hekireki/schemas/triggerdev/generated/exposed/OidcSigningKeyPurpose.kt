package triggerdev.models

/** Each purpose is a separate key set with its own issuer, rotation and publication. */
enum class OidcSigningKeyPurpose(
    val dbName: String,
) {
    EXTERNAL_OIDC("EXTERNAL_OIDC"),
    INTERNAL_RUN("INTERNAL_RUN"),
}
