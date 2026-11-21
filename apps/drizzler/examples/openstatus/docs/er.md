```mermaid
erDiagram
    workspace |o--o{ integration : "(id) - (workspace_id)"
    workspace ||--o{ workspace_sso_domain : "(id) - (workspace_id)"
    user ||--o{ users_to_workspaces : "(id) - (user_id)"
    workspace ||--o{ users_to_workspaces : "(id) - (workspace_id)"
    user ||--o{ account : "(id) - (user_id)"
    user ||--o{ session : "(id) - (user_id)"
    workspace |o--o{ status_report : "(id) - (workspace_id)"
    page |o--o{ status_report : "(id) - (page_id)"
    user |o--o{ status_report : "(id) - (created_by)"
    user |o--o{ status_report : "(id) - (updated_by)"
    status_report ||--o{ status_report_update : "(id) - (status_report_id)"
    user |o--o{ status_report_update : "(id) - (created_by)"
    user |o--o{ status_report_update : "(id) - (updated_by)"
    workspace ||--o{ page : "(id) - (workspace_id)"
    workspace |o--o{ monitor : "(id) - (workspace_id)"
    page ||--o| page_subscriber : "(id) - (page_id)"
    page_subscriber ||--o{ page_subscriber_to_page_component : "(id) - (page_subscriber_id)"
    page_component ||--o{ page_subscriber_to_page_component : "(id) - (page_component_id)"
    workspace |o--o{ notification : "(id) - (workspace_id)"
    monitor |o--o{ notification_trigger : "(id) - (monitor_id)"
    notification |o--o{ notification_trigger : "(id) - (notification_id)"
    monitor ||--o{ notifications_to_monitors : "(id) - (monitor_id)"
    notification ||--o{ notifications_to_monitors : "(id) - (notification_id)"
    monitor ||--o{ notification_outbox : "(id) - (monitor_id)"
    workspace |o--o{ notification_outbox : "(id) - (workspace_id)"
    notification ||--o{ notification_outbox : "(id) - (notification_id)"
    monitor_incident |o--o{ notification_outbox : "(id) - (incident_id)"
    monitor ||--o{ notification_dead_letter : "(id) - (monitor_id)"
    workspace |o--o{ notification_dead_letter : "(id) - (workspace_id)"
    monitor ||--o{ monitor_status : "(id) - (monitor_id)"
    monitor |o--o| monitor_incident : "(id) - (monitor_id)"
    workspace |o--o{ monitor_incident : "(id) - (workspace_id)"
    user |o--o{ monitor_incident : "(id) - (acknowledged_by)"
    user |o--o{ monitor_incident : "(id) - (resolved_by)"
    workspace ||--o{ monitor_tag : "(id) - (workspace_id)"
    monitor ||--o{ monitor_tag_to_monitor : "(id) - (monitor_id)"
    monitor_tag ||--o{ monitor_tag_to_monitor : "(id) - (monitor_tag_id)"
    workspace |o--o{ application : "(id) - (workspace_id)"
    workspace |o--o{ maintenance : "(id) - (workspace_id)"
    page |o--o{ maintenance : "(id) - (page_id)"
    user |o--o{ maintenance : "(id) - (created_by)"
    user |o--o{ maintenance : "(id) - (updated_by)"
    maintenance ||--o{ maintenance_update : "(id) - (maintenance_id)"
    user |o--o{ maintenance_update : "(id) - (created_by)"
    user |o--o{ maintenance_update : "(id) - (updated_by)"
    workspace |o--o{ check : "(id) - (workspace_id)"
    workspace |o--o{ monitor_run : "(id) - (workspace_id)"
    monitor |o--o{ monitor_run : "(id) - (monitor_id)"
    monitor ||--o{ private_location_monitor_status : "(id) - (monitor_id)"
    private_location ||--o{ private_location_monitor_status : "(id) - (private_location_id)"
    workspace |o--o{ private_location : "(id) - (workspace_id)"
    private_location |o--o{ private_location_to_monitor : "(id) - (private_location_id)"
    monitor |o--o{ private_location_to_monitor : "(id) - (monitor_id)"
    workspace ||--o{ monitor_group : "(id) - (workspace_id)"
    page ||--o{ monitor_group : "(id) - (page_id)"
    viewer ||--o{ viewer_session : "(id) - (user_id)"
    viewer ||--o{ viewer_accounts : "(id) - (user_id)"
    workspace ||--o{ api_key : "(id) - (workspace_id)"
    user ||--o{ api_key : "(id) - (created_by_id)"
    workspace ||--o{ page_component : "(id) - (workspace_id)"
    page ||--o{ page_component : "(id) - (page_id)"
    monitor |o--o{ page_component : "(id) - (monitor_id)"
    page_component_groups |o--o{ page_component : "(id) - (group_id)"
    maintenance ||--o{ maintenance_to_page_component : "(id) - (maintenance_id)"
    page_component ||--o{ maintenance_to_page_component : "(id) - (page_component_id)"
    status_report ||--o{ status_report_to_page_component : "(id) - (status_report_id)"
    page_component ||--o{ status_report_to_page_component : "(id) - (page_component_id)"
    status_report_update ||--o{ status_report_update_to_page_component : "(id) - (status_report_update_id)"
    page_component ||--o{ status_report_update_to_page_component : "(id) - (page_component_id)"
    workspace ||--o{ page_component_groups : "(id) - (workspace_id)"
    page ||--o{ page_component_groups : "(id) - (page_id)"
    workspace ||--o{ feedback : "(id) - (workspace_id)"
    user ||--o{ feedback : "(id) - (user_id)"
    external_service ||--o{ external_service_component : "(id) - (external_service_id)"
    external_service ||--o{ external_service_incident : "(id) - (external_service_id)"
    external_service ||--o{ external_service_report : "(id) - (external_service_id)"
    external_service_component |o--o{ external_service_report : "(id) - (external_service_component_id)"
    workspace ||--o{ chat_session : "(id) - (workspace_id)"
    user ||--o{ chat_session : "(id) - (user_id)"
    workspace ||--o{ frozen_monitor_uptime : "(id) - (workspace_id)"
    monitor ||--o{ frozen_monitor_uptime : "(id) - (monitor_id)"
    monitor ||--o{ monitor_transition : "(id) - (monitor_id)"
    oauth_client ||--o{ oauth_session : "(client_id) - (client_id)"
    oauth_client ||--o{ oauth_authorization_code : "(client_id) - (client_id)"
    user ||--o{ oauth_authorization_code : "(id) - (user_id)"
    workspace ||--o{ oauth_authorization_code : "(id) - (workspace_id)"
    oauth_grant |o--o{ oauth_authorization_code : "(id) - (grant_id)"
    oauth_client ||--o{ oauth_grant : "(client_id) - (client_id)"
    user ||--o{ oauth_grant : "(id) - (user_id)"
    workspace ||--o{ oauth_grant : "(id) - (workspace_id)"
    workspace ||--o{ incident : "(id) - (workspace_id)"
    user |o--o{ incident : "(id) - (commander_id)"
    user |o--o{ incident : "(id) - (declared_by)"
    user |o--o{ incident : "(id) - (resolved_by)"
    status_report |o--o| incident : "(id) - (status_report_id)"
    incident ||--o{ incident_event : "(id) - (incident_id)"
    user |o--o{ incident_event : "(id) - (created_by)"
    incident ||--o| incident_postmortem : "(id) - (incident_id)"
    user |o--o{ incident_postmortem : "(id) - (created_by)"
    user |o--o{ incident_postmortem : "(id) - (updated_by)"
    user |o--o{ incident_postmortem : "(id) - (approved_by)"
    workspace ||--o{ slack_user : "(id) - (workspace_id)"
    user ||--o{ slack_user : "(id) - (user_id)"
    monitor ||--o| monitor_heartbeat_state : "(id) - (monitor_id)"
    workspace ||..o{ invitation : "(id) - (workspace_id)"
    integration {
        integer id PK
        text name
        integer workspace_id FK
        text credential "Not used yet but we might need to get store something for the integration  webhook url and or secret"
        text external_id
        integer created_at
        integer updated_at
        text data
    }
    workspace {
        integer id PK
        text slug
        text name
        text stripe_id
        text subscription_id
        text plan
        integer ends_at
        integer paid_until
        integer trial_ends_at
        integer trialed_at
        text limits
        text workos_organization_id
        integer sso_enabled "Our own permission switch, never mirrored from WorkOS: the downgrade cascade clears it, and a later connection webhook must not resurrect it."
        integer created_at
        integer updated_at
        text dsn
    }
    workspace_sso_domain {
        integer id PK
        integer workspace_id FK
        text domain "Globally unique: two workspaces claiming the same domain would make the login-time lookup ambiguous, which is an account-takeover vector."
        integer verified_at "Only ever written from the `organization_domain.verified` webhook."
        integer created_at
    }
    user {
        integer id PK
        text tenant_id "clerk fields"
        text first_name
        text last_name
        text photo_url
        text name "next-auth fields"
        text email
        integer emailVerified
        integer created_at
        integer updated_at
        integer deleted_at
    }
    users_to_workspaces {
        integer user_id PK, FK
        integer workspace_id PK, FK
        text role
        integer created_at
    }
    account {
        integer user_id FK
        text type
        text provider PK
        text provider_account_id PK
        text refresh_token
        text access_token
        integer expires_at
        text token_type
        text scope
        text id_token
        text session_state
    }
    session {
        text session_token PK
        integer user_id FK
        integer expires
    }
    verification_token {
        text identifier PK
        text token PK
        integer expires
    }
    status_report {
        integer id PK
        text status
        text title
        integer workspace_id FK
        integer page_id FK
        integer created_by FK
        integer updated_by FK
        integer created_at
        integer updated_at
    }
    status_report_update {
        integer id PK
        text status
        integer date
        text message
        integer status_report_id FK
        integer created_by FK
        integer updated_by FK
        integer created_at
        integer updated_at
    }
    page {
        integer id PK
        integer workspace_id FK
        text title
        text description
        text icon
        text slug
        text custom_domain
        integer published
        text force_theme
        text custom_theme "Per-mode CSS variable overrides ({ light, dark }) merged over the configured theme vars. Gated by the 'custom-theme' plan limit."
        text password "Password protecting the status page - no specific restriction on password"
        integer password_protected
        text access_type
        text auth_email_domains
        text allowed_ip_ranges
        text homepage_url "links and urls"
        text contact_url
        text default_locale
        text locales
        text default_timezone "IANA zone the public page renders its bars and timestamps in"
        integer legacy_page
        text configuration
        integer allow_index
        integer show_monitor_values "Displays the total and failed request numbers for each monitor TODO: remove this column - we moved into configuration"
        integer created_at
        integer updated_at
    }
    monitor {
        integer id PK
        text job_type
        text periodicity
        text status
        integer active
        text regions
        text url
        text name
        text external_name
        text description
        text headers
        text body
        text method
        integer workspace_id FK
        integer timeout "Custom timeout for this monitor"
        integer degraded_after "Threshold for the monitor to be considered degraded"
        text assertions
        text otel_endpoint
        text otel_headers
        integer public
        integer retry
        integer follow_redirects
        text grpc_service
        text grpc_tls
        text heartbeat_token
        integer heartbeat_period "seconds"
        integer heartbeat_grace
        text heartbeat_cron
        text heartbeat_tz
        integer created_at
        integer updated_at
        integer deleted_at
    }
    page_subscriber {
        integer id PK
        text email
        integer page_id FK
        text channel_type "Added: channel type discriminator"
        text webhook_url "Added: webhook-specific fields (null for email channel)"
        text channel_config
        text slack_channel_id "Slack channel id (null for non-slack channels). teamId + channelName live in channelConfig; the bot token is resolved from the integration row at send time, never persisted here."
        text source "How the subscription was created. Vendor-added rows skip verification and are editable from the dashboard."
        text name "Optional human-readable label (e.g. 'Supabase #incidents'). When set, rendered in place of the raw destination in UI."
        text token
        integer accepted_at
        integer expires_at
        integer unsubscribed_at
        integer created_at
        integer updated_at
    }
    page_subscriber_to_page_component {
        integer page_subscriber_id PK, FK
        integer page_component_id PK, FK
        integer created_at
    }
    notification {
        integer id PK
        text name
        text provider
        text data
        integer workspace_id FK
        integer created_at
        integer updated_at
    }
    notification_trigger {
        integer id PK
        integer monitor_id FK
        integer notification_id FK
        integer cron_timestamp
    }
    notifications_to_monitors {
        integer monitor_id PK, FK
        integer notification_id PK, FK
        integer created_at
    }
    notification_outbox {
        integer id PK
        text dedup_key
        integer monitor_id FK
        integer workspace_id FK
        integer notification_id FK
        text provider
        text event_type
        text from_status
        text to_status
        integer cron_timestamp
        integer incident_id FK
        text payload
        text delivery_status
        text outcome
        integer attempts
        integer next_attempt_at
        integer deadline_at
        text locked_by
        integer locked_until
        integer delivered_at
        text last_error
        integer created_at
    }
    notification_dead_letter {
        integer id PK
        integer outbox_id
        text dedup_key
        integer monitor_id FK
        integer workspace_id FK
        integer notification_id
        text provider
        text event_type
        text from_status
        text to_status
        integer cron_timestamp
        integer incident_id
        text payload
        integer attempts
        text final_error
        integer died_at
    }
    monitor_status {
        integer monitor_id PK, FK
        text region PK
        text status
        integer cron_timestamp
        integer created_at
        integer updated_at
    }
    invitation {
        integer id PK
        text email
        text role
        integer workspace_id FK
        text token
        integer expires_at
        integer created_at
        integer accepted_at
    }
    monitor_incident {
        integer id PK
        text title
        text summary
        text status
        integer monitor_id FK "Service affected by incident"
        integer workspace_id FK "Workspace where the incident happened"
        integer started_at "Data related to incident timeline"
        integer acknowledged_at "Who has acknowledged the incident"
        integer acknowledged_by FK
        integer resolved_at "Who has resolved it"
        integer resolved_by FK
        text incident_screenshot_url
        text recovery_screenshot_url
        integer auto_resolved "If the incident was auto resolved"
        integer created_at
        integer updated_at
    }
    monitor_tag {
        integer id PK
        integer workspace_id FK
        text name
        text color
        integer created_at
        integer updated_at
    }
    monitor_tag_to_monitor {
        integer monitor_id PK, FK
        integer monitor_tag_id PK, FK
        integer created_at
    }
    application {
        integer id PK
        text name
        text dsn
        integer workspace_id FK
        integer created_at
        integer updated_at
    }
    maintenance {
        integer id PK
        text title
        text message "deprecated: mirror of the first `maintenance_update` (backfilled in migration 0092); dropped in a follow-up"
        integer from
        integer to
        integer workspace_id FK
        integer page_id FK
        integer created_by FK
        integer updated_by FK
        integer created_at
        integer updated_at
    }
    maintenance_update {
        integer id PK
        text message
        integer date
        integer maintenance_id FK
        integer created_by FK
        integer updated_by FK
        integer created_at
        integer updated_at
    }
    check {
        integer id PK
        text regions
        text url
        text headers
        text body
        text method
        integer count_requests
        integer workspace_id FK
        integer created_at
    }
    monitor_run {
        integer id PK
        integer workspace_id FK
        integer monitor_id FK
        integer runned_at
        integer created_at
    }
    private_location_monitor_status {
        integer monitor_id PK, FK
        integer private_location_id PK, FK
        text status
        integer cron_timestamp
        integer created_at
        integer updated_at
    }
    private_location {
        integer id PK
        text name
        text token
        text status
        text metadata
        integer last_seen_at
        integer workspace_id FK
        integer created_at
        integer updated_at
    }
    private_location_to_monitor {
        integer private_location_id FK
        integer monitor_id FK
        integer created_at
        integer deleted_at
    }
    monitor_group {
        integer id PK
        integer workspace_id FK
        integer page_id FK
        text name
        integer created_at
        integer updated_at
    }
    viewer {
        integer id PK
        text name
        text email
        integer emailVerified
        text image
        integer created_at
        integer updated_at
    }
    viewer_session {
        text session_token PK
        integer user_id FK
        integer expires
    }
    viewer_accounts {
        text user_id FK
        text type
        text provider PK
        text providerAccountId PK
        text refresh_token
        text access_token
        integer expires_at
        text token_type
        text scope
        text id_token
        text session_state
    }
    api_key {
        integer id PK
        text name
        text description
        text prefix
        text hashed_token
        integer workspace_id FK
        integer created_by_id FK
        text scopes "Stored as a JSON array so per-resource scopes (e.g. 'monitor.read') can be added later without re-shaping the column. v1 stores a single value: ['read'] or ['write']. Existing rows backfill to ['write'] via the column default to preserve current behavior."
        integer created_at
        integer expires_at
        integer last_used_at
    }
    page_component {
        integer id PK
        integer workspace_id FK
        integer page_id FK
        text type
        integer monitor_id FK
        text name
        text description
        integer order
        integer group_id FK
        integer group_order
        integer created_at
        integer updated_at
    }
    maintenance_to_page_component {
        integer maintenance_id PK, FK
        integer page_component_id PK, FK
        integer created_at
    }
    status_report_to_page_component {
        integer status_report_id PK, FK
        integer page_component_id PK, FK
        integer created_at
    }
    status_report_update_to_page_component {
        integer status_report_update_id PK, FK
        integer page_component_id PK, FK
        text impact
        integer created_at
    }
    page_component_groups {
        integer id PK
        integer workspace_id FK
        integer page_id FK
        text name
        integer default_open
        integer created_at
        integer updated_at
    }
    feedback {
        integer id PK
        integer workspace_id FK
        integer user_id FK
        text source
        text message
        text type
        integer blocker
        text path
        text metadata
        integer created_at
    }
    audit_log {
        integer id PK
        integer workspace_id
        text actor_type
        text actor_id
        integer actor_user_id
        text action
        text entity_type
        text entity_id
        text before
        text after
        text metadata
        text changed_fields
        integer created_at
    }
    external_service {
        integer id PK
        text slug
        text aliases
        text name
        text url
        text status_page_url
        text provider
        text industry
        text description
        text api_config
        integer deleted_at
        integer created_at
        integer updated_at
    }
    external_service_component {
        integer id PK
        integer external_service_id FK
        text upstream_component_id
        text slug
        text aliases
        text name
        text description
        text group_name
        integer position
        text indicator
        text status
        integer first_seen_at
        integer updated_at
    }
    external_service_incident {
        integer id PK
        integer external_service_id FK
        text provider_incident_id
        text name
        text status
        text impact
        text shortlink
        integer started_at
        integer created_at
        integer resolved_at
        text affected_component_ids
        text raw_payload
        integer raw_payload_purged_at
        integer first_seen_at
        integer updated_at
    }
    external_service_report {
        integer id PK
        integer external_service_id FK
        integer external_service_component_id FK
        text reporter_hash
        text country
        integer created_at
    }
    chat_session {
        integer id PK
        integer workspace_id FK
        integer user_id FK
        text title
        text messages
        integer created_at
        integer updated_at
    }
    frozen_monitor_uptime {
        integer id PK
        integer workspace_id FK
        integer monitor_id FK
        text month "UTC first-of-month, YYYY-MM-01"
        text days
        integer created_at
    }
    monitor_transition {
        integer id PK
        integer monitor_id FK
        text region
        integer cron_timestamp
        text from_status
        text to_status
        integer quorum_count
        integer region_count
        integer transitioned
        integer outbox_rows
        integer created_at
    }
    oauth_client {
        integer id PK
        text client_id
        text name
        text redirect_uris
        integer revoked_at
        integer created_at
    }
    oauth_session {
        text id PK
        text client_id FK
        text redirect_uri
        text scope
        text state
        text resource
        text code_challenge
        text code_challenge_method
        integer decided_at
        integer expires_at
    }
    oauth_authorization_code {
        integer id PK
        text hash
        text client_id FK
        integer user_id FK
        integer workspace_id FK
        text scope
        text redirect_uri
        text code_challenge
        integer expires_at
        integer consumed_at
        integer grant_id FK "Set on exchange so a replayed code can revoke the grant it produced."
    }
    oauth_grant {
        integer id PK
        text client_id FK
        integer user_id FK
        integer workspace_id FK
        text scope
        text access_token_hash
        integer access_token_expires_at
        text refresh_token_hash
        integer refresh_token_expires_at
        text previous_refresh_token_hash "Grace window: the hash rotated out at `rotatedAt` stays valid briefly so two concurrent refreshes do not kill the grant."
        integer rotated_at
        integer last_used_at
        integer revoked_at
        integer created_at
    }
    incident {
        integer id PK
        integer workspace_id FK
        text title
        text severity
        text status
        text summary
        integer commander_id FK
        integer declared_by FK
        integer declared_at
        integer started_at
        integer mitigated_at
        integer resolved_at
        integer resolved_by FK
        integer closed_at
        integer status_report_id FK
        text slack_team_id
        text slack_channel_id
        integer created_at
        integer updated_at
    }
    incident_event {
        integer id PK
        integer incident_id FK
        text type
        text message
        text source
        text source_ref
        integer created_by FK
        integer created_at
    }
    incident_postmortem {
        integer id PK
        integer incident_id FK
        text status
        text content
        text source_transcript
        text drafted_by
        integer created_by FK
        integer updated_by FK
        integer approved_by FK
        integer approved_at
        integer published_at
        integer created_at
        integer updated_at
    }
    slack_user {
        integer id PK
        integer workspace_id FK
        text slack_team_id
        text slack_user_id
        integer user_id FK
        integer created_at
    }
    monitor_heartbeat_state {
        integer monitor_id PK, FK
        integer last_ping_at
        text last_ping_kind "what the last ping reported: a `/fail` keeps the heartbeat down until a success ping"
        integer next_due_at "NULL until the first ping, so a new heartbeat never alerts"
    }
```
