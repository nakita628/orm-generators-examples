defmodule LangWatch.Passkey do
  @moduledoc """
  The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
  public key and the signature counter are the plugin's to read and write.
  A passkey also becomes an `Identifier` row with provider `passkey`, which
  is what makes it a sign-in method like any other - this table is the
  credential, that one is the identity.

  `aaguid` names the authenticator MODEL, never a device or a person, and is
  only present at registration. Privacy-preserving platforms report all
  zeroes, so a label derived from it is best-effort by construction.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: LangWatch.PrismaDateTime, autogenerate: {LangWatch.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          name: String.t() | nil,
          public_key: String.t(),
          credential_id: String.t(),
          counter: integer(),
          device_type: String.t(),
          backed_up: boolean(),
          transports: String.t() | nil,
          aaguid: String.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: LangWatch.User.t() | nil
        }

  schema "Passkey" do
    field(:id, :string, primary_key: true)
    field(:name, :string)
    field(:public_key, :string, source: :publicKey)
    field(:credential_id, :string, source: :credentialID)
    field(:counter, :integer, default: 0)
    field(:device_type, :string, source: :deviceType)
    field(:backed_up, :boolean, default: false, source: :backedUp)
    field(:transports, :string)
    field(:aaguid, :string)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, LangWatch.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end