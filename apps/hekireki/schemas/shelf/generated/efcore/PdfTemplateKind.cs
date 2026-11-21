#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

/// <summary>
/// The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
/// layouts today; the check-in receipt and the audit receipt can join as new
/// kinds.
/// </summary>
public enum PdfTemplateKind
{
    [PgName("BOOKING")]
    Booking,
}
