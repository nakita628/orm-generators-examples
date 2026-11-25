use super::evaluator_source_code_language::EvaluatorSourceCodeLanguage;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "evaluator_versions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub evaluator_id: String,
    pub version: i32,
    #[sea_orm(column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub prompt: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub prompt_messages: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub partner: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub model: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub provider: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub model_params: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub vars: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub variable_mapping: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub output_definition: Option<Json>,
    #[sea_orm(column_type = "String(StringLen::N(262144))")]
    pub source_code: Option<String>,
    pub source_code_language: Option<EvaluatorSourceCodeLanguage>,
    #[sea_orm(column_type = "JsonBinary")]
    pub questions: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::evaluator::Entity",
        from = "Column::EvaluatorId",
        to = "super::evaluator::Column::Id",
        on_delete = "Cascade",
        fk_name = "evaluator_versions_evaluator_id_fkey"
    )]
    Evaluator,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "evaluator_versions_created_by_user_id_fkey"
    )]
    CreatedByUser,
    #[sea_orm(has_many = "super::topic_clustering_run::Entity")]
    TopicClusteringRuns,
}

impl Related<super::evaluator::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Evaluator.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByUser.def()
    }
}

impl Related<super::topic_clustering_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TopicClusteringRuns.def()
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
        Ok(self)
    }
}
