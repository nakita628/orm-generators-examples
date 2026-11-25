use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Durable per-stage state for budget-staged computed updates, keyed by the
/// continuation chain's root task id (its scope). Rows are written once per
/// record and never copied between continuation tasks:
/// - kind 'excluded': processed-target exclusion ledger for partial batches;
/// - kind 'frontier': seq-ordered queue of sources whose outgoing propagation
/// is not finished (self-referential generations + migrated explicit seeds);
/// - kind 'consumed': retired frontier sources preserved for deferred edge
/// chunks, handed to the continuation as seeds at stage completion.
/// Cleared when the chain's stage completes or the chain dead-letters.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "computed_update_stage_ledger")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub scope_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub kind: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub table_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub record_id: String,
    #[sea_orm(default_value = "0")]
    pub seq: i64,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
