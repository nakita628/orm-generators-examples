defmodule Formbricks.Integration do
  @moduledoc """
  Represents third-party service integrations.
  Enables data flow between Formbricks and external services.

  @property id - Unique identifier for the integration
  @property type - The service being integrated (Google Sheets, Notion, etc.)
  @property workspace - The workspace this integration belongs to
  @property config - Service-specific configuration details
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          type: atom(),
          config: map(),
          workspace: Formbricks.Workspace.t() | nil
        }

  schema "Integration" do
    field(:id, :string, primary_key: true)
    field(:type, Ecto.Enum, values: [:googleSheets, :notion, :airtable, :slack])
    # [IntegrationConfig]
    field(:config, :map)
    field(:workspace_id, :string, source: :workspaceId)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
  end
end