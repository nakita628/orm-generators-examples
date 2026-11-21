defmodule Formbricks.TagsOnResponses do
  @moduledoc """
  Junction table linking tags to responses.
  Enables many-to-many relationship between tags and responses.

  @property response - The tagged response
  @property tag - The tag applied to the response
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          response_id: String.t(),
          tag_id: String.t(),
          response: Formbricks.Response.t() | nil,
          tag: Formbricks.Tag.t() | nil
        }

  schema "TagsOnResponses" do
    field(:response_id, :string, primary_key: true, source: :responseId)
    field(:tag_id, :string, primary_key: true, source: :tagId)
    belongs_to(:response, Formbricks.Response, foreign_key: :response_id, define_field: false, type: :string)
    belongs_to(:tag, Formbricks.Tag, foreign_key: :tag_id, define_field: false, type: :string)
  end
end