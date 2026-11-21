# @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
#
# defaultRecipients:: [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
# emailDocumentSettings:: [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema)
# brandingColors:: [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
# envelopeExpirationPeriod:: [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
# reminderSettings:: [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
class TeamGlobalSettings < ApplicationRecord
  self.table_name = "TeamGlobalSettings"

  enum :documentVisibility, { everyone: "EVERYONE", manager_and_above: "MANAGER_AND_ABOVE", admin: "ADMIN" },
       validate: { allow_nil: true }

  belongs_to :email, class_name: "OrganisationEmail", foreign_key: "emailId", optional: true,
                     inverse_of: :team_global_settings
  has_one :team, foreign_key: "teamGlobalSettingsId", inverse_of: :team_global_settings, dependent: :destroy
end
