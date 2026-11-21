#nullable enable

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZEnvelopeContentMetaSchema } from '@documenso/lib/types/envelope-content-meta';"])
/// </summary>
public partial class EnvelopeContent
{
    public string Id { get; set; } = null!;

    public string EnvelopeId { get; set; } = null!;

    public string EnvelopeItemId { get; set; } = null!;

    public virtual EnvelopeItem EnvelopeItem { get; set; } = null!;

    public virtual Envelope Envelope { get; set; } = null!;

    /// <summary>
    /// [EnvelopeContentMeta] @zod.custom.use(ZEnvelopeContentMetaSchema)
    /// </summary>
    public string ContentMeta { get; set; } = null!;

    /// <summary>
    /// Any number of contents may point at the same data content, e.g. a
    /// document created from a template shares the template's images.
    /// </summary>
    public string? DataContentId { get; set; }

    public virtual DataContent? DataContent { get; set; }
}
