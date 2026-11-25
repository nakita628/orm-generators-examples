use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-recipient cached credential metadata for the configured Cloud Signature
/// Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
/// symmetric-encrypted service-scope access token. One row per recipient on a
/// CSC envelope; absent for SES recipients.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CscCredential")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "providerId", column_type = "Text")]
    pub provider_id: String,
    #[sea_orm(column_name = "credentialId", column_type = "Text")]
    pub credential_id: String,
    /// TSP-validated certificate chain (length-prefixed). Nullable until the
    /// service-scope OAuth callback successfully retrieves and validates it.
    #[sea_orm(column_name = "certCache")]
    pub cert_cache: Option<Vec<u8>>,
    /// Algorithm metadata derived from the CSC `credentials/info` response.
    /// Held as flat scalars so each is independently queryable.
    #[sea_orm(column_name = "signatureAlgorithm", column_type = "Text")]
    pub signature_algorithm: String,
    #[sea_orm(column_name = "keyType", column_type = "Text")]
    pub key_type: String,
    #[sea_orm(column_name = "digestAlgorithm", column_type = "Text")]
    pub digest_algorithm: String,
    #[sea_orm(column_name = "keyLenBits")]
    pub key_len_bits: Option<i32>,
    #[sea_orm(column_name = "signAlgoParams", column_type = "Text")]
    pub sign_algo_params: Option<String>,
    /// Service-scope access token, symmetric-encrypted at the application layer
    /// before persistence. Nullable until the service-scope OAuth callback runs.
    #[sea_orm(column_name = "serviceTokenCiphertext")]
    pub service_token_ciphertext: Option<Vec<u8>>,
    #[sea_orm(
        column_name = "serviceTokenExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub service_token_expires_at: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "recipientId")]
    pub recipient_id: i32,
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
        belongs_to = "super::recipient::Entity",
        from = "Column::RecipientId",
        to = "super::recipient::Column::Id",
        on_delete = "Cascade",
        fk_name = "CscCredential_recipientId_fkey"
    )]
    Recipient,
}

impl Related<super::recipient::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Recipient.def()
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
