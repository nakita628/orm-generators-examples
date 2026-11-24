defmodule Teable.DomainEventOutbox do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Teable.PrismaDateTime, autogenerate: {Teable.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          base_id: String.t(),
          table_id: String.t() | nil,
          message_name: String.t(),
          schema_version: integer(),
          aggregate_id: String.t() | nil,
          payload: map() | list() | String.t() | number() | boolean(),
          payload_bytes: integer(),
          catalog_generation: integer(),
          required_consumers: map() | list() | String.t() | number() | boolean(),
          binding_id: String.t() | nil,
          storage_epoch: integer() | nil,
          unpublished: boolean(),
          settled: String.t() | nil,
          settled_at: DateTime.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "domain_event_outbox" do
    field(:id, :string, primary_key: true)
    field(:base_id, :string)
    field(:table_id, :string)
    field(:message_name, :string)
    field(:schema_version, :integer)
    field(:aggregate_id, :string)
    field(:payload, Teable.PrismaJson)
    field(:payload_bytes, :integer)
    field(:catalog_generation, :integer)
    field(:required_consumers, Teable.PrismaJson)
    field(:binding_id, :string)
    field(:storage_epoch, :integer)
    field(:unpublished, :boolean, default: true)
    field(:settled, :string)
    field(:settled_at, Teable.PrismaDateTime)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end