defmodule LangWatch.RetentionApplyStatus do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          scope_type: atom(),
          scope_id: String.t(),
          changed_at: float() | nil,
          applied_at: float() | nil,
          raised_at: float() | nil,
          shortened_at: float() | nil,
          created_at: float(),
          updated_at: float(),
          occurred_at: float(),
          accepted_at: float(),
          last_event_id: String.t(),
          projection_version: String.t()
        }

  schema "RetentionApplyStatus" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    field(:scope_type, Ecto.Enum, values: [:ORGANIZATION, :TEAM, :PROJECT], source: :scopeType)
    field(:scope_id, :string, source: :scopeId)
    field(:changed_at, :float, source: :changedAt)
    field(:applied_at, :float, source: :appliedAt)
    field(:raised_at, :float, source: :raisedAt)
    field(:shortened_at, :float, source: :shortenedAt)
    field(:created_at, :float, source: :createdAt)
    field(:updated_at, :float, source: :updatedAt)
    field(:occurred_at, :float, source: :occurredAt)
    field(:accepted_at, :float, source: :acceptedAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:projection_version, :string, source: :projectionVersion)
  end
end