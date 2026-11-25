use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a supported language in the system.
/// Used for multilingual survey support.
///
/// @property id - Unique identifier for the language
/// @property code - Language code (e.g., 'en-US')
/// @property alias - Optional friendly name
/// @property workspace - The workspace this language is enabled for
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Language")]
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
    pub code: String,
    #[sea_orm(column_type = "Text")]
    pub alias: Option<String>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Language_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(has_many = "super::survey_language::Entity")]
    SurveyLanguages,
    #[sea_orm(has_many = "super::creative_language::Entity")]
    CreativeLanguages,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::survey_language::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SurveyLanguages.def()
    }
}

impl Related<super::creative_language::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreativeLanguages.def()
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
