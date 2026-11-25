use super::subscription_provider::SubscriptionProvider;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Subscription")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    pub active: bool,
    pub provider: SubscriptionProvider,
    #[sea_orm(
        column_name = "currentPeriodStart",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub current_period_start: DateTime,
    #[sea_orm(
        column_name = "currentPeriodEnd",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub current_period_end: DateTime,
    #[sea_orm(default_value = 1)]
    pub quantity: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "storeOriginalTransactionId", column_type = "Text")]
    pub store_original_transaction_id: Option<String>,
    #[sea_orm(column_name = "storeProductId", column_type = "Text")]
    pub store_product_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "googlePurchaseToken", column_type = "Text")]
    pub google_purchase_token: Option<String>,
    #[sea_orm(column_name = "storeMetadata", column_type = "JsonBinary")]
    pub store_metadata: Option<Json>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeSubscriptionId", column_type = "Text")]
    pub stripe_subscription_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Subscription_userId_fkey"
    )]
    User,
    #[sea_orm(
        has_many = "super::user::Entity",
        from = "Column::Id",
        to = "super::user::Column::ParentSubscriptionId"
    )]
    ChildUsers,
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
