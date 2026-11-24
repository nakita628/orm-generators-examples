defmodule Postiz.MastraTraces do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          parent_span_id: String.t() | nil,
          name: String.t(),
          trace_id: String.t(),
          scope: String.t(),
          kind: integer(),
          attributes: map() | list() | String.t() | number() | boolean() | nil,
          status: map() | list() | String.t() | number() | boolean() | nil,
          events: map() | list() | String.t() | number() | boolean() | nil,
          links: map() | list() | String.t() | number() | boolean() | nil,
          other: String.t() | nil,
          start_time: integer(),
          end_time: integer(),
          created_at: DateTime.t(),
          created_at_z: DateTime.t() | nil
        }

  schema "mastra_traces" do
    field(:id, :string, primary_key: true)
    field(:parent_span_id, :string, source: :parentSpanId)
    field(:name, :string)
    field(:trace_id, :string, source: :traceId)
    field(:scope, :string)
    field(:kind, :integer)
    field(:attributes, Postiz.PrismaJson)
    field(:status, Postiz.PrismaJson)
    field(:events, Postiz.PrismaJson)
    field(:links, Postiz.PrismaJson)
    field(:other, :string)
    field(:start_time, :integer, source: :startTime)
    field(:end_time, :integer, source: :endTime)
    field(:created_at, Postiz.PrismaDateTime, source: :createdAt)
    field(:created_at_z, Postiz.PrismaDateTime, autogenerate: true, source: :createdAtZ)
  end
end