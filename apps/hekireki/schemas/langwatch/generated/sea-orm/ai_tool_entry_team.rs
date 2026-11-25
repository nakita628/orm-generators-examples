use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// @deprecated Per-team scope binding for AiToolEntry. Superseded by
/// AiToolEntryDepartment; kept so the rename migration is non-destructive.
/// No new rows are written here.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AiToolEntryTeam")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "entryId", column_type = "Text")]
    pub entry_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::ai_tool_entry::Entity",
        from = "Column::EntryId",
        to = "super::ai_tool_entry::Column::Id",
        on_delete = "Cascade",
        fk_name = "AiToolEntryTeam_entryId_fkey"
    )]
    Entry,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "AiToolEntryTeam_teamId_fkey"
    )]
    Team,
}

impl Related<super::ai_tool_entry::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Entry.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
