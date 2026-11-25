use super::prisma_date_time::PrismaDateTime;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "snippets")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub uuid: String,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub identifier: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    #[sea_orm(column_type = "Text")]
    pub language: String,
    #[sea_orm(column_name = "parentUuid", column_type = "Text")]
    pub parent_uuid: Option<String>,
    #[sea_orm(column_name = "createdAt", default_expr = "Expr::current_timestamp()")]
    pub created_at: PrismaDateTime,
    #[sea_orm(column_name = "editedAt")]
    pub edited_at: Option<PrismaDateTime>,
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
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.into());
        }
        Ok(self)
    }
}
