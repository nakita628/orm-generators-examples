use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Track")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    pub name: String,
    #[sea_orm(unique)]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
    /// determines if a track can be seen on tracks page
    #[sea_orm(default_value = true)]
    pub visible: bool,
    /// use this to display coming soon UI
    #[sea_orm(column_name = "isComingSoon", default_value = false)]
    pub is_coming_soon: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::track_challenge::Entity")]
    TrackChallenges,
}

impl Related<super::track_challenge::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TrackChallenges.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        super::track_to_user::Relation::User.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::track_to_user::Relation::Track.def().rev())
    }
}

impl ActiveModelBehavior for ActiveModel {}
