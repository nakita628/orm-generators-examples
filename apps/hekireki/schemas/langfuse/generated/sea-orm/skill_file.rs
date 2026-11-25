use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "skill_files")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub skill_id: String,
    #[sea_orm(column_type = "Text")]
    pub blob_id: String,
    #[sea_orm(column_type = "Text")]
    pub path: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "skill_files_project_id_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::skill::Entity",
        from = "Column::SkillId",
        to = "super::skill::Column::Id",
        on_delete = "Cascade",
        fk_name = "skill_files_skill_id_fkey"
    )]
    Skill,
    #[sea_orm(
        belongs_to = "super::skill_blob::Entity",
        from = "Column::BlobId",
        to = "super::skill_blob::Column::Id",
        on_delete = "Restrict",
        fk_name = "skill_files_blob_id_fkey"
    )]
    Blob,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::skill::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Skill.def()
    }
}

impl Related<super::skill_blob::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Blob.def()
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
