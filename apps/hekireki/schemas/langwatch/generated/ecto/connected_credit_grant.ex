defmodule LangWatch.ConnectedCreditGrant do
  @moduledoc """
  A paid credit at the payment provider, which only metered usage draws down.
  Kept here so a grant is created once per reason and the backoffice can show
  what the customer holds without asking the provider.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          account_id: String.t(),
          stripe_credit_grant_id: String.t(),
          amount_usd_cents: integer(),
          kind: atom(),
          term_ends_at: DateTime.t(),
          expires_at: DateTime.t(),
          inserted_at: DateTime.t()
        }

  schema "ConnectedCreditGrant" do
    field(:id, :string, primary_key: true)
    # The `ConnectedBillingAccount` row. A plain column ref, like
    # `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
    field(:account_id, :string, source: :accountId)
    # The grant at the payment provider. Unique, so a retry records one row.
    field(:stripe_credit_grant_id, :string, source: :stripeCreditGrantId)
    field(:amount_usd_cents, :integer, source: :amountUsdCents)
    field(:kind, Ecto.Enum, values: [:commit, :added, :renewal])
    # The term this credit was granted for, and when the provider lets it lapse:
    # after the last invoice of that term has been finalized, never on the day
    # the term ends.
    field(:term_ends_at, LangWatch.PrismaDateTime, source: :termEndsAt)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end