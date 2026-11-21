defmodule Formbricks.Display do
  @moduledoc """
  Records when a survey is shown to a user.
  Tracks survey display history and response status.

  @property id - Unique identifier for the display event
  @property survey - The survey that was displayed
  @property contact - The contact who saw the survey
  @property response - The associated response if one exists
  @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          creative_id: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          survey: Formbricks.Survey.t() | nil,
          contact: Formbricks.Contact.t() | nil,
          response: Formbricks.Response.t() | nil
        }

  schema "Display" do
    field(:id, :string, primary_key: true)
    field(:creative_id, :string, source: :creativeId)
    field(:survey_id, :string, source: :surveyId)
    field(:contact_id, :string, source: :contactId)
    belongs_to(:survey, Formbricks.Survey, foreign_key: :survey_id, define_field: false, type: :string)
    belongs_to(:contact, Formbricks.Contact, foreign_key: :contact_id, define_field: false, type: :string)
    has_one(:response, Formbricks.Response, foreign_key: :display_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end