defmodule LiteLLM.LiteLLMClaudeCodePluginUpload do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          plugin_name: String.t(),
          content: binary(),
          plugin: LiteLLM.LiteLLMClaudeCodePluginTable.t() | nil
        }

  schema "LiteLLM_ClaudeCodePluginUpload" do
    field(:content, :binary)
    field(:plugin_name, :string, primary_key: true)
    belongs_to(:plugin, LiteLLM.LiteLLMClaudeCodePluginTable, foreign_key: :plugin_name, define_field: false, type: :string, references: :name)
  end
end