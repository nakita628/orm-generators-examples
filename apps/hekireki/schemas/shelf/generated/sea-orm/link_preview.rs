use super::link_preview_outcome::LinkPreviewOutcome;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// What Shelf knows about one web link used in a workspace's custom fields,
/// keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
/// the link preview job (fetch columns) and people (customTitle).
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LinkPreview")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// `LinkKey` from @shelf/links, never a raw field value. At most 2048
    /// characters (CHECK constraint in the migration).
    pub url: String,
    #[sea_orm(column_name = "pageTitle")]
    pub page_title: Option<String>,
    /// Null until the first fetch finishes.
    pub outcome: Option<LinkPreviewOutcome>,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(column_name = "fetchedAt")]
    pub fetched_at: Option<DateTime>,
    /// When the link may be fetched next; null means never again. A claim moves
    /// it a few minutes ahead, which is also the lease on a fetch in flight.
    #[sea_orm(column_name = "nextFetchAt")]
    pub next_fetch_at: Option<DateTime>,
    /// The name a person gave the link in this workspace; wins over pageTitle.
    #[sea_orm(column_name = "customTitle")]
    pub custom_title: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.next_fetch_at.is_not_set() {
            self.next_fetch_at = Set(Some(now.naive_utc()));
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
