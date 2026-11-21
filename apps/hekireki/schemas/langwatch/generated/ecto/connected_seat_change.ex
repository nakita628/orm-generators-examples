defmodule LangWatch.ConnectedSeatChange do
  @moduledoc """
  What a mid-term seat change of one connected license owes (ADR-141,
  section 7): the seats added on the reissued license, invoiced prorated to
  the end of the term.

  The row is written before the payment provider is called and completed
  after it, so a call that failed halfway is retried by the next seat
  invoicing pass without inventing a second invoice. One row per reissued license.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          license_id: String.t(),
          account_id: String.t() | nil,
          organization_id: String.t() | nil,
          changed_at: DateTime.t(),
          added_seats: integer(),
          unit_amount_cents: integer(),
          amount_cents: integer(),
          currency: atom() | nil,
          stripe_invoice_id: String.t() | nil,
          state: atom(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "ConnectedSeatChange" do
    field(:id, :string, primary_key: true)
    # The reissued `IssuedLicense` row that carries the new seat count. A
    # plain column ref, like `replacesId`: relationMode="prisma" carries no
    # foreign keys.
    field(:license_id, :string, source: :licenseId)
    # The `ConnectedBillingAccount` row the invoice is raised against. Null
    # when the customer had no billing account (`not_onboarded`).
    field(:account_id, :string, source: :accountId)
    # The customer organization the license belongs to.
    field(:organization_id, :string, source: :organizationId)
    # When the seats changed, which is the day the proration counts from.
    field(:changed_at, LangWatch.PrismaDateTime, source: :changedAt)
    # Seats on the new license above the seats on the one it replaced.
    field(:added_seats, :integer, source: :addedSeats)
    # What one added seat costs for the rest of the term, in cents of `currency`.
    field(:unit_amount_cents, :integer, source: :unitAmountCents)
    # `unitAmountCents` times `addedSeats`. Zero when nothing is owed.
    field(:amount_cents, :integer, source: :amountCents)
    # Null when no billing account named one (`not_onboarded`).
    field(:currency, Ecto.Enum, values: [:USD, :EUR])
    # The one-off invoice the added seats were billed on.
    field(:stripe_invoice_id, :string, source: :stripeInvoiceId)
    field(:state, Ecto.Enum, values: [:intent, :invoiced, :nothing_to_invoice, :not_onboarded])
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end