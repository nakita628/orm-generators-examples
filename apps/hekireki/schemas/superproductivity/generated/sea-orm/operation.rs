use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "operations")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub user_id: i32,
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    pub server_seq: i32,
    #[sea_orm(column_type = "Text")]
    pub action_type: String,
    #[sea_orm(column_type = "Text")]
    pub op_type: String,
    #[sea_orm(column_type = "Text")]
    pub entity_type: String,
    #[sea_orm(column_type = "Text")]
    pub entity_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub entity_ids: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Json,
    #[sea_orm(default_value = "0")]
    pub payload_bytes: i64,
    #[sea_orm(column_type = "JsonBinary")]
    pub vector_clock: Json,
    pub schema_version: i32,
    pub client_timestamp: i64,
    pub received_at: i64,
    #[sea_orm(default_value = false)]
    pub is_payload_encrypted: bool,
    #[sea_orm(column_type = "Text")]
    pub sync_import_reason: Option<String>,
    pub repair_base_server_seq: Option<i32>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "operations_user_id_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
