use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "topic_clustering_runs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub facet_id: String,
    pub facet_version: i32,
    #[sea_orm(column_type = "Text", default_value = "pending")]
    pub status: String,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub config: Json,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamptz(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTimeWithTimeZone,
    #[sea_orm(column_type = "custom(\"timestamptz(3)\")")]
    pub started_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "custom(\"timestamptz(3)\")")]
    pub finished_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub topic_version_ids: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub rule_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub batch_action_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "topic_clustering_runs_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::evaluator_version::Entity",
        from = "(Column::FacetId, Column::FacetVersion)",
        to = "(super::evaluator_version::Column::EvaluatorId, super::evaluator_version::Column::Version)",
        on_delete = "Cascade",
        fk_name = "topic_clustering_runs_facet_id_facet_version_fkey"
    )]
    Facet,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::evaluator_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Facet.def()
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
            self.created_at = Set(now.fixed_offset());
        }
        Ok(self)
    }
}
