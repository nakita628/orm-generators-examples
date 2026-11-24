class OAuthRefreshToken < ApplicationRecord
  self.table_name = "OAuthRefreshToken"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :scope, default: -> { [] }
  alias_attribute :created_at, :createdAt

  validates :hashedToken, presence: true, uniqueness: true
  validates :resource, presence: true
  validates :idleExpiresAt, presence: true

  belongs_to :personal_access_token, foreign_key: "personalAccessTokenId", inverse_of: :oauth_refresh_tokens
  belongs_to :oauth_client, class_name: "OAuthClient", foreign_key: "oauthClientId", inverse_of: :refresh_tokens
end
