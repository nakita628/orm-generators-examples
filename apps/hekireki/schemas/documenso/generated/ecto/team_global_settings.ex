defmodule Documenso.TeamGlobalSettings do
  @moduledoc """
  @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          document_visibility: atom() | nil,
          document_language: String.t() | nil,
          document_timezone: String.t() | nil,
          document_date_format: String.t() | nil,
          delegate_document_ownership: boolean() | nil,
          include_sender_details: boolean() | nil,
          include_signing_certificate: boolean() | nil,
          include_audit_log: boolean() | nil,
          typed_signature_enabled: boolean() | nil,
          upload_signature_enabled: boolean() | nil,
          draw_signature_enabled: boolean() | nil,
          default_recipients: map() | list() | String.t() | number() | boolean() | nil,
          email_reply_to: String.t() | nil,
          email_document_settings: map() | list() | String.t() | number() | boolean() | nil,
          branding_enabled: boolean() | nil,
          branding_logo: String.t() | nil,
          branding_url: String.t() | nil,
          branding_company_details: String.t() | nil,
          branding_colors: map() | list() | String.t() | number() | boolean() | nil,
          branding_css: String.t() | nil,
          envelope_expiration_period: map() | list() | String.t() | number() | boolean() | nil,
          reminder_settings: map() | list() | String.t() | number() | boolean() | nil,
          ai_features_enabled: boolean() | nil,
          email: Documenso.OrganisationEmail.t() | nil,
          team: Documenso.Team.t() | nil
        }

  schema "TeamGlobalSettings" do
    field(:id, :string, primary_key: true)
    field(:document_visibility, Ecto.Enum, values: [:EVERYONE, :MANAGER_AND_ABOVE, :ADMIN], source: :documentVisibility)
    field(:document_language, :string, source: :documentLanguage)
    field(:document_timezone, :string, source: :documentTimezone)
    field(:document_date_format, :string, source: :documentDateFormat)
    field(:delegate_document_ownership, :boolean, source: :delegateDocumentOwnership)
    field(:include_sender_details, :boolean, source: :includeSenderDetails)
    field(:include_signing_certificate, :boolean, source: :includeSigningCertificate)
    field(:include_audit_log, :boolean, source: :includeAuditLog)
    field(:typed_signature_enabled, :boolean, source: :typedSignatureEnabled)
    field(:upload_signature_enabled, :boolean, source: :uploadSignatureEnabled)
    field(:draw_signature_enabled, :boolean, source: :drawSignatureEnabled)
    # [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
    field(:default_recipients, Documenso.PrismaJson, source: :defaultRecipients)
    field(:email_reply_to, :string, source: :emailReplyTo)
    # [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema)
    field(:email_document_settings, Documenso.PrismaJson, source: :emailDocumentSettings)
    field(:branding_enabled, :boolean, source: :brandingEnabled)
    field(:branding_logo, :string, source: :brandingLogo)
    field(:branding_url, :string, source: :brandingUrl)
    field(:branding_company_details, :string, source: :brandingCompanyDetails)
    # [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
    field(:branding_colors, Documenso.PrismaJson, source: :brandingColors)
    field(:branding_css, :string, source: :brandingCss)
    # [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
    field(:envelope_expiration_period, Documenso.PrismaJson, source: :envelopeExpirationPeriod)
    # [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
    field(:reminder_settings, Documenso.PrismaJson, source: :reminderSettings)
    field(:ai_features_enabled, :boolean, source: :aiFeaturesEnabled)
    field(:email_id, :string, source: :emailId)
    belongs_to(:email, Documenso.OrganisationEmail, foreign_key: :email_id, define_field: false, type: :string)
    has_one(:team, Documenso.Team, foreign_key: :team_global_settings_id, references: :id)
  end
end