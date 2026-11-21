package formbricks.models

/** The device type counts are split by: the values of the Show when `deviceType` operand. */
enum class CreativeDeviceType(
    val dbName: String,
) {
    PHONE("phone"),
    TABLET("tablet"),
    DESKTOP("desktop"),
}
