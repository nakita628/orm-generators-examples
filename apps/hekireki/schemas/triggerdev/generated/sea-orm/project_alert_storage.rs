use super::project_alert_type::ProjectAlertType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ProjectAlertStorage")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "alertChannelId", column_type = "Text")]
    pub alert_channel_id: String,
    #[sea_orm(column_name = "alertType")]
    pub alert_type: ProjectAlertType,
    #[sea_orm(column_name = "storageId", column_type = "Text")]
    pub storage_id: String,
    #[sea_orm(column_name = "storageData", column_type = "JsonBinary")]
    pub storage_data: Json,
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
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlertStorage_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::project_alert_channel::Entity",
        from = "Column::AlertChannelId",
        to = "super::project_alert_channel::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "ProjectAlertStorage_alertChannelId_fkey"
    )]
    AlertChannel,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::project_alert_channel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AlertChannel.def()
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
