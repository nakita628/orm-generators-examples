defmodule Rallly.Participant do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Rallly.PrismaDateTime, autogenerate: {Rallly.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t(),
          email: String.t() | nil,
          note: String.t() | nil,
          token: String.t(),
          locale: String.t() | nil,
          time_zone: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t() | nil,
          poll: Rallly.Poll.t() | nil,
          user: Rallly.User.t() | nil,
          invite: Rallly.PollInvite.t() | nil,
          votes: [Rallly.Vote.t()]
        }

  schema "participants" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:email, :string)
    field(:note, :string)
    # Per response edit link: grants edit access to this response without a session. App minted at creation, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet — url safe with no linkifier edge cases, unlike base64url); never a DB default so a code path that forgets to mint fails loudly. A response created through an emailed invite takes the invite's token. Rows that predate minting were backfilled with gen_random_uuid() hex
    field(:token, :string)
    field(:locale, :string)
    field(:time_zone, :string)
    belongs_to(:poll, Rallly.Poll, foreign_key: :poll_id, type: :string)
    belongs_to(:user, Rallly.User, foreign_key: :user_id, type: :string)
    has_one(:invite, Rallly.PollInvite, foreign_key: :participant_id, references: :id)
    has_many(:votes, Rallly.Vote, foreign_key: :participant_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end