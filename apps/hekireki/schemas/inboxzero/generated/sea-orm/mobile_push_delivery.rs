use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "MobilePushDelivery")]
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
    #[sea_orm(column_name = "deduplicationKey", column_type = "Text")]
    pub deduplication_key: String,
    #[sea_orm(column_name = "mobilePushTokenId", column_type = "Text")]
    pub mobile_push_token_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::mobile_push_token::Entity",
        from = "Column::MobilePushTokenId",
        to = "super::mobile_push_token::Column::Id",
        on_delete = "Cascade",
        fk_name = "MobilePushDelivery_mobilePushTokenId_fkey"
    )]
    MobilePushToken,
}

impl Related<super::mobile_push_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MobilePushToken.def()
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
