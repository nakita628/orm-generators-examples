#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// The device type counts are split by: the values of the Show when `deviceType` operand.
/// </summary>
public enum CreativeDeviceType
{
    [PgName("phone")]
    Phone,

    [PgName("tablet")]
    Tablet,

    [PgName("desktop")]
    Desktop,
}
