# @zod.import(["import { ZEnvelopeContentMetaSchema } from '@documenso/lib/types/envelope-content-meta';"])
#
# contentMeta:: [EnvelopeContentMeta] @zod.custom.use(ZEnvelopeContentMetaSchema)
# dataContentId:: Any number of contents may point at the same data content, e.g. a
#   document created from a template shares the template's images.
class EnvelopeContent < ApplicationRecord
  self.table_name = "EnvelopeContent"

  validates :contentMeta, exclusion: { in: [ nil ] }

  belongs_to :envelope_item, foreign_key: "envelopeItemId", inverse_of: :contents
  belongs_to :envelope, foreign_key: "envelopeId", inverse_of: :contents
  belongs_to :data_content, foreign_key: "dataContentId", optional: true, inverse_of: :envelope_contents
end
