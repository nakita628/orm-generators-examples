defmodule Formbricks.RetentionSurveyRestore do
  @moduledoc """
  When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
  and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
  so retention adds nothing to the survey row.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          survey_id: String.t(),
          restored_at: DateTime.t(),
          survey: Formbricks.Survey.t() | nil
        }

  schema "RetentionSurveyRestore" do
    field(:restored_at, Formbricks.PrismaDateTime, source: :restoredAt)
    field(:survey_id, :string, primary_key: true, source: :surveyId)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
  end
end