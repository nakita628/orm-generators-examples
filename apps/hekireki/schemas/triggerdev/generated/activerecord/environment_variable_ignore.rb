# A reason the environment variables "Needs attention" page shows a key in one environment, which
# the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
#
# variable:: For MISSING, a variable of the project with no value in this environment.
class EnvironmentVariableIgnore < ApplicationRecord
  self.table_name = "EnvironmentVariableIgnore"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  enum :reason, { empty: "EMPTY", missing: "MISSING", oidc: "OIDC" }, validate: true

  validates :environmentId, uniqueness: { scope: [ :variableId, :reason ] }

  belongs_to :project, foreign_key: "projectId", inverse_of: :environment_variable_ignores
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId",
                           inverse_of: :environment_variable_ignores
  belongs_to :variable, class_name: "EnvironmentVariable", foreign_key: "variableId", inverse_of: :ignores
  belongs_to :created_by, class_name: "User", foreign_key: "createdById", optional: true,
                          inverse_of: :environment_variable_ignores
end
