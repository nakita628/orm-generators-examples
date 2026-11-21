use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-recipient transient sign-time state held across the credential-scope
/// OAuth round-trip. The `recipientId @unique` constraint caps each recipient
/// to at most one in-flight session — re-clicking Sign UPSERTs the row.
/// Cleaned up transitively when the Recipient is deleted.
/// @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "CscSession")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
    /// flows through the Recipient cascade.
    #[sea_orm(column_name = "envelopeId")]
    pub envelope_id: String,
    /// Pinned at prep time; stable across overlay re-renders so the captured
    /// signedAttrs digest matches the bytes the TSP signs.
    #[sea_orm(column_name = "signingTime")]
    pub signing_time: DateTime,
    /// Contract between prep and sign: an array of
    /// `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
    /// `documentDataId` pins the immutable bytes that produced the captured
    /// digest so sign-time can re-derive the same hash against the same content.
    /// [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
    #[sea_orm(column_name = "itemsJson")]
    pub items_json: Json,
    /// Populated by the credential-scope OAuth callback. Both nullable at row
    /// creation; set together when the SAD is exchanged.
    #[sea_orm(column_name = "encryptedSad")]
    pub encrypted_sad: Option<Vec<u8>>,
    #[sea_orm(column_name = "sadExpiresAt")]
    pub sad_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "recipientId")]
    pub recipient_id: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::recipient::Entity",
        from = "Column::RecipientId",
        to = "super::recipient::Column::Id",
        on_delete = "Cascade"
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
        Ok(self)
    }
}
