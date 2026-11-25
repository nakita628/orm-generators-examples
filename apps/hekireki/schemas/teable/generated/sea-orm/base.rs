use super::provision_state::ProvisionState;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "base")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub space_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    pub order: f64,
    #[sea_orm(column_type = "Text")]
    pub icon: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub schema_pass: Option<String>,
    #[sea_orm(default_value = false)]
    pub v2_enabled: bool,
    #[sea_orm(default_value = "ready")]
    pub provision_state: ProvisionState,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub deleted_time: Option<DateTime>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(column_type = "Text")]
    pub last_modified_by: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_modified_time: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::space::Entity",
        from = "Column::SpaceId",
        to = "super::space::Column::Id",
        fk_name = "base_space_id_fkey"
    )]
    Space,
    #[sea_orm(has_many = "super::table_meta::Entity")]
    Tables,
    #[sea_orm(has_many = "super::base_data_db_move_job::Entity")]
    DataDbMoveJobs,
    #[sea_orm(has_many = "super::user_base_order::Entity")]
    UserOrders,
}

impl Related<super::space::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Space.def()
    }
}

impl Related<super::table_meta::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tables.def()
    }
}

impl Related<super::base_data_db_move_job::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DataDbMoveJobs.def()
    }
}

impl Related<super::user_base_order::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserOrders.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_time.is_not_set() {
            self.created_time = Set(now.naive_utc());
        }
        if !self.last_modified_time.is_set() {
            self.last_modified_time = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
