use super::retention_entity::RetentionEntity;
use super::retention_survey_condition::RetentionSurveyCondition;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
/// `periodDays` after the target's clock. The action follows from the kind of data: responses are
/// deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
/// survey), members deactivated. The dates themselves are never stored; they are computed by the shared
/// helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
/// always agree.
///
/// The allowed ranges are validated by the API. The migration repeats the ones that guard against
/// acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
/// policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
/// CHECK constraints, which Prisma cannot express.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionPolicy")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub entity: RetentionEntity,
    #[sea_orm(default_value = false)]
    pub enabled: bool,
    /// When the current configuration took effect: the policy was switched on or unpaused, or changed
    /// in a way that could bring an action earlier (a shorter period or warning, different
    /// conditions). Lengthening does not move it. The warning always runs in full from here, so no
    /// change can delete data on the night it is made (ENG-3614). Also the clock for members with no
    /// recorded sign-in.
    #[sea_orm(column_name = "enabledAt")]
    pub enabled_at: Option<DateTime>,
    /// How long before the action the notice goes out. The action never runs before the notice has
    /// been out for this long, so the warning always runs in full.
    #[sea_orm(column_name = "warnDays", default_value = 60)]
    pub warn_days: i32,
    /// The period, counted from the target's clock, after which the policy acts: responses are
    /// deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
    #[sea_orm(column_name = "periodDays")]
    pub period_days: i32,
    /// Surveys policy only.
    pub conditions: Vec<RetentionSurveyCondition>,
    /// The last night the sweep's time ran out before it reached this policy's organisation, licensed.
    /// The organisation was due and the sweep was running, so it counts as reached for the gap that
    /// restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
    #[sea_orm(column_name = "deferredAt")]
    pub deferred_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedById")]
    pub updated_by_id: Option<String>,
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
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    UpdatedBy,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UpdatedBy.def()
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
