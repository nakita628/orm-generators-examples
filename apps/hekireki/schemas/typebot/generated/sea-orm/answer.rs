use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Answer")]
pub struct Model {
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "resultId", column_type = "Text")]
    pub result_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "blockId", column_type = "Text")]
    pub block_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: String,
    #[sea_orm(column_name = "variableId", column_type = "Text")]
    pub variable_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub content: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::result::Entity",
        from = "Column::ResultId",
        to = "super::result::Column::Id",
        on_delete = "Cascade",
        fk_name = "Answer_resultId_fkey"
    )]
    Result,
}

impl Related<super::result::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Result.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, _insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.created_at.is_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
