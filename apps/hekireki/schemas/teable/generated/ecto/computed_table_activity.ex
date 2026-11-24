defmodule Teable.ComputedTableActivity do
  @moduledoc """
  Projection of async computed activity per table.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          table_id: String.t(),
          base_id: String.t(),
          status: String.t(),
          calculating_field_count: integer(),
          queued_field_count: integer(),
          estimated_complexity: integer(),
          recent_completions: map() | list() | String.t() | number() | boolean(),
          generation: integer(),
          updated_at: DateTime.t()
        }

  schema "computed_table_activity" do
    field(:table_id, :string, primary_key: true)
    field(:base_id, :string)
    field(:status, :string)
    field(:calculating_field_count, :integer, default: 0)
    field(:queued_field_count, :integer, default: 0)
    field(:estimated_complexity, :integer, default: 0)
    field(:recent_completions, Teable.PrismaJson, default: [])
    field(:generation, :integer, default: 0)
    field(:updated_at, Teable.PrismaDateTime, autogenerate: true)
  end
end