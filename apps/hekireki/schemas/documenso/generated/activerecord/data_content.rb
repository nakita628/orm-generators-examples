# Generic stored data (e.g. an uploaded image) which other models reference.
# The kind of data and its kind specific details live in `metadata`.
# @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
#
# type:: How `data` is stored, e.g. an S3 path or inline bytes.
# metadata:: [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
class DataContent < ApplicationRecord
  self.table_name = "DataContent"
  self.inheritance_column = nil

  enum :type, { s3_path: "S3_PATH", bytes: "BYTES", bytes_64: "BYTES_64" }, validate: true

  validates :data, presence: true
  validates :metadata, exclusion: { in: [ nil ] }

  has_many :envelope_contents, foreign_key: "dataContentId", inverse_of: :data_content, dependent: :nullify
end
