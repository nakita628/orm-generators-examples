defmodule Blinko.McpServers do
  use Ecto.Schema

  @primary_key {:id, :id, autogenerate: true}
  @timestamps_opts [type: Blinko.PrismaDateTime, autogenerate: {Blinko.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: integer(),
          name: String.t(),
          description: String.t() | nil,
          type: String.t(),
          command: String.t() | nil,
          args: map() | list() | String.t() | number() | boolean() | nil,
          url: String.t() | nil,
          env: map() | list() | String.t() | number() | boolean() | nil,
          headers: map() | list() | String.t() | number() | boolean() | nil,
          is_enabled: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t()
        }

  schema "mcpServers" do
    field(:name, :string)
    field(:description, :string)
    field(:type, :string)
    field(:command, :string)
    field(:args, Blinko.PrismaJson)
    field(:url, :string)
    field(:env, Blinko.PrismaJson)
    field(:headers, Blinko.PrismaJson)
    field(:is_enabled, :boolean, default: true, source: :isEnabled)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end