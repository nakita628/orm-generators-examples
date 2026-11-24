defmodule Langfuse.GatewaySpendLimit do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Langfuse.PrismaDateTime, autogenerate: {Langfuse.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          scope: atom(),
          interval: atom(),
          limit_usd: Decimal.t(),
          alert_thresholds: [integer()],
          counting_started_at: DateTime.t(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          organization: Langfuse.Organization.t() | nil,
          api_key: Langfuse.ApiKey.t() | nil,
          user: Langfuse.User.t() | nil,
          owner_user: Langfuse.User.t() | nil,
          alerts: [Langfuse.GatewaySpendLimitAlert.t()]
        }

  schema "gateway_spend_limits" do
    field(:id, :string, primary_key: true)
    field(:scope, Ecto.Enum, values: [:API_KEY, :USER, :API_KEY_DEFAULT, :USER_DEFAULT, :ORGANIZATION])
    field(:interval, Ecto.Enum, values: [:DAILY, :WEEKLY, :MONTHLY, :TOTAL])
    field(:limit_usd, :decimal)
    field(:alert_thresholds, {:array, :integer})
    field(:counting_started_at, Langfuse.PrismaDateTime)
    belongs_to(:organization, Langfuse.Organization, foreign_key: :organization_id, type: :string)
    belongs_to(:api_key, Langfuse.ApiKey, foreign_key: :api_key_id, type: :string)
    belongs_to(:user, Langfuse.User, foreign_key: :user_id, type: :string)
    belongs_to(:owner_user, Langfuse.User, foreign_key: :owner_user_id, type: :string)
    has_many(:alerts, Langfuse.GatewaySpendLimitAlert, foreign_key: :spend_limit_id, references: :id)
    timestamps(inserted_at_source: :created_at)
  end
end