use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The values an `SsoConnection`'s credential references point AT (D09 - see
/// specs/identity/sso-idp-termination.feature).
///
/// The aggregate has always carried `clientIdRef` / `secretRef` / `certRefs`
/// rather than values, because a fact may not carry a secret (ADR-101 §4).
/// Until D09 there was nothing behind the references and every writer passed
/// null. This is the thing behind them.
///
/// `ciphertext` is AES-256-GCM under the deployment's `CREDENTIALS_SECRET`,
/// the same key every other credential in the product is kept under. Nothing
/// reads a row except the projection that builds the engine's dialing
/// configuration; no surface renders one back.
///
/// Written by the command path rather than the fold, and that split is
/// deliberate: the VALUE is not event truth and cannot be, so the log records
/// WHEN a credential changed (a new reference in a new fact) and never what
/// it was. A replay rebuilds every reference from the log and finds the value
/// still here.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoCredential")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// The connection whose registration minted this. Not a relation, matching
    /// `SsoConnection` itself: a projection rebuilt from events must not depend
    /// on relational state the log does not carry.
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: String,
    /// oidc-client-id | oidc-client-secret | saml-idp-config
    pub kind: String,
    pub ciphertext: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
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
