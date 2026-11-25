use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserUpdateRead")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "updateId", column_type = "Text")]
    pub update_id: String,
    #[sea_orm(
        column_name = "readAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub read_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "UserUpdateRead_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::update::Entity",
        from = "Column::UpdateId",
        to = "super::update::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "UserUpdateRead_updateId_fkey"
    )]
    Update,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::update::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Update.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.read_at.is_not_set() {
            self.read_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
