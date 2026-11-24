defmodule Plunk.Domain do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Plunk.PrismaDateTime, autogenerate: {Plunk.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          domain: String.t(),
          verified: boolean(),
          dkim_tokens: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Plunk.Project.t() | nil
        }

  schema "domains" do
    field(:domain, :string)
    field(:verified, :boolean, default: false)
    field(:dkim_tokens, Plunk.PrismaJson, source: :dkimTokens)
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, Plunk.Project, foreign_key: :project_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end