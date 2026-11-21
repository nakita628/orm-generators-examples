# scopeId:: *
#   * Whose cap and discount this invite counts against — the fund's team, or the
#   * sender alone. One non-null key rather than a pair of nullable ones: on a
#   * nullable teamId every personal invite is distinct to Postgres, so the two
#   * keys guarded different scopes and a member could not invite a company both
#   * for their fund and for themselves.
class ReferralInvite < ApplicationRecord
  self.table_name = "ReferralInvite"

  attribute :id, default: -> { Cuid.generate }

  validates :scopeId, presence: true, uniqueness: { scope: :email }
  validates :email, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :referral_invites
  belongs_to :team, foreign_key: "teamId", optional: true, inverse_of: :referral_invites
end
