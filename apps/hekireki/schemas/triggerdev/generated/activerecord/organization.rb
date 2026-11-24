# maximumSchedulesLimit:: This is deprecated and will be removed in the future
# v2Enabled:: @deprecated
# v2MarqsEnabled:: @deprecated
# hasRequestedV3:: @deprecated
# webhookLimitsConfig:: Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults.
# maxSessionDuration:: Maximum session lifetime in seconds for members of this organization.
#   When set, caps each member's User.sessionDuration. The most restrictive cap across
#   all of a user's orgs (excluding nulls) wins.
# streamBasinName:: S2 basin that holds this org's realtime streams. Null until the
#   per-org basin has been provisioned (OSS / s2-lite installs leave
#   it null forever; reads fall back to the global basin env var).
#   Set once at provisioning time; retention is reconfigured in-place
#   when the org's plan changes.
# mfaRequired:: When true, members without MFA are locked out of the dashboard and CLI until they set it up.
# mfaStepUpActions:: Dashboard actions that need an MFA check within the step-up window (see MfaStepUpAction in the webapp).
class Organization < ApplicationRecord
  self.table_name = "Organization"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :maximumExecutionTimePerRunInMs, default: 900_000
  attribute :maximumConcurrencyLimit, default: 10
  attribute :maximumSchedulesLimit, default: 5
  attribute :runsEnabled, default: true
  attribute :v3Enabled, default: false
  attribute :v2Enabled, default: false
  attribute :v2MarqsEnabled, default: false
  attribute :hasRequestedV3, default: false
  attribute :maximumProjectCount, default: 25
  attribute :mfaRequired, default: false
  attribute :mfaStepUpActions, default: -> { [] }
  attribute :mfaStepUpWindowMinutes, default: 5
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :supportAccessMode, { allow: "ALLOW", requires_request: "REQUIRES_REQUEST" }, default: :allow, validate: true

  validates :slug, presence: true, uniqueness: true
  validates :title, presence: true

  has_many :environments, class_name: "RuntimeEnvironment", foreign_key: "organizationId", inverse_of: :organization,
                          dependent: :destroy
  has_many :projects, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :members, class_name: "OrgMember", foreign_key: "organizationId", inverse_of: :organization,
                     dependent: :destroy
  has_many :invites, class_name: "OrgMemberInvite", foreign_key: "organizationId", inverse_of: :organization,
                     dependent: :destroy
  has_many :organization_integrations, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :organization_access_tokens, foreign_key: "organizationId", inverse_of: :organization,
                                        dependent: :restrict_with_error
  has_many :worker_groups, class_name: "WorkerInstanceGroup", foreign_key: "organizationId", inverse_of: :organization,
                           dependent: :destroy
  has_many :worker_instances, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :github_app_installations, foreign_key: "organizationId", inverse_of: :organization,
                                      dependent: :restrict_with_error
  has_many :customer_queries, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :metrics_dashboards, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :prompts, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :platform_notifications, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :error_group_states, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
  has_many :support_access_requests, foreign_key: "organizationId", inverse_of: :organization, dependent: :destroy
end
