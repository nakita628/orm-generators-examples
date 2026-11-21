# Per-recipient transient sign-time state held across the credential-scope
# OAuth round-trip. The `recipientId @unique` constraint caps each recipient
# to at most one in-flight session — re-clicking Sign UPSERTs the row.
# Cleaned up transitively when the Recipient is deleted.
# @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
#
# envelopeId:: Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
#   flows through the Recipient cascade.
# signingTime:: Pinned at prep time; stable across overlay re-renders so the captured
#   signedAttrs digest matches the bytes the TSP signs.
# itemsJson:: Contract between prep and sign: an array of
#   `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
#   `documentDataId` pins the immutable bytes that produced the captured
#   digest so sign-time can re-derive the same hash against the same content.
#   [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
# encryptedSad:: Populated by the credential-scope OAuth callback. Both nullable at row
#   creation; set together when the SAD is exchanged.
class CscSession < ApplicationRecord
  self.table_name = "CscSession"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  validates :envelopeId, presence: true
  validates :signingTime, presence: true
  validates :itemsJson, exclusion: { in: [ nil ] }
  validates :recipientId, uniqueness: true

  belongs_to :recipient, foreign_key: "recipientId", inverse_of: :csc_session
end
