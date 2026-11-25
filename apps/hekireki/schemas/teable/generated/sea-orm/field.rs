use super::provision_state::ProvisionState;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "field")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub options: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub meta: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub ai_config: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub cell_value_type: String,
    pub is_multiple_cell_value: Option<bool>,
    #[sea_orm(column_type = "Text")]
    pub db_field_type: String,
    #[sea_orm(column_type = "Text")]
    pub db_field_name: String,
    #[sea_orm(default_value = "ready")]
    pub provision_state: ProvisionState,
    pub not_null: Option<bool>,
    pub unique: Option<bool>,
    pub is_primary: Option<bool>,
    pub is_computed: Option<bool>,
    pub is_lookup: Option<bool>,
    pub is_conditional_lookup: Option<bool>,
    pub is_pending: Option<bool>,
    pub has_error: Option<bool>,
    #[sea_orm(default_value = false)]
    pub needs_periodic_refresh: bool,
    #[sea_orm(column_type = "String(StringLen::N(16))")]
    pub refresh_cadence: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub lookup_linked_field_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub lookup_options: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub table_id: String,
    pub order: f64,
    pub version: i32,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_modified_time: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub deleted_time: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(column_type = "Text")]
    pub last_modified_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::table_meta::Entity",
        from = "Column::TableId",
        to = "super::table_meta::Column::Id",
        fk_name = "field_table_id_fkey"
    )]
    Table,
    #[sea_orm(has_one = "super::field_periodic_refresh_state::Entity")]
    PeriodicRefreshState,
}

impl Related<super::table_meta::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Table.def()
    }
}

impl Related<super::field_periodic_refresh_state::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PeriodicRefreshState.def()
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
