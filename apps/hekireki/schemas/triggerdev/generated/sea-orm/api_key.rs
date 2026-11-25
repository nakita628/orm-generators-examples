use super::api_key_issuer::ApiKeyIssuer;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "api_keys")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub key_hash: String,
    #[sea_orm(column_type = "Text")]
    pub last_four: String,
    /// Optional one-line note the creator adds to say what the key is for.
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(default_value = "USER")]
    pub issued_by: ApiKeyIssuer,
    #[sea_orm(column_type = "Text")]
    pub preset_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub scopes: Vec<String>,
    /// CIDR ranges this key may be used from. Empty means any IP.
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub allowed_ip_ranges: Vec<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub last_used_ip: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "api_keys_runtime_environment_id_fkey"
    )]
    RuntimeEnvironment,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "api_keys_created_by_user_id_fkey"
    )]
    CreatedBy,
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
