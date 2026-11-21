```mermaid
erDiagram
    user ||--o{ session : "(id) - (user_id)"
    user ||--o{ account : "(id) - (user_id)"
    user |o--o{ oauth_client : "(id) - (user_id)"
    oauth_client ||--o{ oauth_client_resource : "(client_id) - (client_id)"
    oauth_resource ||--o{ oauth_client_resource : "(identifier) - (resource_id)"
    oauth_client ||--o{ oauth_refresh_token : "(client_id) - (client_id)"
    session |o--o{ oauth_refresh_token : "(id) - (session_id)"
    user ||--o{ oauth_refresh_token : "(id) - (user_id)"
    oauth_client ||--o{ oauth_access_token : "(client_id) - (client_id)"
    session |o--o{ oauth_access_token : "(id) - (session_id)"
    user |o--o{ oauth_access_token : "(id) - (user_id)"
    oauth_refresh_token |o--o{ oauth_access_token : "(id) - (refresh_id)"
    oauth_client ||--o{ oauth_consent : "(client_id) - (client_id)"
    user |o--o{ oauth_consent : "(id) - (user_id)"
    user ||--o{ mcp_revocation : "(id) - (user_id)"
    oauth_client ||--o{ mcp_revocation : "(client_id) - (client_id)"
    user ||--o{ followers : "(id) - (follower_id)"
    user ||--o{ followers : "(id) - (followed_id)"
    user ||--o{ user_media_update : "(id) - (user_id)"
    user ||--o{ user_media_monthly_activity : "(id) - (user_id)"
    user ||--o{ user_media_settings : "(id) - (user_id)"
    user ||--o{ user_media_stats_history : "(id) - (user_id)"
    user ||--o{ profile_custom : "(id) - (user_id)"
    user ||--o{ series_list : "(id) - (user_id)"
    series ||--o{ series_list : "(id) - (media_id)"
    series_list ||--o{ series_list_seasons : "(id) - (list_id)"
    series ||--o{ series_genre : "(id) - (media_id)"
    series ||--o{ series_actors : "(id) - (media_id)"
    series ||--o{ series_network : "(id) - (media_id)"
    series ||--o{ series_episodes_per_season : "(id) - (media_id)"
    user ||--o{ series_tags : "(id) - (user_id)"
    series ||--o{ series_tags : "(id) - (media_id)"
    media_tags ||--o{ series_tags : "(id) - (tag_id)"
    user ||--o{ anime_list : "(id) - (user_id)"
    anime ||--o{ anime_list : "(id) - (media_id)"
    anime_list ||--o{ anime_list_seasons : "(id) - (list_id)"
    anime ||--o{ anime_genre : "(id) - (media_id)"
    anime ||--o{ anime_actors : "(id) - (media_id)"
    anime ||--o{ anime_network : "(id) - (media_id)"
    anime ||--o{ anime_episodes_per_season : "(id) - (media_id)"
    user ||--o{ anime_tags : "(id) - (user_id)"
    anime ||--o{ anime_tags : "(id) - (media_id)"
    media_tags ||--o{ anime_tags : "(id) - (tag_id)"
    user ||--o{ movies_list : "(id) - (user_id)"
    movies ||--o{ movies_list : "(id) - (media_id)"
    movies ||--o{ movies_genre : "(id) - (media_id)"
    movies ||--o{ movies_actors : "(id) - (media_id)"
    user ||--o{ movies_tags : "(id) - (user_id)"
    movies ||--o{ movies_tags : "(id) - (media_id)"
    media_tags ||--o{ movies_tags : "(id) - (tag_id)"
    user ||--o{ games_list : "(id) - (user_id)"
    games ||--o{ games_list : "(id) - (media_id)"
    games ||--o{ games_genre : "(id) - (media_id)"
    games ||--o{ games_platforms : "(id) - (media_id)"
    games ||--o{ games_companies : "(id) - (media_id)"
    user ||--o{ games_tags : "(id) - (user_id)"
    games ||--o{ games_tags : "(id) - (media_id)"
    media_tags ||--o{ games_tags : "(id) - (tag_id)"
    user ||--o{ books_list : "(id) - (user_id)"
    books ||--o{ books_list : "(id) - (media_id)"
    books ||--o{ books_genre : "(id) - (media_id)"
    books ||--o{ books_authors : "(id) - (media_id)"
    user ||--o{ books_tags : "(id) - (user_id)"
    books ||--o{ books_tags : "(id) - (media_id)"
    media_tags ||--o{ books_tags : "(id) - (tag_id)"
    user ||--o{ manga_list : "(id) - (user_id)"
    manga ||--o{ manga_list : "(id) - (media_id)"
    manga ||--o{ manga_genre : "(id) - (media_id)"
    manga ||--o{ manga_authors : "(id) - (media_id)"
    user ||--o{ manga_tags : "(id) - (user_id)"
    manga ||--o{ manga_tags : "(id) - (media_id)"
    media_tags ||--o{ manga_tags : "(id) - (tag_id)"
    user ||--o{ mediadle_stats : "(id) - (user_id)"
    user ||--o{ user_mediadle_progress : "(id) - (user_id)"
    daily_mediadle ||--o{ user_mediadle_progress : "(id) - (daily_mediadle_id)"
    user ||--o| which_came_first_runs : "(id) - (user_id)"
    which_came_first_runs ||--o| which_came_first_rounds : "(id) - (run_id)"
    achievement ||--o{ achievement_tier : "(id) - (achievement_id)"
    user ||--o{ user_achievement : "(id) - (user_id)"
    achievement_tier ||--o{ user_achievement : "(id) - (tier_id)"
    user ||--o{ social_notifications : "(id) - (actor_id)"
    user ||--o{ social_notifications : "(id) - (user_id)"
    feature_requests |o--o{ social_notifications : "(id) - (feature_request_id)"
    user ||--o{ media_notifications : "(id) - (user_id)"
    user |o--o{ task_history : "(id) - (user_id)"
    user ||--o{ media_refresh_log : "(id) - (user_id)"
    user |o--o{ feature_requests : "(id) - (created_by)"
    user ||--o{ feature_votes : "(id) - (user_id)"
    feature_requests ||--o{ feature_votes : "(id) - (feature_id)"
    user ||--o{ collections : "(id) - (owner_id)"
    collections ||--o{ collection_items : "(id) - (collection_id)"
    user ||--o{ collection_likes : "(id) - (user_id)"
    collections ||--o{ collection_likes : "(id) - (collection_id)"
    user ||--o| import_jobs : "(id) - (user_id)"
    import_jobs ||--o{ import_items : "(id) - (job_id)"
    user ||--o{ smart_views : "(id) - (user_id)"
    user ||--o{ media_tags : "(id) - (user_id)"
    user {
        integer id PK
        text name
        text email
        text created_at
        text updated_at
        text image
        integer profile_views
        text role
        integer email_verified
        text privacy
        integer hide_aggregate_ratings
        integer grid_list_view
        integer show_onboarding
        integer show_update_modal
        text background_image
        text rating_system
        text search_selector
        integer auto_move_completed_tv_to_on_hold
    }
    session {
        integer id PK
        integer expires_at
        text token
        integer created_at
        integer updated_at
        text ip_address
        text user_agent
        integer user_id FK
    }
    account {
        integer id PK
        text account_id
        text provider_id
        integer user_id FK
        text access_token
        text refresh_token
        text id_token
        integer access_token_expires_at
        integer refresh_token_expires_at
        text scope
        text password
        integer created_at
        integer updated_at
    }
    verification {
        integer id PK
        text identifier
        text value
        integer expires_at
        integer created_at
        integer updated_at
    }
    jwks {
        integer id PK
        text public_key
        text private_key
        integer created_at
        integer expires_at
        text alg
        text crv
    }
    oauth_client {
        integer id PK
        text client_id
        text client_secret
        text client_discovery_id
        integer disabled
        integer skip_consent
        integer enable_end_session
        text subject_type
        text scopes
        text client_credentials_scopes
        integer user_id FK
        integer created_at
        integer updated_at
        text name
        text uri
        text icon
        text contacts
        text tos
        text policy
        text software_id
        text software_version
        text software_statement
        text redirect_uris
        text post_logout_redirect_uris
        text backchannel_logout_uri
        integer backchannel_logout_session_required
        text token_endpoint_auth_method
        text application_type
        text jwks
        text jwks_uri
        text grant_types
        text response_types
        integer require_pkce
        integer dpop_bound_access_tokens
        text reference_id
        text metadata
    }
    oauth_resource {
        integer id PK
        text identifier
        text name
        integer access_token_ttl
        integer refresh_token_ttl
        text signing_algorithm
        text signing_key_id
        text allowed_scopes
        text custom_claims
        integer dpop_bound_access_tokens_required
        integer disabled
        integer created_at
        integer updated_at
        integer policy_version
        text metadata
    }
    oauth_client_resource {
        integer id PK
        text client_id FK
        text resource_id FK
        text metadata
        integer created_at
    }
    oauth_refresh_token {
        integer id PK
        text token
        text client_id FK
        integer session_id FK
        integer user_id FK
        text reference_id
        text authorization_code_id
        text resources
        text requested_user_info_claims
        integer expires_at
        integer created_at
        integer revoked
        integer rotated_at
        text rotation_replay_response
        integer rotation_replay_expires_at
        integer auth_time
        text confirmation
        text scopes
    }
    oauth_access_token {
        integer id PK
        text token
        text client_id FK
        integer session_id FK
        integer user_id FK
        text reference_id
        text authorization_code_id
        text resources
        text requested_user_info_claims
        integer refresh_id FK
        integer expires_at
        integer created_at
        integer revoked
        text confirmation
        text scopes
    }
    oauth_consent {
        integer id PK
        text client_id FK
        integer user_id FK
        text reference_id
        text resources
        text requested_user_info_claims
        text scopes
        integer created_at
        integer updated_at
    }
    oauth_client_assertion {
        text id PK
        integer expires_at
    }
    mcp_revocation {
        integer user_id PK, FK
        text client_id PK, FK
        integer revoked_at
    }
    followers {
        integer follower_id FK
        integer followed_id FK
        text status
    }
    user_media_update {
        integer id PK
        integer user_id FK
        integer media_id
        text media_name
        text media_type
        text update_type
        text payload
        text timestamp
    }
    user_media_monthly_activity {
        integer id PK
        integer user_id FK
        integer media_id
        text month_bucket
        real progress_gained
        text media_type
        integer redo_gained
        text last_activity_at
        integer hidden
        integer had_completion
    }
    user_media_settings {
        integer id PK
        integer user_id FK
        text media_type
        integer time_spent
        integer views
        integer active
        integer total_entries
        integer total_redo
        integer entries_rated
        integer sum_entries_rated
        integer entries_commented
        integer entries_favorites
        integer total_specific
        text status_counts
        integer release_notifications_enabled
    }
    user_media_stats_history {
        integer id PK
        integer user_id FK
        text media_type
        integer time_spent
        integer views
        integer active
        integer total_entries
        integer total_redo
        integer entries_rated
        integer sum_entries_rated
        integer entries_commented
        integer entries_favorites
        integer total_specific
        text status_counts
        text timestamp
        integer media_id
    }
    profile_custom {
        integer id PK
        integer user_id FK
        text key
        text value
        text created_at
        text updated_at
    }
    api_tokens {
        integer id PK
        text provider
        text access_token
        integer expires_at
    }
    series {
        text backdrop_path
        text original_name
        text last_air_date
        text homepage
        text created_by
        integer duration
        integer total_seasons
        integer total_episodes
        text origin_country
        text prod_status
        real vote_average
        real vote_count
        real popularity
        integer api_id
        integer episode_to_air
        integer season_to_air
        text next_episode_to_air
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    series_list {
        integer current_season
        integer current_episode
        integer redo
        integer total
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    series_list_seasons {
        real rating
        integer season PK
        integer redo
        integer list_id PK, FK
    }
    series_genre {
        integer id PK
        integer media_id FK
        text name
    }
    series_actors {
        integer id PK
        integer media_id FK
        text name
    }
    series_network {
        integer id PK
        integer media_id FK
        text name
    }
    series_episodes_per_season {
        integer id PK
        integer media_id FK
        integer season
        integer episodes
    }
    series_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    anime {
        text backdrop_path
        text original_name
        text last_air_date
        text homepage
        text created_by
        integer duration
        integer total_seasons
        integer total_episodes
        text origin_country
        text prod_status
        real vote_average
        real vote_count
        real popularity
        integer season_to_air
        integer episode_to_air
        text next_episode_to_air
        integer api_id
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    anime_list {
        integer current_season
        integer current_episode
        integer redo
        integer total
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    anime_list_seasons {
        real rating
        integer season PK
        integer redo
        integer list_id PK, FK
    }
    anime_genre {
        integer id PK
        integer media_id FK
        text name
    }
    anime_actors {
        integer id PK
        integer media_id FK
        text name
    }
    anime_network {
        integer id PK
        integer media_id FK
        text name
    }
    anime_episodes_per_season {
        integer id PK
        integer media_id FK
        integer season
        integer episodes
    }
    anime_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    movies {
        text backdrop_path
        text original_name
        text homepage
        integer duration
        text original_language
        real vote_average
        real vote_count
        real popularity
        real budget
        real revenue
        text tagline
        integer api_id
        integer collection_id
        text director_name
        text compositor_name
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    movies_list {
        integer redo
        integer total
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    movies_genre {
        integer id PK
        integer media_id FK
        text name
    }
    movies_actors {
        integer id PK
        integer media_id FK
        text name
    }
    movies_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    games {
        text backdrop_image_id
        text game_engine
        text game_modes
        text player_perspective
        real vote_average
        real vote_count
        text igdb_url
        real hltb_main_time
        text hltb_last_checked_at
        real hltb_main_and_extra_time
        real hltb_total_complete_time
        text steam_api_id
        integer collection_id
        integer api_id
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    games_list {
        integer playtime
        text platform
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    games_genre {
        integer id PK
        integer media_id FK
        text name
    }
    games_platforms {
        integer id PK
        integer media_id FK
        text name
    }
    games_companies {
        integer id PK
        integer media_id FK
        text name
        integer publisher
        integer developer
    }
    games_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    books {
        integer pages
        text language
        text publishers
        text api_id
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    books_list {
        integer actual_page
        integer redo
        integer total
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    books_genre {
        integer id PK
        integer media_id FK
        text name
    }
    books_authors {
        integer id PK
        integer media_id FK
        text name
    }
    books_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    manga {
        text original_name
        integer chapters
        text prod_status
        text site_url
        text end_date
        integer volumes
        real vote_average
        real vote_count
        real popularity
        text publishers
        integer api_id
        integer id PK
        text name
        text release_date
        text synopsis
        text image_cover
        integer lock_status
        text added_at
        text last_api_update
    }
    manga_list {
        integer current_chapter
        integer redo
        integer total
        integer id PK
        integer user_id FK
        integer media_id FK
        text status
        integer release_notifications_enabled
        integer favorite
        text comment
        real rating
        text custom_cover
        text added_at
        text last_updated
    }
    manga_genre {
        integer id PK
        integer media_id FK
        text name
    }
    manga_authors {
        integer id PK
        integer media_id FK
        text name
    }
    manga_tags {
        integer id PK
        integer user_id FK
        integer media_id FK
        integer tag_id FK
    }
    daily_mediadle {
        integer id PK
        text media_type
        integer media_id
        text date
        integer pixelation_levels
    }
    mediadle_stats {
        integer id PK
        integer user_id FK
        text media_type
        integer total_played
        integer total_won
        real average_attempts
        integer streak
        integer best_streak
    }
    user_mediadle_progress {
        integer id PK
        integer user_id FK
        integer daily_mediadle_id FK
        integer attempts
        integer completed
        integer succeeded
        text completion_time
    }
    which_came_first_media {
        integer media_id PK
        text release_date
        text media_type PK
    }
    which_came_first_runs {
        integer id PK
        integer user_id FK
        text completed_at
        integer score
        text started_at
        text selected_media_types
        text status
    }
    which_came_first_rounds {
        integer id PK
        integer run_id FK
        integer round_number
        integer left_media_id
        text left_release_date
        integer right_media_id
        text right_release_date
        integer correct
        text selected_side
        text left_media_type
        text right_media_type
        text answered_at
        text created_at
    }
    achievement {
        integer id PK
        text name
        text description
        text code_name
        text media_type
        text value
    }
    achievement_tier {
        integer id PK
        integer achievement_id FK
        text difficulty
        text criteria
        real rarity
    }
    user_achievement {
        integer id PK
        integer user_id FK
        integer tier_id FK
        real progress
        real count
        integer completed
        text completed_at
        text last_calculated_at
    }
    social_notifications {
        integer id PK
        integer actor_id FK
        integer user_id FK
        integer feature_request_id FK
        text type
        integer read
        text created_at
    }
    media_notifications {
        integer id PK
        integer user_id FK
        text name
        integer media_id
        text media_type
        integer season
        integer episode
        integer is_season_finale
        text release_date
        integer read
        text created_at
    }
    task_history {
        integer id PK
        text task_id
        integer user_id FK
        text status
        text error_message
        text task_name
        text triggered_by
        text logs
        text started_at
        text finished_at
    }
    year_recap_release {
        integer year PK
        text mode
        text updated_at
    }
    inactive_account_deletion {
        integer id PK
        text deleted_at
        integer user_id
        text username
        text resurrected_at
        text warning_sent_at
        text last_email_error
        text last_seen_at
        text warning_token_hash
        text last_email_attempt_at
        text deletion_scheduled_at
        text status
        integer email_retry_count
        text created_at
        text updated_at
    }
    media_refresh_log {
        integer id PK
        integer user_id FK
        text media_type
        text api_id
        text refreshed_at
    }
    api_call_rollup {
        integer id PK
        text provider
        integer bucket_start_ms
        text bucket_start
        integer total
        integer errors
        integer duration_ms_total
        integer max_second_burst
        text status_counts
    }
    feature_requests {
        integer id PK
        text title
        text admin_comment
        text description
        text status
        integer created_by FK
        text created_at
    }
    feature_votes {
        integer id PK
        integer user_id FK
        integer feature_id FK
        text created_at
    }
    collections {
        integer id PK
        integer owner_id FK
        integer profile_position
        text title
        text description
        integer view_count
        integer like_count
        integer copied_count
        integer ordered
        text privacy
        text created_at
        text updated_at
    }
    collection_items {
        integer id PK
        integer collection_id FK
        text annotation
        integer media_id
        integer order_index
        text media_type
        text created_at
    }
    collection_likes {
        integer id PK
        integer user_id FK
        integer collection_id FK
        text created_at
    }
    import_jobs {
        integer id PK
        integer user_id FK
        text error
        text source
        integer total_count
        integer failed_count
        integer skipped_count
        integer completed_count
        integer processed_count
        text status
        text created_at
        text updated_at
        text started_at
        text finished_at
        text next_attempt_at
    }
    import_items {
        integer id PK
        integer job_id FK
        text name
        text release_date
        text status_reason
        text external_api_id
        integer row_number
        integer matched_media_id
        text media_type
        text payload_json
        text external_api_source
        text status
        text created_at
        text updated_at
    }
    smart_views {
        integer id PK
        integer user_id FK
        integer profile_position
        text spec
        text created_at
        text updated_at
    }
    media_tags {
        integer id PK
        integer user_id FK
        text media_type
        text name
    }
```
