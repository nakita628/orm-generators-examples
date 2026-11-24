defmodule Formbricks.OutboundEmailAuditErasureIntent do
  @moduledoc """
  Holds an account erasure request across an in-flight email send. The address is never stored.
  The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
  """

  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string
  @timestamps_opts [type: Formbricks.PrismaDateTime, autogenerate: {Formbricks.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          user_id: String.t() | nil,
          user_hash: String.t(),
          recipient_mac: String.t(),
          activated_at: DateTime.t() | nil,
          checked_at: DateTime.t() | nil,
          expires_at: DateTime.t(),
          inserted_at: DateTime.t()
        }

  schema "OutboundEmailAuditErasureIntent" do
    field(:user_id, :string, source: :userId)
    field(:user_hash, :string, source: :userHash)
    field(:recipient_mac, :string, source: :recipientMac)
    field(:activated_at, Formbricks.PrismaDateTime, source: :activatedAt)
    field(:checked_at, Formbricks.PrismaDateTime, source: :checkedAt)
    field(:expires_at, Formbricks.PrismaDateTime, source: :expiresAt)
    timestamps(inserted_at_source: :createdAt, updated_at: false)
  end
end