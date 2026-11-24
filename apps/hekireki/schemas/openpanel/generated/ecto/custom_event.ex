defmodule OpenPanel.CustomEvent do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: OpenPanel.PrismaDateTime, autogenerate: {OpenPanel.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: Ecto.UUID.t(),
          name: String.t(),
          items: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: OpenPanel.Project.t() | nil
        }

  schema "custom_events" do
    field(:id, Ecto.UUID, primary_key: true, read_after_writes: true)
    field(:name, :string)
    # [IPrismaCustomEventItems]
    field(:items, OpenPanel.PrismaJson)
    field(:project_id, :string, source: :projectId)
    belongs_to(:project, OpenPanel.Project, foreign_key: :project_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end