use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ExperimentVersion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "experimentId", column_type = "Text")]
    pub experiment_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// The number the version is addressed and displayed by. Numbered rows hold
    /// a gap-free sequence. The rolling autosave row rides on the counter
    /// instead, which is what keeps it clear of that sequence.
    pub version: i32,
    /// `Experiment.workbenchVersion` at the moment this row was last written.
    /// The rolling row is rewritten in place, so this is the only column that
    /// says how recent its content is, and the version list sorts and pages by
    /// it. It also joins a row to the live state: the row whose `counterVersion`
    /// equals the experiment's counter holds what the workbench shows now.
    #[sea_orm(column_name = "counterVersion")]
    pub counter_version: i32,
    #[sea_orm(column_name = "autoSaved", default_value = false)]
    pub auto_saved: bool,
    #[sea_orm(column_name = "commitMessage", column_type = "Text")]
    pub commit_message: Option<String>,
    #[sea_orm(column_name = "authorId", column_type = "Text")]
    pub author_id: Option<String>,
    /// Who wrote this version: "user", "langy" or "api".
    #[sea_orm(column_name = "authorLabel", column_type = "Text")]
    pub author_label: String,
    /// The run that wrote this version, when a run wrote it.
    ///
    /// A run writes its cells into the workbench state, which advances the
    /// counter. Without this, the page that started that run reads the bump as
    /// somebody else's write, refuses to save, and asks the reader to reload
    /// over their own unsaved edits. Named here, the page adopts the version and
    /// carries on.
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub state: Json,
    /// Schema version of `state`, so a later shape change can be read forward.
    #[sea_orm(
        column_name = "schemaVersion",
        column_type = "Text",
        default_value = "1"
    )]
    pub schema_version: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
