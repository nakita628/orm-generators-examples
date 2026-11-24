class OAuthClient < ApplicationRecord
  self.table_name = "OAuthClient"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :redirectUris, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :kind, { preregistered: "PREREGISTERED", cimd: "CIMD", dcr: "DCR" }, validate: true

  validates :clientId, presence: true, uniqueness: true
  validates :dedupeKey, uniqueness: { allow_nil: true }

  has_many :authorization_codes, foreign_key: "oauthClientId", inverse_of: :oauth_client, dependent: :destroy
  has_many :personal_access_tokens, foreign_key: "oauthClientId", inverse_of: :oauth_client,
                                    dependent: :restrict_with_error
  has_many :refresh_tokens, class_name: "OAuthRefreshToken", foreign_key: "oauthClientId", inverse_of: :oauth_client,
                            dependent: :destroy
end
