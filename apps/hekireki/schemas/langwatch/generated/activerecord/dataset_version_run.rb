# An experiment run that read one dataset version; a version a run read is never thinned.
class DatasetVersionRun < ApplicationRecord
  self.table_name = "DatasetVersionRun"

  alias_attribute :created_at, :createdAt

  validates :projectId, presence: true, uniqueness: { scope: [ :runId, :datasetId ] }
  validates :datasetId, presence: true
  validates :versionId, presence: true
  validates :runId, presence: true
end
