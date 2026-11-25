use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BannedIp")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub ip: String,
    #[sea_orm(column_name = "responsibleTypebotId", column_type = "Text")]
    pub responsible_typebot_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::typebot::Entity",
        from = "Column::ResponsibleTypebotId",
        to = "super::typebot::Column::Id",
        on_delete = "Restrict",
        fk_name = "BannedIp_responsibleTypebotId_fkey"
    )]
    ResponsibleTypebot,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Restrict",
        fk_name = "BannedIp_userId_fkey"
    )]
    User,
}

impl Related<super::typebot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ResponsibleTypebot.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
        Ok(self)
    }
}
