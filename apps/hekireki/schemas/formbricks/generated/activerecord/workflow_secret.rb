# Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
# store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
# version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
# with the workflow. A value is only sent to the origin it was set for (ADR 0051).
#
# value:: `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned.
# boundOrigin:: Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
#   sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
#   until it is set again.
class WorkflowSecret < ApplicationRecord
  self.table_name = "WorkflowSecret"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid2.call }

  validates :value, presence: true

  belongs_to :workflow, foreign_key: [ :workflowId, :workspaceId ], primary_key: [ :id, :workspaceId ],
                        inverse_of: :secrets
end
