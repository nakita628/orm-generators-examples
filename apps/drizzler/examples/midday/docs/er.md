```mermaid
erDiagram
    users |o--o{ transactions : "(id) - (assigned_id)"
    teams ||--o{ transactions : "(id) - (team_id)"
    bank_accounts |o--o{ transactions : "(id) - (bank_account_id)"
    transaction_categories |o--o{ transactions : "(team_id, slug) - (team_id, category_slug)"
    users |o--o{ tracker_entries : "(id) - (assigned_id)"
    tracker_projects |o--o{ tracker_entries : "(id) - (project_id)"
    teams |o--o{ tracker_entries : "(id) - (team_id)"
    customers ||--o{ customer_tags : "(id) - (customer_id)"
    tags ||--o{ customer_tags : "(id) - (tag_id)"
    teams ||--o{ customer_tags : "(id) - (team_id)"
    teams ||--o{ inbox_accounts : "(id) - (team_id)"
    bank_connections |o--o{ bank_accounts : "(id) - (bank_connection_id)"
    users ||--o{ bank_accounts : "(id) - (created_by)"
    teams ||--o{ bank_accounts : "(id) - (team_id)"
    teams ||--o{ invoice_recurring : "(id) - (team_id)"
    users ||--o{ invoice_recurring : "(id) - (user_id)"
    customers |o--o{ invoice_recurring : "(id) - (customer_id)"
    invoice_templates |o--o{ invoice_recurring : "(id) - (template_id)"
    users |o--o{ invoices : "(id) - (user_id)"
    customers |o--o{ invoices : "(id) - (customer_id)"
    teams ||--o{ invoices : "(id) - (team_id)"
    invoice_templates |o--o{ invoices : "(id) - (template_id)"
    invoice_recurring |o--o{ invoices : "(id) - (invoice_recurring_id)"
    teams ||--o{ customers : "(id) - (team_id)"
    teams ||--o{ tags : "(id) - (team_id)"
    users |o--o{ tracker_reports : "(id) - (created_by)"
    tracker_projects |o--o{ tracker_reports : "(id) - (project_id)"
    teams |o--o{ tracker_reports : "(id) - (team_id)"
    tags ||--o{ tracker_project_tags : "(id) - (tag_id)"
    tracker_projects ||--o{ tracker_project_tags : "(id) - (tracker_project_id)"
    teams ||--o{ tracker_project_tags : "(id) - (team_id)"
    users |o--o{ reports : "(id) - (created_by)"
    teams |o--o{ reports : "(id) - (team_id)"
    teams ||--o{ bank_connections : "(id) - (team_id)"
    teams |o--o{ user_invites : "(id) - (team_id)"
    users |o--o{ user_invites : "(id) - (invited_by)"
    teams ||--o{ document_tags : "(id) - (team_id)"
    tags ||--o{ transaction_tags : "(id) - (tag_id)"
    teams ||--o{ transaction_tags : "(id) - (team_id)"
    transactions ||--o{ transaction_tags : "(id) - (transaction_id)"
    teams |o--o{ transaction_attachments : "(id) - (team_id)"
    transactions |o--o{ transaction_attachments : "(id) - (transaction_id)"
    users |o--o{ documents : "(id) - (owner_id)"
    teams |o--o{ documents : "(id) - (team_id)"
    users |o--o{ apps : "(id) - (created_by)"
    teams |o--o{ apps : "(id) - (team_id)"
    teams ||--o{ platform_identities : "(id) - (team_id)"
    users ||--o{ platform_identities : "(id) - (user_id)"
    teams ||--o{ platform_link_tokens : "(id) - (team_id)"
    users ||--o{ platform_link_tokens : "(id) - (user_id)"
    platform_identities ||--o{ provider_notification_batches : "(id) - (platform_identity_id)"
    teams ||--o{ provider_notification_batches : "(id) - (team_id)"
    users ||--o{ provider_notification_batches : "(id) - (user_id)"
    teams ||--o{ invoice_templates : "(id) - (team_id)"
    teams ||--o{ invoice_products : "(id) - (team_id)"
    users |o--o{ invoice_products : "(id) - (created_by)"
    transaction_categories |o--o{ transaction_enrichments : "(team_id, slug) - (team_id, category_slug)"
    teams |o--o{ transaction_enrichments : "(id) - (team_id)"
    users ||--o| users : "(id) - (id)"
    teams |o--o{ users : "(id) - (team_id)"
    customers |o--o{ tracker_projects : "(id) - (customer_id)"
    teams |o--o{ tracker_projects : "(id) - (team_id)"
    transaction_attachments |o--o{ inbox : "(id) - (attachment_id)"
    teams |o--o{ inbox : "(id) - (team_id)"
    transactions |o--o{ inbox : "(id) - (transaction_id)"
    inbox_accounts |o--o{ inbox : "(id) - (inbox_account_id)"
    teams ||--o{ inbox_blocklist : "(id) - (team_id)"
    inbox ||--o{ transaction_match_suggestions : "(id) - (inbox_id)"
    transactions ||--o{ transaction_match_suggestions : "(id) - (transaction_id)"
    teams ||--o{ transaction_match_suggestions : "(id) - (team_id)"
    users |o--o{ transaction_match_suggestions : "(id) - (user_id)"
    documents ||--o{ document_tag_assignments : "(id) - (document_id)"
    document_tags ||--o{ document_tag_assignments : "(id) - (tag_id)"
    teams ||--o{ document_tag_assignments : "(id) - (team_id)"
    teams ||--o{ users_on_team : "(id) - (team_id)"
    users ||--o{ users_on_team : "(id) - (user_id)"
    teams ||--o{ transaction_categories : "(id) - (team_id)"
    transaction_categories |o--o{ transaction_categories : "(id) - (parent_id)"
    users ||--o{ short_links : "(id) - (user_id)"
    teams ||--o{ short_links : "(id) - (team_id)"
    users ||--o{ api_keys : "(id) - (user_id)"
    teams ||--o{ api_keys : "(id) - (team_id)"
    teams |o--o{ oauth_applications : "(id) - (team_id)"
    users |o--o{ oauth_applications : "(id) - (created_by)"
    oauth_applications ||--o{ oauth_authorization_codes : "(id) - (application_id)"
    users ||--o{ oauth_authorization_codes : "(id) - (user_id)"
    teams ||--o{ oauth_authorization_codes : "(id) - (team_id)"
    oauth_applications ||--o{ oauth_access_tokens : "(id) - (application_id)"
    users ||--o{ oauth_access_tokens : "(id) - (user_id)"
    teams ||--o{ oauth_access_tokens : "(id) - (team_id)"
    teams ||--o{ activities : "(id) - (team_id)"
    users |o--o{ activities : "(id) - (user_id)"
    users ||--o{ notification_settings : "(id) - (user_id)"
    teams ||--o{ notification_settings : "(id) - (team_id)"
    transactions ||--o{ accounting_sync_records : "(id) - (transaction_id)"
    teams ||--o{ accounting_sync_records : "(id) - (team_id)"
    teams ||--o{ insights : "(id) - (team_id)"
    insights ||--o{ insight_user_status : "(id) - (insight_id)"
    users ||--o{ insight_user_status : "(id) - (user_id)"
    transaction_categories ||..o{ transactions : "(team_id) - (team_id)"
    "auth.users" ||..o| users : "(id) - (id)"
    transaction_categories |o..o{ transaction_enrichments : "(team_id) - (team_id)"
    document_tag_embeddings {
        text slug PK
        vector embedding
        text name
        text model
    }
    transaction_category_embeddings {
        text name PK
        vector embedding
        text model
        boolean system
        timestamptz created_at
        timestamptz updated_at
    }
    transactions {
        uuid id PK
        timestamptz created_at
        date date
        text name
        transactionMethods method
        numeric amount
        text currency
        uuid team_id FK
        uuid assigned_id FK
        varchar note
        uuid bank_account_id FK
        text internal_id
        transactionStatus status
        numeric balance
        boolean manual
        boolean notified
        boolean internal
        text description
        text category_slug FK
        numeric baseAmount
        text counterparty_name
        text base_currency
        numeric tax_amount
        numeric tax_rate
        text tax_type
        boolean recurring
        transaction_frequency frequency
        text merchant_name
        boolean enrichment_completed
        tsvector fts_vector
    }
    tracker_entries {
        uuid id PK
        timestamptz created_at
        bigint duration "You can use { mode: 'bigint' } if numbers are exceeding js number limitations"
        uuid project_id FK
        timestamptz start
        timestamptz stop
        uuid assigned_id FK
        uuid team_id FK
        text description
        numeric rate
        text currency
        boolean billed
        date date
    }
    customer_tags {
        uuid id PK
        timestamptz created_at
        uuid customer_id FK
        uuid team_id FK
        uuid tag_id FK
    }
    inbox_accounts {
        uuid id PK
        timestamptz created_at
        text email
        text access_token
        text refresh_token
        uuid team_id FK
        timestamptz last_accessed
        inbox_account_providers provider
        text external_id
        timestamptz expiry_date
        text schedule_id
        inbox_account_status status
        text error_message
    }
    bank_accounts {
        uuid id PK
        timestamptz created_at
        uuid created_by FK
        uuid team_id FK
        text name
        text currency
        uuid bank_connection_id FK
        boolean enabled
        text account_id
        numeric balance
        boolean manual
        account_type type
        text base_currency
        numeric baseBalance
        text error_details
        smallint error_retries
        text account_reference
        text iban "Additional account data for reconnect matching and user display"
        text subtype
        text bic
        text routing_number "US bank account details (Teller, Plaid)"
        text wire_routing_number
        text account_number
        text sort_code
        numeric availableBalance "Credit account balances"
        numeric creditLimit
    }
    invoice_recurring {
        uuid id PK
        timestamptz created_at
        timestamptz updated_at
        uuid team_id FK
        uuid user_id FK
        uuid customer_id FK
        invoice_recurring_frequency frequency "Frequency settings"
        integer frequency_day
        integer frequency_week
        integer frequency_interval
        invoice_recurring_end_type end_type "End conditions"
        timestamptz end_date
        integer end_count
        invoice_recurring_status status "Status tracking"
        integer invoices_generated
        integer consecutive_failures
        timestamptz next_scheduled_at
        timestamptz last_generated_at
        text timezone
        integer due_date_offset "Invoice template data"
        numeric amount
        text currency
        jsonb line_items
        jsonb template
        jsonb payment_details
        jsonb from_details
        jsonb note_details
        text customer_name
        numeric vat
        numeric tax
        numeric discount
        numeric subtotal
        jsonb top_block
        jsonb bottom_block
        uuid template_id FK
        timestamptz upcoming_notification_sent_at "Notification tracking"
    }
    invoices {
        uuid id PK
        timestamptz created_at
        timestamptz updated_at
        timestamptz due_date
        text invoice_number
        uuid customer_id FK
        numeric amount
        text currency
        jsonb line_items
        jsonb payment_details
        jsonb customer_details
        jsonb company_datails
        text note
        text internal_note
        uuid team_id FK
        timestamptz paid_at
        tsvector fts
        numeric vat
        numeric tax
        text url
        text file_path
        invoice_status status
        timestamptz viewed_at
        jsonb from_details
        timestamptz issue_date
        jsonb template
        jsonb note_details
        text customer_name
        text token
        text sent_to
        timestamptz reminder_sent_at
        numeric discount
        bigint file_size
        uuid user_id FK
        numeric subtotal
        jsonb top_block
        jsonb bottom_block
        timestamptz sent_at
        timestamptz scheduled_at
        text scheduled_job_id
        uuid template_id FK
        text payment_intent_id
        timestamptz refunded_at
        uuid invoice_recurring_id FK "Recurring invoice fields"
        integer recurring_sequence
    }
    customers {
        uuid id PK
        timestamptz created_at
        text name
        text email
        text billingEmail
        text country
        text address_line_1
        text address_line_2
        text city
        text state
        text zip
        text note
        uuid team_id FK
        text website
        text phone
        text vat_number
        text country_code
        text token
        text contact
        text status "Customer relationship fields"
        text preferred_currency
        integer default_payment_terms
        boolean is_archived
        text source
        text external_id
        text logo_url "Enrichment fields (from Gemini + Google Search grounding)"
        text description
        text industry
        text company_type
        text employee_count
        integer founded_year
        text estimated_revenue
        text funding_stage
        text total_funding
        text headquarters_location
        text timezone
        text linkedin_url
        text twitter_url
        text instagram_url
        text facebook_url
        text ceo_name
        text finance_contact
        text finance_contact_email
        text primary_language
        text fiscal_year_end
        text enrichment_status "Enrichment metadata"
        timestamptz enriched_at
        boolean portal_enabled "Portal fields"
        text portal_id
        tsvector fts
    }
    exchange_rates {
        uuid id PK
        text base
        numeric rate
        text target
        timestamptz updated_at
    }
    tags {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        text name
    }
    tracker_reports {
        uuid id PK
        timestamptz created_at
        text link_id
        text short_link
        uuid team_id FK
        uuid project_id FK
        uuid created_by FK
    }
    invoice_comments {
        uuid id PK
        timestamptz created_at
    }
    tracker_project_tags {
        uuid id PK
        timestamptz created_at
        uuid tracker_project_id FK
        uuid tag_id FK
        uuid team_id FK
    }
    reports {
        uuid id PK
        timestamptz created_at
        text link_id
        uuid team_id FK
        text short_link
        timestamptz from
        timestamptz to
        reportTypes type
        timestamptz expire_at
        text currency
        uuid created_by FK
    }
    institutions {
        text id PK
        text name
        text logo
        bank_providers provider
        text countries
        integer available_history
        integer maximum_consent_validity
        integer popularity
        text type
        institution_status status
        timestamptz created_at
        timestamptz updated_at
    }
    bank_connections {
        uuid id PK
        timestamptz created_at
        text institution_id
        timestamptz expires_at
        uuid team_id FK
        text name
        text logo_url
        text access_token
        text enrollment_id
        bank_providers provider
        timestamptz last_accessed
        text reference_id
        connection_status status
        text error_details
        smallint error_retries
    }
    user_invites {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        text email
        teamRoles role
        text code
        uuid invited_by FK
    }
    document_tags {
        uuid id PK
        timestamptz created_at
        text name
        text slug
        uuid team_id FK
    }
    transaction_tags {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        uuid tag_id FK
        uuid transaction_id FK
    }
    transaction_attachments {
        uuid id PK
        timestamptz created_at
        text type
        uuid transaction_id FK
        uuid team_id FK
        bigint size "You can use { mode: 'bigint' } if numbers are exceeding js number limitations"
        text name
        text path
    }
    teams {
        uuid id PK
        timestamptz created_at
        text name
        text logo_url
        text inbox_id
        text email
        text inbox_email
        boolean inbox_forwarding
        text base_currency
        text country_code
        smallint fiscal_year_start_month
        boolean document_classification
        text flags
        timestamptz canceled_at
        plans plan
        subscription_status subscription_status
        jsonb export_settings
        text stripe_account_id
        text stripe_connect_status
        text company_type
        text heard_about
    }
    documents {
        uuid id PK
        text name
        timestamptz created_at
        jsonb metadata
        text path_tokens
        uuid team_id FK
        text parent_id
        uuid object_id
        uuid owner_id FK
        text tag
        text title
        text body
        tsvector fts
        text summary
        text content
        date date
        text language
        document_processing_status processing_status
        tsvector fts_simple
        tsvector fts_english
        tsvector fts_language
    }
    apps {
        uuid id PK
        uuid team_id FK
        jsonb config
        timestamptz created_at
        text app_id
        uuid created_by FK
        jsonb settings
    }
    platform_identities {
        uuid id PK
        platform_provider provider
        uuid team_id FK
        uuid user_id FK
        text external_user_id
        text external_team_id
        text external_channel_id
        jsonb metadata
        timestamptz created_at
        timestamptz updated_at
    }
    platform_link_tokens {
        uuid id PK
        text code
        platform_provider provider
        uuid team_id FK
        uuid user_id FK
        timestamptz expires_at
        timestamptz used_at
        jsonb metadata
        timestamptz created_at
    }
    provider_notification_batches {
        uuid id PK
        text batch_key
        uuid platform_identity_id FK
        uuid team_id FK
        uuid user_id FK
        platform_provider provider
        text event_family
        jsonb payload
        jsonb notification_context
        timestamptz window_ends_at
        timestamptz sent_at
        timestamptz created_at
        timestamptz updated_at
    }
    invoice_templates {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        text name
        boolean is_default
        text customer_label
        text from_label
        text invoice_no_label
        text issue_date_label
        text due_date_label
        text description_label
        text price_label
        text quantity_label
        text total_label
        text vat_label
        text tax_label
        text payment_label
        text note_label
        text logo_url
        text currency
        jsonb payment_details
        jsonb from_details
        jsonb note_details
        invoice_size size
        text date_format
        boolean include_vat
        boolean include_tax
        numeric tax_rate
        invoice_delivery_type delivery_type
        text discount_label
        boolean include_discount
        boolean include_decimals
        boolean include_qr
        text total_summary_label
        text title
        numeric vat_rate
        boolean include_units
        text subtotal_label
        boolean include_pdf
        boolean send_copy
        boolean include_line_item_tax
        text line_item_tax_label
        boolean payment_enabled
        integer payment_terms_days
        text email_subject
        text email_heading
        text email_body
        text email_button_text
    }
    invoice_products {
        uuid id PK
        timestamptz created_at
        timestamptz updated_at
        uuid team_id FK
        uuid created_by FK
        text name
        text description
        numeric price
        text currency
        text unit
        numeric tax_rate
        boolean isActive
        integer usage_count
        timestamptz last_used_at
        tsvector fts "Full-text search for product names and descriptions"
    }
    transaction_enrichments {
        uuid id PK
        timestamptz created_at
        text name
        uuid team_id FK
        text category_slug FK
        boolean system
    }
    users {
        uuid id PK, FK
        text full_name
        text avatar_url
        text email
        uuid team_id FK
        timestamptz created_at
        text locale
        boolean week_starts_on_monday
        text timezone
        boolean timezone_auto_sync
        numeric time_format
        text date_format
    }
    tracker_projects {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        numeric rate
        text currency
        trackerStatus status
        text description
        text name
        boolean billable
        bigint estimate "You can use { mode: 'bigint' } if numbers are exceeding js number limitations"
        uuid customer_id FK
        tsvector fts
    }
    inbox {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        text file_path
        text file_name
        uuid transaction_id FK
        numeric amount
        text currency
        text content_type
        bigint size "You can use { mode: 'bigint' } if numbers are exceeding js number limitations"
        uuid attachment_id FK
        date date
        text forwarded_to
        text reference_id
        json meta
        inbox_status status
        text website
        text sender_email
        text display_name
        tsvector fts
        inbox_type type
        text description
        numeric base_amount
        text base_currency
        numeric tax_amount
        numeric tax_rate
        text tax_type
        uuid inbox_account_id FK
        text invoice_number
        uuid grouped_inbox_id
    }
    inbox_blocklist {
        uuid id PK
        timestamptz created_at
        uuid team_id FK
        inbox_blocklist_type type
        text value
    }
    transaction_match_suggestions {
        uuid id PK
        timestamptz created_at
        timestamptz updated_at
        uuid team_id FK "Core relationship"
        uuid inbox_id FK
        uuid transaction_id FK
        numeric confidence_score "Match scores for transparency"
        numeric amount_score
        numeric currency_score
        numeric date_score
        numeric name_score
        text match_type "Match context"
        jsonb match_details
        text status "User interaction tracking"
        timestamptz user_action_at
        uuid user_id FK
    }
    document_tag_assignments {
        uuid document_id PK, FK
        uuid tag_id PK, FK
        uuid team_id FK
    }
    users_on_team {
        uuid user_id PK, FK
        uuid team_id PK, FK
        uuid id PK
        teamRoles role
        timestamptz created_at
    }
    transaction_categories {
        uuid id
        text name
        uuid team_id PK, FK
        text color
        timestamptz created_at
        boolean system
        text slug PK
        numeric tax_rate
        text tax_type
        text tax_reporting_code
        boolean excluded
        text description
        uuid parent_id FK
    }
    "auth.users" {
        uuid instance_id
        uuid id PK
        varchar aud
        varchar role
        varchar email
        varchar encrypted_password
        timestamptz email_confirmed_at
        timestamptz invited_at
        varchar confirmation_token
        timestamptz confirmation_sent_at
        varchar recovery_token
        timestamptz recovery_sent_at
        varchar email_change_token_new
        varchar email_change
        timestamptz email_change_sent_at
        timestamptz last_sign_in_at
        jsonb raw_app_meta_data
        jsonb raw_user_meta_data
        boolean is_super_admin
        timestamptz created_at
        timestamptz updated_at
        text phone
        timestamptz phone_confirmed_at
        text phone_change
        varchar phone_change_token
        timestamptz phone_change_sent_at
        timestamptz confirmed_at "Drizzle ORM does not support .stored() for generated columns, so we omit it"
        varchar email_change_token_current
        smallint email_change_confirm_status
        timestamptz banned_until
        varchar reauthentication_token
        timestamptz reauthentication_sent_at
        boolean is_sso_user
        timestamptz deleted_at
        boolean is_anonymous
    }
    short_links {
        uuid id PK
        text short_id
        text url
        text type
        numeric size
        text mime_type
        text file_name
        uuid team_id FK
        uuid user_id FK
        timestamptz expires_at
        timestamptz created_at
    }
    api_keys {
        uuid id PK
        text key_encrypted
        text name
        timestamptz created_at
        uuid user_id FK
        uuid team_id FK
        text key_hash
        text scopes
        timestamptz last_used_at
    }
    oauth_applications {
        uuid id PK
        text name
        text slug
        text description
        text overview
        text developer_name
        text logo_url
        text website
        text install_url
        text screenshots
        text redirect_uris
        text client_id
        text client_secret
        text scopes
        uuid team_id FK
        uuid created_by FK
        timestamptz created_at
        timestamptz updated_at
        boolean is_public
        boolean active
        text status
    }
    oauth_authorization_codes {
        uuid id PK
        text code
        uuid application_id FK
        uuid user_id FK
        uuid team_id FK
        text scopes
        text redirect_uri
        timestamptz expires_at
        timestamptz created_at
        boolean used
        text code_challenge
        text code_challenge_method
    }
    oauth_access_tokens {
        uuid id PK
        text token
        text refresh_token
        uuid application_id FK
        uuid user_id FK
        uuid team_id FK
        text scopes
        timestamptz expires_at
        timestamptz refresh_token_expires_at
        timestamptz created_at
        timestamptz last_used_at
        boolean revoked
        timestamptz revoked_at
    }
    activities {
        uuid id PK
        timestamptz created_at
        uuid team_id FK "Core fields"
        uuid user_id FK
        activity_type type
        smallint priority
        uuid group_id "Group related activities together (e.g., same business event across multiple users)"
        activity_source source "Source of the activity"
        jsonb metadata "All the data"
        activity_status status "Simple lifecycle (only for notifications)"
        timestamptz last_used_at "Timestamp of last system use (e.g. insight generation, digest inclusion)"
    }
    notification_settings {
        uuid id PK
        uuid user_id FK
        uuid team_id FK
        text notification_type
        text channel
        boolean enabled
        timestamptz created_at
        timestamptz updated_at
    }
    accounting_sync_records {
        uuid id PK
        uuid transaction_id FK
        uuid team_id FK
        accounting_provider provider
        text provider_tenant_id
        text provider_transaction_id
        jsonb synced_attachment_mapping "Maps Midday attachment IDs to provider attachment IDs for sync tracking Format: { 'midday-attachment-id': 'provider-attachment-id' }"
        timestamptz synced_at
        accounting_sync_type sync_type
        accounting_sync_status status
        text error_message
        text error_code "Standardized error code for frontend handling (e.g., 'ATTACHMENT_UNSUPPORTED_TYPE', 'AUTH_EXPIRED')"
        text provider_entity_type "Provider-specific entity type (e.g., 'Purchase', 'SalesReceipt', 'Voucher', 'BankTransaction')"
        timestamptz created_at "When the record was first created (synced_at gets updated on every sync)"
    }
    insights {
        uuid id PK
        uuid team_id FK
        insight_period_type period_type "Flexible period definition"
        timestamptz period_start
        timestamptz period_end
        smallint period_year
        smallint period_number
        insight_status status
        jsonb selected_metrics "Selected 4 key metrics (dynamically chosen)"
        jsonb all_metrics "Full metrics snapshot (for drill-down)"
        jsonb anomalies "Detected anomalies and patterns"
        jsonb expense_anomalies "Expense category anomalies (spikes, new categories, etc.)"
        jsonb milestones "Streaks and milestones"
        jsonb activity "Activity context"
        varchar currency
        text title "AI-generated title (for card headers and email subjects)"
        jsonb content "AI-generated content (relief-first structure)"
        jsonb predictions "Forward-looking predictions for follow-through tracking"
        text audio_path "Audio narration storage path: {teamId}/insights/{insightId}.mp3"
        timestamptz generated_at
        timestamptz created_at
        timestamptz updated_at
    }
    insight_user_status {
        uuid insight_id PK, FK
        uuid user_id PK, FK
        timestamptz read_at
        timestamptz dismissed_at
        timestamptz created_at
        timestamptz updated_at
    }
```
