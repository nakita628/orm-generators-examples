# One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage.
#
# name:: A name given by hand; a named version is never thinned by retention.
# trigger:: "autosave", "manual", "experiment_run", "restore" or "initial".
# contentHash:: SHA-256 over the columns and the rows in order; equal content saves no new version.
# authorKind:: "user", "api_key" or "system": who saved the version when authorId is null.
# pinnedAt:: Set when an experiment pins this version; a pinned version is never thinned.
class DatasetVersion < ApplicationRecord
  self.table_name = "DatasetVersion"

  alias_attribute :created_at, :createdAt

  validates :projectId, presence: true
  validates :datasetId, presence: true, uniqueness: { scope: :version }
  validates :version, presence: true
  validates :trigger, presence: true
  validates :contentHash, presence: true
  validates :columnTypes, exclusion: { in: [ nil ] }
  validates :rowCount, presence: true
  validates :sizeBytes, presence: true
  validates :chunkCount, presence: true
end
