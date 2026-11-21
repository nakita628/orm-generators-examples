#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum ActivityAction
{
    [PgName("ASSET_CREATED")]
    AssetCreated,

    [PgName("ASSET_NAME_CHANGED")]
    AssetNameChanged,

    [PgName("ASSET_DESCRIPTION_CHANGED")]
    AssetDescriptionChanged,

    [PgName("ASSET_CATEGORY_CHANGED")]
    AssetCategoryChanged,

    [PgName("ASSET_MODEL_CHANGED")]
    AssetModelChanged,

    [PgName("ASSET_KIT_CHANGED")]
    AssetKitChanged,

    [PgName("ASSET_LOCATION_CHANGED")]
    AssetLocationChanged,

    [PgName("ASSET_TAGS_CHANGED")]
    AssetTagsChanged,

    [PgName("ASSET_STATUS_CHANGED")]
    AssetStatusChanged,

    [PgName("ASSET_VALUATION_CHANGED")]
    AssetValuationChanged,

    [PgName("ASSET_QUANTITY_CHANGED")]
    AssetQuantityChanged,

    [PgName("ASSET_MIN_QUANTITY_CHANGED")]
    AssetMinQuantityChanged,

    [PgName("ASSET_CUSTOM_FIELD_CHANGED")]
    AssetCustomFieldChanged,

    [PgName("ASSET_PREFERRED_BARCODE_CHANGED")]
    AssetPreferredBarcodeChanged,

    [PgName("ASSET_AVAILABLE_TO_BOOK_CHANGED")]
    AssetAvailableToBookChanged,

    [PgName("ASSET_ARCHIVED")]
    AssetArchived,

    [PgName("ASSET_UNARCHIVED")]
    AssetUnarchived,

    [PgName("ASSET_DELETED")]
    AssetDeleted,

    [PgName("CUSTODY_ASSIGNED")]
    CustodyAssigned,

    [PgName("CUSTODY_RELEASED")]
    CustodyReleased,

    [PgName("BOOKING_CREATED")]
    BookingCreated,

    [PgName("BOOKING_STATUS_CHANGED")]
    BookingStatusChanged,

    [PgName("BOOKING_DATES_CHANGED")]
    BookingDatesChanged,

    [PgName("BOOKING_ASSETS_ADDED")]
    BookingAssetsAdded,

    [PgName("BOOKING_ASSETS_REMOVED")]
    BookingAssetsRemoved,

    [PgName("BOOKING_CHECKED_OUT")]
    BookingCheckedOut,

    [PgName("BOOKING_CHECKED_IN")]
    BookingCheckedIn,

    [PgName("BOOKING_PARTIAL_CHECKIN")]
    BookingPartialCheckin,

    [PgName("BOOKING_PARTIAL_CHECKOUT")]
    BookingPartialCheckout,

    [PgName("BOOKING_CANCELLED")]
    BookingCancelled,

    [PgName("BOOKING_ARCHIVED")]
    BookingArchived,

    [PgName("BOOKING_MODEL_REQUESTED")]
    BookingModelRequested,

    [PgName("BOOKING_MODEL_REQUEST_CHANGED")]
    BookingModelRequestChanged,

    [PgName("BOOKING_MODEL_REQUEST_REMOVED")]
    BookingModelRequestRemoved,

    [PgName("BOOKING_MODEL_REQUEST_FULFILLED")]
    BookingModelRequestFulfilled,

    [PgName("AUDIT_CREATED")]
    AuditCreated,

    [PgName("AUDIT_STARTED")]
    AuditStarted,

    [PgName("AUDIT_ASSETS_ADDED")]
    AuditAssetsAdded,

    [PgName("AUDIT_ASSETS_REMOVED")]
    AuditAssetsRemoved,

    [PgName("AUDIT_ASSET_SCANNED")]
    AuditAssetScanned,

    [PgName("AUDIT_ASSET_SCAN_REMOVED")]
    AuditAssetScanRemoved,

    [PgName("AUDIT_DUE_DATE_CHANGED")]
    AuditDueDateChanged,

    [PgName("AUDIT_ASSIGNEE_ADDED")]
    AuditAssigneeAdded,

    [PgName("AUDIT_ASSIGNEE_REMOVED")]
    AuditAssigneeRemoved,

    [PgName("AUDIT_UPDATED")]
    AuditUpdated,

    [PgName("AUDIT_COMPLETED")]
    AuditCompleted,

    [PgName("AUDIT_CANCELLED")]
    AuditCancelled,

    [PgName("AUDIT_ARCHIVED")]
    AuditArchived,

    [PgName("LOCATION_CREATED")]
    LocationCreated,

    [PgName("LOCATION_UPDATED")]
    LocationUpdated,

    [PgName("KIT_CREATED")]
    KitCreated,

    [PgName("KIT_UPDATED")]
    KitUpdated,

    [PgName("ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED")]
    OrganizationQrIdDisplayPreferenceChanged,
}
