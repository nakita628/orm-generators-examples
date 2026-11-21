#nullable enable

using System.Collections.Generic;

namespace Documenso.Models;

/// <summary>
/// Generic stored data (e.g. an uploaded image) which other models reference.
/// The kind of data and its kind specific details live in `metadata`.
/// @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
/// </summary>
public partial class DataContent
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// How `data` is stored, e.g. an S3 path or inline bytes.
    /// </summary>
    public DocumentDataType Type { get; set; }

    public string Data { get; set; } = null!;

    /// <summary>
    /// [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
    /// </summary>
    public string Metadata { get; set; } = null!;

    public virtual ICollection<EnvelopeContent> EnvelopeContents { get; set; } = new List<EnvelopeContent>();
}
