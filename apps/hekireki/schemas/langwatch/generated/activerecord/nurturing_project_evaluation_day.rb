# Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
# written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
#
# day:: UTC `YYYY-MM-DD`.
class NurturingProjectEvaluationDay < ApplicationRecord
  self.table_name = "NurturingProjectEvaluationDay"
  self.primary_key = %w[projectId day]
end
