#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum NoteType
{
    [PgName("COMMENT")]
    Comment,

    [PgName("UPDATE")]
    Update,
}
