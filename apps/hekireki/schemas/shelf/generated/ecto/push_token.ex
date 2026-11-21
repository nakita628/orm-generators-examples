defmodule Shelf.PushToken do
  @moduledoc """
  One Expo push token per installed companion app. A token identifies a
  device, not a person: when someone else signs in on the same phone, the
  row moves to them. Pushes go only to rows where `disabledAt` is null.
  @see apps/webapp/app/modules/push/service.server.ts
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          token: String.t(),
          platform: atom(),
          device_name: String.t() | nil,
          app_version: String.t() | nil,
          last_seen_at: DateTime.t(),
          disabled_at: DateTime.t() | nil,
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "PushToken" do
    field(:id, :string, primary_key: true)
    field(:token, :string)
    field(:platform, Ecto.Enum, values: [:IOS, :ANDROID])
    field(:device_name, :string, source: :deviceName)
    field(:app_version, :string, source: :appVersion)
    field(:last_seen_at, Shelf.PrismaDateTime, source: :lastSeenAt)
    field(:disabled_at, Shelf.PrismaDateTime, source: :disabledAt)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end