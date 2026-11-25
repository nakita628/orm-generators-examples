use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentRedactionJob")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "documentId", column_type = "Text")]
    pub document_id: String,
    #[sea_orm(column_name = "documentVersionId", column_type = "Text")]
    pub document_version_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: String,
    #[sea_orm(column_type = "Text", default_value = "PENDING")]
    pub status: String,
    #[sea_orm(
        column_name = "customTerms",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub custom_terms: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub reasons: Vec<String>,
    #[sea_orm(column_name = "resultVersionId", column_type = "Text")]
    pub result_version_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "detectRunId", column_type = "Text")]
    pub detect_run_id: Option<String>,
    #[sea_orm(column_name = "applyRunId", column_type = "Text")]
    pub apply_run_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::document::Entity",
        from = "Column::DocumentId",
        to = "super::document::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentRedactionJob_documentId_fkey"
    )]
    Document,
    #[sea_orm(
        belongs_to = "super::document_version::Entity",
        from = "Column::DocumentVersionId",
        to = "super::document_version::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentRedactionJob_documentVersionId_fkey"
    )]
    DocumentVersion,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentRedactionJob_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        fk_name = "DocumentRedactionJob_createdById_fkey"
    )]
    CreatedBy,
    #[sea_orm(has_many = "super::document_redaction::Entity")]
    Redactions,
}

impl Related<super::document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Document.def()
    }
}

impl Related<super::document_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DocumentVersion.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::document_redaction::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Redactions.def()
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
