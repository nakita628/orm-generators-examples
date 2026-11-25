use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ReferralInvite")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: Option<String>,
    /// *
    /// * Whose cap and discount this invite counts against — the fund's team, or the
    /// * sender alone. One non-null key rather than a pair of nullable ones: on a
    /// * nullable teamId every personal invite is distinct to Postgres, so the two
    /// * keys guarded different scopes and a member could not invite a company both
    /// * for their fund and for themselves.
    #[sea_orm(column_name = "scopeId", column_type = "Text")]
    pub scope_id: String,
    #[sea_orm(column_type = "Text")]
    pub email: String,
    #[sea_orm(column_name = "referralLink", column_type = "Text")]
    pub referral_link: Option<String>,
    #[sea_orm(
        column_name = "sentAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub sent_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "ReferralInvite_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "ReferralInvite_teamId_fkey"
    )]
    Team,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.sent_at.is_not_set() {
            self.sent_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
