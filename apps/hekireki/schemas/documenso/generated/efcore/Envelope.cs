#nullable enable

using System;
using System.Collections.Generic;

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZDocumentAuthOptionsSchema } from '@documenso/lib/types/document-auth';", "import { ZDocumentFormValuesSchema } from '@documenso/lib/types/document-form-values';", "import { ZSignatureLevelSchema } from '@documenso/lib/types/signature-level';"])
/// </summary>
public partial class Envelope
{
    public string Id { get; set; } = null!;

    public string SecondaryId { get; set; } = null!;

    /// <summary>
    /// @zod.string.describe("A custom external ID you can use to identify the document.")
    /// </summary>
    public string? ExternalId { get; set; }

    public EnvelopeType Type { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? CompletedAt { get; set; }

    public DateTime? DeletedAt { get; set; }

    public string Title { get; set; } = null!;

    public DocumentStatus Status { get; set; } = DocumentStatus.Draft;

    public DocumentSource Source { get; set; }

    /// <summary>
    /// @zod.string.describe("The token for viewing the document using the QR code on the certificate.")
    /// </summary>
    public string? QrToken { get; set; }

    /// <summary>
    /// [SignatureLevel] @zod.custom.use(ZSignatureLevelSchema)
    /// </summary>
    public string SignatureLevel { get; set; } = null!;

    public int InternalVersion { get; set; }

    public bool UseLegacyFieldInsertion { get; set; } = false;

    public virtual ICollection<EnvelopeItem> EnvelopeItems { get; set; } = new List<EnvelopeItem>();

    public virtual ICollection<Recipient> Recipients { get; set; } = new List<Recipient>();

    public virtual ICollection<Field> Fields { get; set; } = new List<Field>();

    public virtual ICollection<EnvelopeContent> Contents { get; set; } = new List<EnvelopeContent>();

    public virtual ICollection<DocumentShareLink> ShareLinks { get; set; } = new List<DocumentShareLink>();

    public virtual ICollection<DocumentAuditLog> AuditLogs { get; set; } = new List<DocumentAuditLog>();

    /// <summary>
    /// [DocumentAuthOptions] @zod.custom.use(ZDocumentAuthOptionsSchema)
    /// </summary>
    public string? AuthOptions { get; set; }

    /// <summary>
    /// [DocumentFormValues] @zod.custom.use(ZDocumentFormValuesSchema)
    /// </summary>
    public string? FormValues { get; set; }

    public DocumentVisibility Visibility { get; set; } = DocumentVisibility.Everyone;

    public TemplateType TemplateType { get; set; } = TemplateType.Private;

    public string PublicTitle { get; set; } = "";

    public string PublicDescription { get; set; } = "";

    public virtual TemplateDirectLink? DirectLink { get; set; }

    public int? TemplateId { get; set; }

    /// <summary>
    /// @zod.number.describe("The ID of the user that created this document.")
    /// </summary>
    public int UserId { get; set; }

    public virtual User User { get; set; } = null!;

    public int TeamId { get; set; }

    public virtual Team Team { get; set; } = null!;

    public virtual Folder? Folder { get; set; }

    public string? FolderId { get; set; }

    public string DocumentMetaId { get; set; } = null!;

    public virtual DocumentMeta DocumentMeta { get; set; } = null!;

    public virtual ICollection<EnvelopeAttachment> EnvelopeAttachments { get; set; } = new List<EnvelopeAttachment>();
}
