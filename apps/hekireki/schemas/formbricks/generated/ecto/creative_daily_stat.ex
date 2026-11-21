defmodule Formbricks.CreativeDailyStat do
  @moduledoc """
  A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
  only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
  running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
  the creative deletes its counts.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          creative_id: String.t(),
          date: Date.t(),
          device_type: atom(),
          impressions: integer(),
          clicks: integer(),
          creative: Formbricks.Creative.t() | nil
        }

  schema "CreativeDailyStat" do
    # UTC day.
    field(:date, Formbricks.PrismaDate, primary_key: true)
    field(:device_type, Ecto.Enum, values: [:phone, :tablet, :desktop], source: :deviceType)
    field(:impressions, :integer, default: 0)
    field(:clicks, :integer, default: 0)
    field(:creative_id, :string, primary_key: true, source: :creativeId)
    belongs_to(:creative, Formbricks.Creative, foreign_key: :creative_id, define_field: false, type: :string)
  end
end