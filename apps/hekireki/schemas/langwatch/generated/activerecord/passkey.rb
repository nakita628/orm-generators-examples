# The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
# public key and the signature counter are the plugin's to read and write.
# A passkey also becomes an `Identifier` row with provider `passkey`, which
# is what makes it a sign-in method like any other - this table is the
# credential, that one is the identity.
#
# `aaguid` names the authenticator MODEL, never a device or a person, and is
# only present at registration. Privacy-preserving platforms report all
# zeroes, so a label derived from it is best-effort by construction.
class Passkey < ApplicationRecord
  self.table_name = "Passkey"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :counter, default: 0
  attribute :backedUp, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :publicKey, presence: true
  validates :credentialID, presence: true
  validates :deviceType, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :passkeys
end
