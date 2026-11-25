use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AttributeToUser")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "memberId")]
    pub member_id: i32,
    #[sea_orm(column_name = "attributeOptionId", column_type = "Text")]
    pub attribute_option_id: String,
    pub weight: Option<i32>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "createdById")]
    pub created_by_id: Option<i32>,
    #[sea_orm(column_name = "createdByDSyncId", column_type = "Text")]
    pub created_by_dsync_id: Option<String>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedById")]
    pub updated_by_id: Option<i32>,
    #[sea_orm(column_name = "updatedByDSyncId", column_type = "Text")]
    pub updated_by_dsync_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::membership::Entity",
        from = "Column::MemberId",
        to = "super::membership::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttributeToUser_memberId_fkey"
    )]
    Member,
    #[sea_orm(
        belongs_to = "super::attribute_option::Entity",
        from = "Column::AttributeOptionId",
        to = "super::attribute_option::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttributeToUser_attributeOptionId_fkey"
    )]
    AttributeOption,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "AttributeToUser_createdById_fkey"
    )]
    CreatedBy,
    #[sea_orm(
        belongs_to = "super::dsync_data::Entity",
        from = "Column::CreatedByDsyncId",
        to = "super::dsync_data::Column::DirectoryId",
        on_delete = "SetNull",
        fk_name = "AttributeToUser_createdByDSyncId_fkey"
    )]
    CreatedByDSync,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UpdatedById",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "AttributeToUser_updatedById_fkey"
    )]
    UpdatedBy,
    #[sea_orm(
        belongs_to = "super::dsync_data::Entity",
        from = "Column::UpdatedByDsyncId",
        to = "super::dsync_data::Column::DirectoryId",
        on_delete = "SetNull",
        fk_name = "AttributeToUser_updatedByDSyncId_fkey"
    )]
    UpdatedByDSync,
}

impl Related<super::membership::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Member.def()
    }
}

impl Related<super::attribute_option::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttributeOption.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedBy.def()
    }
}

impl Related<super::dsync_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByDSync.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
