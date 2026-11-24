use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The one open stop an agent may have: the api takes this row before it sends a stop's request,
/// so two requests at once cannot both open one. A later request takes it over once the stop it
/// names has closed.
/// `sentAt` is set once the request command was sent: from then on the claim passes on only after
/// its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
/// `decisionTakenAt` is set while that decision is in flight and cleared once its command was
/// sent; a decision left in flight is taken over once abandoned.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentStopClaim")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "agentId")]
    pub agent_id: String,
    #[sea_orm(column_name = "stopId")]
    pub stop_id: String,
    #[sea_orm(column_name = "claimedAt")]
    pub claimed_at: DateTime,
    #[sea_orm(column_name = "sentAt")]
    pub sent_at: Option<DateTime>,
    pub decision: Option<String>,
    #[sea_orm(column_name = "decisionTakenAt")]
    pub decision_taken_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
