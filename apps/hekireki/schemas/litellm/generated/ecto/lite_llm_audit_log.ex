defmodule LiteLLM.LiteLLMAuditLog do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}}
  @foreign_key_type :string

  @type t :: %__MODULE__{
          id: String.t(),
          updated_at: DateTime.t(),
          changed_by: String.t(),
          changed_by_api_key: String.t(),
          action: String.t(),
          table_name: String.t(),
          object_id: String.t(),
          before_value: map() | list() | String.t() | number() | boolean() | nil,
          updated_values: map() | list() | String.t() | number() | boolean() | nil
        }

  schema "LiteLLM_AuditLog" do
    field(:updated_at, LiteLLM.PrismaDateTime, autogenerate: true)
    field(:changed_by, :string, default: "")
    field(:changed_by_api_key, :string, default: "")
    field(:action, :string)
    field(:table_name, :string)
    field(:object_id, :string)
    field(:before_value, LiteLLM.PrismaJson)
    field(:updated_values, LiteLLM.PrismaJson)
  end
end