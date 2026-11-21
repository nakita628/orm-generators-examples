defmodule LangWatch.IssuedLicense do
  @moduledoc """
  The registry of licenses LangWatch has issued (ADR-139). It lives on
  LangWatch Cloud; a self-hosted install has the table and no rows.

  The row never holds the token an install presents. `tokenHash` is the
  SHA-256 of that token, so read access to this table is not credential
  access. Status is derived: revoked, superseded, past `expiresAt`, else active.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          license_id: String.t(),
          token_hash: String.t(),
          organization_name: String.t(),
          email: String.t(),
          plan_type: String.t(),
          max_members: integer(),
          max_members_lite: integer() | nil,
          lite_seats_unlimited: boolean(),
          max_members_developer: integer() | nil,
          issued_at: DateTime.t(),
          expires_at: DateTime.t(),
          source: atom(),
          issued_by_id: String.t() | nil,
          revoked_at: DateTime.t() | nil,
          revoked_by_id: String.t() | nil,
          revoked_reason: String.t() | nil,
          superseded_at: DateTime.t() | nil,
          replaces_id: String.t() | nil,
          pending_delivery_license: String.t() | nil,
          services: [String.t()],
          seat_rate_cents: integer() | nil,
          seat_currency: atom() | nil,
          commit_usd_cents: integer(),
          overage_enabled: boolean(),
          overage_max_usd_cents: integer() | nil,
          instance_id: String.t() | nil,
          instance_bound_at: DateTime.t() | nil,
          last_sync_at: DateTime.t() | nil,
          last_sync_version: String.t() | nil,
          reported_members: integer() | nil,
          reported_members_lite: integer() | nil,
          virtual_key_id: String.t() | nil,
          seats_raised_from: integer() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: LangWatch.Organization.t() | nil
        }

  schema "IssuedLicense" do
    field(:id, :string, primary_key: true)
    # The `licenseId` inside the signed license (`lic-<uuid>`).
    field(:license_id, :string, source: :licenseId)
    # SHA-256 of the `lwl_` token. Looked up on every credential resolution.
    field(:token_hash, :string, source: :tokenHash)
    # As signed into the license.
    field(:organization_name, :string, source: :organizationName)
    field(:email, :string)
    field(:plan_type, :string, source: :planType)
    field(:max_members, :integer, source: :maxMembers)
    # The lite seats sold: null when not stated or when unlimited. Never a stand-in number.
    field(:max_members_lite, :integer, source: :maxMembersLite)
    # True when the license sells unlimited lite seats; maxMembersLite is then null.
    field(:lite_seats_unlimited, :boolean, default: false, source: :liteSeatsUnlimited)
    # The Developer seats sold: null when not stated, which sells none.
    field(:max_members_developer, :integer, source: :maxMembersDeveloper)
    field(:issued_at, LangWatch.PrismaDateTime, source: :issuedAt)
    field(:expires_at, LangWatch.PrismaDateTime, source: :expiresAt)
    field(:source, Ecto.Enum, values: [:BACKOFFICE, :PURCHASE, :SCRIPT, :LEGACY_IMPORT])
    # The operator who issued or registered it. Null for the purchase flow and the script.
    field(:issued_by_id, :string, source: :issuedById)
    field(:revoked_at, LangWatch.PrismaDateTime, source: :revokedAt)
    field(:revoked_by_id, :string, source: :revokedById)
    field(:revoked_reason, :string, source: :revokedReason)
    # Set when the install first presents the license that replaced this one.
    field(:superseded_at, LangWatch.PrismaDateTime, source: :supersededAt)
    # The id of the IssuedLicense this one replaces, when it was reissued. A
    # plain column ref (relationMode="prisma" cannot declare a self-relation on
    # Postgres); the registry service validates it. Unique, so a license is
    # replaced by one license at most.
    field(:replaces_id, :string, source: :replacesId)
    # A reissued license waiting to be delivered over sync, encrypted. Erased
    # when the install first presents it. The only license text the registry holds.
    field(:pending_delivery_license, :string, source: :pendingDeliveryLicense)
    # Hosted services this license may call, such as `instant_evals`.
    field(:services, {:array, :string})
    # The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats.
    field(:seat_rate_cents, :integer, source: :seatRateCents)
    field(:seat_currency, Ecto.Enum, values: [:USD, :EUR], source: :seatCurrency)
    # The prepaid usage commit for the term, in USD cents.
    field(:commit_usd_cents, :integer, default: 0, source: :commitUsdCents)
    # Whether usage may continue past the commit, up to `overageMaxUsdCents`.
    field(:overage_enabled, :boolean, default: false, source: :overageEnabled)
    field(:overage_max_usd_cents, :integer, source: :overageMaxUsdCents)
    # The install this license is bound to. Set on first use, cleared by an operator.
    field(:instance_id, :string, source: :instanceId)
    field(:instance_bound_at, LangWatch.PrismaDateTime, source: :instanceBoundAt)
    # The last license sync this install reported: when it arrived, the
    # LangWatch version that sent it, and the seats in use at that moment.
    # Null until the install syncs for the first time.
    field(:last_sync_at, LangWatch.PrismaDateTime, source: :lastSyncAt)
    field(:last_sync_version, :string, source: :lastSyncVersion)
    field(:reported_members, :integer, source: :reportedMembers)
    field(:reported_members_lite, :integer, source: :reportedMembersLite)
    # The managed gateway key (purpose CONNECT) that hosted-service traffic of
    # this license runs under. One per license, so spend rows name the install.
    # Created when the license first resolves. A plain column, not a relation.
    field(:virtual_key_id, :string, source: :virtualKeyId)
    # The seats of the license this one replaced, set only when a seat change
    # raised a linked license. Billing reads it through licensing to invoice
    # the added seats, so the fact commits with the license or not at all.
    field(:seats_raised_from, :integer, source: :seatsRaisedFrom)
    # The customer organization on LangWatch Cloud. Null for a license that is
    # recorded but not linked to a customer yet, such as a self-serve purchase:
    # it resolves to nothing until an operator links it.
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end