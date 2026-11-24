defmodule LangWatch.GatewayCacheRule do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          priority: integer(),
          enabled: boolean(),
          matchers: map() | list() | String.t() | number() | boolean(),
          action: map() | list() | String.t() | number() | boolean(),
          mode_enum: atom(),
          archived_at: DateTime.t() | nil,
          created_by_id: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "GatewayCacheRule" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:name, :string)
    field(:description, :string)
    field(:priority, :integer, default: 100)
    field(:enabled, :boolean, default: true)
    field(:matchers, LangWatch.PrismaJson)
    field(:action, LangWatch.PrismaJson)
    field(:mode_enum, Ecto.Enum, values: [:RESPECT, :FORCE, :DISABLE], source: :modeEnum)
    field(:archived_at, LangWatch.PrismaDateTime, source: :archivedAt)
    field(:created_by_id, :string, source: :createdById)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end