# Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
# JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
#
# authorizationCodeId:: Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
#   binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
class OauthAccessToken < ApplicationRecord
  self.table_name = "oauthAccessToken"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :scopes, default: -> { [] }
  attribute :resources, default: -> { [] }
  attribute :requestedUserInfoClaims, default: -> { [] }
  alias_attribute :created_at, :createdAt

  validates :token, presence: true, uniqueness: true
  validates :expiresAt, presence: true

  belongs_to :client, class_name: "OauthClient", foreign_key: "clientId", primary_key: "clientId",
                      inverse_of: :access_tokens
  belongs_to :session, foreign_key: "sessionId", optional: true, inverse_of: :oauth_access_tokens
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :oauth_access_tokens
  belongs_to :refresh, class_name: "OauthRefreshToken", foreign_key: "refreshId", optional: true,
                       inverse_of: :access_tokens
end
