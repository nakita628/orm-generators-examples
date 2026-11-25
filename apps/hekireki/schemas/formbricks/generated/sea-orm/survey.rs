use super::display_options::displayOptions;
use super::survey_script_mode::SurveyScriptMode;
use super::survey_status::SurveyStatus;
use super::survey_type::SurveyType;
use super::survey_visibility::SurveyVisibility;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a complete survey configuration including questions, styling, and display rules.
/// Core model for the survey functionality in Formbricks.
///
/// @property id - Unique identifier for the survey
/// @property name - Display name of the survey
/// @property type - Survey delivery method (link, web, website, app)
/// @property status - Current state of the survey (draft, active, completed, etc)
/// @property workspace - The workspace this survey belongs to
/// @property questions - JSON array containing survey questions configuration
/// @property displayOption - Rules for how often the survey can be shown
/// @property triggers - Actions that can trigger this survey
/// @property attributeFilters - Rules for targeting specific contacts
/// @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
/// @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Survey")]
pub struct Model {
    /// [SurveyHiddenFields]
    #[sea_orm(
        column_name = "hiddenFields",
        column_type = "JsonBinary",
        default_value = "{\"enabled\": false}"
    )]
    pub hidden_fields: Json,
    /// [SurveyVariables]
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub variables: Json,
    #[sea_orm(column_name = "ownershipInitialized", default_value = false)]
    pub ownership_initialized: bool,
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
    #[sea_orm(column_name = "redirectUrl", column_type = "Text")]
    pub redirect_url: Option<String>,
    #[sea_orm(default_value = "app")]
    pub r#type: SurveyType,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    /// ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
    /// the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
    #[sea_orm(default_value = "workspace")]
    pub visibility: SurveyVisibility,
    #[sea_orm(column_name = "ownerId", column_type = "Text")]
    pub owner_id: Option<String>,
    /// Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
    /// with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
    /// projection, decided from PostgreSQL facts until the projector acknowledges it.
    #[sea_orm(column_name = "visibilityVersion", default_value = 0)]
    pub visibility_version: i32,
    /// The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was.
    #[sea_orm(column_name = "visibilityProjectedVersion", default_value = 0)]
    pub visibility_projected_version: i32,
    /// "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
    /// trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
    /// in create/update data.
    #[sea_orm(column_name = "visibilityPending", default_value = false)]
    pub visibility_pending: bool,
    #[sea_orm(
        column_name = "visibilityChangedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub visibility_changed_at: Option<DateTime>,
    #[sea_orm(column_name = "visibilityChangedById", column_type = "Text")]
    pub visibility_changed_by_id: Option<String>,
    #[sea_orm(default_value = "draft")]
    pub status: SurveyStatus,
    /// [SurveyWelcomeCard]
    #[sea_orm(
        column_name = "welcomeCard",
        column_type = "JsonBinary",
        default_value = "{\"enabled\": false}"
    )]
    pub welcome_card: Json,
    /// [SurveyQuestions]
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub questions: Json,
    /// [SurveyBlocks]
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::JsonBinary))")]
    pub blocks: Vec<Json>,
    /// [SurveyEnding]
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::JsonBinary))")]
    pub endings: Vec<Json>,
    #[sea_orm(column_name = "displayOption", default_value = "displayOnce")]
    pub display_option: displayOptions,
    #[sea_orm(column_name = "recontactDays")]
    pub recontact_days: Option<i32>,
    #[sea_orm(column_name = "displayLimit")]
    pub display_limit: Option<i32>,
    /// [SurveyInlineTriggers]
    #[sea_orm(column_name = "inlineTriggers", column_type = "JsonBinary")]
    pub inline_triggers: Option<Json>,
    #[sea_orm(column_name = "autoClose")]
    pub auto_close: Option<i32>,
    #[sea_orm(column_name = "keepEndingCardOpen", default_value = false)]
    pub keep_ending_card_open: bool,
    #[sea_orm(column_name = "autoComplete")]
    pub auto_complete: Option<i32>,
    #[sea_orm(default_value = 0)]
    pub delay: i32,
    #[sea_orm(column_name = "publishOn", column_type = "custom(\"timestamp(3)\")")]
    pub publish_on: Option<DateTime>,
    #[sea_orm(column_name = "closeOn", column_type = "custom(\"timestamp(3)\")")]
    pub close_on: Option<DateTime>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    /// [SurveyClosedMessage]
    #[sea_orm(column_name = "surveyClosedMessage", column_type = "JsonBinary")]
    pub survey_closed_message: Option<Json>,
    #[sea_orm(column_name = "segmentId", column_type = "Text")]
    pub segment_id: Option<String>,
    /// [SurveyWorkspaceOverwrites]
    #[sea_orm(column_name = "workspaceOverwrites", column_type = "JsonBinary")]
    pub workspace_overwrites: Option<Json>,
    /// [SurveyStyling]
    #[sea_orm(column_type = "JsonBinary")]
    pub styling: Option<Json>,
    /// [SurveySingleUse]
    #[sea_orm(
        column_name = "singleUse",
        column_type = "JsonBinary",
        default_value = "{\"enabled\": false, \"isEncrypted\": true}"
    )]
    pub single_use: Option<Json>,
    #[sea_orm(column_name = "isVerifyEmailEnabled", default_value = false)]
    pub is_verify_email_enabled: bool,
    #[sea_orm(column_name = "isBackButtonHidden", default_value = false)]
    pub is_back_button_hidden: bool,
    #[sea_orm(column_name = "isAutoProgressingEnabled", default_value = false)]
    pub is_auto_progressing_enabled: bool,
    #[sea_orm(column_name = "isCaptureIpEnabled", default_value = false)]
    pub is_capture_ip_enabled: bool,
    #[sea_orm(default_value = false)]
    pub is_anonymize_responses_enabled: bool,
    #[sea_orm(column_type = "Text")]
    pub pin: Option<String>,
    #[sea_orm(
        column_name = "displayPercentage",
        column_type = "Decimal(Some((65, 30)))"
    )]
    pub display_percentage: Option<Decimal>,
    #[sea_orm(column_name = "showLanguageSwitch")]
    pub show_language_switch: Option<bool>,
    /// Open the survey in the respondent's browser language when it matches an enabled survey
    /// language. Null on surveys that predate the setting, which reads as off.
    #[sea_orm(column_name = "autoSelectLanguage")]
    pub auto_select_language: Option<bool>,
    /// [SurveyRecaptcha]
    #[sea_orm(
        column_type = "JsonBinary",
        default_value = "{\"enabled\": false, \"threshold\":0.1}"
    )]
    pub recaptcha: Option<Json>,
    /// [SurveyLinkMetadata]
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub metadata: Json,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub slug: Option<String>,
    #[sea_orm(column_name = "customHeadScripts", column_type = "Text")]
    pub custom_head_scripts: Option<String>,
    #[sea_orm(column_name = "customHeadScriptsMode", default_value = "add")]
    pub custom_head_scripts_mode: Option<SurveyScriptMode>,
    /// [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
    /// by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
    #[sea_orm(column_name = "customCss", column_type = "JsonBinary")]
    pub custom_css: Option<Json>,
    /// [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
    /// (ENG-3873). null = no conditions, which is what every existing survey keeps.
    #[sea_orm(column_name = "showConditions", column_type = "JsonBinary")]
    pub show_conditions: Option<Json>,
    /// [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
    /// survey on load. null = none, which is what every existing survey keeps.
    #[sea_orm(column_name = "startLogic", column_type = "JsonBinary")]
    pub start_logic: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Survey_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        fk_name = "Survey_createdBy_fkey"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::OwnerId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "Survey_ownerId_fkey"
    )]
    Owner,
    #[sea_orm(
        belongs_to = "super::segment::Entity",
        from = "Column::SegmentId",
        to = "super::segment::Column::Id",
        fk_name = "Survey_segmentId_fkey"
    )]
    Segment,
    #[sea_orm(has_many = "super::response::Entity")]
    Responses,
    #[sea_orm(has_many = "super::survey_quota::Entity")]
    Quotas,
    #[sea_orm(has_many = "super::survey_trigger::Entity")]
    Triggers,
    #[sea_orm(has_many = "super::survey_attribute_filter::Entity")]
    AttributeFilters,
    #[sea_orm(has_many = "super::display::Entity")]
    Displays,
    #[sea_orm(has_many = "super::survey_language::Entity")]
    Languages,
    #[sea_orm(has_many = "super::embedded_data::Entity")]
    EmbeddedData,
    #[sea_orm(has_many = "super::survey_embedded_data::Entity")]
    EmbeddedDataLinks,
    #[sea_orm(has_many = "super::survey_follow_up::Entity")]
    FollowUps,
    #[sea_orm(has_many = "super::feedback_source_formbricks_mapping::Entity")]
    FeedbackSourceMappings,
    #[sea_orm(has_many = "super::retention_exemption::Entity")]
    RetentionExemptions,
    #[sea_orm(has_many = "super::retention_notice::Entity")]
    RetentionNotices,
    #[sea_orm(has_many = "super::creative::Entity")]
    Creatives,
    #[sea_orm(has_many = "super::attachment_export::Entity")]
    AttachmentExports,
    #[sea_orm(has_one = "super::retention_survey_restore::Entity")]
    RetentionRestore,
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

impl Related<super::segment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Segment.def()
    }
}

impl Related<super::response::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Responses.def()
    }
}

impl Related<super::survey_quota::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Quotas.def()
    }
}

impl Related<super::survey_trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Triggers.def()
    }
}

impl Related<super::survey_attribute_filter::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttributeFilters.def()
    }
}

impl Related<super::display::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Displays.def()
    }
}

impl Related<super::survey_language::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Languages.def()
    }
}

impl Related<super::embedded_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmbeddedData.def()
    }
}

impl Related<super::survey_embedded_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmbeddedDataLinks.def()
    }
}

impl Related<super::survey_follow_up::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FollowUps.def()
    }
}

impl Related<super::feedback_source_formbricks_mapping::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSourceMappings.def()
    }
}

impl Related<super::retention_exemption::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RetentionExemptions.def()
    }
}

impl Related<super::retention_notice::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RetentionNotices.def()
    }
}

impl Related<super::creative::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creatives.def()
    }
}

impl Related<super::attachment_export::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttachmentExports.def()
    }
}

impl Related<super::retention_survey_restore::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RetentionRestore.def()
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
