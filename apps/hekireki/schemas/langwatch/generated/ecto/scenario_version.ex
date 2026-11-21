defmodule LangWatch.ScenarioVersion do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          version: integer(),
          author_id: String.t() | nil,
          author_label: String.t(),
          change_description: String.t() | nil,
          snapshot: map(),
          schema_version: integer(),
          inserted_at: DateTime.t(),
          scenario: LangWatch.Scenario.t() | nil
        }

  schema "ScenarioVersion" do
    field(:id, :string, primary_key: true)
    # The project of the parent case, carried here so every read filters by
    # tenant without a join. The service is what keeps the pair honest: no
    # composite foreign key can, under relationMode="prisma".
    field(:project_id, :string, source: :projectId)
    field(:version, :integer)
    field(:author_id, :string, source: :authorId)
    # Who wrote this version: "user", "api", "cli" or "langy".
    field(:author_label, :string, source: :authorLabel)
    field(:change_description, :string, source: :changeDescription)
    # Envelope { schemaVersion, fields, changedFields } holding the editable
    # field set as saved and the names of the fields this save changed.
    field(:snapshot, :map)
    # Schema version of `snapshot`, so a later shape change can be read forward.
    field(:schema_version, :integer, default: 1, source: :schemaVersion)
    field(:scenario_id, :string, source: :scenarioId)
    # Deleting the case takes its history with it. Under relationMode="prisma"
    # the cascade is the client's, so a raw SQL delete still has to clear these
    # rows itself.
    belongs_to(:scenario, LangWatch.Scenario, foreign_key: :scenario_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end