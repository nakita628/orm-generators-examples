defmodule LangWatch.LicenseRenewalClaim do
  @moduledoc """
  One renewed licence per paid Stripe invoice of a self-serve licence
  subscription (specs/licensing/self-serving-license-purchase.feature). A
  redelivered invoice event finds the row and mints nothing. A claim never
  delivered is taken over once `claimedAt` is older than the claim lease.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          stripe_invoice_id: String.t(),
          stripe_subscription_id: String.t(),
          claim_token: String.t(),
          claimed_at: DateTime.t(),
          issued_license: map() | nil,
          delivered_at: DateTime.t() | nil,
          license_id: String.t() | nil
        }

  schema "LicenseRenewalClaim" do
    field(:stripe_invoice_id, :string, primary_key: true, source: :stripeInvoiceId)
    field(:stripe_subscription_id, :string, source: :stripeSubscriptionId)
    # Fences completion and release to the delivery that holds the claim.
    field(:claim_token, :string, source: :claimToken)
    field(:claimed_at, LangWatch.PrismaDateTime, source: :claimedAt)
    # The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered.
    field(:issued_license, :map, source: :issuedLicense)
    field(:delivered_at, LangWatch.PrismaDateTime, source: :deliveredAt)
    # The `licenseId` inside the delivered licence.
    field(:license_id, :string, source: :licenseId)
  end
end