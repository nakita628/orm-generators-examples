#nullable enable

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
/// </summary>
public partial class OrganisationGlobalSettings
{
    public string Id { get; set; } = null!;

    public virtual Organisation? Organisation { get; set; }

    public DocumentVisibility DocumentVisibility { get; set; } = DocumentVisibility.Everyone;

    public string DocumentLanguage { get; set; } = "en";

    public bool IncludeSenderDetails { get; set; } = true;

    public bool IncludeSigningCertificate { get; set; } = true;

    public bool IncludeAuditLog { get; set; } = false;

    public string? DocumentTimezone { get; set; }

    public string DocumentDateFormat { get; set; } = "yyyy-MM-dd hh:mm a";

    public bool DelegateDocumentOwnership { get; set; } = false;

    public bool TypedSignatureEnabled { get; set; } = true;

    public bool UploadSignatureEnabled { get; set; } = true;

    public bool DrawSignatureEnabled { get; set; } = true;

    /// <summary>
    /// [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
    /// </summary>
    public string? DefaultRecipients { get; set; }

    public string? EmailId { get; set; }

    public virtual OrganisationEmail? Email { get; set; }

    public string? EmailReplyTo { get; set; }

    /// <summary>
    /// [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
    /// </summary>
    public string EmailDocumentSettings { get; set; } = null!;

    public bool BrandingEnabled { get; set; } = false;

    public string BrandingLogo { get; set; } = "";

    public string BrandingUrl { get; set; } = "";

    public string BrandingCompanyDetails { get; set; } = "";

    /// <summary>
    /// [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
    /// </summary>
    public string? BrandingColors { get; set; }

    public string BrandingCss { get; set; } = "";

    /// <summary>
    /// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    /// </summary>
    public string? EnvelopeExpirationPeriod { get; set; }

    /// <summary>
    /// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    /// </summary>
    public string? ReminderSettings { get; set; }

    public bool AiFeaturesEnabled { get; set; } = false;
}
