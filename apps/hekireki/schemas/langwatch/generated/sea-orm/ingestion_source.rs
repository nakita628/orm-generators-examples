use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-platform fleet configuration that connects a closed SaaS
/// product's audit / OTel / S3 stream to LangWatch's Activity
/// Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
/// + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).
///
/// One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
/// "Acme Workato production"). The runtime ingests events on whatever
/// cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
/// pull for poll-based admin APIs (Tier C). Per-source-type adapters
/// normalise platform-specific shapes into OCSF + AOS before they
/// land in trace_summaries with SourceType + SourceId tags.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IngestionSource")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// Optional team-scope. null = org-wide.
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<String>,
    /// 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
    /// 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
    /// String not enum so adapters can be added without migrations - the
    /// supported set lives in code (activity-monitor adapter registry).
    #[sea_orm(column_name = "sourceType")]
    pub source_type: String,
    /// Human-friendly label for the admin UI. Unique per-org so two
    /// "Cowork" sources can coexist with different names.
    pub name: String,
    pub description: Option<String>,
    /// HMAC bearer token the upstream operator pastes into the platform
    /// (or LangWatch sends in pull mode). Stored encrypted at rest via
    /// the same pepper-derived KEK as VirtualKey secrets. Auto-generated
    /// on create, rotatable via the admin UI with a 24h grace window.
    #[sea_orm(column_name = "ingestSecretHash")]
    pub ingest_secret_hash: String,
    /// Per-source-type configuration. Shape is platform-specific:
    /// - otel_generic: { allowedSourceTypeLabel?: string }
    /// - claude_cowork: { workspaceId: string }
    /// - workato: { sharedSecretLastFour: string }
    /// - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
    /// - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
    /// - claude_compliance: { workspaceApiKeyHash, pollEverySec }
    /// - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
    #[sea_orm(column_name = "parserConfig", default_value = "{}")]
    pub parser_config: Json,
    /// For pull-based sources, persisted between polls so the next poll
    /// resumes where the last left off (timestamp / cursor / file marker).
    #[sea_orm(column_name = "pollerCursor")]
    pub poller_cursor: Option<Json>,
    /// Phase 10 - number of consecutive errors observed by the puller
    /// worker. Reset to 0 on a successful run. Drives admin-UI banners
    /// ("Microsoft authentication failed - re-authorize") + back-off
    /// scheduling when the source has been broken for a while.
    /// Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
    #[sea_orm(column_name = "errorCount", default_value = 0)]
    pub error_count: i32,
    /// Phase 10 - cron schedule for pull-mode sources. Validated by the
    /// adapter's pullConfig schema; mirrored here for fast worker
    /// dispatch. Null for push-mode sources.
    #[sea_orm(column_name = "pullSchedule")]
    pub pull_schedule: Option<String>,
    /// Status flag. 'active' (default), 'disabled' (admin paused), or
    /// 'awaiting_first_event' (created but never received).
    #[sea_orm(default_value = "awaiting_first_event")]
    pub status: String,
    /// ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
    /// (Databricks Genie). Null = don't route. Same shape as
    /// VirtualKey.traceProjectId: a destination is NOT a scope row and grants
    /// no access to the source; it only decides where routed conversation
    /// traces land. Must be a project of this org (validated at write time).
    /// Freely editable; already-routed history stays where it landed.
    /// NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
    #[sea_orm(column_name = "traceProjectId")]
    pub trace_project_id: Option<String>,
    /// The account the provider itself named when this connection was saved -
    /// an OpenAI or Anthropic organization id, read once per save and compared
    /// so a second connection cannot read one account's spend twice. NOT
    /// derived from the administrator key: nothing worked out from a customer
    /// credential is stored here or anywhere else, because the account id is
    /// not a secret and cannot be turned back into a key.
    /// Nullable because every connection that predates the guard has none, and
    /// because Azure and Power Platform name what they read in parserConfig
    /// instead. "Unknown account" has to be a storable state.
    /// Spec: specs/ai-gateway/governance/ingestion-sources.feature
    #[sea_orm(column_name = "providerAccountId")]
    pub provider_account_id: Option<String>,
    /// Last event timestamp observed (any direction). Surfaces in the
    /// admin UI as "last event 2 minutes ago".
    #[sea_orm(column_name = "lastEventAt")]
    pub last_event_at: Option<DateTime>,
    /// ADR-128 - when a pull run last reached the provider without erroring,
    /// mirrored from the ingestionPullRunStatus projection. Deliberately NOT
    /// lastEventAt: an empty period is a working puller, and a run that
    /// delivered data and then failed is not. Drives the health badge and
    /// the "no data since" cut-off, below which a day's spend is unknown
    /// rather than zero.
    /// Spec: specs/governance/ingestion-source-health.feature
    #[sea_orm(column_name = "lastSuccessAt")]
    pub last_success_at: Option<DateTime>,
    /// ADR-128 - how far the last run actually read, and whether it reached the
    /// end. Mirrored from the ingestionPullRunStatus projection, in the same way
    /// and for the same reason as lastSuccessAt above: every reader of this
    /// surface selects from this row, and two of them render a LIST of sources,
    /// so reaching a per-run projection would be one query per row on a screen
    /// built to show many at once.
    ///
    /// These two are the only evidence separating a source that truncated once
    /// and then finished from one that is stuck. A stuck source's run clock
    /// advances on every attempt, so the clock cannot be that evidence and is
    /// deliberately not stored here - only the point the run REACHED.
    ///
    /// Completeness is "complete" or "truncated". Null on both is unknown, which
    /// is what every row written before this existed reads as, and unknown must
    /// never render as either answer.
    /// Spec: specs/governance/ingestion-source-health.feature
    #[sea_orm(column_name = "lastReadThroughAt")]
    pub last_read_through_at: Option<DateTime>,
    #[sea_orm(column_name = "lastRunCompleteness")]
    pub last_run_completeness: Option<String>,
    /// ADR-088 - the window this source read while pulled cost recording was
    /// off for the organization, so its spend was dropped rather than priced.
    /// Both null = nothing dropped. The pull cursor advances whether or not the
    /// money path is live (audit-only is a supported mode), so without this the
    /// loss is unrecoverable AND invisible: those days read as no spend rather
    /// than as unknown. Widened by any run that drops a priced item, cleared
    /// only when a later run re-reads back across the whole window.
    /// Spec: specs/governance/pulled-usage-cost-gap.feature
    #[sea_orm(column_name = "unpricedUsageSince")]
    pub unpriced_usage_since: Option<DateTime>,
    #[sea_orm(column_name = "unpricedUsageThrough")]
    pub unpriced_usage_through: Option<DateTime>,
    #[sea_orm(column_name = "archivedAt")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "SetNull"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    CreatedBy,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
