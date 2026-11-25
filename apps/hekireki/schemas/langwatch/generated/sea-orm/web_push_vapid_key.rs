use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// This installation's VAPID key pair, generated on first use so Web Push works
/// without configuration. Fixed id, so the table holds one row: a second
/// insert loses the race on the primary key and reads the winner's pair.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WebPushVapidKey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The uncompressed P-256 public key, base64url. The browser subscribes with it.
    #[sea_orm(column_name = "publicKey", column_type = "Text")]
    pub public_key: String,
    /// The private key, base64url, encrypted with the process encryption key.
    #[sea_orm(column_name = "privateKeyEncrypted", column_type = "Text")]
    pub private_key_encrypted: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
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
        Ok(self)
    }
}
