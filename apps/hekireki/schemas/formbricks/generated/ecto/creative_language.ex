defmodule Formbricks.CreativeLanguage do
  @moduledoc """
  The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          language_id: String.t(),
          creative_id: String.t(),
          default: boolean(),
          enabled: boolean(),
          language: Formbricks.Language.t() | nil,
          creative: Formbricks.Creative.t() | nil
        }

  schema "CreativeLanguage" do
    field(:default, :boolean, default: false)
    field(:enabled, :boolean, default: true)
    field(:language_id, :string, primary_key: true, source: :languageId)
    field(:creative_id, :string, primary_key: true, source: :creativeId)
    belongs_to(:language, Formbricks.Language, foreign_key: :language_id, define_field: false, type: :string)
    belongs_to(:creative, Formbricks.Creative, foreign_key: :creative_id, define_field: false, type: :string)
  end
end