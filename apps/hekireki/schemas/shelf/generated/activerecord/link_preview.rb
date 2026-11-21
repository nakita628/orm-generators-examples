# What Shelf knows about one web link used in a workspace's custom fields,
# keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
# the link preview job (fetch columns) and people (customTitle).
#
# url:: `LinkKey` from @shelf/links, never a raw field value. At most 2048
#   characters (CHECK constraint in the migration).
# outcome:: Null until the first fetch finishes.
# nextFetchAt:: When the link may be fetched next; null means never again. A claim moves
#   it a few minutes ahead, which is also the lease on a fetch in flight.
# customTitle:: The name a person gave the link in this workspace; wins over pageTitle.
class LinkPreview < ApplicationRecord
  self.table_name = "LinkPreview"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :attempts, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :outcome, { title: "TITLE", no_title: "NO_TITLE", sign_in: "SIGN_IN", blocked: "BLOCKED", unreachable: "UNREACHABLE" },
       validate: { allow_nil: true }

  validates :url, presence: true, uniqueness: { scope: :organizationId }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :link_previews
end
