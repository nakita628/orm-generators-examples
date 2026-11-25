use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScenarioVersion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "scenarioId", column_type = "Text")]
    pub scenario_id: String,
    /// The project of the parent case, carried here so every read filters by
    /// tenant without a join. The service is what keeps the pair honest: no
    /// composite foreign key can, under relationMode="prisma".
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    pub version: i32,
    #[sea_orm(column_name = "authorId", column_type = "Text")]
    pub author_id: Option<String>,
    /// Who wrote this version: "user", "api", "cli" or "langy".
    #[sea_orm(column_name = "authorLabel", column_type = "Text")]
    pub author_label: String,
    #[sea_orm(column_name = "changeDescription", column_type = "Text")]
    pub change_description: Option<String>,
    /// Envelope { schemaVersion, fields, changedFields } holding the editable
    /// field set as saved and the names of the fields this save changed.
    #[sea_orm(column_type = "JsonBinary")]
    pub snapshot: Json,
    /// Schema version of `snapshot`, so a later shape change can be read forward.
    #[sea_orm(column_name = "schemaVersion", default_value = 1)]
    pub schema_version: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::scenario::Entity",
        from = "Column::ScenarioId",
        to = "super::scenario::Column::Id",
        on_delete = "Cascade",
        fk_name = "ScenarioVersion_scenarioId_fkey"
    )]
    Scenario,
}

impl Related<super::scenario::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scenario.def()
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
