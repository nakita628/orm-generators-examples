defmodule Formbricks.CreativeDailyPageStat do
  @moduledoc """
  A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
  `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          creative_id: String.t(),
          date: Date.t(),
          page: String.t(),
          impressions: integer(),
          clicks: integer(),
          creative: Formbricks.Creative.t() | nil
        }

  schema "CreativeDailyPageStat" do
    # UTC day.
    field(:date, Formbricks.PrismaDate, primary_key: true)
    # Host + path, without query string or hash, or `(other)`.
    field(:page, :string, primary_key: true)
    field(:impressions, :integer, default: 0)
    field(:clicks, :integer, default: 0)
    field(:creative_id, :string, primary_key: true, source: :creativeId)
    belongs_to(:creative, Formbricks.Creative, foreign_key: :creative_id, define_field: false, type: :string)
  end
end