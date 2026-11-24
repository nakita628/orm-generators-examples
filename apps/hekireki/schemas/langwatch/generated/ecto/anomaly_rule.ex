defmodule LangWatch.AnomalyRule do
  @moduledoc """
  Admin-defined anomaly detection rule. Evaluated server-side at
  ingest time (Option C - not in this slice yet) against
  gateway_activity_events. When triggered, an alert is dispatched
  to the destinations configured on the rule (Slack / SIEM / webhook
  / PagerDuty / email).

  This is a configuration entity - the actual rule evaluation +
  alert state lives in a future ClickHouse table
  (gateway_anomaly_alerts) and a server-side worker. Spec:
  specs/ai-gateway/governance/anomaly-rules.feature
  (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          scope: String.t(),
          scope_id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          severity: String.t(),
          rule_type: String.t(),
          threshold_config: map() | list() | String.t() | number() | boolean(),
          destination_config: map() | list() | String.t() | number() | boolean(),
          status: String.t(),
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil,
          created_by: LangWatch.User.t() | nil,
          alerts: [LangWatch.AnomalyAlert.t()]
        }

  schema "AnomalyRule" do
    field(:id, :string, primary_key: true)
    # 'organization' | 'team' | 'project' | 'source_type' | 'source'.
    # String not enum so new scopes ship without migrations.
    field(:scope, :string)
    # Org id / team id / project id / source-type string / source id
    # - narrows what the rule applies to.
    field(:scope_id, :string, source: :scopeId)
    field(:name, :string)
    field(:description, :string)
    # 'critical' | 'warning' | 'info'. Drives default destination
    # routing (e.g. critical → PagerDuty, warning → Slack only).
    field(:severity, :string)
    # Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
    # 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
    # | 'custom'. New types ship without migrations - evaluator
    # picks a per-type implementation from a code registry.
    field(:rule_type, :string, source: :ruleType)
    # Per-rule-type threshold shape. Examples:
    # spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
    # rate_limit: { windowSec, requestsThreshold }
    # after_hours: { startHour, endHour, timezone, requestsThreshold }
    field(:threshold_config, LangWatch.PrismaJson, default: %{}, source: :thresholdConfig)
    # Alert routing per destination. Shape:
    # { slack?: { channelId }, siem?: { endpointUrl, authHeader },
    # webhook?: { url, authHeader }, pagerduty?: { routingKey },
    # email?: { to: string[] } }
    # Empty map = use org defaults (when defaults exist; today no-op).
    field(:destination_config, LangWatch.PrismaJson, default: %{}, source: :destinationConfig)
    # 'active' (default - evaluator picks it up) | 'disabled' (admin
    # paused; rule rows kept for audit). 'archived' uses archivedAt
    # timestamp instead so we can filter `archivedAt IS NULL` cheaply.
    field(:status, :string, default: "active")
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:organization_id, :string, source: :organizationId)
    field(:created_by_id, :string, source: :createdById)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:created_by, LangWatch.User, foreign_key: :created_by_id, define_field: false, type: :string)
    has_many(:alerts, LangWatch.AnomalyAlert, foreign_key: :rule_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end