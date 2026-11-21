# A saved layout for a generated PDF, kept per team workspace under a name:
# which columns, rows and header lines print, and in what order the rows
# come. A row is a layout, not a generated file; the person generating a PDF
# picks a template by name.
#
# config:: The layout's switches. Read through `parseBookingTemplateConfig`,
#   which fills any key a stored row lacks from the Checklist template's
#   values.
# isDefault:: The template the PDF dialog offers first. One per workspace and kind,
#   enforced by a partial unique index in raw SQL (migration
#   20261103120000_add_pdf_templates). Names are also unique per workspace
#   and kind ignoring case, through an expression index in the same
#   migration.
class PdfTemplate < ApplicationRecord
  self.table_name = "PdfTemplate"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :config, default: -> { JSON.parse("{}") }
  attribute :isDefault, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :kind, { booking: "BOOKING" }, default: :booking, validate: true

  validates :name, presence: true, uniqueness: { scope: [ :organizationId, :kind ] }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :pdf_templates
end
