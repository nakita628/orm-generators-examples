#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum ActivityEntity
{
    [PgName("ASSET")]
    Asset,

    [PgName("BOOKING")]
    Booking,

    [PgName("AUDIT")]
    Audit,

    [PgName("KIT")]
    Kit,

    [PgName("LOCATION")]
    Location,

    [PgName("TEAM_MEMBER")]
    TeamMember,

    [PgName("CUSTODY")]
    Custody,

    [PgName("USER")]
    User,

    [PgName("ORGANIZATION")]
    Organization,
}
