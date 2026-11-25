use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "prompts")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "text")]
    pub r#type: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(column_name = "filePath", column_type = "Text")]
    pub file_path: Option<String>,
    #[sea_orm(column_name = "exportName", column_type = "Text")]
    pub export_name: Option<String>,
    #[sea_orm(column_name = "variableSchema", column_type = "JsonBinary")]
    pub variable_schema: Option<Json>,
    #[sea_orm(column_name = "defaultModel", column_type = "Text")]
    pub default_model: Option<String>,
    #[sea_orm(column_name = "defaultConfig", column_type = "JsonBinary")]
    pub default_config: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub tags: Vec<String>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
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
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "prompts_organizationId_fkey"
    )]
    Organization,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "prompts_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "prompts_runtimeEnvironmentId_fkey"
    )]
    RuntimeEnvironment,
    #[sea_orm(has_many = "super::prompt_version::Entity")]
    Versions,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
    }
}

impl Related<super::prompt_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Versions.def()
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
