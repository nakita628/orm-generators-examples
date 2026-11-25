use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "MessagesGroup")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "buyerOrganizationId", column_type = "Text")]
    pub buyer_organization_id: String,
    #[sea_orm(column_name = "buyerId", column_type = "Text")]
    pub buyer_id: String,
    #[sea_orm(column_name = "sellerId", column_type = "Text")]
    pub seller_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::BuyerId",
        to = "super::user::Column::Id",
        fk_name = "MessagesGroup_buyerId_fkey"
    )]
    Buyer,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::BuyerOrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "MessagesGroup_buyerOrganizationId_fkey"
    )]
    BuyerOrganization,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::SellerId",
        to = "super::user::Column::Id",
        fk_name = "MessagesGroup_sellerId_fkey"
    )]
    Seller,
    #[sea_orm(has_many = "super::messages::Entity")]
    Messages,
    #[sea_orm(has_many = "super::orders::Entity")]
    Orders,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Buyer.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BuyerOrganization.def()
    }
}

impl Related<super::messages::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Messages.def()
    }
}

impl Related<super::orders::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Orders.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
