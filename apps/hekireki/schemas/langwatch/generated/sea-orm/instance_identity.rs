use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The identity this install presents to LangWatch (ADR-139, section 2).
///
/// Exactly one row, minted on first use and never rewritten. It survives
/// restarts, backups and hostname changes, and an install restored into
/// another database mints another one, which is what the registry refuses as
/// the wrong instance.
///
/// It replaced an identity built from the organization name and id, which
/// leaked the customer's name in cleartext, named one organization rather than
/// one install, and could not be joined to a license row without splitting a
/// string.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "InstanceIdentity")]
pub struct Model {
    /// Fixed, so this table holds one row. A second insert loses the race on the
    /// primary key and re-reads the winner's id.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// A random UUID. Carries nothing about the customer.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "instanceId", column_type = "Text")]
    pub instance_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    /// When the usage report last reached LangWatch, and why the last attempt
    /// did not. Both are what the checkup page reads to say whether reporting
    /// works, which nothing recorded before: the sender logged success as soon
    /// as `fetch` resolved and never looked at the response.
    #[sea_orm(column_name = "lastReportAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_report_at: Option<DateTime>,
    #[sea_orm(column_name = "lastReportError", column_type = "Text")]
    pub last_report_error: Option<String>,
    /// The optional category of the usage report, switched off from the checkup
    /// page. With it off the report still says which release runs and how many
    /// projects the install carries, and says nothing about what anyone does
    /// with it (ADR-139, section 10).
    #[sea_orm(column_name = "optionalMetricsOptOut", default_value = false)]
    pub optional_metrics_opt_out: bool,
    /// Hostname, which names the customer's own network and so has a switch of
    /// its own rather than travelling with its category.
    #[sea_orm(column_name = "hostnameOptOut", default_value = false)]
    pub hostname_opt_out: bool,
    /// Sending a failed upgrade step's diagnostics to LangWatch on its own,
    /// switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
    #[sea_orm(column_name = "upgradeDiagnosticsOptOut", default_value = false)]
    pub upgrade_diagnostics_opt_out: bool,
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
        Ok(self)
    }
}
