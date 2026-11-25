use super::oidc_signing_key_purpose::OidcSigningKeyPurpose;
use super::oidc_signing_key_status::OidcSigningKeyStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OidcSigningKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub kid: String,
    #[sea_orm(default_value = "EXTERNAL_OIDC")]
    pub purpose: OidcSigningKeyPurpose,
    /// At most one STAGED and one ACTIVE key per purpose: enforced by a partial unique index in the migration.
    pub status: OidcSigningKeyStatus,
    #[sea_orm(column_name = "publicJwk", column_type = "JsonBinary")]
    pub public_jwk: Json,
    #[sea_orm(column_name = "readyAt", column_type = "custom(\"timestamp(3)\")")]
    pub ready_at: DateTime,
    #[sea_orm(column_name = "publicUntil", column_type = "custom(\"timestamp(3)\")")]
    pub public_until: Option<DateTime>,
    #[sea_orm(column_name = "generatedByUserId", column_type = "Text")]
    pub generated_by_user_id: String,
    #[sea_orm(column_name = "activatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub activated_at: Option<DateTime>,
    #[sea_orm(column_name = "activatedByUserId", column_type = "Text")]
    pub activated_by_user_id: Option<String>,
    #[sea_orm(column_name = "retiredAt", column_type = "custom(\"timestamp(3)\")")]
    pub retired_at: Option<DateTime>,
    #[sea_orm(column_name = "retiredByUserId", column_type = "Text")]
    pub retired_by_user_id: Option<String>,
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
pub enum Relation {}

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
