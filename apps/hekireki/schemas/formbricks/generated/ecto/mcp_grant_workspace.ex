defmodule Formbricks.McpGrantWorkspace do
  @moduledoc """
  A workspace a grant with `workspaces` access reaches.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          grant_id: String.t(),
          workspace_id: String.t(),
          grant: Formbricks.McpGrant.t() | nil,
          workspace: Formbricks.Workspace.t() | nil
        }

  schema "McpGrantWorkspace" do
    field(:grant_id, :string, primary_key: true, source: :grantId)
    field(:workspace_id, :string, primary_key: true, source: :workspaceId)
    belongs_to(:grant, Formbricks.McpGrant, foreign_key: :grant_id, define_field: false, type: :string)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
  end
end