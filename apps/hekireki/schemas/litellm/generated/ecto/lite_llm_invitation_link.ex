defmodule LiteLLM.LiteLLMInvitationLink do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          is_accepted: boolean(),
          accepted_at: DateTime.t() | nil,
          expires_at: DateTime.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t(),
          lite_llm_user_table_user: LiteLLM.LiteLLMUserTable.t() | nil,
          lite_llm_user_table_created: LiteLLM.LiteLLMUserTable.t() | nil,
          lite_llm_user_table_updated: LiteLLM.LiteLLMUserTable.t() | nil
        }

  schema "LiteLLM_InvitationLink" do
    field(:is_accepted, :boolean, default: false)
    field(:accepted_at, LiteLLM.PrismaDateTime)
    field(:expires_at, LiteLLM.PrismaDateTime)
    field(:created_at, LiteLLM.PrismaDateTime)
    field(:updated_at, LiteLLM.PrismaDateTime)
    belongs_to(:lite_llm_user_table_user, LiteLLM.LiteLLMUserTable, foreign_key: :user_id, references: :user_id)
    belongs_to(:lite_llm_user_table_created, LiteLLM.LiteLLMUserTable, foreign_key: :created_by, references: :user_id)
    belongs_to(:lite_llm_user_table_updated, LiteLLM.LiteLLMUserTable, foreign_key: :updated_by, references: :user_id)
  end
end