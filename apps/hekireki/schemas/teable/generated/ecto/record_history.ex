defmodule Teable.RecordHistory do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          table_id: String.t(),
          record_id: String.t(),
          field_id: String.t(),
          before: String.t(),
          after: String.t(),
          created_time: DateTime.t(),
          created_by: String.t(),
          actor_type: String.t() | nil,
          actor_id: String.t() | nil
        }

  schema "record_history" do
    field(:id, :string, primary_key: true)
    field(:table_id, :string)
    field(:record_id, :string)
    field(:field_id, :string)
    field(:before, :string)
    field(:after, :string)
    field(:created_time, Teable.PrismaDateTime, autogenerate: true)
    field(:created_by, :string)
    field(:actor_type, :string)
    field(:actor_id, :string)
  end
end