use super::pdf_template_kind::PdfTemplateKind;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A saved layout for a generated PDF, kept per team workspace under a name:
/// which columns, rows and header lines print, and in what order the rows
/// come. A row is a layout, not a generated file; the person generating a PDF
/// picks a template by name.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "PdfTemplate")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub name: String,
    #[sea_orm(default_value = "BOOKING")]
    pub kind: PdfTemplateKind,
    /// The layout's switches. Read through `parseBookingTemplateConfig`,
    /// which fills any key a stored row lacks from the Checklist template's
    /// values.
    #[sea_orm(default_value = "{}")]
    pub config: Json,
    /// The template the PDF dialog offers first. One per workspace and kind,
    /// enforced by a partial unique index in raw SQL (migration
    /// 20261103120000_add_pdf_templates). Names are also unique per workspace
    /// and kind ignoring case, through an expression index in the same
    /// migration.
    #[sea_orm(column_name = "isDefault", default_value = false)]
    pub is_default: bool,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
