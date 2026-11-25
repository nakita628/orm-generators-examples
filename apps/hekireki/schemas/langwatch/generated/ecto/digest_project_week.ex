defmodule LangWatch.DigestProjectWeek do
  @moduledoc """
  A project's ISO week, folded from its owners' events. The id is the fold's aggregate key.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          project_id: String.t(),
          organization_id: String.t(),
          iso_week: String.t(),
          state: map() | list() | String.t() | number() | boolean(),
          occurred_at: DateTime.t(),
          last_event_id: String.t(),
          accepted_at: DateTime.t(),
          projection_version: String.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "DigestProjectWeek" do
    field(:id, :string, primary_key: true)
    field(:project_id, :string, source: :projectId)
    field(:organization_id, :string, source: :organizationId)
    field(:iso_week, :string, source: :isoWeek)
    # The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps.
    field(:state, LangWatch.PrismaJson)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:accepted_at, LangWatch.PrismaDateTime, source: :acceptedAt)
    field(:projection_version, :string, source: :projectionVersion)
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
  end
end