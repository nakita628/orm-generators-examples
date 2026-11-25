use super::gateway_connection_status::GatewayConnectionStatus;
use super::gateway_provider::GatewayProvider;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "gateway_ai_connections")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub provider: GatewayProvider,
    #[sea_orm(column_name = "encrypted_credentials", column_type = "Text")]
    pub encrypted_credential: String,
    #[sea_orm(column_name = "display_secret_key", column_type = "Text")]
    pub display_secret: String,
    #[sea_orm(column_type = "Text")]
    pub created_by_id: Option<String>,
    #[sea_orm(default_value = "enabled")]
    pub status: GatewayConnectionStatus,
    #[sea_orm(column_type = "JsonBinary")]
    pub allowed_models: Option<Json>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub custom_models: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub base_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub extra_headers: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub extra_header_keys: Vec<String>,
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
        fk_name = "gateway_ai_connections_organization_id_fkey"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
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
