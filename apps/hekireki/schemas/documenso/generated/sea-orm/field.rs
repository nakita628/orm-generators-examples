use super::field_type::FieldType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// @zod.import(["import { ZFieldMetaNotOptionalSchema } from '@documenso/lib/types/field-meta';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Field")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "secondaryId", column_type = "Text")]
    pub secondary_id: String,
    #[sea_orm(column_name = "externalId", column_type = "String(StringLen::N(255))")]
    pub external_id: Option<String>,
    #[sea_orm(column_name = "groupId", column_type = "String(StringLen::N(255))")]
    pub group_id: Option<String>,
    #[sea_orm(column_name = "envelopeId", column_type = "Text")]
    pub envelope_id: String,
    #[sea_orm(column_name = "envelopeItemId", column_type = "Text")]
    pub envelope_item_id: String,
    #[sea_orm(column_name = "recipientId")]
    pub recipient_id: i32,
    pub r#type: FieldType,
    /// @zod.number.describe("The page number of the field on the document. Starts from 1.")
    pub page: i32,
    #[sea_orm(
        column_name = "positionX",
        column_type = "Decimal(Some((65, 30)))",
        default_value = 0
    )]
    pub position_x: Decimal,
    #[sea_orm(
        column_name = "positionY",
        column_type = "Decimal(Some((65, 30)))",
        default_value = 0
    )]
    pub position_y: Decimal,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))", default_value = -1)]
    pub width: Decimal,
    #[sea_orm(column_type = "Decimal(Some((65, 30)))", default_value = -1)]
    pub height: Decimal,
    #[sea_orm(column_name = "customText", column_type = "Text")]
    pub custom_text: String,
    pub inserted: bool,
    /// [FieldMeta] @zod.custom.use(ZFieldMetaNotOptionalSchema)
    #[sea_orm(column_name = "fieldMeta", column_type = "JsonBinary")]
    pub field_meta: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::envelope_item::Entity",
        from = "Column::EnvelopeItemId",
        to = "super::envelope_item::Column::Id",
        on_delete = "Cascade",
        fk_name = "Field_envelopeItemId_fkey"
    )]
    EnvelopeItem,
    #[sea_orm(
        belongs_to = "super::envelope::Entity",
        from = "Column::EnvelopeId",
        to = "super::envelope::Column::Id",
        on_delete = "Cascade",
        fk_name = "Field_envelopeId_fkey"
    )]
    Envelope,
    #[sea_orm(
        belongs_to = "super::recipient::Entity",
        from = "Column::RecipientId",
        to = "super::recipient::Column::Id",
        on_delete = "Cascade",
        fk_name = "Field_recipientId_fkey"
    )]
    Recipient,
    #[sea_orm(has_one = "super::signature::Entity")]
    Signature,
}

impl Related<super::envelope_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvelopeItem.def()
    }
}

impl Related<super::envelope::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Envelope.def()
    }
}

impl Related<super::recipient::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Recipient.def()
    }
}

impl Related<super::signature::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Signature.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
