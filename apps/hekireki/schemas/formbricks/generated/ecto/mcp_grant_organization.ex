defmodule Formbricks.McpGrantOrganization do
  @moduledoc """
  An organization a grant with `organizations` access reaches, workspaces created there later included.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          grant_id: String.t(),
          organization_id: String.t(),
          grant: Formbricks.McpGrant.t() | nil,
          organization: Formbricks.Organization.t() | nil
        }

  schema "McpGrantOrganization" do
    field(:grant_id, :string, primary_key: true, source: :grantId)
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    belongs_to(:grant, Formbricks.McpGrant, foreign_key: :grant_id, define_field: false, type: :string)
    belongs_to(:organization, Formbricks.Organization, foreign_key: :organization_id, define_field: false, type: :string)
  end
end