use super::retention_entity::RetentionEntity;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The notice for one target under one policy, so each notice goes out once: a survey or a member, or
/// the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
/// claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
/// members) before the target's clock: then a later cycle claims it again and sends a new one.
///
/// A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
/// `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
/// hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
/// 'surveys' AND "clockAt" IS NOT NULL`) in
/// `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
/// that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
/// "Indexes Prisma cannot express" in the README).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionNotice")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    pub entity: RetentionEntity,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    /// When the notice was claimed for sending. A claim that never got delivered can be taken over.
    #[sea_orm(
        column_name = "sentAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub sent_at: DateTime,
    /// When the notice was handed to the mail transport, or recorded without one (no SMTP). The
    /// warning only counts from a delivered notice, so a crash between claim and send can't let the
    /// action run without it.
    #[sea_orm(column_name = "deliveredAt", column_type = "custom(\"timestamp(3)\")")]
    pub delivered_at: Option<DateTime>,
    /// Whether an email actually went out. False when SMTP isn't configured.
    #[sea_orm(column_name = "emailSent", default_value = false)]
    pub email_sent: bool,
    /// Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
    /// can't mark a newer claim delivered.
    #[sea_orm(column_name = "claimToken", column_type = "Text")]
    pub claim_token: Option<String>,
    /// The target's clock the notice was computed for (surveys and members). A notice counts only while
    /// the target's clock is still this one, so activity between the sweep's read and its claim voids
    /// it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
    /// for a survey notice voided for good: a person deleted or moved one of its responses, or the
    /// responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
    #[sea_orm(column_name = "clockAt", column_type = "custom(\"timestamp(3)\")")]
    pub clock_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "RetentionNotice_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "RetentionNotice_surveyId_fkey"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "RetentionNotice_userId_fkey"
    )]
    User,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
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
        if insert && self.sent_at.is_not_set() {
            self.sent_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
