defmodule Shelf.ActivityEvent do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          occurred_at: DateTime.t(),
          actor_user_id: String.t() | nil,
          actor_snapshot: map() | nil,
          action: atom(),
          entity_type: atom(),
          entity_id: String.t(),
          asset_id: String.t() | nil,
          booking_id: String.t() | nil,
          audit_session_id: String.t() | nil,
          audit_asset_id: String.t() | nil,
          kit_id: String.t() | nil,
          location_id: String.t() | nil,
          team_member_id: String.t() | nil,
          target_user_id: String.t() | nil,
          field: String.t() | nil,
          from_value: map() | nil,
          to_value: map() | nil,
          meta: map() | nil,
          organization: Shelf.Organization.t() | nil
        }

  schema "ActivityEvent" do
    field(:id, :string, primary_key: true)
    field(:occurred_at, Shelf.PrismaDateTime, autogenerate: true, source: :occurredAt)
    field(:actor_user_id, :string, source: :actorUserId)
    field(:actor_snapshot, :map, source: :actorSnapshot)
    field(:action, Ecto.Enum, values: [:ASSET_CREATED, :ASSET_NAME_CHANGED, :ASSET_DESCRIPTION_CHANGED, :ASSET_CATEGORY_CHANGED, :ASSET_MODEL_CHANGED, :ASSET_KIT_CHANGED, :ASSET_LOCATION_CHANGED, :ASSET_TAGS_CHANGED, :ASSET_STATUS_CHANGED, :ASSET_VALUATION_CHANGED, :ASSET_QUANTITY_CHANGED, :ASSET_MIN_QUANTITY_CHANGED, :ASSET_CUSTOM_FIELD_CHANGED, :ASSET_PREFERRED_BARCODE_CHANGED, :ASSET_AVAILABLE_TO_BOOK_CHANGED, :ASSET_ARCHIVED, :ASSET_UNARCHIVED, :ASSET_DELETED, :CUSTODY_ASSIGNED, :CUSTODY_RELEASED, :BOOKING_CREATED, :BOOKING_STATUS_CHANGED, :BOOKING_DATES_CHANGED, :BOOKING_ASSETS_ADDED, :BOOKING_ASSETS_REMOVED, :BOOKING_CHECKED_OUT, :BOOKING_CHECKED_IN, :BOOKING_PARTIAL_CHECKIN, :BOOKING_PARTIAL_CHECKOUT, :BOOKING_CANCELLED, :BOOKING_ARCHIVED, :BOOKING_MODEL_REQUESTED, :BOOKING_MODEL_REQUEST_CHANGED, :BOOKING_MODEL_REQUEST_REMOVED, :BOOKING_MODEL_REQUEST_FULFILLED, :AUDIT_CREATED, :AUDIT_STARTED, :AUDIT_ASSETS_ADDED, :AUDIT_ASSETS_REMOVED, :AUDIT_ASSET_SCANNED, :AUDIT_ASSET_SCAN_REMOVED, :AUDIT_DUE_DATE_CHANGED, :AUDIT_ASSIGNEE_ADDED, :AUDIT_ASSIGNEE_REMOVED, :AUDIT_UPDATED, :AUDIT_COMPLETED, :AUDIT_CANCELLED, :AUDIT_ARCHIVED, :LOCATION_CREATED, :LOCATION_UPDATED, :KIT_CREATED, :KIT_UPDATED, :ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED])
    field(:entity_type, Ecto.Enum, values: [:ASSET, :BOOKING, :AUDIT, :KIT, :LOCATION, :TEAM_MEMBER, :CUSTODY, :USER, :ORGANIZATION], source: :entityType)
    field(:entity_id, :string, source: :entityId)
    field(:asset_id, :string, source: :assetId)
    field(:booking_id, :string, source: :bookingId)
    field(:audit_session_id, :string, source: :auditSessionId)
    field(:audit_asset_id, :string, source: :auditAssetId)
    field(:kit_id, :string, source: :kitId)
    field(:location_id, :string, source: :locationId)
    field(:team_member_id, :string, source: :teamMemberId)
    field(:target_user_id, :string, source: :targetUserId)
    field(:field, :string)
    field(:from_value, :map, source: :fromValue)
    field(:to_value, :map, source: :toValue)
    field(:meta, :map)
    field(:organization_id, :string, source: :organizationId)
    belongs_to(:organization, Shelf.Organization, foreign_key: :organization_id, define_field: false, type: :string)
  end
end