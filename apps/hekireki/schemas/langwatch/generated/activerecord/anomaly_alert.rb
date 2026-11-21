# Persisted anomaly detection. Inserted by the anomaly-detection
# reactor (event-sourcing pipeline; modeled on PR #3351's
# alertTrigger pattern) when a rule's threshold trips while
# processing a freshly-arrived ActivityEventReceived event.
# Read by api.activityMonitor.recentAnomalies for the admin dashboard.
#
# Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
# window updates the existing row rather than creating a duplicate.
# Acknowledgement / resolution are admin actions on the row.
#
# Spec: specs/ai-gateway/governance/anomaly-detection.feature
#
# severity:: Denormalised from rule for fast WHERE filtering on the
#   dashboard query (no join needed for the common path).
# triggerWindowStart:: Window the trigger refers to (UTC).
# triggerSpendUsd:: Numeric trigger context. Per-rule-type detail in `detail` JSONB.
# detail:: Per-rule-type extra context: {topActors: [...], baseline: ...,
#   observed: ..., timezone: ..., etc}.
# state:: 'open' | 'acknowledged' | 'resolved'.
# destinationStatus:: Per-destination dispatch outcome:
#   { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
#   slack: ..., pagerduty: ..., siem: ..., email: ... }
class AnomalyAlert < ApplicationRecord
  self.table_name = "AnomalyAlert"

  attribute :id, default: -> { Nanoid.generate }
  attribute :detail, default: -> { JSON.parse("{}") }
  attribute :state, default: "open"
  attribute :destinationStatus, default: -> { JSON.parse("{}") }

  validates :severity, presence: true
  validates :ruleName, presence: true
  validates :ruleType, presence: true
  validates :triggerWindowStart, presence: true, uniqueness: { scope: :ruleId }
  validates :triggerWindowEnd, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :anomaly_alerts
  belongs_to :rule, class_name: "AnomalyRule", foreign_key: "ruleId", inverse_of: :alerts
end
