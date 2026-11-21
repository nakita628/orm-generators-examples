defmodule LangWatch.IngestionSource do
  @moduledoc """
  Per-platform fleet configuration that connects a closed SaaS
  product's audit / OTel / S3 stream to LangWatch's Activity
  Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
  + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).

  One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
  "Acme Workato production"). The runtime ingests events on whatever
  cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
  pull for poll-based admin APIs (Tier C). Per-source-type adapters
  normalise platform-specific shapes into OCSF + AOS before they
  land in trace_summaries with SourceType + SourceId tags.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          source_type: String.t(),
          name: String.t(),
          description: String.t() | nil,
          ingest_secret_hash: String.t(),
          parser_config: map(),
          poller_cursor: map() | nil,
          error_count: integer(),
          pull_schedule: String.t() | nil,
          status: String.t(),
          trace_project_id: String.t() | nil,
          provider_account_id: String.t() | nil,
          last_event_at: DateTime.t() | nil,
          last_success_at: DateTime.t() | nil,
          last_read_through_at: DateTime.t() | nil,
          last_run_completeness: String.t() | nil,
          unpriced_usage_since: DateTime.t() | nil,
          unpriced_usage_through: DateTime.t() | nil,
          archived_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil,
          team: LangWatch.Team.t() | nil,
          created_by: LangWatch.User.t() | nil
        }

  schema "IngestionSource" do
    field(:id, :string, primary_key: true)
    # 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
    # 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
    # String not enum so adapters can be added without migrations - the
    # supported set lives in code (activity-monitor adapter registry).
    field(:source_type, :string, source: :sourceType)
    # Human-friendly label for the admin UI. Unique per-org so two
    # "Cowork" sources can coexist with different names.
    field(:name, :string)
    field(:description, :string)
    # HMAC bearer token the upstream operator pastes into the platform
    # (or LangWatch sends in pull mode). Stored encrypted at rest via
    # the same pepper-derived KEK as VirtualKey secrets. Auto-generated
    # on create, rotatable via the admin UI with a 24h grace window.
    field(:ingest_secret_hash, :string, source: :ingestSecretHash)
    # Per-source-type configuration. Shape is platform-specific:
    # - otel_generic: { allowedSourceTypeLabel?: string }
    # - claude_cowork: { workspaceId: string }
    # - workato: { sharedSecretLastFour: string }
    # - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
    # - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
    # - claude_compliance: { workspaceApiKeyHash, pollEverySec }
    # - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
    field(:parser_config, :map, default: %{}, source: :parserConfig)
    # For pull-based sources, persisted between polls so the next poll
    # resumes where the last left off (timestamp / cursor / file marker).
    field(:poller_cursor, :map, source: :pollerCursor)
    # Phase 10 - number of consecutive errors observed by the puller
    # worker. Reset to 0 on a successful run. Drives admin-UI banners
    # ("Microsoft authentication failed - re-authorize") + back-off
    # scheduling when the source has been broken for a while.
    # Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
    field(:error_count, :integer, default: 0, source: :errorCount)
    # Phase 10 - cron schedule for pull-mode sources. Validated by the
    # adapter's pullConfig schema; mirrored here for fast worker
    # dispatch. Null for push-mode sources.
    field(:pull_schedule, :string, source: :pullSchedule)
    # Status flag. 'active' (default), 'disabled' (admin paused), or
    # 'awaiting_first_event' (created but never received).
    field(:status, :string, default: "awaiting_first_event")
    # ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
    # (Databricks Genie). Null = don't route. Same shape as
    # VirtualKey.traceProjectId: a destination is NOT a scope row and grants
    # no access to the source; it only decides where routed conversation
    # traces land. Must be a project of this org (validated at write time).
    # Freely editable; already-routed history stays where it landed.
    # NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
    field(:trace_project_id, :string, source: :traceProjectId)
    # The account the provider itself named when this connection was saved -
    # an OpenAI or Anthropic organization id, read once per save and compared
    # so a second connection cannot read one account's spend twice. NOT
    # derived from the administrator key: nothing worked out from a customer
    # credential is stored here or anywhere else, because the account id is
    # not a secret and cannot be turned back into a key.
    # Nullable because every connection that predates the guard has none, and
    # because Azure and Power Platform name what they read in parserConfig
    # instead. "Unknown account" has to be a storable state.
    # Spec: specs/ai-gateway/governance/ingestion-sources.feature
    field(:provider_account_id, :string, source: :providerAccountId)
    # Last event timestamp observed (any direction). Surfaces in the
    # admin UI as "last event 2 minutes ago".
    field(:last_event_at, LangWatch.PrismaDateTime, source: :lastEventAt)
    # ADR-128 - when a pull run last reached the provider without erroring,
    # mirrored from the ingestionPullRunStatus projection. Deliberately NOT
    # lastEventAt: an empty period is a working puller, and a run that
    # delivered data and then failed is not. Drives the health badge and
    # the "no data since" cut-off, below which a day's spend is unknown
    # rather than zero.
    # Spec: specs/governance/ingestion-source-health.feature
    field(:last_success_at, LangWatch.PrismaDateTime, source: :lastSuccessAt)
    # ADR-128 - how far the last run actually read, and whether it reached the
    # end. Mirrored from the ingestionPullRunStatus projection, in the same way
    # and for the same reason as lastSuccessAt above: every reader of this
    # surface selects from this row, and two of them render a LIST of sources,
    # so reaching a per-run projection would be one query per row on a screen
    # built to show many at once.
    #
    # These two are the only evidence separating a source that truncated once
    # and then finished from one that is stuck. A stuck source's run clock
    # advances on every attempt, so the clock cannot be that evidence and is
    # deliberately not stored here - only the point the run REACHED.
    #
    # Completeness is "complete" or "truncated". Null on both is unknown, which
    # is what every row written before this existed reads as, and unknown must
    # never render as either answer.
    # Spec: specs/governance/ingestion-source-health.feature
    field(:last_read_through_at, LangWatch.PrismaDateTime, source: :lastReadThroughAt)
    field(:last_run_completeness, :string, source: :lastRunCompleteness)
    # ADR-088 - the window this source read while pulled cost recording was
    # off for the organization, so its spend was dropped rather than priced.
    # Both null = nothing dropped. The pull cursor advances whether or not the
    # money path is live (audit-only is a supported mode), so without this the
    # loss is unrecoverable AND invisible: those days read as no spend rather
    # than as unknown. Widened by any run that drops a priced item, cleared
    # only when a later run re-reads back across the whole window.
    # Spec: specs/governance/pulled-usage-cost-gap.feature
    field(:unpriced_usage_since, LangWatch.PrismaDateTime, source: :unpricedUsageSince)
    field(:unpriced_usage_through, LangWatch.PrismaDateTime, source: :unpricedUsageThrough)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:organization_id, :string, source: :organizationId)
    # Optional team-scope. null = org-wide.
    field(:team_id, :string, source: :teamId)
    field(:created_by_id, :string, source: :createdById)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    belongs_to(:team, LangWatch.Team, foreign_key: :team_id, define_field: false, type: :string)
    belongs_to(:created_by, LangWatch.User, foreign_key: :created_by_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end