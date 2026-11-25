use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Links a survey to an Embedded Data field and records the key that field's value is stored
/// under inside that survey.
///
/// `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
/// independent, so nothing would stop a survey in one workspace linking a field defined in another.
/// Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
/// puts the workspace scope every query needs on the row instead of behind a join through Survey.
///
/// @property storageKey - The key the value lives under: a cuid for computed fields, the field
/// name for ingested ones. Migrated fields keep their original key, which is what lets existing
/// recall tokens, logic conditions and stored responses keep resolving untouched.
///
/// @property order - Position within this survey's field list, dense from 0. Lives on the link
/// rather than on `EmbeddedData` because a shared definition sits at a different position in each
/// survey that links it. Deliberately has no default: every writer knows the position it is writing,
/// so a missing one should be a type error rather than a row that silently sorts to the front.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SurveyEmbeddedData")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "surveyId", column_type = "Text")]
    pub survey_id: String,
    #[sea_orm(column_name = "embeddedDataId", column_type = "Text")]
    pub embedded_data_id: String,
    #[sea_orm(column_name = "storageKey", column_type = "Text")]
    pub storage_key: String,
    pub order: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::survey::Entity",
        from = "(Column::SurveyId, Column::WorkspaceId)",
        to = "(super::survey::Column::Id, super::survey::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "SurveyEmbeddedData_surveyId_workspaceId_fkey"
    )]
    Survey,
    #[sea_orm(
        belongs_to = "super::embedded_data::Entity",
        from = "(Column::EmbeddedDataId, Column::WorkspaceId)",
        to = "(super::embedded_data::Column::Id, super::embedded_data::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "SurveyEmbeddedData_embeddedDataId_workspaceId_fkey"
    )]
    EmbeddedData,
}

impl Related<super::survey::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Survey.def()
    }
}

impl Related<super::embedded_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmbeddedData.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
