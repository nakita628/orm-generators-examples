use super::action_execution_status::ActionExecutionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "automation_executions")]
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
    pub source_id: String,
    #[sea_orm(column_type = "Text")]
    pub automation_id: String,
    #[sea_orm(column_type = "Text")]
    pub trigger_id: String,
    #[sea_orm(column_type = "Text")]
    pub action_id: String,
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: ActionExecutionStatus,
    #[sea_orm(column_type = "JsonBinary")]
    pub input: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub output: Option<Json>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub finished_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::automation::Entity",
        from = "Column::AutomationId",
        to = "super::automation::Column::Id",
        on_delete = "Cascade",
        fk_name = "automation_executions_automation_id_fkey"
    )]
    Automation,
    #[sea_orm(
        belongs_to = "super::trigger::Entity",
        from = "Column::TriggerId",
        to = "super::trigger::Column::Id",
        on_delete = "Cascade",
        fk_name = "automation_executions_trigger_id_fkey"
    )]
    Trigger,
    #[sea_orm(
        belongs_to = "super::action::Entity",
        from = "Column::ActionId",
        to = "super::action::Column::Id",
        on_delete = "Cascade",
        fk_name = "automation_executions_action_id_fkey"
    )]
    Action,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "automation_executions_project_id_fkey"
    )]
    Project,
}

impl Related<super::automation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Automation.def()
    }
}

impl Related<super::trigger::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Trigger.def()
    }
}

impl Related<super::action::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Action.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
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
