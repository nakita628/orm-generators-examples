defmodule LiteLLM.LiteLLMVerifiedSubject do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :subject_id}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          issuer: String.t(),
          tenant_id: String.t(),
          oid: String.t(),
          kind: String.t(),
          verified_via: String.t(),
          verified_at: DateTime.t(),
          user: LiteLLM.LiteLLMUserTable.t() | nil
        }

  schema "LiteLLM_VerifiedSubject" do
    field(:issuer, :string)
    field(:tenant_id, :string)
    field(:oid, :string)
    field(:kind, :string, default: "human")
    field(:verified_via, :string, default: "sso_interactive")
    field(:verified_at, LiteLLM.PrismaDateTime, autogenerate: true)
    belongs_to(:user, LiteLLM.LiteLLMUserTable, foreign_key: :user_id, references: :user_id)
  end
end