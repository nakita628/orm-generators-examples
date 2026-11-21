defmodule Rallly.PollInvite do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Rallly.PrismaDateTime, autogenerate: {Rallly.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          email: String.t(),
          token: String.t(),
          opened_at: DateTime.t() | nil,
          reminded_at: DateTime.t() | nil,
          revoked_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          poll: Rallly.Poll.t() | nil,
          participant: Rallly.Participant.t() | nil
        }

  schema "poll_invites" do
    field(:id, :string, primary_key: true)
    field(:email, :string)
    # Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access
    field(:token, :string)
    field(:opened_at, Rallly.PrismaDateTime)
    field(:reminded_at, Rallly.PrismaDateTime)
    # Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token
    field(:revoked_at, Rallly.PrismaDateTime)
    belongs_to(:poll, Rallly.Poll, foreign_key: :poll_id, type: :string)
    # Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite
    belongs_to(:participant, Rallly.Participant, foreign_key: :participant_id, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end