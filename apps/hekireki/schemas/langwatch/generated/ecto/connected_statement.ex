defmodule LangWatch.ConnectedStatement do
  @moduledoc """
  The monthly usage statement one connected customer was sent (ADR-139,
  section 7). One row per account per month keeps the statement to one send,
  whatever runs the tick.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          account_id: String.t(),
          month: DateTime.t(),
          sent_at: DateTime.t()
        }

  schema "ConnectedStatement" do
    field(:id, :string, primary_key: true)
    # The `ConnectedBillingAccount` row. A plain column ref.
    field(:account_id, :string, source: :accountId)
    # First day of the month the statement covers, in UTC.
    field(:month, LangWatch.PrismaDateTime)
    field(:sent_at, LangWatch.PrismaDateTime, source: :sentAt)
  end
end