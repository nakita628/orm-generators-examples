use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
/// public key and the signature counter are the plugin's to read and write.
/// A passkey also becomes an `Identifier` row with provider `passkey`, which
/// is what makes it a sign-in method like any other - this table is the
/// credential, that one is the identity.
///
/// `aaguid` names the authenticator MODEL, never a device or a person, and is
/// only present at registration. Privacy-preserving platforms report all
/// zeroes, so a label derived from it is best-effort by construction.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Passkey")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_name = "publicKey", column_type = "Text")]
    pub public_key: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "credentialID", column_type = "Text")]
    pub credential_id: String,
    #[sea_orm(default_value = 0)]
    pub counter: i32,
    #[sea_orm(column_name = "deviceType", column_type = "Text")]
    pub device_type: String,
    #[sea_orm(column_name = "backedUp", default_value = false)]
    pub backed_up: bool,
    #[sea_orm(column_type = "Text")]
    pub transports: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub aaguid: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Passkey_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
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
