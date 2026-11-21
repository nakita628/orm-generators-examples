# One upgrade diagnostics bundle a self-hosted install sent over Connect, on
# LangWatch Cloud. The bytes live in object storage under
# `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
# points at them and feeds Cloud admin and the Slack signal. Plain column
# refs, no relation, as the usage report rows.
#
# instanceId:: The `SelfHostedInstance.instanceId` the licence caller presented.
# trigger:: `operator` or `automatic`.
# expiresAt:: When the lifecycle rule removes the bundle; the row is listed until then.
# slackAnnouncedAt:: Set on the report that posted this step and release's one Slack message.
class SelfHostedInstanceDiagnostic < ApplicationRecord
  self.table_name = "SelfHostedInstanceDiagnostic"

  validates :instanceId, presence: true
  validates :stepId, presence: true
  validates :attempt, presence: true
  validates :trigger, presence: true
  validates :byteLength, presence: true
  validates :objectKey, presence: true
  validates :expiresAt, presence: true
end
