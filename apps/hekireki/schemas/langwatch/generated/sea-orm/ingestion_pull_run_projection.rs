use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Rebuildable per-source pull checkpoint and operator-facing run status.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "IngestionPullRunProjection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "sourceId")]
    pub source_id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    #[sea_orm(default_value = false)]
    pub enabled: bool,
    pub cron: Option<String>,
    pub cursor: Option<String>,
    #[sea_orm(column_name = "lastRunAt")]
    pub last_run_at: Option<f64>,
    #[sea_orm(column_name = "lastRunOutcome")]
    pub last_run_outcome: Option<String>,
    #[sea_orm(column_name = "lastRunEventCount", default_value = 0)]
    pub last_run_event_count: i32,
    #[sea_orm(column_name = "lastRunError")]
    pub last_run_error: Option<String>,
    #[sea_orm(column_name = "lastRunErrorCode")]
    pub last_run_error_code: Option<String>,
    #[sea_orm(column_name = "consecutiveErrors", default_value = 0)]
    pub consecutive_errors: i32,
    /// When a run last completed without erroring. Backfilled by replay, which
    /// is why the projection version moved with this column.
    #[sea_orm(column_name = "lastSuccessAt")]
    pub last_success_at: Option<f64>,
    /// Identity of the run whose outcome this row reflects. Outcomes from a
    /// superseded run arrive with a smaller value and are ignored, so a late
    /// completion cannot regress Cursor or the status fields.
    #[sea_orm(column_name = "lastRunScheduledFor")]
    pub last_run_scheduled_for: Option<f64>,
    /// ADR-128 - how far the last run read, in epoch milliseconds, and whether
    /// it reached the end. The fold writes both; the IngestionSource row mirrors
    /// them for the screens. Null read-through means no run has reported one.
    /// Null completeness reads as unknown, never as either answer, which is what
    /// every row written before this existed holds.
    ///
    /// Nullable on purpose, so replay is not required to make this migration
    /// safe: a row that has never seen a run carrying these stays null and every
    /// reader already handles unknown.
    /// Spec: specs/governance/ingestion-source-health.feature
    #[sea_orm(column_name = "lastReadThroughAt")]
    pub last_read_through_at: Option<f64>,
    #[sea_orm(column_name = "lastRunCompleteness")]
    pub last_run_completeness: Option<String>,
    /// The last time somebody asked this source to enumerate its agents, and how
    /// that ask ended. A listing is a different scope on the same connection as
    /// the pull above: a credential that cannot enumerate a directory still pulls
    /// cost perfectly, so no health, coverage or completeness reader may treat a
    /// refusal here as the source failing to pull.
    ///
    /// Count is set ONLY when the outcome is listed, and stays null on a refusal.
    /// Zero is a real answer meaning the provider returned an empty list, so a
    /// zero default would manufacture that answer for a source that refused.
    /// Reason is set only on a refusal and is a plain string, never an enum and
    /// never a foreign key, because the log outlives the vocabulary. Status is
    /// nullable even when refused, because a refusal we generated ourselves has
    /// no HTTP status, and it is never rendered to a customer.
    ///
    /// The two `...ListingAt` columns hold the instant the ASK was accepted, not
    /// the instant its answer was recorded, and the fold orders outcomes by them
    /// so a superseded listing cannot overwrite the one that replaced it. The
    /// name predates that job and is kept because renaming a column is a
    /// migration; the fold's own field docs are the statement of record.
    #[sea_orm(column_name = "lastAgentsListingAt")]
    pub last_agents_listing_at: Option<f64>,
    #[sea_orm(column_name = "lastAgentsListingOutcome")]
    pub last_agents_listing_outcome: Option<String>,
    #[sea_orm(column_name = "lastAgentsListingCount")]
    pub last_agents_listing_count: Option<i32>,
    #[sea_orm(column_name = "lastAgentsListingReason")]
    pub last_agents_listing_reason: Option<String>,
    #[sea_orm(column_name = "lastAgentsListingStatus")]
    pub last_agents_listing_status: Option<i32>,
    /// The same for people, in its own set of columns rather than sharing the
    /// agents set with a kind discriminator. A source can be refused for people
    /// and fine for agents, and one shared set holds only the most recent listing
    /// of either kind, so syncing people would erase an agents refusal and the
    /// agents page would go back to not knowing after briefly knowing.
    ///
    /// LastPeopleDirectoryCount is everyone the provider's directory named.
    /// LastPeopleWithheldCount is how many of those this deployment does not hold
    /// because erasure suppression removed them: a SUBSET of the directory count,
    /// never an addition to it. Adding the two counts the same people twice.
    /// Subtracting is the valid arithmetic, and the surviving total is that
    /// subtraction rather than a third column that could disagree.
    ///
    /// The withheld count is safe to show as a CURRENT figure and never as a
    /// series: it moving from zero to one at a known moment says an erasure
    /// happened then, which on a small tenant identifies the person as surely as
    /// a name would. So no trend line, no history drawer, no per-run export, and
    /// never beside a per-person list.
    #[sea_orm(column_name = "lastPeopleListingAt")]
    pub last_people_listing_at: Option<f64>,
    #[sea_orm(column_name = "lastPeopleListingOutcome")]
    pub last_people_listing_outcome: Option<String>,
    #[sea_orm(column_name = "lastPeopleDirectoryCount")]
    pub last_people_directory_count: Option<i32>,
    #[sea_orm(column_name = "lastPeopleWithheldCount")]
    pub last_people_withheld_count: Option<i32>,
    #[sea_orm(column_name = "lastPeopleListingReason")]
    pub last_people_listing_reason: Option<String>,
    #[sea_orm(column_name = "lastPeopleListingStatus")]
    pub last_people_listing_status: Option<i32>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: f64,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: f64,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: f64,
    #[sea_orm(column_name = "acceptedAt")]
    pub accepted_at: f64,
    #[sea_orm(column_name = "lastEventId")]
    pub last_event_id: String,
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
