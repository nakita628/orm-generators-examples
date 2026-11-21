#nullable enable

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"])
/// </summary>
public partial class DocumentMeta
{
    public string Id { get; set; } = null!;

    public string? Subject { get; set; }

    public string? Message { get; set; }

    public string? Timezone { get; set; } = "Etc/UTC";

    public string? DateFormat { get; set; } = "yyyy-MM-dd hh:mm a";

    public string? RedirectUrl { get; set; }

    public DocumentSigningOrder SigningOrder { get; set; } = DocumentSigningOrder.Parallel;

    public bool AllowDictateNextSigner { get; set; } = false;

    public bool TypedSignatureEnabled { get; set; } = true;

    public bool UploadSignatureEnabled { get; set; } = true;

    public bool DrawSignatureEnabled { get; set; } = true;

    public string Language { get; set; } = "en";

    public DocumentDistributionMethod DistributionMethod { get; set; } = DocumentDistributionMethod.Email;

    /// <summary>
    /// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
    /// </summary>
    public string? EmailSettings { get; set; }

    public string? EmailReplyTo { get; set; }

    public string? EmailId { get; set; }

    /// <summary>
    /// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    /// </summary>
    public string? EnvelopeExpirationPeriod { get; set; }

    /// <summary>
    /// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    /// </summary>
    public string? ReminderSettings { get; set; }

    public virtual Envelope? Envelope { get; set; }
}
