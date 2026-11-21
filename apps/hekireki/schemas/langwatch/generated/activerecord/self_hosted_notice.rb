# A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
# report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
# tenant: every install receives every published notice that has not expired.
#
# severity:: `info`, `warning` or `security`.
# url:: An https link to read more; null where the notice names none.
# fixedIn:: The release that fixes what the notice is about; an install on it or newer hides it.
# status:: `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent.
# publishedBy:: The staff member who published it, by user id.
class SelfHostedNotice < ApplicationRecord
  self.table_name = "SelfHostedNotice"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :severity, presence: true
  validates :title, presence: true
  validates :body, presence: true
  validates :status, presence: true
end
