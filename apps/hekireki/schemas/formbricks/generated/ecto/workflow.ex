defmodule Formbricks.Workflow do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          description: String.t() | nil,
          status: atom(),
          definition: map() | list() | String.t() | number() | boolean(),
          webhook_token_hash: String.t() | nil,
          webhook_token_encrypted: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil,
          creator: Formbricks.User.t() | nil
        }

  schema "Workflow" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:description, :string)
    field(:status, Ecto.Enum, values: [:draft, :enabled, :disabled, :archived], default: :draft)
    # [WorkflowDefinition]
    field(:definition, Formbricks.PrismaJson)
    # Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
    # the editor can show the URL again. Never part of the definition, its versions or its copies.
    field(:webhook_token_hash, :string, source: :webhookTokenHash)
    field(:webhook_token_encrypted, :string, source: :webhookTokenEncrypted)
    field(:workspace_id, :string, source: :workspaceId)
    field(:created_by, :string, source: :createdBy)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    belongs_to(:creator, Formbricks.User, foreign_key: :created_by, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end