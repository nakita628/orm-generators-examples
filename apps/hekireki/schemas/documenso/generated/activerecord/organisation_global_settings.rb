# @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
#
# defaultRecipients:: [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
# emailDocumentSettings:: [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
# brandingColors:: [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
# envelopeExpirationPeriod:: [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
# reminderSettings:: [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
class OrganisationGlobalSettings < ApplicationRecord
  self.table_name = "OrganisationGlobalSettings"

  attribute :documentLanguage, default: "en"
  attribute :includeSenderDetails, default: true
  attribute :includeSigningCertificate, default: true
  attribute :includeAuditLog, default: false
  attribute :documentDateFormat, default: "yyyy-MM-dd hh:mm a"
  attribute :delegateDocumentOwnership, default: false
  attribute :typedSignatureEnabled, default: true
  attribute :uploadSignatureEnabled, default: true
  attribute :drawSignatureEnabled, default: true
  attribute :brandingEnabled, default: false
  attribute :brandingLogo, default: ""
  attribute :brandingUrl, default: ""
  attribute :brandingCompanyDetails, default: ""
  attribute :brandingCss, default: ""
  attribute :aiFeaturesEnabled, default: false

  enum :documentVisibility, { everyone: "EVERYONE", manager_and_above: "MANAGER_AND_ABOVE", admin: "ADMIN" },
       default: :everyone, validate: true

  validates :emailDocumentSettings, exclusion: { in: [ nil ] }

  belongs_to :email, class_name: "OrganisationEmail", foreign_key: "emailId", optional: true,
                     inverse_of: :organisation_global_settings
  has_one :organisation, foreign_key: "organisationGlobalSettingsId", inverse_of: :organisation_global_settings,
                         dependent: :restrict_with_error
end
