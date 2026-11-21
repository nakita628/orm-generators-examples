defmodule Shelf.TierLimit do
  use Ecto.Schema

  @primary_key false
  @timestamps_opts [type: Shelf.PrismaDateTime, autogenerate: {Shelf.PrismaDateTime, :autogenerate, []}]

  @type t :: %__MODULE__{
          id: atom(),
          can_import_assets: boolean(),
          can_export_assets: boolean(),
          can_import_nrm: boolean(),
          can_hide_shelf_branding: boolean(),
          can_use_reports: boolean(),
          can_use_recurring_reminders: boolean(),
          can_use_push_notifications: boolean(),
          max_custom_fields: integer(),
          max_organizations: integer(),
          inserted_at: DateTime.t(),
          updated_at: DateTime.t(),
          tier: Shelf.Tier.t() | nil
        }

  schema "TierLimit" do
    field(:id, Ecto.Enum, values: [:free, :tier_1, :tier_2, :custom])
    field(:can_import_assets, :boolean, default: false, source: :canImportAssets)
    field(:can_export_assets, :boolean, default: false, source: :canExportAssets)
    field(:can_import_nrm, :boolean, default: false, source: :canImportNRM)
    field(:can_hide_shelf_branding, :boolean, default: false, source: :canHideShelfBranding)
    field(:can_use_reports, :boolean, default: false, source: :canUseReports)
    field(:can_use_recurring_reminders, :boolean, default: false, source: :canUseRecurringReminders)
    field(:can_use_push_notifications, :boolean, default: false, source: :canUsePushNotifications)
    field(:max_custom_fields, :integer, default: 0, source: :maxCustomFields)
    field(:max_organizations, :integer, default: 1, source: :maxOrganizations)
    has_one(:tier, Shelf.Tier, foreign_key: :tier_limit_id, references: :id)
    timestamps(inserted_at_source: :createdAt, updated_at_source: :updatedAt)
  end
end