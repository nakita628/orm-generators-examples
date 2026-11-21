# token:: Per response edit link: grants edit access to this response without a session. App minted at creation, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet — url safe with no linkifier edge cases, unlike base64url); never a DB default so a code path that forgets to mint fails loudly. A response created through an emailed invite takes the invite's token. Rows that predate minting were backfilled with gen_random_uuid() hex
class Participant < ApplicationRecord
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :name, presence: true
  validates :token, presence: true, uniqueness: true

  belongs_to :poll, inverse_of: :participants
  belongs_to :user, optional: true, inverse_of: :participants
  has_one :invite, class_name: "PollInvite", inverse_of: :participant, dependent: :nullify
  has_many :votes, dependent: :destroy
end
