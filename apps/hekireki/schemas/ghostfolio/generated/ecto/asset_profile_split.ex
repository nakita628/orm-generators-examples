defmodule Ghostfolio.AssetProfileSplit do
  @moduledoc """
  A stock split of an asset profile. The numerator is the number of shares
  held after the split, the denominator the number of shares held before, for
  example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.

  The ratio is stored as two integers instead of the resulting factor, so that
  it stays exact, for example 1/3 for a 1:3 reverse split.

  Activities before the date are adjusted by the ratio: the quantity is
  multiplied by it and the unit price divided by it, so that the total value
  of an activity stays the same.

  Market data is already split-adjusted by the data providers and must not be
  adjusted again.
  """

  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          date: DateTime.t(),
          denominator: integer(),
          numerator: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          symbol_profile: Ghostfolio.SymbolProfile.t() | nil
        }

  schema "AssetProfileSplit" do
    field(:date, Ghostfolio.PrismaDateTime)
    field(:denominator, :integer)
    field(:numerator, :integer)
    field(:symbol_profile_id, :string, source: :symbolProfileId)
    belongs_to(:symbol_profile, Ghostfolio.SymbolProfile, foreign_key: :symbol_profile_id, define_field: false)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end