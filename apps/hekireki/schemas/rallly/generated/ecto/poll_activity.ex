defmodule Rallly.PollActivity do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Rallly.PrismaDateTime, autogenerate: {Rallly.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          type: String.t(),
          user_id: String.t() | nil,
          participant_id: String.t() | nil,
          invite_id: String.t() | nil,
          option_id: String.t() | nil,
          payload: map() | list() | String.t() | number() | boolean(),
          inserted_at: DateTime.t(),
          poll: Rallly.Poll.t() | nil
        }

  schema "poll_activities" do
    field(:id, :string, primary_key: true)
    # Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations
    field(:type, :string)
    # Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question
    field(:user_id, :string)
    field(:participant_id, :string)
    field(:invite_id, :string)
    field(:option_id, :string)
    field(:payload, Rallly.PrismaJson, default: %{})
    belongs_to(:poll, Rallly.Poll, foreign_key: :poll_id, type: :string)
    timestamps(inserted_at_source: :created_at, updated_at: false)
  end
end