use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
/// A read model, one row per organization: rebuildable by replaying the event log, so no relation.
/// Epochs in milliseconds.
/// Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OrganizationSeatPeak")]
pub struct Model {
    /// The organization, which is also the aggregate id.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// When the opening snapshot was taken, epoch milliseconds.
    #[sea_orm(column_name = "openedAt")]
    pub opened_at: f64,
    #[sea_orm(column_name = "fullSeats")]
    pub full_seats: i32,
    #[sea_orm(column_name = "liteSeats")]
    pub lite_seats: i32,
    #[sea_orm(column_name = "developerSeats")]
    pub developer_seats: i32,
    /// The latest change time applied to any holder; invitations expire up to it.
    #[sea_orm(column_name = "latestChangeAt")]
    pub latest_change_at: f64,
    /// Changes whose `from` disagreed with the seat the fold held.
    pub mismatches: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
