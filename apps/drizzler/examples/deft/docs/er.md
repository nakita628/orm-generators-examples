```mermaid
erDiagram
    native_reference_states ||--o{ native_mention_deliveries : "(id) - (source_state_id)"
    users ||--o{ native_mention_deliveries : "(id) - (recipient_user_id)"
    users ||--o{ native_mention_deliveries : "(id) - (actor_user_id)"
    native_reference_states ||--o{ native_mention_deliveries : "(org_id, id) - (org_id, source_state_id)"
    orgs ||--o{ native_create_requests : "(id) - (org_id)"
    users ||--o{ native_create_requests : "(id) - (user_id)"
    users ||--o{ org_members : "(id) - (user_id)"
    users ||--o{ invites : "(id) - (invited_by)"
    users |o--o{ invites : "(id) - (accepted_by)"
    users |o--o{ spaces : "(id) - (created_by)"
    spaces ||--o{ space_members : "(id) - (space_id)"
    users ||--o{ space_members : "(id) - (user_id)"
    spaces ||--o{ messages : "(id) - (space_id)"
    users ||--o{ messages : "(id) - (user_id)"
    spaces ||--o| app_private_defty_seals : "(org_id, id) - (org_id, space_id)"
    users ||--o{ app_private_defty_seals : "(id) - (owner_user_id)"
    users ||--o{ app_private_defty_seals : "(id) - (defty_user_id)"
    app_private_defty_seals ||--o| app_private_defty_grants : "(org_id, id) - (org_id, seal_id)"
    app_resource_bindings ||--o{ app_private_defty_grants : "(org_id, id, owner_user_id) - (org_id, resource_binding_id, owner_user_id)"
    app_sync_checkpoints ||--o{ app_private_defty_grants : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_resource_projections ||--o{ app_private_defty_grants : "(org_id, id) - (org_id, projection_id)"
    messages ||--o{ reactions : "(id) - (message_id)"
    users ||--o{ reactions : "(id) - (user_id)"
    users ||--o{ files : "(id) - (uploaded_by)"
    messages |o--o{ files : "(id) - (message_id)"
    files ||--o{ attachment_derivatives : "(org_id, id) - (org_id, file_id)"
    users |o--o{ projects : "(id) - (lead_id)"
    projects ||--o{ project_spaces : "(id) - (project_id)"
    spaces ||--o{ project_spaces : "(id) - (space_id)"
    projects ||--o{ tasks : "(id) - (project_id)"
    users |o--o{ tasks : "(id) - (assignee_id)"
    users ||--o{ tasks : "(id) - (created_by)"
    messages |o--o{ tasks : "(id) - (source_message_id)"
    tasks |o--o{ tasks : "(id) - (parent_task_id)"
    messages ||--o{ message_attachments : "(org_id, id) - (org_id, message_id)"
    files ||--o{ message_attachments : "(org_id, id) - (org_id, file_id)"
    tasks ||--o{ task_attachments : "(org_id, id) - (org_id, task_id)"
    files ||--o{ task_attachments : "(org_id, id) - (org_id, file_id)"
    tasks ||--o{ task_labels : "(id) - (task_id)"
    labels ||--o{ task_labels : "(id) - (label_id)"
    orgs ||--o{ task_comments : "(id) - (org_id)"
    tasks ||--o{ task_comments : "(id) - (task_id)"
    users ||--o{ task_comments : "(id) - (user_id)"
    orgs ||--o{ task_reactions : "(id) - (org_id)"
    tasks ||--o{ task_reactions : "(id) - (task_id)"
    users ||--o{ task_reactions : "(id) - (user_id)"
    orgs ||--o{ task_activity : "(id) - (org_id)"
    tasks ||--o{ task_activity : "(id) - (task_id)"
    users |o--o{ task_activity : "(id) - (user_id)"
    tasks ||--o{ task_relationships : "(id) - (source_task_id)"
    tasks ||--o{ task_relationships : "(id) - (target_task_id)"
    tasks ||--o{ task_watchers : "(id) - (task_id)"
    users ||--o{ task_watchers : "(id) - (user_id)"
    tasks ||--o{ task_assignees : "(id) - (task_id)"
    users ||--o{ task_assignees : "(id) - (user_id)"
    users ||--o{ notifications : "(id) - (user_id)"
    projects |o--o{ saved_views : "(id) - (project_id)"
    users ||--o{ saved_views : "(id) - (user_id)"
    users ||--o{ favorites : "(id) - (user_id)"
    orgs ||--o{ capability_provider_snapshots : "(id) - (org_id)"
    orgs ||--o{ app_runs : "(id) - (org_id)"
    capability_provider_snapshots ||--o{ app_runs : "(org_id, id) - (org_id, provider_snapshot_id)"
    app_runs ||--o{ app_runs : "(org_id, id) - (org_id, root_run_id)"
    app_runs |o--o{ app_runs : "(org_id, id) - (org_id, parent_run_id)"
    app_installations |o--o{ app_runs : "(org_id, id) - (org_id, origin_app_installation_id)"
    app_versions |o--o{ app_runs : "(org_id, installation_id, id) - (org_id, origin_app_installation_id, origin_app_version_id)"
    app_grant_snapshots |o--o{ app_runs : "(org_id, app_installation_id, app_version_id, id) - (org_id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id)"
    app_action_bindings |o--o{ app_runs : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, action_key, provider_kind, mcp_connection_id, operation_name, provider_snapshot_id) - (org_id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id, origin_app_binding_key, provider_kind, provider_instance_id, operation_name, provider_snapshot_id)"
    app_runtime_bindings |o--o{ app_runs : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, risk_class, review_requirement, retry_class, retention_class) - (org_id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id, origin_runtime_binding_id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, risk_class, review_requirement, retry_class, retention_class)"
    app_resource_bindings |o--o{ app_runs : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, risk_class, review_requirement, review_scope, retry_class, retention_class) - (org_id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id, origin_resource_binding_id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, risk_class, review_requirement, review_scope, retry_class, retention_class)"
    app_native_bindings |o--o{ app_runs : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, owner_user_id, risk_class, review_requirement, review_scope, retry_class, retention_class) - (org_id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id, origin_native_binding_id, provider_kind, provider_instance_id, operation_name, provider_snapshot_id, execution_actor_id, risk_class, review_requirement, review_scope, retry_class, retention_class)"
    app_automation_definitions |o--o{ app_runs : "(org_id, id) - (org_id, origin_app_automation_definition_id)"
    app_public_ingress |o--o| app_runs : "(org_id, endpoint_id, id) - (org_id, origin_public_endpoint_id, origin_public_ingress_id)"
    app_automation_fires |o--o| app_runs : "(org_id, definition_id, id) - (org_id, origin_app_automation_definition_id, origin_app_automation_fire_id)"
    app_runs ||--o| app_run_attempts : "(org_id, id) - (org_id, run_id)"
    app_runtime_sessions |o--o{ app_run_attempts : "(org_id, runtime_binding_id, id, session_epoch, runtime_epoch) - (org_id, runtime_binding_id, runtime_session_id, runtime_session_epoch, runtime_epoch)"
    app_runtime_sessions |o--o{ app_run_attempts : "(org_id, resource_binding_id, id, session_epoch, runtime_epoch) - (org_id, resource_binding_id, runtime_session_id, runtime_session_epoch, runtime_epoch)"
    app_runs |o--o{ app_run_attempts : "(org_id, id, origin_resource_binding_id) - (org_id, run_id, resource_binding_id)"
    app_runs ||--o{ app_run_secret_payloads : "(org_id, id) - (org_id, run_id)"
    app_run_attempts |o--o{ app_run_secret_payloads : "(org_id, run_id, id) - (org_id, run_id, attempt_id)"
    app_runs ||--o{ app_run_events : "(org_id, id) - (org_id, run_id)"
    app_runs ||--o{ app_run_receipts : "(org_id, id) - (org_id, run_id)"
    app_run_attempts |o--o{ app_run_receipts : "(org_id, run_id, id) - (org_id, run_id, attempt_id)"
    users ||--o{ agent_actions : "(id) - (user_id)"
    users |o--o{ agent_actions : "(id) - (approved_by_user_id)"
    app_runs |o--o| agent_actions : "(org_id, id) - (org_id, app_run_id)"
    users ||--o{ attention_items : "(id) - (user_id)"
    attention_items ||--o{ attention_events : "(id) - (attention_item_id)"
    users ||--o{ attention_events : "(id) - (user_id)"
    users |o--o{ attention_events : "(id) - (actor_user_id)"
    attention_items ||--o{ attention_deliveries : "(id) - (attention_item_id)"
    users ||--o{ attention_deliveries : "(id) - (user_id)"
    users ||--o{ web_push_subscriptions : "(id) - (user_id)"
    agent_actions ||--o{ agent_action_approvers : "(id) - (action_id)"
    users ||--o{ agent_action_approvers : "(id) - (user_id)"
    orgs ||--o{ work_intents : "(id) - (org_id)"
    spaces |o--o{ work_intents : "(id) - (space_id)"
    messages |o--o{ work_intents : "(id) - (source_message_id)"
    users |o--o{ work_intents : "(id) - (source_user_id)"
    agent_employees |o--o{ work_intents : "(id) - (agent_employee_id)"
    tasks |o--o{ work_intents : "(id) - (converted_task_id)"
    users |o--o{ work_intents : "(id) - (converted_by)"
    users |o--o{ work_intents : "(id) - (dismissed_by)"
    orgs ||--o{ message_observations : "(id) - (org_id)"
    messages ||--o{ message_observations : "(id) - (message_id)"
    spaces |o--o{ message_observations : "(id) - (space_id)"
    users |o--o{ message_observations : "(id) - (user_id)"
    users |o--o{ skills : "(id) - (created_by)"
    users |o--o{ task_templates : "(id) - (created_by)"
    orgs ||--o{ module_installations : "(id) - (org_id)"
    users |o--o{ module_installations : "(id) - (installed_by_user_id)"
    module_installations ||--o| module_versions : "(org_id, id) - (org_id, installation_id)"
    orgs ||--o{ app_installations : "(id) - (org_id)"
    users |o--o{ app_installations : "(id) - (installed_by_user_id)"
    app_installations ||--o| app_versions : "(org_id, id) - (org_id, installation_id)"
    app_installations ||--o{ app_module_bindings : "(org_id, id) - (org_id, app_installation_id)"
    app_versions ||--o{ app_module_bindings : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    module_installations ||--o{ app_module_bindings : "(org_id, id) - (org_id, module_installation_id)"
    module_versions ||--o{ app_module_bindings : "(org_id, installation_id, id) - (org_id, module_installation_id, module_version_id)"
    app_installations ||--o{ app_grant_snapshots : "(org_id, id, app_id) - (org_id, app_installation_id, app_id)"
    app_versions ||--o{ app_grant_snapshots : "(org_id, installation_id, id, version, manifest_digest, package_digest) - (org_id, app_installation_id, app_version_id, app_version, manifest_digest, package_digest)"
    app_grant_snapshots |o--o{ app_grant_snapshots : "(org_id, app_installation_id, app_version_id, id) - (org_id, app_installation_id, app_version_id, requested_snapshot_id)"
    app_grant_snapshots |o--o| app_grant_snapshots : "(org_id, app_installation_id, id) - (org_id, app_installation_id, supersedes_snapshot_id)"
    app_versions ||--o{ app_dependency_locks : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_dependency_locks : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    app_installations ||--o{ app_dependency_locks : "(org_id, id, app_id) - (org_id, dependency_installation_id, required_app_id)"
    app_versions ||--o{ app_dependency_locks : "(org_id, installation_id, id, version, manifest_digest, package_digest) - (org_id, dependency_installation_id, dependency_version_id, required_version, dependency_manifest_digest, dependency_package_digest)"
    app_versions ||--o{ app_action_bindings : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_action_bindings : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    mcp_connections ||--o{ app_action_bindings : "(org_id, id) - (org_id, mcp_connection_id)"
    capability_provider_snapshots ||--o{ app_action_bindings : "(org_id, provider_kind, provider_instance_id, id) - (org_id, provider_kind, mcp_connection_id, provider_snapshot_id)"
    app_versions ||--o{ app_owner_connection_admissions : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    org_members ||--o{ app_owner_connection_admissions : "(org_id, user_id) - (org_id, operator_user_id)"
    app_grant_snapshots ||--o{ app_owner_connection_admissions : "(org_id, app_installation_id, app_version_id, id) - (org_id, app_installation_id, app_version_id, grant_snapshot_id)"
    org_members ||--o{ app_owner_connection_admissions : "(org_id, user_id) - (org_id, reviewed_by_user_id)"
    app_owner_connection_admissions |o--o{ app_runtime_registrations : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, operator_user_id, id, review_digest) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, operator_user_id, owner_connection_admission_id, owner_connection_admission_digest)"
    app_versions ||--o{ app_runtime_registrations : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_runtime_registrations : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    org_members ||--o{ app_runtime_registrations : "(org_id, user_id) - (org_id, operator_user_id)"
    app_runtime_registrations ||--o{ app_runtime_bindings : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, contract_version) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, runtime_registration_id, registration_contract_version)"
    capability_provider_snapshots ||--o{ app_runtime_bindings : "(org_id, provider_kind, provider_instance_id, id) - (org_id, provider_kind, provider_instance_id, provider_snapshot_id)"
    app_grant_snapshots ||--o{ app_native_bindings : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    capability_provider_snapshots ||--o{ app_native_bindings : "(org_id, provider_kind, provider_instance_id, id) - (org_id, provider_kind, provider_instance_id, provider_snapshot_id)"
    org_members ||--o{ app_native_bindings : "(org_id, user_id) - (org_id, owner_user_id)"
    org_members ||--o{ app_native_bindings : "(org_id, user_id) - (org_id, stage_manager_user_id)"
    app_grant_snapshots ||--o{ app_resource_bindings : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    app_runtime_registrations ||--o{ app_resource_bindings : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, contract_version) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, runtime_registration_id, registration_contract_version)"
    capability_provider_snapshots ||--o{ app_resource_bindings : "(org_id, provider_kind, provider_instance_id, id) - (org_id, provider_kind, provider_instance_id, provider_snapshot_id)"
    org_members ||--o{ app_resource_bindings : "(org_id, user_id) - (org_id, owner_user_id)"
    org_members |o--o{ app_resource_bindings : "(org_id, user_id) - (org_id, reviewed_by_user_id)"
    app_runtime_bindings |o--o{ app_runtime_sessions : "(org_id, runtime_registration_id, id) - (org_id, runtime_registration_id, runtime_binding_id)"
    app_resource_bindings |o--o{ app_runtime_sessions : "(org_id, runtime_registration_id, id) - (org_id, runtime_registration_id, resource_binding_id)"
    org_members ||--o{ app_runtime_sessions : "(org_id, user_id) - (org_id, operator_user_id)"
    app_resource_bindings ||--o| app_sync_checkpoints : "(org_id, id) - (org_id, resource_binding_id)"
    app_runs ||--o{ app_sync_intents : "(org_id, id, origin_app_installation_id, origin_app_version_id, origin_app_grant_snapshot_id, origin_resource_binding_id, provider_snapshot_id) - (org_id, run_id, app_installation_id, app_version_id, grant_snapshot_id, resource_binding_id, provider_snapshot_id)"
    app_sync_checkpoints ||--o{ app_sync_intents : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_resource_bindings ||--o{ app_sync_intents : "(org_id, id, owner_user_id) - (org_id, resource_binding_id, owner_user_id)"
    app_resource_bindings ||--o{ app_sync_intents : "(org_id, id, owner_user_id, descriptor_digest) - (org_id, resource_binding_id, owner_user_id, descriptor_digest)"
    app_sync_checkpoints ||--o{ app_resource_projections : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_sync_checkpoints ||--o{ app_attachment_stages : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_run_attempts ||--o{ app_attachment_stages : "(org_id, run_id, id) - (org_id, run_id, attempt_id)"
    app_resource_projections |o--o{ app_attachment_stages : "(org_id, id, checkpoint_id, resource_binding_id) - (org_id, projection_id, checkpoint_id, resource_binding_id)"
    app_installations ||--o{ app_automation_definitions : "(org_id, id) - (org_id, app_installation_id)"
    app_versions ||--o{ app_automation_definitions : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_automation_definitions : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    app_action_bindings ||--o{ app_automation_definitions : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, action_key, id) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, action_key, action_binding_id)"
    app_automation_definitions ||--o| app_automation_fires : "(org_id, id) - (org_id, definition_id)"
    org_members ||--o{ app_developer_pairings : "(org_id, user_id) - (org_id, created_by_user_id)"
    module_installations ||--o{ module_records : "(org_id, id) - (org_id, installation_id)"
    module_versions ||--o{ module_records : "(org_id, installation_id, id) - (org_id, installation_id, validated_version_id)"
    module_installations ||--o{ module_record_relations : "(org_id, id) - (org_id, installation_id)"
    module_records ||--o{ module_record_relations : "(org_id, installation_id, id) - (org_id, installation_id, source_record_id)"
    module_records ||--o{ module_record_relations : "(org_id, installation_id, id) - (org_id, installation_id, target_record_id)"
    module_records ||--o{ module_record_merges : "(org_id, installation_id, id) - (org_id, installation_id, source_record_id)"
    module_records ||--o{ module_record_merges : "(org_id, installation_id, id) - (org_id, installation_id, target_record_id)"
    users ||--o{ module_saved_views : "(id) - (owner_user_id)"
    module_installations ||--o{ module_saved_views : "(org_id, id) - (org_id, installation_id)"
    org_members ||--o{ module_saved_views : "(org_id, user_id) - (org_id, owner_user_id)"
    agent_actions |o--o{ module_mutation_receipts : "(id) - (agent_action_id)"
    module_installations ||--o{ module_mutation_receipts : "(org_id, id) - (org_id, installation_id)"
    module_records ||--o{ module_mutation_receipts : "(org_id, installation_id, id) - (org_id, installation_id, record_id)"
    orgs ||--o{ resource_relation_sets : "(id) - (org_id)"
    resource_relation_sets ||--o{ resource_relation_edges : "(org_id, id) - (org_id, relation_set_id)"
    resource_relation_sets ||--o{ resource_relation_receipts : "(org_id, id) - (org_id, relation_set_id)"
    orgs ||--o{ agent_webhooks : "(id) - (org_id)"
    agent_employees ||--o{ agent_webhooks : "(id) - (agent_employee_id)"
    users |o--o{ agent_webhooks : "(id) - (created_by)"
    users ||--o{ triggers : "(id) - (created_by)"
    users ||--o{ connected_accounts : "(id) - (user_id)"
    users |o--o{ events : "(id) - (user_id)"
    connected_accounts |o--o{ events : "(id) - (connected_account_id)"
    users ||--o{ ics_subscriptions : "(id) - (user_id)"
    users ||--o{ reminders : "(id) - (user_id)"
    messages |o--o{ reminders : "(id) - (source_message_id)"
    messages ||--o{ pinned_messages : "(id) - (message_id)"
    spaces ||--o{ pinned_messages : "(id) - (space_id)"
    users ||--o{ pinned_messages : "(id) - (pinned_by)"
    users ||--o{ scheduled_messages : "(id) - (user_id)"
    spaces ||--o{ scheduled_messages : "(id) - (space_id)"
    spaces ||--o| canvases : "(id) - (space_id)"
    users |o--o{ canvases : "(id) - (last_edited_by)"
    users ||--o{ message_bookmarks : "(id) - (user_id)"
    messages ||--o{ message_bookmarks : "(id) - (message_id)"
    spaces ||--o{ message_bookmarks : "(id) - (space_id)"
    users ||--o{ user_groups : "(id) - (created_by)"
    user_groups ||--o{ user_group_members : "(id) - (group_id)"
    users ||--o{ user_group_members : "(id) - (user_id)"
    users |o--o{ teams : "(id) - (lead_user_id)"
    spaces |o--o{ teams : "(id) - (default_space_id)"
    users |o--o{ teams : "(id) - (created_by)"
    teams ||--o{ team_members : "(id) - (team_id)"
    users ||--o{ team_members : "(id) - (user_id)"
    teams ||--o{ team_resources : "(id) - (team_id)"
    users |o--o{ team_resources : "(id) - (created_by)"
    teams ||--o{ team_dashboard_snapshots : "(id) - (team_id)"
    users ||--o{ custom_emoji : "(id) - (uploaded_by)"
    users ||--o{ workflow_rules : "(id) - (created_by)"
    workflow_rules ||--o{ workflow_runs : "(id) - (rule_id)"
    messages |o--o{ workflow_runs : "(id) - (triggered_by_message_id)"
    users |o--o{ workflow_runs : "(id) - (triggered_by_user_id)"
    users |o--o{ automation_runs : "(id) - (user_id)"
    users ||--o| onboarding_state : "(id) - (user_id)"
    users ||--o{ agent_memory : "(id) - (user_id)"
    users ||--o{ cross_references : "(id) - (created_by)"
    orgs ||--o{ duplicate_flags : "(id) - (org_id)"
    tasks ||--o{ duplicate_flags : "(id) - (task_a_id)"
    tasks ||--o{ duplicate_flags : "(id) - (task_b_id)"
    users ||--o{ agent_nudges : "(id) - (user_id)"
    tasks ||--o{ agent_nudges : "(id) - (task_id)"
    users ||--o{ meeting_briefs : "(id) - (user_id)"
    events ||--o{ meeting_briefs : "(id) - (event_id)"
    users ||--o{ people_interactions : "(id) - (user_a_id)"
    users ||--o{ people_interactions : "(id) - (user_b_id)"
    users ||--o{ people_expertise : "(id) - (user_id)"
    users ||--o{ people_influence : "(id) - (user_id)"
    users ||--o{ people_patterns : "(id) - (user_id)"
    users ||--o{ people_relationships : "(id) - (user_a_id)"
    users ||--o{ people_relationships : "(id) - (user_b_id)"
    users ||--o{ oneone_preps : "(id) - (manager_id)"
    users ||--o{ oneone_preps : "(id) - (report_id)"
    users ||--o{ burnout_alerts : "(id) - (user_id)"
    users ||--o{ burnout_alerts : "(id) - (alerted_to)"
    users ||--o{ note_folders : "(id) - (user_id)"
    users ||--o{ notes : "(id) - (user_id)"
    note_folders |o--o{ notes : "(id) - (folder_id)"
    spaces |o--o{ notes : "(id) - (visibility_space_id)"
    notes ||--o{ note_versions : "(id) - (note_id)"
    notes ||--o{ note_shares : "(id) - (note_id)"
    users ||--o{ note_shares : "(id) - (shared_with_user_id)"
    tags ||--o{ entity_tags : "(id) - (tag_id)"
    spaces |o--o{ clips : "(id) - (space_id)"
    messages |o--o{ clips : "(id) - (message_id)"
    users ||--o{ clips : "(id) - (created_by)"
    users ||--o{ manager_settings : "(id) - (user_id)"
    users ||--o{ thread_reads : "(id) - (user_id)"
    messages ||--o{ thread_reads : "(id) - (parent_message_id)"
    messages ||--o{ message_versions : "(id) - (message_id)"
    spaces |o--o{ wiki_pages : "(id) - (space_id)"
    spaces |o--o{ wiki_pages : "(id) - (origin_space_id)"
    messages |o--o{ wiki_pages : "(id) - (origin_message_id)"
    users |o--o{ wiki_pages : "(id) - (origin_user_id)"
    users |o--o{ wiki_pages : "(id) - (user_id)"
    wiki_pages ||--o{ wiki_links : "(id) - (source_page_id)"
    wiki_pages ||--o{ wiki_links : "(id) - (target_page_id)"
    orgs |o--o{ wiki_citations : "(id) - (org_id)"
    wiki_pages ||--o{ wiki_citations : "(id) - (page_id)"
    spaces |o--o{ wiki_citations : "(id) - (source_space_id)"
    users |o--o{ wiki_citations : "(id) - (source_user_id)"
    agent_employees ||--o{ wiki_memory_syncs : "(id) - (agent_employee_id)"
    wiki_pages ||--o{ wiki_memory_syncs : "(id) - (page_id)"
    wiki_pages |o--o{ wiki_ops_log : "(id) - (page_id)"
    wiki_pages ||--o{ wiki_page_versions : "(id) - (page_id)"
    users ||--o{ mcp_connections : "(id) - (created_by)"
    mcp_connections ||--o{ mcp_tool_overrides : "(id) - (mcp_connection_id)"
    mcp_connections ||--o{ mcp_tool_overrides : "(org_id, id) - (org_id, mcp_connection_id)"
    users ||--o{ agent_employees : "(id) - (user_id)"
    users ||--o{ agent_employees : "(id) - (created_by)"
    agent_employees ||--o{ agent_certification_challenges : "(id) - (employee_id)"
    agent_employees ||--o{ agent_mcp_call_audit : "(id) - (employee_id)"
    users |o--o{ mcp_tokens : "(id) - (user_id)"
    agent_employees |o--o{ mcp_tokens : "(id) - (agent_employee_id)"
    users |o--o{ mcp_tokens : "(id) - (created_by)"
    agent_employees ||--o{ agent_channel_connections : "(id) - (agent_employee_id)"
    agent_employees ||--o{ agent_channel_tokens : "(id) - (agent_employee_id)"
    users |o--o{ agent_channel_tokens : "(id) - (created_by)"
    agent_employees ||--o{ agent_channel_events : "(id) - (agent_employee_id)"
    spaces |o--o{ agent_channel_events : "(id) - (space_id)"
    messages |o--o{ agent_channel_events : "(id) - (thread_id)"
    users |o--o{ agent_channel_events : "(id) - (actor_user_id)"
    agent_employees ||--o{ agent_channel_cursors : "(id) - (agent_employee_id)"
    agent_channel_connections |o--o{ agent_channel_cursors : "(id) - (connection_id)"
    agent_employees ||--o{ agent_channel_sessions : "(id) - (agent_employee_id)"
    agent_employees ||--o{ agent_channel_delivery_attempts : "(id) - (agent_employee_id)"
    agent_channel_events |o--o{ agent_channel_delivery_attempts : "(id) - (event_id)"
    users ||--o{ oauth_authorization_codes : "(id) - (user_id)"
    users ||--o{ oauth_grants : "(id) - (user_id)"
    oauth_grants ||--o{ oauth_access_tokens : "(id) - (grant_id)"
    users ||--o{ oauth_access_tokens : "(id) - (user_id)"
    oauth_grants ||--o{ oauth_refresh_tokens : "(id) - (grant_id)"
    users |o--o{ oauth_audit_events : "(id) - (user_id)"
    agent_employees ||--o{ agent_employee_skills : "(id) - (agent_employee_id)"
    skills ||--o{ agent_employee_skills : "(id) - (skill_id)"
    users ||--o{ agent_plans : "(id) - (user_id)"
    users ||--o{ api_keys : "(id) - (created_by)"
    orgs |o--o{ agent_employee_templates : "(id) - (org_id)"
    users |o--o{ agent_employee_templates : "(id) - (created_by)"
    agent_employees ||--o{ agent_cooperative_log : "(id) - (employee_id)"
    agent_employees ||--o{ agent_session_turns : "(id) - (employee_id)"
    agent_employees ||--o{ agent_heartbeat_turns : "(id) - (agent_employee_id)"
    agent_actions ||--o{ action_receipts : "(id) - (action_id)"
    agent_employees |o--o{ action_receipts : "(id) - (employee_id)"
    users |o--o{ action_receipts : "(id) - (approver_id)"
    spaces ||--o{ space_memory : "(id) - (space_id)"
    agent_employees |o--o{ space_memory : "(id) - (updated_by_employee_id)"
    users |o--o{ integrations : "(id) - (connected_by)"
    messages ||--o{ message_classifications : "(id) - (message_id)"
    users ||--o{ web_sessions : "(id) - (user_id)"
    orgs ||--o{ web_sessions : "(id) - (org_id)"
    org_members ||--o{ app_experience_consent_grants : "(org_id, user_id) - (org_id, owner_user_id)"
    app_versions ||--o{ app_experience_consent_grants : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    org_members ||--o{ app_runtime_agent_policies : "(org_id, user_id) - (org_id, owner_user_id)"
    app_runtime_bindings ||--o{ app_runtime_agent_policies : "(org_id, id) - (org_id, runtime_binding_id)"
    org_members ||--o{ app_experience_sessions : "(org_id, user_id) - (org_id, user_id)"
    web_sessions ||--o{ app_experience_sessions : "(id) - (web_session_id)"
    app_experience_consent_grants |o--o{ app_experience_sessions : "(org_id, id, owner_user_id) - (org_id, consent_grant_id, user_id)"
    app_versions ||--o{ app_experience_sessions : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_experience_sessions : "(org_id, app_installation_id, app_version_id, id, snapshot_kind) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, grant_snapshot_kind)"
    app_experience_sessions ||--o{ app_experience_resource_exposures : "(org_id, id, user_id, web_session_id) - (org_id, experience_session_id, owner_user_id, web_session_id)"
    users ||--o{ app_action_batches : "(id) - (owner_user_id)"
    app_runtime_bindings ||--o{ app_action_batches : "(org_id, id) - (org_id, runtime_binding_id)"
    app_experience_consent_grants ||--o{ app_action_batches : "(org_id, id, owner_user_id) - (org_id, consent_grant_id, owner_user_id)"
    app_action_batches ||--o{ app_action_batch_items : "(org_id, id) - (org_id, batch_id)"
    app_runs ||--o| app_action_batch_items : "(org_id, id) - (org_id, run_id)"
    app_installations ||--o{ app_private_files : "(org_id, id) - (org_id, installation_id)"
    users ||--o{ app_private_files : "(id) - (owner_user_id)"
    app_installations ||--o{ app_action_payloads : "(org_id, id) - (org_id, installation_id)"
    users ||--o{ app_action_payloads : "(id) - (owner_user_id)"
    app_action_payloads ||--o{ app_action_payload_files : "(org_id, owner_user_id, id) - (org_id, owner_user_id, snapshot_id)"
    app_private_files ||--o{ app_action_payload_files : "(org_id, owner_user_id, id) - (org_id, owner_user_id, file_id)"
    app_action_payloads |o--o{ app_run_human_authorizations : "(org_id, owner_user_id, id, snapshot_digest) - (org_id, owner_user_id, snapshot_id, snapshot_digest)"
    app_runs ||--o| app_run_human_authorizations : "(org_id, id) - (org_id, run_id)"
    app_experience_consent_grants |o--o{ app_run_human_authorizations : "(org_id, id, owner_user_id) - (org_id, consent_grant_id, owner_user_id)"
    app_experience_resource_exposures |o--o{ app_run_human_authorizations : "(org_id, id, owner_user_id) - (org_id, exposure_id, owner_user_id)"
    app_experience_resource_exposures ||--o{ app_experience_resource_exposure_resources : "(org_id, id) - (org_id, exposure_id)"
    app_resource_bindings ||--o{ app_experience_resource_exposure_resources : "(org_id, runtime_registration_id, id) - (org_id, runtime_registration_id, resource_binding_id)"
    app_versions ||--o{ app_public_endpoints : "(org_id, installation_id, id) - (org_id, app_installation_id, app_version_id)"
    app_grant_snapshots ||--o{ app_public_endpoints : "(org_id, app_installation_id, app_version_id, id) - (org_id, app_installation_id, app_version_id, grant_snapshot_id)"
    module_installations ||--o{ app_public_endpoints : "(org_id, id) - (org_id, module_installation_id)"
    app_runtime_bindings |o--o{ app_public_endpoints : "(org_id, id) - (org_id, runtime_binding_id)"
    app_native_bindings |o--o{ app_public_endpoints : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, owner_user_id) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, native_binding_id, approver_user_id)"
    app_native_bindings |o--o{ app_public_endpoints : "(org_id, app_installation_id, app_version_id, grant_snapshot_id, id, owner_user_id) - (org_id, app_installation_id, app_version_id, grant_snapshot_id, cancel_native_binding_id, approver_user_id)"
    org_members |o--o{ app_public_endpoints : "(org_id, user_id) - (org_id, approver_user_id)"
    app_public_endpoints ||--o{ app_public_ingress : "(org_id, id) - (org_id, endpoint_id)"
    app_public_endpoints ||--o{ app_public_hmac_keys : "(org_id, id) - (org_id, endpoint_id)"
    app_public_hmac_keys ||--o{ app_public_hmac_nonces : "(org_id, endpoint_id, id) - (org_id, endpoint_id, key_id)"
    app_public_ingress ||--o{ app_canonical_claims : "(org_id, endpoint_id, id) - (org_id, endpoint_id, ingress_id)"
    module_records ||--o{ app_canonical_claims : "(org_id, installation_id, id) - (org_id, provider_instance_id, resource_id)"
    app_canonical_claims ||--o{ app_public_cancellations : "(org_id, endpoint_id, id) - (org_id, endpoint_id, claim_id)"
    app_public_endpoints ||--o{ app_public_cancellations : "(org_id, app_installation_id, id) - (org_id, app_installation_id, endpoint_id)"
    app_runs |o--o{ app_public_cancellations : "(org_id, id) - (org_id, original_run_id)"
    app_public_cancellations ||--o{ app_public_cancellation_selections : "(org_id, app_installation_id, id, original_run_id) - (org_id, app_installation_id, cancellation_id, original_run_id)"
    app_native_bindings ||--o{ app_public_cancellation_selections : "(org_id, app_installation_id, id, owner_user_id) - (org_id, app_installation_id, native_binding_id, owner_user_id)"
    app_runs |o--o| app_public_cancellation_selections : "(org_id, id) - (org_id, cancel_run_id)"
    app_resource_bindings ||--o{ app_resource_access_grants : "(org_id, id, owner_user_id) - (org_id, resource_binding_id, owner_user_id)"
    app_sync_checkpoints ||--o{ app_resource_access_grants : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_resource_projections ||--o{ app_resource_access_grants : "(org_id, id) - (org_id, projection_id)"
    org_members ||--o{ app_resource_access_grants : "(org_id, user_id) - (org_id, recipient_user_id)"
    app_resource_bindings ||--o{ app_private_mcp_grants : "(org_id, id, owner_user_id) - (org_id, resource_binding_id, owner_user_id)"
    app_sync_checkpoints ||--o{ app_private_mcp_grants : "(org_id, id, resource_binding_id) - (org_id, checkpoint_id, resource_binding_id)"
    app_resource_projections ||--o{ app_private_mcp_grants : "(org_id, id) - (org_id, projection_id)"
    org_members ||--o{ app_private_mcp_grants : "(org_id, user_id) - (org_id, subject_user_id)"
    mcp_tokens ||--o{ app_private_mcp_grants : "(org_id, id) - (org_id, mcp_token_id)"
    app_installations ||--o{ app_private_state_records : "(org_id, id) - (org_id, installation_id)"
    users ||--o{ app_private_state_records : "(id) - (owner_user_id)"
    email_accounts ||--o| email_checkpoints : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_messages : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_messages ||--o{ email_memberships : "(org_id, owner_user_id, account_id, logical_message_id) - (org_id, owner_user_id, account_id, logical_message_id)"
    email_accounts ||--o{ email_sync_receipts : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_provider_states : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o| email_reading_preferences : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o| email_imap_credentials : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_drafts : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_send_operations : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_drafts ||--o{ email_send_operations : "(org_id, owner_user_id, account_id, draft_id) - (org_id, owner_user_id, account_id, draft_id)"
    email_accounts ||--o{ email_provider_draft_mirrors : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_drafts ||--o| email_provider_draft_mirrors : "(org_id, owner_user_id, account_id, draft_id) - (org_id, owner_user_id, account_id, draft_id)"
    email_accounts ||--o| email_account_consents : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_native_connection_owners ||--o{ email_native_connection_requests : "(org_id, owner_user_id) - (org_id, owner_user_id)"
    email_accounts ||--o| email_native_account_connections : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_operations : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_drafts ||--o{ email_attachments : "(org_id, owner_user_id, account_id, draft_id) - (org_id, owner_user_id, account_id, draft_id)"
    email_drafts ||--o{ email_outbox : "(org_id, owner_user_id, account_id, draft_id) - (org_id, owner_user_id, account_id, draft_id)"
    email_outbox ||--o{ email_outbox_events : "(org_id, owner_user_id, account_id, outbox_id) - (org_id, owner_user_id, account_id, outbox_id)"
    email_accounts ||--o{ email_templates : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o| email_signatures : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_messages ||--o{ email_imap_aliases : "(org_id, owner_user_id, account_id, logical_message_id) - (org_id, owner_user_id, account_id, logical_message_id)"
    email_messages ||--o| email_imap_move_reservations : "(org_id, owner_user_id, account_id, logical_message_id) - (org_id, owner_user_id, account_id, logical_message_id)"
    email_operations ||--o{ email_imap_move_reservations : "(org_id, owner_user_id, account_id, operation_id) - (org_id, owner_user_id, account_id, operation_id)"
    email_accounts ||--o{ email_timers : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_presentation_cache : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_notification_events : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o| email_notification_states : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_calendar_invitations : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    events |o--o{ email_calendar_invitations : "(id) - (event_id)"
    email_calendar_invitations ||--o{ email_calendar_invitation_actions : "(org_id, owner_user_id, account_id, invitation_id) - (org_id, owner_user_id, account_id, invitation_id)"
    email_outbox |o--o{ email_calendar_invitation_actions : "(org_id, owner_user_id, account_id, outbox_id) - (org_id, owner_user_id, account_id, outbox_id)"
    orgs ||--o{ email_context_transfers : "(id) - (org_id)"
    users ||--o{ email_context_transfers : "(id) - (owner_user_id)"
    email_accounts ||--o{ email_context_transfers : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_rules : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_rule_receipts : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_rules ||--o{ email_rule_receipts : "(org_id, owner_user_id, account_id, rule_id) - (org_id, owner_user_id, account_id, rule_id)"
    email_accounts ||--o{ email_agent_grants : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o| email_knowledge_policies : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_knowledge_observations : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_knowledge_sources : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    email_accounts ||--o{ email_knowledge_threads : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    orgs ||--o{ app_integration_capture_policies : "(id) - (org_id)"
    users ||--o{ app_integration_capture_policies : "(id) - (owner_user_id)"
    app_installations ||--o{ app_integration_capture_policies : "(org_id, id) - (org_id, installation_id)"
    app_integration_capture_policies ||--o{ app_integration_capture_claims : "(org_id, id) - (org_id, policy_id)"
    wiki_pages ||--o{ app_integration_capture_threads : "(id) - (page_id)"
    app_integration_capture_claims ||--o| app_integration_capture_sources : "(org_id, id) - (org_id, claim_id)"
    wiki_pages ||--o{ app_integration_capture_sources : "(id) - (page_id)"
    orgs ||--o{ app_integration_destination_receipts : "(id) - (org_id)"
    users ||--o{ app_integration_destination_receipts : "(id) - (owner_user_id)"
    orgs ||--o{ app_integration_notification_policies : "(id) - (org_id)"
    users ||--o{ app_integration_notification_policies : "(id) - (owner_user_id)"
    orgs ||--o{ app_restore_access_reviews : "(id) - (org_id)"
    users ||--o{ app_restore_access_reviews : "(id) - (owner_user_id)"
    email_accounts ||--o| native_email_adoptions : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    native_email_adoptions ||--o| native_email_adoption_destinations : "(org_id, owner_user_id, account_id) - (org_id, owner_user_id, account_id)"
    orgs {
        text id PK
        text name
        text slug
        text logo_url
        text timezone
        trust_level trust_level
        text agent_name
        boolean agent_enabled
        jsonb ai_config "Per-org AI provider config (BYOK). Read via apps/api/src/lib/org-ai-config.ts. Schema documented in packages/db/drizzle/0061_org_ai_config.sql."
        timestamp created_at
        timestamp updated_at
    }
    users {
        text id PK
        text email
        text name
        user_kind kind
        boolean is_agent
        text agent_employee_id
        text avatar_url
        text title
        text profile_summary
        text expertise_tags
        text timezone
        text status_emoji
        text status_text
        timestamp status_expires_at
        text password_hash
        integer password_version
        boolean email_verified
        timestamp last_seen_at
        text notification_keywords
        jsonb notification_preferences
        boolean show_read_receipts
        text ics_publish_token "Per-user secret token for the outbound ICS feed. Lazily generated when the user first opens Settings → Calendar. See migration 0062."
        timestamp created_at
        timestamp updated_at
    }
    native_reference_states {
        text id PK
        text org_id
        text source_kind
        text source_id
        text content_hash
        integer revision
        jsonb current_refs
        jsonb published_person_ids
        integer publication_revision
        boolean is_deleted
        timestamp created_at
        timestamp updated_at
    }
    native_mention_deliveries {
        text id PK
        text org_id FK
        text source_state_id FK
        integer publication_revision
        text recipient_user_id FK
        text actor_user_id FK
        text status
        text attention_id
        text reason
        timestamp created_at
        timestamp updated_at
    }
    native_create_requests {
        text id PK
        text org_id FK
        text user_id FK
        text operation
        text request_hash
        text resource_id
        timestamp created_at
    }
    org_members {
        text id PK
        text org_id
        text user_id FK
        org_role role
        boolean is_active
        integer app_run_authorization_version
        timestamp joined_at
        timestamp created_at
        timestamp updated_at
    }
    invites {
        text id PK
        text org_id
        text email
        text token
        text type
        text invited_by FK
        text accepted_by FK
        timestamp accepted_at
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
    }
    spaces {
        text id PK
        text org_id
        text name
        text description
        text topic
        space_type type
        boolean is_default
        boolean is_archived
        boolean agent_enabled
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    space_members {
        text id PK
        text space_id FK
        text user_id FK
        boolean is_muted
        text notification_level
        text last_read_message_id
        timestamp last_read_at
        timestamp joined_at
    }
    messages {
        text id PK
        text org_id
        text space_id FK
        text user_id FK
        text content
        text parent_id
        boolean is_pinned
        boolean is_deleted
        timestamp edited_at
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    app_private_defty_seals {
        text id PK
        text org_id FK
        text space_id FK
        text owner_user_id FK
        text defty_user_id FK
        timestamptz created_at
    }
    app_private_defty_grants {
        text id PK
        text org_id FK
        text seal_id FK
        text owner_user_id FK
        text resource_binding_id FK
        text checkpoint_id FK
        text projection_id FK
        text review_digest
        jsonb snapshot
        timestamptz accepted_at
        timestamptz expires_at
        timestamptz revoked_at
        text active_request_id
        text active_prompt_digest
    }
    reactions {
        text id PK
        text message_id FK
        text user_id FK
        text emoji
        timestamp created_at
        timestamp updated_at
    }
    files {
        text id PK
        text org_id
        text uploaded_by FK
        text filename
        text mime_type
        integer size_bytes
        text storage_key
        text thumbnail_key
        text detected_mime_type
        text attachment_kind
        text content_sha256
        attachment_processing_status processing_status
        text processing_error
        timestamp processed_at
        timestamp staged_expires_at
        text message_id FK
        text task_id
        timestamp created_at
        timestamp updated_at
    }
    attachment_derivatives {
        text org_id FK
        text file_id PK, FK
        text kind PK
        text mime_type
        text content
        integer size_bytes
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    projects {
        text id PK
        text org_id
        text name
        text description
        text prefix
        text icon
        text color
        text lead_id FK
        boolean is_archived
        boolean is_deleted
        timestamp deleted_at
        integer task_counter
        timestamp created_at
        timestamp updated_at
    }
    project_spaces {
        text id PK
        text project_id FK
        text space_id FK
    }
    tasks {
        text id PK
        text org_id
        text project_id FK
        integer number
        text title
        text description
        task_status status
        task_priority priority
        text assignee_id FK "Primary assignee (singular). Used by board columns, dashboard 'My Tasks', nudge targeting, and status-change notifications. Every task has exactly one primary assignee or null."
        text created_by FK
        timestamp due_date
        timestamp start_date
        text estimation
        boolean is_template
        text recurrence
        text recurrence_source_id
        real sort_order
        text source_message_id FK
        text parent_task_id FK
        boolean is_deleted
        vector embedding "Task 3.8 — pgvector embedding over (title + description) for semantic search via retrieveContext({ types: ['tasks'] }). Populated by the embed-content worker (source_type: 'task') and backfilled via backfill-task-embeddings.ts."
        jsonb metadata "Task 4.11 — skill-defined custom-field payload. Keys match the `id`s in the resolved skill config's `custom_fields[]` (e.g. Sales skill stores `deal_value`, `contact_name`). Null for tasks in projects with no custom fields. Not covered by tasks.search_vector (FTS still scans title + description only)."
        timestamp created_at
        timestamp updated_at
    }
    message_attachments {
        text org_id FK
        text message_id PK, FK
        text file_id PK, FK
        integer position
        timestamp created_at
    }
    task_attachments {
        text org_id FK
        text task_id PK, FK
        text file_id PK, FK
        integer position
        timestamp created_at
    }
    labels {
        text id PK
        text org_id
        text name
        text color
        timestamp created_at
        timestamp updated_at
    }
    task_labels {
        text task_id PK, FK
        text label_id PK, FK
    }
    task_comments {
        text id PK
        text org_id FK
        text task_id FK
        text user_id FK
        text content
        boolean is_deleted
        timestamp created_at
        timestamp updated_at
    }
    task_reactions {
        text id PK
        text org_id FK
        text task_id FK
        text user_id FK
        text emoji
        timestamp created_at
    }
    task_activity {
        text id PK
        text org_id FK
        text task_id FK
        text user_id FK
        text action
        text field
        text old_value
        text new_value
        text agent_action_id "Task 3.3 — attribution to the specific agent action + employee that produced this activity row. Both nullable; when set they link the activity log entry to an agentActions row and/or an agentEmployees row so the UI can show 'done by agent X via plan Y'."
        text acting_agent_employee_id
        timestamp created_at
        timestamp updated_at
    }
    task_relationships {
        text id PK
        text source_task_id FK
        text target_task_id FK
        task_relationship_type type
        timestamp created_at
        timestamp updated_at
    }
    task_watchers {
        text id PK
        text task_id FK
        text user_id FK
        timestamp created_at
    }
    task_assignees {
        text id PK
        text task_id FK
        text user_id FK
        timestamp created_at
    }
    notifications {
        text id PK
        text org_id
        text user_id FK
        notification_type type
        text title
        text body
        text link
        boolean is_read
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    saved_views {
        text id PK
        text org_id
        text project_id FK
        text user_id FK
        text name
        jsonb config
        boolean is_shared
        timestamp created_at
        timestamp updated_at
    }
    favorites {
        text id PK
        text user_id FK
        text entity_type
        text entity_id
        real sort_order
        timestamp created_at
        timestamp updated_at
    }
    capability_provider_snapshots {
        text id PK
        text org_id FK
        text provider_kind
        text provider_instance_id
        text adapter_contract_version
        text snapshot_digest
        jsonb safe_snapshot
        timestamp captured_at
        timestamp created_at
    }
    app_runs {
        text id PK
        text org_id FK
        text contract_version
        text origin_kind
        text initiating_actor_type
        text initiating_actor_id
        text execution_actor_type
        text execution_actor_id FK
        text provider_kind FK
        text provider_instance_id FK
        text operation_name FK
        text provider_snapshot_id FK
        text origin_app_installation_id FK
        text origin_app_version_id FK
        text origin_app_binding_key FK
        text origin_runtime_binding_id FK
        text origin_resource_binding_id FK
        text origin_native_binding_id FK
        text origin_public_endpoint_id FK
        text origin_public_ingress_id FK
        text origin_app_grant_snapshot_id FK
        text origin_app_automation_definition_id FK
        text origin_app_automation_fire_id FK
        text state
        text risk_class FK
        text review_requirement FK
        text review_scope FK
        text retry_class FK
        text retention_class FK
        text idempotency_key_version
        text idempotency_fingerprint
        text input_fingerprint_key_version
        text input_fingerprint
        jsonb authorization_snapshot
        jsonb safe_preview
        jsonb safe_outcome
        text root_run_id FK
        text parent_run_id FK
        integer depth
        timestamp input_expires_at
        timestamp result_expires_at
        timestamp idempotency_expires_at
        integer attempt_limit
        text execution_release_kind
        timestamp execution_released_at
        timestamp budget_reserved_at
        integer budget_reserved_count
        integer budget_limit_at_reservation
        timestamp input_purged_at
        timestamp result_purged_at
        timestamp started_at
        timestamp terminal_at
        timestamp unknown_outcome_at
        timestamp reconciled_at
        timestamp cancelled_at
        timestamp cancel_requested_at
        timestamp created_at
        timestamp updated_at
    }
    app_run_attempts {
        text id PK
        text org_id FK
        text run_id FK
        integer attempt_number
        text retry_of_attempt_id
        text state
        text claim_owner
        text claim_token
        timestamp claimed_at
        timestamp lease_expires_at
        timestamp provider_call_started_at
        timestamp provider_call_finished_at
        text provider_idempotency_key_version
        text provider_idempotency_fingerprint
        text runtime_binding_id FK
        text resource_binding_id FK
        text runtime_session_id FK
        integer runtime_session_epoch FK
        integer runtime_epoch FK
        integer runtime_sequence
        text runtime_result_hmac
        jsonb safe_outcome
        text error_code
        timestamp created_at
        timestamp updated_at
    }
    app_run_secret_payloads {
        text id PK
        text org_id FK
        text run_id FK
        text attempt_id FK
        text payload_kind
        text envelope_version
        text algorithm
        text key_version
        text nonce_b64
        text ciphertext_b64
        text auth_tag_b64
        integer payload_bytes
        timestamp expires_at
        timestamp created_at
    }
    app_run_events {
        text id PK
        text org_id FK
        text run_id FK
        text event_version
        integer sequence
        text event_type
        text actor_type
        text actor_id
        jsonb payload
        timestamp created_at
    }
    app_run_receipts {
        text id PK
        text org_id FK
        text run_id FK
        text attempt_id FK
        text receipt_version
        text receipt_key
        text receipt_kind
        jsonb envelope
        text envelope_digest
        text signing_key_version
        text signature_hmac
        timestamp signed_at
        timestamp created_at
    }
    agent_actions {
        text id PK
        text org_id FK
        text user_id FK
        text conversation_id
        text message_id
        text agent_employee_id
        text tool_use_id
        text source
        text app_run_id FK "Nullable compatibility link only. Governed Run submission owns the one app_run_invoke row; the action is never the execution source of truth."
        text mcp_connection_id
        text plan_id
        text plan_step_id
        text channel_event_id
        text runtime_request_key
        text action
        jsonb params
        jsonb result
        approval_tier approval_tier
        approval_status approval_status
        timestamp approved_at
        text approved_by_user_id FK
        timestamp executed_at
        text error
        jsonb before_state
        jsonb after_state
        timestamp undone_at
        timestamp created_at
        timestamp updated_at
    }
    attention_items {
        text id PK
        text org_id
        text user_id FK
        text kind
        text lane
        text priority
        text state
        text dedupe_key
        text source_type
        text source_id
        text source_event_id
        text title
        text body
        text link
        jsonb metadata
        timestamp due_at
        timestamp urgent_at
        timestamp last_event_at
        integer event_count
        integer version
        timestamp seen_at
        timestamp acknowledged_at
        timestamp snoozed_until
        timestamp resolved_at
        text resolution
        timestamp created_at
        timestamp updated_at
    }
    attention_events {
        text id PK
        text org_id
        text attention_item_id FK
        text user_id FK
        text event_type
        text source_event_id
        text actor_user_id FK
        jsonb payload
        timestamp created_at
    }
    attention_deliveries {
        text id PK
        text org_id
        text attention_item_id FK
        text user_id FK
        text channel
        text status
        integer delivery_version
        integer attempt_count
        text provider_message_id
        text last_error
        timestamp next_attempt_at
        timestamp sent_at
        timestamp delivered_at
        timestamp created_at
        timestamp updated_at
    }
    web_push_subscriptions {
        text id PK
        text org_id
        text user_id FK
        text endpoint
        text endpoint_hash
        text p256dh
        text auth
        text device_name
        text user_agent
        boolean is_active
        integer failure_count
        timestamp last_used_at
        timestamp created_at
        timestamp updated_at
    }
    agent_action_approvers {
        text id PK
        text org_id
        text action_id FK
        text user_id FK
        text decision
        timestamp decided_at
        timestamp created_at
        timestamp updated_at
    }
    work_intents {
        text id PK
        text org_id FK
        text space_id FK
        text source_message_id FK
        text source_user_id FK
        text agent_employee_id FK
        work_intent_kind kind
        work_intent_status status
        text title
        text summary
        real confidence
        text proposed_action
        jsonb proposed_params
        text dedupe_key
        text converted_action_id
        text converted_task_id FK
        text converted_by FK
        timestamp converted_at
        text dismissed_by FK
        timestamp dismissed_at
        text failure_reason
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    message_observations {
        text id PK
        text org_id FK
        text message_id FK
        text space_id FK
        text user_id FK
        integer observation_version
        message_observation_status status
        text ignored_reason
        jsonb classifier_result
        jsonb downstream_jobs
        integer capture_count
        text last_error
        timestamp started_at
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }
    skills {
        text id PK
        text org_id "org_id is nullable: bundled + marketplace skills are cross-tenant."
        text name
        text description
        text slug
        text system_prompt "`system_prompt` + `param_schema` retained for back-compat with pre-Phase-4 rows. Canonical home for both is now `agent_config` (see skill-config.ts)."
        jsonb param_schema
        text source
        text version
        text icon
        jsonb agent_config
        text source_url
        boolean is_deleted
        text default_agent_employee_id
        text created_by FK
        integer usage_count
        timestamp created_at
        timestamp updated_at
    }
    task_templates {
        text id PK
        text org_id
        text name
        text description
        text icon
        text slug
        text source
        text version
        jsonb tasks
        text created_by FK
        boolean is_deleted
        integer usage_count
        timestamp created_at
        timestamp updated_at
    }
    module_installations {
        text id PK
        text org_id FK
        text module_id
        text slug
        text source
        boolean is_enabled
        timestamp disabled_at
        text agent_access "Manifests never grant agent access. An org admin chooses this installation policy independently; the service still applies principal trust/approval."
        text installed_by_user_id FK
        text installed_by_actor_type
        text installed_by_actor_id
        text updated_by_actor_type
        text updated_by_actor_id
        boolean is_deleted
        timestamp deleted_at
        text deleted_by_actor_type
        text deleted_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    module_versions {
        text id PK
        text org_id FK
        text installation_id FK
        text version
        jsonb manifest
        text manifest_digest
        boolean is_active
        timestamp activated_at
        text created_by_actor_type
        text created_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    app_installations {
        text id PK
        text org_id FK
        text app_id
        text lineage_key
        text lineage_authority_type
        text lineage_authority_id
        text source
        text state
        text active_version_id
        text active_grant_snapshot_id
        text active_grant_snapshot_kind
        integer lifecycle_epoch
        integer grant_epoch
        text installed_by_user_id FK
        text installed_by_actor_type
        text installed_by_actor_id
        text updated_by_actor_type
        text updated_by_actor_id
        timestamp disabled_at
        timestamp created_at
        timestamp updated_at
    }
    app_versions {
        text id PK
        text org_id FK
        text installation_id FK
        text version
        text protocol_version
        jsonb manifest
        text manifest_digest
        text package_digest
        jsonb package
        text requested_grant_snapshot_id "Nullable preserves pre-.23 Protocol v0 rows. New v0/v1/v2 staging writes an immutable compatibility/request projection; connected protocols require the pointer."
        jsonb provenance
        text state
        timestamp staged_at
        timestamp activated_at
        timestamp failed_at
        timestamp superseded_at
        text created_by_actor_type
        text created_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    app_module_bindings {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text module_installation_id FK
        text module_version_id FK
        text module_id
        text ownership
        timestamp created_at
    }
    app_grant_snapshots {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text app_id FK
        text app_version FK
        text manifest_digest FK
        text package_digest FK
        text snapshot_kind
        text snapshot_version
        text requested_snapshot_id FK
        text supersedes_snapshot_id FK
        jsonb resource_rights
        jsonb classification
        jsonb canonical_snapshot
        text snapshot_digest
        text reviewed_by_actor_type
        text reviewed_by_actor_id "The upgrade trigger verifies this human is an active owner/admin in org_id at insert time. It remains an immutable historical actor snapshot rather than an FK that would prevent a former reviewer from leaving the workspace."
        timestamp reviewed_at
        timestamp created_at
    }
    app_dependency_locks {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text dependency_key
        text required_app_id FK
        text required_version FK
        text dependency_installation_id FK
        text dependency_version_id FK
        text dependency_manifest_digest FK
        text dependency_package_digest FK
        integer dependency_lifecycle_epoch
        text ownership
        jsonb canonical_lock
        text lock_digest
        timestamp created_at
    }
    app_action_bindings {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text action_key
        text capability_requirement_key
        text connector_requirement_key
        text interface_identity
        text provider_kind FK
        text mcp_connection_id FK
        text provider_snapshot_id FK
        text operation_name
        text operation_schema_digest
        integer connector_authorization_version
        text risk_class
        text review_requirement
        text review_scope
        text egress_class
        text retry_class
        text retention_class
        text automation_eligibility
        boolean provider_idempotency_key_required
        jsonb canonical_binding
        text binding_digest
        timestamp created_at
    }
    app_owner_connection_admissions {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text operator_user_id FK
        text policy_key
        text review_digest
        jsonb snapshot
        text reviewed_by_user_id FK
        timestamptz reviewed_at
        text state
        timestamptz revoked_at
        timestamp created_at
        timestamp updated_at
    }
    app_runtime_registrations {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text operator_user_id FK
        text owner_connection_admission_id FK
        text owner_connection_admission_digest FK
        text contract_version
        text state
        integer runtime_epoch
        text reviewed_by_user_id
        timestamp reviewed_at
        timestamp created_at
        timestamp updated_at
    }
    app_runtime_bindings {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text runtime_registration_id FK
        text registration_contract_version FK
        text action_key
        text owner_resource_binding_id "Forward composite resource FK is installed by apply-extras/preview.57."
        text interface_identity
        text provider_kind FK
        text provider_instance_id FK
        text provider_snapshot_id FK
        text operation_name
        text risk_class
        text review_requirement
        text retry_class
        text retention_class
        text state
        text reviewed_by_user_id
        timestamp reviewed_at
        timestamp created_at
        timestamp updated_at
    }
    app_native_bindings {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text action_key
        text operation_name
        text provider_kind FK
        text provider_instance_id FK
        text provider_snapshot_id FK
        text owner_user_id FK
        text stage_manager_user_id FK
        integer stage_manager_authorization_version
        integer owner_authorization_version
        integer installation_lifecycle_epoch
        integer installation_grant_epoch
        text package_digest
        text grant_snapshot_digest
        jsonb target
        jsonb historical_create_policy
        text proposal_digest
        text consent_digest
        text reviewed_contract_digest
        text risk_class
        text review_requirement
        text review_scope
        text retry_class
        text retention_class
        text state
        timestamp reviewed_at
        timestamp created_at
        timestamp updated_at
    }
    app_resource_bindings {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text runtime_registration_id FK
        text registration_contract_version FK
        text provider_kind FK
        text provider_instance_id FK
        text provider_snapshot_id FK
        text resource_key
        text target_key
        text resource_family
        text operation_name
        text interface_identity
        jsonb reviewed_descriptor
        text descriptor_digest
        jsonb attachment_policy
        text attachment_consent_digest
        text owner_user_id FK
        text owner_scope
        text risk_class
        text review_requirement
        text review_scope
        text retry_class
        text retention_class
        integer max_records_per_page
        integer max_page_bytes
        integer max_retained_records
        integer max_retained_bytes
        integer min_interval_seconds
        timestamp consent_expires_at
        text consent_lifetime
        text state
        text reviewed_by_user_id FK
        timestamp reviewed_at
        timestamp created_at
        timestamp updated_at
    }
    app_runtime_sessions {
        text id PK
        text org_id FK
        text runtime_registration_id FK
        text runtime_binding_id FK
        text resource_binding_id FK
        text operator_user_id FK
        text token_hash
        text audience
        integer session_epoch
        integer runtime_epoch
        integer lifecycle_epoch
        integer grant_epoch
        integer next_sequence
        timestamp expires_at
        timestamptz renewable_until
        text renewal_lifetime
        text renewal_web_session_id
        text renewal_operator_role
        timestamp revoked_at
        timestamp created_at
        timestamp updated_at
    }
    app_sync_checkpoints {
        text id PK
        text org_id FK
        text resource_binding_id FK
        integer generation
        text state
        integer cursor_sequence
        text cursor_hmac_key_version
        text cursor_hmac
        text cursor_state
        text cursor_envelope_version
        text cursor_algorithm
        text cursor_key_version
        text cursor_nonce_b64
        text cursor_ciphertext_b64
        text cursor_auth_tag_b64
        integer cursor_bytes
        integer retained_record_count
        integer retained_bytes
        text last_applied_run_id
        text last_applied_page_digest
        timestamp last_applied_at
        timestamp last_checked_at
        timestamp fresh_until
        timestamp created_at
        timestamp updated_at
    }
    app_sync_intents {
        text id PK
        text org_id FK
        text run_id FK
        text resource_binding_id FK
        text checkpoint_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text provider_snapshot_id FK
        text owner_user_id FK
        text descriptor_digest FK
        integer generation
        integer expected_cursor_sequence
        text expected_cursor_hmac_key_version
        text expected_cursor_hmac
        timestamp created_at
    }
    app_resource_projections {
        text id PK
        text org_id FK
        text resource_binding_id FK
        text checkpoint_id FK
        integer generation
        text resource_id_hmac_key_version
        text resource_id_hmac
        text provider_id_envelope_version
        text provider_id_algorithm
        text provider_id_key_version
        text provider_id_nonce_b64
        text provider_id_ciphertext_b64
        text provider_id_auth_tag_b64
        integer provider_id_bytes
        text body_envelope_version
        text body_algorithm
        text body_key_version
        text body_nonce_b64
        text body_ciphertext_b64
        text body_auth_tag_b64
        integer body_bytes
        text state
        integer applied_sequence
        timestamp first_seen_at
        timestamp last_seen_at
        timestamp source_updated_at
        timestamp fresh_until
        timestamp tombstoned_at
        timestamp created_at
        timestamp updated_at
    }
    app_attachment_stages {
        text id PK
        text org_id FK
        text resource_binding_id FK
        text checkpoint_id FK
        integer generation
        text run_id FK
        text attempt_id FK
        text claim_token
        integer reservation_sequence
        text fingerprint_key_version
        text parent_locator_hmac
        text parent_revision_hmac
        text attachment_key_hmac
        text content_hmac
        integer declared_size_bytes
        jsonb metadata_envelope
        text binary_key_version
        text binary_nonce_b64
        text binary_auth_tag_b64
        text object_id
        text state
        timestamp stage_expires_at
        timestamp linked_expires_at
        text projection_id FK
        text parent_body_hmac
        timestamp accepted_at
        timestamp retired_at
        timestamp purged_at
        timestamp created_at
        timestamp updated_at
    }
    app_automation_definitions {
        text id PK
        text org_id FK
        text app_installation_id FK
        text app_version_id FK
        text app_manifest_digest
        text app_package_digest
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text grant_snapshot_digest
        text action_binding_id FK
        text action_key FK
        text interface_identity
        text automation_request_key
        text automation_request_digest
        integer installation_lifecycle_epoch
        integer installation_grant_epoch
        text provider_kind
        text mcp_connection_id
        text provider_snapshot_id
        text provider_snapshot_digest
        text operation_name
        text operation_schema_digest
        text binding_digest
        integer connector_authorization_version
        jsonb placement_resource_ref
        text placement_resource_revision
        text placement_content_digest
        jsonb selected_resource_ref
        text selected_resource_revision
        text selected_content_digest
        text selected_relation_input_key
        text selected_relation_key
        integer selected_relation_revision
        text schedule_kind
        text local_time
        text timezone
        text misfire_policy
        integer catch_up_window_minutes
        integer max_actions_per_fire
        integer max_org_runs_per_utc_day
        integer max_pending_org_fires
        timestamp valid_from
        timestamp valid_until
        text policy_version
        text policy_digest
        jsonb authorization_vector
        text authorization_digest
        jsonb canonical_definition
        text definition_digest
        text state
        integer definition_epoch
        text created_by_user_id
        text approved_by_user_id
        integer approver_authorization_version
        timestamp approved_at
        timestamp state_changed_at
        timestamp revoked_at
        timestamp expired_at
        timestamp created_at
        timestamp updated_at
    }
    app_automation_fires {
        text id PK
        text org_id FK
        text definition_id FK
        integer definition_epoch
        text logical_local_date
        text local_time
        text timezone
        timestamp resolved_at_utc
        text fire_identity
        text state
        integer attempt_count
        text claim_owner
        text claim_token
        timestamp claimed_at
        timestamp lease_expires_at
        text app_run_id
        text terminal_reason
        timestamp terminal_at
        timestamp created_at
        timestamp updated_at
    }
    app_developer_pairings {
        text id PK
        text org_id FK
        text code_hash
        text created_by_user_id FK
        timestamp expires_at
        timestamp consumed_at
        timestamp revoked_at
        text session_token_hash
        timestamp session_expires_at
        timestamp install_used_at
        timestamp created_at
        timestamp updated_at
    }
    module_records {
        text id PK
        text org_id FK
        text installation_id FK
        text collection_key
        text validated_version_id FK
        jsonb data
        integer revision
        text create_idempotency_key "Durable create dedupe. Update/archive safety is provided by revision CAS; storing only a mutable 'last request' key would make old retries unsafe."
        text search_title "Only manifest-declared fields are projected into these columns by the module service. Source JSON remains untouched and is never blanket-indexed."
        text search_subtitle
        text search_text
        tsvector search_vector
        text created_by_actor_type
        text created_by_actor_id
        text updated_by_actor_type
        text updated_by_actor_id
        boolean is_deleted
        timestamp deleted_at
        text deleted_by_actor_type
        text deleted_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    module_record_relations {
        text id PK
        text org_id FK
        text installation_id FK
        text field_key
        text source_record_id FK
        text target_record_id FK
        integer position
        text created_by_actor_type
        text created_by_actor_id
        text updated_by_actor_type
        text updated_by_actor_id
        boolean is_deleted
        timestamp deleted_at
        text deleted_by_actor_type
        text deleted_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    module_record_merges {
        text id PK
        text org_id FK
        text installation_id FK
        text source_record_id FK
        text target_record_id FK
        integer source_revision
        integer target_revision
        jsonb source_data
        jsonb target_data
        jsonb link_snapshot
        jsonb choices
        text created_by
        timestamp created_at
    }
    module_saved_views {
        text id PK
        text org_id FK
        text installation_id FK
        text collection_key
        text owner_user_id FK
        text name
        text view_type
        jsonb config
        boolean is_deleted
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    module_mutation_receipts {
        text id PK
        text org_id FK
        text installation_id FK
        text agent_action_id FK
        text actor_type
        text actor_id
        text operation
        text idempotency_key
        text input_digest
        text record_id FK
        integer result_revision
        text result_manifest_digest
        boolean result_archived
        text changed_fields
        timestamp created_at
    }
    resource_relation_sets {
        text id PK
        text org_id FK
        text source_provider_kind
        text source_provider_instance_id
        text source_resource_type
        text source_resource_id
        text relation_key
        integer revision
        text updated_by_actor_type
        text updated_by_actor_id
        timestamp created_at
        timestamp updated_at
    }
    resource_relation_edges {
        text id PK
        text org_id FK
        text relation_set_id FK
        text target_provider_kind
        text target_provider_instance_id
        text target_resource_type
        text target_resource_id
        integer position
        text created_by_actor_type
        text created_by_actor_id
        boolean is_deleted
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    resource_relation_receipts {
        text id PK
        text org_id FK
        text relation_set_id FK
        text actor_type
        text actor_id
        text operation
        text idempotency_key
        text input_digest
        integer result_revision
        jsonb result_refs
        timestamp created_at
    }
    agent_webhooks {
        text id PK
        text org_id FK
        text agent_employee_id FK
        text slug
        text secret_hash
        text hmac_key_encrypted "Per-webhook AES-encrypted HMAC key (fix #7). New webhooks issue this alongside the legacy scrypt secret so callers can sign with HMAC-SHA256 (`x-deft-webhook-signature: sha256=<hex>`) instead of shipping the raw secret. Pre-existing rows have NULL until they're rotated."
        text label
        boolean enabled
        timestamp last_fired_at
        integer fire_count
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    tools {
        text id PK
        text name
        text description
        text category
        jsonb params_schema
        approval_tier approval_tier
        boolean is_active
        timestamp created_at
        timestamp updated_at
    }
    triggers {
        text id PK
        text org_id
        text name
        text event_type
        jsonb condition
        jsonb actions
        boolean is_active
        text schedule
        text agent_employee_id
        timestamp last_fired_at
        integer fire_count
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    connected_accounts {
        text id PK
        text org_id
        text user_id FK
        text provider
        text provider_account_id
        text access_token_encrypted
        text refresh_token_encrypted
        timestamp token_expires_at
        text scopes
        jsonb metadata
        timestamp last_sync_at
        text sync_error
        timestamp created_at
        timestamp updated_at
    }
    events {
        text id PK
        text org_id
        event_source source
        text event_type
        text external_id
        text title
        text body
        text url
        text actor
        timestamp timestamp
        jsonb metadata
        text user_id FK
        text connected_account_id FK
        timestamp created_at
        timestamp updated_at
    }
    ics_subscriptions {
        text id PK
        text org_id
        text user_id FK
        text ics_url
        text label
        integer sync_interval_min
        boolean is_active
        timestamp last_synced_at
        text last_error
        integer last_event_count
        timestamp created_at
        timestamp updated_at
    }
    reminders {
        text id PK
        text org_id
        text user_id FK
        text message
        timestamp remind_at
        text source_message_id FK
        boolean is_sent
        timestamp created_at
        timestamp updated_at
    }
    pinned_messages {
        text id PK
        text message_id FK
        text space_id FK
        text pinned_by FK
        timestamp pinned_at
    }
    scheduled_messages {
        text id PK
        text org_id
        text user_id FK
        text space_id FK
        text content
        timestamp scheduled_for
        text status
        timestamp sent_at
        timestamp created_at
        timestamp updated_at
    }
    canvases {
        text id PK
        text org_id
        text space_id FK
        text title
        jsonb content
        text last_edited_by FK
        timestamp last_edited_at
        timestamp created_at
        timestamp updated_at
    }
    message_bookmarks {
        text id PK
        text org_id
        text user_id FK
        text message_id FK
        text space_id FK
        timestamp created_at
        timestamp updated_at
    }
    user_groups {
        text id PK
        text org_id
        text name
        text handle
        text description
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    user_group_members {
        text id PK
        text group_id FK
        text user_id FK
    }
    teams {
        text id PK
        text org_id
        text name
        text handle
        text description
        text type
        team_visibility visibility
        text avatar_url
        text color
        text lead_user_id FK
        text default_space_id FK
        boolean is_archived
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    team_members {
        text id PK
        text org_id
        text team_id FK
        text user_id FK
        team_role role
        timestamp joined_at
        timestamp created_at
        timestamp updated_at
    }
    team_resources {
        text id PK
        text org_id
        text team_id FK
        team_resource_type resource_type
        text resource_id
        text label
        text created_by FK
        timestamp created_at
    }
    team_dashboard_snapshots {
        text id PK
        text org_id
        text team_id FK
        text snapshot_type
        jsonb payload_json
        timestamp generated_at
        timestamp created_at
    }
    custom_emoji {
        text id PK
        text org_id
        text name
        text image_url
        text uploaded_by FK
        timestamp created_at
        timestamp updated_at
    }
    workflow_rules {
        text id PK
        text org_id
        text name
        text trigger_type
        jsonb trigger_config
        text action_type
        jsonb action_config
        text created_by FK
        boolean is_active
        timestamp created_at
        timestamp updated_at
    }
    workflow_runs {
        text id PK
        text rule_id FK
        text triggered_by_message_id FK
        text triggered_by_user_id FK
        jsonb result
        text status
        timestamp executed_at
    }
    automation_runs {
        text id PK
        text org_id
        text kind
        text subject_id
        text user_id FK
        text agent_employee_id
        text idempotency_key
        timestamp scheduled_for
        text status
        text generator
        jsonb context
        jsonb output
        text result_entity_id
        text error
        timestamp started_at
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }
    onboarding_state {
        text id PK
        text user_id FK
        boolean org_created
        boolean profile_set
        boolean first_space_created
        boolean first_message_sent
        boolean first_invite_sent
        boolean first_task_created
        boolean agent_tried
        boolean completed
        timestamp created_at
        timestamp updated_at
    }
    standups {
        text id PK
        text org_id
        timestamp date
        text generated_by
        text summary
        jsonb raw_data
        timestamp created_at
        timestamp updated_at
    }
    agent_memory {
        text id PK
        text org_id
        text user_id FK
        text conversation_id
        text scope
        text key
        text value
        timestamp created_at
        timestamp updated_at
    }
    cross_references {
        text id PK
        text org_id
        text source_type
        text source_id
        text target_type
        text target_id
        text context
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    duplicate_flags {
        text id PK
        text org_id FK
        text task_a_id FK
        text task_b_id FK
        numeric similarity
        timestamp created_at
    }
    audit_log {
        text id PK
        text org_id
        text actor_type
        text actor_id
        text action
        text entity_type
        text entity_id
        jsonb before_state
        jsonb after_state
        jsonb metadata
        timestamp created_at
    }
    agent_nudges {
        text id PK
        text org_id
        text user_id FK
        text task_id FK
        text nudge_type
        text message
        boolean is_dismissed
        timestamp created_at
        timestamp updated_at
    }
    job_queue {
        text id PK
        text org_id "System-wide jobs (for example cron scans) intentionally have no org. Product jobs should set this so queue health and dedupe remain tenant-aware."
        text queue
        text name
        jsonb data
        text status
        integer attempts
        integer max_attempts
        timestamp run_at
        timestamp started_at
        timestamp completed_at
        text error
        text cron_key
        text dedupe_key "Idempotency is retention-bound: terminal rows keep their key until the queue retention sweep removes them."
        text locked_by
        text lock_token
        timestamp lock_expires_at
        timestamp created_at
    }
    meeting_briefs {
        text id PK
        text org_id
        text user_id FK
        text event_id FK
        text brief_text
        timestamp created_at
        timestamp updated_at
    }
    people_interactions {
        text id PK
        text org_id
        text user_a_id FK
        text user_b_id FK
        integer interaction_count
        real recency_weighted_score
        integer dm_count
        integer shared_space_count
        integer mention_count
        integer thread_co_participation
        timestamp last_interaction_at
        timestamp updated_at
    }
    people_expertise {
        text id PK
        text org_id
        text user_id FK
        text topic
        integer message_count
        integer question_answered_count
        integer mentioned_for_help_count
        integer tasks_completed_count
        real expertise_score
        timestamp first_seen_at
        timestamp updated_at
    }
    people_influence {
        text id PK
        text org_id
        text user_id FK
        text influence_type
        text context
        real score
        integer evidence_count
        jsonb evidence_samples
        timestamp updated_at
    }
    people_patterns {
        text id PK
        text org_id
        text user_id FK
        text pattern_type
        jsonb pattern_data
        jsonb baseline_data
        real confidence
        timestamp updated_at
    }
    people_relationships {
        text id PK
        text org_id
        text user_a_id FK
        text user_b_id FK
        text relationship_type
        real strength
        text direction
        jsonb evidence
        timestamp first_detected_at
        timestamp updated_at
    }
    team_health_snapshots {
        text id PK
        text org_id
        timestamp snapshot_date
        jsonb team_data
        text generated_by
        timestamp created_at
    }
    oneone_preps {
        text id PK
        text org_id
        text manager_id FK
        text report_id FK
        timestamp meeting_date
        jsonb prep_content
        text status
        timestamp created_at
        timestamp updated_at
    }
    burnout_alerts {
        text id PK
        text org_id
        text user_id FK
        text alerted_to FK
        jsonb signals
        real confidence
        text status
        timestamp created_at
        timestamp acknowledged_at
    }
    note_folders {
        text id PK
        text org_id
        text user_id FK
        text name
        text icon
        text parent_folder_id
        integer sort_order
        boolean is_deleted
        timestamp created_at
        timestamp updated_at
    }
    notes {
        text id PK
        text org_id
        text user_id FK
        text folder_id FK
        text title
        text content
        text icon
        boolean is_pinned
        boolean is_template
        boolean is_deleted
        integer version
        text visibility
        text visibility_space_id FK
        timestamp created_at
        timestamp updated_at
    }
    note_versions {
        text id PK
        text note_id FK
        integer version
        text title
        text content
        text edited_by
        timestamp created_at
    }
    note_shares {
        text id PK
        text note_id FK
        text shared_with_user_id FK
        text permission
        timestamp created_at
    }
    tags {
        text id PK
        text org_id
        text name
        text color
        timestamp created_at
        timestamp updated_at
    }
    entity_tags {
        text id PK
        text org_id
        text tag_id FK
        entity_tag_type entity_type
        text entity_id
        timestamp created_at
        timestamp updated_at
    }
    clips {
        text id PK
        text org_id
        text space_id FK
        text message_id FK
        clip_context_type context_type
        text context_id
        clip_mode mode
        text created_by FK
        integer duration_s
        text file_key
        integer file_size
        text mime_type
        clip_status status
        text transcript
        jsonb segments
        jsonb summary
        jsonb participants
        text whisper_model
        text error
        boolean is_deleted
        timestamp created_at
        timestamp updated_at
    }
    manager_settings {
        text id PK
        text user_id FK
        text org_id
        text team_pulse_frequency
        boolean oneone_prep_enabled
        boolean burnout_alerts_enabled
        integer overload_threshold
        integer blocked_threshold_hours
        boolean weekly_digest_enabled
        timestamp created_at
        timestamp updated_at
    }
    thread_reads {
        text id PK
        text user_id FK
        text parent_message_id FK
        timestamp last_read_at
        timestamp created_at
        timestamp updated_at
    }
    message_versions {
        text id PK
        text message_id FK
        text content
        timestamp edited_at
    }
    wiki_pages {
        text id PK
        text org_id
        wiki_page_scope scope
        text space_id FK
        text origin_space_id FK
        text origin_message_id FK
        text origin_user_id FK
        text created_via
        text user_id FK
        text agent_employee_id
        wiki_page_type type
        text title
        text slug
        text summary
        text content
        jsonb metadata
        real confidence
        integer version
        text previous_content
        boolean is_deleted
        vector embedding
        text tags
        text referenced_user_ids
        timestamp created_at
        timestamp updated_at
    }
    wiki_links {
        text id PK
        text org_id
        text source_page_id FK
        text target_page_id FK
        text context
        timestamp created_at
    }
    wiki_citations {
        text id PK
        text org_id FK
        text page_id FK
        text source_type
        text source_id
        text source_space_id FK
        text source_user_id FK
        text excerpt
        timestamp created_at
    }
    wiki_memory_syncs {
        text id PK
        text org_id
        text agent_employee_id FK
        text idempotency_key
        text content_digest
        text page_id FK
        integer page_version
        text runtime_session_id
        jsonb provenance
        timestamp created_at
        timestamp updated_at
    }
    wiki_ops_log {
        text id PK
        text org_id
        text operation
        text page_id FK
        jsonb details
        text performed_by
        timestamp created_at
    }
    wiki_page_versions {
        text id PK
        text page_id FK
        integer version
        text title
        text content
        text summary
        text edited_by
        timestamp created_at
    }
    mcp_connections {
        text id PK
        text org_id
        text name
        text slug
        text server_url
        mcp_transport transport
        text stdio_command
        jsonb stdio_args
        text auth_type
        jsonb auth_config_encrypted
        boolean is_active
        timestamp last_connected_at
        text connection_error
        jsonb tools_cache
        timestamp tools_cached_at
        approval_tier default_trust_tier
        text enabled_tools
        integer app_run_authorization_version
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    mcp_tool_overrides {
        text id PK
        text org_id FK
        text mcp_connection_id FK
        text tool_name
        approval_tier trust_tier_override
        boolean is_disabled
        integer app_run_authorization_version
        timestamp created_at
        timestamp updated_at
    }
    agent_employees {
        text id PK
        text org_id
        text user_id FK
        text name
        text slug
        agent_employee_role role
        text avatar_url
        text system_prompt
        text expertise_description
        text starter_prompts
        text mcp_connection_ids "native_tools[] removed in Task 4.12 (migration 0038) — per-employee tool selection moved to the skills primitive (migrations 0035-0037)."
        text disabled_tools
        text space_ids
        text project_ids
        trust_level trust_level
        integer max_daily_actions
        integer daily_action_count
        timestamp daily_action_reset_at
        boolean heartbeat_enabled
        integer heartbeat_interval_min
        jsonb heartbeat_config
        jsonb heartbeat_overrides "Task 8.2 — per-employee overlay for the heartbeat prompt builder. Shape (all fields optional): { checklist?: string[];          // extra checklist items cadence_minutes?: number;      // per-employee cadence override (Task 8.3) }"
        integer daily_budget_cents "Task 8.5 — daily cost guardrails. `daily_cost_cents` is reset at UTC midnight alongside `daily_action_count`; `daily_budget_cents` is the soft cap enforced by the heartbeat + trigger dispatchers. Default 10000 (=$100/day) mirrors the PR plan, tunable per-employee via the PATCH endpoint."
        integer daily_cost_cents
        boolean unhealthy "Task 8.5 / 8.6 — circuit breaker. Set by the heartbeat handler after 3 consecutive errors OR by Task 8.6's loop detector; cleared via PATCH { mark_healthy: true }. Blocks every autonomous dispatcher from firing until cleared."
        text unhealthy_reason
        timestamp last_heartbeat_at
        boolean is_active
        boolean is_deleted "is_active=false is 'pause' (resumable). is_deleted=true is a real soft-delete — the row is hidden from every list endpoint and cannot be restored from the UI. Separate semantics added in migration 0058."
        timestamp deleted_at
        boolean is_byoa
        text byoa_model_info
        text mcp_token_hash
        text runtime_kind "BYOA runtime/control-plane metadata. Deft registers an already-running agent as a workplace employee; it does not own the runtime's identity."
        text job_title
        text wake_mode
        text certification_status
        timestamp last_verified_at
        timestamp last_mcp_call_at
        timestamp last_work_outcome_at
        text connection_notes
        text trigger_subscriptions "trigger_subscriptions is the routing key for the trigger system (e.g. member.joined, cron:standup) — kept as part of Phase 9."
        integer app_run_authorization_version
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    agent_certification_challenges {
        text id PK
        text org_id
        text employee_id FK
        text nonce
        text required_tools
        text status
        text failure_reason
        timestamp started_at
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }
    agent_mcp_call_audit {
        text id PK
        text org_id
        text employee_id FK
        text tool_name
        boolean success
        text error
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    mcp_tokens {
        text id PK
        text org_id
        text user_id FK
        text agent_employee_id FK
        text principal_kind
        text name
        text token_hash
        text token_prefix
        text scopes
        timestamp last_used_at
        timestamp revoked_at
        integer app_run_authorization_version
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_connections {
        text id PK
        text org_id
        text agent_employee_id FK
        text runtime_kind
        text status
        text protocol_version
        timestamp last_seen_at
        text last_event_id
        text last_error
        timestamp paused_at
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_tokens {
        text id PK
        text org_id
        text agent_employee_id FK
        text name
        text token_hash
        text token_prefix
        jsonb scopes
        boolean is_active
        timestamp last_used_at
        timestamp revoked_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_events {
        text id PK
        text org_id
        text agent_employee_id FK
        text kind
        text source_kind
        text source_id
        text space_id FK
        text thread_id FK
        text actor_user_id FK
        jsonb payload
        text idempotency_key
        text status
        integer delivery_count
        text claim_owner
        text claim_token
        timestamp claimed_at
        timestamp lease_expires_at
        timestamp delivered_at
        timestamp acked_at
        timestamp completed_at
        timestamp failed_at
        text work_outcome
        text outcome_detail
        timestamp outcome_at
        text runtime_session_key
        text error
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_cursors {
        text id PK
        text org_id
        text agent_employee_id FK
        text connection_id FK
        text last_delivered_event_id
        text last_acked_event_id
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_sessions {
        text id PK
        text org_id
        text agent_employee_id FK
        text deft_scope
        text deft_scope_id
        text runtime_session_key
        text busy_mode
        timestamp last_active_at
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    agent_channel_delivery_attempts {
        text id PK
        text org_id
        text agent_employee_id FK
        text event_id FK
        text direction
        text idempotency_key
        text status
        jsonb request_json
        jsonb response_json
        text error
        timestamp created_at
        timestamp updated_at
    }
    oauth_clients {
        text id PK
        text client_id
        text client_secret_hash
        text client_name
        text client_uri
        text logo_uri
        text redirect_uris
        text grant_types
        text response_types
        text token_endpoint_auth_method
        jsonb metadata
        timestamp created_at
        timestamp updated_at
    }
    oauth_authorization_codes {
        text id PK
        text code_hash
        text org_id
        text user_id FK
        text client_id
        text redirect_uri
        text code_challenge
        text code_challenge_method
        text resource
        text scopes
        timestamp expires_at
        timestamp used_at
        timestamp created_at
        timestamp updated_at
    }
    oauth_grants {
        text id PK
        text org_id
        text user_id FK
        text client_id
        text app_name
        text connector_profile
        text scopes
        timestamp revoked_at
        timestamp created_at
        timestamp updated_at
    }
    oauth_access_tokens {
        text id PK
        text token_hash
        text grant_id FK
        text org_id
        text user_id FK
        text client_id
        text resource
        text scopes
        timestamp expires_at
        timestamp last_used_at
        timestamp revoked_at
        integer app_run_authorization_version
        timestamp created_at
        timestamp updated_at
    }
    oauth_refresh_tokens {
        text id PK
        text token_hash
        text grant_id FK
        text rotated_from
        timestamp expires_at
        timestamp revoked_at
        timestamp created_at
        timestamp updated_at
    }
    oauth_audit_events {
        text id PK
        text org_id
        text user_id FK
        text client_id
        text event
        jsonb metadata
        timestamp created_at
    }
    agent_employee_skills {
        text agent_employee_id PK, FK
        text skill_id PK, FK
        timestamp installed_at
        text installed_version
    }
    agent_plans {
        text id PK
        text org_id
        text user_id FK
        text agent_employee_id
        text conversation_id
        text title
        text description
        jsonb steps
        plan_status status
        integer current_step
        jsonb context
        text error
        boolean fail_fast "Task 3.9 — fail-fast mode. When true, the executor marks every later step 'skipped_due_to_failure' and stops as soon as any step fails, instead of asking the agent for an alternative path. Default false preserves existing recovery-and-continue behavior."
        boolean rollback_on_fail "Task 3.9 — rollback-on-fail mode. Only meaningful when fail_fast=true. When set, successful write-action steps taken earlier in the plan are reversed on failure (create_task → soft-delete, post_message → mark deleted). Steps without a safe reversal (update_task_*) log a warning and are left as-is."
        timestamp created_at
        timestamp updated_at
    }
    api_keys {
        text id PK
        text org_id
        text agent_employee_id
        text name
        text key_hash
        text key_prefix
        text permissions
        integer rate_limit_per_minute
        integer rate_limit_per_day
        timestamp last_used_at
        integer request_count
        boolean is_active
        timestamp expires_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    agent_employee_templates {
        text id PK
        text org_id FK "Block 3.1 — nullable. NULL = first-party/community seed; non-NULL = org-scoped 'Save as template' clone. Uniqueness is (org_id, slug) declared via the SQL migration's COALESCE-keyed partial index."
        text slug
        text name
        text version
        agent_employee_role role
        text description
        text soul_md
        text agents_md
        text user_md_template
        text tools_md
        text default_tools
        text default_capability_packs "Phase 9 — pack slugs matching `CAPABILITY_PACKS` in capability-packs.ts. Nullable for backward compatibility with rows seeded before migration 0016."
        trust_level default_trust_level
        text default_trigger_subscriptions
        text model_recommendation
        text fallback_models
        text source
        text source_attribution
        integer download_count
        boolean is_public
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    agent_cooperative_log {
        text id PK
        text org_id
        text employee_id FK
        text kind
        text summary "Free-form narrative the agent sends. Never truncated — the point of the log is to receive the agent's voice verbatim."
        jsonb metadata "Optional structured metadata (decision alternatives, outcome code, attempted-action name, etc.). Shape is intentionally unscoped."
        text session_turn_id "Optional pointer to the turn this record belongs to, when the agent can provide one. Allows a future rollup to thread records."
        timestamp created_at
        timestamp updated_at
    }
    agent_session_turns {
        text id PK
        text org_id
        text employee_id FK
        text trigger_kind
        text triggering_message_id
        text space_id
        jsonb input_messages_json
        text raw_reply_text
        jsonb tool_calls_json
        integer latency_ms
        text model_name
        integer tokens_in
        integer tokens_out
        text result
        text error
        timestamp created_at
        timestamp updated_at
    }
    agent_heartbeat_turns {
        text id PK
        text org_id
        text agent_employee_id FK
        timestamp fired_at
        integer cadence_minutes
        text prompt_sha
        integer action_count
        integer tokens_in
        integer tokens_out
        integer cost_cents
        text outcome "Outcome vocabulary: - 'dispatched'          — succeeded, agent ran - 'no_op'               — agent returned HEARTBEAT_OK - 'skipped_budget'      — daily action / cost cap hit - 'skipped_idempotent'  — same prompt_sha as last no_op - 'skipped_unhealthy'   — circuit breaker tripped - 'skipped_disconnected'— Gateway not connected - 'error'               — dispatcher threw"
        text outcome_reason
        text summary
        jsonb raw_response
    }
    action_receipts {
        text id PK
        text org_id
        text action_id FK
        text employee_id FK
        text proposer
        text proposer_id
        text approver_id FK
        text decision
        text decision_reason
        text action_name
        jsonb action_params_json
        jsonb result_json
        text signature_hmac
        timestamp signed_at
        timestamp created_at
        timestamp updated_at
    }
    space_memory {
        text id PK
        text org_id
        text space_id FK
        text key
        jsonb value
        text updated_by_employee_id FK
        timestamp created_at
        timestamp updated_at
    }
    integrations {
        text id PK
        text org_id
        text provider
        text account_label
        text access_token_encrypted
        text refresh_token_encrypted
        timestamp access_token_expires_at
        text scopes
        text external_workspace_id
        text external_workspace_name
        text external_default_project_id
        text status
        text connected_by FK
        timestamp last_used_at
        timestamp created_at
        timestamp updated_at
    }
    message_classifications {
        text id PK
        text org_id
        text message_id FK
        text intent
        real confidence
        boolean agent_mentioned
        boolean blocked
        text task_references
        jsonb entities
        text memorable_facts
        text decision
        timestamp created_at
    }
    web_sessions {
        text id PK
        text user_id FK
        text org_id FK
        text refresh_token_hash
        timestamptz expires_at
        timestamptz revoked_at
        timestamptz created_at
    }
    app_experience_consent_grants {
        text id PK
        text org_id FK
        text owner_user_id FK
        text app_installation_id FK
        text app_version_id FK
        text experience_key
        text scope_digest
        jsonb snapshot
        integer epoch
        timestamptz created_at
        timestamptz revoked_at
    }
    app_runtime_agent_policies {
        text org_id PK, FK
        text owner_user_id PK, FK
        text runtime_binding_id PK, FK
        text mode
        integer revision
        timestamptz created_at
        timestamptz updated_at
    }
    app_experience_sessions {
        text id PK
        text org_id FK
        text user_id FK
        text web_session_id FK
        text consent_grant_id FK
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        text grant_snapshot_kind FK
        text experience_key
        text resource_target_key
        text artifact_digest
        integer lifecycle_epoch
        integer grant_epoch
        timestamptz created_at
        timestamptz expires_at
        timestamptz revoked_at
    }
    app_experience_resource_exposures {
        text id PK
        text org_id FK
        text experience_session_id FK
        text owner_user_id FK
        text web_session_id FK
        text review_digest
        jsonb snapshot
        text payload_policy_version
        integer exposure_epoch
        timestamptz created_at
        timestamptz expires_at
        timestamptz revoked_at
    }
    app_action_batches {
        text id PK
        text org_id FK
        text owner_user_id FK
        text runtime_binding_id FK
        text source
        text employee_id
        text token_id
        text token_kind
        integer token_version
        text idempotency_digest
        text content_digest
        text title
        text consent_grant_id FK
        integer consent_epoch
        integer policy_revision
        text state
        timestamptz created_at
        timestamptz approved_at
        timestamptz cancelled_at
    }
    app_action_batch_items {
        text org_id PK, FK
        text batch_id PK, FK
        text item_key PK
        text label
        integer ordinal
        text run_id FK
        text input_digest
    }
    app_private_files {
        text id PK
        text org_id FK
        text owner_user_id FK
        text installation_id FK
        text resource_target_key
        text state_key
        text record_id
        text artifact_digest
        text declaration_digest
        text upload_operation_id
        integer upload_revision
        text content_digest
        integer byte_length
        jsonb metadata_envelope
        bytea ciphertext
        text key_version
        text nonce_b64
        text auth_tag_b64
        timestamptz created_at
        timestamptz updated_at
        timestamptz expires_at
        timestamptz detached_at
        timestamptz purged_at
    }
    app_action_payloads {
        text id PK
        text org_id FK
        text owner_user_id FK
        text installation_id FK
        text app_version_id
        text grant_snapshot_id
        text resource_target_key
        text state_key
        text draft_id
        integer draft_revision
        text artifact_digest
        text declaration_digest
        text runtime_binding_id
        text action_key
        text action_contract_digest
        text input_digest
        text snapshot_digest
        text bytes_digest
        integer byte_length
        jsonb metadata_envelope
        bytea ciphertext
        text key_version
        text nonce_b64
        text auth_tag_b64
        timestamptz created_at
        timestamptz updated_at
        timestamptz expires_at
        timestamptz purged_at
    }
    app_action_payload_files {
        text org_id PK, FK
        text owner_user_id FK
        text snapshot_id PK, FK
        text file_id PK, FK
        integer position
    }
    app_run_human_authorizations {
        text org_id PK, FK
        text run_id PK, FK
        text owner_user_id FK
        text consent_grant_id FK
        integer consent_epoch
        text exposure_id FK
        integer exposure_epoch
        text snapshot_id FK
        text snapshot_digest FK
        timestamptz created_at
    }
    app_experience_resource_exposure_resources {
        text org_id PK, FK
        text exposure_id PK, FK
        text resource_key PK
        text resource_binding_id FK
        text runtime_registration_id FK
        integer runtime_epoch
        text descriptor_digest
        text resource_type
        jsonb allowed_operations
        jsonb allowed_fields
    }
    app_experience_resource_exposure_audit {
        text id PK
        text org_id
        text exposure_id
        text experience_session_id
        text owner_user_id
        text review_digest
        text event
        jsonb safe_snapshot
        timestamptz created_at
    }
    revoked_tokens {
        text id PK
        text token_hash
        text user_id
        text org_id
        timestamp revoked_at
    }
    app_public_endpoints {
        text id PK
        text org_id FK
        text slug_digest
        text app_installation_id FK
        text app_version_id FK
        text grant_snapshot_id FK
        integer installation_lifecycle_epoch
        integer installation_grant_epoch
        text module_installation_id FK
        text collection_key
        text public_action_key
        text runtime_binding_id FK
        text native_binding_id FK
        jsonb native_input_mapping
        text approver_user_id FK
        jsonb input_mapping
        text mapping_digest
        jsonb availability_policy
        jsonb budget_policy
        jsonb cancellation_policy
        text cancel_native_binding_id FK
        jsonb authentication_policy
        text hmac_key_id
        text state
        integer endpoint_epoch
        text review_digest
        text reviewed_by_user_id
        timestamp reviewed_at
        text public_label
        integer max_body_bytes
        timestamp created_at
        timestamp updated_at
    }
    app_public_ingress {
        text id PK
        text org_id FK
        text endpoint_id FK
        integer endpoint_epoch
        text request_key_digest
        text input_digest
        text state
        text follow_up_state
        text follow_up_code
        timestamp handled_at
        timestamp created_at
    }
    app_public_hmac_keys {
        text id PK
        text org_id FK
        text endpoint_id FK
        text sealed_secret
        timestamp created_at
    }
    app_public_hmac_nonces {
        text id PK
        text org_id FK
        text endpoint_id FK
        text key_id FK
        text nonce_digest
        timestamp signed_at
        timestamp accepted_at
        timestamp expires_at
    }
    app_canonical_claims {
        text id PK
        text org_id FK
        text endpoint_id FK
        text ingress_id FK
        text provider_kind
        text provider_instance_id FK
        text resource_type
        text resource_id FK
        text claim_kind
        integer claimed_resource_revision
        timestamp released_at
        timestamp budget_reserved_at "Fresh PostgreSQL admission instant; historical NULL rows retain their original created_at as an explicitly approximate budget-day fallback."
        text control_digest
        timestamp control_expires_at
        timestamp created_at
    }
    app_public_cancellations {
        text id PK
        text org_id FK
        text app_installation_id FK
        text endpoint_id FK
        text claim_id FK
        text original_run_id FK
        text request_key_digest
        text state
        timestamp accepted_at
        timestamp settled_at
    }
    app_public_cancellation_selections {
        text id PK
        text org_id FK
        text cancellation_id FK
        text app_installation_id FK
        text original_run_id FK
        text owner_user_id FK
        text native_binding_id FK
        text consent_digest
        text selection_digest
        jsonb historical_create_pin
        text input_digest
        text original_output_digest
        text cancel_run_id FK
        timestamp created_at
    }
    app_resource_access_grants {
        text id PK
        text org_id FK
        text owner_user_id FK
        text recipient_user_id FK
        text app_installation_id
        text resource_binding_id FK
        text checkpoint_id FK
        text projection_id FK
        text review_digest
        jsonb snapshot
        bigserial accepted_sequence
        timestamptz accepted_at
        timestamptz expires_at
        timestamptz revoked_at
        text revoked_by_user_id
    }
    app_private_mcp_grants {
        text id PK
        text org_id FK
        text owner_user_id FK
        text subject_user_id FK
        text mcp_token_id FK
        text app_installation_id
        text resource_binding_id FK
        text checkpoint_id FK
        text projection_id FK
        text review_digest
        jsonb snapshot
        bigserial accepted_sequence
        timestamptz accepted_at
        timestamptz expires_at
        timestamptz revoked_at
        text revoked_by_user_id
    }
    app_private_state_records {
        text org_id PK, FK
        text owner_user_id PK, FK
        text installation_id PK, FK
        text state_key PK
        text record_id PK
        text resource_target_key
        text artifact_digest
        text declaration_digest
        integer revision
        jsonb body
        text key_version
        integer byte_length
        timestamptz created_at
        timestamptz updated_at
        timestamptz expires_at
        timestamptz deleted_at
    }
    email_accounts {
        text org_id PK
        text owner_user_id PK
        text account_id PK
        text provider
        boolean provider_draft_sync_enabled
        text credential_source
        text resource_binding_id
        text retained_resource_binding_id
        text principal_digest
        text state
        integer authority_revision
        timestamptz consented_at
        timestamptz revoked_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_checkpoints {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        integer revision
        jsonb cursor_envelope
        timestamptz observed_at
        timestamptz fresh_until
        text coverage
        boolean has_more
        text sync_state
        timestamptz created_at
        timestamptz updated_at
    }
    email_messages {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text logical_message_id PK
        text revision
        integer storage_revision
        timestamptz received_at
        timestamptz sort_at
        boolean seen
        boolean flagged
        boolean answered
        jsonb summary_envelope
        jsonb source_envelope
        text inventory_generation
        timestamptz deleted_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_memberships {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text logical_message_id PK, FK
        text folder PK
        text inventory_generation
        timestamptz created_at
        timestamptz updated_at
    }
    email_sync_receipts {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text page_id PK
        text request_digest
        text page_digest
        integer checkpoint_revision
        timestamptz observed_at
        boolean complete
        boolean has_more
        timestamptz created_at
        timestamptz updated_at
    }
    email_provider_states {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text state_key PK
        integer storage_revision
        text generation
        jsonb state_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_reading_preferences {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text external_images
        integer revision
        timestamptz created_at
        timestamptz updated_at
    }
    email_imap_credentials {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text principal_digest
        integer connection_revision
        jsonb envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_drafts {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text draft_id PK
        integer revision
        integer content_revision
        text origin
        text created_by_user_id
        text origin_run_id
        text last_edited_by_user_id
        text status
        jsonb content_envelope
        timestamptz submitted_at
        timestamptz archived_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_send_operations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text operation_id PK
        text draft_id FK
        integer draft_revision
        text operation_digest
        text state
        jsonb receipt_envelope
        text initiated_by_user_id
        timestamptz created_at
        timestamptz updated_at
    }
    email_provider_draft_mirrors {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text draft_id PK, FK
        integer desired_revision
        integer synced_revision
        integer generation
        integer envelope_revision
        text claim_token_hash
        integer claim_generation
        integer claim_revision
        text claim_operation_id
        text state
        jsonb provider_envelope
        text last_error_code
        timestamptz created_at
        timestamptz updated_at
    }
    email_account_consents {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text provider
        text principal_digest
        integer revision
        integer native_revision
        integer credential_revision
        text state
        jsonb capabilities
        text bridge_digest
        jsonb bridge_envelope
        timestamptz consented_at
        timestamptz updated_at
        timestamptz revoked_at
    }
    email_native_connection_owners {
        text org_id PK
        text owner_user_id PK
        integer request_count
        integer active_count
        timestamptz created_at
        timestamptz updated_at
    }
    email_native_connection_requests {
        text org_id PK, FK
        text owner_user_id PK, FK
        text request_id PK
        text account_id
        text provider
        text request_digest
        text binding_digest
        text session_digest
        text state
        integer revision
        text state_digest
        jsonb envelope
        timestamptz expires_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_native_account_connections {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text provider
        text principal_digest
        text bridge_digest
        integer revision
        jsonb descriptor_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_operations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text operation_id PK
        text kind
        integer revision
        jsonb content_envelope
        text lease_token
        timestamptz lease_until
        timestamptz created_at
        timestamptz updated_at
    }
    email_attachments {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text draft_id PK, FK
        text file_id PK
        integer upload_revision
        integer size_bytes
        jsonb metadata_envelope
        bytea ciphertext
        text key_version
        text nonce_b64
        text auth_tag_b64
        timestamptz created_at
        timestamptz detached_at
    }
    email_outbox {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text outbox_id PK
        text request_id
        text operation_id
        text draft_id FK
        integer draft_revision
        integer revision
        text state
        text job_id
        timestamptz earliest_dispatch_at
        jsonb snapshot_envelope
        jsonb state_envelope
        text mime_digest
        integer byte_length
        bytea ciphertext
        text key_version
        text nonce_b64
        text auth_tag_b64
        text attempt_id
        text attempt_lock_hash
        timestamptz started_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_outbox_events {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text outbox_id PK, FK
        integer revision PK
        jsonb envelope
        timestamptz created_at
    }
    email_templates {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text template_id PK
        integer revision
        text status
        jsonb content_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_signatures {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        integer revision
        jsonb preferences_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_imap_aliases {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text logical_message_id FK
        text locator_key PK
        integer revision
        jsonb evidence_envelope
        timestamptz updated_at
    }
    email_imap_move_reservations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text logical_message_id PK, FK
        text operation_id FK
        text source_digest
        jsonb reservation_envelope
        timestamptz created_at
    }
    email_timers {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text timer_id PK
        text logical_message_digest
        text kind
        integer revision
        text state
        timestamptz scheduled_at
        text display_timezone
        jsonb timer_envelope
        text firing_job_id
        text firing_lock_token
        timestamptz firing_lease_expires_at
        timestamptz completed_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_presentation_cache {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text cache_key PK
        integer revision
        jsonb content_envelope
        integer byte_length
        timestamptz expires_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_notification_events {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        uuid notification_id PK
        text logical_message_id
        integer revision
        integer authority_revision
        text state
        timestamptz created_at
        timestamptz projected_at
    }
    email_notification_states {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        boolean initial_inventory_complete
        timestamptz created_at
        timestamptz updated_at
    }
    email_calendar_invitations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text invitation_id PK
        integer revision
        integer sequence
        timestamptz dtstamp
        text content_digest
        text logical_message_id
        text source_digest
        jsonb model_envelope
        text event_id FK
        text event_fingerprint
        timestamptz created_at
        timestamptz updated_at
    }
    email_calendar_invitation_actions {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        uuid request_id PK
        text invitation_id FK
        text payload_digest
        integer invitation_revision
        jsonb receipt_envelope
        text outbox_id FK
        timestamptz created_at
    }
    email_context_transfers {
        text org_id PK, FK
        text owner_user_id PK, FK
        text request_id PK
        text request_digest
        text account_id FK
        text logical_message_id
        text source_digest
        text mode
        text review_digest
        text destination_kind
        text destination_id
        text destination_url
        timestamptz created_at
        timestamptz updated_at
    }
    email_rules {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text rule_id PK
        integer revision
        text state
        integer order_index
        jsonb content_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_rule_receipts {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text rule_id PK, FK
        integer rule_revision PK
        text target_digest PK
        integer action_index PK
        text operation_id
        integer storage_revision
        text state
        jsonb receipt_envelope
        timestamptz created_at
        timestamptz updated_at
    }
    email_agent_grants {
        text id PK
        text org_id FK
        text owner_user_id FK
        text account_id FK
        text principal_kind
        text principal_id
        jsonb operations
        integer revision
        timestamptz revoked_at
        timestamptz created_at
        timestamptz updated_at
    }
    email_knowledge_policies {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        boolean enabled
        integer revision
        jsonb destination
        integer history_days
        boolean include_bulk
        jsonb excluded_senders
        jsonb excluded_folders
        timestamptz history_since
        timestamptz created_at
        timestamptz updated_at
    }
    email_knowledge_observations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text logical_message_id PK
        integer storage_revision
        integer policy_revision
        text content_digest
        text thread_key
        text state
        boolean suppressed
        text last_error_code
        timestamptz retry_at
        integer attempts
        timestamptz created_at
        timestamptz updated_at
    }
    email_knowledge_sources {
        text id PK
        text org_id FK
        text owner_user_id FK
        text account_id FK
        text logical_message_id
        text message_revision
        text thread_key
        text content_digest
        integer policy_revision
        text provider
        text source_digest
        text sender
        timestamptz received_at
        boolean partial
        timestamptz created_at
        timestamptz updated_at
    }
    email_knowledge_threads {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        text thread_key PK
        text destination_key PK
        text page_id
        integer page_version
        text last_source_order
        timestamptz created_at
        timestamptz updated_at
    }
    email_knowledge_budgets {
        text org_id PK
        text day PK
        text budget_key PK
        integer reserved_tokens
        timestamptz created_at
        timestamptz updated_at
    }
    app_connection_sessions {
        text org_id PK
        text owner_user_id PK
        text session_id PK
        text web_session_id
        timestamptz access_expires_at
        integer revision
        jsonb record
        timestamptz created_at
        timestamptz updated_at
    }
    app_data_consents {
        text org_id PK
        text owner_user_id PK
        text consent_id PK
        text connection_session_id
        text store_key
        text state
        text consent_digest
        jsonb snapshot
        timestamptz expires_at
        timestamptz created_at
        timestamptz updated_at
    }
    app_data_invocations {
        text org_id PK
        text owner_user_id PK
        text request_id PK
        text connection_session_id
        text operation
        text input_digest
        text state
        integer revision
        text web_session_id
        timestamptz access_expires_at
        timestamptz expires_at
        jsonb context
        jsonb result
        timestamptz created_at
        timestamptz updated_at
    }
    app_resource_observation_consents {
        text org_id PK
        text owner_user_id
        text consent_id PK
        text resource_binding_id
        text connection_session_id
        text state
        text consent_digest
        jsonb snapshot
        timestamptz expires_at
        timestamptz created_at
        timestamptz updated_at
    }
    app_resource_observation_runs {
        text org_id PK
        text run_id PK
        text consent_id
        text consent_digest
        timestamptz created_at
        timestamptz updated_at
    }
    app_connection_forms {
        text org_id PK
        text owner_user_id PK
        text session_id PK
        text web_session_id
        timestamptz access_expires_at
        integer revision
        jsonb record
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_capture_policies {
        text id PK
        text org_id FK
        text owner_user_id FK
        text installation_id FK
        text account_ref
        text store_key
        integer revision
        boolean enabled
        jsonb input
        timestamptz history_since
        text scan_checkpoint_id
        text scan_after_projection_id
        timestamptz retired_at
        text retired_reason
        text last_error
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_capture_claims {
        text id PK
        text org_id FK
        text owner_user_id
        text installation_id
        text account_ref
        text store_key
        text policy_id FK
        integer policy_revision
        text thread_key
        text body_digest
        jsonb target
        text state
        text page_id
        text error_code
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_capture_threads {
        text id PK
        text org_id
        text owner_user_id
        text installation_id
        text account_ref
        text store_key
        text thread_key
        text destination_key
        text page_id FK
        integer page_version
        timestamptz last_observed_at
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_capture_sources {
        text id PK
        text org_id FK
        text owner_user_id
        text installation_id
        text account_ref
        text store_key
        text claim_id FK
        text page_id FK
        jsonb target
        jsonb publication
        text thread_key
        text body_digest
        timestamptz observed_at
        boolean partial
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_destination_receipts {
        text org_id PK, FK
        text owner_user_id PK, FK
        text request_id PK
        text request_digest
        text review_digest
        jsonb destination
        jsonb source
        jsonb result
        timestamptz created_at
        timestamptz updated_at
    }
    app_integration_notification_policies {
        text id
        text org_id PK, FK
        text owner_user_id PK, FK
        text installation_id PK
        text account_ref PK
        text store_key PK
        boolean enabled
        integer revision
        timestamptz created_at
        timestamptz updated_at
    }
    app_connection_consent_continuations {
        text org_id PK
        text owner_user_id PK
        text session_id PK
        text source_session_id
        text resource_binding_id
        text web_session_id
        timestamptz access_expires_at
        text authority_operation
        timestamptz authority_expires_at
        integer revision
        jsonb record
        timestamptz created_at
        timestamptz updated_at
    }
    app_restore_access_reviews {
        text id PK
        text org_id FK
        text owner_user_id FK
        text installation_id
        text target_key
        text session_id
        text account_ref
        jsonb offered_grants
        text review_digest
        timestamptz issued_at
        timestamptz expires_at
        timestamptz consumed_at
        timestamptz created_at
        timestamptz updated_at
    }
    native_email_adoptions {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        uuid adoption_id
        integer revision
        text state
        integer source_generation
        integer source_authority_revision
        integer source_consent_revision
        text source_principal_digest
        text source_binding_digest
        text snapshot_digest
        text destination_binding_digest
        text destination_admission_digest
        text runtime_receipt_digest
        timestamptz created_at
        timestamptz updated_at
    }
    native_email_adoption_destinations {
        text org_id PK, FK
        text owner_user_id PK, FK
        text account_id PK, FK
        uuid adoption_id
        text session_id
        jsonb admission
        jsonb profile
        jsonb receipt
        timestamptz created_at
        timestamptz published_at
    }
```
