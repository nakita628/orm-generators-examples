```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    users ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    users |o--o{ token_blacklist_outstandingtoken : "(id) - (user_id)"
    token_blacklist_outstandingtoken ||--o| token_blacklist_blacklistedtoken : "(id) - (token_id)"
    users }o--o{ auth_group : "(groups) - (user_set)"
    users }o--o{ auth_permission : "(user_permissions) - (user_set)"
    users |o--o{ address : "(id) - (created_by_id)"
    users |o--o{ address : "(id) - (updated_by_id)"
    organization ||--o{ address : "(id) - (org_id)"
    users |o--o{ organization : "(id) - (created_by_id)"
    users |o--o{ organization : "(id) - (updated_by_id)"
    users |o--o{ tags : "(id) - (created_by_id)"
    users |o--o{ tags : "(id) - (updated_by_id)"
    organization ||--o{ tags : "(id) - (org_id)"
    users |o--o{ profile : "(id) - (created_by_id)"
    users |o--o{ profile : "(id) - (updated_by_id)"
    users ||--o{ profile : "(id) - (user_id)"
    organization ||--o{ profile : "(id) - (org_id)"
    address |o--o{ profile : "(id) - (address_id)"
    users |o--o{ comment : "(id) - (created_by_id)"
    users |o--o{ comment : "(id) - (updated_by_id)"
    django_content_type ||--o{ comment : "(id) - (content_type_id)"
    profile |o--o{ comment : "(id) - (commented_by_id)"
    contacts |o--o{ comment : "(id) - (commented_by_contact_id)"
    organization ||--o{ comment : "(id) - (org_id)"
    users |o--o{ commentFiles : "(id) - (created_by_id)"
    users |o--o{ commentFiles : "(id) - (updated_by_id)"
    comment ||--o{ commentFiles : "(id) - (comment_id)"
    organization |o--o{ commentFiles : "(id) - (org_id)"
    users |o--o{ attachments : "(id) - (created_by_id)"
    users |o--o{ attachments : "(id) - (updated_by_id)"
    django_content_type ||--o{ attachments : "(id) - (content_type_id)"
    organization ||--o{ attachments : "(id) - (org_id)"
    users |o--o{ document : "(id) - (created_by_id)"
    users |o--o{ document : "(id) - (updated_by_id)"
    organization ||--o{ document : "(id) - (org_id)"
    document }o--o{ profile : "(shared_to) - (document_shared_to)"
    document }o--o{ teams : "(teams) - (document_teams)"
    users |o--o{ apiSettings : "(id) - (created_by_id)"
    users |o--o{ apiSettings : "(id) - (updated_by_id)"
    organization ||--o{ apiSettings : "(id) - (org_id)"
    apiSettings }o--o{ profile : "(lead_assigned_to) - (lead_assignee_users)"
    apiSettings }o--o{ tags : "(tags) - (apisettings_set)"
    organization ||--o{ portal_login_token : "(id) - (org_id)"
    contacts ||--o{ portal_login_token : "(id) - (contact_id)"
    users |o--o{ activity : "(id) - (created_by_id)"
    users |o--o{ activity : "(id) - (updated_by_id)"
    profile |o--o{ activity : "(id) - (user_id)"
    organization ||--o{ activity : "(id) - (org_id)"
    users |o--o{ notification : "(id) - (created_by_id)"
    users |o--o{ notification : "(id) - (updated_by_id)"
    profile ||--o{ notification : "(id) - (recipient_id)"
    profile |o--o{ notification : "(id) - (actor_id)"
    organization ||--o{ notification : "(id) - (org_id)"
    users |o--o{ device_token : "(id) - (created_by_id)"
    users |o--o{ device_token : "(id) - (updated_by_id)"
    organization ||--o{ device_token : "(id) - (org_id)"
    users ||--o{ device_token : "(id) - (user_id)"
    users |o--o{ teams : "(id) - (created_by_id)"
    users |o--o{ teams : "(id) - (updated_by_id)"
    organization ||--o{ teams : "(id) - (org_id)"
    teams }o--o{ profile : "(users) - (user_teams)"
    users |o--o{ contact_form_submission : "(id) - (created_by_id)"
    users |o--o{ contact_form_submission : "(id) - (updated_by_id)"
    users |o--o{ contact_form_submission : "(id) - (replied_by_id)"
    users |o--o{ custom_field_definition : "(id) - (created_by_id)"
    users |o--o{ custom_field_definition : "(id) - (updated_by_id)"
    organization ||--o{ custom_field_definition : "(id) - (org_id)"
    users |o--o{ personal_access_token : "(id) - (created_by_id)"
    users |o--o{ personal_access_token : "(id) - (updated_by_id)"
    organization ||--o{ personal_access_token : "(id) - (org_id)"
    profile ||--o{ personal_access_token : "(id) - (profile_id)"
    users |o--o{ calendar_feed_token : "(id) - (created_by_id)"
    users |o--o{ calendar_feed_token : "(id) - (updated_by_id)"
    organization ||--o{ calendar_feed_token : "(id) - (org_id)"
    profile ||--o| calendar_feed_token : "(id) - (profile_id)"
    organization ||--o{ portal_access_token : "(id) - (org_id)"
    users |o--o{ security_audit_log : "(id) - (created_by_id)"
    users |o--o{ security_audit_log : "(id) - (updated_by_id)"
    users |o--o{ security_audit_log : "(id) - (user_id)"
    organization |o--o{ security_audit_log : "(id) - (org_id)"
    users |o--o{ pack_application : "(id) - (created_by_id)"
    users |o--o{ pack_application : "(id) - (updated_by_id)"
    organization ||--o{ pack_application : "(id) - (org_id)"
    profile |o--o{ pack_application : "(id) - (applied_by_id)"
    users |o--o{ installed_app : "(id) - (created_by_id)"
    users |o--o{ installed_app : "(id) - (updated_by_id)"
    organization ||--o{ installed_app : "(id) - (org_id)"
    profile |o--o{ installed_app : "(id) - (installed_by_id)"
    users |o--o{ saved_view : "(id) - (created_by_id)"
    users |o--o{ saved_view : "(id) - (updated_by_id)"
    organization ||--o{ saved_view : "(id) - (org_id)"
    profile ||--o{ saved_view : "(id) - (profile_id)"
    users |o--o{ accounts : "(id) - (created_by_id)"
    users |o--o{ accounts : "(id) - (updated_by_id)"
    organization ||--o{ accounts : "(id) - (org_id)"
    accounts }o--o{ profile : "(assigned_to) - (account_assigned_users)"
    accounts }o--o{ teams : "(teams) - (account_teams)"
    accounts }o--o{ contacts : "(contacts) - (account_contacts)"
    accounts }o--o{ tags : "(tags) - (account_tags)"
    users |o--o{ account_email : "(id) - (created_by_id)"
    users |o--o{ account_email : "(id) - (updated_by_id)"
    accounts |o--o{ account_email : "(id) - (from_account_id)"
    organization ||--o{ account_email : "(id) - (org_id)"
    account_email }o--o{ contacts : "(recipients) - (recieved_email)"
    users |o--o{ emailLogs : "(id) - (created_by_id)"
    users |o--o{ emailLogs : "(id) - (updated_by_id)"
    account_email |o--o{ emailLogs : "(id) - (email_id)"
    contacts |o--o{ emailLogs : "(id) - (contact_id)"
    organization ||--o{ emailLogs : "(id) - (org_id)"
    users |o--o{ case : "(id) - (created_by_id)"
    users |o--o{ case : "(id) - (updated_by_id)"
    accounts |o--o{ case : "(id) - (account_id)"
    organization ||--o{ case : "(id) - (org_id)"
    case_stage |o--o{ case : "(id) - (stage_id)"
    case |o--o{ case : "(id) - (merged_into_id)"
    profile |o--o{ case : "(id) - (merged_by_id)"
    case |o--o{ case : "(id) - (parent_id)"
    case }o--o{ contacts : "(contacts) - (case_contacts)"
    case }o--o{ profile : "(assigned_to) - (case_assigned_users)"
    case }o--o{ teams : "(teams) - (cases_teams)"
    case }o--o{ tags : "(tags) - (case_tags)"
    users |o--o{ case_watcher : "(id) - (created_by_id)"
    users |o--o{ case_watcher : "(id) - (updated_by_id)"
    case ||--o{ case_watcher : "(id) - (case_id)"
    profile ||--o{ case_watcher : "(id) - (profile_id)"
    organization ||--o{ case_watcher : "(id) - (org_id)"
    users |o--o{ csat_survey : "(id) - (created_by_id)"
    users |o--o{ csat_survey : "(id) - (updated_by_id)"
    case ||--o| csat_survey : "(id) - (case_id)"
    contacts |o--o{ csat_survey : "(id) - (contact_id)"
    organization ||--o{ csat_survey : "(id) - (org_id)"
    users |o--o{ solution : "(id) - (created_by_id)"
    users |o--o{ solution : "(id) - (updated_by_id)"
    organization ||--o{ solution : "(id) - (org_id)"
    solution }o--o{ case : "(cases) - (solutions)"
    solution }o--o{ tags : "(tags) - (solution_tags)"
    users |o--o{ case_pipeline : "(id) - (created_by_id)"
    users |o--o{ case_pipeline : "(id) - (updated_by_id)"
    organization ||--o{ case_pipeline : "(id) - (org_id)"
    users |o--o{ case_stage : "(id) - (created_by_id)"
    users |o--o{ case_stage : "(id) - (updated_by_id)"
    case_pipeline ||--o{ case_stage : "(id) - (pipeline_id)"
    organization ||--o{ case_stage : "(id) - (org_id)"
    users |o--o{ reopen_policy : "(id) - (created_by_id)"
    users |o--o{ reopen_policy : "(id) - (updated_by_id)"
    organization ||--o| reopen_policy : "(id) - (org_id)"
    users |o--o{ escalation_policy : "(id) - (created_by_id)"
    users |o--o{ escalation_policy : "(id) - (updated_by_id)"
    organization ||--o{ escalation_policy : "(id) - (org_id)"
    profile |o--o{ escalation_policy : "(id) - (first_response_target_id)"
    profile |o--o{ escalation_policy : "(id) - (resolution_target_id)"
    teams |o--o{ escalation_policy : "(id) - (notify_team_id)"
    users |o--o{ inbound_mailbox : "(id) - (created_by_id)"
    users |o--o{ inbound_mailbox : "(id) - (updated_by_id)"
    organization ||--o{ inbound_mailbox : "(id) - (org_id)"
    profile |o--o{ inbound_mailbox : "(id) - (default_assignee_id)"
    users |o--o{ email_message : "(id) - (created_by_id)"
    users |o--o{ email_message : "(id) - (updated_by_id)"
    organization ||--o{ email_message : "(id) - (org_id)"
    case |o--o{ email_message : "(id) - (case_id)"
    inbound_mailbox |o--o{ email_message : "(id) - (mailbox_id)"
    users |o--o{ routing_rule : "(id) - (created_by_id)"
    users |o--o{ routing_rule : "(id) - (updated_by_id)"
    organization ||--o{ routing_rule : "(id) - (org_id)"
    teams |o--o{ routing_rule : "(id) - (target_team_id)"
    routing_rule }o--o{ profile : "(target_assignees) - (routing_rules)"
    users |o--o{ routing_rule_state : "(id) - (created_by_id)"
    users |o--o{ routing_rule_state : "(id) - (updated_by_id)"
    organization ||--o{ routing_rule_state : "(id) - (org_id)"
    routing_rule ||--o| routing_rule_state : "(id) - (rule_id)"
    users |o--o{ time_entry : "(id) - (created_by_id)"
    users |o--o{ time_entry : "(id) - (updated_by_id)"
    organization ||--o{ time_entry : "(id) - (org_id)"
    case |o--o{ time_entry : "(id) - (case_id)"
    project |o--o{ time_entry : "(id) - (project_id)"
    task |o--o{ time_entry : "(id) - (task_id)"
    profile ||--o{ time_entry : "(id) - (profile_id)"
    invoice |o--o{ time_entry : "(id) - (invoice_id)"
    users |o--o{ approval_rule : "(id) - (created_by_id)"
    users |o--o{ approval_rule : "(id) - (updated_by_id)"
    organization ||--o{ approval_rule : "(id) - (org_id)"
    teams |o--o{ approval_rule : "(id) - (match_team_id)"
    approval_rule }o--o{ profile : "(approvers) - (approval_rules)"
    users |o--o{ approval : "(id) - (created_by_id)"
    users |o--o{ approval : "(id) - (updated_by_id)"
    organization ||--o{ approval : "(id) - (org_id)"
    case ||--o{ approval : "(id) - (case_id)"
    approval_rule ||--o{ approval : "(id) - (rule_id)"
    profile ||--o{ approval : "(id) - (requested_by_id)"
    profile |o--o{ approval : "(id) - (approver_id)"
    users |o--o{ contacts : "(id) - (created_by_id)"
    users |o--o{ contacts : "(id) - (updated_by_id)"
    organization ||--o{ contacts : "(id) - (org_id)"
    accounts |o--o{ contacts : "(id) - (account_id)"
    contacts }o--o{ profile : "(assigned_to) - (contact_assigned_users)"
    contacts }o--o{ teams : "(teams) - (contact_teams)"
    contacts }o--o{ tags : "(tags) - (contact_tags)"
    users |o--o{ lead : "(id) - (created_by_id)"
    users |o--o{ lead : "(id) - (updated_by_id)"
    organization ||--o{ lead : "(id) - (org_id)"
    lead_stage |o--o{ lead : "(id) - (stage_id)"
    profile |o--o{ lead : "(id) - (first_response_by_id)"
    lead }o--o{ profile : "(assigned_to) - (lead_assigned_users)"
    lead }o--o{ teams : "(teams) - (lead_teams)"
    lead }o--o{ tags : "(tags) - (lead_tags)"
    lead }o--o{ contacts : "(contacts) - (lead_contacts)"
    users |o--o{ lead_pipeline : "(id) - (created_by_id)"
    users |o--o{ lead_pipeline : "(id) - (updated_by_id)"
    organization ||--o{ lead_pipeline : "(id) - (org_id)"
    users |o--o{ lead_stage : "(id) - (created_by_id)"
    users |o--o{ lead_stage : "(id) - (updated_by_id)"
    lead_pipeline ||--o{ lead_stage : "(id) - (pipeline_id)"
    organization ||--o{ lead_stage : "(id) - (org_id)"
    users |o--o{ lead_source : "(id) - (created_by_id)"
    users |o--o{ lead_source : "(id) - (updated_by_id)"
    organization ||--o{ lead_source : "(id) - (org_id)"
    users |o--o{ lead_intake : "(id) - (created_by_id)"
    users |o--o{ lead_intake : "(id) - (updated_by_id)"
    organization ||--o{ lead_intake : "(id) - (org_id)"
    profile |o--o{ lead_intake : "(id) - (assignee_id)"
    lead_intake }o--o{ tags : "(tags) - (lead_intakes)"
    users |o--o{ lead_intake_event : "(id) - (created_by_id)"
    users |o--o{ lead_intake_event : "(id) - (updated_by_id)"
    organization ||--o{ lead_intake_event : "(id) - (org_id)"
    lead_intake ||--o{ lead_intake_event : "(id) - (intake_id)"
    lead |o--o{ lead_intake_event : "(id) - (lead_id)"
    users |o--o{ opportunity_pipeline : "(id) - (created_by_id)"
    users |o--o{ opportunity_pipeline : "(id) - (updated_by_id)"
    organization ||--o{ opportunity_pipeline : "(id) - (org_id)"
    users |o--o{ opportunity_pipeline_stage : "(id) - (created_by_id)"
    users |o--o{ opportunity_pipeline_stage : "(id) - (updated_by_id)"
    organization ||--o{ opportunity_pipeline_stage : "(id) - (org_id)"
    opportunity_pipeline ||--o{ opportunity_pipeline_stage : "(id) - (pipeline_id)"
    users |o--o{ opportunity : "(id) - (created_by_id)"
    users |o--o{ opportunity : "(id) - (updated_by_id)"
    accounts |o--o{ opportunity : "(id) - (account_id)"
    opportunity_pipeline ||--o{ opportunity : "(id) - (pipeline_id)"
    profile |o--o{ opportunity : "(id) - (closed_by_id)"
    organization ||--o{ opportunity : "(id) - (org_id)"
    opportunity }o--o{ contacts : "(contacts) - (opportunity_contacts)"
    opportunity }o--o{ profile : "(assigned_to) - (opportunity_assigned_users)"
    opportunity }o--o{ teams : "(teams) - (opportunity_teams)"
    opportunity }o--o{ tags : "(tags) - (opportunity_tags)"
    users |o--o{ opportunity_line_item : "(id) - (created_by_id)"
    users |o--o{ opportunity_line_item : "(id) - (updated_by_id)"
    opportunity ||--o{ opportunity_line_item : "(id) - (opportunity_id)"
    product |o--o{ opportunity_line_item : "(id) - (product_id)"
    organization ||--o{ opportunity_line_item : "(id) - (org_id)"
    users |o--o{ sales_goal : "(id) - (created_by_id)"
    users |o--o{ sales_goal : "(id) - (updated_by_id)"
    profile |o--o{ sales_goal : "(id) - (assigned_to_id)"
    teams |o--o{ sales_goal : "(id) - (team_id)"
    organization ||--o{ sales_goal : "(id) - (org_id)"
    users |o--o{ board : "(id) - (created_by_id)"
    users |o--o{ board : "(id) - (updated_by_id)"
    profile ||--o{ board : "(id) - (owner_id)"
    organization ||--o{ board : "(id) - (org_id)"
    users |o--o{ board_member : "(id) - (created_by_id)"
    users |o--o{ board_member : "(id) - (updated_by_id)"
    board ||--o{ board_member : "(id) - (board_id)"
    profile ||--o{ board_member : "(id) - (profile_id)"
    organization ||--o{ board_member : "(id) - (org_id)"
    users |o--o{ board_column : "(id) - (created_by_id)"
    users |o--o{ board_column : "(id) - (updated_by_id)"
    board ||--o{ board_column : "(id) - (board_id)"
    organization ||--o{ board_column : "(id) - (org_id)"
    users |o--o{ board_task : "(id) - (created_by_id)"
    users |o--o{ board_task : "(id) - (updated_by_id)"
    board_column ||--o{ board_task : "(id) - (column_id)"
    accounts |o--o{ board_task : "(id) - (account_id)"
    contacts |o--o{ board_task : "(id) - (contact_id)"
    opportunity |o--o{ board_task : "(id) - (opportunity_id)"
    organization ||--o{ board_task : "(id) - (org_id)"
    board_task }o--o{ profile : "(assigned_to) - (assigned_board_tasks)"
    users |o--o{ task_pipeline : "(id) - (created_by_id)"
    users |o--o{ task_pipeline : "(id) - (updated_by_id)"
    organization ||--o{ task_pipeline : "(id) - (org_id)"
    users |o--o{ task_stage : "(id) - (created_by_id)"
    users |o--o{ task_stage : "(id) - (updated_by_id)"
    task_pipeline ||--o{ task_stage : "(id) - (pipeline_id)"
    organization ||--o{ task_stage : "(id) - (org_id)"
    users |o--o{ task : "(id) - (created_by_id)"
    users |o--o{ task : "(id) - (updated_by_id)"
    accounts |o--o{ task : "(id) - (account_id)"
    opportunity |o--o{ task : "(id) - (opportunity_id)"
    case |o--o{ task : "(id) - (case_id)"
    lead |o--o{ task : "(id) - (lead_id)"
    project |o--o{ task : "(id) - (project_id)"
    organization ||--o{ task : "(id) - (org_id)"
    task_stage |o--o{ task : "(id) - (stage_id)"
    task }o--o{ contacts : "(contacts) - (task_contacts)"
    task }o--o{ profile : "(assigned_to) - (task_assigned_users)"
    task }o--o{ teams : "(teams) - (tasks_teams)"
    task }o--o{ tags : "(tags) - (task_tags)"
    users |o--o{ invoice_template : "(id) - (created_by_id)"
    users |o--o{ invoice_template : "(id) - (updated_by_id)"
    organization ||--o{ invoice_template : "(id) - (org_id)"
    users |o--o{ product : "(id) - (created_by_id)"
    users |o--o{ product : "(id) - (updated_by_id)"
    organization ||--o{ product : "(id) - (org_id)"
    users |o--o{ invoice : "(id) - (created_by_id)"
    users |o--o{ invoice : "(id) - (updated_by_id)"
    accounts |o--o{ invoice : "(id) - (account_id)"
    contacts |o--o{ invoice : "(id) - (contact_id)"
    opportunity |o--o{ invoice : "(id) - (opportunity_id)"
    project |o--o{ invoice : "(id) - (project_id)"
    invoice_template |o--o{ invoice : "(id) - (template_id)"
    organization ||--o{ invoice : "(id) - (org_id)"
    invoice }o--o{ profile : "(assigned_to) - (invoice_assigned_to)"
    invoice }o--o{ teams : "(teams) - (invoices_teams)"
    users |o--o{ invoice_line_item : "(id) - (created_by_id)"
    users |o--o{ invoice_line_item : "(id) - (updated_by_id)"
    invoice ||--o{ invoice_line_item : "(id) - (invoice_id)"
    product |o--o{ invoice_line_item : "(id) - (product_id)"
    organization ||--o{ invoice_line_item : "(id) - (org_id)"
    users |o--o{ payment : "(id) - (created_by_id)"
    users |o--o{ payment : "(id) - (updated_by_id)"
    invoice ||--o{ payment : "(id) - (invoice_id)"
    organization ||--o{ payment : "(id) - (org_id)"
    users |o--o{ estimate : "(id) - (created_by_id)"
    users |o--o{ estimate : "(id) - (updated_by_id)"
    accounts |o--o{ estimate : "(id) - (account_id)"
    contacts |o--o{ estimate : "(id) - (contact_id)"
    opportunity |o--o{ estimate : "(id) - (opportunity_id)"
    invoice |o--o{ estimate : "(id) - (converted_to_invoice_id)"
    project |o--o{ estimate : "(id) - (project_id)"
    invoice |o--o{ estimate : "(id) - (deposit_invoice_id)"
    organization ||--o{ estimate : "(id) - (org_id)"
    estimate }o--o{ profile : "(assigned_to) - (estimate_assigned_to)"
    estimate }o--o{ teams : "(teams) - (estimate_teams)"
    users |o--o{ estimate_line_item : "(id) - (created_by_id)"
    users |o--o{ estimate_line_item : "(id) - (updated_by_id)"
    estimate ||--o{ estimate_line_item : "(id) - (estimate_id)"
    product |o--o{ estimate_line_item : "(id) - (product_id)"
    organization ||--o{ estimate_line_item : "(id) - (org_id)"
    users |o--o{ recurring_invoice : "(id) - (created_by_id)"
    users |o--o{ recurring_invoice : "(id) - (updated_by_id)"
    accounts |o--o{ recurring_invoice : "(id) - (account_id)"
    contacts |o--o{ recurring_invoice : "(id) - (contact_id)"
    opportunity |o--o{ recurring_invoice : "(id) - (opportunity_id)"
    organization ||--o{ recurring_invoice : "(id) - (org_id)"
    recurring_invoice }o--o{ profile : "(assigned_to) - (recurring_invoice_assigned_to)"
    recurring_invoice }o--o{ teams : "(teams) - (recurring_invoice_teams)"
    users |o--o{ recurring_invoice_line_item : "(id) - (created_by_id)"
    users |o--o{ recurring_invoice_line_item : "(id) - (updated_by_id)"
    recurring_invoice ||--o{ recurring_invoice_line_item : "(id) - (recurring_invoice_id)"
    product |o--o{ recurring_invoice_line_item : "(id) - (product_id)"
    organization ||--o{ recurring_invoice_line_item : "(id) - (org_id)"
    users |o--o{ invoice_history : "(id) - (created_by_id)"
    invoice ||--o{ invoice_history : "(id) - (invoice_id)"
    profile |o--o{ invoice_history : "(id) - (updated_by_id)"
    organization ||--o{ invoice_history : "(id) - (org_id)"
    users |o--o{ orders : "(id) - (created_by_id)"
    users |o--o{ orders : "(id) - (updated_by_id)"
    organization ||--o{ orders : "(id) - (org_id)"
    accounts ||--o{ orders : "(id) - (account_id)"
    contacts |o--o{ orders : "(id) - (contact_id)"
    opportunity |o--o{ orders : "(id) - (opportunity_id)"
    users |o--o{ order_line_item : "(id) - (created_by_id)"
    users |o--o{ order_line_item : "(id) - (updated_by_id)"
    organization ||--o{ order_line_item : "(id) - (org_id)"
    orders ||--o{ order_line_item : "(id) - (order_id)"
    product |o--o{ order_line_item : "(id) - (product_id)"
    users |o--o{ business_calendar : "(id) - (created_by_id)"
    users |o--o{ business_calendar : "(id) - (updated_by_id)"
    organization ||--o{ business_calendar : "(id) - (org_id)"
    users |o--o{ business_holiday : "(id) - (created_by_id)"
    users |o--o{ business_holiday : "(id) - (updated_by_id)"
    business_calendar ||--o{ business_holiday : "(id) - (calendar_id)"
    organization ||--o{ business_holiday : "(id) - (org_id)"
    users |o--o{ macro : "(id) - (created_by_id)"
    users |o--o{ macro : "(id) - (updated_by_id)"
    profile |o--o{ macro : "(id) - (owner_id)"
    organization ||--o{ macro : "(id) - (org_id)"
    macro }o--o{ profile : "(set_assignees) - (macros_assigning)"
    macro }o--o{ tags : "(tags) - (macros_tagging)"
    users |o--o{ web_form : "(id) - (created_by_id)"
    users |o--o{ web_form : "(id) - (updated_by_id)"
    organization ||--o{ web_form : "(id) - (org_id)"
    profile |o--o{ web_form : "(id) - (assign_to_id)"
    profile |o--o{ web_form : "(id) - (rotation_last_assigned_id)"
    apiSettings |o--o| web_form : "(id) - (legacy_api_setting_id)"
    web_form }o--o{ profile : "(rotation_members) - (webforms_rotated)"
    web_form }o--o{ profile : "(notify_profiles) - (webforms_notified)"
    web_form }o--o{ tags : "(tags) - (webforms)"
    users |o--o{ web_form_field : "(id) - (created_by_id)"
    users |o--o{ web_form_field : "(id) - (updated_by_id)"
    organization ||--o{ web_form_field : "(id) - (org_id)"
    web_form ||--o{ web_form_field : "(id) - (form_id)"
    custom_field_definition |o--o{ web_form_field : "(id) - (custom_field_id)"
    users |o--o{ web_form_submission : "(id) - (created_by_id)"
    users |o--o{ web_form_submission : "(id) - (updated_by_id)"
    organization ||--o{ web_form_submission : "(id) - (org_id)"
    web_form ||--o{ web_form_submission : "(id) - (form_id)"
    lead |o--o{ web_form_submission : "(id) - (lead_id)"
    case |o--o{ web_form_submission : "(id) - (case_id)"
    users |o--o{ web_form_daily_stat : "(id) - (created_by_id)"
    users |o--o{ web_form_daily_stat : "(id) - (updated_by_id)"
    organization ||--o{ web_form_daily_stat : "(id) - (org_id)"
    web_form ||--o{ web_form_daily_stat : "(id) - (form_id)"
    users |o--o{ webhook_endpoint : "(id) - (created_by_id)"
    users |o--o{ webhook_endpoint : "(id) - (updated_by_id)"
    organization ||--o{ webhook_endpoint : "(id) - (org_id)"
    profile |o--o{ webhook_endpoint : "(id) - (app_id)"
    users |o--o{ webhook_delivery : "(id) - (created_by_id)"
    users |o--o{ webhook_delivery : "(id) - (updated_by_id)"
    organization ||--o{ webhook_delivery : "(id) - (org_id)"
    webhook_endpoint ||--o{ webhook_delivery : "(id) - (endpoint_id)"
    users |o--o{ automation_rule : "(id) - (created_by_id)"
    users |o--o{ automation_rule : "(id) - (updated_by_id)"
    organization ||--o{ automation_rule : "(id) - (org_id)"
    profile |o--o{ automation_rule : "(id) - (owner_id)"
    users |o--o{ automation_run : "(id) - (created_by_id)"
    users |o--o{ automation_run : "(id) - (updated_by_id)"
    organization ||--o{ automation_run : "(id) - (org_id)"
    automation_rule ||--o{ automation_run : "(id) - (rule_id)"
    automation_run |o--o{ automation_run : "(id) - (triggered_by_id)"
    users |o--o{ automation_rotation : "(id) - (created_by_id)"
    users |o--o{ automation_rotation : "(id) - (updated_by_id)"
    organization ||--o{ automation_rotation : "(id) - (org_id)"
    automation_rule ||--o{ automation_rotation : "(id) - (rule_id)"
    users |o--o{ interaction : "(id) - (created_by_id)"
    users |o--o{ interaction : "(id) - (updated_by_id)"
    organization ||--o{ interaction : "(id) - (org_id)"
    profile |o--o{ interaction : "(id) - (logged_by_id)"
    lead |o--o{ interaction : "(id) - (lead_id)"
    contacts |o--o{ interaction : "(id) - (contact_id)"
    accounts |o--o{ interaction : "(id) - (account_id)"
    opportunity |o--o{ interaction : "(id) - (opportunity_id)"
    case |o--o{ interaction : "(id) - (case_id)"
    users |o--o{ project : "(id) - (created_by_id)"
    users |o--o{ project : "(id) - (updated_by_id)"
    organization ||--o{ project : "(id) - (org_id)"
    accounts ||--o{ project : "(id) - (account_id)"
    opportunity |o--o{ project : "(id) - (deal_id)"
    project }o--o{ profile : "(assigned_to) - (assigned_projects)"
    project }o--o{ tags : "(tags) - (projects)"
    users |o--o{ project_rate : "(id) - (created_by_id)"
    users |o--o{ project_rate : "(id) - (updated_by_id)"
    organization ||--o{ project_rate : "(id) - (org_id)"
    project ||--o{ project_rate : "(id) - (project_id)"
    profile ||--o{ project_rate : "(id) - (profile_id)"
    auth_permission {
        AutoField id PK
        CharField name
        ForeignKey content_type_id FK
        CharField codename
    }
    auth_group {
        AutoField id PK
        CharField name UK
    }
    django_admin_log {
        AutoField id PK
        DateTimeField action_time
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        TextField object_id
        CharField object_repr
        LogEntryActionFlag action_flag
        TextField change_message
    }
    django_content_type {
        AutoField id PK
        CharField app_label
        CharField model
    }
    django_session {
        CharField session_key PK
        TextField session_data
        DateTimeField expire_date
    }
    token_blacklist_outstandingtoken {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField jti UK
        TextField token
        DateTimeField created_at
        DateTimeField expires_at
    }
    token_blacklist_blacklistedtoken {
        BigAutoField id PK
        OneToOneField token_id FK, UK
        DateTimeField blacklisted_at
    }
    django_ses_sesstat {
        AutoField id PK
        DateField date UK
        PositiveIntegerField delivery_attempts
        PositiveIntegerField bounces
        PositiveIntegerField complaints
        PositiveIntegerField rejects
    }
    django_ses_blacklistedemail {
        AutoField id PK
        EmailField email UK
    }
    users {
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        UUIDField id PK
        EmailField email UK
        CharField name
        CharField profile_pic
        CharField activation_key
        DateTimeField key_expires
        BooleanField is_active
        BooleanField is_staff
    }
    address {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField address_line
        CharField street
        CharField city
        CharField state
        CharField postcode
        AddressCountry country
        ForeignKey org_id FK
    }
    organization {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField api_key UK
        BooleanField is_active
        CharField company_name "Legal company name for invoices"
        ImageField logo "Company logo"
        CharField address_line
        CharField city
        CharField state
        CharField postcode
        OrgCountry country
        CharField phone
        EmailField email
        URLField website
        CharField tax_id "Tax ID / VAT / Registration number"
        OrgDefaultCurrency default_currency
        OrgDefaultCountry default_country
        CharField timezone "IANA timezone (e.g. America/New_York). Sets the org's day."
        BooleanField csat_enabled
        BooleanField auto_close_children_on_parent_close
        JSONField signup_source
        CharField vertical "Pack id applied to this org, e.g. 'real-estate'. Descriptive only."
        JSONField terminology "Label overrides, e.g. {'lead.plural': 'Enquiries'}."
        JSONField modules_off
        JSONField module_order
        JSONField modules_on_at
        DateTimeField lead_response_since
        DateTimeField lead_outcomes_since
        OrgLeadReminderMinutes lead_reminder_minutes
        DateTimeField lead_reminder_since
        BooleanField help_center_enabled
        CharField help_center_slug
        CharField invoice_number_prefix
        CharField estimate_number_prefix
        OrgNumberDatePart number_date_part
        PositiveSmallIntegerField number_padding
        PositiveIntegerField invoice_number_start
    }
    tags {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        CharField slug
        TagsColor color
        TextField description
        BooleanField is_active
        ForeignKey org_id FK
    }
    profile {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey user_id FK
        ForeignKey org_id FK
        CharField phone
        CharField alternate_phone
        ForeignKey address_id FK
        ProfileRole role
        BooleanField has_sales_access
        BooleanField has_marketing_access
        BooleanField is_active
        BooleanField is_organization_admin
        DateField date_of_joining
        BooleanField lead_alerts_email
        BooleanField lead_alerts_push
        BooleanField is_app
    }
    comment {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey content_type_id FK
        UUIDField object_id
        TextField comment
        DateTimeField commented_on
        ForeignKey commented_by_id FK
        ForeignKey commented_by_contact_id FK
        BooleanField is_internal
        ForeignKey org_id FK
    }
    commentFiles {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey comment_id FK
        FileField comment_file
        ForeignKey org_id FK
    }
    attachments {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey content_type_id FK
        UUIDField object_id
        CharField file_name
        FileField attachment
        ForeignKey org_id FK
    }
    document {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        TextField title
        FileField document_file
        DocumentStatus status
        ForeignKey org_id FK
    }
    apiSettings {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        TextField title
        CharField apikey
        URLField website
        ForeignKey org_id FK
    }
    magic_link_token {
        UUIDField id PK
        EmailField email
        CharField token UK
        MagicLinkTokenDelivery delivery
        CharField code_hash
        PositiveSmallIntegerField attempts
        DateTimeField created_at
        DateTimeField expires_at
        BooleanField is_used
        DateTimeField used_at
        GenericIPAddressField ip_address
    }
    portal_login_token {
        UUIDField id PK
        ForeignKey org_id FK
        ForeignKey contact_id FK
        CharField code_hash
        PositiveSmallIntegerField attempts
        DateTimeField created_at
        DateTimeField expires_at
        BooleanField is_used
        DateTimeField used_at
        GenericIPAddressField ip_address
    }
    activity {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey user_id FK
        ActivityAction action
        ActivityEntityType entity_type
        UUIDField entity_id
        CharField entity_name
        TextField description
        JSONField metadata
        ForeignKey org_id FK
    }
    notification {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey recipient_id FK
        CharField verb
        ForeignKey actor_id FK
        CharField entity_type
        UUIDField entity_id
        CharField entity_name
        JSONField data
        CharField link
        DateTimeField read_at
        ForeignKey org_id FK
    }
    device_token {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey user_id FK
        CharField token
        DeviceTokenPlatform platform
        CharField app_version
        DateTimeField last_seen_at
    }
    teams {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        ForeignKey org_id FK
    }
    contact_form_submission {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        EmailField email
        TextField message
        ContactFormSubmissionReason reason
        GenericIPAddressField ip_address
        TextField user_agent
        URLField referrer
        ContactFormSubmissionStatus status
        ForeignKey replied_by_id FK
        DateTimeField replied_at
    }
    custom_field_definition {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        CharField target_model
        SlugField key
        CharField label
        CustomFieldDefinitionFieldType field_type
        JSONField options "List of {value, label} pairs for dropdown fields."
        BooleanField is_required
        BooleanField is_filterable
        PositiveIntegerField display_order
        BooleanField is_active
    }
    personal_access_token {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey profile_id FK
        CharField name
        CharField token_hash UK
        CharField token_prefix
        JSONField scopes
        DateTimeField expires_at
        DateTimeField last_used_at
        DateTimeField revoked_at
    }
    calendar_feed_token {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        OneToOneField profile_id FK, UK
        CharField token_hash UK
        DateTimeField last_used_at
    }
    portal_access_token {
        UUIDField id PK
        CharField token_hash UK
        ForeignKey org_id FK
        PortalAccessTokenResourceType resource_type
        UUIDField resource_id
        DateTimeField created_at
    }
    security_audit_log {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        SecurityAuditLogEventType event_type
        ForeignKey user_id FK
        ForeignKey org_id FK
        TextField description
        JSONField metadata
        GenericIPAddressField ip_address
        TextField user_agent
        CharField request_path
        CharField request_method
        BooleanField success
    }
    pack_application {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField pack_id
        PositiveIntegerField pack_version
        ForeignKey applied_by_id FK
        JSONField result "{'created': [...], 'skipped': [...], 'failed': [...]}"
    }
    installed_app {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField app_id
        PositiveIntegerField installed_version
        InstalledAppStatus status
        ForeignKey installed_by_id FK
        DateTimeField installed_at
        DateTimeField removed_at
        JSONField owned
        JSONField offered
    }
    saved_view {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey profile_id FK
        CharField module
        CharField name
        JSONField filters
    }
    accounts {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        EmailField email
        CharField phone
        URLField website
        AccountIndustry industry
        PositiveIntegerField number_of_employees
        DecimalField annual_revenue
        AccountCurrency currency
        CharField address_line
        CharField city
        CharField state
        CharField postcode
        AccountCountry country
        TextField description
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        BooleanField is_active
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        ForeignKey org_id FK
    }
    account_email {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey from_account_id FK
        TextField message_subject
        TextField message_body
        CharField timezone
        DateTimeField scheduled_date_time
        BooleanField scheduled_later
        EmailField from_email
        TextField rendered_message_body
        ForeignKey org_id FK
    }
    emailLogs {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey email_id FK
        ForeignKey contact_id FK
        BooleanField is_sent
        ForeignKey org_id FK
    }
    case {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        CaseStatus status
        CasePriority priority
        CaseCaseType case_type
        ForeignKey account_id FK
        DateField closed_on
        TextField description
        BooleanField is_active
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        ForeignKey org_id FK
        DateTimeField first_response_at
        DateTimeField resolved_at
        PositiveIntegerField sla_first_response_hours "Target hours for first response"
        PositiveIntegerField sla_resolution_hours "Target hours for resolution"
        ForeignKey stage_id FK "Current pipeline stage (null = use status-based kanban)"
        DecimalField kanban_order "Order within the kanban column for drag-drop positioning"
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        DateTimeField sla_paused_at
        PositiveIntegerField sla_paused_seconds
        DateTimeField last_escalation_fired_at
        PositiveSmallIntegerField escalation_count
        CharField external_thread_id
        ForeignKey merged_into_id FK
        DateTimeField merged_at
        ForeignKey merged_by_id FK
        JSONField alt_thread_ids
        JSONField merge_record
        ForeignKey parent_id FK
        BooleanField is_problem "Marks this case as an ITIL 'problem' (umbrella ticket)."
    }
    case_watcher {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey case_id FK
        ForeignKey profile_id FK
        CaseWatcherSubscribedVia subscribed_via
        ForeignKey org_id FK
    }
    csat_survey {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        OneToOneField case_id FK, UK
        ForeignKey contact_id FK
        CharField token_hash UK
        DateTimeField sent_at
        PositiveSmallIntegerField rating
        TextField comment
        DateTimeField responded_at
        DateTimeField expires_at
        ForeignKey org_id FK
    }
    solution {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField title
        TextField description
        SolutionStatus status
        BooleanField is_published
        ForeignKey org_id FK
    }
    case_pipeline {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        ForeignKey org_id FK
        BooleanField is_default "Sorts this pipeline first in listings (see Meta.ordering) and is limited to one per org. It does NOT route new records: a case created without a stage keeps stage=NULL, which is the supported status-based kanban mode."
        BooleanField is_active
    }
    case_stage {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey pipeline_id FK
        CharField name
        PositiveIntegerField order
        CharField color
        CaseStageStageType stage_type
        CaseStageMapsToStatus maps_to_status "When case enters this stage, also update Case.status"
        PositiveIntegerField wip_limit "Maximum cases allowed in this stage (null = unlimited)"
        ForeignKey org_id FK
    }
    reopen_policy {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        OneToOneField org_id FK, UK
        BooleanField is_enabled
        PositiveSmallIntegerField reopen_window_days "Replies within this many days of closed_on reopen the case."
        ReopenPolicyReopenToStatus reopen_to_status "Status the case flips to on reopen (must be non-terminal)."
        BooleanField notify_assigned
    }
    escalation_policy {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        EscalationPolicyPriority priority
        PositiveIntegerField first_response_hours "Target hours for first response. Blank uses the built-in default."
        PositiveIntegerField resolution_hours "Target hours for resolution. Blank uses the built-in default."
        PositiveIntegerField next_response_hours "Target hours for each reply after the first. Blank uses the built-in default."
        EscalationPolicyFirstResponseAction first_response_action
        EscalationPolicyResolutionAction resolution_action
        ForeignKey first_response_target_id FK "Profile reassigned to and/or notified on first-response breach."
        ForeignKey resolution_target_id FK "Profile reassigned to and/or notified on resolution breach."
        ForeignKey notify_team_id FK "Optional team CC'd on either breach action."
        BooleanField is_active
    }
    inbound_mailbox {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        EmailField address
        InboundMailboxProvider provider
        CharField webhook_secret "Postmark's basic-auth password or Mailgun's HTTP webhook signing key, which the webhook checks on every delivery; a Postmark or Mailgun mailbox without one accepts nothing. Unused for SES, which is proven by the SNS signature and the topic_arn pin below. Write-only: the API never returns it."
        CharField topic_arn "SNS Topic ARN this mailbox accepts mail from. A valid SNS signature only proves a message came from some topic in some AWS account, so the webhook also requires this exact match. An admin can enter it. Left blank, it is pinned from the first verified SubscriptionConfirmation whose topic belongs to an AWS account listed in the INBOUND_SNS_ACCOUNT_IDS setting, and a topic in any other account pins nothing. Until it is set, every message is rejected."
        CharField imap_host
        PositiveIntegerField imap_port
        CharField imap_username
        CharField imap_password_enc
        InboundMailboxDefaultPriority default_priority
        InboundMailboxDefaultCaseType default_case_type
        ForeignKey default_assignee_id FK
        BooleanField is_active
        InboundMailboxCreates creates
        CharField lead_source
        CharField lead_source_detail
        BooleanField is_generated
    }
    email_message {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        ForeignKey case_id FK "Null when the message was dropped (spam/bounce/auto-reply) but we still want an audit trail."
        ForeignKey mailbox_id FK "The inbound address this message arrived through. Set at ingest; SET_NULL keeps the audit row if the mailbox is later deleted. Null on historical rows the backfill could not attribute."
        EmailMessageDirection direction
        CharField message_id
        CharField in_reply_to
        TextField references "Whitespace-separated list of RFC-5322 Message-IDs from the References header."
        EmailField from_address
        TextField to_addresses
        TextField cc_addresses
        CharField subject
        TextField body_text
        TextField body_html
        DateTimeField received_at
        CharField drop_reason "If non-empty, the message was rejected (spam/bounce/etc.) before any Case was touched. `case` will be null in this case."
    }
    routing_rule {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        CharField name
        PositiveIntegerField priority_order "Lower runs first; ties broken by created_at."
        BooleanField is_active
        JSONField conditions "List of {field, op, value} triples; AND across entries."
        RoutingRuleStrategy strategy
        ForeignKey target_team_id FK
        BooleanField stop_processing "If a rule matches and this is True, skip lower-priority rules."
    }
    routing_rule_state {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        OneToOneField rule_id FK, UK
        PositiveIntegerField last_assigned_index
    }
    time_entry {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        ForeignKey case_id FK
        ForeignKey project_id FK
        ForeignKey task_id FK
        ForeignKey profile_id FK
        DateTimeField started_at
        DateTimeField ended_at "Null while the timer is running."
        PositiveIntegerField duration_minutes "Recomputed on save when both timestamps are present."
        TextField description
        BooleanField billable
        DecimalField hourly_rate "Snapshot at log time so future rate changes do not alter history."
        TimeEntryCurrency currency
        ForeignKey invoice_id FK "Set when this entry has been pushed to a draft invoice."
        BooleanField auto_stopped "Set by the auto-stop Celery task when a forgotten timer is killed."
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
    }
    approval_rule {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        ForeignKey org_id FK
        ApprovalRuleTriggerEvent trigger_event
        ApprovalRuleMatchPriority match_priority
        ApprovalRuleMatchCaseType match_case_type
        ForeignKey match_team_id FK
        ApprovalRuleApproverRole approver_role
        BooleanField is_active
    }
    approval {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK
        ForeignKey case_id FK
        ForeignKey rule_id FK
        ForeignKey requested_by_id FK
        ForeignKey approver_id FK
        ApprovalState state
        TextField note
        TextField reason
        DateTimeField decided_at
    }
    contacts {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField first_name
        CharField last_name
        EmailField email
        CharField phone
        CharField phone_e164
        CharField organization
        CharField title
        CharField department
        BooleanField do_not_call
        URLField linkedin_url
        CharField address_line
        CharField city
        CharField state
        CharField postcode
        ContactCountry country
        TextField description
        CharField source
        CharField source_detail
        CharField source_ref
        BooleanField is_active
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        BooleanField auto_created "True when the contact was created by an inbound channel (e.g. email-to-ticket) rather than a human; flagged for admin review."
        ForeignKey org_id FK
        ForeignKey account_id FK "Primary account this contact belongs to"
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
    }
    lead {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField title "Lead name/subject (e.g., 'Enterprise Deal', 'Website Inquiry')"
        CharField salutation "e.g., Mr, Mrs, Ms, Dr"
        CharField first_name
        CharField last_name
        EmailField email
        CharField phone
        CharField phone_e164
        CharField job_title "Person's job title (e.g., 'VP of Sales', 'CTO')"
        CharField website
        URLField linkedin_url
        LeadStatus status
        CharField source
        CharField source_detail
        CharField source_ref
        LeadIndustry industry
        LeadRating rating
        DecimalField opportunity_amount
        LeadCurrency currency
        IntegerField probability
        DateField close_date
        CharField address_line
        CharField city
        CharField state
        CharField postcode
        LeadCountry country
        DateField last_contacted
        DateField next_follow_up
        TextField description
        BooleanField is_active
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        ForeignKey org_id FK
        CharField company_name
        ForeignKey stage_id FK "Current pipeline stage (null = use status-based kanban)"
        DecimalField kanban_order "Order within the kanban column for drag-drop"
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        CharField external_source
        CharField external_id
        DateTimeField last_enquiry_at
        DateTimeField first_response_at
        ForeignKey first_response_by_id FK
        DateTimeField last_outbound_at
        DateTimeField waiting_reminded_at
        DateTimeField won_at
        DateTimeField lost_at
    }
    lead_pipeline {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        ForeignKey org_id FK
        BooleanField is_default "Sorts this pipeline first in listings (see Meta.ordering) and is limited to one per org. It does NOT route new records: a lead created without a stage keeps stage=NULL, which is the supported status-based kanban mode."
        BooleanField is_active
    }
    lead_stage {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey pipeline_id FK
        CharField name
        PositiveIntegerField order
        CharField color
        LeadStageStageType stage_type
        LeadStageMapsToStatus maps_to_status "When lead enters this stage, also update Lead.status"
        IntegerField win_probability "Default probability when lead enters this stage"
        PositiveIntegerField wip_limit "Maximum leads allowed in this stage (null = unlimited)"
        ForeignKey org_id FK
    }
    lead_source {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField key
        CharField label
        BooleanField is_active
        IntegerField order
    }
    lead_intake {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        LeadIntakePreset preset
        CharField key
        CharField previous_key
        DateTimeField previous_key_expires_at
        JSONField mapping
        CharField default_source
        CharField source_detail
        ForeignKey assignee_id FK
        BooleanField is_active
        DateTimeField last_received_at
        TextField last_error
    }
    lead_intake_event {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey intake_id FK
        DateTimeField received_at
        LeadIntakeEventResult result
        ForeignKey lead_id FK
        CharField external_id
        TextField error
        JSONField payload
    }
    opportunity_pipeline {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        BooleanField is_default
    }
    opportunity_pipeline_stage {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey pipeline_id FK
        CharField code
        CharField label
        PositiveIntegerField order
        DealStageKind kind
        PositiveIntegerField expected_days
        PositiveIntegerField warning_days
    }
    opportunity {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        ForeignKey account_id FK
        CharField stage
        ForeignKey pipeline_id FK
        OpportunityOpportunityType opportunity_type
        OpportunityCurrency currency
        DecimalField amount
        OpportunityAmountSource amount_source
        IntegerField probability
        DateField closed_on
        CharField lead_source
        CharField source_detail
        CharField source_ref
        ForeignKey closed_by_id FK
        TextField description
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        DateTimeField stage_changed_at
        DecimalField kanban_order "Order within the kanban column for drag-drop positioning"
        BooleanField is_active
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        ForeignKey org_id FK
    }
    opportunity_line_item {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey opportunity_id FK
        ForeignKey product_id FK
        CharField name
        CharField description
        DecimalField quantity
        DecimalField unit_price
        OpportunityLineItemDiscountType discount_type
        DecimalField discount_value
        DecimalField discount_amount
        DecimalField subtotal
        DecimalField total
        PositiveIntegerField order
        ForeignKey org_id FK
    }
    sales_goal {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        SalesGoalGoalType goal_type
        DecimalField target_value
        SalesGoalCurrency currency
        SalesGoalPeriodType period_type
        DateField period_start
        DateField period_end
        ForeignKey assigned_to_id FK
        ForeignKey team_id FK
        JSONField type_weights "Optional multiplier per opportunity type, e.g. {'RENEWAL': 0.5, 'NEW_BUSINESS': 1.5}. A type left out of the map, and a deal with no type at all, counts at its full value. An empty map is an unweighted goal. Not accepted on an ACTIVITIES goal, which has no deal type to weigh."
        BooleanField is_active
        BooleanField milestone_50_notified
        BooleanField milestone_90_notified
        BooleanField milestone_100_notified
        ForeignKey org_id FK
    }
    board {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        ForeignKey owner_id FK
        ForeignKey org_id FK
        BooleanField is_archived
    }
    board_member {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey board_id FK
        ForeignKey profile_id FK
        BoardMemberRole role
        ForeignKey org_id FK
    }
    board_column {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey board_id FK
        CharField name
        PositiveIntegerField order
        CharField color
        PositiveIntegerField limit
        ForeignKey org_id FK
    }
    board_task {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey column_id FK
        CharField title
        TextField description
        PositiveIntegerField order
        BoardTaskPriority priority
        DateTimeField due_date
        DateTimeField completed_at
        ForeignKey account_id FK
        ForeignKey contact_id FK
        ForeignKey opportunity_id FK
        ForeignKey org_id FK
    }
    task_pipeline {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        ForeignKey org_id FK
        BooleanField is_default "Sorts this pipeline first in listings (see Meta.ordering) and is limited to one per org. It does NOT route new records: a task created without a stage keeps stage=NULL, which is the supported status-based kanban mode."
        BooleanField is_active
    }
    task_stage {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey pipeline_id FK
        CharField name
        PositiveIntegerField order
        CharField color
        TaskStageStageType stage_type
        CharField maps_to_status "When task enters this stage, also update Task.status"
        PositiveIntegerField wip_limit "Maximum tasks allowed in this stage (null = unlimited)"
        ForeignKey org_id FK
    }
    task {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField title
        TaskStatus status
        TaskPriority priority
        DateField due_date
        TextField description
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        ForeignKey account_id FK
        ForeignKey opportunity_id FK
        ForeignKey case_id FK
        ForeignKey lead_id FK
        ForeignKey project_id FK
        ForeignKey org_id FK
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
        DecimalField kanban_order "Order within the kanban column for drag-drop positioning"
        ForeignKey stage_id FK "Custom pipeline stage (if using pipeline mode)"
    }
    invoice_template {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        ImageField logo
        CharField primary_color
        CharField secondary_color
        TextField template_html
        TextField template_css
        TextField default_notes
        TextField default_terms
        TextField footer_text
        BooleanField is_default
        ForeignKey org_id FK
    }
    product {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        TextField description
        CharField sku
        DecimalField price
        ProductCurrency currency
        CharField category
        BooleanField is_active
        ForeignKey org_id FK
    }
    invoice {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField invoice_title
        CharField invoice_number
        InvoiceStatus status
        ForeignKey account_id FK "The legal entity responsible for payment (REQUIRED via validation)"
        ForeignKey contact_id FK "The 'Bill To' contact (REQUIRED via validation)"
        ForeignKey opportunity_id FK "The sales opportunity this invoice relates to (optional)"
        ForeignKey project_id FK
        CharField client_name
        EmailField client_email
        CharField client_phone
        CharField billing_address_line
        CharField billing_city
        CharField billing_state
        CharField billing_postcode
        InvoiceBillingCountry billing_country
        CharField client_address_line
        CharField client_city
        CharField client_state
        CharField client_postcode
        InvoiceClientCountry client_country
        DecimalField subtotal
        InvoiceDiscountType discount_type
        DecimalField discount_value
        DecimalField discount_amount
        DecimalField tax_rate
        DecimalField tax_amount
        DecimalField shipping_amount
        DecimalField total_amount
        InvoiceCurrency currency
        DecimalField amount_paid
        DecimalField amount_due
        DateField issue_date
        DateField due_date
        InvoicePaymentTerms payment_terms
        DateTimeField sent_at
        DateTimeField viewed_at
        DateTimeField paid_at
        DateTimeField cancelled_at
        BooleanField reminder_enabled
        PositiveIntegerField reminder_days_before
        PositiveIntegerField reminder_days_after
        InvoiceReminderFrequency reminder_frequency
        DateTimeField last_reminder_sent
        PositiveIntegerField reminder_count
        CharField public_token UK
        BooleanField public_link_enabled
        ForeignKey template_id FK
        TextField notes
        TextField terms
        TextField details
        CharField billing_period
        CharField po_number
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        BooleanField is_email_sent
        ForeignKey org_id FK
    }
    invoice_line_item {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey invoice_id FK
        ForeignKey product_id FK
        CharField name
        CharField description
        DecimalField quantity
        DecimalField unit_price
        InvoiceLineItemDiscountType discount_type
        DecimalField discount_value
        DecimalField discount_amount
        DecimalField tax_rate
        DecimalField tax_amount
        DecimalField subtotal
        DecimalField total
        PositiveIntegerField order
        ForeignKey org_id FK
    }
    payment {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey invoice_id FK
        DecimalField amount
        DateField payment_date
        PaymentPaymentMethod payment_method
        CharField reference_number
        TextField notes
        ForeignKey org_id FK
    }
    estimate {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField estimate_number
        CharField title
        EstimateStatus status
        ForeignKey account_id FK "The legal entity (REQUIRED via validation)"
        ForeignKey contact_id FK "The contact person (REQUIRED via validation)"
        ForeignKey opportunity_id FK
        CharField client_name
        EmailField client_email
        CharField client_phone
        CharField client_address_line
        CharField client_city
        CharField client_state
        CharField client_postcode
        EstimateClientCountry client_country
        DecimalField subtotal
        EstimateDiscountType discount_type
        DecimalField discount_value
        DecimalField discount_amount
        DecimalField tax_rate
        DecimalField tax_amount
        DecimalField total_amount
        EstimateCurrency currency
        DateField issue_date
        DateField expiry_date
        DateTimeField sent_at
        DateTimeField viewed_at
        DateTimeField accepted_at
        DateTimeField declined_at
        CharField accepted_by_name
        EmailField accepted_by_email
        GenericIPAddressField accepted_ip
        TextField accepted_user_agent
        ForeignKey converted_to_invoice_id FK
        TextField introduction
        DecimalField deposit_percent
        BooleanField start_project
        EstimateProjectBillingType project_billing_type
        ForeignKey project_id FK
        ForeignKey deposit_invoice_id FK
        CharField public_token UK
        BooleanField public_link_enabled
        TextField notes
        TextField terms
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        ForeignKey org_id FK
    }
    estimate_line_item {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey estimate_id FK
        ForeignKey product_id FK
        CharField name
        CharField description
        DecimalField quantity
        DecimalField unit_price
        EstimateLineItemDiscountType discount_type
        DecimalField discount_value
        DecimalField discount_amount
        DecimalField tax_rate
        DecimalField tax_amount
        DecimalField subtotal
        DecimalField total
        PositiveIntegerField order
        ForeignKey org_id FK
    }
    recurring_invoice {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField title
        BooleanField is_active
        ForeignKey account_id FK
        ForeignKey contact_id FK
        ForeignKey opportunity_id FK
        CharField client_name
        EmailField client_email
        RecurringInvoiceFrequency frequency
        PositiveIntegerField custom_days
        DateField start_date
        DateField end_date
        DateField next_generation_date
        RecurringInvoicePaymentTerms payment_terms
        BooleanField auto_send
        RecurringInvoiceCurrency currency
        DecimalField subtotal
        RecurringInvoiceDiscountType discount_type
        DecimalField discount_value
        DecimalField tax_rate
        DecimalField total_amount
        TextField notes
        TextField terms
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        PositiveIntegerField invoices_generated
        ForeignKey org_id FK
    }
    recurring_invoice_line_item {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey recurring_invoice_id FK
        ForeignKey product_id FK
        CharField name
        CharField description
        DecimalField quantity
        DecimalField unit_price
        RecurringInvoiceLineItemDiscountType discount_type
        DecimalField discount_value
        DecimalField tax_rate
        PositiveIntegerField order
        ForeignKey org_id FK
    }
    invoice_history {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        UUIDField id PK
        ForeignKey invoice_id FK
        CharField invoice_title
        CharField invoice_number
        InvoiceHistoryStatus status
        CharField client_name
        EmailField client_email
        DecimalField total_amount
        DecimalField amount_due
        InvoiceHistoryCurrency currency
        DateField due_date
        ForeignKey updated_by_id FK
        TextField details
        ForeignKey org_id FK
    }
    orders {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        CharField order_number
        OrderStatus status
        ForeignKey account_id FK "The account this order belongs to"
        ForeignKey contact_id FK "The contact associated with this order"
        ForeignKey opportunity_id FK "The opportunity this order originated from"
        CharField currency
        DecimalField subtotal
        DecimalField discount_amount
        DecimalField tax_amount
        DecimalField total_amount
        DateField order_date
        DateField activated_date
        DateField shipped_date
        CharField billing_address_line
        CharField billing_city
        CharField billing_state
        CharField billing_postcode
        OrderBillingCountry billing_country
        CharField shipping_address_line
        CharField shipping_city
        CharField shipping_state
        CharField shipping_postcode
        OrderShippingCountry shipping_country
        TextField description
    }
    order_line_item {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey order_id FK
        ForeignKey product_id FK
        CharField name
        TextField description
        DecimalField quantity
        DecimalField unit_price
        DecimalField discount_amount
        DecimalField total
        IntegerField sort_order
    }
    business_calendar {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField name
        CharField timezone "IANA timezone (e.g. America/New_York)."
        BooleanField is_default
        TimeField monday_open
        TimeField monday_close
        TimeField tuesday_open
        TimeField tuesday_close
        TimeField wednesday_open
        TimeField wednesday_close
        TimeField thursday_open
        TimeField thursday_close
        TimeField friday_open
        TimeField friday_close
        TimeField saturday_open
        TimeField saturday_close
        TimeField sunday_open
        TimeField sunday_close
        ForeignKey org_id FK
    }
    business_holiday {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey calendar_id FK
        DateField date
        CharField name
        ForeignKey org_id FK
    }
    macro {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        CharField title
        TextField body
        MacroScope scope
        ForeignKey owner_id FK
        BooleanField is_active
        PositiveIntegerField usage_count
        MacroSetStatus set_status
        MacroSetPriority set_priority
        ForeignKey org_id FK
    }
    web_form {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        WebFormTarget target "What an accepted submission creates. Fixed once the form has submissions, so its history never mixes leads and tickets."
        BooleanField is_published "Only published forms accept submissions. Flipped through the publish and unpublish endpoints, never through a plain update."
        JSONField allowed_origins "Origins permitted to frame this form or POST to it. An empty list means unrestricted."
        CharField submit_button_label
        WebFormSuccessMode success_mode
        TextField success_message
        URLField redirect_url
        WebFormAssignmentMode assignment_mode
        ForeignKey assign_to_id FK
        PositiveIntegerField rotation_cap
        ForeignKey rotation_last_assigned_id FK
        CharField lead_source
        WebFormTicketPriority ticket_priority
        WebFormTicketType ticket_type
        WebFormCaptchaProvider captcha_provider
        CharField captcha_site_key
        CharField captcha_secret
        BooleanField reject_disposable_email
        OneToOneField legacy_api_setting_id FK, UK "Set by the backfill migration. Links a legacy web-to-lead API key to the form that now backs it."
    }
    web_form_field {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey form_id FK
        PositiveIntegerField order
        WebFormFieldSource source
        WebFormFieldLeadField lead_field
        WebFormFieldTicketField ticket_field
        ForeignKey custom_field_id FK
        CharField label
        CharField placeholder
        BooleanField is_required
    }
    web_form_submission {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey form_id FK
        ForeignKey lead_id FK
        ForeignKey case_id FK
        JSONField payload
        WebFormSubmissionStatus status
        CharField reject_reason
        GenericIPAddressField submitted_ip
        CharField referer
    }
    web_form_daily_stat {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey form_id FK
        DateField date
        PositiveIntegerField views
    }
    webhook_endpoint {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        URLField url
        CharField description
        JSONField events
        WebhookEndpointFormat format
        BooleanField is_active
        CharField disabled_reason
        CharField secret
        ForeignKey app_id FK
    }
    webhook_delivery {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey endpoint_id FK
        CharField event
        UUIDField event_id
        JSONField payload
        WebhookDeliveryStatus status
        PositiveSmallIntegerField attempts
        PositiveSmallIntegerField response_status
        CharField error
        DateTimeField next_attempt_at
        DateTimeField last_attempt_at
    }
    automation_rule {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        CharField object_key
        JSONField trigger
        AutomationRuleMatch match
        JSONField conditions
        JSONField actions
        BooleanField is_active
        PositiveIntegerField position
        ForeignKey owner_id FK
        DateTimeField switched_on_at
        CharField paused_reason
        DateTimeField last_failure_notice_at
    }
    automation_run {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey rule_id FK
        CharField object_key
        UUIDField record_id
        CharField record_name
        AutomationRunStatus status
        CharField reason
        UUIDField chain_id
        PositiveSmallIntegerField depth
        ForeignKey triggered_by_id FK
        JSONField steps
        PositiveIntegerField skipped_count
        DateTimeField window_start
    }
    automation_rotation {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey rule_id FK
        PositiveSmallIntegerField action_index
        PositiveIntegerField next_index
    }
    interaction {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        InteractionKind kind
        InteractionDirection direction
        DateTimeField occurred_at
        PositiveIntegerField duration_seconds
        BooleanField duration_estimated
        InteractionOutcome outcome
        CharField subject
        TextField body
        ForeignKey logged_by_id FK
        JSONField participants
        ForeignKey lead_id FK
        ForeignKey contact_id FK
        ForeignKey account_id FK
        ForeignKey opportunity_id FK
        ForeignKey case_id FK
        InteractionHome home
        CharField origin
        CharField external_id
    }
    project {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        CharField name
        ForeignKey account_id FK
        ForeignKey deal_id FK
        ProjectStatus status
        ProjectBillingType billing_type
        ProjectCurrency currency
        DecimalField budget_hours
        DecimalField budget_amount
        DecimalField hourly_rate
        DateField start_date
        DateField due_date
        TextField description
        JSONField custom_fields "Per-org schema extension; values are validated against common.CustomFieldDefinition."
        DecimalField retainer_fee
        DecimalField retainer_hours "Hours each month's fee includes; none means the fee alone."
        BooleanField retainer_rollover
        DecimalField retainer_overage_rate
        DecimalField retainer_carried_hours
        DateField retainer_next_month "First day of the month the next retainer invoice bills."
        DateField retainer_last_month "First day of the last month to bill; empty while the retainer is open."
        ProjectHoursAlert hours_alert
        ProjectAmountAlert amount_alert
        BooleanField is_sample "True only for demo rows created by common.packs.applier._apply_sample_data. Server-set exclusively, never expose this as a writable field on any serializer. It is the sole key common.packs.applier.clear_sample_data uses to decide what to delete, so a client-writable path here would let a user mark an arbitrary real record as sample and have it deleted."
    }
    project_rate {
        DateTimeField created_at
        DateTimeField updated_at
        ForeignKey created_by_id FK
        ForeignKey updated_by_id FK
        UUIDField id PK
        ForeignKey org_id FK "Organization this record belongs to"
        ForeignKey project_id FK
        ForeignKey profile_id FK
        DecimalField hourly_rate
    }
```
