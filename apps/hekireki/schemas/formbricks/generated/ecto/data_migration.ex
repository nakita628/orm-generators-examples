defmodule Formbricks.DataMigration do
  @moduledoc """
  Tracks database schema migrations.
  Used to manage and track the state of data structure changes.

  @property id - Unique identifier for the migration
  @property name - Name of the migration
  @property status - Current state of the migration (pending, applied, failed)
  @property startedAt - When the migration began
  @property finishedAt - When the migration completed
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          started_at: DateTime.t(),
          finished_at: DateTime.t() | nil,
          name: String.t(),
          status: atom()
        }

  schema "DataMigration" do
    field(:id, :string, primary_key: true)
    field(:started_at, Formbricks.PrismaDateTime, autogenerate: true)
    field(:finished_at, Formbricks.PrismaDateTime)
    field(:name, :string)
    field(:status, Ecto.Enum, values: [:pending, :applied, :failed])
  end
end