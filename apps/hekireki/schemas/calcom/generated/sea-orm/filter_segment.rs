use super::filter_segment_scope::FilterSegmentScope;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FilterSegment")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "tableIdentifier", column_type = "Text")]
    pub table_identifier: String,
    pub scope: FilterSegmentScope,
    #[sea_orm(column_name = "activeFilters", column_type = "JsonBinary")]
    pub active_filters: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub sorting: Option<Json>,
    #[sea_orm(column_name = "columnVisibility", column_type = "JsonBinary")]
    pub column_visibility: Option<Json>,
    #[sea_orm(column_name = "columnSizing", column_type = "JsonBinary")]
    pub column_sizing: Option<Json>,
    #[sea_orm(column_name = "perPage")]
    pub per_page: i32,
    #[sea_orm(column_name = "searchTerm", column_type = "Text")]
    pub search_term: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "FilterSegment_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "FilterSegment_teamId_fkey"
    )]
    Team,
    #[sea_orm(has_many = "super::user_filter_segment_preference::Entity")]
    UserPreferences,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::user_filter_segment_preference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserPreferences.def()
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
