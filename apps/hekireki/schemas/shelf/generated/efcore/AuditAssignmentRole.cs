#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum AuditAssignmentRole
{
    [PgName("LEAD")]
    Lead,

    [PgName("PARTICIPANT")]
    Participant,
}
