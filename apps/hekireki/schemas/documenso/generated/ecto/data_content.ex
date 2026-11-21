defmodule Documenso.DataContent do
  @moduledoc """
  Generic stored data (e.g. an uploaded image) which other models reference.
  The kind of data and its kind specific details live in `metadata`.
  @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          type: atom(),
          data: String.t(),
          metadata: map(),
          envelope_contents: [Documenso.EnvelopeContent.t()]
        }

  schema "DataContent" do
    field(:id, :string, primary_key: true)
    # How `data` is stored, e.g. an S3 path or inline bytes.
    field(:type, Ecto.Enum, values: [:S3_PATH, :BYTES, :BYTES_64])
    field(:data, :string)
    # [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
    field(:metadata, :map)
    has_many(:envelope_contents, Documenso.EnvelopeContent, foreign_key: :data_content_id, references: :id)
  end
end