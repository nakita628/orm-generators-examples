use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Deployment")]
pub struct Model {
    /// This is a single row table, so we use a fixed id
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: i32,
    pub logo: Option<String>,
    /// @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
    pub theme: Option<Json>,
    #[sea_orm(column_name = "licenseKey")]
    pub license_key: Option<String>,
    #[sea_orm(column_name = "signatureTokenEncrypted")]
    pub signature_token_encrypted: Option<String>,
    #[sea_orm(column_name = "agreedLicenseAt")]
    pub agreed_license_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
