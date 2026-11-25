use super::subscription_interval::SubscriptionInterval;
use super::subscription_status::SubscriptionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "subscriptions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub price_id: String,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
    #[sea_orm(column_type = "Text")]
    pub subscription_item_id: String,
    pub amount: i32,
    pub discount_percent_off: Option<f64>,
    pub discount_amount_off: Option<i32>,
    pub status: SubscriptionStatus,
    pub active: bool,
    #[sea_orm(column_type = "Text")]
    pub currency: String,
    pub interval: SubscriptionInterval,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub period_start: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub period_end: DateTime,
    #[sea_orm(default_value = false)]
    pub cancel_at_period_end: bool,
    #[sea_orm(column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub space_id: String,
    #[sea_orm(column_type = "Text")]
    pub api_addon_item_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub api_addon_price_id: Option<String>,
    pub api_addon_amount: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub api_addon_tier: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub api_poll_retained_tier: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub api_poll_retained_until: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "subscriptions_user_id_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::space::Entity",
        from = "Column::SpaceId",
        to = "super::space::Column::Id",
        on_delete = "Cascade",
        fk_name = "subscriptions_space_id_fkey"
    )]
    Space,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::space::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Space.def()
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
