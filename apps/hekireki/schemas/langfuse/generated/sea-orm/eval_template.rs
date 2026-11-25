use super::eval_template_source_code_language::EvalTemplateSourceCodeLanguage;
use super::eval_template_type::EvalTemplateType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "eval_templates")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub version: i32,
    #[sea_orm(column_type = "Text")]
    pub prompt: Option<String>,
    #[sea_orm(default_value = "LLM_AS_JUDGE")]
    pub r#type: EvalTemplateType,
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
    #[sea_orm(column_name = "output_schema", column_type = "JsonBinary")]
    pub output_definition: Option<Json>,
    #[sea_orm(column_type = "String(StringLen::N(262144))")]
    pub source_code: Option<String>,
    pub source_code_language: Option<EvalTemplateSourceCodeLanguage>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "eval_templates_project_id_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::job_configuration::Entity")]
    JobConfiguration,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::job_configuration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::JobConfiguration.def()
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
