# cliVersion:: CLI version that initiated the deploy, stamped at initialization or worker upload
# templateKey:: Key of the microVM template prebuilt from this deployment's image, set when the builder succeeds
# indexObjectKey:: Object key of the full index snapshot in the deployment index store. Set only after the upload succeeds.
# indexSha256:: sha256 hex of the snapshot's uncompressed JSON
# buildEnvVars:: Encrypted build-time env vars for pre-bundled (fromBundle) deploys, as an
#   EncryptedSecretValue envelope. Kept for the redeploy window, then cleared by a sweep.
# initiatedByCredential:: The kind of credential that started the deployment. Null when unknown.
# supersededById:: The newer GitHub build that superseded this one while it was still queued
class WorkerDeployment < ApplicationRecord
  self.table_name = "WorkerDeployment"
  self.inheritance_column = nil
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :imagePlatform, default: "linux/amd64"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :status, { pending: "PENDING", installing: "INSTALLING", building: "BUILDING", deploying: "DEPLOYING", deployed: "DEPLOYED", failed: "FAILED", canceled: "CANCELED", timed_out: "TIMED_OUT" },
       default: :pending, validate: true
  enum :type, { managed: "MANAGED", unmanaged: "UNMANAGED", v1: "V1" }, default: :v1, validate: true
  enum :initiatedByCredential, { root_api_key: "ROOT_API_KEY", scoped_api_key: "SCOPED_API_KEY", personal_access_token: "PERSONAL_ACCESS_TOKEN", organization_access_token: "ORGANIZATION_ACCESS_TOKEN", trusted_publishing: "TRUSTED_PUBLISHING" },
       validate: { allow_nil: true }

  validates :contentHash, presence: true
  validates :friendlyId, presence: true, uniqueness: true
  validates :shortCode, presence: true, uniqueness: { scope: :projectId }
  validates :version, presence: true, uniqueness: { scope: :environmentId }
  validates :workerId, uniqueness: { allow_nil: true }

  belongs_to :project, foreign_key: "projectId", inverse_of: :worker_deployment
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId",
                           inverse_of: :worker_deployments
  belongs_to :worker, class_name: "BackgroundWorker", foreign_key: "workerId", optional: true, inverse_of: :deployment
  belongs_to :triggered_by, class_name: "User", foreign_key: "triggeredById", optional: true, inverse_of: :deployments
  has_many :promotions, class_name: "WorkerDeploymentPromotion", foreign_key: "deploymentId", inverse_of: :deployment,
                        dependent: :destroy
  has_many :alerts, class_name: "ProjectAlert", foreign_key: "workerDeploymentId", inverse_of: :worker_deployment,
                    dependent: :destroy
  has_many :worker_instance, foreign_key: "deploymentId", inverse_of: :deployment, dependent: :nullify
  has_many :integration_deployments, foreign_key: "deploymentId", inverse_of: :worker_deployment, dependent: :nullify
end
