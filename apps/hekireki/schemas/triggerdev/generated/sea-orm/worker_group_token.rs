use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkerGroupToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "tokenHash", column_type = "Text")]
    pub token_hash: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "previousTokenHash", column_type = "Text")]
    pub previous_token_hash: Option<String>,
    #[sea_orm(
        column_name = "previousTokenValidUntil",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub previous_token_valid_until: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "workerGroupId", column_type = "Text")]
    pub worker_group_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::worker_instance_group::Entity",
        from = "Column::WorkerGroupId",
        to = "super::worker_instance_group::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "WorkerGroupToken_workerGroupId_fkey"
    )]
    ChildOf,
    #[sea_orm(
        has_one = "super::worker_instance_group::Entity",
        from = "Column::Id",
        to = "super::worker_instance_group::Column::TokenId"
    )]
    WorkerGroup,
}

impl Related<super::worker_instance_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ChildOf.def()
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
