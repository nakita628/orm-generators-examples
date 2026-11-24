defmodule LangWatch.AnomalyAlert do
  @moduledoc """
  Persisted anomaly detection. Inserted by the anomaly-detection
  reactor (event-sourcing pipeline; modeled on PR #3351's
  alertTrigger pattern) when a rule's threshold trips while
  processing a freshly-arrived ActivityEventReceived event.
  Read by api.activityMonitor.recentAnomalies for the admin dashboard.

  Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
  window updates the existing row rather than creating a duplicate.
  Acknowledgement / resolution are admin actions on the row.

  Spec: specs/ai-gateway/governance/anomaly-detection.feature
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          severity: String.t(),
          rule_name: String.t(),
          rule_type: String.t(),
          trigger_window_start: DateTime.t(),
          trigger_window_end: DateTime.t(),
          trigger_spend_usd: Decimal.t() | nil,
          trigger_event_count: integer() | nil,
          detail: map() | list() | String.t() | number() | boolean(),
          state: String.t(),
          destination_status: map() | list() | String.t() | number() | boolean(),
          detected_at: DateTime.t(),
          acknowledged_at: DateTime.t() | nil,
          resolved_at: DateTime.t() | nil,
          organization: LangWatch.Organization.t() | nil,
          rule: LangWatch.AnomalyRule.t() | nil
        }

  schema "AnomalyAlert" do
    field(:id, :string, primary_key: true)
    # Denormalised from rule for fast WHERE filtering on the
    # dashboard query (no join needed for the common path).
    field(:severity, :string)
    field(:rule_name, :string, source: :ruleName)
    field(:rule_type, :string, source: :ruleType)
    # Window the trigger refers to (UTC).
    field(:trigger_window_start, LangWatch.PrismaDateTime, source: :triggerWindowStart)
    field(:trigger_window_end, LangWatch.PrismaDateTime, source: :triggerWindowEnd)
    # Numeric trigger context. Per-rule-type detail in `detail` JSONB.
    field(:trigger_spend_usd, :decimal, source: :triggerSpendUsd)
    field(:trigger_event_count, :integer, source: :triggerEventCount)
    # Per-rule-type extra context: {topActors: [...], baseline: ...,
    # observed: ..., timezone: ..., etc}.
    field(:detail, LangWatch.PrismaJson, default: %{})
    # 'open' | 'acknowledged' | 'resolved'.
    field(:state, :string, default: "open")
    # Per-destination dispatch outcome:
    # { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
    # slack: ..., pagerduty: ..., siem: ..., email: ... }
    field(:destination_status, LangWatch.PrismaJson, default: %{}, source: :destinationStatus)
    field(:detected_at, LangWatch.PrismaDateTime, autogenerate: true, source: :detectedAt)
    field(:acknowledged_at, LangWatch.PrismaDateTime, source: :acknowledgedAt)
    field(:resolved_at, LangWatch.PrismaDateTime, source: :resolvedAt)
    field(:organization_id, :string, source: :organizationId)
    field(:rule_id, :string, source: :ruleId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:rule, LangWatch.AnomalyRule, foreign_key: :rule_id, define_field: false, type: :string)
  end
end