```mermaid
erDiagram
    user ||--o{ session : "(id) - (user_id)"
    organization |o--o{ session : "(id) - (active_organization_id)"
    user ||--o{ account : "(id) - (user_id)"
    project |o--o{ folder : "(id) - (project_id)"
    user ||--o{ folder : "(id) - (user_id)"
    workspace |o--o{ folder : "(id) - (workspace_id)"
    folder |o--o{ folder : "(id) - (parent_id)"
    folder |o--o{ folder : "(id, workspace_id, resource_type) - (parent_id, workspace_id, resource_type)"
    folder |o--o{ folder : "(id, project_id, resource_type) - (parent_id, project_id, resource_type)"
    user ||--o{ pinned_item : "(id) - (user_id)"
    workspace ||--o{ pinned_item : "(id) - (workspace_id)"
    user ||--o{ workspace_visit : "(id) - (user_id)"
    workspace ||--o{ workspace_visit : "(id) - (workspace_id)"
    user ||--o{ workflow : "(id) - (user_id)"
    workspace |o--o{ workflow : "(id) - (workspace_id)"
    folder |o--o{ workflow : "(id) - (folder_id)"
    workflow ||--o{ workflow_blocks : "(id) - (workflow_id)"
    workflow ||--o{ workflow_edges : "(id) - (workflow_id)"
    workflow_blocks ||--o{ workflow_edges : "(id) - (source_block_id)"
    workflow_blocks ||--o{ workflow_edges : "(id) - (target_block_id)"
    workflow ||--o{ workflow_subflows : "(id) - (workflow_id)"
    workflow |o--o{ workflow_execution_snapshots : "(id) - (workflow_id)"
    workflow |o--o{ workflow_execution_logs : "(id) - (workflow_id)"
    workspace ||--o{ workflow_execution_logs : "(id) - (workspace_id)"
    workflow_execution_snapshots ||--o{ workflow_execution_logs : "(id) - (state_snapshot_id)"
    workflow_deployment_version |o--o{ workflow_execution_logs : "(id) - (deployment_version_id)"
    workspace ||--o{ execution_large_values : "(id) - (workspace_id)"
    workflow |o--o{ execution_large_values : "(id) - (workflow_id)"
    workspace ||--o{ execution_large_value_references : "(id) - (workspace_id)"
    workflow |o--o{ execution_large_value_references : "(id) - (workflow_id)"
    workspace ||--o{ execution_large_value_dependencies : "(id) - (workspace_id)"
    workflow ||--o{ paused_executions : "(id) - (workflow_id)"
    paused_executions ||--o{ resume_queue : "(id) - (paused_execution_id)"
    user ||--o| environment : "(id) - (user_id)"
    organization ||--o| organization_secret_source : "(id) - (organization_id)"
    organization_secret_source ||--o{ organization_secret : "(id) - (source_id)"
    user |o--o{ organization_secret : "(id) - (owner_user_id)"
    organization ||--o{ sim_link : "(id) - (organization_id)"
    user |o--o{ sim_link : "(id) - (created_by)"
    workspace ||--o| workspace_environment : "(id) - (workspace_id)"
    workspace ||--o{ secret_usage : "(id) - (workspace_id)"
    workspace ||--o{ workspace_byok_keys : "(id) - (workspace_id)"
    user |o--o{ workspace_byok_keys : "(id) - (created_by)"
    organization ||--o{ organization_byok_keys : "(id) - (organization_id)"
    user |o--o{ organization_byok_keys : "(id) - (created_by)"
    user ||--o| settings : "(id) - (user_id)"
    workflow |o--o{ workflow_schedule : "(id) - (workflow_id)"
    workflow_deployment_version |o--o{ workflow_schedule : "(id) - (deployment_version_id)"
    workflow_deployment_operation |o--o{ workflow_schedule : "(id) - (deployment_operation_id)"
    user |o--o{ workflow_schedule : "(id) - (source_user_id)"
    workspace |o--o{ workflow_schedule : "(id) - (source_workspace_id)"
    workflow_schedule |o--o{ job_execution_logs : "(id) - (schedule_id)"
    workspace ||--o{ job_execution_logs : "(id) - (workspace_id)"
    workflow ||--o{ webhook : "(id) - (workflow_id)"
    workflow_deployment_version |o--o{ webhook : "(id) - (deployment_version_id)"
    workflow ||--o{ webhook_path_claim : "(id) - (workflow_id)"
    workflow ||--o{ sim_trigger_state : "(id) - (workflow_id)"
    user ||--o{ api_key : "(id) - (user_id)"
    workspace |o--o{ api_key : "(id) - (workspace_id)"
    user |o--o{ api_key : "(id) - (created_by)"
    user ||--o| user_stats : "(id) - (user_id)"
    workspace |o--o{ custom_tools : "(id) - (workspace_id)"
    user |o--o{ custom_tools : "(id) - (user_id)"
    workspace |o--o{ skill : "(id) - (workspace_id)"
    user |o--o{ skill : "(id) - (user_id)"
    skill ||--o{ skill_member : "(id) - (skill_id)"
    user ||--o{ skill_member : "(id) - (user_id)"
    user |o--o{ skill_member : "(id) - (invited_by)"
    workflow ||--o{ chat : "(id) - (workflow_id)"
    user ||--o{ chat : "(id) - (user_id)"
    organization |o--o{ aws_marketplace_license : "(id) - (organization_id)"
    user |o--o{ aws_marketplace_license : "(id) - (linked_by_user_id)"
    organization ||--o{ aws_marketplace_usage_tally : "(id) - (organization_id)"
    organization |o--o{ aws_marketplace_usage_report : "(id) - (organization_id)"
    aws_marketplace_license ||--o{ aws_marketplace_usage_report : "(id) - (license_id)"
    organization ||--o{ organization_search_history : "(id) - (organization_id)"
    user ||--o{ organization_search_history : "(id) - (user_id)"
    user ||--o| member : "(id) - (user_id)"
    organization ||--o{ member : "(id) - (organization_id)"
    member ||--o| member_profile : "(id) - (member_id)"
    member |o--o{ member_profile : "(id) - (manager_member_id)"
    user |o--o{ member_profile : "(id) - (updated_by)"
    organization ||--o| people_directory_sync : "(id) - (organization_id)"
    user |o--o{ people_directory_sync : "(id) - (configured_by)"
    organization ||--o| people_directory_scim_change : "(id) - (organization_id)"
    people_directory_sync ||--o{ people_directory_entry : "(organization_id) - (organization_id)"
    people_directory_sync ||--o{ people_directory_stage : "(organization_id) - (organization_id)"
    people_directory_sync ||--o{ people_directory_page : "(organization_id) - (organization_id)"
    organization ||--o{ organization_member_usage_limit : "(id) - (organization_id)"
    user ||--o{ organization_member_usage_limit : "(id) - (user_id)"
    user |o--o{ organization_member_usage_limit : "(id) - (set_by)"
    organization ||--o| organization_access_request_settings : "(id) - (organization_id)"
    user |o--o{ organization_access_request_settings : "(id) - (updated_by)"
    organization ||--o{ permission_access_request : "(id) - (organization_id)"
    user ||--o{ permission_access_request : "(id) - (requester_id)"
    user |o--o{ permission_access_request : "(id) - (decided_by)"
    user ||--o{ invitation : "(id) - (inviter_id)"
    organization |o--o{ invitation : "(id) - (organization_id)"
    user ||--o{ workspace : "(id) - (owner_id)"
    organization |o--o{ workspace : "(id) - (organization_id)"
    user ||--o{ workspace : "(id) - (billed_account_user_id)"
    workspace |o--o{ workspace : "(id) - (forked_from_workspace_id)"
    project |o--o{ workspace : "(id, organization_scope_key) - (projectId, organization_scope_key)"
    workspace |o--o{ workspace : "(id, projectId) - (forked_from_workspace_id, projectId)"
    organization |o--o{ project : "(id) - (organization_id)"
    user ||--o{ project : "(id) - (owner_id)"
    workspace ||--o{ workspace_fork_resource_map : "(id) - (child_workspace_id)"
    user |o--o{ workspace_fork_resource_map : "(id) - (created_by)"
    workspace ||--o{ workspace_fork_block_map : "(id) - (child_workspace_id)"
    workspace ||--o{ workspace_fork_dependent_value : "(id) - (child_workspace_id)"
    workspace ||--o{ workspace_fork_promote_run : "(id) - (child_workspace_id)"
    user |o--o{ workspace_fork_promote_run : "(id) - (created_by)"
    workspace ||--o{ workspace_fork_workflow_sync : "(id) - (child_workspace_id)"
    workflow ||--o{ workspace_fork_workflow_sync : "(id) - (source_workflow_id)"
    workflow ||--o{ workspace_fork_workflow_sync : "(id) - (target_workflow_id)"
    workspace ||--o{ background_work_status : "(id) - (workspace_id)"
    workflow |o--o{ background_work_status : "(id) - (workflow_id)"
    workspace ||--o{ workspace_operation_receipt : "(id) - (workspace_id)"
    workspace ||--o{ workspace_file : "(id) - (workspace_id)"
    user ||--o{ workspace_file : "(id) - (uploaded_by)"
    workspace ||--o| dashboard : "(id) - (workspace_id)"
    user |o--o{ dashboard : "(id) - (created_by)"
    user |o--o{ dashboard : "(id) - (updated_by)"
    workspace ||--o{ workflow_test : "(id) - (workspace_id)"
    workspace_files ||--o| workflow_test : "(id) - (body_file_id)"
    user |o--o{ workflow_test : "(id) - (created_by_user_id)"
    workflow_test ||--o{ workflow_test_run : "(id) - (test_id)"
    workspace ||--o{ workflow_test_run : "(id) - (workspace_id)"
    user |o--o{ workflow_test_run : "(id) - (triggered_by_user_id)"
    workspace ||--o{ changelog_release : "(id) - (workspace_id)"
    workspace_files ||--o| changelog_release : "(id) - (body_file_id)"
    user |o--o{ changelog_release : "(id) - (created_by)"
    user |o--o{ changelog_release : "(id) - (updated_by)"
    changelog_release ||--o{ changelog_change : "(id) - (release_id)"
    workflow |o--o{ changelog_change : "(id) - (workflow_id)"
    workflow_deployment_version |o--o{ changelog_change : "(id) - (deployment_version_id)"
    copilot_chats |o--o{ changelog_change : "(id) - (chat_id)"
    project |o--o{ workspace_files : "(id) - (project_id)"
    user ||--o{ workspace_files : "(id) - (user_id)"
    workspace |o--o{ workspace_files : "(id) - (workspace_id)"
    organization |o--o{ workspace_files : "(id) - (organization_id)"
    folder |o--o{ workspace_files : "(id) - (folder_id)"
    copilot_chats |o--o{ workspace_files : "(id) - (chat_id)"
    workspace_files ||--o{ workspace_file_search_index : "(id) - (file_id)"
    workspace ||--o{ workspace_file_search_index : "(id) - (workspace_id)"
    workspace ||--o| workspace_file_search_dispatch_queue : "(id) - (workspace_id)"
    workspace_files ||--o{ workspace_file_search_segment : "(id) - (file_id)"
    workspace ||--o{ workspace_file_search_segment : "(id) - (workspace_id)"
    workspace_files ||--o| workspace_file_search_revision : "(id) - (file_id)"
    workspace_file_search_build |o--o{ workspace_file_search_revision : "(id) - (build_id)"
    workspace_file_search_build ||--o{ workspace_file_search_chunk : "(id) - (build_id)"
    workspace_file_search_build ||--o{ file_search_dependency : "(id) - (build_id)"
    workspace_files ||--o| workspace_file_secret_provenance : "(id) - (file_id)"
    workspace_files ||--o| workspace_file_collab_state : "(id) - (file_id)"
    workspace_files ||--o{ workspace_file_version : "(id) - (file_id)"
    workspace |o--o{ workspace_file_version : "(id) - (workspace_id)"
    workspace |o--o{ public_share : "(id) - (workspace_id)"
    user |o--o{ public_share : "(id) - (created_by)"
    invitation ||--o{ invitation_workspace_grant : "(id) - (invitation_id)"
    workspace ||--o{ invitation_workspace_grant : "(id) - (workspace_id)"
    user ||--o{ permissions : "(id) - (user_id)"
    workspace ||--o{ memory : "(id) - (workspace_id)"
    memory ||--o| memory_secret_provenance : "(id) - (memory_id)"
    memory ||--o{ memory_item : "(id) - (memory_id)"
    memory ||--o{ agent_memory_turn : "(id) - (memory_id)"
    workflow ||--o{ agent_memory_turn : "(id) - (workflow_id)"
    memory ||--o{ memory_artifact : "(id) - (memory_id)"
    execution_large_values ||--o{ memory_artifact : "(key) - (key)"
    organization ||--o{ organization_search_integration : "(id) - (organization_id)"
    user ||--o{ knowledge_base : "(id) - (user_id)"
    workspace |o--o| knowledge_base : "(id) - (workspace_id)"
    organization |o--o| knowledge_base : "(id) - (organization_id)"
    folder |o--o{ knowledge_base : "(id) - (folder_id)"
    workspace ||--o{ knowledge_base_import : "(id) - (workspace_id)"
    knowledge_base |o--o{ knowledge_base_import : "(id) - (knowledge_base_id)"
    knowledge_base ||--o{ document : "(id) - (knowledge_base_id)"
    knowledge_connector |o--o{ document : "(id) - (connector_id)"
    user |o--o{ document : "(id) - (uploaded_by)"
    document ||--o| document_secret_provenance : "(id) - (document_id)"
    knowledge_base ||--o{ knowledge_base_tag_definitions : "(id) - (knowledge_base_id)"
    knowledge_base ||--o{ embedding : "(id) - (knowledge_base_id)"
    document ||--o{ embedding : "(id) - (document_id)"
    embedding ||--o| embedding_search : "(id) - (id)"
    document ||--o| knowledge_projection_dirty : "(id) - (document_id)"
    embedding ||--o| embedding_secret_provenance : "(id) - (embedding_id)"
    organization ||--o{ mothership_benchmarks : "(id) - (organization_id)"
    user ||--o{ mothership_benchmarks : "(id) - (user_id)"
    user |o--o{ mothership_benchmarks : "(id) - (run_as_user_id)"
    workspace ||--o{ mothership_benchmarks : "(id) - (source_workspace_id)"
    mothership_benchmarks ||--o{ mothership_benchmark_runs : "(id) - (benchmark_id)"
    user ||--o{ mothership_memory_spaces : "(id) - (user_id)"
    organization ||--o{ mothership_memory_spaces : "(id) - (organization_id)"
    user ||--o{ mothership_memory_selections : "(id) - (user_id)"
    organization ||--o{ mothership_memory_selections : "(id) - (organization_id)"
    mothership_memory_spaces |o--o{ mothership_memory_selections : "(id) - (space_id)"
    user ||--o{ copilot_chats : "(id) - (user_id)"
    workflow |o--o{ copilot_chats : "(id) - (workflow_id)"
    workspace |o--o{ copilot_chats : "(id) - (workspace_id)"
    organization |o--o{ copilot_chats : "(id) - (organization_id)"
    copilot_chats ||--o{ mothership_resource_effects : "(id) - (chat_id)"
    copilot_chats ||--o{ copilot_messages : "(id) - (chat_id)"
    user ||--o{ workflow_checkpoints : "(id) - (user_id)"
    workflow ||--o{ workflow_checkpoints : "(id) - (workflow_id)"
    copilot_chats ||--o{ workflow_checkpoints : "(id) - (chat_id)"
    user ||--o{ copilot_request_stops : "(id) - (user_id)"
    workspace ||--o{ copilot_request_stops : "(id) - (workspace_id)"
    user ||--o{ copilot_organization_request_stops : "(id) - (user_id)"
    organization ||--o{ copilot_organization_request_stops : "(id) - (organization_id)"
    user ||--o{ desktop_devices : "(id) - (user_id)"
    session |o--o{ desktop_devices : "(id) - (session_id)"
    copilot_chats ||--o{ copilot_runs : "(id) - (chat_id)"
    user ||--o{ copilot_runs : "(id) - (user_id)"
    workflow |o--o{ copilot_runs : "(id) - (workflow_id)"
    workspace |o--o{ copilot_runs : "(id) - (workspace_id)"
    organization |o--o{ copilot_runs : "(id) - (organization_id)"
    desktop_devices |o--o{ copilot_runs : "(id) - (desktop_device_id)"
    copilot_runs ||--o{ copilot_run_checkpoints : "(id) - (run_id)"
    copilot_runs ||--o{ copilot_async_tool_calls : "(id) - (run_id)"
    copilot_run_checkpoints |o--o{ copilot_async_tool_calls : "(id) - (checkpoint_id)"
    user ||--o{ copilot_feedback : "(id) - (user_id)"
    copilot_chats ||--o{ copilot_feedback : "(id) - (chat_id)"
    workflow ||--o{ workflow_deployment_version : "(id) - (workflow_id)"
    workflow ||--o| workflow_deployment_operation : "(id) - (workflow_id)"
    workflow_deployment_version ||--o{ workflow_deployment_operation : "(id) - (deployment_version_id)"
    workflow_deployment_version |o--o{ workflow_deployment_operation : "(id) - (previous_active_version_id)"
    workspace |o--o{ mcp_servers : "(id) - (workspace_id)"
    organization |o--o{ mcp_servers : "(id) - (organization_id)"
    credential_group |o--o{ mcp_servers : "(id) - (credential_group_id)"
    user |o--o{ mcp_servers : "(id) - (created_by)"
    mcp_servers ||--o| mcp_server_oauth : "(id) - (mcp_server_id)"
    user |o--o{ mcp_server_oauth : "(id) - (user_id)"
    workspace |o--o{ mcp_server_oauth : "(id) - (workspace_id)"
    organization |o--o{ mcp_server_oauth : "(id) - (organization_id)"
    user ||--o{ sso_provider : "(id) - (user_id)"
    organization |o--o{ sso_provider : "(id) - (organization_id)"
    organization ||--o{ sso_domain : "(id) - (organization_id)"
    user |o--o{ sso_domain : "(id) - (created_by)"
    user |o--o{ oauth_client : "(id) - (user_id)"
    oauth_client ||--o{ oauth_consent : "(client_id) - (client_id)"
    user |o--o{ oauth_consent : "(id) - (user_id)"
    oauth_client ||--o{ oauth_token_family : "(client_id) - (client_id)"
    session |o--o{ oauth_token_family : "(id) - (session_id)"
    user ||--o{ oauth_token_family : "(id) - (user_id)"
    oauth_consent |o--o{ oauth_token_family : "(id) - (consent_id)"
    oauth_client ||--o{ oauth_refresh_token : "(client_id) - (client_id)"
    session |o--o{ oauth_refresh_token : "(id) - (session_id)"
    user ||--o{ oauth_refresh_token : "(id) - (user_id)"
    oauth_token_family ||--o{ oauth_refresh_token : "(id) - (family_id)"
    oauth_client ||--o{ oauth_access_token : "(client_id) - (client_id)"
    session |o--o{ oauth_access_token : "(id) - (session_id)"
    user |o--o{ oauth_access_token : "(id) - (user_id)"
    oauth_refresh_token |o--o{ oauth_access_token : "(id) - (refresh_id)"
    workspace ||--o{ workflow_mcp_server : "(id) - (workspace_id)"
    user ||--o{ workflow_mcp_server : "(id) - (created_by)"
    workflow_mcp_server ||--o{ workflow_mcp_tool : "(id) - (server_id)"
    workflow ||--o{ workflow_mcp_tool : "(id) - (workflow_id)"
    organization ||--o{ custom_block : "(id) - (organization_id)"
    workflow ||--o{ custom_block : "(id) - (workflow_id)"
    user |o--o{ custom_block : "(id) - (created_by)"
    workspace |o--o{ audit_log : "(id) - (workspace_id)"
    user |o--o{ audit_log : "(id) - (actor_id)"
    organization ||--o{ organization_search_mcp_invocation : "(id) - (organization_id)"
    user |o--o{ organization_search_mcp_invocation : "(id) - (user_id)"
    user ||--o{ usage_log : "(id) - (user_id)"
    workspace |o--o{ usage_log : "(id) - (workspace_id)"
    workflow |o--o{ usage_log : "(id) - (workflow_id)"
    workspace |o--o{ credential : "(id) - (workspace_id)"
    organization |o--o{ credential : "(id) - (organization_id)"
    slack_app |o--o{ credential : "(id) - (slack_app_id)"
    account |o--o{ credential : "(id) - (account_id)"
    user |o--o{ credential : "(id) - (env_owner_user_id)"
    credential_group_enrollment |o--o{ credential : "(id) - (credential_group_enrollment_id)"
    mcp_servers |o--o{ credential : "(id) - (mcp_server_id)"
    user |o--o{ credential : "(id) - (created_by)"
    workspace ||--o{ workload_identity_pending_subject : "(id) - (workspace_id)"
    user ||--o{ workload_identity_pending_subject : "(id) - (user_id)"
    workspace |o--o| credential_group : "(id) - (workspace_id)"
    organization |o--o| credential_group : "(id) - (organization_id)"
    user |o--o{ credential_group : "(id) - (created_by)"
    organization |o--o{ slack_app : "(id) - (organization_id)"
    organization ||--o{ slack_search_installation : "(id) - (organization_id)"
    credential ||--o| slack_search_installation : "(id) - (credential_id)"
    slack_app |o--o{ slack_search_installation : "(id) - (slack_app_id)"
    slack_search_installation ||--o{ slack_search_turn : "(id) - (installation_id)"
    credential_group ||--o{ credential_group_enrollment : "(id) - (credential_group_id)"
    user |o--o{ credential_group_enrollment : "(id) - (user_id)"
    user |o--o{ credential_group_enrollment : "(id) - (created_by)"
    organization ||--o{ assistant_tool_permission : "(id) - (organization_id)"
    user ||--o{ assistant_tool_permission : "(id) - (user_id)"
    credential ||--o{ assistant_tool_permission : "(id) - (credential_id)"
    credential ||--o{ credential_member : "(id) - (credential_id)"
    user ||--o{ credential_member : "(id) - (user_id)"
    user |o--o{ credential_member : "(id) - (invited_by)"
    user ||--o{ pending_credential_draft : "(id) - (user_id)"
    workspace |o--o{ pending_credential_draft : "(id) - (workspace_id)"
    organization |o--o{ pending_credential_draft : "(id) - (organization_id)"
    credential |o--o{ pending_credential_draft : "(id) - (credential_id)"
    organization ||--o| permission_group : "(id) - (organization_id)"
    user ||--o{ permission_group : "(id) - (created_by)"
    permission_group ||--o{ permission_group_workspace : "(id) - (permission_group_id)"
    workspace ||--o{ permission_group_workspace : "(id) - (workspace_id)"
    organization ||--o{ permission_group_workspace : "(id) - (organization_id)"
    permission_group ||--o{ permission_group_member : "(id) - (permission_group_id)"
    organization ||--o{ permission_group_member : "(id) - (organization_id)"
    user ||--o{ permission_group_member : "(id) - (user_id)"
    user |o--o{ permission_group_member : "(id) - (assigned_by)"
    workspace |o--o{ resource_policy : "(id) - (workspace_id)"
    organization |o--o{ resource_policy : "(id) - (organization_id)"
    user |o--o{ resource_policy : "(id) - (created_by)"
    user |o--o{ resource_policy : "(id) - (updated_by)"
    knowledge_base ||--o{ knowledge_connector : "(id) - (knowledge_base_id)"
    credential_group |o--o{ knowledge_connector : "(id) - (credential_group_id)"
    knowledge_connector ||--o{ knowledge_connector_partition : "(id) - (connector_id)"
    knowledge_connector ||--o| knowledge_connector_permission_snapshot : "(id) - (connector_id)"
    knowledge_connector_permission_snapshot ||--o{ knowledge_connector_permission_grant : "(connector_id) - (connector_id)"
    workspace |o--o{ knowledge_connector_member : "(id) - (workspace_id)"
    organization |o--o{ knowledge_connector_member : "(id) - (organization_id)"
    knowledge_connector ||--o{ knowledge_connector_member : "(id) - (connector_id)"
    credential ||--o{ knowledge_connector_member : "(id) - (credential_id)"
    document ||--o{ knowledge_document_observation : "(id) - (document_id)"
    knowledge_connector_member ||--o{ knowledge_document_observation : "(id) - (member_id)"
    workspace |o--o{ knowledge_external_directory : "(id) - (workspace_id)"
    organization |o--o{ knowledge_external_directory : "(id) - (organization_id)"
    organization |o--o{ knowledge_external_group : "(id) - (organization_id)"
    workspace |o--o{ knowledge_external_group : "(id) - (workspace_id)"
    knowledge_external_group ||--o{ knowledge_external_group_member : "(id) - (group_id)"
    knowledge_connector ||--o{ knowledge_connector_member_sync_log : "(id) - (connector_id)"
    knowledge_connector ||--o{ knowledge_connector_sync_log : "(id) - (connector_id)"
    workspace ||--o{ user_table_definitions : "(id) - (workspace_id)"
    folder |o--o{ user_table_definitions : "(id) - (folder_id)"
    user ||--o{ user_table_definitions : "(id) - (created_by)"
    user_table_definitions ||--o{ user_table_rows : "(id) - (table_id)"
    workspace ||--o{ user_table_rows : "(id) - (workspace_id)"
    user |o--o{ user_table_rows : "(id) - (created_by)"
    user_table_definitions ||--o{ user_table_row_changes : "(id) - (table_id)"
    user_table_rows ||--o| user_table_row_secret_provenance : "(id) - (row_id)"
    user_table_definitions ||--o| table_views : "(id) - (table_id)"
    workspace ||--o{ table_views : "(id) - (workspace_id)"
    user |o--o{ table_views : "(id) - (created_by)"
    user_table_definitions ||--o| table_jobs : "(id) - (table_id)"
    workspace ||--o{ table_jobs : "(id) - (workspace_id)"
    user_table_definitions ||--o{ table_row_executions : "(id) - (table_id)"
    user_table_rows ||--o{ table_row_executions : "(id) - (row_id)"
    user |o--o{ table_row_executions : "(id) - (capability_governed_user_id)"
    user_table_definitions ||--o{ table_run_dispatches : "(id) - (table_id)"
    workspace ||--o{ table_run_dispatches : "(id) - (workspace_id)"
    user |o--o{ table_run_dispatches : "(id) - (triggered_by_user_id)"
    user |o--o{ table_run_dispatches : "(id) - (capability_governed_user_id)"
    workspace ||--o{ mothership_inbox_allowed_sender : "(id) - (workspace_id)"
    user ||--o{ mothership_inbox_allowed_sender : "(id) - (added_by)"
    workspace ||--o{ mothership_inbox_task : "(id) - (workspace_id)"
    copilot_chats |o--o{ mothership_inbox_task : "(id) - (chat_id)"
    workspace ||--o| mothership_inbox_webhook : "(id) - (workspace_id)"
    organization ||--o{ data_drains : "(id) - (organization_id)"
    user ||--o{ data_drains : "(id) - (created_by)"
    data_drains ||--o{ data_drain_runs : "(id) - (drain_id)"
    workspace ||--o{ workspace_sandbox : "(id) - (workspace_id)"
    user |o--o{ workspace_sandbox : "(id) - (created_by)"
    organization ||--o| scim_connection : "(id) - (organization_id)"
    user |o--o{ scim_connection : "(id) - (created_by)"
    scim_connection ||--o{ scim_credential : "(id) - (connection_id)"
    user |o--o{ scim_credential : "(id) - (revoked_by)"
    user |o--o{ scim_credential : "(id) - (created_by)"
    scim_connection ||--o{ scim_user : "(id) - (connection_id)"
    user ||--o{ scim_user : "(id) - (user_id)"
    scim_connection ||--o{ scim_user_tombstone : "(id) - (connection_id)"
    user ||--o{ scim_user_tombstone : "(id) - (user_id)"
    scim_connection ||--o{ scim_group : "(id) - (connection_id)"
    scim_group ||--o{ scim_group_member : "(id) - (group_id)"
    scim_user ||--o{ scim_group_member : "(id) - (scim_user_id)"
    scim_group ||--o{ scim_group_mapping : "(id) - (group_id)"
    permission_group |o--o{ scim_group_mapping : "(id) - (permission_group_id)"
    workspace |o--o{ scim_group_mapping : "(id) - (workspace_id)"
    user |o--o{ scim_group_mapping : "(id) - (created_by)"
    scim_connection ||--o{ scim_projection_grant : "(id) - (connection_id)"
    scim_user ||--o{ scim_projection_grant : "(id) - (scim_user_id)"
    scim_connection ||--o{ scim_request_log : "(id) - (connection_id)"
    user {
        text id PK
        text name
        text email "Unique byte-for-byte only. The identity an address names is `foldedEmail(email)`, which `user_email_lower_idx` indexes."
        text normalized_email "Legacy signup normalization strips Gmail dots and tags and must never be used for identity. Retained for Better Auth and full-row readers that still select this column; removal needs a separate projection migration."
        boolean email_verified
        text image
        timestamp created_at
        timestamp updated_at
        text stripe_customer_id
        text role
        boolean banned
        text ban_reason
        timestamp ban_expires
        timestamp suspended_at "When set, the account is suspended: sign-in is refused and API keys stop authenticating, while every resource the user owns is left untouched. Deliberately not `banned`. A ban is a platform-admin action whose `user.update.after` hook runs `disableUserResources`, archiving every workspace the user owns and deleting their API keys, and Sim has no server-side unban to reverse it. SCIM `active: false` is a reversible organization-level suspension that must preserve ownership for a later reactivation, so it needs a state of its own."
        text suspension_source "Who suspended the account. Only `scim` exists today; a source only ever lifts its own suspension, so a later source cannot have its suspensions undone by a directory sync."
    }
    session {
        text id PK
        timestamp expires_at
        text token
        timestamp created_at
        timestamp updated_at
        text ip_address
        text user_agent
        text user_id FK
        text active_organization_id FK
        text impersonated_by
    }
    account {
        text id PK
        text account_id
        text provider_id
        text user_id FK
        text access_token
        text refresh_token
        text id_token
        timestamp access_token_expires_at
        timestamp refresh_token_expires_at
        text scope
        text password
        text oauth_config
        timestamp refresh_revoked_at "The provider rejected `refresh_token` as revoked (`invalid_grant` and kin). The marker holds only while `refresh_revoked_token_hash` still fingerprints the stored refresh token, so any writer that stores a new chain supersedes it; a reconnect clears it explicitly, since a provider may reauthorize without issuing a new refresh token."
        text refresh_revoked_code
        text refresh_revoked_token_hash
        timestamp created_at
        timestamp updated_at
    }
    verification {
        text id PK
        text identifier
        text value
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
    }
    folder {
        text id PK
        text project_id FK "File folders may belong to a Project; other resource folders remain workspace-owned."
        folder_resource_type resource_type FK
        text name
        text user_id FK
        text workspace_id FK
        text parent_id FK
        boolean locked
        integer sort_order
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    pinned_item {
        text id PK
        text user_id FK
        text workspace_id FK
        text resource_type "'workflow' | 'file' | 'knowledge_base' | 'table' | 'folder' | 'workspace'"
        text resource_id
        timestamp pinned_at
    }
    workspace_visit {
        text user_id PK, FK
        text workspace_id PK, FK
        timestamp visited_at
    }
    workflow {
        text id PK
        text user_id FK "Creator and owner. Legitimate as ownership: it anchors personal (workspace-less) workflows, cascades the workflow away with the account, and names the owner for webhook config and deploy-as-block resolution. acts as, what it may read, or what it may authorize. The acting principal is `ExecutionMetadata.userId`, which the principal layer (`resolvePrincipalAttribution`) resolves to the caller for a session, personal API key, or delegated run, and to the workspace billing account for a workspace API key, schedule, or webhook. Exactly one execution use survives, carried as `ExecutionMetadata.workflowUserId`: the personal-environment fallback in `executeWorkflowCore`, for runs with no identifiable caller — workspace API keys, schedules, webhooks, and unauthenticated public-API calls. Those have nobody to resolve personal variables as, and a deployed workflow is routinely authored against its owner's personal keys, so dropping the fallback would break them. Workspace variables never fall back here; they always authorize against the actor."
        text workspace_id FK
        text folder_id FK
        integer sort_order
        text name
        text description
        timestamp last_synced
        timestamp created_at
        timestamp updated_at
        boolean is_deployed
        timestamp deployed_at
        boolean is_public_api
        boolean locked
        boolean fork_sync_excluded
        integer run_count
        timestamp last_run_at
        json variables
        timestamp archived_at
    }
    workflow_blocks {
        text id PK
        text workflow_id FK
        text type "'starter', 'agent', 'api', 'function'"
        text name
        decimal position_x
        decimal position_y
        boolean enabled
        boolean horizontal_handles
        boolean is_wide
        boolean advanced_mode
        boolean trigger_mode
        boolean error_enabled
        jsonb retry "Opt-in {@link BlockRetryConfig}; NULL means the block never retries."
        boolean locked
        decimal height
        jsonb sub_blocks
        jsonb outputs
        jsonb data
        timestamp created_at
        timestamp updated_at
    }
    workflow_edges {
        text id PK
        text workflow_id FK
        text source_block_id FK
        text target_block_id FK
        text source_handle
        text target_handle
        timestamp created_at
    }
    workflow_subflows {
        text id PK
        text workflow_id FK
        text type "'loop' or 'parallel'"
        jsonb config
        timestamp created_at
        timestamp updated_at
    }
    waitlist {
        text id PK
        text email
        text status "pending, approved, rejected"
        timestamp created_at
        timestamp updated_at
    }
    workflow_execution_snapshots {
        text id PK
        text workflow_id FK
        text state_hash
        jsonb state_data
        timestamp created_at
    }
    workflow_execution_logs {
        text id PK
        text workflow_id FK
        text workspace_id FK
        text execution_id
        text state_snapshot_id FK
        text deployment_version_id FK
        text level "'info' | 'error'"
        text status "See `PERSISTED_WORKFLOW_EXECUTION_STATUSES` in `apps/sim/lib/logs/types.ts`."
        text trigger "'api' | 'webhook' | 'schedule' | 'manual' | 'chat'"
        timestamp started_at
        timestamp execution_deadline_at "Absolute deadline for the current active attempt; cleared while paused or terminal."
        timestamp ended_at
        integer total_duration_ms "Wall clock from `started_at` for a terminal row; for a `pending` (paused) row, the active duration recorded at the checkpoint, which excludes the time the run sits waiting. Resuming leaves that checkpoint value in place while the row accrues time again, so a `running` row's value is stale until the next terminal write recomputes it."
        jsonb execution_data "Heavy trace data (traceSpans, finalOutput, workflowInput, executionState) is externalized to object storage; this column then holds a slim payload: a `traceStoreRef` (__simLargeValueRef) pointer to the stored object plus inline markers (hasTraceSpans, traceSpanCount, environment, trigger, truncation flags). It also still holds the FULL payload inline for legacy / not-yet-backfilled rows, for the storage-write-failure fallback, and for job_execution_logs. Required — not droppable. Read it via `materializeExecutionData`, which resolves the pointer."
        decimal cost_total "Faithful, write-once projection of the run's usage_log ledger sum (dollars). Backs list cost display/filter/sort without live aggregation; never an independently-computed value (cost_total == SUM(usage_log) for the run)."
        text models_used "Model names used by the run (incl. zero-cost/BYOK), for the v1 model filter."
        jsonb files "File metadata for execution files"
        timestamp created_at
    }
    execution_large_values {
        text key PK
        text workspace_id FK
        text workflow_id FK
        text owner_execution_id
        integer size
        timestamp created_at
        timestamp deleted_at
    }
    execution_large_value_references {
        text key PK
        text execution_id PK
        execution_large_value_reference_source source PK
        text workspace_id FK
        text workflow_id FK
        timestamp created_at
    }
    execution_large_value_dependencies {
        text parent_key PK
        text child_key PK
        text workspace_id FK
        timestamp created_at
    }
    paused_executions {
        text id PK
        text workflow_id FK
        text execution_id
        jsonb execution_snapshot
        jsonb pause_points
        integer total_pause_count
        integer resumed_count
        integer automatic_resume_retry_count
        text status
        jsonb metadata
        timestamp paused_at
        timestamp updated_at
        timestamp expires_at
        timestamp next_resume_at "Earliest `resumeAt` across this row's time-based pause points. NULL for human-only pauses."
    }
    resume_queue {
        text id PK
        text paused_execution_id FK
        text parent_execution_id
        text new_execution_id
        text context_id
        jsonb resume_input
        text status
        timestamp queued_at
        timestamp claimed_at
        timestamp completed_at
        text failure_reason
    }
    client_credential_token {
        text id PK
        text encrypted_value
        text access_token_digest
        timestamp expires_at
        timestamp updated_at
    }
    environment {
        text id PK "Use the user id as the key"
        text user_id FK
        json variables
        timestamp updated_at
    }
    organization_secret_source {
        text id PK
        text organization_id FK
        text mode
        timestamp created_at
        timestamp updated_at
    }
    organization_secret {
        text id PK
        text source_id FK
        text owner_user_id FK
        text name
        text encrypted_value
        timestamp updated_at
    }
    sim_link {
        text id PK
        text organization_id FK
        text name
        text normalized_name
        text url
        text default_url
        text description
        boolean unlisted
        boolean restrict_edit
        integer visit_count
        timestamp last_visited_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    workspace_environment {
        text id PK
        text workspace_id FK
        json variables
        timestamp created_at
        timestamp updated_at
    }
    secret_usage {
        text id PK
        text workspace_id FK
        text secret_name "Matches `credential.envKey`; the trail is keyed by name, not by credential row."
        secret_usage_scope secret_scope
        text secret_owner_user_id "Whose personal secret this was; empty for a workspace one, which the workspace owns. Two people can hold personal secrets under the same name, and a personal secret shared with the workspace resolves for callers who do not own it, so name and scope alone do not identify a secret. Without this column one person's trail would show another's runs. Not the same as `actorUserId`: a scheduled run resolves the workflow owner's personal slice under the workspace's execution actor."
        secret_usage_source source
        text workflow_id "Empty for a Copilot or MCP resolution, which has no workflow. Empty string rather than null because both this and `actorUserId` sit inside the unique key below, and Postgres treats nulls as distinct — two Copilot rows would never collide, so the upsert would insert forever instead of incrementing. A sentinel keeps the upsert identity explicit and null-free rather than relying on nullable uniqueness semantics. Deliberately not a foreign key, and neither is `actorUserId`. An `onDelete: 'set null'` would rewrite a key column, so two rows differing only by the deleted id would collide and an ordinary workflow or account deletion would fail on this constraint. They are historical facts in a usage ledger rather than live references, so they are stored as plain ids and joined leniently; a row outliving its workflow is the point of a trail."
        text actor_user_id "Whose access authorized the resolution — the run's actor; empty when there is none."
        date usage_date "UTC day bucket."
        integer use_count
        timestamp last_used_at
        text last_execution_id "Deep-links the most recent run in Logs, where the block and its code are visible."
        text last_trigger "The surface the most recent run came in through (`api`, `webhook`, `schedule`, `manual`, `chat`, `copilot`). There is deliberately no separate 'triggered by' column: for every trigger kind the executor can name a caller, that caller *is* `actorUserId`, and for the rest (schedule, webhook, workspace key) no human triggered the run at all."
        timestamp created_at
        timestamp updated_at
    }
    workspace_byok_keys {
        text id PK
        text workspace_id FK
        text provider_id
        text encrypted_api_key
        text name
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    organization_byok_keys {
        text id PK
        text organization_id FK
        text provider_id
        text encrypted_api_key
        text name
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    settings {
        text id PK "Use the user id as the key"
        text user_id FK
        text theme
        boolean auto_connect
        boolean telemetry_enabled
        json email_preferences
        boolean billing_usage_notifications_enabled
        boolean super_user_mode_enabled
        text mothership_environment
        boolean error_notifications_enabled
        integer snap_to_grid_size "0 = off, 10-50 = grid size"
        boolean show_action_bar
        boolean auto_focus_on_click
        text timezone
        jsonb copilot_auto_allowed_tools "Copilot auto-allowed integration tools - array of tool IDs that can run without confirmation"
        text last_active_workspace_id
        timestamp updated_at
    }
    workflow_schedule {
        text id PK
        text workflow_id FK
        text deployment_version_id FK
        text deployment_operation_id FK
        text block_id
        text cron_expression
        timestamp next_run_at
        timestamp last_ran_at
        timestamp last_queued_at
        text trigger_type "'manual', 'webhook', 'schedule'"
        text timezone
        integer failed_count
        integer infra_retry_count
        text status "'active', 'disabled', or 'completed'"
        timestamp last_failed_at
        text source_type "'workflow' or 'job'"
        text job_title
        text prompt
        text lifecycle "'persistent' or 'until_complete'"
        text success_condition
        integer max_runs
        integer run_count
        text source_chat_id
        text source_task_name
        text source_user_id FK
        text source_workspace_id FK
        text secret_scope
        jsonb mounted_secrets
        jsonb job_history
        jsonb contexts "`@`-mentioned resources / `/`-invoked skills captured with the prompt, resolved into the agent run at fire time."
        jsonb excluded_dates "ISO timestamps of recurring occurrences the user deleted individually (EXDATE); the executor skips them."
        timestamp ends_at "Recurrence end boundary: the schedule completes once its next run would fall after this instant."
        decimal run_budget_usd "A job's spend cap per run in USD; the run pauses once its settled cost reaches it. Null takes the default cap. Unused by workflow schedules."
        timestamp archived_at
        timestamp created_at
        timestamp updated_at
    }
    job_execution_logs {
        text id PK
        text schedule_id FK
        text workspace_id FK
        text execution_id
        text level
        text status
        text trigger
        timestamp started_at
        timestamp ended_at
        integer total_duration_ms
        jsonb execution_data
        jsonb cost
        timestamp occurrence_at "The schedule occurrence this run is for. With the schedule it names the run once: `execution_id` is derived from both and is the job turn's run id. Null on legacy rows."
        text outcome "'met' | 'not_met' | 'blocked' | 'failed' | 'skipped_busy' | 'skipped_budget' | 'interrupted'"
        text summary
        text billing_request_id "The billing request the run's charges settle under, recorded when its first leg is admitted."
        timestamp created_at
    }
    webhook {
        text id PK
        text workflow_id FK
        text deployment_version_id FK
        text registration_status
        integer registration_generation
        text config_fingerprint
        timestamp prepared_at
        text block_id
        text path "URL-addressable webhook path. NULL for shared-app providers (e.g. the native Slack and TikTok triggers) whose events arrive on a single shared endpoint and route by `routingKey` instead of a per-workflow path."
        text routing_key "Tenant routing key for shared-app providers, such as Slack `team_id` or TikTok `open_id`, derived server-side from the connected credential at deploy time — never user input. Inbound events match on this after HMAC verification."
        text provider "e.g., 'whatsapp', 'github', etc."
        json provider_config
        boolean is_active
        integer failed_count "Track consecutive failures"
        timestamp last_failed_at
        timestamp archived_at
        timestamp created_at
        timestamp updated_at
    }
    webhook_path_claim {
        text path PK
        text workflow_id FK
        integer generation
        timestamp created_at
        timestamp updated_at
    }
    sim_trigger_state {
        text workflow_id PK, FK
        text block_id PK
        text scope_key PK
        timestamp last_fired_at
        timestamp updated_at
    }
    api_key {
        text id PK
        text user_id FK
        text workspace_id FK "Only set for workspace keys"
        text created_by FK
        text name
        text key
        text key_hash
        text type
        api_key_access access
        timestamp last_used
        timestamp created_at
        timestamp updated_at
        timestamp expires_at
    }
    user_stats {
        text id PK
        text user_id FK
        decimal current_usage_limit "Default $5 (1,000 credits) for free plan, null for team/enterprise"
        timestamp usage_limit_updated_at
        decimal last_period_cost "Previous-period usage; written by the cycle-close sweep from ledger sums."
        decimal billed_overage_this_period "Threshold/final billing tracker. Incremented when threshold billing collects overage mid-period; reset to zero by the cycle-close sweep at period rollover. It is not incremented by the ordinary per-usage ledger write path. Amount of overage already billed via threshold billing"
        decimal credit_balance "Credit balance tracker. Still debited/credited by billing lifecycle paths and threshold/final overage collection. It is not a per-usage aggregate counter."
        decimal last_period_copilot_cost "Previous-period Copilot cost; written by the cycle-close sweep from copilot-source ledger sums."
        bigint storage_used_bytes "Storage upload/delete hot-path tracker for personal plans. This remains a direct aggregate write for personal file storage changes; org-scoped storage writes update `organization.storageUsedBytes`."
        boolean free_tier_eligible "Whether the personal free-tier allowance counts toward this user's limit. False when the account reuses an identity from a recently deleted account (see `deletedAccountIdentity`): the free allowance is granted once per person, not once per account row. Prepaid credit and paid plans are unaffected."
        boolean billing_blocked
        billing_blocked_reason billing_blocked_reason
        jsonb limit_notifications "Highest usage-limit threshold already emailed per category (e.g. `{ storage: 80, tables: 100 }`). Prevents re-spamming the same warning; re-arms when usage drops back below the re-arm band. Keyed by limit category ('storage' | 'tables'); seats live on `organization`. `credits` instead holds the threshold emailed for the billing period and limit in `creditsPeriod` (start, epoch seconds) and `creditsLimit` (cents), so a new period or a changed limit re-arms it without a reset (see `claimCreditsThreshold`). Dedup granularity is per billing account per category — intentionally NOT per table, so a user hitting the row limit on several tables gets one 'tables' warning, not one per table (the email still names the table that triggered it)."
    }
    deleted_account_identity {
        text identity_hash PK "HMAC-SHA256 (hex) of `<kind>:<canonical identity>`."
        timestamp deleted_at
        timestamp expires_at
    }
    custom_tools {
        text id PK
        text workspace_id FK
        text user_id FK
        text title
        json schema
        text code
        timestamp created_at
        timestamp updated_at
    }
    skill {
        text id PK
        text workspace_id FK
        text user_id FK
        text name
        text description
        text content
        timestamp created_at
        timestamp updated_at
    }
    skill_member {
        text id PK
        text skill_id FK
        text user_id FK
        text invited_by FK
        timestamp created_at
        timestamp updated_at
    }
    subscription {
        text id PK
        text plan
        text reference_id
        text stripe_customer_id
        text stripe_subscription_id
        text status
        timestamp period_start
        timestamp period_end
        boolean cancel_at_period_end
        timestamp cancel_at
        timestamp canceled_at
        timestamp ended_at
        integer seats
        timestamp trial_start
        timestamp trial_end
        text billing_interval
        text stripe_schedule_id
        json metadata
        timestamp last_closed_period_start "Durable cycle-close marker: the `periodStart` of the most recent period whose close (final overage collection, `billedOverageThisPeriod` reset, last-period bookkeeping) has been committed. The daily cycle-close sweep closes the previous period whenever this lags the row's `periodStart`, then advances it. Null = never initialized; the first sweep initializes it to the current `periodStart` without billing so historical periods are never retroactively closed. A deleted subscription's terminal settlement advances it to `periodEnd`: every period ending at or before the marker is settled, and a later charge into one is refused."
    }
    rate_limit_bucket {
        text key PK
        decimal tokens
        timestamp last_refill_at
        timestamp blocked_until
        jsonb capacity_state "Bounded adaptive provider budgets and expiring request leases; ordinary buckets leave it null."
        timestamp updated_at
    }
    chat {
        text id PK
        text workflow_id FK
        text user_id FK
        text identifier
        text title
        text description
        boolean is_active
        json customizations
        text auth_type "'public', 'password', 'email', 'sso'"
        text password "Stored hashed, populated when authType is 'password'"
        json allowed_emails "Array of allowed emails or domains when authType is 'email' or 'sso'"
        json output_configs "Array of {blockId, path} objects"
        boolean include_thinking "When true, public chat SSE exposes provider thinking events. Independent of the `X-Sim-Stream-Protocol` header, which governs answer-text cadence rather than frame exposure. Default off — never derived from auth type or isSecureMode."
        boolean include_tool_calls "When true, public chat SSE exposes tool lifecycle events. Independent of includeThinking and of the protocol header."
        timestamp archived_at
        timestamp created_at
        timestamp updated_at
    }
    organization {
        text id PK
        text name
        text slug
        text logo
        json metadata
        json session_policy_settings
        integer security_policy_version "Monotonic counter embedded in the Better Auth cookie-cache version for this org's members. Bumped on any security-policy change or org-wide session revocation so every cached session cookie in the org is invalidated (falls through to a DB session read) within the policy cache TTL instead of the 24h cookie-cache lifetime."
        boolean require_sso "Whether members must sign in through this organization's identity provider. Checked only when a session is created, so turning it on ends no session that already exists; signing everyone out stays the separate revoke action. Owners keep password sign-in as a break-glass path for a broken identity provider."
        json whitelabel_settings
        json data_retention_settings
        decimal org_usage_limit
        bigint storage_used_bytes "Storage upload/delete hot-path tracker for org-scoped plans. This remains a direct aggregate write for organization file storage changes; personal storage writes update `user_stats.storageUsedBytes`."
        jsonb limit_notifications "Highest usage-limit threshold already emailed per category for this org (e.g. `{ seats: 80, storage: 100 }`). Mirrors `user_stats.limitNotifications` for org-scoped (pooled) limits. Re-arms when usage drops below the re-arm band; `credits` instead re-arms with a new billing period or a changed limit (see `claimCreditsThreshold`)."
        decimal credit_balance "Organization credit balance tracker. Still debited/credited by billing lifecycle paths and threshold/final overage collection. It is not a per-usage aggregate counter."
        timestamp created_at
        timestamp updated_at
    }
    aws_marketplace_license {
        text id PK
        text license_arn
        text product_code
        text customer_aws_account_id
        text agreement_id
        timestamp agreement_seen_at "When the source of `agreement_id` was true, so a redelivered old event cannot roll it back."
        text renewed_agreement_id "The agreement that replaced this license's, from a Purchase Agreement Created event."
        text organization_id FK
        aws_marketplace_license_status status
        text dimension
        integer quantity
        timestamp expires_at
        timestamp last_verified_at
        timestamp last_verification_attempt_at "Every verification attempt, failed ones too, so a failing license cannot starve the rest."
        timestamp license_event_received_at "AWS says not to activate a subscription until this event has arrived, so staff can see it."
        text claim_token_hash
        timestamp claim_expires_at
        text linked_by_user_id FK
        timestamp organization_linked_at "When `organization_id` last changed to its current value; usage counts from here."
        timestamp ended_at "When a read first found the license no longer granting access; cleared if it returns."
        text contact_email "Whom to contact about provisioning, as the buyer entered it during setup."
        text company_name
        timestamp created_at
        timestamp updated_at
    }
    aws_marketplace_usage_tally {
        text organization_id PK, FK
        text dimension PK
        timestamp period_start PK
        timestamp counted_from
        decimal cost_usd
        integer units
        timestamp updated_at
    }
    aws_marketplace_usage_report {
        text id PK
        text organization_id FK "Kept null when the organization is deleted, so usage already charged is still sent."
        text license_id FK
        text dimension
        timestamp hour_start
        integer quantity
        aws_marketplace_usage_report_status status
        timestamp sent_at "When the record first went to AWS. AWS keeps the first record it accepts for a license, dimension and hour, so the quantity is fixed from then on."
        text metering_record_id
        text error "Why AWS refused the record, or why it can no longer be sent."
        timestamp reported_at
        timestamp created_at
        timestamp updated_at
    }
    organization_search_history {
        text organization_id PK, FK
        text user_id PK, FK
        jsonb sources
        jsonb queries
    }
    member {
        text id PK
        text user_id FK
        text organization_id FK
        text role "'admin' or 'member' - team-level permissions only"
        timestamp created_at
    }
    member_profile {
        text member_id PK, FK
        text title
        text department
        text location
        date start_date
        text manager_member_id FK
        text updated_by FK
        timestamp updated_at
    }
    people_directory_sync {
        text organization_id PK, FK
        text source
        text source_target
        text client_id
        text encrypted_credential
        text configured_by FK
        text lease_token
        timestamp lease_expires_at
        text run_id
        timestamp run_started_at
        text phase
        text cursor
        integer users_read
        integer retry_count
        timestamp next_sync_at
        timestamp last_attempt_at
        timestamp last_success_at
        timestamp last_failed_at
        text last_error
    }
    people_directory_scim_change {
        text organization_id PK, FK
        timestamp changed_at
    }
    people_directory_entry {
        text id PK
        text organization_id FK
        text source
        text external_id
        text email
        text name
        text title
        text department
        text location
        text manager_external_id
        timestamp created_at
    }
    people_directory_stage {
        text organization_id PK, FK
        text run_id PK
        text external_id PK
        text proposed_id
        text email
        text name
        text title
        text department
        text location
        text manager_external_id
        text reports_cursor
        boolean reports_complete
    }
    people_directory_page {
        text organization_id PK, FK
        text run_id PK
        text page_hash PK
    }
    organization_member_usage_limit {
        text id PK
        text organization_id FK
        text user_id FK
        decimal usage_limit
        text set_by FK "Admin who set the cap (audit only). Soft FK: nulled if that user is deleted so the member's limit row survives — never cascade-deleted."
        timestamp created_at
        timestamp updated_at
    }
    organization_access_request_settings {
        text organization_id PK, FK
        boolean allow_requests
        timestamp updated_at
        text updated_by FK
    }
    permission_access_request {
        text id PK
        text organization_id FK
        text requester_id FK
        text workspace_id
        text scope_key
        text target_key
        jsonb target
        text target_label
        text membership_id
        text group_id
        text group_name
        text reason
        text status
        text decision_reason
        text decided_by FK
        jsonb decision
        timestamp created_at
        timestamp updated_at
        timestamp decided_at
    }
    invitation {
        text id PK
        invitation_kind kind
        text email
        text inviter_id FK
        text organization_id FK
        invitation_membership_intent membership_intent
        text role
        invitation_status status
        text token
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
    }
    workspace {
        text id PK
        text name
        text color
        text logo_url
        text owner_id FK "checks. The owner→admin derivation is redundant: every workspace owner already has an explicit `admin` row in `permissions` (verified across all production workspaces) and all creation paths add one. Retained only as the lifecycle anchor — `onDelete: 'cascade'` cleans up a user's workspaces on account deletion — and the ownership-transfer target when an owner is removed. For admin checks use explicit `permissions` rows; for the workspace's principal billing identity use `billedAccountUserId`. DO NOT DELETE."
        text organization_id FK
        text organization_scope_key FK "Encodes nullable organization identity for the composite membership foreign key."
        workspaceProjectIdColumn projectId FK
        workspace_mode workspace_mode
        text billed_account_user_id FK
        bigint storage_used_bytes "Durable workspace-first storage ledger. Invariant: this non-negative total and the currently routed payer aggregate change atomically while the workspace row is locked. A payer identity change moves this entire total old payer -> new payer in the same transaction."
        boolean allow_personal_api_keys
        boolean inbox_enabled
        text inbox_address
        text inbox_provider_id
        text inbox_secret_scope
        jsonb inbox_mounted_secrets
        timestamp archived_at
        timestamp organization_assigned_at
        text forked_from_workspace_id FK
        boolean fork_sync_new_workflows_excluded "Whether a newly created workflow in this workspace starts outside fork sync. `false` (default): new workflows join sync once deployed. `true`: they land with `workflow.forkSyncExcluded` set and are opted in on the Forks page. Uniform across a fork lineage (written to every member, inherited by new forks) and forward-only: flipping it never rewrites an existing workflow's `forkSyncExcluded`."
        timestamp created_at
        timestamp updated_at
    }
    project {
        text id PK
        text name
        text organization_id FK
        text organization_scope_key "Encodes nullable organization identity for the composite membership foreign key."
        text owner_id FK "Lifecycle owner for personal and organization Projects; never an implicit access grant."
        timestamp archived_at
        timestamp created_at
        timestamp updated_at
    }
    workspace_fork_resource_map {
        text id PK
        text child_workspace_id FK
        workspace_fork_resource_type resource_type
        text parent_resource_id
        text child_resource_id
        text created_by FK "SET NULL (not CASCADE): deleting the creating user must not delete the fork's identity mappings, which the edge depends on for every future promote."
        timestamp created_at
        timestamp updated_at
    }
    workspace_fork_block_map {
        text id PK
        text child_workspace_id FK
        text parent_workflow_id
        text parent_block_id
        text child_workflow_id
        text child_block_id
        timestamp created_at
        timestamp updated_at
    }
    workspace_fork_dependent_value {
        text id PK
        text child_workspace_id FK
        text target_workflow_id
        text target_block_id
        text sub_block_key
        text value
        timestamp created_at
        timestamp updated_at
    }
    workspace_fork_promote_run {
        text id PK
        text child_workspace_id FK
        text source_workspace_id
        text target_workspace_id
        workspace_fork_promote_direction direction
        jsonb snapshot
        text created_by FK "SET NULL (not CASCADE): deleting the creating user must not delete a pending undo point for a target workspace."
        timestamp created_at
    }
    workspace_fork_workflow_sync {
        text deployment_operation_id PK
        text child_workspace_id FK
        text source_workflow_id FK
        text target_workflow_id FK
        text source_deployment_version_id
        bigint sequence
        text promote_run_id
        timestamp activated_at
        text rollback_operation_id
        timestamp rolled_back_at
    }
    background_work_status {
        text id PK
        text workspace_id FK
        text workflow_id FK
        background_work_kind kind
        background_work_status_value status
        text message
        text error
        jsonb metadata
        timestamp started_at
        timestamp completed_at
        timestamp updated_at
    }
    workspace_operation_receipt {
        text id PK
        text workspace_id FK
        text request_id
        text request_hash
        text kind
        jsonb report
        timestamp created_at
        timestamp updated_at
    }
    workspace_file {
        text id PK
        text workspace_id FK
        text name
        text key
        integer size
        text type
        text uploaded_by FK
        timestamp deleted_at
        timestamp uploaded_at
    }
    dashboard {
        text id PK
        text workspace_id FK
        text content
        integer revision
        text created_by FK
        text updated_by FK
        timestamp created_at
        timestamp updated_at
    }
    workflow_test {
        text id PK
        text workspace_id FK
        text name
        text title
        text description
        text body_file_id FK
        jsonb cases "Every `describe > it` path in the file with its mode (`run`, `only`, `skip`)."
        text source_hash "sha256 of the source `cases` was read from; a run reporting another hash is stale."
        text created_by_user_id FK
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    workflow_test_run {
        text id PK
        text test_id FK
        text workspace_id FK
        text version
        text status
        integer passed
        integer failed
        integer skipped
        jsonb report "The sandbox report: every case's checks, judge reasons, logs, and executions."
        text error "Why the file produced no report: it did not load, or the run itself failed."
        text source_hash "sha256 of the source this run executed; null when it failed before reading the file."
        jsonb ran_against "`{ workflowId, deploymentVersionId }` for each workflow the run executed; null draft ids."
        jsonb progress "Each case's status (`running`, `pass`, `fail`, `skip`) by `describe > it` path, as it runs."
        jsonb triggered_by_actor
        text triggered_by_user_id FK
        timestamp started_at
        timestamp completed_at
    }
    changelog_release {
        text id PK
        text workspace_id FK
        text title
        integer version_major
        integer version_minor
        integer version_patch
        text bump_reason
        text body_file_id FK
        integer revision
        timestamp published_at
        text created_by FK
        text updated_by FK
        timestamp created_at
        timestamp updated_at
    }
    changelog_change {
        text id PK
        text release_id FK
        integer position
        text text
        text workflow_id FK
        text deployment_version_id FK
        uuid chat_id FK
    }
    workspace_files {
        text id PK
        text project_id FK "Direct owner, independent of creator, caller, and billing payer."
        text key
        text user_id FK
        text workspace_id FK
        text organization_id FK
        text folder_id FK
        text context "'workspace', 'mothership', 'copilot', 'chat', 'knowledge-base', 'profile-pictures', 'general', 'execution'"
        uuid chat_id FK
        text message_id "Logical id of the copilot message this file was born in (the user message the upload was attached to). Plain text with no FK: message ids are only unique per chat — the same id legitimately exists in the source chat and every fork of it, which is what lets a fork's 'copy files at-or-before this message' cut match rows in both. NULL means 'birth unknown / not tracked': rows predating this column and contexts that don't stamp it. Nulled together with chatId when a file is materialized to the workspace."
        text original_name
        text display_name "Collision-disambiguated name exposed to the copilot VFS as `uploads/<displayName>`. For mothership chat uploads, identical originalNames within a chat get suffixed `(2)`, `(3)`, ... in upload order so the VFS path is unique per chat. NULL on legacy rows that predate this column — readers must coalesce to originalName. Stable for the row's lifetime; the partial unique index below enforces uniqueness for new (non-NULL) rows. NULLs are treated as distinct in PG unique indexes, so legacy collisions remain (acceptable: those uploads have already happened)."
        text content_type
        bigint size_bytes "Exact byte size."
        integer width "Intrinsic pixel dimensions of an image file, captured lazily on first view (and stored so later views reserve layout space before the image loads, via aspect-ratio). NULL for non-images and for rows not yet backfilled. Purely a rendering hint — never affects stored file content."
        integer height
        timestamp deleted_at
        timestamp uploaded_at
        timestamp updated_at
        timestamp content_updated_at "Content-scoped version: advances ONLY when the file's CONTENT changes (upload / content overwrite), never on metadata writes (rename, move, soft-delete, restore). It is the optimistic-concurrency validator the collaborative-document persist guards on (RFC 7232 `If-Match` semantics: validate the representation, not the row) — so a rename can't make a racing live-doc persist see a stale token, reconcile stale durable content, and clobber in-flight edits. NOT NULL with a `now()` default: Postgres applies this as a fast-default (no table rewrite), existing rows get a stable timestamp that — like every metadata write — never advances it, and every insert path is covered without per-call plumbing. Only a content write (upload / overwrite) advances it. MILLISECOND precision is an invariant, not an incidental detail. This value is a revision identity that round-trips through JavaScript `Date` and JSON — search-index dispatch payloads, If-Match tokens, realtime versions — all of which truncate to milliseconds, while `now()` stores microseconds. A sub-millisecond value therefore stops comparing equal to its own round-trip, and every SQL equality keyed on it matches zero rows: a file whose `workspace_file_search_index.source_content_updated_at` came from such a round trip can never be claimed, indexed, or cleaned up. The default truncates, and the `workspace_files_content_version_millisecond` trigger enforces it for the writers a default cannot reach — explicit `CURRENT_TIMESTAMP` expressions, raw SQL inserts, and updates that change row data. Relationship fences that only advance the row version preserve an existing revision."
        integer secret_provenance_version "Durable cutover marker for content secret provenance. NULL is reserved for legacy rows and writes from app versions that predate tracking. Provenance-aware writers set version 1 in the same transaction as the matching sidecar. A tracked version without a matching sidecar fails closed."
    }
    workspace_file_search_index {
        text file_id PK, FK
        text workspace_id FK
        timestamp source_content_updated_at PK
        workspace_file_search_index_status status
        boolean partial
        text failure_reason
        integer line_count
        integer indexed_bytes
        timestamp dispatched_at
        timestamp created_at
        timestamp updated_at
    }
    workspace_file_search_dispatch_queue {
        text workspace_id PK, FK
        timestamp enqueued_at
        timestamp last_dispatched_at
        timestamp updated_at
    }
    file_search_dispatch_queue {
        text entity_type PK
        text entity_id PK
        timestamp enqueued_at
        timestamp last_dispatched_at
        timestamp updated_at
    }
    workspace_file_search_backfill {
        text id PK
        text after_workspace_id
        text after_entity_type
        text after_entity_id
        text after_file_id
        timestamp completed_at
        timestamp updated_at
    }
    workspace_file_search_segment {
        text file_id PK, FK
        text workspace_id FK
        timestamp source_content_updated_at PK
        integer line_number PK
        integer segment_number PK
        integer segment_start
        integer line_length
        text content
    }
    workspace_file_search_build {
        text id PK
        text file_id
        text workspace_id
        text entity_type
        text entity_id
        text artifact_key
        timestamp source_content_updated_at
        timestamp created_at
        timestamp expires_at "Null only for a completely published build; abandoned workers expire automatically."
    }
    workspace_file_search_revision {
        text file_id PK, FK
        text workspace_id
        text entity_type
        text entity_id
        timestamp source_content_updated_at
        workspace_file_search_index_status status
        text build_id FK
        text failure_reason
        integer line_count
        integer indexed_bytes
        integer chunk_count
        timestamp dispatched_at
        timestamp handoff_expires_at "Deadline for a claim's run to be handed off, in PostgreSQL time. Set with the claim and cleared once a run is known to exist: Trigger.dev accepted it, in-process indexing took it, or it began its build. A claim still carrying an expired deadline has no run known to exist and is released, and its token fences out any run it did get. NULL otherwise, including claims made before this column existed, which fall back to the stale-dispatch window."
        timestamp updated_at
    }
    workspace_file_search_chunk {
        text build_id PK, FK
        text workspace_id
        text entity_type
        text entity_id
        integer ordinal PK
        integer line_start
        boolean fragment
        integer overlap
        text content
    }
    file_search_dependency {
        text build_id PK, FK
        text file_id PK
        text key
        timestamp source_content_updated_at
        text entity_type "The build's canonical owner, always derived from the build by the `file_search_dependency_owner` trigger, so the dispatcher can find an owner's stale dependencies without testing every ready revision."
        text entity_id
    }
    upload_session {
        text id PK
        text token_hash
        text user_id
        text workspace_id
        text knowledge_base_id
        text workflow_id
        text execution_id
        upload_session_purpose purpose
        upload_session_method method
        text storage_context
        text final_key
        upload_session_provider storage_provider
        text provider_upload_id
        text provider_object_version
        text file_name
        text content_type
        bigint file_size
        integer part_size
        integer part_count
        upload_session_status status
        jsonb metadata
        text processing_lease_id
        timestamp processing_lease_expires_at
        text completed_file_id
        text error
        timestamp created_at
        timestamp expires_at
        timestamp completed_at
        timestamp updated_at
    }
    workspace_file_secret_provenance {
        text file_id PK, FK
        timestamp content_updated_at
        text status
        jsonb entries
        timestamp updated_at
    }
    workspace_file_collab_state {
        text file_id PK, FK
        bytea doc_state "`Y.encodeStateAsUpdate` of the collaborative doc at last persist — apply with `Y.applyUpdate`."
        text source_hash "sha256 (hex) of the markdown `docState` was derived from — the freshness tag for cold-start."
        timestamp updated_at
    }
    workspace_file_version {
        text id PK
        text file_id FK
        text workspace_id FK "Compatibility scope, checked against fileId; Project versions derive their owner from the file."
        integer version
        text key
        bigint size_bytes
        text content_type
        text content_hash "sha256 (hex) of the bytes; NULL for an implicit version materialized without reading them."
        timestamp superseded_at
        workspace_file_version_source source
        text author_user_ids "Users who wrote these bytes, in first-contribution order; empty when unattributable."
        integer restored_from_version
        text secret_provenance_status
        jsonb secret_provenance_entries
        timestamp created_at
        timestamp updated_at
    }
    public_share {
        text id PK
        text resource_type "'file' | 'folder' (folder reserved for future)"
        text resource_id
        text entity_type
        text entity_id
        text workspace_id FK
        text created_by FK "SET NULL (not CASCADE) so a share — and its public link — outlives the user who created it; the shared file keeps its durable owner."
        text token
        boolean is_active
        text auth_type "'public' (anyone with the link) | 'password' | 'email' (OTP) | 'sso'."
        text password "AES-256-GCM encrypted share password; null unless authType is 'password'."
        json allowed_emails "Allowed emails/domains (e.g. '@acme.com') when authType is 'email' or 'sso'."
        timestamp created_at
        timestamp updated_at
    }
    invitation_workspace_grant {
        text id PK
        text invitation_id FK
        text workspace_id FK
        permission_type permission
        timestamp created_at
        timestamp updated_at
    }
    permissions {
        text id PK
        text user_id FK
        text entity_type "'workspace', 'workflow', 'organization', etc."
        text entity_id
        permission_type permission_type
        timestamp created_at
        timestamp updated_at
    }
    memory {
        text id PK
        text workspace_id FK
        text key
        jsonb data
        integer storage_version "Version 2 keeps data as an immutable prefix and appends ordered memory items."
        text encrypted_context_summary "One replaceable derived context summary; never part of the public message projection."
        integer secret_provenance_version "NULL is a legacy/untracked record; version 1 requires a fresh private sidecar."
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }
    memory_secret_provenance {
        text memory_id PK, FK
        text content_hash
        text status
        jsonb entries
        timestamp updated_at
    }
    memory_item {
        text id PK
        text memory_id FK
        bigint sequence
        text append_key
        text turn_id
        text kind
        jsonb data
        text content_hash
        text provenance_status
        jsonb provenance_entries
        timestamp created_at
    }
    agent_memory_turn {
        text id PK
        text memory_id FK
        text workflow_id FK
        text execution_id
        text block_id
        text node_id
        integer execution_order
        text encrypted_state
        integer revision
        timestamp created_at
        timestamp updated_at
    }
    memory_artifact {
        text memory_id PK, FK
        text key PK, FK
    }
    organization_search_integration {
        text organization_id PK, FK
        text connector_type PK
        boolean approved
        timestamp updated_at
    }
    knowledge_base {
        text id PK
        text user_id FK
        text workspace_id FK
        text organization_id FK
        text folder_id FK
        text name
        text description
        boolean is_search_index "Search indexes retain their identity independently of their display name."
        integer token_count
        text embedding_model
        integer embedding_dimension
        json chunking_config
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    knowledge_base_import {
        text id PK
        text workspace_id FK
        text user_id "Attribution only, with no foreign key: the row is the janitor's handle on the bundle and the base, so it must outlive its importer's account."
        text source_key "The uploaded archive, deleted once the import settles."
        text knowledge_base_id FK
        text status
        integer documents_total
        integer documents_imported
        integer documents_failed
        boolean vectors_reused "Whether the bundle's vectors matched this deployment's embedding model and were kept."
        jsonb failures "The first few documents that could not be imported, as `{ filename, message }`."
        text error
        timestamp created_at
        timestamp started_at "When a worker claimed the import; unset, a retried completion dispatches it again."
        timestamp updated_at
        timestamp completed_at
    }
    document {
        text id PK
        text knowledge_base_id FK
        text filename
        text file_url
        text storage_key "Canonical storage key derived from fileUrl at write time (e.g. 'kb/<...>'), or null for external/data: ingestion URLs. KB file authorization matches on this exact key rather than re-parsing the URL at read time."
        integer file_size "Size in bytes"
        text mime_type "e.g., 'application/pdf', 'text/plain'"
        integer chunk_count
        integer token_count
        integer character_count
        text processing_status "'pending', 'processing', 'completed', 'failed'"
        integer processing_attempts "Dispatches spent on this document since its last successful pass. A bounded retry budget, not a dispatch generation. The stuck-document sweep re-dispatches a failing document every sync for the whole retry window, and each dispatch re-parses and re-embeds it — so a document that fails deterministically was billed once per sync indefinitely. Past the budget it becomes a dead letter: still visible and still user-retryable, but no longer swept. Reset to 0 whenever a pass completes."
        timestamp processing_queued_at "When indexing was last dispatched to a worker, which is not when a worker picked it up — a document sits at `pending` in between. Recovery sweeps measure queue wait from here; `processingStartedAt` is written only once a worker actually starts. NULL means never dispatched, or dispatched before this column existed."
        text processing_queue_token "Opaque dispatch generation; NULL identifies payloads created before token rollout."
        timestamp processing_started_at
        timestamp processing_deferred_until "Scheduled execution time of an accepted durable quota continuation."
        timestamp processing_completed_at
        text processing_error
        timestamp processing_recovery_after "Retry admission backoff, separate from an accepted worker continuation."
        boolean enabled
        timestamp archived_at "Parent KB/workspace archive marker"
        timestamp deleted_at
        boolean user_excluded "User explicitly excluded — skip on sync"
        text tag1 "Document tags for filtering (inherited by all chunks) Text tags (7 slots)"
        text tag2
        text tag3
        text tag4
        text tag5
        text tag6
        text tag7
        double_precision number1 "Number tags (5 slots)"
        double_precision number2
        double_precision number3
        double_precision number4
        double_precision number5
        timestamp date1 "Date tags (2 slots)"
        timestamp date2
        boolean boolean1 "Boolean tags (3 slots)"
        boolean boolean2
        boolean boolean3
        text connector_id FK
        text external_id
        text content_hash
        text source_url
        integer secret_provenance_version "NULL is a legacy/untracked source; version 1 requires a matching source sidecar."
        text uploaded_by FK "User who uploaded the document, for usage attribution. Null for connector/cron-synced docs (and pre-migration rows) → indexing billing falls back to the workspace billed account."
        text acl "Sorted access-token list applied by every document read; the vocabulary is owned by `apps/sim/lib/knowledge/access/tokens.ts`. `{ws}` (the default) is any workspace member and `{}` is nobody. Uploads, API-created documents, and workspace-mode connectors keep the default; a members-mode connector materialises it from `knowledge_document_observation`."
        jsonb acl_requirements "Additional OR clauses, all of which must match; preserves source permission intersections."
        timestamp acl_verified_at "Last authoritative source ACL evidence; NULL fails closed for mirrored permissions."
        timestamp source_modified_at "Source last-modified time when the connector reports one; NULL for uploads."
        timestamp source_seen_at "Start of the durable source-listing cycle that last observed this document."
        timestamp uploaded_at
    }
    document_secret_provenance {
        text document_id PK, FK
        text source_hash
        text status
        jsonb entries
        timestamp updated_at
    }
    knowledge_base_tag_definitions {
        text id PK
        text knowledge_base_id FK
        text tag_slot
        text display_name
        text field_type "'text', future: 'date', 'number', 'range'"
        timestamp created_at
        timestamp updated_at
    }
    embedding {
        text id PK
        text knowledge_base_id FK
        text document_id FK
        integer chunk_index
        text chunk_hash
        text content
        integer secret_provenance_version "NULL is a legacy/untracked chunk; version 1 requires a fresh private sidecar."
        integer content_length
        integer token_count
        vector embedding "Vector embeddings. A chunk populates exactly the one column matching its knowledge base's `embedding_dimension`, and the others stay NULL: pgvector fixes a column's width, so one width per column is the only way to store models that emit different sizes in the same table. The widths cover what the popular embedding models emit — 384 (all-minilm), 768 (nomic-embed-text, embeddinggemma), 1024 (mxbai-embed-large, bge-m3, Voyage), 1536 (OpenAI's small model), 3072 (OpenAI's large model, gemini-embedding-001). `embedding` is the original 1536 column, kept under its bare name so every row written before the other widths existed stays exactly where it is."
        vector embedding_384
        vector embedding_768
        vector embedding_1024
        vector embedding_3072
        text embedding_model
        integer start_offset
        integer end_offset
        text tag1 "Tag columns inherited from document for efficient filtering Text tags (7 slots)"
        text tag2
        text tag3
        text tag4
        text tag5
        text tag6
        text tag7
        double_precision number1 "Number tags (5 slots)"
        double_precision number2
        double_precision number3
        double_precision number4
        double_precision number5
        timestamp date1 "Date tags (2 slots)"
        timestamp date2
        boolean boolean1 "Boolean tags (3 slots)"
        boolean boolean2
        boolean boolean3
        boolean enabled
        tsvector content_tsv
        timestamp created_at
        timestamp updated_at
    }
    embedding_search {
        text id PK, FK
        text knowledge_base_id
        text document_id
        boolean enabled
        text connector_id "contract-pending(after #8528 is fully deployed and source/ACL projection writers have drained): drop connector_id — regular KB retrieval checks the parent document."
        text acl "contract-pending(after #8528 is fully deployed and source/ACL projection writers have drained): drop acl — regular KB retrieval retains document-level access checks."
        bit binary "contract-pending(after vector writers stop computing binary projections and embedding_search_width_check is replaced): drop binary and all binary_* columns — their ANN indexes were dropped in 0372 and no reader uses them."
        bit binary_384
        bit binary_768
        bit binary_1024
        bit binary_3072
        halfvec vector
        halfvec vector_384
        halfvec vector_512
        halfvec vector_768
        halfvec vector_1024
        halfvec vector_3072
    }
    knowledge_projection_dirty {
        text document_id PK, FK
        bigint generation "Bumped by every mark; the projector removes the row only on the generation it read."
        boolean content "Whether chunk content changed, not only the document's source or ACL."
        timestamptz marked_at
    }
    embedding_secret_provenance {
        text embedding_id PK, FK
        text content_hash
        text status
        jsonb entries
        timestamp updated_at
    }
    docs_embeddings {
        uuid chunk_id PK
        text chunk_text
        text source_document
        text source_link
        text header_text
        integer header_level
        integer token_count
        vector embedding "Vector embedding - optimized for text-embedding-3-small with HNSW support"
        text embedding_model
        jsonb metadata
        tsvector chunk_text_tsv
        timestamp created_at
        timestamp updated_at
    }
    mothership_benchmarks {
        text id PK
        text organization_id FK
        text user_id FK
        text run_as_user_id FK "Null preserves the original owner's execution identity for existing cases."
        text source_workspace_id FK
        text name
        jsonb artifacts
        integer version
        text running_stage
        text attempt_id
        timestamp lease_expires_at
        text planner_chat_id
        text error
        timestamp created_at
        timestamp updated_at
    }
    mothership_benchmark_runs {
        text id PK
        text benchmark_id FK
        text label
        text evaluation_key
        jsonb reviews
        integer version
        timestamp reviewed_at
        integer correct
        integer automatic_correct
        integer total
        jsonb artifacts
        timestamp created_at
    }
    mothership_memory_spaces {
        uuid id PK
        text user_id FK
        text organization_id FK
        text name
        timestamp created_at
    }
    mothership_memory_selections {
        text user_id PK, FK
        text organization_id PK, FK
        uuid space_id FK
    }
    copilot_chats {
        uuid id PK
        text user_id FK
        text workflow_id FK
        text workspace_id FK
        text organization_id FK
        uuid memory_space_id "Bound once, without an FK: workspace chats can outlive their organization; missing graph ownership fails closed."
        chat_type type
        text title
        text model
        text conversation_id
        text external_conversation_key "Stable provider conversation identity; only trusted ingress can bind it to this private chat."
        jsonb external_conversation_metadata
        text preview_yaml
        jsonb config
        jsonb resources
        jsonb auto_allowed_tools "Copilot tool ids the user allowed for the rest of this chat only, as opposed to the account-wide list on `settings.copilotAutoAllowedTools`."
        timestamp last_seen_at
        boolean pinned
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    mothership_resource_effects {
        uuid chat_id PK, FK
        text effect_id PK
        timestamp created_at
    }
    copilot_messages {
        uuid id PK
        uuid chat_id FK
        text message_id
        text role
        jsonb content
        jsonb display_content "The chat UI's projection of `content` (tool-call params truncated), or null when it equals `content` or has not been backfilled. Transcript reads for the UI select `coalesce(display_content, content)` so they never detoast the full params."
        text stream_id
        text model
        integer seq
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    workflow_checkpoints {
        uuid id PK
        text user_id FK
        text workflow_id FK
        uuid chat_id FK
        text message_id "ID of the user message that triggered this checkpoint"
        json workflow_state
        timestamp created_at
        timestamp updated_at
    }
    copilot_request_stops {
        text user_id PK, FK
        text workspace_id PK, FK
        text stream_id PK
        timestamptz stopped_at
    }
    copilot_organization_request_stops {
        text user_id PK, FK
        text organization_id PK, FK
        text stream_id PK
        timestamptz stopped_at
    }
    desktop_devices {
        text id PK
        text user_id FK
        text session_id FK
        text name
        text app_version
        text platform
        jsonb capabilities
        timestamptz last_seen_at
        timestamptz revoked_at
        timestamptz created_at
        timestamptz updated_at
    }
    copilot_runs {
        uuid id PK
        text execution_id
        integer tool_execution_version "Rows predating the active ownership protocol cannot certify tool settlement."
        timestamp tool_admission_closed_at
        uuid parent_run_id
        uuid chat_id FK
        text user_id FK
        text workflow_id FK
        text workspace_id FK
        text stream_id
        text organization_id FK
        text agent
        text model
        text provider
        copilot_run_status status
        jsonb request_context
        text desktop_device_id FK "Set at admission when the turn's desktop runs its tools in the background executor."
        timestamp started_at
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
        text error
    }
    copilot_run_checkpoints {
        uuid id PK
        uuid run_id FK
        text pending_tool_call_id
        jsonb conversation_snapshot
        jsonb agent_state
        jsonb provider_request
        timestamp created_at
        timestamp updated_at
    }
    copilot_async_tool_calls {
        uuid id PK
        uuid run_id FK
        uuid checkpoint_id FK
        text tool_call_id
        text tool_name
        jsonb args
        text executable_name "Canonical dispatch identity; null only on calls recorded before invocation binding."
        text target_workspace_id
        jsonb approved_args "Validated user edits, committed once with the decision; original args stay immutable."
        jsonb assistant_action "Server-resolved authority sealed with the original proposal, never supplied by the worker."
        boolean execution_retryable "External writes without a provider idempotency key must never be replayed after a crash."
        copilot_async_tool_status status
        jsonb result
        text error
        copilot_tool_permission_decision permission_decision "A null decision on a gated call means its prompt is still outstanding after a reload."
        timestamp permission_decided_at
        timestamp permission_requested_at "Set when this call was held for the user's decision. Whether a call is gated depends on the turn (the feature flag, the tool's arguments, the user's allow lists), so the row records it: every execution claim refuses a gated call until `permissionDecision` allows it."
        timestamptz pickup_deadline_at "When a desktop call on a device-bound run must be claimed by: set once the call may run and is offered to its device. A call still unclaimed after it fails as never started."
        bigint persist_seq "When this call was persisted, as a strictly increasing number: calls from one turn can share a millisecond, and a device must run them in the order the model emitted them. Null on rows persisted before the column existed."
        timestamp claimed_at
        text claimed_by
        timestamp browser_download_started_at "One-use download-save admission; never released after an uncertain storage outcome."
        timestamp execution_started_at "Separate from the model-facing terminal result, which can precede cleanup."
        timestamp execution_settled_at
        text execution_owner_token "Independent of stream ownership and terminal result delivery."
        timestamptz execution_lease_expires_at
        timestamptz execution_revoked_at
        integer execution_attempts "Sim executions started for this call, counting each re-run after its owner died."
        timestamptz execution_deadline_at "When Sim must have settled this call: the worker's wait from the frame, less the settle margin, fixed at first receipt. Every wait, execution and re-run of the call ends by it. Null when the frame named no deadline."
        text client_workflow_execution_id "Assigned only after the workflow HTTP executor has reserved this execution identity."
        jsonb sandbox_processes
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }
    copilot_feedback {
        uuid feedback_id PK
        text user_id FK
        uuid chat_id FK
        text user_query
        text agent_response
        boolean is_positive
        text feedback
        text workflow_yaml "Optional workflow YAML if edit/build workflow was triggered"
        timestamp created_at
        timestamp updated_at
    }
    workflow_deployment_version {
        text id PK
        text workflow_id FK
        integer version
        text name
        text description
        json state
        boolean is_active
        timestamp created_at
        text created_by
    }
    workflow_deployment_operation {
        text id PK
        text workflow_id FK
        text deployment_version_id FK
        integer version
        text previous_active_version_id FK
        text action
        integer protocol_version
        integer generation
        text status
        jsonb component_readiness
        text error_code
        text error_message
        text idempotency_key
        text request_hash
        text actor_id
        timestamp completed_at
        timestamp created_at
        timestamp updated_at
    }
    idempotency_key {
        text key PK
        json result
        timestamp created_at
    }
    outbox_event {
        text id PK
        text event_type
        json payload
        text status
        integer attempts
        integer max_attempts
        timestamp available_at
        timestamp locked_at
        text last_error
        timestamp created_at
        timestamp processed_at
    }
    mcp_servers {
        text id PK
        text workspace_id FK
        text organization_id FK
        text credential_group_id FK
        text managed_connector_id
        integer oauth_config_version
        text created_by FK "Track who created the server, but workspace owns it"
        text name
        text description
        text transport
        text url
        text auth_type
        text oauth_client_id "Optional pre-registered OAuth credentials for servers that don't support Dynamic Client Registration (RFC 7591). When set, these shortcut the SDK's DCR step. `oauthClientSecret` is encrypted."
        text oauth_client_secret
        json headers
        integer timeout
        integer retries
        boolean enabled
        timestamp last_connected
        text connection_status
        text last_error
        jsonb status_config
        integer tool_count
        timestamp last_tools_refresh
        integer total_requests
        timestamp last_used
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    mcp_server_oauth {
        text id PK
        text mcp_server_id FK
        text user_id FK "Last workspace user who initiated/completed authorization."
        text workspace_id FK
        text organization_id FK
        text client_information "Encrypted JSON of the RFC 7591 dynamic client registration result. Encrypted because some authorization servers may issue a client_secret even for clients advertising `token_endpoint_auth_method: 'none'`."
        text tokens "Encrypted JSON of the OAuth tokens (access + refresh)."
        text code_verifier "PKCE verifier held only between /authorize and /callback."
        text state "Opaque state mint to correlate the callback."
        timestamp state_created_at "When `state` was minted. Used to expire the active-flow window and the state replay window independently of `updatedAt`, which is touched by token refreshes and other writes."
        timestamp last_refreshed_at
        timestamp created_at
        timestamp updated_at
    }
    sso_provider {
        text id PK
        text issuer
        text domain
        text oidc_config
        text saml_config
        text user_id FK
        text provider_id
        text organization_id FK
        boolean domain_verified "Better Auth's SSO `domainVerification` flag. Sim proves ownership itself via {@link ssoDomain} before registration, so this mirrors that decision rather than driving a second flow. It makes Better Auth treat the provider as authoritative for its domain and auto-link same-email accounts; without it, IdPs omitting `email_verified` (notably Entra) strand those users. Defaults to true so pre-existing providers keep signing in across deploy."
        boolean jit_provisioning_enabled "Whether a successful SSO sign-in may provision a new organization membership. Sim owns this admission path so seat checks, billing effects, session policy, and audit all use the same transaction as every other join. Defaults to true to preserve existing providers during a rolling deploy."
    }
    sso_domain {
        text id PK
        text organization_id FK
        text domain "Normalized (lowercase, registrable) domain — see `normalizeSSODomain`."
        text status "`'pending'` until the DNS TXT record is observed, then `'verified'`."
        text verification_token "High-entropy token placed in the domain's `_sim-challenge` TXT record."
        timestamp verified_at
        text primary_provider_id "The provider sign-in uses for this domain when the organization has more than one on it, such as while moving from one identity provider to another. Holds the provider id, not a foreign key: it is honored only while that provider still belongs to this organization and serves this domain, and deleting the provider clears it. Null means the domain's first verified provider by id, which is also the only one when there is just one. See `sso-primary-provider.ts`."
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    oauth_client {
        text id PK
        text client_id
        text client_secret
        boolean disabled
        boolean skip_consent
        boolean enable_end_session
        text subject_type
        text scopes
        text user_id FK
        timestamp created_at
        timestamp updated_at
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
        text token_endpoint_auth_method
        text grant_types
        text response_types
        boolean public
        text type
        boolean require_pkce
        text reference_id
        jsonb metadata
    }
    oauth_consent {
        text id PK
        text client_id FK
        text user_id FK
        text reference_id
        text scopes
        timestamp created_at
        timestamp updated_at
    }
    oauth_token_family {
        text id PK
        text client_id FK
        text session_id FK
        text user_id FK
        text reference_id
        text consent_id FK
        integer current_generation
        timestamp created_at
        timestamp expires_at
    }
    oauth_refresh_token {
        text id PK
        text token
        text client_id FK
        text session_id FK
        text user_id FK
        text reference_id
        timestamp expires_at
        timestamp created_at
        timestamp revoked
        timestamp auth_time
        text scopes
        text resource
        text family_id FK
        integer generation
    }
    oauth_access_token {
        text id PK
        text token
        text client_id FK
        text session_id FK
        text user_id FK
        text reference_id
        text refresh_id FK
        timestamp expires_at
        timestamp created_at
        text scopes
        text resource
    }
    workflow_mcp_server {
        text id PK
        text workspace_id FK
        text created_by FK
        text name
        text description
        boolean is_public
        timestamp deleted_at
        timestamp created_at
        timestamp updated_at
    }
    workflow_mcp_tool {
        text id PK
        text server_id FK
        text workflow_id FK
        text tool_name
        text tool_description
        json parameter_schema
        json parameter_description_overrides
        timestamp archived_at
        timestamp created_at
        timestamp updated_at
    }
    custom_block {
        text id PK
        text organization_id FK
        text workflow_id FK
        text type
        text name
        text description
        text icon_url "Uploaded icon image URL (workspace storage), or null for the default icon."
        json inputs "Per-input authored overrides keyed by the source Start field's stable `id`: `Array<{ id, placeholder?, required? }>`. Only the placeholder and required flag are authored — the input field set and its name/type/description are always derived live from the deployed Start (so they can never go stale); an override whose field was removed is ignored. Absent/empty → no overrides; every deployed Start input is still exposed."
        json outputs "Curated outputs exposed to consumers: `Array<{ blockId, path, name }>`. Each maps a child-workflow block output (blockId + dot-path) to a friendly output name on the block. Empty/absent → expose the child's whole `result`. Internal plumbing (child workflow id, trace spans) is never exposed."
        boolean enabled
        boolean trace_child_runs "The publisher's org-wide decision on whether this block's runs are joined into a consumer's trace. It is the ONLY policy — no viewer check runs downstream — so turning it on publishes the source workflow's block names, inputs, outputs, and prompts to anyone who can read a consuming workflow's logs, including consumers with no access to the source workspace. Defaults false because of that: this is the same boundary curated outputs and redacted errors exist to hold, and it may only open by an affirmative act of the publisher, never by a column default applied to rows nobody revisited."
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    audit_log {
        text id PK
        text workspace_id FK
        text actor_id FK
        text action
        text resource_type
        text resource_id
        text actor_name
        text actor_email
        text resource_name
        text description
        jsonb metadata
        text ip_address
        text user_agent
        text surface "The official client the request came from (`web`, `desktop`, `cli`, `sdk-js`, `sdk-python`), as resolved from `X-Sim-Client-Info`. Null for background work and callers that do not identify themselves."
        timestamp created_at
    }
    organization_search_mcp_invocation {
        text id PK
        text organization_id FK
        text user_id FK
        text auth_kind
        text oauth_client_id "Snapshots survive OAuth client deletion; names are client-declared, not verified branding."
        text client_name
        text tool_name
        text outcome
        integer duration_ms
        timestamptz created_at
    }
    usage_log {
        text id PK
        text user_id FK
        usage_log_category category
        usage_log_source source
        text description
        jsonb metadata
        decimal cost
        text event_key
        billing_entity_type billing_entity_type
        text billing_entity_id
        timestamp billing_period_start
        timestamp billing_period_end
        text workspace_id FK
        text workflow_id FK
        text execution_id
        timestamp created_at
    }
    credential {
        text id PK
        text workspace_id FK
        text organization_id FK
        text slack_app_id FK
        credential_type type
        text display_name
        text description
        boolean unredacted "Opts an env_workspace secret out of resolved-secret redaction: its value renders in plaintext across surfaces (logs, model-visible content, sandbox file exports) instead of `{{NAME}}`, and is not recorded into durable provenance. Meaningful only for type = 'env_workspace'; writes for other types are rejected in the orchestration layer."
        text provider_id
        text account_id FK
        text env_key
        text env_owner_user_id FK
        text encrypted_service_account_key
        text encrypted_personal_token "Encrypted provider token bound immutably to createdBy, providerSubjectId, and providerTenantId."
        text authorization_app_id
        text credential_group_enrollment_id FK
        text credential_group_option_id
        text mcp_server_id FK
        integer mcp_oauth_config_version
        integer managed_oauth_scope_version
        text provider_subject_id
        text provider_tenant_id
        managed_oauth_credential_status managed_oauth_status
        text granted_scopes
        jsonb provider_metadata
        text encrypted_oauth_token_set
        jsonb mcp_tools
        timestamp mcp_tools_refreshed_at
        timestamp granted_at
        timestamp revoked_at
        timestamp access_token_expires_at
        timestamp refresh_token_expires_at
        timestamp last_refreshed_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    workload_identity_pending_subject {
        text id PK
        text workspace_id FK
        text user_id FK "Who the subject was issued to; only they can attach it to a credential."
        text provider "The keyless provider family (`aws`, `databricks`, …) the subject was issued for."
        text subject
        timestamp created_at
        timestamp expires_at "Renewed whenever the subject is shown again; once it passes, the subject is no longer offered."
    }
    credential_group {
        text id PK
        text workspace_id FK "contract-pending(org connected accounts fully deployed and legacy Search migrated): remove workspace ownership."
        text organization_id FK
        text public_id
        text name
        text description
        jsonb options
        text encrypted_provider_configuration
        credential_group_status status
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    slack_app {
        text id PK
        text organization_id FK
        text kind
        text client_id "Custom app credentials; company app credentials come from the deployment environment."
        text encrypted_client_secret
        text encrypted_signing_secret
        text revision
        timestamp created_at
        timestamp updated_at
    }
    slack_search_installation {
        text id PK
        text organization_id FK
        text credential_id FK
        text app_id
        text slack_app_id FK
        text team_id
        text team_name
        text bot_user_id
        text enterprise_id
        boolean enabled
        text credential_version
        text revision
        text last_outcome
        timestamp last_event_at
        timestamp created_at
        timestamp updated_at
    }
    slack_search_turn {
        text id PK
        integer ordinal
        text installation_id FK
        text conversation_key
        text event_id
        jsonb payload
        text status
        text lease_id
        timestamp lease_expires_at
        text outcome
        timestamp created_at
        timestamp updated_at
    }
    credential_group_enrollment {
        text id PK
        text credential_group_id FK
        text email
        text user_id FK "Bound once after the invitee signs in with the verified invitation email."
        credential_group_enrollment_status status
        text invitation_token_hash
        timestamp invitation_expires_at
        timestamp invited_at
        timestamp sent_at
        timestamp completed_at
        timestamp revoked_at
        text last_delivery_error
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    assistant_tool_permission {
        text organization_id PK, FK
        text user_id PK, FK
        text credential_id PK, FK
        integer revision
        text read_policy
        text write_policy
        jsonb overrides
        timestamp updated_at
    }
    credential_member {
        text id PK
        text credential_id FK
        text user_id FK
        credential_member_role role
        credential_member_status status
        timestamp joined_at
        text invited_by FK
        timestamp created_at
        timestamp updated_at
    }
    pending_credential_draft {
        text id PK
        text user_id FK
        text workspace_id FK
        text organization_id FK
        text provider_id
        text display_name
        text description
        text credential_id FK
        text oauth_config
        timestamp expires_at
        timestamp created_at
    }
    permission_group {
        text id PK
        text organization_id FK
        text name
        text description
        jsonb config
        text created_by FK
        timestamp created_at
        timestamp updated_at
        boolean is_default
        text membership_mode "How an empty non-default group behaves. `inherit` (the default, and every pre-existing group) keeps the member invariant above: no member rows means the group governs every member of its workspaces. `explicit` means the group governs exactly its member rows and therefore governs nobody when empty. Directory-managed groups must be `explicit`. Under `inherit`, an identity provider removing the last member would silently widen the group from 'these three people' to 'everyone in these workspaces' — the opposite of what the administrator asked for."
    }
    permission_group_workspace {
        text id PK
        text permission_group_id FK
        text workspace_id FK
        text organization_id FK
        timestamp created_at
    }
    permission_group_member {
        text id PK
        text permission_group_id FK
        text organization_id FK
        text user_id FK
        text assigned_by FK
        timestamp assigned_at
    }
    resource_policy {
        text id PK
        text workspace_id FK
        text organization_id FK
        text resource_type
        text resource_id
        integer revision
        jsonb document
        text created_by FK
        text updated_by FK
        timestamp created_at
        timestamp updated_at
    }
    async_jobs {
        text id PK
        text type
        jsonb payload
        text status
        timestamp created_at
        timestamp started_at
        timestamp completed_at
        timestamp run_at
        integer attempts
        integer max_attempts
        text error
        jsonb output
        jsonb metadata
        timestamp updated_at
    }
    knowledge_connector {
        text id PK
        text knowledge_base_id FK
        text connector_type
        text credential_id "The credential used to index content. In members mode it is optional: NULL indexes the union of members' listings; a dedicated credential indexes content while members only establish document visibility. Not yet a foreign key: rows written before the `credential` table existed may still hold a raw `account.id`, which script migration 0011 remaps. contract-pending(after 0011 has run in production): add the reference to `credential.id` with ON DELETE SET NULL and remove the raw account-id fallback in lib/oauth/credential-service.ts."
        text encrypted_api_key
        json source_config
        text sync_mode
        integer sync_interval_minutes
        text access_mode "How document access is derived. `workspace`: every synced document is `{ws}`. `members`: a document's ACL is the members whose own listing returned it, optionally with a dedicated credential for content. `admin`: source permissions and identity groups are mirrored independently of content changes."
        text credential_group_id FK "Members mode: the credential group whose option supplies the member credentials."
        text credential_group_option_id "Members mode: the option within the group; must map to this connector's provider."
        text member_sync_status "Members-mode run state. Mirrors `status` for the content engine but is independent of it: the two engines never run for the same connector (`kc_sync_lock_exclusive_check`), yet share no columns so neither can misread the other's lease."
        text member_sync_lock_token
        timestamp member_sync_lock_lease_at
        timestamp next_member_sync_at "Millisecond precision: the scheduler round-trips this through a JavaScript `Date` and claims the run by equality, so stored microseconds from a SQL writer could never be matched back. Any column compared that way has to stay within what a `Date` can carry."
        timestamp last_member_sync_at
        text last_member_sync_error
        integer member_sync_consecutive_failures
        boolean access_rewrite_pending "Set by a mode switch whose ACL rewrite exceeded the request budget; the member-sync job finishes the rewrite before the mode takes effect."
        boolean image_removal_pending "Set when a settings change stops indexing images the connector already indexed; the next sync removes them and clears it."
        jsonb member_tombstone_cursor "Where the members-mode absence reconcile resumes: the external id of the last live document it checked, in `doc_connector_external_id_idx` order. NULL starts a new pass from the beginning."
        text member_resurrection_cursor "Where the members-mode resurrection walk resumes: the last document id it covered. NULL starts a new walk from the connector's first document."
        text status "One of `active`, `pending`, `syncing`, `error`, `paused`, `disabled`. `pending` and `syncing` are the two halves of a sync in flight: `pending` is written as the sync is handed to the queue, `syncing` when a worker takes the lock. The split exists because the queue depth between them is unbounded — without `pending` the row is indistinguishable from idle for as long as the hand-off takes, which is what forced readers to guess from `created_at`. A row left `pending` past the lock TTL is reclaimed by the scheduler, which is the only thing that ever observes a lost hand-off."
        timestamp last_sync_at
        text last_sync_error
        integer last_sync_doc_count
        jsonb listing_checkpoint "Durable content-listing progress; contains cursors and counters, never credentials."
        jsonb directory_checkpoint "Member account enumeration resumes independently of the content listing."
        timestamp next_sync_at "Millisecond precision for the same round-trip reason as `next_member_sync_at`."
        timestamp next_directory_sync_at
        integer consecutive_failures
        text sync_lock_token "Identifies the sync run that currently holds this connector's lock. `status = 'syncing'` only says *a* run holds it. After the scheduler reclaims a stale lock and dispatches a replacement, the original run would still see `syncing` and overwrite the replacement's state. Terminal writes match this token so a run can prove the lock is still *its own*."
        timestamp sync_lock_lease_at "When the run holding this connector's lock last proved it was alive. Split off `updated_at`, which the stale-lock reaper used to read as a lease. `updated_at` is the row's modification time, so every unrelated write — a config edit, a status change — renewed the lease of a wedged run and pushed its recovery out by another full TTL. Only lock acquisition and the heartbeat write this column; both terminal helpers clear it alongside `sync_lock_token`. NULL on a row locked before this column existed, and on any future writer that forgets it, so every reader compares `COALESCE(lease, updated_at)` rather than the lease alone — a `lease <= cutoff` test is NULL-false and would make such a row permanently unreclaimable."
        timestamp created_at
        timestamp updated_at
        timestamp archived_at
        timestamp deleted_at
        timestamp detached_at "Set when the connector is removed but its documents are kept. The connector stops syncing and leaves every management surface at once, while its documents stay readable; a background job releases them as standalone entries in bounded pages and then deletes the row. Releasing a document rewrites every search projection row of it, so the release cannot run inside the removal request."
        bigint detach_reserved_bytes "Storage admitted and charged when the connector was detached but not yet matched by a released document. Each released page consumes its bytes; whatever remains when the row is deleted, such as a document deleted before its release, is settled then. Billing recomputations count it alongside standalone documents, since the workspace ledger already includes it."
    }
    knowledge_connector_partition {
        text connector_id PK, FK
        text partition_key PK
        text generation_id
        jsonb context
        text cursor
        text status
        integer attempts
        timestamp retry_at
        timestamp last_served_at
        jsonb failure
        text permission_cursor
        integer permission_attempts
        timestamp permission_retry_at
        timestamp permission_last_served_at
        timestamp permission_started_at
        jsonb permission_failure
    }
    knowledge_connector_permission_snapshot {
        text connector_id PK, FK
        integer revision
        jsonb metadata
        jsonb payload
    }
    knowledge_connector_permission_grant {
        text connector_id PK, FK
        text group_key PK
        text subject_token PK
    }
    knowledge_connector_member {
        text id PK
        text workspace_id FK
        text organization_id FK
        text connector_id FK
        text credential_id FK
        text subject_token "Snapshot of the member's identity token, derived from the credential row by `lib/knowledge/access/tokens.ts`. Reconciliation rewrites it and rematerialises the member's documents if the credential's subject changes."
        text status "`active`, `suspended`, or `disabled`."
        integer consecutive_failures
        timestamp next_attempt_at "When the member is next due. NULL means 'with the connector's next run': a member that completed on a manual-only connector. A new member is due now. An explicit time that has passed is what keeps a connector re-dispatching itself."
        timestamp last_started_at
        timestamp last_complete_listing_at "Last listing that was full, complete, and not suspect — the only kind that may remove observations."
        integer last_listed_count
        text last_error
        timestamp member_synced_through "Authorization watermark: complete nonsuspect full listing or completely drained change feed."
        timestamp scope_renewed_at "When every observation under the containers the source still grants this member was last renewed, for connectors that grant access per container; NULL until the first renewal completes."
        text scope_renewal_cursor "Where an unfinished scope renewal resumes in the source's container listing, and when that renewal pass began; both NULL when no pass is in progress."
        timestamp scope_renewal_started_at
        text change_cursor "Where the member's change feed resumes. Opened just before a full listing and stored once that listing lands, so every later run reads the feed instead of relisting; NULL when the connector has no feed or the feed has to be reopened."
        jsonb listing_checkpoint "Durable full-listing progress, separate from the provider's incremental change token."
        timestamp suspended_at
        timestamp created_at
        timestamp updated_at
    }
    knowledge_document_observation {
        text document_id PK, FK
        text member_id PK, FK
        timestamp last_seen_at
        text run_id "Member-sync run (`knowledge_connector_member_sync_log.id`) that last asserted this observation."
    }
    knowledge_external_directory {
        text workspace_id FK
        text organization_id FK
        text provider_id
        text tenant_id
        text sync_lock_token
        timestamp sync_lock_lease_at
        timestamp last_started_at "A newer attempt invalidates cached completion until that whole pass succeeds."
        timestamp last_complete_sync_at "Complete directory enumeration, including empty directories; independent of individual group freshness."
    }
    knowledge_external_group {
        text id PK
        text workspace_id FK
        text organization_id FK
        text provider_id "Matches the provider segment of the `g:` token, e.g. `google-drive`."
        text tenant_id "The directory this group belongs to: a Workspace domain for Google, a site's cloud id for Confluence."
        text external_group_id
        timestamp last_synced_at "When this group's membership was last enumerated in full, and the only thing that decides whether it still grants access. A failed or partial enumeration writes nothing at all — not the membership, not this column — which is what makes a transient directory outage harmless. It is also why the column has to exist: without an age bound, a group whose sync stopped running would keep granting forever from membership nobody has checked since. A group unconfirmed for longer than `EXTERNAL_GROUP_STALE_AFTER_MS` grants nothing."
        timestamp created_at
        timestamp updated_at
    }
    knowledge_external_group_member {
        text group_id PK, FK
        text subject_token PK "Includes `u:*@<domain>` for directory-wide grants; source account IDs retain their case."
        timestamp created_at
    }
    knowledge_connector_member_sync_log {
        text id PK
        text connector_id FK
        text status "`started`, `partial`, `completed`, or `failed`."
        timestamp started_at
        timestamp completed_at
        integer members_claimed
        integer members_completed
        integer members_incomplete
        integer members_failed
        integer docs_failed "Null on historical runs that did not record document failure counts."
        integer processing_dispatch_failed
        integer docs_listed
        integer docs_added
        integer docs_updated
        integer docs_unchanged
        integer docs_hydrated_once
        integer observations_added
        integer observations_renewed "Observations kept fresh by per-container renewal rather than relisting."
        integer observations_removed
        integer docs_tombstoned
        integer docs_resurrected
        integer docs_purged
        integer credentials_audited
        text error_message
        text database_failure_class "The transient database failure class (`capacity`, `conflict`, or `connection`) that failed the run; null on every other outcome and on runs logged before it was recorded. Only these runs count toward the database retry streak."
    }
    knowledge_connector_sync_log {
        text id PK
        text connector_id FK
        text status
        timestamp started_at
        timestamp completed_at
        integer docs_added
        integer docs_updated
        integer docs_deleted
        integer docs_unchanged
        integer docs_skipped
        integer docs_failed
        integer listed_count "Complete listing-cycle size; per-worker counters may cover only its last page batch."
        text error_message
        text database_failure_class "The transient database failure class (`capacity`, `conflict`, or `connection`) that failed the run; null on every other outcome and on runs logged before it was recorded. Only these runs count toward the database retry streak."
    }
    user_table_definitions {
        text id PK
        text workspace_id FK
        text folder_id FK
        text name
        text description
        jsonb schema "Stores the table schema definition. Example: { columns: [{ name: string, type: string, required: boolean }] }"
        jsonb metadata "Stores UI-specific metadata separate from the data schema. Example: { columnWidths: { name: 200, age: 100 } }"
        integer max_rows
        integer row_count
        bigint rows_version "Folded part of a monotonic counter: the triggers on `user_table_rows` log one `user_table_row_changes` row per INSERT/DELETE statement, and one per transaction when an UPDATE changes `data` or `order_key`; the fold adds their count here. Keys the versioned table-snapshot cache so a stored CSV under `v{rows_version}` is reused until the table mutates. Never written from application code — the fold is the only writer. Read the live value through `lib/table/row-changes.ts`, never this column alone."
        boolean schema_locked "Per-table mutation locks. Each guards one mutation verb; an admin toggles them independently. Enforced at the `lib/table` service layer (see `lib/table/mutation-locks.ts`), which covers every entry point — routes, workflow blocks, and Mothership — since all funnel through those helpers. A locked verb rejects with 423; toggling a lock requires workspace admin. Append-only = update + delete locked; read-only = all four locked. These are integrity controls (against accidental/agentic mutation), not confidentiality controls — reads and exports are never blocked."
        boolean insert_locked
        boolean update_locked
        boolean delete_locked
        timestamp archived_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    user_table_rows {
        text id PK
        text table_id FK "The foreign key is `DEFERRABLE INITIALLY DEFERRED` (the `table_rows_version_at_commit` migration), so a row updated twice in one transaction does not key-share the definition row mid-transaction. drizzle can't express deferrability, so it lives only in the migration."
        text workspace_id FK
        jsonb data
        integer position
        text order_key "Fractional order key (base-62 string) — the authoritative row order. Nullable during the backfill window. Ordered with `id` as a deterministic tiebreaker. Stored with `COLLATE 'C'` (migration 0228) so Postgres compares it bytewise, matching the fractional-indexing library's ASCII ordering. drizzle can't express column collation, so the collation lives only in the migration."
        integer secret_provenance_version
        timestamp created_at
        timestamp updated_at
        text created_by FK
    }
    user_table_row_changes {
        bigint id PK
        text table_id FK
        integer row_delta "`+n` for an insert, `-n` for a delete, `0` for an update of row content or order."
    }
    user_table_row_secret_provenance {
        text row_id PK, FK
        timestamp content_updated_at
        text status
        jsonb entries
        timestamp updated_at
    }
    table_views {
        text id PK
        text table_id FK
        text workspace_id FK
        text name
        jsonb config "`TableViewConfig` — `{ filter, sort, hiddenColumns, columnOrder, columnWidths, pinnedColumns }`. Every column reference is keyed by stable column id, so a rename never touches a saved view; ids of deleted columns are pruned on read."
        boolean is_default
        text created_by FK "Nullable with `set null` rather than the `cascade` used for table ownership: a view is workspace-shared, so it must outlive the member who created it."
        timestamp created_at
        timestamp updated_at
    }
    table_jobs {
        text id PK
        text table_id FK
        text workspace_id FK
        text type "`'import'` | `'delete'`."
        text status "`'running'` → `'ready'` | `'failed'` | `'canceled'`."
        jsonb payload "Type-specific descriptor (e.g. delete filter/exclusions). Nullable; reserved for future resumability — today's workers carry their payload in-process via `runDetached`."
        integer rows_processed
        text error
        timestamp started_at
        timestamp updated_at
        timestamp completed_at
    }
    table_row_executions {
        text table_id FK
        text row_id PK, FK
        text group_id PK
        text status
        text execution_id
        text job_id
        text workflow_id
        text error
        text running_block_ids
        jsonb block_errors
        timestamp cancelled_at
        text capability_governed_user_id FK "Person whose permission group gates this cell's tools, persisted with the dispatcher's `pending` pre-stamp. The stamp and the worker that runs it are not the same run: a cell task that finds the row's cascade lock held bails, and the lock owner drains the marker itself. That owner belongs to whatever dispatch queued IT, so without the subject on the marker the drained cell would run under the wrong person's group — or under none, when the owner is an actorless auto-fire. Read only while the marker is unclaimed (`pending` with a null `execution_id`); later writes on the same cell carry no subject and null it, which is why nothing reads it after pickup. `ON DELETE SET NULL`, matching `table_run_dispatches`: a deleted person's runs are stopped by that table's cancel, not held open by this reference."
        jsonb enrichment_details "Enrichment cascade breakdown (provider outcomes, cost, timing) for `enrichment`-type groups. Null for workflow groups and pre-feature runs. Deliberately excluded from the hot grid read (`loadExecutionsByRow`) — read on demand for the enrichment details panel."
        timestamp updated_at
    }
    table_run_dispatches {
        text id PK
        text table_id FK
        text workspace_id FK
        text request_id
        text mode "`'all'` re-runs completed cells; `'incomplete'` skips them."
        jsonb scope "`{ groupIds: string[], rowIds?: string[] }` — the run's scope."
        text status "`pending` → `dispatching` → `complete` | `cancelled`."
        integer cursor "Highest `user_table_rows.position` we've already enqueued cells for."
        jsonb limit "Optional cap on how much work the dispatch does before completing. `{ type: 'rows', max: number }` today; the discriminated shape lets future caps (cells, cost, duration) extend without a schema change. Null = unbounded (process every row in scope)."
        integer processed_count "Units of `limit.type` already consumed (eligible rows dispatched, for `type: 'rows'`). Mutable counter the dispatcher advances per window so the budget survives across the checkpointed waits between windows."
        boolean is_manual_run "When true, eligibility bypasses `autoRun: false` skip and treats terminal states as re-runnable. Auto-fire paths (row inserts, CSV import, addWorkflowGroup) set this to false so the dispatch honors the autoRun toggle."
        text triggered_by_user_id FK "User who triggered the run, for per-member usage attribution. Null for auto-fire (row insert/update, CSV import) with no human initiator — those fall back to the workspace billed account."
        text capability_governed_user_id FK "The person whose permission group governs what this run's cells may do. Distinct from `triggered_by_user_id`, which is an *attribution* and substitutes the workspace billed account when the credential names no human — right for a meter, wrong for a gate, since it would run a bystander's tool denylist against an actorless request. Null when the run has no acting person (workspace API key, schedule, auto-fire), which means no per-tool gate applies. Producers set it explicitly, `null` included: it is required on every dispatch input precisely so a new one cannot inherit the attribution by omission. `set null` on delete, paired with a cancel of this account's non-terminal dispatches inside `deleteUserAccount`. Nulling alone would be a silent un-gate — the worker cannot tell a subject erased by deletion from one that was never there — and `restrict` would block account deletion behind background work. Going terminal first makes the nulled row unreachable."
        timestamp requested_at
        timestamp heartbeat_at "Last time the dispatcher loop made progress on this dispatch. Stamped by the same per-window writes that advance `cursor` and `processed_count`, so it means 'a holder is alive' rather than 'this started a while ago' — the distinction the cleanup sweep needs, since `requested_at` never moves and a legitimately long dispatch would otherwise be reclaimed under a live holder. Null on rows written before this column existed; the sweep reads `COALESCE(heartbeat_at, requested_at)` so those stay reclaimable."
        timestamp completed_at
        timestamp cancelled_at
    }
    mothership_inbox_allowed_sender {
        text id PK
        text workspace_id FK
        text email
        text label
        text added_by FK
        timestamp created_at
    }
    mothership_inbox_task {
        text id PK
        text workspace_id FK
        text from_email
        boolean sender_authenticated "Whether the receiving provider authenticated `fromEmail`'s domain (SPF, DKIM, or DMARC)."
        text from_name
        text subject
        text body_preview
        text body_text
        text body_html
        text email_message_id
        text in_reply_to
        text response_message_id
        text agentmail_message_id
        text status
        uuid chat_id FK
        text trigger_job_id
        text result_summary
        text error_message
        text rejection_reason
        boolean has_attachments
        text cc_recipients
        timestamp created_at
        timestamp processing_started_at
        timestamp completed_at
    }
    mothership_inbox_webhook {
        text id PK
        text workspace_id FK
        text webhook_id
        text secret
        timestamp created_at
    }
    data_drains {
        text id PK
        text organization_id FK
        text name
        data_drain_source source
        data_drain_destination destination_type
        jsonb destination_config "Non-secret destination config (bucket, region, prefix, url, ...). Validated by destination registry."
        text destination_credentials "Encrypted JSON blob containing destination credentials. Never returned to clients."
        data_drain_cadence schedule_cadence
        boolean enabled
        text cursor "Opaque cursor — JSON-encoded, source-defined. Advances after each acknowledged delivery checkpoint."
        timestamp last_run_at
        timestamp last_success_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    data_drain_runs {
        text id PK
        text drain_id FK
        data_drain_run_status status
        data_drain_run_trigger trigger
        timestamp started_at
        timestamp finished_at
        integer rows_exported
        bigint bytes_written
        text cursor_before
        text cursor_after
        text error
        jsonb locators "Destination-specific delivery locators for this run (e.g. S3 keys, webhook response ids)."
    }
    workspace_sandbox {
        text id PK
        text workspace_id FK
        text name
        sandbox_language language
        jsonb dependencies
        jsonb cli_tools
        jsonb system_packages
        text spec_hash
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    sandbox_image {
        text id PK
        text provider
        text spec_hash
        jsonb spec
        sandbox_image_status status
        text image_ref "Passed to the provider at create time once `status` is `ready`."
        text provider_image_id "Provider-side image identifier, when it differs from `imageRef`."
        text build_id
        bigint materialization_generation "Monotonic target release for this provider materialization; legacy rows are generation 0."
        text error_code "Classified taxonomy code; see lib/execution/remote-sandbox/build-errors.ts."
        text error_message "User-facing copy rendered from the code at classification time."
        text error_detail "Installer log tail, shown behind a disclosure."
        timestamp last_used_at
        timestamp created_at
        timestamp updated_at
    }
    scim_connection {
        text id PK
        text organization_id FK
        text status "`active` or `disabled`. Disabling refuses every credential immediately."
        jsonb settings
        timestamp last_request_at
        text reconcile_lock_token "Reconcile-job lease, in the shape the connector member sync already uses."
        timestamp reconcile_lease_at
        timestamp reconciled_at
        text created_by FK
        timestamp created_at
        timestamp updated_at
    }
    scim_credential {
        text id PK
        text connection_id FK
        text token_hash
        text token_prefix "Leading characters of the token, for identifying it in the settings list."
        jsonb scopes
        timestamp expires_at
        timestamp revoked_at
        text revoked_by FK
        timestamp last_used_at
        text created_by FK
        timestamp created_at
    }
    scim_user {
        text id PK
        text connection_id FK
        text user_id FK
        text external_id
        text user_name "Lower-cased `userName`; the uniqueness and lookup key within a connection."
        boolean active
        jsonb attributes
        text order_key "Stable ascending sort key for pagination. A provider pages with `startIndex`, so the order must not shift between pages the way `created_at` alone can when rows share a timestamp."
        timestamp created_at
        timestamp updated_at
    }
    scim_user_tombstone {
        text id PK
        text connection_id FK
        text external_id
        text user_id FK
        timestamp deleted_at
    }
    scim_group {
        text id PK
        text connection_id FK
        text external_id
        text display_name
        text display_name_key "Lower-cased `displayName`. Uniqueness is case-insensitive because Microsoft Entra treats a group name as its match key and will otherwise create a duplicate whose only difference is capitalization."
        text order_key
        timestamp created_at
        timestamp updated_at
    }
    scim_group_member {
        text id PK
        text group_id FK
        text scim_user_id FK
        timestamp created_at
    }
    scim_group_mapping {
        text id PK
        text group_id FK
        text target_kind "`permission_group`, `workspace`, or `org_role`."
        text permission_group_id FK
        text workspace_id FK
        permission_type permission_type "Permission granted on `workspaceId`, for workspace targets."
        text role "Organization role granted, for `org_role` targets. Only `admin` is accepted."
        text source "`automatic` mappings were made by name matching and are replaced when the group is renamed; `manual` ones were made by an administrator and are never removed by a sync. Kept apart from `createdBy`, which goes null when its author's account is deleted."
        text created_by FK
        timestamp created_at
    }
    scim_projection_grant {
        text id PK
        text connection_id FK
        text scim_user_id FK
        text target_kind
        text target_id "Permission group id, workspace id, or the granted organization role."
        permission_type permission_type "The permission SCIM set, so a later manual upgrade stays detectable."
        permission_type baseline_permission "Manual workspace access to restore when an unlocked directory withdraws its grant."
        text origin "`directory` when the directory created the access; `adopted` when the person already held it by hand and a mapping merely covers it. Adopted access is left in place when the mapping goes away, unless the directory is the organization's source of truth."
        timestamp created_at
        timestamp updated_at
    }
    scim_request_log {
        text id PK
        text connection_id FK
        text credential_id
        text method
        text path "Resource path only. Query strings can carry directory attribute values."
        integer status
        text scim_type
        text detail
        text user_agent
        integer duration_ms
        timestamp created_at
    }
    copilot_service_usage {
        uuid id PK
        uuid stream_id
        text tool_call_id
        text service
        decimal cost_usd
        text worker_origin
        timestamptz created_at
        timestamptz next_attempt_at
        integer attempts
        timestamptz delivered_at
        text last_error
    }
```
