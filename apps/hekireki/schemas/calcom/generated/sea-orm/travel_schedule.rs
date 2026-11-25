use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TravelSchedule")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_name = "timeZone", column_type = "Text")]
    pub time_zone: String,
    #[sea_orm(column_name = "startDate", column_type = "custom(\"timestamp(3)\")")]
    pub start_date: DateTime,
    #[sea_orm(column_name = "endDate", column_type = "custom(\"timestamp(3)\")")]
    pub end_date: Option<DateTime>,
    #[sea_orm(column_name = "prevTimeZone", column_type = "Text")]
    pub prev_time_zone: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "TravelSchedule_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
