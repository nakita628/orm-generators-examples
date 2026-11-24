defmodule Ghostfolio.MarketData do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Ghostfolio.PrismaDateTime, autogenerate: {Ghostfolio.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          data_source: atom(),
          date: DateTime.t(),
          is_carried_forward: boolean(),
          market_price: float(),
          state: atom(),
          symbol: String.t(),
          inserted_at: DateTime.t()
        }

  schema "MarketData" do
    field(:data_source, Ecto.Enum, values: [:ALPHA_VANTAGE, :COINGECKO, :EOD_HISTORICAL_DATA, :FINANCIAL_MODELING_PREP, :GHOSTFOLIO, :GOOGLE_SHEETS, :MANUAL, :RAPID_API, :YAHOO], source: :dataSource)
    field(:date, Ghostfolio.PrismaDateTime)
    field(:is_carried_forward, :boolean, default: false, source: :isCarriedForward)
    field(:market_price, :float, source: :marketPrice)
    field(:state, Ecto.Enum, values: [:CLOSE, :INTRADAY], default: :CLOSE)
    field(:symbol, :string)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end