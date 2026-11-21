use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
/// assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
/// table, so omitting it risks an adapter error on a code path we do not control.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "oauthClientAssertion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "expiresAt")]
    pub expires_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
