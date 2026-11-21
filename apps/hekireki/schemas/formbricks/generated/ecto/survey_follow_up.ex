defmodule Formbricks.SurveyFollowUp do
  @moduledoc """
  Defines follow-up actions for survey responses.
  Enables automated actions based on specific survey response conditions.

  @property id - Unique identifier for the follow-up
  @property survey - The associated survey
  @property name - Display name for the follow-up
  @property trigger - Conditions that activate the follow-up
  @property action - Actions to take when triggered
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          trigger: map(),
          action: map(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          survey: Formbricks.Survey.t() | nil
        }

  schema "SurveyFollowUp" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    # [SurveyFollowUpTrigger]
    field(:trigger, :map)
    # [SurveyFollowUpAction]
    field(:action, :map)
    field(:survey_id, :string, source: :surveyId)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end