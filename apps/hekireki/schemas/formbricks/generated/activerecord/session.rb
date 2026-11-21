# Stores active authentication sessions for revocable server-side login state.
#
# @property sessionToken - Opaque token stored in the browser cookie
# @property user - The Formbricks user who owns this session
# @property expires - Hard expiry for the session
class Session < ApplicationRecord
  self.table_name = "Session"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :sessionToken, presence: true, uniqueness: true
  validates :expires, presence: true

  belongs_to :user, foreign_key: "userId", inverse_of: :sessions
  has_many :oauth_access_tokens, foreign_key: "sessionId", inverse_of: :session, dependent: :nullify
  has_many :oauth_refresh_tokens, foreign_key: "sessionId", inverse_of: :session, dependent: :nullify
end
