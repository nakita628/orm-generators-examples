use super::vote_type::VoteType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "votes")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub participant_id: String,
    #[sea_orm(column_type = "Text")]
    pub option_id: String,
    #[sea_orm(column_type = "Text")]
    pub poll_id: String,
    #[sea_orm(default_value = "yes")]
    pub r#type: VoteType,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::participant::Entity",
        from = "Column::ParticipantId",
        to = "super::participant::Column::Id",
        on_delete = "Cascade",
        fk_name = "votes_participant_id_fkey"
    )]
    Participant,
    #[sea_orm(
        belongs_to = "super::option::Entity",
        from = "Column::OptionId",
        to = "super::option::Column::Id",
        on_delete = "Cascade",
        fk_name = "votes_option_id_fkey"
    )]
    Option,
    #[sea_orm(
        belongs_to = "super::poll::Entity",
        from = "Column::PollId",
        to = "super::poll::Column::Id",
        on_delete = "Cascade",
        fk_name = "votes_poll_id_fkey"
    )]
    Poll,
}

impl Related<super::participant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Participant.def()
    }
}

impl Related<super::option::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Option.def()
    }
}

impl Related<super::poll::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Poll.def()
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
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
