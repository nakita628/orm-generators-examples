use super::analytics_integration_export_source::AnalyticsIntegrationExportSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mixpanel_integrations")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub encrypted_mixpanel_project_token: String,
    #[sea_orm(column_type = "Text")]
    pub mixpanel_region: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_sync_at: Option<DateTime>,
    pub enabled: bool,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(default_value = "TRACES_OBSERVATIONS")]
    pub export_source: AnalyticsIntegrationExportSource,
    #[sea_orm(default_value = true)]
    pub backfill: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "mixpanel_integrations_project_id_fkey"
    )]
    Project,
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
        Ok(self)
    }
}
