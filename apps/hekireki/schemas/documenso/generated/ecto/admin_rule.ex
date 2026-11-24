defmodule Documenso.AdminRule do
  @moduledoc """
  @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Documenso.PrismaDateTime, autogenerate: {Documenso.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          mode: atom(),
          trigger: String.t(),
          definition: map() | list() | String.t() | number() | boolean(),
          priority: integer(),
          stop_on_match: boolean(),
          exempt_organisation_ids: [String.t()],
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          matches: [Documenso.AdminRuleMatch.t()],
          admin_tags: [Documenso.AdminTag.t()]
        }

  schema "AdminRule" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:mode, Ecto.Enum, values: [:ENFORCE, :MONITOR, :DISABLED])
    # [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
    field(:trigger, :string)
    # [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema)
    field(:definition, Documenso.PrismaJson)
    field(:priority, :integer)
    field(:stop_on_match, :boolean, source: :stopOnMatch)
    # Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
    # organisation
    field(:exempt_organisation_ids, {:array, :string}, source: :exemptOrganisationIds)
    has_many(:matches, Documenso.AdminRuleMatch, foreign_key: :rule_id, references: :id)
    has_many(:admin_tags, Documenso.AdminTag, foreign_key: :admin_rule_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end