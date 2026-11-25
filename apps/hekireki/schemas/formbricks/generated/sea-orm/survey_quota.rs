use super::survey_quota_action::SurveyQuotaAction;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a quota configuration for a survey.
/// Defines response limits and conditions for quota management.
///
/// @property id - Unique identifier for the quota
/// @property survey - The survey this quota belongs to
/// @property name - Display name for the quota
/// @property limit - Maximum number of responses allowed
/// @property conditions - JSON object containing quota conditions
/// @property action - Action to take when quota is reached
/// @property endingCardId - Optional ending card to show when quota is reached
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SurveyQuota")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub limit: i32,
    /// [SurveyQuotaLogic]
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub logic: Json,
    pub action: SurveyQuotaAction,
    #[sea_orm(column_name = "endingCardId", column_type = "Text")]
    pub ending_card_id: Option<String>,
    #[sea_orm(column_name = "countPartialSubmissions", default_value = false)]
    pub count_partial_submissions: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "Column::SurveyId",
        to = "super::survey::Column::Id",
        on_delete = "Cascade",
        fk_name = "SurveyQuota_surveyId_fkey"
    )]
    Survey,
    #[sea_orm(has_many = "super::response_quota_link::Entity")]
    QuotaLinks,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::response_quota_link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::QuotaLinks.def()
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
