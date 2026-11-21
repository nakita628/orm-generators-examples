# Junction table linking tags to responses.
# Enables many-to-many relationship between tags and responses.
#
# @property response - The tagged response
# @property tag - The tag applied to the response
class TagsOnResponses < ApplicationRecord
  self.table_name = "TagsOnResponses"
  self.primary_key = %w[responseId tagId]

  belongs_to :response, foreign_key: "responseId", inverse_of: :tags
  belongs_to :tag, foreign_key: "tagId", inverse_of: :responses
end
