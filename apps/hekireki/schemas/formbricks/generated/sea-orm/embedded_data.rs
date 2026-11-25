use super::embedded_data_source::EmbeddedDataSource;
use super::embedded_data_type::EmbeddedDataType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Defines one Embedded Data field: the extra context stored on a response that isn't an
/// answer to a question (brand, plan, page type, a computed score, ...).
///
/// A field is either **local** — used by a single survey, which is how existing variables and
/// hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
/// be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
/// is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
/// other way round. Deriving it rather than storing a boolean removes a column that could drift
/// out of step with the two fields it stands for.
///
/// @property key - Library name, unique per workspace and immutable. Null for local fields
/// @property name - Human-readable display label
/// @property source - Where the value comes from (computed | ingested). `reserved` is part of the
/// shared enum but never stored as a row: those fields are a static catalog in code
/// @property dataType - Declared value type, used for coercion at ingest
/// @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
/// @property locked - Ingested only: ignore external writes and use the default instead
/// @property survey - The owning survey for a local field; cascades so nothing is orphaned
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EmbeddedData")]
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
    pub key: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    pub source: EmbeddedDataSource,
    #[sea_orm(column_name = "dataType", default_value = "string")]
    pub data_type: EmbeddedDataType,
    /// [EmbeddedDataDefaultValue]
    #[sea_orm(column_name = "defaultValue", column_type = "JsonBinary")]
    pub default_value: Option<Json>,
    #[sea_orm(default_value = false)]
    pub locked: bool,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: Option<String>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "(Column::SurveyId, Column::WorkspaceId)",
        to = "(super::survey::Column::Id, super::survey::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "EmbeddedData_surveyId_workspaceId_fkey"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "EmbeddedData_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(has_many = "super::survey_embedded_data::Entity")]
    SurveyLinks,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::survey_embedded_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SurveyLinks.def()
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
