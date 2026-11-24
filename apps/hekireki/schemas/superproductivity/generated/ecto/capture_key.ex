defmodule SuperProductivity.CaptureKey do
  @moduledoc """
  Public half of the user's capture key. The row outlives switching the
  feature off: `disabledAt` stops a stale device re-registering that key.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          user_id: integer(),
          public_key: String.t(),
          fingerprint: String.t(),
          created_at: integer(),
          disabled_at: integer() | nil,
          user: SuperProductivity.User.t() | nil
        }

  schema "capture_keys" do
    # Base64 SPKI DER of an ECDH P-256 public key.
    field(:public_key, :string)
    field(:fingerprint, :string)
    field(:created_at, :integer)
    field(:disabled_at, :integer)
    field(:user_id, :id, primary_key: true)
    belongs_to(:user, SuperProductivity.User, foreign_key: :user_id, define_field: false)
  end
end