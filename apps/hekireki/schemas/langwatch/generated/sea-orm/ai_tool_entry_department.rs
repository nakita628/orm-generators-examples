use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Per-department scope binding for AiToolEntry. Empty set = org-wide.
/// Non-empty = entry visible only to members of those departments (an
/// OrganizationUser whose `departmentId` is in the set). Mirrors
/// AiToolEntryTeam's shape. `departmentId` is a plain column ref
/// (relationMode="prisma"), matching how Department is referenced
/// elsewhere; the service validates the id belongs to the org.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AiToolEntryDepartment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "entryId")]
    pub entry_id: String,
    #[sea_orm(column_name = "departmentId")]
    pub department_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::ai_tool_entry::Entity",
        from = "Column::EntryId",
        to = "super::ai_tool_entry::Column::Id",
        on_delete = "Cascade"
    )]
    Entry,
}

impl Related<super::ai_tool_entry::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Entry.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
