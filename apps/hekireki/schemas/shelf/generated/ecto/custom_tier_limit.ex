defmodule Shelf.CustomTierLimit do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: String.t(),
          can_import_assets: boolean(),
          can_export_assets: boolean(),
          can_import_nrm: boolean(),
          can_hide_shelf_branding: boolean(),
          can_use_reports: boolean(),
          can_use_recurring_reminders: boolean(),
          can_use_push_notifications: boolean(),
          max_custom_fields: integer(),
          max_organizations: integer(),
          is_enterprise: boolean(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          user: Shelf.User.t() | nil
        }

  schema "CustomTierLimit" do
    field(:id, :string, primary_key: true)
    field(:can_import_assets, :boolean, default: true, source: :canImportAssets)
    field(:can_export_assets, :boolean, default: true, source: :canExportAssets)
    field(:can_import_nrm, :boolean, default: true, source: :canImportNRM)
    field(:can_hide_shelf_branding, :boolean, default: true, source: :canHideShelfBranding)
    field(:can_use_reports, :boolean, default: true, source: :canUseReports)
    field(:can_use_recurring_reminders, :boolean, default: true, source: :canUseRecurringReminders)
    field(:can_use_push_notifications, :boolean, default: true, source: :canUsePushNotifications)
    field(:max_custom_fields, :integer, default: 1000, source: :maxCustomFields)
    field(:max_organizations, :integer, default: 1, source: :maxOrganizations)
    field(:is_enterprise, :boolean, default: false, source: :isEnterprise)
    field(:user_id, :string, source: :userId)
    belongs_to(:user, Shelf.User, foreign_key: :user_id, define_field: false, type: :string)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end