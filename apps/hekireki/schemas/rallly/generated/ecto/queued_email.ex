defmodule Rallly.QueuedEmail do
  @moduledoc """
  An email waiting to be sent by the queue, one row per recipient. The
  email is built when it is sent, from the row its subject points at.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Rallly.PrismaDateTime, autogenerate: {Rallly.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          kind: String.t(),
          subject_id: String.t(),
          batch_id: String.t(),
          status: atom(),
          send_after: DateTime.t() | nil,
          attempts: integer(),
          claimed_at: DateTime.t() | nil,
          sent_at: DateTime.t() | nil,
          last_error: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Rallly.User.t() | nil
        }

  schema "queued_emails" do
    field(:id, :string, primary_key: true)
    # Which handler builds the email. A string rather than an enum so a new
    # kind needs no migration; the claim only takes kinds the build knows.
    field(:kind, :string)
    # The row the email is about, interpreted by the kind. Not a foreign key:
    # kinds point at different tables, and a subject deleted before the send
    # makes its handler skip the email.
    field(:subject_id, :string)
    # Groups the emails one action queued, so the action can send its own
    # first batch at once and its progress can be read back.
    field(:batch_id, :string)
    field(:status, Ecto.Enum, values: [:pending, :sent, :skipped, :failed], default: :pending)
    # Not claimed before this time. Unset sends at the next run.
    field(:send_after, Rallly.PrismaDateTime)
    field(:attempts, :integer, default: 0)
    field(:claimed_at, Rallly.PrismaDateTime)
    field(:sent_at, Rallly.PrismaDateTime)
    field(:last_error, :string)
    # The account whose action queued the email. Banning it skips whatever
    # is still pending.
    belongs_to(:user, Rallly.User, foreign_key: :user_id, type: :string)
    timestamps(inserted_at_source: :created_at)
  end
end