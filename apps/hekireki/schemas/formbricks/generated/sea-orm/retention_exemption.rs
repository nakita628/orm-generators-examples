use super::retention_entity::RetentionEntity;
use super::retention_exemption_source::RetentionExemptionSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
/// policy id, so removing and re-adding a policy doesn't orphan its exemptions.
///
/// Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
/// per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
/// (`WHERE "revokedAt" IS NULL`) in
/// `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
/// `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
/// README). An expired row still holds that slot until it is revoked, so creating a new exemption
/// must close an expired one in the same transaction.
///
/// `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
/// column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
/// (a survey carries its workspace, not its organisation): writers derive it from the survey.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RetentionExemption")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub entity: RetentionEntity,
    #[sea_orm(column_name = "surveyId")]
    pub survey_id: Option<String>,
    pub until: DateTime,
    #[sea_orm(default_value = "manual")]
    pub source: RetentionExemptionSource,
    /// Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`).
    pub reason: Option<String>,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<String>,
    #[sea_orm(column_name = "revokedAt")]
    pub revoked_at: Option<DateTime>,
    #[sea_orm(column_name = "revokedById")]
    pub revoked_by_id: Option<String>,
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
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::RevokedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    RevokedBy,
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
        Relation::CreatedBy.def()
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
        Ok(self)
    }
}
