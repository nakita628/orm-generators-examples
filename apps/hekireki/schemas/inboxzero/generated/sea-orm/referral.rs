use super::referral_status::ReferralStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Referral")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "referrerUserId", column_type = "Text")]
    pub referrer_user_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "referredUserId", column_type = "Text")]
    pub referred_user_id: String,
    #[sea_orm(column_name = "referralCodeUsed", column_type = "Text")]
    pub referral_code_used: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: ReferralStatus,
    #[sea_orm(
        column_name = "rewardGrantedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub reward_granted_at: Option<DateTime>,
    #[sea_orm(column_name = "stripeBalanceTransactionId", column_type = "Text")]
    pub stripe_balance_transaction_id: Option<String>,
    #[sea_orm(column_name = "rewardAmount")]
    pub reward_amount: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ReferrerUserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Referral_referrerUserId_fkey"
    )]
    ReferrerUser,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ReferredUserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Referral_referredUserId_fkey"
    )]
    ReferredUser,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReferrerUser.def()
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
