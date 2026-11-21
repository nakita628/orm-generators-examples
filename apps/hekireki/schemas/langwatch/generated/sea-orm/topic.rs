use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// PROJECTION of the topic-clustering event stream (ADR-051): topics are
/// `topics_recorded` events; this table is written ONLY by
/// TopicModelFoldProjection (cursor: TopicModelProjection) and is
/// rebuildable by replay. Kept named `Topic` because it is the domain
/// entity every surface reads — the `*Projection` suffix is reserved for
/// cursor/read-model tables.
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Topic")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    pub name: String,
    #[sea_orm(column_name = "parentId")]
    pub parent_id: Option<String>,
    pub embeddings_model: String,
    pub centroid: Json,
    #[sea_orm(column_name = "p95Distance")]
    pub p95_distance: f64,
    #[sea_orm(column_name = "automaticallyGenerated", default_value = true)]
    pub automatically_generated: bool,
    /// The `topics_recorded` event that recorded this topic (its id in the
    /// event log). Null means the row predates event ownership and has not
    /// been seeded onto the stream yet; the projection sets it on every row
    /// it writes.
    #[sea_orm(column_name = "lastEventId")]
    pub last_event_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::topic::Entity",
        from = "Column::ParentId",
        to = "super::topic::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict"
    )]
    Parent,
    #[sea_orm(
        has_many = "super::topic::Entity",
        from = "Column::Id",
        to = "super::topic::Column::ParentId"
    )]
    Subtopics,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::topic::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Parent.def()
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
