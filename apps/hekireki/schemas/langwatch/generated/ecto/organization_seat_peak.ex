defmodule LangWatch.OrganizationSeatPeak do
  @moduledoc """
  Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
  A read model, one row per organization: rebuildable by replaying the event log, so no relation.
  Epochs in milliseconds.
  Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          organization_id: String.t(),
          opened_at: float(),
          full_seats: integer(),
          lite_seats: integer(),
          developer_seats: integer(),
          latest_change_at: float(),
          mismatches: integer()
        }

  schema "OrganizationSeatPeak" do
    # The organization, which is also the aggregate id.
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    # When the opening snapshot was taken, epoch milliseconds.
    field(:opened_at, :float, source: :openedAt)
    field(:full_seats, :integer, source: :fullSeats)
    field(:lite_seats, :integer, source: :liteSeats)
    field(:developer_seats, :integer, source: :developerSeats)
    # The latest change time applied to any holder; invitations expire up to it.
    field(:latest_change_at, :float, source: :latestChangeAt)
    # Changes whose `from` disagreed with the seat the fold held.
    field(:mismatches, :integer)
  end
end