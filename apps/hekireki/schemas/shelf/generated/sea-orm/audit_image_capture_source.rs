use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// How an audit photo reached Shelf, as reported by the client. The phone app
/// reports CAMERA only for its own camera, which cannot hand back an old
/// picture, so it is the strongest evidence short of a modified client; a
/// library photo's date is only what its file claims; the web app cannot tell
/// a camera shot from a file; PHONE is a phone app version that does not say
/// which.
#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(
    rs_type = "String",
    db_type = "Enum",
    enum_name = "AuditImageCaptureSource"
)]
pub enum AuditImageCaptureSource {
    #[sea_orm(string_value = "CAMERA")]
    Camera,
    #[sea_orm(string_value = "LIBRARY")]
    Library,
    #[sea_orm(string_value = "WEB_UPLOAD")]
    WebUpload,
    #[sea_orm(string_value = "PHONE")]
    Phone,
}
