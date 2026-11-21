use super::audit_status::AuditStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AuditSession")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    pub description: Option<String>,
    #[sea_orm(column_name = "targetId")]
    pub target_id: Option<String>,
    #[sea_orm(default_value = "PENDING")]
    pub status: AuditStatus,
    #[sea_orm(column_name = "scopeMeta")]
    pub scope_meta: Option<Json>,
    #[sea_orm(column_name = "expectedAssetCount", default_value = 0)]
    pub expected_asset_count: i32,
    #[sea_orm(column_name = "foundAssetCount", default_value = 0)]
    pub found_asset_count: i32,
    #[sea_orm(column_name = "missingAssetCount", default_value = 0)]
    pub missing_asset_count: i32,
    #[sea_orm(column_name = "unexpectedAssetCount", default_value = 0)]
    pub unexpected_asset_count: i32,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "dueDate")]
    pub due_date: Option<DateTime>,
    #[sea_orm(column_name = "completedAt")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "cancelledAt")]
    pub cancelled_at: Option<DateTime>,
    #[sea_orm(column_name = "activeSchedulerReference")]
    pub active_scheduler_reference: Option<String>,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
    #[sea_orm(has_many = "super::audit_assignment::Entity")]
    Assignments,
    #[sea_orm(has_many = "super::audit_asset::Entity")]
    Assets,
    #[sea_orm(has_many = "super::audit_scan::Entity")]
    Scans,
    #[sea_orm(has_many = "super::audit_note::Entity")]
    Notes,
    #[sea_orm(has_many = "super::audit_image::Entity")]
    Images,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::audit_assignment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assignments.def()
    }
}

impl Related<super::audit_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assets.def()
    }
}

impl Related<super::audit_scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::audit_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
    }
}

impl Related<super::audit_image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Images.def()
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
