use super::mail_split_filter_kind::MailSplitFilterKind;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// One condition row in the split builder. A split may hold several of the same
/// kind (two senders, two labels), so kind is deliberately not unique per split.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "MailSplitFilter")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub kind: MailSplitFilterKind,
    #[sea_orm(column_type = "Text")]
    pub value: Option<String>,
    pub order: i32,
    #[sea_orm(column_name = "mailSplitId", column_type = "Text")]
    pub mail_split_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::mail_split::Entity",
        from = "Column::MailSplitId",
        to = "super::mail_split::Column::Id",
        on_delete = "Cascade",
        fk_name = "MailSplitFilter_mailSplitId_fkey"
    )]
    MailSplit,
}

impl Related<super::mail_split::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MailSplit.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
