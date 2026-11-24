defmodule Shelf.NotificationPreference do
  @moduledoc """
  A person's push notification switches. A missing row means every switch
  is on, so no backfill is needed. The three category columns are named
  after `PushCategory` in apps/webapp/app/modules/push/types.ts.
  """

  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          user_id: String.t(),
          push_enabled: boolean(),
          bookings: boolean(),
          audits: boolean(),
          stock_and_reminders: boolean(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "NotificationPreference" do
    field(:push_enabled, :boolean, default: true, source: :pushEnabled)
    field(:bookings, :boolean, default: true)
    field(:audits, :boolean, default: true)
    field(:stock_and_reminders, :boolean, default: true, source: :stockAndReminders)
    field(:user_id, :string, primary_key: true, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at: false, updated_at_source: :updatedAt)
  end
end