class Answer < ApplicationRecord
  self.table_name = "Answer"
  self.primary_key = %w[resultId blockId groupId]

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :createdAt

  validates :blockId, presence: true, uniqueness: { scope: [ :resultId, :groupId ] }
  validates :groupId, presence: true
  validates :content, presence: true

  belongs_to :result, foreign_key: "resultId", inverse_of: :answers
end
