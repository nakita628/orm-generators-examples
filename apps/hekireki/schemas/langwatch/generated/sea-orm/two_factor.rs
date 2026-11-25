use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The two-factor plugin's own table (D06). Protocol state, and row-truth for
/// good (ADR-101 R12): the shared secret and the backup codes live here and
/// in no event, ever. Nothing in this repo reads or writes these columns -
/// the plugin owns them end to end, including the lockout counters, which is
/// why `failedVerificationCount` and `lockedUntil` are here rather than
/// rebuilt on the projection.
///
/// `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
/// compared against what somebody types, and the plugin decrypts and compares;
/// hashing them would leave it unable to verify its own codes.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TwoFactor")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub secret: String,
    #[sea_orm(column_name = "backupCodes", column_type = "Text")]
    pub backup_codes: String,
    #[sea_orm(default_value = true)]
    pub verified: bool,
    #[sea_orm(column_name = "failedVerificationCount", default_value = 0)]
    pub failed_verification_count: i32,
    #[sea_orm(column_name = "lockedUntil", column_type = "custom(\"timestamp(3)\")")]
    pub locked_until: Option<DateTime>,
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
        fk_name = "TwoFactor_userId_fkey"
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
