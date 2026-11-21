defmodule Trigger.IntegrationDeployment do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Trigger.PrismaDateTime, autogenerate: {Trigger.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          integration_name: String.t(),
          integration_deployment_id: String.t(),
          commit_sha: String.t(),
          status: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          worker_deployment: Trigger.WorkerDeployment.t() | nil
        }

  schema "IntegrationDeployment" do
    field(:id, :string, primary_key: true)
    # For example Vercel
    field(:integration_name, :string, source: :integrationName)
    # External ID
    field(:integration_deployment_id, :string, source: :integrationDeploymentId)
    field(:commit_sha, :string, source: :commitSHA)
    # External deployment status
    field(:status, :string)
    field(:deployment_id, :string, source: :deploymentId)
    belongs_to(:worker_deployment, Trigger.WorkerDeployment, foreign_key: :deployment_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end