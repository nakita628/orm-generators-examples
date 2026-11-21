# The dated match table: "`user_abc123` on Anthropic is platform user #42,
# valid from March 3", closed when it stops being true.
#
# Dating is what makes a re-issued email address survivable: the leaver's link
# closes at offboarding, the new hire gets a new link, and January's spend
# stays with January's person. A correction closes the old row and opens a new
# one; nothing is ever silently rewritten.
#
# discoveredPersonId:: The provider-side identity this link is about.
# userId:: The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
#   the row and its dates remain, so the history of the link is not rewritten
#   by the erasure either.
# evidenceKind:: directory_id | verified_email | human_confirmed. Exact verified-email
#   equality is the primary evidence; a directory id corroborates but never
#   stands alone (ADR-128 §12).
# validTo:: Null means the link is open. Offboarding and corrections close it.
class IdentityMatch < ApplicationRecord
  self.table_name = "IdentityMatch"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true
  validates :discoveredPersonId, presence: true
  validates :evidenceKind, presence: true
  validates :validFrom, presence: true
end
