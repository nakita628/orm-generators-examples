defmodule Hoppscotch.TeamRequest do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Hoppscotch.PrismaDateTime, autogenerate: {Hoppscotch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          title: String.t(),
          request: map() | list() | String.t() | number() | boolean(),
          mock_examples: map() | list() | String.t() | number() | boolean() | nil,
          order_index: integer(),
          created_on: DateTime.t(),
          updated_at: DateTime.t(),
          collection: Hoppscotch.TeamCollection.t() | nil,
          team: Hoppscotch.Team.t() | nil
        }

  schema "TeamRequest" do
    field(:id, :string, primary_key: true)
    field(:title, :string)
    field(:request, Hoppscotch.PrismaJson)
    field(:mock_examples, Hoppscotch.PrismaJson, source: :mockExamples)
    field(:order_index, :integer, source: :orderIndex)
    field(:created_on, Hoppscotch.PrismaDateTime, autogenerate: true, source: :createdOn)
    field(:collection_id, :string, source: :collectionID)
    field(:team_id, :string, source: :teamID)
    belongs_to(:collection, Hoppscotch.TeamCollection, foreign_key: :collection_id, define_field: false, type: :string)
    belongs_to(:team, Hoppscotch.Team, foreign_key: :team_id, define_field: false, type: :string)
    timestamps(inserted_at: false, updated_at_source: :updatedOn)
  end
end