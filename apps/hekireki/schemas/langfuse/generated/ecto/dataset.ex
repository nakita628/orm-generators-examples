defmodule Langfuse.Dataset do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          project_id: String.t(),
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          metadata: map() | list() | String.t() | number() | boolean() | nil,
          remote_experiment_url: String.t() | nil,
          remote_experiment_payload: map() | list() | String.t() | number() | boolean() | nil,
          remote_experiment_enabled: boolean(),
          remote_experiment_secret_key: String.t() | nil,
          remote_experiment_display_secret_key: String.t() | nil,
          remote_experiment_request_headers: map() | list() | String.t() | number() | boolean() | nil,
          remote_experiment_display_headers: map() | list() | String.t() | number() | boolean() | nil,
          input_schema: map() | list() | String.t() | number() | boolean() | nil,
          expected_output_schema: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          project: Langfuse.Project.t() | nil
        }

  schema "datasets" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:metadata, Langfuse.PrismaJson)
    field(:remote_experiment_url, :string)
    field(:remote_experiment_payload, Langfuse.PrismaJson)
    field(:remote_experiment_enabled, :boolean, default: true)
    field(:remote_experiment_secret_key, :string)
    field(:remote_experiment_display_secret_key, :string)
    field(:remote_experiment_request_headers, Langfuse.PrismaJson)
    field(:remote_experiment_display_headers, Langfuse.PrismaJson)
    field(:input_schema, Langfuse.PrismaJson)
    field(:expected_output_schema, Langfuse.PrismaJson)
    field(:project_id, :string, primary_key: true)
    belongs_to(:project, Langfuse.Project, foreign_key: :project_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end