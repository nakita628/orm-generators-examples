# capturedAt:: When the photo was taken, from its EXIF data (or the phone app's copy of
#   it). Null when the photo carried no usable date. EXIF can be edited, so
#   this is the best signal available, not proof.
# captureSource:: How the photo reached Shelf. Null only for photos uploaded before this
#   field existed; a phone upload with no or an unknown source is PHONE.
class AuditImage < ApplicationRecord
  self.table_name = "AuditImage"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :captureSource, { camera: "CAMERA", library: "LIBRARY", web_upload: "WEB_UPLOAD", phone: "PHONE" },
       validate: { allow_nil: true }

  validates :imageUrl, presence: true

  belongs_to :audit_session, foreign_key: "auditSessionId", inverse_of: :images
  belongs_to :audit_asset, foreign_key: "auditAssetId", optional: true, inverse_of: :images
  belongs_to :uploaded_by, class_name: "User", foreign_key: "uploadedById", optional: true,
                           inverse_of: :audit_images_uploaded
  belongs_to :organization, foreign_key: "organizationId", inverse_of: :audit_images
end
