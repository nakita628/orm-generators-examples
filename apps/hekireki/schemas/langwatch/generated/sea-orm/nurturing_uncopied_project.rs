use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
/// retry clock reads or writes this table; an operator step may drop it once old writers are gone.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingUncopiedProject")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(
        column_name = "firstSeenAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub first_seen_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.first_seen_at.is_not_set() {
            self.first_seen_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
