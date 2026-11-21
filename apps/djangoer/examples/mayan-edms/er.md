```mermaid
erDiagram
    auth_user ||--o{ events_eventsubscription : "(id) - (user_id)"
    events_storedeventtype ||--o{ events_eventsubscription : "(id) - (stored_event_type_id)"
    auth_user ||--o{ events_notification : "(id) - (user_id)"
    actstream_action ||--o{ events_notification : "(id) - (action_id)"
    django_content_type ||--o{ events_objecteventsubscription : "(id) - (content_type_id)"
    auth_user ||--o{ events_objecteventsubscription : "(id) - (user_id)"
    events_storedeventtype ||--o{ events_objecteventsubscription : "(id) - (stored_event_type_id)"
    auth_user ||--o{ appearance_usersetting : "(id) - (user_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_group ||..o{ actstream_action : "(id) - (actor_object_id)"
    auth_group ||..o{ actstream_action : "(id) - (target_object_id)"
    auth_group ||..o{ actstream_action : "(id) - (action_object_object_id)"
    auth_group ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||..o{ actstream_action : "(id) - (actor_object_id)"
    auth_user ||..o{ actstream_action : "(id) - (target_object_id)"
    auth_user ||..o{ actstream_action : "(id) - (action_object_object_id)"
    auth_user ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    auth_user ||--o{ actstream_follow : "(id) - (user_id)"
    django_content_type ||--o{ actstream_follow : "(id) - (content_type_id)"
    django_content_type ||--o{ actstream_action : "(id) - (actor_content_type_id)"
    django_content_type |o--o{ actstream_action : "(id) - (target_content_type_id)"
    django_content_type |o--o{ actstream_action : "(id) - (action_object_content_type_id)"
    axes_accessattempt ||--o| axes_accessattemptexpiration : "(id) - (access_attempt_id)"
    django_celery_beat_intervalschedule |o--o{ django_celery_beat_periodictask : "(id) - (interval_id)"
    django_celery_beat_crontabschedule |o--o{ django_celery_beat_periodictask : "(id) - (crontab_id)"
    django_celery_beat_solarschedule |o--o{ django_celery_beat_periodictask : "(id) - (solar_id)"
    django_celery_beat_clockedschedule |o--o{ django_celery_beat_periodictask : "(id) - (clocked_id)"
    auth_user ||--o| authtoken_token : "(id) - (user_id)"
    logging_storederrorlog ||--o{ logging_errorlogpartition : "(id) - (error_log_id)"
    django_content_type ||--o{ logging_errorlogpartition : "(id) - (content_type_id)"
    logging_errorlogpartition ||--o{ logging_errorlogpartitionentry : "(id) - (error_log_partition_id)"
    django_content_type ||--o{ acls_accesscontrollist : "(id) - (content_type_id)"
    permissions_role ||--o{ acls_accesscontrollist : "(id) - (role_id)"
    acls_accesscontrollist }o--o{ permissions_storedpermission : "(permissions) - (acls)"
    acls_accesscontrollist ||..o{ actstream_action : "(id) - (actor_object_id)"
    acls_accesscontrollist ||..o{ actstream_action : "(id) - (target_object_id)"
    acls_accesscontrollist ||..o{ actstream_action : "(id) - (action_object_object_id)"
    auth_user ||--o| user_management_useroptions : "(id) - (user_id)"
    auth_user |o--o{ authentication_attempts_loginattempt : "(id) - (user_id)"
    authentication_attempts_loginattempt ||..o{ actstream_action : "(id) - (actor_object_id)"
    authentication_attempts_loginattempt ||..o{ actstream_action : "(id) - (target_object_id)"
    authentication_attempts_loginattempt ||..o{ actstream_action : "(id) - (action_object_object_id)"
    auth_user ||--o| authentication_otp_userotpdata : "(id) - (user_id)"
    auth_user |o--o{ autoadmin_autoadminsingleton : "(id) - (account_id)"
    converter_asset ||..o{ actstream_action : "(id) - (actor_object_id)"
    converter_asset ||..o{ actstream_action : "(id) - (target_object_id)"
    converter_asset ||..o{ actstream_action : "(id) - (action_object_object_id)"
    converter_asset ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    django_content_type ||--o{ converter_objectlayer : "(id) - (content_type_id)"
    converter_storedlayer ||--o{ converter_objectlayer : "(id) - (stored_layer_id)"
    converter_objectlayer ||--o{ converter_layertransformation : "(id) - (object_layer_id)"
    credentials_storedcredential ||..o{ actstream_action : "(id) - (actor_object_id)"
    credentials_storedcredential ||..o{ actstream_action : "(id) - (target_object_id)"
    credentials_storedcredential ||..o{ actstream_action : "(id) - (action_object_object_id)"
    credentials_storedcredential ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    dashboards_storeddashboard ||..o{ actstream_action : "(id) - (actor_object_id)"
    dashboards_storeddashboard ||..o{ actstream_action : "(id) - (target_object_id)"
    dashboards_storeddashboard ||..o{ actstream_action : "(id) - (action_object_object_id)"
    dashboards_storeddashboard ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    django_gpg_key ||..o{ actstream_action : "(id) - (actor_object_id)"
    django_gpg_key ||..o{ actstream_action : "(id) - (target_object_id)"
    django_gpg_key ||..o{ actstream_action : "(id) - (action_object_object_id)"
    django_gpg_key ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user ||--o{ dynamic_search_savedresultset : "(id) - (user_id)"
    dynamic_search_savedresultset ||..o{ actstream_action : "(id) - (actor_object_id)"
    dynamic_search_savedresultset ||..o{ actstream_action : "(id) - (target_object_id)"
    dynamic_search_savedresultset ||..o{ actstream_action : "(id) - (action_object_object_id)"
    dynamic_search_savedresultset ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    dynamic_search_savedresultset ||--o{ dynamic_search_savedresultsetentry : "(id) - (saved_resultset_id)"
    file_caching_cache ||..o{ actstream_action : "(id) - (actor_object_id)"
    file_caching_cache ||..o{ actstream_action : "(id) - (target_object_id)"
    file_caching_cache ||..o{ actstream_action : "(id) - (action_object_object_id)"
    file_caching_cache ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    file_caching_cache ||--o{ file_caching_cachepartition : "(id) - (cache_id)"
    file_caching_cachepartition ||..o{ actstream_action : "(id) - (actor_object_id)"
    file_caching_cachepartition ||..o{ actstream_action : "(id) - (target_object_id)"
    file_caching_cachepartition ||..o{ actstream_action : "(id) - (action_object_object_id)"
    file_caching_cachepartition ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    file_caching_cachepartition ||--o{ file_caching_cachepartitionfile : "(id) - (partition_id)"
    auth_user ||--o| locales_userlocaleprofile : "(id) - (user_id)"
    django_content_type |o--o{ messaging_message : "(id) - (sender_content_type_id)"
    auth_user ||--o{ messaging_message : "(id) - (user_id)"
    messaging_message ||..o{ actstream_action : "(id) - (actor_object_id)"
    messaging_message ||..o{ actstream_action : "(id) - (target_object_id)"
    messaging_message ||..o{ actstream_action : "(id) - (action_object_object_id)"
    messaging_message ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    permissions_role }o--o{ permissions_storedpermission : "(permissions) - (roles)"
    permissions_role }o--o{ auth_group : "(groups) - (roles)"
    permissions_role ||..o{ actstream_action : "(id) - (actor_object_id)"
    permissions_role ||..o{ actstream_action : "(id) - (target_object_id)"
    permissions_role ||..o{ actstream_action : "(id) - (action_object_object_id)"
    permissions_role ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    permissions_storedpermission ||..o{ actstream_action : "(id) - (actor_object_id)"
    permissions_storedpermission ||..o{ actstream_action : "(id) - (target_object_id)"
    permissions_storedpermission ||..o{ actstream_action : "(id) - (action_object_object_id)"
    quotas_quota ||..o{ actstream_action : "(id) - (actor_object_id)"
    quotas_quota ||..o{ actstream_action : "(id) - (target_object_id)"
    quotas_quota ||..o{ actstream_action : "(id) - (action_object_object_id)"
    quotas_quota ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user ||--o{ server_side_events_streamevent : "(id) - (user_id)"
    auth_user ||--o{ storage_downloadfile : "(id) - (user_id)"
    storage_downloadfile ||..o{ actstream_action : "(id) - (actor_object_id)"
    storage_downloadfile ||..o{ actstream_action : "(id) - (target_object_id)"
    storage_downloadfile ||..o{ actstream_action : "(id) - (action_object_object_id)"
    storage_downloadfile ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user ||--o{ views_userconfirmview : "(id) - (user_id)"
    auth_user ||--o{ views_userviewmode : "(id) - (user_id)"
    announcements_announcement ||..o{ actstream_action : "(id) - (actor_object_id)"
    announcements_announcement ||..o{ actstream_action : "(id) - (target_object_id)"
    announcements_announcement ||..o{ actstream_action : "(id) - (action_object_object_id)"
    announcements_announcement ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documenttype ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documenttype ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documenttype ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documenttype ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documenttype ||--o{ documents_documenttypefilename : "(id) - (document_type_id)"
    documents_documenttypefilename ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documenttypefilename ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documenttypefilename ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documenttypefilename ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentfile |o--o| documents_document : "(id) - (file_latest_id)"
    documents_documenttype ||--o{ documents_document : "(id) - (document_type_id)"
    documents_documentversion |o--o| documents_document : "(id) - (version_active_id)"
    documents_document ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_document ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_document ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_document ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_document ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    documents_document ||--o{ documents_documentfile : "(id) - (document_id)"
    documents_documentfile ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documentfile ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documentfile ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documentfile ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentfile ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    documents_document ||--o{ documents_documentversion : "(id) - (document_id)"
    documents_documentversion ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documentversion ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documentversion ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documentversion ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentversion ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    documents_documentversion ||--o{ documents_documentversionpage : "(id) - (document_version_id)"
    django_content_type ||--o{ documents_documentversionpage : "(id) - (content_type_id)"
    documents_documentversionpage ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documentversionpage ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documentversionpage ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documentversionpage ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentversionpage ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    documents_documentfile ||--o{ documents_documentfilepage : "(id) - (document_file_id)"
    documents_documentfilepage ||..o{ actstream_action : "(id) - (actor_object_id)"
    documents_documentfilepage ||..o{ actstream_action : "(id) - (target_object_id)"
    documents_documentfilepage ||..o{ actstream_action : "(id) - (action_object_object_id)"
    documents_documentfilepage ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentfilepage ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    auth_user ||--o{ documents_recentlyaccesseddocument : "(id) - (user_id)"
    documents_document ||--o{ documents_recentlyaccesseddocument : "(id) - (document_id)"
    cabinets_cabinet |o--o{ cabinets_cabinet : "(id) - (parent_id)"
    cabinets_cabinet }o--o{ documents_document : "(documents) - (cabinets)"
    cabinets_cabinet ||..o{ actstream_action : "(id) - (actor_object_id)"
    cabinets_cabinet ||..o{ actstream_action : "(id) - (target_object_id)"
    cabinets_cabinet ||..o{ actstream_action : "(id) - (action_object_object_id)"
    cabinets_cabinet ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_document ||--o| checkouts_documentcheckout : "(id) - (document_id)"
    auth_user ||--o{ checkouts_documentcheckout : "(id) - (user_id)"
    documents_document ||--o{ document_comments_comment : "(id) - (document_id)"
    auth_user ||--o{ document_comments_comment : "(id) - (user_id)"
    document_comments_comment ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_comments_comment ||..o{ actstream_action : "(id) - (target_object_id)"
    document_comments_comment ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_comments_comment ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    auth_user ||--o{ documents_favoritedocument : "(id) - (user_id)"
    documents_document ||--o{ documents_favoritedocument : "(id) - (document_id)"
    document_indexing_indextemplate }o--o{ documents_documenttype : "(document_types) - (index_templates)"
    document_indexing_indextemplate ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_indexing_indextemplate ||..o{ actstream_action : "(id) - (target_object_id)"
    document_indexing_indextemplate ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_indexing_indextemplate ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_indexing_indextemplate ||--o{ document_indexing_indextemplateeventtrigger : "(id) - (index_template_id)"
    events_storedeventtype ||--o{ document_indexing_indextemplateeventtrigger : "(id) - (stored_event_type_id)"
    document_indexing_indextemplatenode |o--o{ document_indexing_indextemplatenode : "(id) - (parent_id)"
    document_indexing_indextemplate ||--o{ document_indexing_indextemplatenode : "(id) - (index_id)"
    document_indexing_indexinstancenode |o--o{ document_indexing_indexinstancenode : "(id) - (parent_id)"
    document_indexing_indextemplatenode ||--o{ document_indexing_indexinstancenode : "(id) - (index_template_node_id)"
    document_indexing_indexinstancenode }o--o{ documents_document : "(documents) - (index_instance_nodes)"
    documents_documentfilepage ||--o| document_parsing_documentfilepagecontent : "(id) - (document_file_page_id)"
    documents_documenttype ||--o| document_parsing_documenttypesettings : "(id) - (document_type_id)"
    documents_documentfile ||--o{ document_signatures_signaturebasemodel : "(id) - (document_file_id)"
    document_signatures_signaturebasemodel ||--o| document_signatures_detachedsignature : "is a, by signaturebasemodel_ptr_id"
    document_signatures_detachedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (actor_object_id)"
    document_signatures_detachedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (target_object_id)"
    document_signatures_detachedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (action_object_object_id)"
    document_signatures_detachedsignature ||..o{ acls_accesscontrollist : "(signaturebasemodel_ptr_id) - (object_id)"
    document_signatures_signaturebasemodel ||--o| document_signatures_embeddedsignature : "is a, by signaturebasemodel_ptr_id"
    document_signatures_embeddedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (actor_object_id)"
    document_signatures_embeddedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (target_object_id)"
    document_signatures_embeddedsignature ||..o{ actstream_action : "(signaturebasemodel_ptr_id) - (action_object_object_id)"
    document_signatures_embeddedsignature ||..o{ acls_accesscontrollist : "(signaturebasemodel_ptr_id) - (object_id)"
    document_states_workflow }o--o{ documents_documenttype : "(document_types) - (workflows)"
    document_states_workflow ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflow ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflow ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflow ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflow ||--o{ document_states_workflowstate : "(id) - (workflow_id)"
    document_states_workflowstate ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowstate ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowstate ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowstate ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflow ||--o{ document_states_workflowtransition : "(id) - (workflow_id)"
    document_states_workflowstate ||--o{ document_states_workflowtransition : "(id) - (origin_state_id)"
    document_states_workflowstate ||--o{ document_states_workflowtransition : "(id) - (destination_state_id)"
    document_states_workflowtransition ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowtransition ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowtransition ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowtransition ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflowtransition ||--o{ document_states_workflowtransitionfield : "(id) - (transition_id)"
    document_states_workflowtransitionfield ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowtransitionfield ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowtransitionfield ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowtransitionfield ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflow ||--o{ document_states_workflowinstance : "(id) - (workflow_id)"
    documents_document ||--o{ document_states_workflowinstance : "(id) - (document_id)"
    document_states_workflowstate ||--o{ document_states_workflowinstance : "(id) - (state_active_id)"
    document_states_workflowinstance ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowinstance ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowinstance ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowinstance ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflowinstance ||--o{ document_states_workflowinstancelogentry : "(id) - (workflow_instance_id)"
    document_states_workflowtransition |o--o{ document_states_workflowinstancelogentry : "(id) - (transition_id)"
    auth_user |o--o{ document_states_workflowinstancelogentry : "(id) - (user_id)"
    document_states_workflowstate ||--o{ document_states_workflowstateaction : "(id) - (state_id)"
    document_states_workflowstateaction ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowstateaction ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowstateaction ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowstateaction ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflowstateaction ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    document_states_workflowstate ||--o{ document_states_workflowstateescalation : "(id) - (state_id)"
    document_states_workflowtransition ||--o{ document_states_workflowstateescalation : "(id) - (transition_id)"
    document_states_workflowstateescalation ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowstateescalation ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowstateescalation ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowstateescalation ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    document_states_workflowtransition ||--o{ document_states_workflowtransitiontriggerevent : "(id) - (transition_id)"
    events_storedeventtype ||--o{ document_states_workflowtransitiontriggerevent : "(id) - (event_type_id)"
    document_states_workflowtransitiontriggerevent ||..o{ actstream_action : "(id) - (actor_object_id)"
    document_states_workflowtransitiontriggerevent ||..o{ actstream_action : "(id) - (target_object_id)"
    document_states_workflowtransitiontriggerevent ||..o{ actstream_action : "(id) - (action_object_object_id)"
    document_states_workflowtransitiontriggerevent ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    duplicates_storedduplicatebackend ||--o{ duplicates_duplicatebackendentry : "(id) - (stored_backend_id)"
    documents_document ||--o{ duplicates_duplicatebackendentry : "(id) - (document_id)"
    duplicates_duplicatebackendentry }o--o{ documents_document : "(documents) - (as_duplicate)"
    file_metadata_storeddriver ||--o{ file_metadata_documentfiledriverentry : "(id) - (driver_id)"
    documents_documentfile ||--o{ file_metadata_documentfiledriverentry : "(id) - (document_file_id)"
    documents_documenttype ||--o{ file_metadata_documenttypedriverconfiguration : "(id) - (document_type_id)"
    file_metadata_storeddriver ||--o{ file_metadata_documenttypedriverconfiguration : "(id) - (stored_driver_id)"
    file_metadata_documentfiledriverentry ||--o{ file_metadata_filemetadataentry : "(id) - (document_file_driver_entry_id)"
    file_metadata_storeddriver ||..o{ actstream_action : "(id) - (actor_object_id)"
    file_metadata_storeddriver ||..o{ actstream_action : "(id) - (target_object_id)"
    file_metadata_storeddriver ||..o{ actstream_action : "(id) - (action_object_object_id)"
    file_metadata_storeddriver ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    linking_smartlink }o--o{ documents_documenttype : "(document_types) - (smart_links)"
    linking_smartlink ||..o{ actstream_action : "(id) - (actor_object_id)"
    linking_smartlink ||..o{ actstream_action : "(id) - (target_object_id)"
    linking_smartlink ||..o{ actstream_action : "(id) - (action_object_object_id)"
    linking_smartlink ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    linking_smartlink ||--o{ linking_smartlinkcondition : "(id) - (smart_link_id)"
    linking_smartlinkcondition ||..o{ actstream_action : "(id) - (actor_object_id)"
    linking_smartlinkcondition ||..o{ actstream_action : "(id) - (target_object_id)"
    linking_smartlinkcondition ||..o{ actstream_action : "(id) - (action_object_object_id)"
    linking_smartlinkcondition ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    mailer_usermailer ||..o{ actstream_action : "(id) - (actor_object_id)"
    mailer_usermailer ||..o{ actstream_action : "(id) - (target_object_id)"
    mailer_usermailer ||..o{ actstream_action : "(id) - (action_object_object_id)"
    mailer_usermailer ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    mailer_usermailer ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    metadata_metadatatype ||..o{ actstream_action : "(id) - (actor_object_id)"
    metadata_metadatatype ||..o{ actstream_action : "(id) - (target_object_id)"
    metadata_metadatatype ||..o{ actstream_action : "(id) - (action_object_object_id)"
    metadata_metadatatype ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documenttype ||--o{ metadata_documenttypemetadatatype : "(id) - (document_type_id)"
    metadata_metadatatype ||--o{ metadata_documenttypemetadatatype : "(id) - (metadata_type_id)"
    metadata_documenttypemetadatatype ||..o{ actstream_action : "(id) - (actor_object_id)"
    metadata_documenttypemetadatatype ||..o{ actstream_action : "(id) - (target_object_id)"
    metadata_documenttypemetadatatype ||..o{ actstream_action : "(id) - (action_object_object_id)"
    metadata_documenttypemetadatatype ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_document ||--o{ metadata_documentmetadata : "(id) - (document_id)"
    metadata_metadatatype ||--o{ metadata_documentmetadata : "(id) - (metadata_type_id)"
    documents_documenttype ||--o| ocr_documenttypeocrsettings : "(id) - (document_type_id)"
    documents_documentversionpage ||--o| ocr_documentversionpageocrcontent : "(id) - (document_version_page_id)"
    sequences_sequence }o--o{ documents_documenttype : "(document_types) - (sequences)"
    sequences_sequence ||..o{ actstream_action : "(id) - (actor_object_id)"
    sequences_sequence ||..o{ actstream_action : "(id) - (target_object_id)"
    sequences_sequence ||..o{ actstream_action : "(id) - (action_object_object_id)"
    sequences_sequence ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_document ||--o{ signature_captures_signaturecapture : "(id) - (document_id)"
    auth_user ||--o{ signature_captures_signaturecapture : "(id) - (user_id)"
    signature_captures_signaturecapture ||..o{ actstream_action : "(id) - (actor_object_id)"
    signature_captures_signaturecapture ||..o{ actstream_action : "(id) - (target_object_id)"
    signature_captures_signaturecapture ||..o{ actstream_action : "(id) - (action_object_object_id)"
    signature_captures_signaturecapture ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    documents_documentfile ||--o{ sources_documentfilesourcemetadata : "(id) - (document_file_id)"
    sources_source ||..o{ actstream_action : "(id) - (actor_object_id)"
    sources_source ||..o{ actstream_action : "(id) - (target_object_id)"
    sources_source ||..o{ actstream_action : "(id) - (action_object_object_id)"
    sources_source ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    sources_source ||..o{ logging_errorlogpartition : "(id) - (object_id)"
    tags_tag }o--o{ documents_document : "(documents) - (tags)"
    tags_tag ||..o{ actstream_action : "(id) - (actor_object_id)"
    tags_tag ||..o{ actstream_action : "(id) - (target_object_id)"
    tags_tag ||..o{ actstream_action : "(id) - (action_object_object_id)"
    tags_tag ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    web_links_weblink }o--o{ documents_documenttype : "(document_types) - (web_links)"
    web_links_weblink ||..o{ actstream_action : "(id) - (actor_object_id)"
    web_links_weblink ||..o{ actstream_action : "(id) - (target_object_id)"
    web_links_weblink ||..o{ actstream_action : "(id) - (action_object_object_id)"
    web_links_weblink ||..o{ acls_accesscontrollist : "(id) - (object_id)"
    events_storedeventtype {
        AutoField id PK
        CharField name UK
    }
    events_eventsubscription {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey stored_event_type_id FK
    }
    events_notification {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey action_id FK
        BooleanField read
    }
    events_objecteventsubscription {
        AutoField id PK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        ForeignKey user_id FK
        ForeignKey stored_event_type_id FK
    }
    appearance_usersetting {
        AutoField id PK
        ForeignKey user_id FK
        CharField namespace "Axis or frontend that owns this setting."
        CharField key
        TextField value
    }
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
    auth_user {
        AutoField id PK
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField username UK "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        CharField first_name "The given name of the user."
        CharField last_name "The family name of the user."
        EmailField email
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
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
    actstream_follow {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey content_type_id FK
        CharField object_id
        BooleanField actor_only
        CharField flag
        DateTimeField started
    }
    actstream_action {
        AutoField id PK
        ForeignKey actor_content_type_id FK
        CharField actor_object_id
        CharField verb
        TextField description
        ForeignKey target_content_type_id FK
        CharField target_object_id
        ForeignKey action_object_content_type_id FK
        CharField action_object_object_id
        DateTimeField timestamp
        BooleanField public
    }
    axes_accessfailurelog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        BooleanField locked_out
    }
    axes_accessattempt {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        TextField get_data
        TextField post_data
        PositiveIntegerField failures_since_start
    }
    axes_accessattemptexpiration {
        OneToOneField access_attempt_id PK
        DateTimeField expires_at "The time when access attempt expires and is no longer valid."
    }
    axes_accesslog {
        AutoField id PK
        CharField user_agent
        GenericIPAddressField ip_address
        CharField username
        CharField http_accept
        CharField path_info
        DateTimeField attempt_time
        DateTimeField logout_time
        CharField session_hash
    }
    django_celery_beat_solarschedule {
        AutoField id PK
        SolarScheduleEvent event "The type of solar event when the job should run"
        DecimalField latitude "Run the task when the event happens at this latitude"
        DecimalField longitude "Run the task when the event happens at this longitude"
    }
    django_celery_beat_intervalschedule {
        AutoField id PK
        IntegerField every "Number of interval periods to wait before running the task again"
        IntervalSchedulePeriod period "The type of period between task runs (Example: days)"
    }
    django_celery_beat_clockedschedule {
        AutoField id PK
        DateTimeField clocked_time "Run the task at clocked time"
    }
    django_celery_beat_crontabschedule {
        AutoField id PK
        CharField minute "Cron Minutes to Run. Use '*' for 'all'. (Example: '0,30')"
        CharField hour "Cron Hours to Run. Use '*' for 'all'. (Example: '8,20')"
        CharField day_of_month "Cron Days Of The Month to Run. Use '*' for 'all'. (Example: '1,15')"
        CharField month_of_year "Cron Months (1-12) Of The Year to Run. Use '*' for 'all'. (Example: '1,12')"
        CharField day_of_week "Cron Days Of The Week to Run. Use '*' for 'all', Sunday is 0 or 7, Monday is 1. (Example: '0,5')"
        CrontabScheduleTimezone timezone "Timezone to Run the Cron Schedule on. Default is UTC."
    }
    django_celery_beat_periodictasks {
        SmallIntegerField ident PK
        DateTimeField last_update
    }
    django_celery_beat_periodictask {
        AutoField id PK
        CharField name UK "Short Description For This Task"
        CharField task "The Name of the Celery Task that Should be Run.  (Example: 'proj.tasks.import_contacts')"
        ForeignKey interval_id FK "Interval Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey crontab_id FK "Crontab Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey solar_id FK "Solar Schedule to run the task on.  Set only one schedule type, leave the others null."
        ForeignKey clocked_id FK "Clocked Schedule to run the task on.  Set only one schedule type, leave the others null."
        TextField args "JSON encoded positional arguments (Example: ['arg1', 'arg2'])"
        TextField kwargs "JSON encoded keyword arguments (Example: {'argument': 'value'})"
        CharField queue "Queue defined in CELERY_TASK_QUEUES. Leave None for default queuing."
        CharField exchange "Override Exchange for low-level AMQP routing"
        CharField routing_key "Override Routing Key for low-level AMQP routing"
        TextField headers "JSON encoded message headers for the AMQP message."
        PositiveIntegerField priority "Priority Number between 0 and 255. Supported by: RabbitMQ, Redis (priority reversed, 0 is highest)."
        DateTimeField expires "Datetime after which the schedule will no longer trigger the task to run"
        PositiveIntegerField expire_seconds "Timedelta with seconds which the schedule will no longer trigger the task to run"
        BooleanField one_off "If True, the schedule will only run the task a single time"
        DateTimeField start_time "Datetime when the schedule should begin triggering the task to run"
        BooleanField enabled "Set to False to disable the schedule"
        DateTimeField last_run_at "Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False."
        PositiveIntegerField total_run_count "Running count of how many times the schedule has triggered the task"
        DateTimeField date_changed "Datetime that this PeriodicTask was last modified"
        TextField description "Detailed description about the details of this Periodic Task"
    }
    authtoken_token {
        CharField key PK
        OneToOneField user_id FK, UK
        DateTimeField created
    }
    logging_storederrorlog {
        AutoField id PK
        CharField name UK
    }
    logging_errorlogpartition {
        AutoField id PK
        ForeignKey error_log_id FK
        CharField name
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
    }
    logging_errorlogpartitionentry {
        AutoField id PK
        ForeignKey error_log_partition_id FK
        DateTimeField datetime
        CharField domain_name
        TextField text
    }
    task_manager_taskdeduplicationentry {
        AutoField id PK
        CharField dotted_path "Python path of the task type the marker suppresses."
        TextField key "Serialized keyword arguments identifying the unit of work requested."
        CharField key_hash "Hash of the key. The unique constraint is applied to this field and not to the key itself because the length of a key is decided by the task requesting the work, and an index over an unbounded column exceeds the size limit of some database managers."
        DateTimeField datetime_created "Date and time the work was requested or last re-dispatched. A marker whose work has not begun stops suppressing requests once this is older than the queued interval, which is what allows a worker that ended before publishing its task to recover without intervention."
        DateTimeField datetime_started "Date and time a worker began performing the work. Until it is set, the task the marker stands for is still waiting in the queue and has not read the object it will act on, so a request arriving in that period asks for work the queued task is already going to perform."
        BooleanField is_dirty "A new request for the same unit of work arrived after the task began performing it, therefore too late for the task to take it into account. The work is dispatched once more when the task completes."
    }
    acls_accesscontrollist {
        AutoField id PK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        ForeignKey role_id FK "Role to which the access is granted for the specified object."
    }
    user_management_useroptions {
        AutoField id PK
        OneToOneField user_id FK, UK
        BooleanField block_password_change
    }
    authentication_attempts_loginattempt {
        AutoField id PK
        CharField username "The username exactly as it was supplied for the authentication attempt. For failed attempts this may not match any existing account."
        ForeignKey user_id FK "The existing account that the supplied username resolved to, if any."
        LoginAttemptResult result "Whether the authentication attempt succeeded or failed."
        LoginAttemptSource source "The subsystem through which the authentication was attempted, such as the interactive session or the API."
        GenericIPAddressField ip_address "The client address as resolved by the lockout layer. Behind a reverse proxy this is only meaningful when the proxy trust is configured; treat it as best effort."
        TextField user_agent "Information about the web browser or application used to make the login attempt."
        DateTimeField datetime
    }
    authentication_otp_userotpdata {
        AutoField id PK
        OneToOneField user_id FK, UK
        CharField secret
    }
    autoadmin_autoadminsingleton {
        AutoField id PK
        ForeignKey account_id FK
        CharField password
        CharField password_hash
    }
    converter_asset {
        AutoField id PK
        CharField label UK
        CharField internal_name UK "This value will be used when referencing this asset. Can only contain letters, numbers, and underscores."
        FileField file
        CharField file_hash "A hash/checksum generated from the asset file. Used to invalidate cached images that use the asset when the asset file changes."
    }
    converter_storedlayer {
        AutoField id PK
        CharField name UK
        PositiveIntegerField order UK
    }
    converter_objectlayer {
        AutoField id PK
        ForeignKey content_type_id FK
        PositiveIntegerField object_id
        BooleanField enabled
        ForeignKey stored_layer_id FK
    }
    converter_layertransformation {
        AutoField id PK
        ForeignKey object_layer_id FK
        PositiveIntegerField order "Order in which the transformations will be executed. If left unchanged, an automatic order value will be assigned."
        CharField name
        TextField arguments "Enter the arguments for the transformation as a YAML dictionary. ie: {'degrees': 180}"
        BooleanField enabled
    }
    credentials_storedcredential {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
        CharField label UK "Short description of this credential."
        CharField internal_name UK "This value will be used by other apps to reference this credential. Can only contain letters, numbers, and underscores."
    }
    dashboards_storeddashboard {
        AutoField id PK
        CharField name UK
    }
    django_gpg_key {
        AutoField id PK
        TextField key_data "ASCII armored version of the key."
        DateTimeField creation_date
        DateTimeField expiration_date
        CharField fingerprint UK
        PositiveIntegerField length
        PositiveIntegerField algorithm
        TextField user_id
        KeyKeyType key_type
    }
    dynamic_search_savedresultset {
        AutoField id PK
        PositiveIntegerField time_to_live "Time to keep the resultset in seconds. This value is increased every time the resultset is accessed."
        DateTimeField timestamp "The server date and time when the resultset was created."
        PositiveBigIntegerField epoch
        CharField app_label
        CharField model_name
        TextField search_query_text
        TextField search_explainer_text
        PositiveIntegerField result_count "Number of results stored in the resultset."
        ForeignKey user_id FK "User for which the resultset was created."
    }
    dynamic_search_savedresultsetentry {
        AutoField id PK
        ForeignKey saved_resultset_id FK
        BigIntegerField object_id
    }
    file_caching_cache {
        AutoField id PK
        CharField defined_storage_name UK "Internal name of the defined storage for this cache."
        PositiveBigIntegerField maximum_size "Maximum size of the cache in bytes."
        PositiveBigIntegerField maximum_size_old "Previous maximum size of the cache in bytes."
    }
    file_caching_cachepartition {
        AutoField id PK
        ForeignKey cache_id FK
        CharField name
    }
    file_caching_cachepartitionfile {
        AutoField id PK
        ForeignKey partition_id FK
        DateTimeField datetime
        DateTimeField accessed "Last date and time this cache partition file was read. Used to evict the least recently used files first."
        CharField filename
        PositiveIntegerField file_size
        PositiveIntegerField hits "Times this cache partition file has been accessed."
    }
    locales_userlocaleprofile {
        AutoField id PK
        OneToOneField user_id FK, UK
        UserLocaleProfileTimezone timezone
        UserLocaleProfileLanguage language
    }
    lock_manager_lock {
        AutoField id PK
        DateTimeField creation_datetime
        IntegerField timeout
        CharField name UK
    }
    messaging_message {
        AutoField id PK
        ForeignKey sender_content_type_id FK
        PositiveIntegerField sender_object_id
        ForeignKey user_id FK
        CharField subject "Short description of this message."
        TextField body "The actual content of the message."
        BooleanField read "This field determines if the message has been read or not."
        DateTimeField date_time "Date and time of the message creation."
    }
    permissions_role {
        AutoField id PK
        CharField label UK "A short text describing the role."
    }
    permissions_storedpermission {
        AutoField id PK
        CharField namespace
        CharField name
    }
    quotas_quota {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data
        BooleanField enabled "Allow quick disable or enable of the quota."
    }
    server_side_events_streamevent {
        AutoField id PK
        ForeignKey user_id FK
        CharField event_type
        TextField payload
        DateTimeField datetime_created
    }
    smart_settings_updatedstoredsetting {
        AutoField id PK
        CharField key UK "A unique identifier."
        TextField value
    }
    storage_downloadfile {
        AutoField id PK
        CharField filename
        DateTimeField datetime
        FileField file
        CharField label
        ForeignKey user_id FK
    }
    storage_shareduploadedfile {
        AutoField id PK
        DateTimeField datetime
        FileField file
        CharField filename
    }
    views_userconfirmview {
        AutoField id PK
        CharField namespace
        CharField name "Full name of the view including the namespace."
        ForeignKey user_id FK
        BooleanField remember "Remember the last confirmation of the view."
    }
    views_userviewmode {
        AutoField id PK
        CharField namespace
        CharField name "Full name of the view including the namespace."
        ForeignKey user_id FK
        CharField value "Stored value used to identify the display mode of the view."
    }
    announcements_announcement {
        AutoField id PK
        CharField label "Short description of this announcement."
        TextField text "The actual text to be displayed."
        BooleanField enabled
        DateTimeField start_datetime "Date and time after which this announcement will be displayed."
        DateTimeField end_datetime "Date and time until when this announcement is to be displayed."
    }
    documents_documenttype {
        AutoField id PK
        CharField label UK "A short text identifying the document type."
        CharField filename_generator_backend "The class responsible for producing the actual filename used to store the uploaded documents."
        TextField filename_generator_backend_arguments "The arguments for the filename generator backend as a YAML dictionary."
        PositiveIntegerField trash_time_period "Amount of time after which documents of this type will be moved to the trash."
        DocumentTypeTrashTimeUnit trash_time_unit
        PositiveIntegerField delete_time_period "Amount of time after which documents of this type in the trash will be deleted."
        DocumentTypeDeleteTimeUnit delete_time_unit
        BooleanField document_stub_pruning_enabled "Delete documents that do not contain any files after a configured expiration interval."
        PositiveBigIntegerField document_stub_expiration_interval "Time (in seconds) after which a document stub will be considered invalid and deleted, if pruning is enabled. This an optimization setting and should only be changed for specific circumstances."
    }
    documents_documenttypefilename {
        AutoField id PK
        ForeignKey document_type_id FK
        CharField filename
        BooleanField enabled
    }
    documents_document {
        AutoField id PK
        OneToOneField file_latest_id FK, UK
        UUIDField uuid "UUID of a document, universally Unique ID. An unique identifier generated for each document."
        ForeignKey document_type_id FK "The document type of the document."
        CharField label "A short text identifying the document. By default, will be set to the filename of the first file uploaded to the document."
        TextField description "An optional short text describing a document."
        DateTimeField datetime_created "The date and time of the document creation."
        CharField language "The primary language in the document."
        BooleanField in_trash "Whether or not this document is in the trash."
        DateTimeField trashed_date_time "The server date and time when the document was moved to the trash."
        BooleanField is_stub "A document stub is a document with an entry on the database but no file uploaded. This could be an interrupted upload or a deferred upload via the API."
        OneToOneField version_active_id FK, UK
    }
    documents_documentfile {
        AutoField id PK
        ForeignKey document_id FK
        DateTimeField timestamp "The server date and time when the document file was processed."
        TextField comment "An optional short text describing the document file."
        FileField file
        CharField filename
        CharField mimetype "The document file's file mimetype. MIME types are a standard way to describe the format of a file, in this case the file format of the document. Some examples: 'text/plain' or 'image/jpeg'."
        CharField encoding "The document file file encoding. binary 7-bit, binary 8-bit, text, base64, etc."
        CharField checksum "A hash/checkdigit/fingerprint generated from the document's binary data. Only identical documents will have the same checksum."
        PositiveBigIntegerField size "The size of the file in bytes."
    }
    documents_documentversion {
        AutoField id PK
        ForeignKey document_id FK
        DateTimeField timestamp "The server date and time when the document version was created."
        TextField comment "An optional short text describing the document version."
        BooleanField active "Determines the active version of the document."
    }
    documents_documentversionpage {
        AutoField id PK
        ForeignKey document_version_id FK
        PositiveIntegerField page_number "Unique integer number for the page. Pages are ordered by this number."
        ForeignKey content_type_id FK "Content type for the source object of the page."
        PositiveIntegerField object_id "ID for the source object of the page."
    }
    documents_documentfilepage {
        AutoField id PK
        ForeignKey document_file_id FK
        PositiveIntegerField page_number
    }
    documents_recentlyaccesseddocument {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey document_id FK
        DateTimeField datetime_accessed
    }
    cabinets_cabinet {
        AutoField id PK
        TreeForeignKey parent_id FK
        CharField label "A short text used to identify the cabinet."
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    checkouts_documentcheckout {
        AutoField id PK
        OneToOneField document_id FK, UK
        DateTimeField checkout_datetime
        DateTimeField expiration_datetime "Amount of time to hold the document checked out in minutes."
        ForeignKey user_id FK
        BooleanField block_new_file "Do not allow new file of this document to be uploaded."
    }
    document_comments_comment {
        AutoField id PK
        ForeignKey document_id FK
        ForeignKey user_id FK "The user account that made the comment."
        TextField text "Actual text content of the comment."
        DateTimeField submit_date
    }
    documents_favoritedocument {
        AutoField id PK
        ForeignKey user_id FK
        ForeignKey document_id FK
        DateTimeField datetime_added
    }
    document_indexing_indextemplate {
        AutoField id PK
        CharField label UK "Short description of this index."
        SlugField slug UK "This value will be used by other apps to reference this index."
        BooleanField enabled "Causes this index to be visible and updated when document data changes."
    }
    document_indexing_indextemplateeventtrigger {
        AutoField id PK
        ForeignKey index_template_id FK
        ForeignKey stored_event_type_id FK
    }
    document_indexing_indextemplatenode {
        AutoField id PK
        TreeForeignKey parent_id FK "Parent index template node of this node."
        ForeignKey index_id FK
        TextField expression "Enter a template to render. Use Django's default templating language."
        BooleanField enabled "Causes this node to be visible and updated when document data changes."
        BooleanField link_documents "Check this option to have this node act as a container for documents and not as a parent for further nodes."
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    document_indexing_indexinstancenode {
        AutoField id PK
        TreeForeignKey parent_id FK
        ForeignKey index_template_node_id FK
        CharField value
        PositiveIntegerField lft
        PositiveIntegerField rght
        PositiveIntegerField tree_id
        PositiveIntegerField level
    }
    document_parsing_documentfilepagecontent {
        AutoField id PK
        OneToOneField document_file_page_id FK, UK
        TextField content "The actual text content as extracted by the document parsing backend."
    }
    document_parsing_documenttypesettings {
        AutoField id PK
        OneToOneField document_type_id FK, UK
        BooleanField auto_parsing "Automatically queue newly created documents for parsing."
    }
    document_signatures_signaturebasemodel {
        AutoField id PK
        ForeignKey document_file_id FK
        DateTimeField date_time
        CharField key_id "ID of the key that will be used to sign the document."
        CharField signature_id
        CharField public_key_fingerprint
    }
    document_signatures_detachedsignature {
        OneToOneField signaturebasemodel_ptr_id PK
        FileField signature_file "Signature file previously generated."
    }
    document_signatures_embeddedsignature {
        OneToOneField signaturebasemodel_ptr_id PK
    }
    document_states_workflow {
        AutoField id PK
        BooleanField auto_launch "Launch workflow when document is created."
        CharField internal_name UK "This value will be used by other apps to reference this workflow. Can only contain letters, numbers, and underscores."
        CharField label UK "A short text to describe the workflow."
        BooleanField ignore_completed "Ignore workflow instances if they are in their final state."
    }
    document_states_workflowstate {
        AutoField id PK
        ForeignKey workflow_id FK
        CharField label "A short text to describe the workflow state."
        BooleanField initial "The state at which the workflow will start in. Only one state can be the initial state."
        BooleanField final "The state at which the workflow will stop. Only one state can be the final state."
        IntegerField completion "The percent of completion that this state represents in relation to the workflow. Use numbers without the percent sign."
    }
    document_states_workflowtransition {
        AutoField id PK
        TextField condition "The condition that will determine if this transition is enabled or not. The condition is evaluated against the workflow instance. Conditions that do not return any value, that return the Python logical None, or an empty string ('') are considered to be logical false, any other value is considered to be the logical true."
        ForeignKey workflow_id FK
        CharField label "A short text to describe the transition."
        ForeignKey origin_state_id FK
        ForeignKey destination_state_id FK
    }
    document_states_workflowtransitionfield {
        AutoField id PK
        ForeignKey transition_id FK
        WorkflowTransitionFieldFieldType field_type
        CharField name "The name that will be used to identify this field in other parts of the workflow system."
        CharField label "The field name that will be shown on the user interface."
        TextField default "Optional default value for the field. Can be a template."
        TextField lookup "Enter a template to render. Must result in a comma delimited string."
        TextField help_text "An optional message that will help users better understand the purpose of the field and data to provide."
        BooleanField required "Whether this field needs to be filled out or not to proceed."
        WorkflowTransitionFieldWidget widget "An optional class to change the default presentation of the field."
        TextField widget_kwargs "A group of keyword arguments to customize the widget. Use YAML format."
    }
    document_states_workflowinstance {
        AutoField id PK
        ForeignKey workflow_id FK "The workflow template."
        DateTimeField datetime "Workflow instance creation date time."
        ForeignKey document_id FK "The document to which the workflow instance is attached."
        TextField context
        ForeignKey state_active_id FK "The currently active state of the workflow instance."
    }
    document_states_workflowinstancelogentry {
        AutoField id PK
        ForeignKey workflow_instance_id FK
        DateTimeField datetime
        ForeignKey transition_id FK
        ForeignKey user_id FK
        TextField comment
        TextField extra_data
    }
    document_states_workflowstateaction {
        AutoField id PK
        TextField condition "The condition that will determine if this state action is executed or not. The condition is evaluated against the workflow instance. Conditions that do not return any value, that return the Python logical None, or an empty string ('') are considered to be logical false, any other value is considered to be the logical true."
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
        ForeignKey state_id FK
        CharField label "A short text describing the action. Actions are execute by alphabetical order."
        BooleanField enabled
        WorkflowStateActionWhen when "At which moment of the state this action will execute."
    }
    document_states_workflowstateescalation {
        AutoField id PK
        TextField condition "The condition that will determine if this object is executed or not. Conditions that do not return any value, that return the Python logical None, or an empty string ('') are considered to be logical false, any other value is considered to be the logical true."
        ForeignKey state_id FK
        ForeignKey transition_id FK
        IntegerField priority "Determine the order in which the escalations will be evaluated. Escalations with a lower priority number are executed before escalations with a higher priority number."
        BooleanField enabled "Enable automatic transition the workflow after a specified amount of time has elapsed in the state without change."
        WorkflowStateEscalationUnit unit
        PositiveIntegerField amount "Amount of the selected escalation units of time."
        TextField comment "Comment to save to the workflow instance when the escalation is executed."
    }
    document_states_workflowtransitiontriggerevent {
        AutoField id PK
        ForeignKey transition_id FK
        ForeignKey event_type_id FK
    }
    duplicates_storedduplicatebackend {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
    }
    duplicates_duplicatebackendentry {
        AutoField id PK
        ForeignKey stored_backend_id FK
        ForeignKey document_id FK
    }
    file_metadata_documentfiledriverentry {
        AutoField id PK
        ForeignKey driver_id FK
        ForeignKey document_file_id FK
    }
    file_metadata_documenttypedriverconfiguration {
        AutoField id PK
        TextField arguments "Enter the arguments for the drive for the specific document type as a YAML dictionary. ie: {'degrees': 180}"
        ForeignKey document_type_id FK
        ForeignKey stored_driver_id FK
        BooleanField enabled "Enable this driver to process document files of the selected document type."
    }
    file_metadata_filemetadataentry {
        AutoField id PK
        ForeignKey document_file_driver_entry_id FK
        CharField internal_name "Normalized name of the file metadata entry."
        CharField key "Name of the file metadata entry as provided by the driver."
        TextField value "Value of the file metadata entry."
    }
    file_metadata_storeddriver {
        AutoField id PK
        CharField driver_path UK
        CharField internal_name UK
        BooleanField exists "The class defined by this instance is valid and active."
    }
    linking_smartlink {
        AutoField id PK
        CharField label "A short text describing the smart link."
        CharField dynamic_label "Use this field to show a unique label depending on the document from which the smart link is being accessed."
        BooleanField enabled
    }
    linking_smartlinkcondition {
        AutoField id PK
        ForeignKey smart_link_id FK
        SmartLinkConditionInclusion inclusion "The inclusion is ignored for the first item."
        CharField foreign_document_data "This represents the metadata of all other documents."
        SmartLinkConditionOperator operator
        TextField expression "The expression using document properties to be evaluated against the foreign document field."
        BooleanField negated "Inverts the logic of the operator."
        BooleanField enabled
    }
    mailer_usermailer {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
        CharField label UK "A short text describing the mailing profile."
        BooleanField default "If default, this mailing profile will be pre-selected on the document mailing form."
        BooleanField enabled
    }
    mayan_statistics_statisticresult {
        AutoField id PK
        SlugField slug UK
        DateTimeField datetime
        TextField serialize_data
    }
    metadata_metadatatype {
        AutoField id PK
        CharField name UK "Name used by other apps to reference this metadata type. Do not use python reserved words, or spaces."
        CharField label "Short description of this metadata type."
        CharField default "Enter a template to render."
        TextField lookup "Enter a template to render. Must result in a comma delimited string."
        CharField validation "The validator will reject data entry if the value entered does not conform to the expected format."
        TextField validation_arguments "Enter the arguments for the validator in YAML format."
        CharField parser "The parser will reformat the value entered to conform to the expected format."
        TextField parser_arguments "Enter the arguments for the parser in YAML format."
    }
    metadata_documenttypemetadatatype {
        AutoField id PK
        ForeignKey document_type_id FK
        ForeignKey metadata_type_id FK
        BooleanField required
    }
    metadata_documentmetadata {
        AutoField id PK
        ForeignKey document_id FK
        ForeignKey metadata_type_id FK
        TextField value "The actual value stored in the metadata type field for the document."
    }
    ocr_documenttypeocrsettings {
        AutoField id PK
        OneToOneField document_type_id FK, UK
        BooleanField auto_ocr "Automatically queue newly created documents for OCR."
    }
    ocr_documentversionpageocrcontent {
        AutoField id PK
        OneToOneField document_version_page_id FK, UK
        TextField content "The actual text content extracted by the OCR backend."
    }
    sequences_sequence {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
        CharField label UK "Short description of this sequence."
        CharField internal_name UK "This value will be used by other apps to reference this sequence. Can only contain letters, numbers, and underscores."
        BigIntegerField position "Zero based counter of how many values have been consumed. Set this to continue the numbering of a system being replaced."
        BooleanField value_is_unique "Whether a value of this sequence identifies a single position. Recorded when the sequence is saved and not entered by hand."
        SequenceOnLimit on_limit "What to do when the sequence reaches the last position its backend is able to represent."
    }
    signature_captures_signaturecapture {
        AutoField id PK
        ForeignKey document_id FK
        TextField data "Data representing the handwritten signature."
        TextField svg "Vector representation of the handwritten signature."
        CharField text "Print version of the captured signature."
        ForeignKey user_id FK
        DateTimeField date_time_created
        DateTimeField date_time_edited
        CharField internal_name "This value will be used when referencing this signature capture in relationship to the document. Can only contain letters, numbers, and underscores."
    }
    sources_documentfilesourcemetadata {
        AutoField id PK
        ForeignKey document_file_id FK
        CharField key "Name of the source metadata entry."
        TextField value "The actual value stored in the source metadata for the document file."
    }
    sources_source {
        AutoField id PK
        CharField backend_path "The dotted Python path to the backend class."
        TextField backend_data "JSON encoded data for the backend class."
        CharField label UK "A short text to describe this source."
        BooleanField enabled
    }
    tags_tag {
        AutoField id PK
        CharField label UK "A short text used as the tag name."
        CharField color "The RGB color values for the tag."
    }
    web_links_weblink {
        AutoField id PK
        CharField label UK "A short text describing the web link."
        TextField template "Template that will be used to craft the final URL of the web link."
        BooleanField enabled
    }
```
