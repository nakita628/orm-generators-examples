use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Customer")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub name: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub email: Option<String>,
    #[sea_orm(
        column_name = "normalizedEmail",
        column_type = "String(StringLen::N(191))"
    )]
    pub normalized_email: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub avatar: Option<String>,
    #[sea_orm(column_name = "externalId", column_type = "String(StringLen::N(191))")]
    pub external_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(
        column_name = "stripeCustomerId",
        column_type = "String(StringLen::N(191))"
    )]
    pub stripe_customer_id: Option<String>,
    #[sea_orm(column_name = "linkId", column_type = "String(StringLen::N(191))")]
    pub link_id: Option<String>,
    #[sea_orm(column_name = "clickId", column_type = "String(StringLen::N(191))")]
    pub click_id: Option<String>,
    #[sea_orm(column_name = "clickedAt", column_type = "custom(\"datetime(3)\")")]
    pub clicked_at: Option<DateTime>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub country: Option<String>,
    #[sea_orm(default_value = 0)]
    pub sales: i32,
    #[sea_orm(column_name = "saleAmount", default_value = "0")]
    pub sale_amount: i64,
    #[sea_orm(column_name = "firstSaleAt", column_type = "custom(\"datetime(3)\")")]
    pub first_sale_at: Option<DateTime>,
    #[sea_orm(
        column_name = "subscriptionCanceledAt",
        column_type = "custom(\"datetime(3)\")"
    )]
    pub subscription_canceled_at: Option<DateTime>,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(
        column_name = "projectConnectId",
        column_type = "String(StringLen::N(191))"
    )]
    pub project_connect_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "Customer_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::link::Entity",
        from = "Column::LinkId",
        to = "super::link::Column::Id",
        fk_name = "Customer_linkId_fkey"
    )]
    Link,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Link.def()
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
