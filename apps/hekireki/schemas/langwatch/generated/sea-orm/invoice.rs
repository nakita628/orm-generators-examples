use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Invoice")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "subscriptionId", column_type = "Text")]
    pub subscription_id: String,
    #[sea_orm(column_name = "totalAmount")]
    pub total_amount: f64,
    #[sea_orm(column_type = "Text")]
    pub currency: String,
    #[sea_orm(column_name = "dueDate", column_type = "custom(\"timestamp(3)\")")]
    pub due_date: DateTime,
    #[sea_orm(default_value = false)]
    pub paid: bool,
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
        belongs_to = "super::subscription::Entity",
        from = "Column::SubscriptionId",
        to = "super::subscription::Column::Id",
        fk_name = "Invoice_subscriptionId_fkey"
    )]
    Subscription,
    #[sea_orm(has_many = "super::invoice_item::Entity")]
    LineItems,
}

impl Related<super::subscription::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Subscription.def()
    }
}

impl Related<super::invoice_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LineItems.def()
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
