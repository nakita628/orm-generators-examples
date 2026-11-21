# Admin-defined anomaly detection rule. Evaluated server-side at
# ingest time (Option C - not in this slice yet) against
# gateway_activity_events. When triggered, an alert is dispatched
# to the destinations configured on the rule (Slack / SIEM / webhook
# / PagerDuty / email).
#
# This is a configuration entity - the actual rule evaluation +
# alert state lives in a future ClickHouse table
# (gateway_anomaly_alerts) and a server-side worker. Spec:
# specs/ai-gateway/governance/anomaly-rules.feature
# (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
#
# scope:: 'organization' | 'team' | 'project' | 'source_type' | 'source'.
#   String not enum so new scopes ship without migrations.
# scopeId:: Org id / team id / project id / source-type string / source id
#   - narrows what the rule applies to.
# severity:: 'critical' | 'warning' | 'info'. Drives default destination
#   routing (e.g. critical → PagerDuty, warning → Slack only).
# ruleType:: Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
#   'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
#   | 'custom'. New types ship without migrations - evaluator
#   picks a per-type implementation from a code registry.
# thresholdConfig:: Per-rule-type threshold shape. Examples:
#   spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
#   rate_limit: { windowSec, requestsThreshold }
#   after_hours: { startHour, endHour, timezone, requestsThreshold }
# destinationConfig:: Alert routing per destination. Shape:
#   { slack?: { channelId }, siem?: { endpointUrl, authHeader },
#   webhook?: { url, authHeader }, pagerduty?: { routingKey },
#   email?: { to: string[] } }
#   Empty map = use org defaults (when defaults exist; today no-op).
# status:: 'active' (default - evaluator picks it up) | 'disabled' (admin
#   paused; rule rows kept for audit). 'archived' uses archivedAt
#   timestamp instead so we can filter `archivedAt IS NULL` cheaply.
class AnomalyRule < ApplicationRecord
  self.table_name = "AnomalyRule"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :thresholdConfig, default: -> { JSON.parse("{}") }
  attribute :destinationConfig, default: -> { JSON.parse("{}") }
  attribute :status, default: "active"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :scope, presence: true
  validates :scopeId, presence: true
  validates :name, presence: true, uniqueness: { scope: :organizationId }
  validates :severity, presence: true
  validates :ruleType, presence: true

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :anomaly_rules
  belongs_to :created_by, class_name: "User", foreign_key: "createdById", optional: true,
                          inverse_of: :anomaly_rules_created
  has_many :alerts, class_name: "AnomalyAlert", foreign_key: "ruleId", inverse_of: :rule, dependent: :destroy
end
