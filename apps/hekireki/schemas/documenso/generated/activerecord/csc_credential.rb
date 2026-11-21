# Per-recipient cached credential metadata for the configured Cloud Signature
# Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
# symmetric-encrypted service-scope access token. One row per recipient on a
# CSC envelope; absent for SES recipients.
#
# certCache:: TSP-validated certificate chain (length-prefixed). Nullable until the
#   service-scope OAuth callback successfully retrieves and validates it.
# signatureAlgorithm:: Algorithm metadata derived from the CSC `credentials/info` response.
#   Held as flat scalars so each is independently queryable.
# serviceTokenCiphertext:: Service-scope access token, symmetric-encrypted at the application layer
#   before persistence. Nullable until the service-scope OAuth callback runs.
class CscCredential < ApplicationRecord
  self.table_name = "CscCredential"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :providerId, presence: true
  validates :credentialId, presence: true
  validates :signatureAlgorithm, presence: true
  validates :keyType, presence: true
  validates :digestAlgorithm, presence: true
  validates :recipientId, uniqueness: true

  belongs_to :recipient, foreign_key: "recipientId", inverse_of: :csc_credential
end
