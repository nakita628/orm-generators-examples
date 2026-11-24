defmodule LiteLLM.LiteLLMSearchToolsTable do
  use Ecto.Schema

  @primary_key {:id, :string, autogenerate: {Ecto.UUID, :generate, []}, source: :search_tool_id}
  @foreign_key_type :string
  @timestamps_opts [type: LiteLLM.PrismaDateTime, autogenerate: {LiteLLM.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          search_tool_name: String.t(),
          litellm_params: map() | list() | String.t() | number() | boolean(),
          search_tool_info: map() | list() | String.t() | number() | boolean() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "LiteLLM_SearchToolsTable" do
    field(:search_tool_name, :string)
    field(:litellm_params, LiteLLM.PrismaJson)
    field(:search_tool_info, LiteLLM.PrismaJson)
    timestamps(inserted_at_source: :created_at)
  end
end