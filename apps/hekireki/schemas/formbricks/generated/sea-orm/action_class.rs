use super::action_type::ActionType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a user action that can trigger surveys.
/// Used to define points in the user journey where surveys can be shown.
///
/// @property id - Unique identifier for the action
/// @property name - Display name of the action
/// @property type - Whether this is a code or no-code action
/// @property key - Unique identifier used in code implementation
/// @property noCodeConfig - Configuration for no-code setup
/// @property workspace - The workspace this action belongs to
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ActionClass")]
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
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    pub r#type: ActionType,
    #[sea_orm(column_type = "Text")]
    pub key: Option<String>,
    /// [ActionClassNoCodeConfig]
    #[sea_orm(column_name = "noCodeConfig", column_type = "JsonBinary")]
    pub no_code_config: Option<Json>,
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
        fk_name = "ActionClass_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(has_many = "super::survey_trigger::Entity")]
    SurveyTriggers,
    #[sea_orm(has_many = "super::creative_trigger::Entity")]
    CreativeTriggers,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::survey_trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SurveyTriggers.def()
    }
}

impl Related<super::creative_trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreativeTriggers.def()
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
