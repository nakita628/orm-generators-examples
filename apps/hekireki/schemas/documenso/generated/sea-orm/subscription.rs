use super::subscription_status::SubscriptionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Subscription")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(default_value = "INACTIVE")]
    pub status: SubscriptionStatus,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "planId", column_type = "Text")]
    pub plan_id: String,
    #[sea_orm(column_name = "priceId", column_type = "Text")]
    pub price_id: String,
    #[sea_orm(column_name = "periodEnd", column_type = "custom(\"timestamp(3)\")")]
    pub period_end: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "cancelAtPeriodEnd", default_value = false)]
    pub cancel_at_period_end: bool,
    #[sea_orm(column_name = "customerId", column_type = "Text")]
    pub customer_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "organisationId", column_type = "Text")]
    pub organisation_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organisation::Entity",
        from = "Column::OrganisationId",
        to = "super::organisation::Column::Id",
        on_delete = "Cascade",
        fk_name = "Subscription_organisationId_fkey"
    )]
    Organisation,
}

impl Related<super::organisation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organisation.def()
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
