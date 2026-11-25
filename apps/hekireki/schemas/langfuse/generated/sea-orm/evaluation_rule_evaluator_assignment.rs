use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "evaluation_rule_evaluator_assignments")]
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
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub evaluation_rule_id: String,
    #[sea_orm(column_type = "Text")]
    pub evaluator_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub variable_mapping: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "evaluation_rule_evaluator_assignments_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::evaluation_rule::Entity",
        from = "Column::EvaluationRuleId",
        to = "super::evaluation_rule::Column::Id",
        on_delete = "Cascade",
        fk_name = "evaluation_rule_evaluator_assignments_evaluation_rule_id_fkey"
    )]
    EvaluationRule,
    #[sea_orm(
        belongs_to = "super::evaluator::Entity",
        from = "Column::EvaluatorId",
        to = "super::evaluator::Column::Id",
        on_delete = "Cascade",
        fk_name = "evaluation_rule_evaluator_assignments_evaluator_id_fkey"
    )]
    Evaluator,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::evaluation_rule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EvaluationRule.def()
    }
}

impl Related<super::evaluator::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Evaluator.def()
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
