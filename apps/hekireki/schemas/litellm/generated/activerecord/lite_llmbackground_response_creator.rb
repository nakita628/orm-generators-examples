class LiteLLMBackgroundResponseCreator < ApplicationRecord
  self.table_name = "LiteLLM_BackgroundResponseCreator"
  self.primary_key = "response_id"

  validates :upstream_url, presence: true
  validates :create_context, exclusion: { in: [ nil ] }
end
