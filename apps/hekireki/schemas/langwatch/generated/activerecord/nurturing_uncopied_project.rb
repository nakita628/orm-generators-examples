# retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
# retry clock reads or writes this table; an operator step may drop it once old writers are gone.
class NurturingUncopiedProject < ApplicationRecord
  self.table_name = "NurturingUncopiedProject"
  self.primary_key = "projectId"
end
