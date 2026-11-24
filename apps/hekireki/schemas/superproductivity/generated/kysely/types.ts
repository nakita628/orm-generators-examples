import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

export interface User {
  id: Generated<number>
  email: string
  password_hash: string | null
  is_verified: Generated<number>
  verification_token: string | null
  verification_token_expires_at: ColumnType<
    string,
    bigint | number | string,
    bigint | number | string
  > | null
  verification_resend_count: Generated<number>
  reset_password_token: string | null
  reset_password_token_expires_at: ColumnType<
    string,
    bigint | number | string,
    bigint | number | string
  > | null
  passkey_recovery_token: string | null
  passkey_recovery_token_expires_at: ColumnType<
    string,
    bigint | number | string,
    bigint | number | string
  > | null
  login_token: string | null
  login_token_expires_at: ColumnType<
    string,
    bigint | number | string,
    bigint | number | string
  > | null
  failed_login_attempts: Generated<number>
  locked_until: ColumnType<string, bigint | number | string, bigint | number | string> | null
  token_version: Generated<number>
  terms_accepted_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
  created_at: Generated<Timestamp>
  storage_quota_bytes: Generated<
    ColumnType<string, bigint | number | string, bigint | number | string>
  >
  storage_used_bytes: Generated<
    ColumnType<string, bigint | number | string, bigint | number | string>
  >
}

export interface Passkey {
  id: string
  credential_id: Buffer
  public_key: Buffer
  counter: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  transports: string | null
  created_at: Generated<Timestamp>
  last_used_at: Timestamp | null
  user_id: number
}

export interface PendingPasskeyRegistration {
  id: string
  verification_token: string
  verification_token_expires_at: ColumnType<
    string,
    bigint | number | string,
    bigint | number | string
  >
  credential_id: Buffer
  public_key: Buffer
  counter: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  transports: string | null
  created_at: Generated<Timestamp>
  user_id: number
}

export interface Operation {
  id: string
  user_id: number
  client_id: string
  server_seq: number
  action_type: string
  op_type: string
  entity_type: string
  entity_id: string | null
  entity_ids: Generated<string[]>
  payload: unknown
  payload_bytes: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  vector_clock: unknown
  schema_version: number
  client_timestamp: ColumnType<string, bigint | number | string, bigint | number | string>
  received_at: ColumnType<string, bigint | number | string, bigint | number | string>
  is_payload_encrypted: Generated<boolean>
  sync_import_reason: string | null
  repair_base_server_seq: number | null
}

export interface UserSyncState {
  user_id: number
  last_seq: Generated<number>
  last_snapshot_seq: number | null
  snapshot_data: Buffer | null
  snapshot_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
  snapshot_schema_version: Generated<number | null>
  latest_full_state_seq: number | null
  latest_full_state_vector_clock: unknown
  latest_state_replacement_seq: number | null
}

export interface SyncDevice {
  client_id: string
  user_id: number
  device_name: string | null
  user_agent: string | null
  /**
   * Bare semver the client last reported on download (`18.22.0`), NULL until
   * it does. Read by the checkpoint gate (#9962); never shown to users.
   */
  app_version: string | null
  last_seen_at: ColumnType<string, bigint | number | string, bigint | number | string>
  /** Reserved: written once at device registration, never read or updated. */
  last_acked_seq: Generated<number>
  created_at: ColumnType<string, bigint | number | string, bigint | number | string>
}

/**
 * Public half of the user's capture key. The row outlives switching the
 * feature off: `disabledAt` stops a stale device re-registering that key.
 */
export interface CaptureKey {
  user_id: number
  /** Base64 SPKI DER of an ECDH P-256 public key. */
  public_key: string
  fingerprint: string
  created_at: ColumnType<string, bigint | number | string, bigint | number | string>
  disabled_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
}

/** Capture tokens and OAuth grants store only hashes of bearer credentials. */
export interface CaptureGrant {
  id: string
  user_id: number
  kind: string
  client_name: string
  secret_hash: string
  created_at: ColumnType<string, bigint | number | string, bigint | number | string>
  expires_at: ColumnType<string, bigint | number | string, bigint | number | string>
  last_used_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
  revoked_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
  client_id: string | null
  resource: string | null
  refresh_hash: string | null
  prev_refresh_hash: string | null
  rotated_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
  refresh_expires_at: ColumnType<string, bigint | number | string, bigint | number | string> | null
}

/**
 * One sealed capture waiting for a device. No relation to its grant: items
 * outlive a grant that is revoked without discarding them.
 */
export interface CaptureItem {
  id: string
  user_id: number
  grant_id: string
  ciphertext: string
  /** Task id the claiming device must use; replaced when another device takes over. */
  task_id: string
  created_at: ColumnType<string, bigint | number | string, bigint | number | string>
  lease_device_id: string | null
  lease_until: ColumnType<string, bigint | number | string, bigint | number | string> | null
}

export interface CaptureConnectCode {
  user_id: number
  code_hash: string
  token_version: number
  expires_at: ColumnType<string, bigint | number | string, bigint | number | string>
}

export interface CaptureAuth {
  id: string
  user_id: number | null
  client_id: string
  client_name: string
  redirect_uri: string
  code_challenge: string
  resource: string
  state: string
  scope: string
  code_hash: string | null
  token_version: number | null
  expires_at: ColumnType<string, bigint | number | string, bigint | number | string>
}

export interface ConnectRequest {
  id: string
  poll_secret_hash: string
  match_code: string
  device_name: string
  platform: string
  status: Generated<string>
  user_id: number | null
  token_version: number | null
  expires_at: ColumnType<string, bigint | number | string, bigint | number | string>
}

export interface DB {
  users: User
  passkeys: Passkey
  pending_passkey_registrations: PendingPasskeyRegistration
  operations: Operation
  user_sync_state: UserSyncState
  sync_devices: SyncDevice
  capture_keys: CaptureKey
  capture_grants: CaptureGrant
  capture_items: CaptureItem
  capture_connect_codes: CaptureConnectCode
  capture_auth: CaptureAuth
  connect_requests: ConnectRequest
}
