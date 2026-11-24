defmodule Umami.ErrorEvent do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          website_id: Ecto.UUID.t(),
          id: Ecto.UUID.t(),
          issue_id: Ecto.UUID.t(),
          session_id: Ecto.UUID.t(),
          visit_id: Ecto.UUID.t(),
          created_at: DateTime.t(),
          received_at: DateTime.t(),
          name: String.t(),
          message: String.t(),
          stack: String.t(),
          frames: map() | list() | String.t() | number() | boolean(),
          tags: map() | list() | String.t() | number() | boolean(),
          handled: boolean(),
          url_path: String.t(),
          release: String.t(),
          environment: String.t(),
          browser: String.t(),
          os: String.t(),
          device: String.t()
        }

  schema "error_event" do
    field(:website_id, Ecto.UUID, primary_key: true)
    field(:id, Ecto.UUID, primary_key: true, source: :event_id)
    field(:issue_id, Ecto.UUID)
    field(:session_id, Ecto.UUID)
    field(:visit_id, Ecto.UUID)
    field(:created_at, Umami.PrismaDateTime)
    field(:received_at, Umami.PrismaDateTime, autogenerate: true)
    field(:name, :string)
    field(:message, :string)
    field(:stack, :string)
    field(:frames, Umami.PrismaJson)
    field(:tags, Umami.PrismaJson)
    field(:handled, :boolean)
    field(:url_path, :string)
    field(:release, :string)
    field(:environment, :string)
    field(:browser, :string)
    field(:os, :string)
    field(:device, :string)
  end
end