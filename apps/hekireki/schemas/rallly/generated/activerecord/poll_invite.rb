# token:: Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
# revoked_at:: Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
# participant_id:: Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
class PollInvite < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :email, presence: true, uniqueness: { scope: :poll_id }
  validates :token, presence: true, uniqueness: true
  validates :participant_id, uniqueness: { allow_nil: true }

  belongs_to :poll, inverse_of: :invites
  belongs_to :participant, optional: true, inverse_of: :invite
end
