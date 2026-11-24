defmodule Postiz.MastraDatasetItems do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          dataset_id: String.t(),
          dataset_version: integer(),
          valid_to: integer() | nil,
          is_deleted: boolean(),
          input: map() | list() | String.t() | number() | boolean(),
          ground_truth: map() | list() | String.t() | number() | boolean() | nil,
          request_context: map() | list() | String.t() | number() | boolean() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          source: map() | list() | String.t() | number() | boolean() | nil,
          created_at: DateTime.t(),
          updated_at: DateTime.t(),
          created_at_z: DateTime.t() | nil,
          updated_at_z: DateTime.t() | nil,
          expected_trajectory: map() | list() | String.t() | number() | boolean() | nil,
          organization_id: String.t() | nil,
          project_id: String.t() | nil,
          tool_mocks: map() | list() | String.t() | number() | boolean() | nil,
          unmocked_tool_policy: String.t() | nil,
          scorer_ids: map() | list() | String.t() | number() | boolean() | nil,
          external_id: String.t() | nil
        }

  schema "mastra_dataset_items" do
    field(:id, :string, primary_key: true)
    field(:dataset_id, :string, source: :datasetId)
    field(:dataset_version, :integer, primary_key: true, source: :datasetVersion)
    field(:valid_to, :integer, source: :validTo)
    field(:is_deleted, :boolean, source: :isDeleted)
    field(:input, Postiz.PrismaJson)
    field(:ground_truth, Postiz.PrismaJson, source: :groundTruth)
    field(:request_context, Postiz.PrismaJson, source: :requestContext)
    field(:metadata, Postiz.PrismaJson)
    field(:source, Postiz.PrismaJson)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:updated_at, Postiz.PrismaDateTime, source: :updatedAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
    field(:updated_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :updatedAtZ)
    field(:expected_trajectory, Postiz.PrismaJson, source: :expectedTrajectory)
    field(:organization_id, :string, source: :organizationId)
    field(:project_id, :string, source: :projectId)
    field(:tool_mocks, Postiz.PrismaJson, source: :toolMocks)
    field(:unmocked_tool_policy, :string, source: :unmockedToolPolicy)
    field(:scorer_ids, Postiz.PrismaJson, source: :scorerIds)
    field(:external_id, :string, source: :externalId)
  end
end