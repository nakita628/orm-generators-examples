use super::document_filing_status::DocumentFilingStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentFiling")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "messageId", column_type = "Text")]
    pub message_id: String,
    #[sea_orm(column_name = "attachmentId", column_type = "Text")]
    pub attachment_id: String,
    #[sea_orm(column_type = "Text")]
    pub filename: String,
    #[sea_orm(column_name = "folderId", column_type = "Text")]
    pub folder_id: Option<String>,
    #[sea_orm(column_name = "folderPath", column_type = "Text")]
    pub folder_path: String,
    #[sea_orm(column_name = "fileId", column_type = "Text")]
    pub file_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub reasoning: Option<String>,
    pub confidence: Option<f64>,
    #[sea_orm(default_value = "FILED")]
    pub status: DocumentFilingStatus,
    #[sea_orm(column_name = "wasAsked", default_value = false)]
    pub was_asked: bool,
    #[sea_orm(column_name = "wasCorrected", default_value = false)]
    pub was_corrected: bool,
    #[sea_orm(column_name = "originalPath", column_type = "Text")]
    pub original_path: Option<String>,
    #[sea_orm(column_name = "correctedAt", column_type = "custom(\"timestamp(3)\")")]
    pub corrected_at: Option<DateTime>,
    #[sea_orm(column_name = "feedbackPositive")]
    pub feedback_positive: Option<bool>,
    #[sea_orm(column_name = "feedbackAt", column_type = "custom(\"timestamp(3)\")")]
    pub feedback_at: Option<DateTime>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "notificationMessageId", column_type = "Text")]
    pub notification_message_id: Option<String>,
    #[sea_orm(
        column_name = "notificationSentAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub notification_sent_at: Option<DateTime>,
    #[sea_orm(column_name = "notificationBatchId", column_type = "Text")]
    pub notification_batch_id: Option<String>,
    #[sea_orm(column_name = "driveConnectionId", column_type = "Text")]
    pub drive_connection_id: String,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::drive_connection::Entity",
        from = "Column::DriveConnectionId",
        to = "super::drive_connection::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentFiling_driveConnectionId_fkey"
    )]
    DriveConnection,
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentFiling_emailAccountId_fkey"
    )]
    EmailAccount,
}

impl Related<super::drive_connection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DriveConnection.def()
    }
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
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
