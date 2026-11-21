defmodule LangWatch.ActivationCode do
  @moduledoc """
  A short code a fresh install pastes instead of a license blob (ADR-139,
  section 5). Redeeming it mints the license the code describes and returns
  it signed, which is what identifies a new customer without anybody emailing
  a license key around.

  The row holds a hash of the code, never the code. Every lookup hashes what
  the caller presented first, so the stored value is a step away from anything
  that could be presented: read access to this table is not credential access.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          code_hash: String.t(),
          code_hint: String.t(),
          organization_id: String.t(),
          organization_name: String.t(),
          email: String.t(),
          plan_type: String.t(),
          max_members: integer(),
          max_members_lite: integer() | nil,
          lite_seats_unlimited: boolean(),
          max_members_developer: integer() | nil,
          license_term_days: integer(),
          services: [String.t()],
          expires_at: DateTime.t(),
          reusable: boolean(),
          redeemed_at: DateTime.t() | nil,
          redeemed_by_instance_id: String.t() | nil,
          issued_license_id: String.t() | nil,
          redemption_count: integer(),
          revoked_at: DateTime.t() | nil,
          revoked_by_id: String.t() | nil,
          created_by_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "ActivationCode" do
    field(:id, :string, primary_key: true)
    # SHA-256 of the normalised code. Looked up on every redemption.
    field(:code_hash, :string, source: :codeHash)
    # The last four characters, so an operator tells two codes apart in a list
    # without the list holding either one.
    field(:code_hint, :string, source: :codeHint)
    # The customer this code belongs to, and the terms it mints.
    field(:organization_id, :string, source: :organizationId)
    field(:organization_name, :string, source: :organizationName)
    field(:email, :string)
    field(:plan_type, :string, source: :planType)
    field(:max_members, :integer, source: :maxMembers)
    # The lite seats the minted license carries: null when unlimited. Never a stand-in number.
    field(:max_members_lite, :integer, default: 0, source: :maxMembersLite)
    # True when the minted license carries unlimited lite seats; maxMembersLite is then null.
    field(:lite_seats_unlimited, :boolean, default: false, source: :liteSeatsUnlimited)
    # The Developer seats the minted license carries: null when not stated, which sells none.
    field(:max_members_developer, :integer, source: :maxMembersDeveloper)
    # How long the license this code mints runs for.
    field(:license_term_days, :integer, source: :licenseTermDays)
    # Hosted services the minted license may call.
    field(:services, {:array, :string})
    # When the code stops working, which is not the term of what it mints.
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    # A code a customer may redeem on more than one install, for a rollout.
    field(:reusable, :boolean, default: false)
    # Set by the redemption that won, which is what makes a single-use code
    # single use: the write itself carries the condition.
    field(:redeemed_at, LangWatch.PrismaDateTime, source: :redeemedAt)
    field(:redeemed_by_instance_id, :string, source: :redeemedByInstanceId)
    # The license the winning redemption minted.
    field(:issued_license_id, :string, source: :issuedLicenseId)
    # How many installs have redeemed it, which only moves on a reusable code.
    field(:redemption_count, :integer, default: 0, source: :redemptionCount)
    field(:revoked_at, LangWatch.PrismaDateTime, source: :revokedAt)
    field(:revoked_by_id, :string, source: :revokedById)
    field(:created_by_id, :string, source: :createdById)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end