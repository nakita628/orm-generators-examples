#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// A saved layout for a generated PDF, kept per team workspace under a name:
/// which columns, rows and header lines print, and in what order the rows
/// come. A row is a layout, not a generated file; the person generating a PDF
/// picks a template by name.
/// </summary>
public partial class PdfTemplate
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public PdfTemplateKind Kind { get; set; } = PdfTemplateKind.Booking;

    /// <summary>
    /// The layout's switches. Read through `parseBookingTemplateConfig`,
    /// which fills any key a stored row lacks from the Checklist template's
    /// values.
    /// </summary>
    public string Config { get; set; } = "{}";

    /// <summary>
    /// The template the PDF dialog offers first. One per workspace and kind,
    /// enforced by a partial unique index in raw SQL (migration
    /// 20261103120000_add_pdf_templates). Names are also unique per workspace
    /// and kind ignoring case, through an expression index in the same
    /// migration.
    /// </summary>
    public bool IsDefault { get; set; } = false;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
