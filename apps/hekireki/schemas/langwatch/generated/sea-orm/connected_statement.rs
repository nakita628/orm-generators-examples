use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The monthly usage statement one connected customer was sent (ADR-139,
/// section 7). One row per account per month keeps the statement to one send,
/// whatever runs the tick.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ConnectedStatement")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The `ConnectedBillingAccount` row. A plain column ref.
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: String,
    /// First day of the month the statement covers, in UTC.
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub month: DateTime,
    #[sea_orm(column_name = "sentAt", column_type = "custom(\"timestamp(3)\")")]
    pub sent_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
