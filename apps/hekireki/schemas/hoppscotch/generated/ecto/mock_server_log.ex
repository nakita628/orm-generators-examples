defmodule Hoppscotch.MockServerLog do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          request_method: String.t(),
          request_path: String.t(),
          request_headers: map() | list() | String.t() | number() | boolean(),
          request_body: map() | list() | String.t() | number() | boolean() | nil,
          request_query: map() | list() | String.t() | number() | boolean() | nil,
          response_status: integer(),
          response_headers: map() | list() | String.t() | number() | boolean(),
          response_body: map() | list() | String.t() | number() | boolean() | nil,
          response_time: integer(),
          ip_address: String.t() | nil,
          user_agent: String.t() | nil,
          executed_at: DateTime.t(),
          mock_server: Hoppscotch.MockServer.t() | nil
        }

  schema "MockServerLog" do
    field(:id, :string, primary_key: true)
    field(:request_method, :string, source: :requestMethod)
    field(:request_path, :string, source: :requestPath)
    field(:request_headers, Hoppscotch.PrismaJson, source: :requestHeaders)
    field(:request_body, Hoppscotch.PrismaJson, source: :requestBody)
    field(:request_query, Hoppscotch.PrismaJson, source: :requestQuery)
    field(:response_status, :integer, source: :responseStatus)
    field(:response_headers, Hoppscotch.PrismaJson, source: :responseHeaders)
    field(:response_body, Hoppscotch.PrismaJson, source: :responseBody)
    field(:response_time, :integer, source: :responseTime)
    field(:ip_address, :string, source: :ipAddress)
    field(:user_agent, :string, source: :userAgent)
    field(:executed_at, Hoppscotch.PrismaDateTime, autogenerate: true, source: :executedAt)
    field(:mock_server_id, :string, source: :mockServerID)
    belongs_to(:mock_server, Hoppscotch.MockServer, foreign_key: :mock_server_id, define_field: false, type: :string)
  end
end