defmodule Teable.FieldPeriodicRefreshState do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Teable.PrismaDateTime, autogenerate: {Teable.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          field_id: String.t(),
          table_id: String.t(),
          base_id: String.t(),
          enabled: boolean(),
          time_zone: String.t(),
          contains_now_precision: boolean(),
          last_refreshed_at: DateTime.t() | nil,
          last_refreshed_local_date: String.t() | nil,
          last_range_hash: String.t() | nil,
          last_status: String.t(),
          last_error: String.t() | nil,
          pending_task_id: String.t() | nil,
          pending_local_date: String.t() | nil,
          pending_range_hash: String.t() | nil,
          pending_started_at: DateTime.t() | nil,
          created_time: DateTime.t(),
          updated_at: DateTime.t() | nil,
          field: Teable.Field.t() | nil
        }

  schema "field_periodic_refresh_state" do
    field(:table_id, :string)
    field(:base_id, :string)
    field(:enabled, :boolean, default: true)
    field(:time_zone, :string, default: "UTC")
    field(:contains_now_precision, :boolean, default: false)
    field(:last_refreshed_at, Teable.PrismaDateTime)
    field(:last_refreshed_local_date, :string)
    field(:last_range_hash, :string)
    field(:last_status, :string, default: "idle")
    field(:last_error, :string)
    field(:pending_task_id, :string)
    field(:pending_local_date, :string)
    field(:pending_range_hash, :string)
    field(:pending_started_at, Teable.PrismaDateTime)
    field(:created_time, Teable.PrismaDateTime, autogenerate: true)
    field(:field_id, :string, primary_key: true)
    belongs_to(:field, Teable.Field, foreign_key: :field_id, define_field: false, type: :string)
    timestamps(inserted_at: false, updated_at_source: :last_modified_time)
  end
end