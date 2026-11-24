defmodule Formbricks.EmbeddedFieldLabel do
  @moduledoc """
  A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
  place of the key in chart legends, headers and the chart builder's field pickers. One per key and
  workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
  are never renamed here: they keep their translated names.

  @property key - The embedded-data key exactly as collected, case included
  @property label - The name shown for the key across the workspace
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          key: String.t(),
          label: String.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          workspace: Formbricks.Workspace.t() | nil
        }

  schema "EmbeddedFieldLabel" do
    field(:id, :string, primary_key: true)
    field(:key, :string)
    field(:label, :string)
    field(:workspace_id, :string, source: :workspaceId)
    belongs_to(:workspace, Formbricks.Workspace, foreign_key: :workspace_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end