use super::creative_show_on::CreativeShowOn;
use super::creative_status::CreativeStatus;
use super::creative_type::CreativeType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
/// (`linkUrl`), never both.
///
/// @property type - Fixed at creation.
/// @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
/// @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
/// @property survey - The Destination survey, in the same workspace by construction of the composite key.
/// @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
/// @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Creative")]
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
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub r#type: CreativeType,
    #[sea_orm(default_value = "draft")]
    pub status: CreativeStatus,
    #[sea_orm(column_name = "publishedAt", column_type = "custom(\"timestamp(3)\")")]
    pub published_at: Option<DateTime>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub content: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub styling: Json,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: Option<String>,
    #[sea_orm(column_name = "linkUrl", column_type = "Text")]
    pub link_url: Option<String>,
    #[sea_orm(column_name = "showOn", default_value = "pageView")]
    pub show_on: CreativeShowOn,
    #[sea_orm(column_name = "showConditions", column_type = "JsonBinary")]
    pub show_conditions: Option<Json>,
    #[sea_orm(column_name = "displayPercentage")]
    pub display_percentage: Option<f64>,
    #[sea_orm(column_name = "startAt", column_type = "custom(\"timestamp(3)\")")]
    pub start_at: Option<DateTime>,
    #[sea_orm(column_name = "endAt", column_type = "custom(\"timestamp(3)\")")]
    pub end_at: Option<DateTime>,
    #[sea_orm(column_name = "autoSelectLanguage")]
    pub auto_select_language: Option<bool>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Creative_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "Creative_createdBy_fkey"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "(Column::SurveyId, Column::WorkspaceId)",
        to = "(super::survey::Column::Id, super::survey::Column::WorkspaceId)",
        on_update = "NoAction",
        on_delete = "NoAction",
        fk_name = "Creative_surveyId_workspaceId_fkey"
    )]
    Survey,
    #[sea_orm(has_many = "super::creative_trigger::Entity")]
    Triggers,
    #[sea_orm(has_many = "super::creative_language::Entity")]
    Languages,
    #[sea_orm(has_many = "super::creative_daily_stat::Entity")]
    DailyStats,
    #[sea_orm(has_many = "super::creative_daily_page_stat::Entity")]
    DailyPageStats,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::creative_trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Triggers.def()
    }
}

impl Related<super::creative_language::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Languages.def()
    }
}

impl Related<super::creative_daily_stat::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DailyStats.def()
    }
}

impl Related<super::creative_daily_page_stat::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DailyPageStats.def()
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
