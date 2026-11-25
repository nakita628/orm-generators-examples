use super::phone_number_subscription_status::PhoneNumberSubscriptionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CalAiPhoneNumber")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "phoneNumber", column_type = "Text")]
    pub phone_number: String,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "providerPhoneNumberId", column_type = "Text")]
    pub provider_phone_number_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "stripeCustomerId", column_type = "Text")]
    pub stripe_customer_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeSubscriptionId", column_type = "Text")]
    pub stripe_subscription_id: Option<String>,
    #[sea_orm(column_name = "subscriptionStatus")]
    pub subscription_status: Option<PhoneNumberSubscriptionStatus>,
    #[sea_orm(column_name = "inboundAgentId", column_type = "Text")]
    pub inbound_agent_id: Option<String>,
    #[sea_orm(column_name = "outboundAgentId", column_type = "Text")]
    pub outbound_agent_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "CalAiPhoneNumber_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "CalAiPhoneNumber_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::agent::Entity",
        from = "Column::InboundAgentId",
        to = "super::agent::Column::Id",
        on_delete = "SetNull",
        fk_name = "CalAiPhoneNumber_inboundAgentId_fkey"
    )]
    InboundAgent,
    #[sea_orm(
        belongs_to = "super::agent::Entity",
        from = "Column::OutboundAgentId",
        to = "super::agent::Column::Id",
        on_delete = "SetNull",
        fk_name = "CalAiPhoneNumber_outboundAgentId_fkey"
    )]
    OutboundAgent,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::agent::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InboundAgent.def()
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
