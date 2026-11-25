use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkingHoursOverride")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Date")]
    pub date: Date,
    #[sea_orm(column_name = "isOpen", default_value = false)]
    pub is_open: bool,
    #[sea_orm(column_name = "openTime", column_type = "Text")]
    pub open_time: Option<String>,
    #[sea_orm(column_name = "closeTime", column_type = "Text")]
    pub close_time: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
    #[sea_orm(column_name = "workingHoursId", column_type = "Text")]
    pub working_hours_id: String,
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
        belongs_to = "super::working_hours::Entity",
        from = "Column::WorkingHoursId",
        to = "super::working_hours::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "WorkingHoursOverride_workingHoursId_fkey"
    )]
    WorkingHours,
}

impl Related<super::working_hours::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkingHours.def()
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
