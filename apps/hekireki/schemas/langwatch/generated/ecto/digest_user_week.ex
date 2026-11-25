defmodule LangWatch.DigestUserWeek do
  @moduledoc """
  A person's ISO week of coding-agent activity in one organization, folded from the agents' facts.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t(),
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

  schema "DigestUserWeek" do
    field(:id, :string, primary_key: true)
    field(:user_id, :string, source: :userId)
    field(:organization_id, :string, source: :organizationId)
    field(:iso_week, :string, source: :isoWeek)
    # Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins.
    field(:state, LangWatch.PrismaJson)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:accepted_at, LangWatch.PrismaDateTime, source: :acceptedAt)
    field(:projection_version, :string, source: :projectionVersion)
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
  end
end