use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Better Auth JWT signing keys for OAuth/OIDC access and ID tokens.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "jwks")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "publicKey")]
    pub public_key: String,
    #[sea_orm(column_name = "privateKey")]
    pub private_key: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: Option<DateTime>,
    /// Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
    /// than one alg and `getLatestKeyByAlg` can pick the right key per resource.
    ///
    /// NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
    /// puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
    /// rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
    /// OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
    /// 1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
    /// keys keep working untouched.
    pub alg: Option<String>,
    /// Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
    /// the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
    /// same always-written path as `alg` above.
    pub crv: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
