#nullable enable

namespace Documenso.Models;

/// <summary>
/// @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
/// </summary>
public partial class TeamGlobalSettings
{
    public string Id { get; set; } = null!;

    public virtual Team? Team { get; set; }

    public DocumentVisibility? DocumentVisibility { get; set; }

    public string? DocumentLanguage { get; set; }

    public string? DocumentTimezone { get; set; }

    public string? DocumentDateFormat { get; set; }

    public bool? DelegateDocumentOwnership { get; set; }

    public bool? IncludeSenderDetails { get; set; }

    public bool? IncludeSigningCertificate { get; set; }

    public bool? IncludeAuditLog { get; set; }

    public bool? TypedSignatureEnabled { get; set; }

    public bool? UploadSignatureEnabled { get; set; }

    public bool? DrawSignatureEnabled { get; set; }

    /// <summary>
    /// [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
    /// </summary>
    public string? DefaultRecipients { get; set; }

    public string? EmailId { get; set; }

    public virtual OrganisationEmail? Email { get; set; }

    public string? EmailReplyTo { get; set; }

    /// <summary>
    /// [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema)
    /// </summary>
    public string? EmailDocumentSettings { get; set; }

    public bool? BrandingEnabled { get; set; }

    public string? BrandingLogo { get; set; }

    public string? BrandingUrl { get; set; }

    public string? BrandingCompanyDetails { get; set; }

    /// <summary>
    /// [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
    /// </summary>
    public string? BrandingColors { get; set; }

    public string? BrandingCss { get; set; }

    /// <summary>
    /// [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    /// </summary>
    public string? EnvelopeExpirationPeriod { get; set; }

    /// <summary>
    /// [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    /// </summary>
    public string? ReminderSettings { get; set; }

    public bool? AiFeaturesEnabled { get; set; }
}
