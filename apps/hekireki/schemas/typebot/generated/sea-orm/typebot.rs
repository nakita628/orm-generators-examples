use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Typebot")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub version: Option<String>,
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
    #[sea_orm(column_type = "Text")]
    pub icon: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "folderId", column_type = "Text")]
    pub folder_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub groups: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub variables: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub edges: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub theme: Json,
    #[sea_orm(column_name = "selectedThemeTemplateId", column_type = "Text")]
    pub selected_theme_template_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub settings: Json,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "publicId", column_type = "Text")]
    pub public_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "customDomain", column_type = "Text")]
    pub custom_domain: Option<String>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "resultsTablePreferences", column_type = "JsonBinary")]
    pub results_table_preferences: Option<Json>,
    #[sea_orm(column_name = "isArchived", default_value = false)]
    pub is_archived: bool,
    #[sea_orm(column_name = "isClosed", default_value = false)]
    pub is_closed: bool,
    #[sea_orm(column_name = "whatsAppCredentialsId", column_type = "Text")]
    pub whats_app_credentials_id: Option<String>,
    #[sea_orm(column_name = "riskLevel")]
    pub risk_level: Option<i32>,
    #[sea_orm(column_name = "spaceId", column_type = "Text")]
    pub space_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::dashboard_folder::Entity",
        from = "Column::FolderId",
        to = "super::dashboard_folder::Column::Id",
        fk_name = "Typebot_folderId_fkey"
    )]
    Folder,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Typebot_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::space::Entity",
        from = "Column::SpaceId",
        to = "super::space::Column::Id",
        on_delete = "SetNull",
        fk_name = "Typebot_spaceId_fkey"
    )]
    Space,
    #[sea_orm(has_many = "super::collaborators_on_typebots::Entity")]
    Collaborators,
    #[sea_orm(has_many = "super::invitation::Entity")]
    Invitations,
    #[sea_orm(has_many = "super::result::Entity")]
    Results,
    #[sea_orm(has_many = "super::webhook::Entity")]
    Webhooks,
    #[sea_orm(has_many = "super::banned_ip::Entity")]
    BannedIps,
    #[sea_orm(has_one = "super::public_typebot::Entity")]
    PublishedTypebot,
}

impl Related<super::dashboard_folder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Folder.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::space::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Space.def()
    }
}

impl Related<super::collaborators_on_typebots::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Collaborators.def()
    }
}

impl Related<super::invitation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Invitations.def()
    }
}

impl Related<super::result::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Results.def()
    }
}

impl Related<super::webhook::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhooks.def()
    }
}

impl Related<super::banned_ip::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BannedIps.def()
    }
}

impl Related<super::public_typebot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PublishedTypebot.def()
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
