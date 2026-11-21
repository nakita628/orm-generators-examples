# Junction table linking responses to quotas.
# Tracks which responses counted towards specific quotas and their status.
#
# @property response - The response that counted towards the quota
# @property quota - The quota this response counted towards
# @property status - Whether the response was screened in or out
class ResponseQuotaLink < ApplicationRecord
  self.table_name = "ResponseQuotaLink"
  self.primary_key = %w[responseId quotaId]

  enum :status, { screened_in: "screenedIn", screened_out: "screenedOut" }, validate: true

  belongs_to :response, foreign_key: "responseId", inverse_of: :quota_links
  belongs_to :quota, class_name: "SurveyQuota", foreign_key: "quotaId", inverse_of: :quota_links
end
