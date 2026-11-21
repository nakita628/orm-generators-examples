use super::document_data_type::DocumentDataType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Generic stored data (e.g. an uploaded image) which other models reference.
/// The kind of data and its kind specific details live in `metadata`.
/// @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DataContent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    /// How `data` is stored, e.g. an S3 path or inline bytes.
    pub r#type: DocumentDataType,
    pub data: String,
    /// [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
    pub metadata: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::envelope_content::Entity")]
    EnvelopeContents,
}

impl Related<super::envelope_content::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EnvelopeContents.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
