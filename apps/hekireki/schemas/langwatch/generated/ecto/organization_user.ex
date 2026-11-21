defmodule LangWatch.OrganizationUser do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          organization_id: String.t(),
          role: atom(),
          department_id: String.t() | nil,
          disabled_at: DateTime.t() | nil,
          pending_sso_grant_id: String.t() | nil,
          membership_stamp: String.t(),
          full_seat_since: DateTime.t() | nil,
          seat_pending_since: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil,
          organization: LangWatch.Organization.t() | nil
        }

  schema "OrganizationUser" do
    field(:role, Ecto.Enum, values: [:ADMIN, :MEMBER, :EXTERNAL, :DEVELOPER])
    # Department assignment for spend attribution (the people lens).
    # Single-valued, org-scoped, accounting-only (never gates access).
    # Plain column ref (relationMode="prisma"); the service validates the
    # id belongs to this org. See departments.feature.
    field(:department_id, :string, source: :departmentId)
    # Set when an admin disables this membership to stay within the licensed
    # seat count. The person loses access to this organization only, keeps
    # their role, department and everything they did, and frees their seat.
    # Reversible: re-enabling restores access if a seat is available.
    # Seat counting excludes disabled memberships. See
    # seat-reconciliation.feature.
    field(:disabled_at, LangWatch.PrismaDateTime, source: :disabledAt)
    # Stable grant intent for an unfinished automatic SSO admission.
    # Cleared after the authoritative grant is observed; never inferred from missing access.
    field(:pending_sso_grant_id, :string, source: :pendingSsoGrantId)
    # A new generation for each membership lifetime. Stamped USER grant facts
    # must match the current active membership before they can be projected;
    # removing and rejoining therefore cannot revive a delayed old attach.
    field(:membership_stamp, :string, read_after_writes: true, source: :membershipStamp)
    # When this membership started holding the full seat it holds now; null means createdAt.
    # Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
    field(:full_seat_since, LangWatch.PrismaDateTime, source: :fullSeatSince)
    # Set when admission wrote this membership disabled because no seat was free: the person
    # is waiting for one. Cleared when the membership is enabled.
    field(:seat_pending_since, LangWatch.PrismaDateTime, source: :seatPendingSince)
    field(:user_id, :string, primary_key: true, source: :userId)
    field(:organization_id, :string, primary_key: true, source: :organizationId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    belongs_to(:organization, LangWatch.Organization, foreign_key: :organization_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end