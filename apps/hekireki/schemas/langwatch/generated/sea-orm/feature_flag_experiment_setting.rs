use super::feature_flag_experiment_subject_type::FeatureFlagExperimentSubjectType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One experiment setting per (flag, subject). A USER row is that person's
/// own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
/// for that exact scope. No row means inherit.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeatureFlagExperimentSetting")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "flagKey", column_type = "Text")]
    pub flag_key: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "subjectType")]
    pub subject_type: FeatureFlagExperimentSubjectType,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "subjectId", column_type = "Text")]
    pub subject_id: String,
    pub enabled: bool,
    #[sea_orm(column_name = "changedByUserId", column_type = "Text")]
    pub changed_by_user_id: Option<String>,
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
