use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EnvelopeItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    pub order: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "documentDataId", column_type = "Text")]
    pub document_data_id: String,
    #[sea_orm(column_name = "envelopeId", column_type = "Text")]
    pub envelope_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::document_data::Entity",
        from = "Column::DocumentDataId",
        to = "super::document_data::Column::Id",
        on_delete = "Cascade",
        fk_name = "EnvelopeItem_documentDataId_fkey"
    )]
    DocumentData,
    #[sea_orm(
        belongs_to = "super::envelope::Entity",
        from = "Column::EnvelopeId",
        to = "super::envelope::Column::Id",
        on_delete = "Cascade",
        fk_name = "EnvelopeItem_envelopeId_fkey"
    )]
    Envelope,
    #[sea_orm(has_many = "super::field::Entity")]
    Field,
    #[sea_orm(has_many = "super::envelope_content::Entity")]
    Contents,
}

impl Related<super::document_data::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DocumentData.def()
    }
}

impl Related<super::envelope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Envelope.def()
    }
}

impl Related<super::field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Field.def()
    }
}

impl Related<super::envelope_content::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contents.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
