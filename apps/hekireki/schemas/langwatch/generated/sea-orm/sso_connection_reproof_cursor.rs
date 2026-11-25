use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Operational round-robin position for ADR-123's domain re-proof sweep.
///
/// This is deliberately not a column on `SsoConnection`: that row is a
/// whole-row event projection and its fold is its only writer, while a
/// healthy re-read states no fact. Keeping the scheduling cursor beside the
/// projection preserves both rules and still lets a fixed-size sweep revisit
/// every connection, including one whose proof has started wavering.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoConnectionReproofCursor")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: String,
    #[sea_orm(
        column_name = "lastReproofAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_reproof_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::sso_connection::Entity",
        from = "Column::ConnectionId",
        to = "super::sso_connection::Column::Id",
        on_delete = "Cascade",
        fk_name = "SsoConnectionReproofCursor_connectionId_fkey"
    )]
    Connection,
}

impl Related<super::sso_connection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Connection.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
