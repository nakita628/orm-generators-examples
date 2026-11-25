use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "users")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub name: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub email: Option<String>,
    #[sea_orm(column_name = "emailVerified", column_type = "custom(\"datetime(3)\")")]
    pub email_verified: Option<DateTime>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub image: Option<String>,
    #[sea_orm(
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub updated_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub stripe_customer_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub stripe_subscription_id: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub stripe_price_id: Option<String>,
    #[sea_orm(column_type = "custom(\"datetime(3)\")")]
    pub stripe_current_period_end: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::account::Entity")]
    Accounts,
    #[sea_orm(has_many = "super::session::Entity")]
    Sessions,
    #[sea_orm(has_many = "super::post::Entity")]
    Post,
}

impl Related<super::account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Accounts.def()
    }
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Sessions.def()
    }
}

impl Related<super::post::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Post.def()
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
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
