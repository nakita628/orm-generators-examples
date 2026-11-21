defmodule Shelf.PrintBatch do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          printed: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          qr_codes: [Shelf.Qr.t()]
        }

  schema "PrintBatch" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:printed, :boolean, default: false)
    has_many(:qr_codes, Shelf.Qr, foreign_key: :batch_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end