# name:: For example "us-east-1"
# masterQueue:: If managed, it will default to the name, e.g. "us-east-1"
#   If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
# description:: "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes
# location:: "usa", "europe", etc. Used like a pseudo enum for things like flags
# region:: Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs.
# enableFastPath:: When true, runs enqueued to this worker queue may skip the intermediate queue
#   and be pushed directly to the worker queue when concurrency is available.
class WorkerInstanceGroup < ApplicationRecord
  self.table_name = "WorkerInstanceGroup"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :hidden, default: false
  attribute :enableFastPath, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :type, { managed: "MANAGED", unmanaged: "UNMANAGED" }, validate: true
  enum :workloadType, { container: "CONTAINER", microvm: "MICROVM" }, default: :container, validate: true

  validates :name, presence: true
  validates :masterQueue, presence: true, uniqueness: true
  validates :tokenId, uniqueness: true

  belongs_to :token, class_name: "WorkerGroupToken", foreign_key: "tokenId", inverse_of: :worker_group
  belongs_to :organization, foreign_key: "organizationId", optional: true, inverse_of: :worker_groups
  belongs_to :project, foreign_key: "projectId", optional: true, inverse_of: :worker_groups
  has_many :workers, class_name: "WorkerInstance", foreign_key: "workerGroupId", inverse_of: :worker_group,
                     dependent: :restrict_with_error
  has_many :background_workers, foreign_key: "workerGroupId", inverse_of: :worker_group, dependent: :nullify
  has_many :default_for_projects, class_name: "Project", foreign_key: "defaultWorkerGroupId",
                                  inverse_of: :default_worker_group, dependent: :nullify
end
