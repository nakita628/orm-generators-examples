use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
/// copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
/// fact only when it changes this row.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingOrganizationPlan")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub plan: Option<String>,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamp(3)\")")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "fromCatchUp")]
    pub from_catch_up: bool,
    /// The event ids folded in, so a redelivery folds once.
    #[sea_orm(
        column_name = "appliedEventIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub applied_event_ids: Vec<String>,
    /// The fact that last changed the row; only it may tell Customer.io.
    #[sea_orm(column_name = "winnerEventId", column_type = "Text")]
    pub winner_event_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
