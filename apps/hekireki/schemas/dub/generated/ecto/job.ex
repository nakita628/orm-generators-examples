defmodule Dub.Job do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Dub.PrismaDateTime, autogenerate: {Dub.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          payload: map() | list() | String.t() | number() | boolean(),
          options: map() | list() | String.t() | number() | boolean() | nil,
          scheduled_at: DateTime.t(),
          attempts: integer(),
          last_error: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "Job" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:payload, Dub.PrismaJson)
    field(:options, Dub.PrismaJson)
    field(:scheduled_at, Dub.PrismaDateTime, autogenerate: true, source: :scheduledAt)
    field(:attempts, :integer, default: 0)
    field(:last_error, :string, source: :lastError)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end