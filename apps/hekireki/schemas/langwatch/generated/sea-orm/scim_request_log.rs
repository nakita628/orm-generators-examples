use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Every SCIM request we could attribute, and what we answered (ADR-126).
///
/// EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
/// was refused decided nothing — so this is a table and not an event: nothing
/// downstream derives from it, rebuilding the world from the log must not
/// depend on how often a provider retried a GET, and rows age out of a
/// retention window an event log would never have.
///
/// NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
/// answered 401 and recorded nowhere: we do not know whose organization to
/// file it under, and a table unauthenticated traffic can write is a table
/// anybody on the internet can fill. What answers that case instead is
/// `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
/// a credential we recognize.
///
/// THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
/// and our own sentence. Never a token, never a header, never a provider's
/// raw message — the same reason `scim_apply_failed` carries a code and not a
/// message.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScimRequestLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// Null when the refusal resolved an organization and stopped before the
    /// connection — a lapsed plan is answered off the token alone.
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: Option<String>,
    pub method: String,
    /// The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
    /// raw path: a path carries ids and query strings, and this is read by
    /// people rather than matched by machines.
    pub resource: String,
    pub status: i32,
    /// A stable slug, only on a refusal, so a reader branches on this and never
    /// on the sentence beside it.
    pub reason: Option<String>,
    /// Our own short sentence, customer-safe.
    pub detail: Option<String>,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade"
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
        if insert && self.occurred_at.is_not_set() {
            self.occurred_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
