use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_WorkflowEvent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub event_id: String,
    #[sea_orm(column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_type = "Text")]
    pub event_type: String,
    #[sea_orm(column_type = "Text")]
    pub step_name: String,
    pub sequence_number: i32,
    #[sea_orm(column_type = "JsonBinary")]
    pub data: Option<Json>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_workflow_run::Entity",
        from = "Column::RunId",
        to = "super::lite_llm_workflow_run::Column::RunId",
        fk_name = "LiteLLM_WorkflowEvent_run_id_fkey"
    )]
    Run,
}

impl Related<super::lite_llm_workflow_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            event_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
