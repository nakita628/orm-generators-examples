defmodule Formbricks.FeedbackSourceFormbricksMapping do
  @moduledoc """
  Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
  Each row represents one element that will create FeedbackRecords when answered.

  @property id - Unique identifier for the mapping
  @property feedbackSource - The parent feedback source
  @property survey - The survey containing the element
  @property elementId - The element ID within the survey (from blocks[].elements[].id)
  @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
  @property customFieldLabel - Optional override for the element headline as field_label in Hub
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          feedback_source_id: String.t(),
          workspace_id: String.t(),
          survey_id: String.t(),
          element_id: String.t(),
          hub_field_type: atom(),
          custom_field_label: String.t() | nil,
          inserted_at: DateTime.t()
        }

  schema "FeedbackSourceFormbricksMapping" do
    field(:id, :string, primary_key: true)
    field(:feedback_source_id, :string)
    field(:workspace_id, :string, source: :workspaceId)
    field(:survey_id, :string, source: :surveyId)
    field(:element_id, :string, source: :elementId)
    field(:hub_field_type, Ecto.Enum, values: [:text, :categorical, :nps, :csat, :ces, :rating, :number, :boolean, :date], source: :hubFieldType)
    field(:custom_field_label, :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end