use super::gateway_spend_limit_interval::GatewaySpendLimitInterval;
use super::gateway_spend_limit_scope::GatewaySpendLimitScope;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "gateway_spend_limits")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub organization_id: String,
    pub scope: GatewaySpendLimitScope,
    #[sea_orm(column_type = "Text")]
    pub api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_id: Option<String>,
    pub interval: GatewaySpendLimitInterval,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))")]
    pub limit_usd: Decimal,
    pub alert_thresholds: Vec<i32>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub counting_started_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub owner_user_id: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_spend_limits_organization_id_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::api_key::Entity",
        from = "Column::ApiKeyId",
        to = "super::api_key::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_spend_limits_api_key_id_fkey"
    )]
    ApiKey,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "gateway_spend_limits_user_id_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::OwnerUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "gateway_spend_limits_owner_user_id_fkey"
    )]
    OwnerUser,
    #[sea_orm(has_many = "super::gateway_spend_limit_alert::Entity")]
    Alerts,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::api_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiKey.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::gateway_spend_limit_alert::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Alerts.def()
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
