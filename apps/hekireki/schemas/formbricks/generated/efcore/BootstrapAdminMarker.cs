#nullable enable

using NpgsqlTypes;

namespace Formbricks.Models;

/// <summary>
/// The bootstrap marker's only value (ENG-2247). A one-variant enum rather than a boolean on
/// purpose: a unique index over a nullable boolean has a slot for `true` AND one for `false`, so the
/// first `false` ever written would take the second slot and break every later sign-up. There is no
/// `false` to write here, which is the same guarantee without a CHECK constraint Prisma cannot model.
/// </summary>
public enum BootstrapAdminMarker
{
    [PgName("bootstrapAdmin")]
    BootstrapAdmin,
}
