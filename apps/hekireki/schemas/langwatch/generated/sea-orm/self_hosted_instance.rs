use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One self-hosted install, as its daily usage report describes it (ADR-139,
/// section 10). The row is created by the first report that names an instance
/// id and is updated by every report after that, so the backoffice reads one
/// row per install rather than folding a history on every page load.
///
/// The report carries no credential, so it identifies an install and never a
/// customer. `organizationId` and `issuedLicenseId` are resolved from the
/// license bound to this instance, which the install proved by presenting its
/// token to license sync. A field in the posted body never reaches them.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SelfHostedInstance")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// The UUID the install minted for itself. The join key for everything else.
    #[sea_orm(unique)]
    #[sea_orm(column_name = "instanceId", column_type = "Text")]
    pub instance_id: String,
    #[sea_orm(
        column_name = "firstSeenAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub first_seen_at: DateTime,
    #[sea_orm(column_name = "lastSeenAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_seen_at: DateTime,
    /// What the install says about itself.
    #[sea_orm(column_type = "Text")]
    pub version: Option<String>,
    #[sea_orm(column_name = "installMethod", column_type = "Text")]
    pub install_method: Option<String>,
    #[sea_orm(column_name = "chartVersion", column_type = "Text")]
    pub chart_version: Option<String>,
    /// The host it answers on, when the operator left hostname reporting on.
    #[sea_orm(column_type = "Text")]
    pub hostname: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub environment: Option<String>,
    /// The day the install itself says it first came up, which predates the
    /// first report on an install that switched reporting on later.
    #[sea_orm(column_name = "installedAt", column_type = "custom(\"timestamp(3)\")")]
    pub installed_at: Option<DateTime>,
    #[sea_orm(column_name = "reportSchemaVersion")]
    pub report_schema_version: Option<i32>,
    /// The customer, and the license that named them. Plain column refs rather
    /// than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
    /// carries no foreign keys, and the receiver resolves both from the license
    /// registry rather than from the report.
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "issuedLicenseId", column_type = "Text")]
    pub issued_license_id: Option<String>,
    /// Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
    /// install that switched the optional category off.
    #[sea_orm(column_name = "userEmailDomains", column_type = "JsonBinary")]
    pub user_email_domains: Option<Json>,
    /// The same domains as the keys of `userEmailDomains`, flattened so an
    /// operator can search for one. A JSON object's keys are the domains
    /// themselves, and no query language addresses a key it does not already
    /// know the name of.
    #[sea_orm(
        column_name = "userDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub user_domains: Vec<String>,
    /// What the last report carried, so the list and the drawer read one row:
    /// the onboarding ladder, the usage counts and the operational block, under
    /// the same keys the dictionary declares.
    #[sea_orm(column_name = "latestReport", column_type = "JsonBinary")]
    pub latest_report: Option<Json>,
    /// What the last report said the operator had switched off, read from which
    /// categories arrived rather than from a field the install could assert.
    #[sea_orm(column_name = "optionalMetricsReported", default_value = true)]
    pub optional_metrics_reported: bool,
    #[sea_orm(column_name = "hostnameReported", default_value = true)]
    pub hostname_reported: bool,
    /// How many reports have landed, and how many fields the last one carried
    /// that this release has no name for. A count above zero says the senders
    /// are ahead of the receiver.
    #[sea_orm(column_name = "reportCount", default_value = 0)]
    pub report_count: i32,
    #[sea_orm(column_name = "lastUnknownFields", default_value = 0)]
    pub last_unknown_fields: i32,
    /// The lead signals already raised for this install. A signal on this list
    /// is never raised again, which is what keeps a large install from
    /// reporting that it is large every morning.
    #[sea_orm(
        column_name = "raisedSignals",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub raised_signals: Vec<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
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
        if insert && self.first_seen_at.is_not_set() {
            self.first_seen_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
