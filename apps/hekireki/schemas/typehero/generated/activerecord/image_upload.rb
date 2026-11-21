# id:: the id of the db row
# userId:: the user id fk
# user:: The user that this is related to
# url:: the url to the s3 uploaded assets
class ImageUpload < ApplicationRecord
  self.table_name = "ImageUpload"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :url, presence: true, length: { maximum: 256 }

  belongs_to :user, foreign_key: "userId", inverse_of: :image_upload
end
