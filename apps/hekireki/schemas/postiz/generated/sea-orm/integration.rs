use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Integration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "internalId", column_type = "Text")]
    pub internal_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "customName", column_type = "Text")]
    pub custom_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub picture: Option<String>,
    #[sea_orm(column_name = "providerIdentifier", column_type = "Text")]
    pub provider_identifier: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub token: String,
    #[sea_orm(default_value = false)]
    pub disabled: bool,
    #[sea_orm(
        column_name = "tokenExpiration",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub token_expiration: Option<DateTime>,
    #[sea_orm(column_name = "refreshToken", column_type = "Text")]
    pub refresh_token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub profile: Option<String>,
    #[sea_orm(column_name = "deletedAt", column_type = "custom(\"timestamp(3)\")")]
    pub deleted_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_name = "inBetweenSteps", default_value = false)]
    pub in_between_steps: bool,
    #[sea_orm(column_name = "refreshNeeded", default_value = false)]
    pub refresh_needed: bool,
    #[sea_orm(
        column_name = "postingTimes",
        column_type = "Text",
        default_value = "[{\"time\":120}, {\"time\":400}, {\"time\":700}]"
    )]
    pub posting_times: String,
    #[sea_orm(column_name = "customInstanceDetails", column_type = "Text")]
    pub custom_instance_details: Option<String>,
    #[sea_orm(column_name = "customerId", column_type = "Text")]
    pub customer_id: Option<String>,
    #[sea_orm(column_name = "rootInternalId", column_type = "Text")]
    pub root_internal_id: Option<String>,
    #[sea_orm(
        column_name = "additionalSettings",
        column_type = "Text",
        default_value = "[]"
    )]
    pub additional_settings: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::customer::Entity",
        from = "Column::CustomerId",
        to = "super::customer::Column::Id",
        fk_name = "Integration_customerId_fkey"
    )]
    Customer,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "Integration_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(has_many = "super::exising_plug_data::Entity")]
    ExisingPlugData,
    #[sea_orm(has_many = "super::integrations_webhooks::Entity")]
    Webhooks,
    #[sea_orm(has_many = "super::order_items::Entity")]
    OrderItems,
    #[sea_orm(has_many = "super::plugs::Entity")]
    Plugs,
    #[sea_orm(has_many = "super::post::Entity")]
    Posts,
}

impl Related<super::customer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Customer.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::exising_plug_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExisingPlugData.def()
    }
}

impl Related<super::integrations_webhooks::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhooks.def()
    }
}

impl Related<super::order_items::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OrderItems.def()
    }
}

impl Related<super::plugs::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Plugs.def()
    }
}

impl Related<super::post::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Posts.def()
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
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
