use super::survey_overlay::SurveyOverlay;
use super::widget_placement::WidgetPlacement;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Main grouping mechanism for resources in Formbricks.
/// Each organization can have multiple workspaces to separate different applications or products.
///
/// @property id - Unique identifier for the workspace
/// @property name - Display name of the workspace
/// @property organization - Reference to parent organization
/// @property surveys - Collection of surveys in this workspace
/// @property styling - Workspace-wide styling configuration
/// @property config - Workspace-specific configuration
/// @property recontactDays - Default recontact delay for surveys
/// @property placement - Default widget placement for in-app surveys
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Workspace")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "legacyEnvironmentId")]
    pub legacy_environment_id: Option<String>,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// [Styling]
    #[sea_orm(default_value = "{\"allowStyleOverwrite\":true}")]
    pub styling: Json,
    /// [WorkspaceConfig]
    #[sea_orm(default_value = "{}")]
    pub config: Json,
    #[sea_orm(column_name = "recontactDays", default_value = 7)]
    pub recontact_days: i32,
    #[sea_orm(column_name = "linkSurveyBranding", default_value = true)]
    pub link_survey_branding: bool,
    #[sea_orm(column_name = "inAppSurveyBranding", default_value = true)]
    pub in_app_survey_branding: bool,
    #[sea_orm(default_value = "bottomRight")]
    pub placement: WidgetPlacement,
    #[sea_orm(column_name = "clickOutsideClose", default_value = true)]
    pub click_outside_close: bool,
    #[sea_orm(default_value = "none")]
    pub overlay: SurveyOverlay,
    /// Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`.
    #[sea_orm(column_name = "overlayColor")]
    pub overlay_color: Option<String>,
    #[sea_orm(column_name = "overlayOpacity")]
    pub overlay_opacity: Option<i32>,
    /// [Logo]
    pub logo: Option<Json>,
    #[sea_orm(column_name = "appSetupCompleted", default_value = false)]
    pub app_setup_completed: bool,
    #[sea_orm(column_name = "customHeadScripts")]
    pub custom_head_scripts: Option<String>,
    /// [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service.
    #[sea_orm(column_name = "customCss")]
    pub custom_css: Option<Json>,
    /// [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
    /// submitting its source again through the current processor and authorization checks.
    #[sea_orm(column_name = "customCssPrevious")]
    pub custom_css_previous: Option<Json>,
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
    #[sea_orm(has_many = "super::language::Entity")]
    Languages,
    #[sea_orm(has_many = "super::embedded_data::Entity")]
    EmbeddedData,
    #[sea_orm(has_many = "super::workspace_team::Entity")]
    WorkspaceTeams,
    #[sea_orm(has_many = "super::feedback_directory_workspace::Entity")]
    FeedbackDirectoryWorkspaces,
    #[sea_orm(has_many = "super::chart::Entity")]
    Charts,
    #[sea_orm(has_many = "super::dashboard::Entity")]
    Dashboards,
    #[sea_orm(has_many = "super::embedded_field_label::Entity")]
    EmbeddedFieldLabels,
    #[sea_orm(has_many = "super::creative::Entity")]
    Creatives,
    #[sea_orm(has_many = "super::attachment_export::Entity")]
    AttachmentExports,
    #[sea_orm(has_many = "super::survey::Entity")]
    Surveys,
    #[sea_orm(has_many = "super::contact::Entity")]
    Contacts,
    #[sea_orm(has_many = "super::action_class::Entity")]
    ActionClasses,
    #[sea_orm(has_many = "super::contact_attribute_key::Entity")]
    ContactAttributeKeys,
    #[sea_orm(has_many = "super::webhook::Entity")]
    Webhooks,
    #[sea_orm(has_many = "super::tag::Entity")]
    Tags,
    #[sea_orm(has_many = "super::segment::Entity")]
    Segments,
    #[sea_orm(has_many = "super::integration::Entity")]
    Integrations,
    #[sea_orm(has_many = "super::api_key_workspace::Entity")]
    ApiKeyWorkspaces,
    #[sea_orm(has_many = "super::feedback_source::Entity")]
    FeedbackSources,
    #[sea_orm(has_many = "super::workflow::Entity")]
    Workflows,
    #[sea_orm(has_many = "super::workflow_run::Entity")]
    WorkflowRuns,
    #[sea_orm(has_many = "super::mcp_grant_workspace::Entity")]
    McpGrants,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::language::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Languages.def()
    }
}

impl Related<super::embedded_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmbeddedData.def()
    }
}

impl Related<super::workspace_team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkspaceTeams.def()
    }
}

impl Related<super::feedback_directory_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackDirectoryWorkspaces.def()
    }
}

impl Related<super::chart::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Charts.def()
    }
}

impl Related<super::dashboard::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dashboards.def()
    }
}

impl Related<super::embedded_field_label::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmbeddedFieldLabels.def()
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

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Surveys.def()
    }
}

impl Related<super::contact::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contacts.def()
    }
}

impl Related<super::action_class::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ActionClasses.def()
    }
}

impl Related<super::contact_attribute_key::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ContactAttributeKeys.def()
    }
}

impl Related<super::webhook::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhooks.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tags.def()
    }
}

impl Related<super::segment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Segments.def()
    }
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integrations.def()
    }
}

impl Related<super::api_key_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiKeyWorkspaces.def()
    }
}

impl Related<super::feedback_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSources.def()
    }
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflows.def()
    }
}

impl Related<super::workflow_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkflowRuns.def()
    }
}

impl Related<super::mcp_grant_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::McpGrants.def()
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
