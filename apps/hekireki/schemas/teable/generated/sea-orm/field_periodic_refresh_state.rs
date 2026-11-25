use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "field_periodic_refresh_state")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub field_id: String,
    #[sea_orm(column_type = "Text")]
    pub table_id: String,
    #[sea_orm(column_type = "Text")]
    pub base_id: String,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(column_type = "String(StringLen::N(64))", default_value = "UTC")]
    pub time_zone: String,
    #[sea_orm(default_value = false)]
    pub contains_now_precision: bool,
    #[sea_orm(column_type = "TimestampWithTimeZone")]
    pub last_refreshed_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_type = "String(StringLen::N(16))")]
    pub last_refreshed_local_date: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(128))")]
    pub last_range_hash: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(32))", default_value = "idle")]
    pub last_status: String,
    #[sea_orm(column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(64))")]
    pub pending_task_id: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(16))")]
    pub pending_local_date: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(128))")]
    pub pending_range_hash: Option<String>,
    #[sea_orm(column_type = "TimestampWithTimeZone")]
    pub pending_started_at: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_time: DateTimeWithTimeZone,
    #[sea_orm(column_type = "TimestampWithTimeZone")]
    pub last_modified_time: Option<DateTimeWithTimeZone>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::field::Entity",
        from = "Column::FieldId",
        to = "super::field::Column::Id",
        on_delete = "Cascade",
        fk_name = "field_periodic_refresh_state_field_id_fkey"
    )]
    Field,
}

impl Related<super::field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Field.def()
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
            self.created_time = Set(now.fixed_offset());
        }
        if !self.last_modified_time.is_set() {
            self.last_modified_time = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
