defmodule LangWatch.NurturingVirtualKey do
  @moduledoc """
  Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
  creator comes from the created fact alone; the holder and purpose are each the latest non-null
  value by the fact's time (stamped beside them); the state times keep their latest.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          virtual_key_id: String.t(),
          organization_id: String.t(),
          name: String.t(),
          occurred_at: DateTime.t(),
          key_created_at: DateTime.t() | nil,
          creator_user_id: String.t() | nil,
          principal_user_id: String.t() | nil,
          principal_user_id_at: DateTime.t() | nil,
          purpose: String.t() | nil,
          purpose_at: DateTime.t() | nil,
          disabled_at: DateTime.t() | nil,
          enabled_at: DateTime.t() | nil,
          revoked_at: DateTime.t() | nil
        }

  schema "NurturingVirtualKey" do
    field(:virtual_key_id, :string, primary_key: true, source: :virtualKeyId)
    field(:organization_id, :string, source: :organizationId)
    field(:name, :string)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:key_created_at, LangWatch.PrismaDateTime, source: :keyCreatedAt)
    field(:creator_user_id, :string, source: :creatorUserId)
    field(:principal_user_id, :string, source: :principalUserId)
    field(:principal_user_id_at, LangWatch.PrismaDateTime, source: :principalUserIdAt)
    field(:purpose, :string)
    field(:purpose_at, LangWatch.PrismaDateTime, source: :purposeAt)
    field(:disabled_at, LangWatch.PrismaDateTime, source: :disabledAt)
    field(:enabled_at, LangWatch.PrismaDateTime, source: :enabledAt)
    field(:revoked_at, LangWatch.PrismaDateTime, source: :revokedAt)
  end
end