use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
/// and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
/// organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingProjectUsageView")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    /// The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
    /// or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
    #[sea_orm(primary_key, auto_increment = false)]
    pub owner: String,
    /// The project's usage as the fold folded it (NurturingProjectUsageState).
    pub folded: Json,
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    #[sea_orm(column_name = "appliedEventIds")]
    pub applied_event_ids: Vec<String>,
    /// A row of another version reads as undecodable and is re-folded from the log.
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
    /// Event-derived, so a replayed row equals the row it rebuilds.
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
