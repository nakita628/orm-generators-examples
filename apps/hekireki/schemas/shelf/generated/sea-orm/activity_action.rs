use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "ActivityAction")]
pub enum ActivityAction {
    #[sea_orm(string_value = "ASSET_CREATED")]
    AssetCreated,
    #[sea_orm(string_value = "ASSET_NAME_CHANGED")]
    AssetNameChanged,
    #[sea_orm(string_value = "ASSET_DESCRIPTION_CHANGED")]
    AssetDescriptionChanged,
    #[sea_orm(string_value = "ASSET_CATEGORY_CHANGED")]
    AssetCategoryChanged,
    #[sea_orm(string_value = "ASSET_MODEL_CHANGED")]
    AssetModelChanged,
    #[sea_orm(string_value = "ASSET_KIT_CHANGED")]
    AssetKitChanged,
    #[sea_orm(string_value = "ASSET_LOCATION_CHANGED")]
    AssetLocationChanged,
    #[sea_orm(string_value = "ASSET_TAGS_CHANGED")]
    AssetTagsChanged,
    #[sea_orm(string_value = "ASSET_STATUS_CHANGED")]
    AssetStatusChanged,
    #[sea_orm(string_value = "ASSET_VALUATION_CHANGED")]
    AssetValuationChanged,
    #[sea_orm(string_value = "ASSET_QUANTITY_CHANGED")]
    AssetQuantityChanged,
    #[sea_orm(string_value = "ASSET_MIN_QUANTITY_CHANGED")]
    AssetMinQuantityChanged,
    #[sea_orm(string_value = "ASSET_CUSTOM_FIELD_CHANGED")]
    AssetCustomFieldChanged,
    #[sea_orm(string_value = "ASSET_PREFERRED_BARCODE_CHANGED")]
    AssetPreferredBarcodeChanged,
    #[sea_orm(string_value = "ASSET_AVAILABLE_TO_BOOK_CHANGED")]
    AssetAvailableToBookChanged,
    #[sea_orm(string_value = "ASSET_ARCHIVED")]
    AssetArchived,
    #[sea_orm(string_value = "ASSET_UNARCHIVED")]
    AssetUnarchived,
    #[sea_orm(string_value = "ASSET_DELETED")]
    AssetDeleted,
    #[sea_orm(string_value = "CUSTODY_ASSIGNED")]
    CustodyAssigned,
    #[sea_orm(string_value = "CUSTODY_RELEASED")]
    CustodyReleased,
    #[sea_orm(string_value = "BOOKING_CREATED")]
    BookingCreated,
    #[sea_orm(string_value = "BOOKING_STATUS_CHANGED")]
    BookingStatusChanged,
    #[sea_orm(string_value = "BOOKING_DATES_CHANGED")]
    BookingDatesChanged,
    #[sea_orm(string_value = "BOOKING_ASSETS_ADDED")]
    BookingAssetsAdded,
    #[sea_orm(string_value = "BOOKING_ASSETS_REMOVED")]
    BookingAssetsRemoved,
    #[sea_orm(string_value = "BOOKING_CHECKED_OUT")]
    BookingCheckedOut,
    #[sea_orm(string_value = "BOOKING_CHECKED_IN")]
    BookingCheckedIn,
    #[sea_orm(string_value = "BOOKING_PARTIAL_CHECKIN")]
    BookingPartialCheckin,
    #[sea_orm(string_value = "BOOKING_PARTIAL_CHECKOUT")]
    BookingPartialCheckout,
    #[sea_orm(string_value = "BOOKING_CANCELLED")]
    BookingCancelled,
    #[sea_orm(string_value = "BOOKING_ARCHIVED")]
    BookingArchived,
    #[sea_orm(string_value = "BOOKING_MODEL_REQUESTED")]
    BookingModelRequested,
    #[sea_orm(string_value = "BOOKING_MODEL_REQUEST_CHANGED")]
    BookingModelRequestChanged,
    #[sea_orm(string_value = "BOOKING_MODEL_REQUEST_REMOVED")]
    BookingModelRequestRemoved,
    #[sea_orm(string_value = "BOOKING_MODEL_REQUEST_FULFILLED")]
    BookingModelRequestFulfilled,
    #[sea_orm(string_value = "AUDIT_CREATED")]
    AuditCreated,
    #[sea_orm(string_value = "AUDIT_STARTED")]
    AuditStarted,
    #[sea_orm(string_value = "AUDIT_ASSETS_ADDED")]
    AuditAssetsAdded,
    #[sea_orm(string_value = "AUDIT_ASSETS_REMOVED")]
    AuditAssetsRemoved,
    #[sea_orm(string_value = "AUDIT_ASSET_SCANNED")]
    AuditAssetScanned,
    #[sea_orm(string_value = "AUDIT_ASSET_SCAN_REMOVED")]
    AuditAssetScanRemoved,
    #[sea_orm(string_value = "AUDIT_DUE_DATE_CHANGED")]
    AuditDueDateChanged,
    #[sea_orm(string_value = "AUDIT_ASSIGNEE_ADDED")]
    AuditAssigneeAdded,
    #[sea_orm(string_value = "AUDIT_ASSIGNEE_REMOVED")]
    AuditAssigneeRemoved,
    #[sea_orm(string_value = "AUDIT_UPDATED")]
    AuditUpdated,
    #[sea_orm(string_value = "AUDIT_COMPLETED")]
    AuditCompleted,
    #[sea_orm(string_value = "AUDIT_CANCELLED")]
    AuditCancelled,
    #[sea_orm(string_value = "AUDIT_ARCHIVED")]
    AuditArchived,
    #[sea_orm(string_value = "LOCATION_CREATED")]
    LocationCreated,
    #[sea_orm(string_value = "LOCATION_UPDATED")]
    LocationUpdated,
    #[sea_orm(string_value = "KIT_CREATED")]
    KitCreated,
    #[sea_orm(string_value = "KIT_UPDATED")]
    KitUpdated,
    #[sea_orm(string_value = "ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED")]
    OrganizationQrIdDisplayPreferenceChanged,
}
