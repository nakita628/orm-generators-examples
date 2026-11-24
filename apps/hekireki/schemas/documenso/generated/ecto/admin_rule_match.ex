defmodule Documenso.AdminRuleMatch do
  @moduledoc """
  @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          mode: atom(),
          trigger: String.t(),
          organisation_id: String.t() | nil,
          envelope_id: String.t() | nil,
          event_snapshot: map() | list() | String.t() | number() | boolean(),
          actions_run: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          rule: Documenso.AdminRule.t() | nil,
          actor_user: Documenso.User.t() | nil
        }

  schema "AdminRuleMatch" do
    field(:id, :string, primary_key: true)
    field(:mode, Ecto.Enum, values: [:ENFORCE, :MONITOR, :DISABLED])
    # [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    field(:trigger, :string)
    field(:organisation_id, :string, source: :organisationId)
    field(:envelope_id, :string, source: :envelopeId)
    # [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema)
    field(:event_snapshot, Documenso.PrismaJson, source: :eventSnapshot)
    # [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema)
    field(:actions_run, Documenso.PrismaJson, source: :actionsRun)
    field(:rule_id, :string, source: :ruleId)
    field(:actor_user_id, :id, source: :actorUserId)
    belongs_to(:rule, Documenso.AdminRule, foreign_key: :rule_id, define_field: false, type: :string)
    belongs_to(:actor_user, Documenso.User, foreign_key: :actor_user_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end