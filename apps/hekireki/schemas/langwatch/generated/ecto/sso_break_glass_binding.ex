defmodule LangWatch.SsoBreakGlassBinding do
  @moduledoc """
  A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).

  One named person who may still sign in without the identity provider
  while their organization's sign-in belongs to one, and the date on which
  that stops. Activation asks whether such a person exists, because the
  failure it prevents is the only one that cannot be recovered from inside
  the product: a misconfigured connection goes live, every member is handed
  to an identity provider that will not authenticate them, and nobody is
  left who can turn it off.

  Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
  writes a NEW row naming the one it replaced, so the date a way in
  previously ended is still readable in the history - which is what makes a
  renewal auditable rather than a column somebody moved.

  No relation to Organization or User, matching the identity heads above:
  the binding is the record of who could get in, and a cascade would erase
  exactly the history a lockout post-mortem reads.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          user_id: String.t(),
          granted_by_user_id: String.t(),
          granted_at: DateTime.t(),
          expires_at: DateTime.t(),
          superseded_at: DateTime.t() | nil,
          renewed_from_id: String.t() | nil,
          warned_days: [integer()],
          inserted_at: DateTime.t()
        }

  schema "SsoBreakGlassBinding" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # Who holds the way in.
    field(:user_id, :string, source: :userId)
    # Who granted it. A way back in is never self-served.
    field(:granted_by_user_id, :string, source: :grantedByUserId)
    field(:granted_at, LangWatch.PrismaDateTime, source: :grantedAt)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    # Set when a renewal replaced this row; the row itself is never edited.
    field(:superseded_at, LangWatch.PrismaDateTime, source: :supersededAt)
    # The binding this one renewed, so the previous end date stays readable.
    field(:renewed_from_id, :string, source: :renewedFromId)
    # Which of the fourteen, seven and one day warnings have been sent, so a
    # sweep that runs twice in a day never says the same thing twice.
    field(:warned_days, {:array, :integer}, source: :warnedDays)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end