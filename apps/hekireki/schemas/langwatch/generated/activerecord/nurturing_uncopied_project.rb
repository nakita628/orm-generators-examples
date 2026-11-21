# A project nurturing saw traces of with no placement copy, and when it first did: every worker
# retries its trace batches only within one window from then, across restarts. Deleted when the
# copy lands.
class NurturingUncopiedProject < ApplicationRecord
  self.table_name = "NurturingUncopiedProject"
  self.primary_key = "projectId"
end
