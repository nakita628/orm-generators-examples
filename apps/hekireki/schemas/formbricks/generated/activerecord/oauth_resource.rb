# Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
# its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
# a token is now bound to the resource it was approved for instead of to a static allow-list.
# `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
class OauthResource < ApplicationRecord
  self.table_name = "oauthResource"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :allowedScopes, default: -> { [] }
  attribute :dpopBoundAccessTokensRequired, default: false
  attribute :disabled, default: false
  attribute :policyVersion, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :identifier, presence: true, uniqueness: true
  validates :name, presence: true

  has_many :clients, class_name: "OauthClientResource", foreign_key: "resourceId", primary_key: "identifier",
                     inverse_of: :resource, dependent: :destroy
end
