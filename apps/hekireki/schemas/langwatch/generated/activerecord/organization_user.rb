# departmentId:: Department assignment for spend attribution (the people lens).
#   Single-valued, org-scoped, accounting-only (never gates access).
#   Plain column ref (relationMode="prisma"); the service validates the
#   id belongs to this org. See departments.feature.
# billingOnly:: A Billing-only member (ADR-205): role EXTERNAL, the Billing grant, no team and
#   no seat. Any role change clears it. See billing-only-invite.feature.
# disabledAt:: Set when an admin disables this membership to stay within the licensed
#   seat count. The person loses access to this organization only, keeps
#   their role, department and everything they did, and frees their seat.
#   Reversible: re-enabling restores access if a seat is available.
#   Seat counting excludes disabled memberships. See
#   seat-reconciliation.feature.
# pendingSsoGrantId:: Stable grant intent for an unfinished automatic SSO admission.
#   Cleared after the authoritative grant is observed; never inferred from missing access.
# membershipStamp:: A new generation for each membership lifetime. Stamped USER grant facts
#   must match the current active membership before they can be projected;
#   removing and rejoining therefore cannot revive a delayed old attach.
# fullSeatSince:: When this membership started holding the full seat it holds now; null means createdAt.
#   Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
# seatPendingSince:: Set when admission wrote this membership disabled because no seat was free: the person
#   is waiting for one. Cleared when the membership is enabled.
class OrganizationUser < ApplicationRecord
  self.table_name = "OrganizationUser"
  self.primary_key = %w[userId organizationId]

  attribute :billingOnly, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :role, { admin: "ADMIN", member: "MEMBER", external: "EXTERNAL", developer: "DEVELOPER" }, validate: true

  belongs_to :user, foreign_key: "userId", inverse_of: :org_memberships
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :members
end
