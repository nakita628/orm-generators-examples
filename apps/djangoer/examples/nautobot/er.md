```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user ||--o{ social_auth_usersocialauth : "(id) - (user_id)"
    taggit_tag ||--o{ taggit_taggeditem : "(id) - (tag_id)"
    django_content_type ||--o{ taggit_taggeditem : "(id) - (content_type_id)"
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_celery_beat_intervalschedule |o--o{ django_celery_beat_periodictask : "(id) - (interval_id)"
    django_celery_beat_crontabschedule |o--o{ django_celery_beat_periodictask : "(id) - (crontab_id)"
    django_celery_beat_solarschedule |o--o{ django_celery_beat_periodictask : "(id) - (solar_id)"
    django_celery_beat_clockedschedule |o--o{ django_celery_beat_periodictask : "(id) - (clocked_id)"
    circuits_provider ||--o{ circuits_providernetwork : "(id) - (provider_id)"
    circuits_providernetwork ||..o{ extras_taggeditem : "(id) - (object_id)"
    circuits_providernetwork ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    circuits_providernetwork ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    circuits_providernetwork ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    circuits_providernetwork ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    circuits_providernetwork ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    circuits_providernetwork ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    circuits_provider ||..o{ extras_taggeditem : "(id) - (object_id)"
    circuits_provider ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    circuits_provider ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    circuits_provider ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    circuits_provider ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    circuits_provider ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    circuits_provider ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    circuits_circuittype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    circuits_circuittype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    circuits_circuittype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    circuits_circuittype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    circuits_circuittype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    circuits_circuittype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ circuits_circuit : "(id) - (status_id)"
    circuits_provider ||--o{ circuits_circuit : "(id) - (provider_id)"
    circuits_circuittype ||--o{ circuits_circuit : "(id) - (circuit_type_id)"
    tenancy_tenant |o--o{ circuits_circuit : "(id) - (tenant_id)"
    circuits_circuittermination |o--o{ circuits_circuit : "(id) - (circuit_termination_a_id)"
    circuits_circuittermination |o--o{ circuits_circuit : "(id) - (circuit_termination_z_id)"
    circuits_circuit ||..o{ extras_taggeditem : "(id) - (object_id)"
    circuits_circuit ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    circuits_circuit ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    circuits_circuit ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    circuits_circuit ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    circuits_circuit ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    circuits_circuit ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    circuits_circuit ||--o{ circuits_circuittermination : "(id) - (circuit_id)"
    dcim_location |o--o{ circuits_circuittermination : "(id) - (location_id)"
    circuits_providernetwork |o--o{ circuits_circuittermination : "(id) - (provider_network_id)"
    cloud_cloudnetwork |o--o{ circuits_circuittermination : "(id) - (cloud_network_id)"
    circuits_circuittermination ||..o{ extras_taggeditem : "(id) - (object_id)"
    circuits_circuittermination ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    circuits_circuittermination ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    circuits_circuittermination ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    circuits_circuittermination ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    circuits_circuittermination ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    circuits_circuittermination ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    circuits_circuittermination ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_manufacturer ||--o{ cloud_cloudaccount : "(id) - (provider_id)"
    extras_secretsgroup |o--o{ cloud_cloudaccount : "(id) - (secrets_group_id)"
    cloud_cloudaccount ||..o{ extras_taggeditem : "(id) - (object_id)"
    cloud_cloudaccount ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    cloud_cloudaccount ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    cloud_cloudaccount ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    cloud_cloudaccount ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    cloud_cloudaccount ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudaccount ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer ||--o{ cloud_cloudresourcetype : "(id) - (provider_id)"
    cloud_cloudresourcetype }o--o{ django_content_type : "(content_types) - (cloud_resource_types)"
    cloud_cloudresourcetype ||..o{ extras_taggeditem : "(id) - (object_id)"
    cloud_cloudresourcetype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    cloud_cloudresourcetype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    cloud_cloudresourcetype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    cloud_cloudresourcetype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    cloud_cloudresourcetype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudresourcetype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    cloud_cloudresourcetype ||--o{ cloud_cloudnetwork : "(id) - (cloud_resource_type_id)"
    cloud_cloudaccount ||--o{ cloud_cloudnetwork : "(id) - (cloud_account_id)"
    cloud_cloudnetwork |o--o{ cloud_cloudnetwork : "(id) - (parent_id)"
    cloud_cloudnetwork ||..o{ extras_taggeditem : "(id) - (object_id)"
    cloud_cloudnetwork ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    cloud_cloudnetwork ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    cloud_cloudnetwork ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    cloud_cloudnetwork ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    cloud_cloudnetwork ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudnetwork ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    cloud_cloudnetwork ||--o{ cloud_cloudnetworkprefixassignment : "(id) - (cloud_network_id)"
    ipam_prefix ||--o{ cloud_cloudnetworkprefixassignment : "(id) - (prefix_id)"
    cloud_cloudnetworkprefixassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudnetworkprefixassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    cloud_cloudresourcetype ||--o{ cloud_cloudservice : "(id) - (cloud_resource_type_id)"
    cloud_cloudaccount |o--o{ cloud_cloudservice : "(id) - (cloud_account_id)"
    cloud_cloudservice ||..o{ extras_taggeditem : "(id) - (object_id)"
    cloud_cloudservice ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    cloud_cloudservice ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    cloud_cloudservice ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    cloud_cloudservice ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    cloud_cloudservice ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudservice ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    cloud_cloudnetwork ||--o{ cloud_cloudservicenetworkassignment : "(id) - (cloud_network_id)"
    cloud_cloudservice ||--o{ cloud_cloudservicenetworkassignment : "(id) - (cloud_service_id)"
    cloud_cloudservicenetworkassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    cloud_cloudservicenetworkassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ data_validation_regexrule : "(id) - (content_type_id)"
    data_validation_regexrule ||..o{ extras_taggeditem : "(id) - (object_id)"
    data_validation_regexrule ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    data_validation_regexrule ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    data_validation_regexrule ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    data_validation_regexrule ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    data_validation_regexrule ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    data_validation_regexrule ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ data_validation_minmaxrule : "(id) - (content_type_id)"
    data_validation_minmaxrule ||..o{ extras_taggeditem : "(id) - (object_id)"
    data_validation_minmaxrule ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    data_validation_minmaxrule ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    data_validation_minmaxrule ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    data_validation_minmaxrule ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    data_validation_minmaxrule ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    data_validation_minmaxrule ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ data_validation_requiredrule : "(id) - (content_type_id)"
    data_validation_requiredrule ||..o{ extras_taggeditem : "(id) - (object_id)"
    data_validation_requiredrule ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    data_validation_requiredrule ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    data_validation_requiredrule ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    data_validation_requiredrule ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    data_validation_requiredrule ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    data_validation_requiredrule ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ data_validation_uniquerule : "(id) - (content_type_id)"
    data_validation_uniquerule ||..o{ extras_taggeditem : "(id) - (object_id)"
    data_validation_uniquerule ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    data_validation_uniquerule ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    data_validation_uniquerule ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    data_validation_uniquerule ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    data_validation_uniquerule ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    data_validation_uniquerule ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ data_validation_datacompliance : "(id) - (content_type_id)"
    data_validation_datacompliance ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    data_validation_datacompliance ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    data_validation_datacompliance ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ dcim_consoleport : "(id) - (device_id)"
    dcim_module |o--o{ dcim_consoleport : "(id) - (module_id)"
    dcim_consoleport ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_consoleport ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_consoleport ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_consoleport ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_consoleport ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_consoleport ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_consoleport ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_consoleport ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_device |o--o{ dcim_consoleserverport : "(id) - (device_id)"
    dcim_module |o--o{ dcim_consoleserverport : "(id) - (module_id)"
    dcim_consoleserverport ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_consoleserverport ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_consoleserverport ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_consoleserverport ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_consoleserverport ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_consoleserverport ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_consoleserverport ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_consoleserverport ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_device |o--o{ dcim_powerport : "(id) - (device_id)"
    dcim_module |o--o{ dcim_powerport : "(id) - (module_id)"
    dcim_powerport ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_powerport ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_powerport ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_powerport ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_powerport ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_powerport ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_powerport ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_powerport ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_device |o--o{ dcim_poweroutlet : "(id) - (device_id)"
    dcim_module |o--o{ dcim_poweroutlet : "(id) - (module_id)"
    dcim_powerport |o--o{ dcim_poweroutlet : "(id) - (power_port_id)"
    dcim_poweroutlet ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_poweroutlet ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_poweroutlet ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_poweroutlet ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_poweroutlet ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_poweroutlet ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_poweroutlet ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_poweroutlet ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_device |o--o{ dcim_interface : "(id) - (device_id)"
    dcim_module |o--o{ dcim_interface : "(id) - (module_id)"
    extras_status ||--o{ dcim_interface : "(id) - (status_id)"
    extras_role |o--o{ dcim_interface : "(id) - (role_id)"
    dcim_interface |o--o{ dcim_interface : "(id) - (parent_interface_id)"
    dcim_interface |o--o{ dcim_interface : "(id) - (bridge_id)"
    dcim_interface |o--o{ dcim_interface : "(id) - (lag_id)"
    ipam_vlan |o--o{ dcim_interface : "(id) - (untagged_vlan_id)"
    ipam_vrf |o--o{ dcim_interface : "(id) - (vrf_id)"
    dcim_interface }o--o{ ipam_vlan : "(tagged_vlans) - (interfaces_as_tagged)"
    dcim_interface ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_interface ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_interface ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_interface ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_interface ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_interface ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_interface ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_interface ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_interface ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_interface ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_status ||--o{ dcim_interfaceredundancygroup : "(id) - (status_id)"
    extras_secretsgroup |o--o{ dcim_interfaceredundancygroup : "(id) - (secrets_group_id)"
    ipam_ipaddress |o--o{ dcim_interfaceredundancygroup : "(id) - (virtual_ip_id)"
    dcim_interfaceredundancygroup ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_interfaceredundancygroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_interfaceredundancygroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_interfaceredundancygroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_interfaceredundancygroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_interfaceredundancygroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_interfaceredundancygroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_interfaceredundancygroup ||--o{ dcim_interfaceredundancygroupassociation : "(id) - (interface_redundancy_group_id)"
    dcim_interface ||--o{ dcim_interfaceredundancygroupassociation : "(id) - (interface_id)"
    dcim_interfaceredundancygroupassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_interfaceredundancygroupassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ dcim_frontport : "(id) - (device_id)"
    dcim_module |o--o{ dcim_frontport : "(id) - (module_id)"
    dcim_rearport ||--o{ dcim_frontport : "(id) - (rear_port_id)"
    dcim_frontport ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_frontport ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_frontport ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_frontport ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_frontport ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_frontport ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_frontport ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ dcim_rearport : "(id) - (device_id)"
    dcim_module |o--o{ dcim_rearport : "(id) - (module_id)"
    dcim_rearport ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_rearport ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_rearport ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_rearport ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_rearport ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_rearport ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_rearport ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device ||--o{ dcim_devicebay : "(id) - (device_id)"
    dcim_device |o--o| dcim_devicebay : "(id) - (installed_device_id)"
    dcim_devicebay ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_devicebay ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_devicebay ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_devicebay ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_devicebay ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_devicebay ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_devicebay ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_inventoryitem |o--o{ dcim_inventoryitem : "(id) - (parent_id)"
    dcim_device ||--o{ dcim_inventoryitem : "(id) - (device_id)"
    dcim_manufacturer |o--o{ dcim_inventoryitem : "(id) - (manufacturer_id)"
    dcim_softwareversion |o--o{ dcim_inventoryitem : "(id) - (software_version_id)"
    dcim_inventoryitem }o--o{ dcim_softwareimagefile : "(software_image_files) - (inventory_items)"
    dcim_inventoryitem ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_inventoryitem ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_inventoryitem ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_inventoryitem ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_inventoryitem ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_inventoryitem ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_inventoryitem ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ dcim_modulebay : "(id) - (parent_device_id)"
    dcim_module |o--o{ dcim_modulebay : "(id) - (parent_module_id)"
    dcim_modulefamily |o--o{ dcim_modulebay : "(id) - (module_family_id)"
    dcim_modulebay ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_modulebay ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_modulebay ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_modulebay ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_modulebay ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_modulebay ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_modulebay ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer |o--o{ dcim_cabletype : "(id) - (manufacturer_id)"
    dcim_cabletype ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_cabletype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_cabletype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_cabletype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_cabletype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_cabletype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_cabletype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_cabletype |o--o{ dcim_cable : "(id) - (cable_type_id)"
    extras_status ||--o{ dcim_cable : "(id) - (status_id)"
    dcim_cable ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_cable ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_cable ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_cable ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_cable ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_cable ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_cable ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_cable ||--o{ dcim_cabletocabletermination : "(id) - (cable_id)"
    circuits_circuittermination |o--o| dcim_cabletocabletermination : "(id) - (circuit_termination_id)"
    dcim_consoleport |o--o| dcim_cabletocabletermination : "(id) - (console_port_id)"
    dcim_consoleserverport |o--o| dcim_cabletocabletermination : "(id) - (console_server_port_id)"
    dcim_frontport |o--o| dcim_cabletocabletermination : "(id) - (front_port_id)"
    dcim_interface |o--o| dcim_cabletocabletermination : "(id) - (interface_id)"
    dcim_powerfeed |o--o| dcim_cabletocabletermination : "(id) - (power_feed_id)"
    dcim_poweroutlet |o--o| dcim_cabletocabletermination : "(id) - (power_outlet_id)"
    dcim_powerport |o--o| dcim_cabletocabletermination : "(id) - (power_port_id)"
    dcim_rearport |o--o| dcim_cabletocabletermination : "(id) - (rear_port_id)"
    dcim_cabletocabletermination ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_cabletocabletermination ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ dcim_cablepath : "(id) - (origin_type_id)"
    django_content_type |o--o{ dcim_cablepath : "(id) - (destination_type_id)"
    dcim_cablepath ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_cablepath ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_consoleporttemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_consoleporttemplate : "(id) - (module_type_id)"
    dcim_consoleporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_consoleporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_consoleporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_consoleporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_consoleporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_consoleporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_consoleserverporttemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_consoleserverporttemplate : "(id) - (module_type_id)"
    dcim_consoleserverporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_consoleserverporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_consoleserverporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_consoleserverporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_consoleserverporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_consoleserverporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_powerporttemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_powerporttemplate : "(id) - (module_type_id)"
    dcim_powerporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_powerporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_powerporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_powerporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_powerporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_powerporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_poweroutlettemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_poweroutlettemplate : "(id) - (module_type_id)"
    dcim_powerporttemplate |o--o{ dcim_poweroutlettemplate : "(id) - (power_port_template_id)"
    dcim_poweroutlettemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_poweroutlettemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_poweroutlettemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_poweroutlettemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_poweroutlettemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_poweroutlettemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_interfacetemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_interfacetemplate : "(id) - (module_type_id)"
    dcim_interfacetemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_interfacetemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_interfacetemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_interfacetemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_interfacetemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_interfacetemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_frontporttemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_frontporttemplate : "(id) - (module_type_id)"
    dcim_rearporttemplate ||--o{ dcim_frontporttemplate : "(id) - (rear_port_template_id)"
    dcim_frontporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_frontporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_frontporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_frontporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_frontporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_frontporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_rearporttemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_rearporttemplate : "(id) - (module_type_id)"
    dcim_rearporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_rearporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_rearporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_rearporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_rearporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_rearporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype ||--o{ dcim_devicebaytemplate : "(id) - (device_type_id)"
    dcim_devicebaytemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_devicebaytemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_devicebaytemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_devicebaytemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_devicebaytemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_devicebaytemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype |o--o{ dcim_modulebaytemplate : "(id) - (device_type_id)"
    dcim_moduletype |o--o{ dcim_modulebaytemplate : "(id) - (module_type_id)"
    dcim_modulefamily |o--o{ dcim_modulebaytemplate : "(id) - (module_family_id)"
    dcim_modulebaytemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_modulebaytemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_modulebaytemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_modulebaytemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_modulebaytemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_modulebaytemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_manufacturer ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_manufacturer ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_manufacturer ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_manufacturer ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_manufacturer ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicefamily ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_devicefamily ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_devicefamily ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_devicefamily ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_devicefamily ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_devicefamily ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_devicefamily ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_devicetype ||--o{ dcim_devicetypetosoftwareimagefile : "(id) - (device_type_id)"
    dcim_softwareimagefile ||--o{ dcim_devicetypetosoftwareimagefile : "(id) - (software_image_file_id)"
    dcim_devicetypetosoftwareimagefile ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_devicetypetosoftwareimagefile ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer ||--o{ dcim_devicetype : "(id) - (manufacturer_id)"
    dcim_devicefamily |o--o{ dcim_devicetype : "(id) - (device_family_id)"
    dcim_devicetype ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_devicetype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_devicetype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_devicetype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_devicetype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_devicetype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_devicetype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer |o--o{ dcim_platform : "(id) - (manufacturer_id)"
    dcim_platform ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_platform ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_platform ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_platform ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_platform ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_platform ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_configcontextschema |o--o{ dcim_device : "(id) - (local_config_context_schema_id)"
    django_content_type |o--o{ dcim_device : "(id) - (local_config_context_data_owner_content_type_id)"
    dcim_devicetype ||--o{ dcim_device : "(id) - (device_type_id)"
    extras_status ||--o{ dcim_device : "(id) - (status_id)"
    extras_role ||--o{ dcim_device : "(id) - (role_id)"
    tenancy_tenant |o--o{ dcim_device : "(id) - (tenant_id)"
    dcim_platform |o--o{ dcim_device : "(id) - (platform_id)"
    dcim_location ||--o{ dcim_device : "(id) - (location_id)"
    dcim_rack |o--o{ dcim_device : "(id) - (rack_id)"
    ipam_ipaddress |o--o{ dcim_device : "(id) - (primary_ip4_id)"
    ipam_ipaddress |o--o{ dcim_device : "(id) - (primary_ip6_id)"
    dcim_virtualchassis |o--o{ dcim_device : "(id) - (virtual_chassis_id)"
    dcim_deviceredundancygroup |o--o{ dcim_device : "(id) - (device_redundancy_group_id)"
    dcim_softwareversion |o--o{ dcim_device : "(id) - (software_version_id)"
    extras_secretsgroup |o--o{ dcim_device : "(id) - (secrets_group_id)"
    dcim_controllermanageddevicegroup |o--o{ dcim_device : "(id) - (controller_managed_device_group_id)"
    dcim_device }o--o{ dcim_softwareimagefile : "(software_image_files) - (devices)"
    dcim_device ||..o{ extras_imageattachment : "(id) - (object_id)"
    dcim_device ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_device ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_device ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_device ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_device ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_device ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_device ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device ||--o{ dcim_deviceclusterassignment : "(id) - (device_id)"
    virtualization_cluster ||--o{ dcim_deviceclusterassignment : "(id) - (cluster_id)"
    dcim_deviceclusterassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_deviceclusterassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o| dcim_virtualchassis : "(id) - (master_id)"
    dcim_virtualchassis ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_virtualchassis ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_virtualchassis ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_virtualchassis ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_virtualchassis ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_virtualchassis ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_virtualchassis ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ dcim_deviceredundancygroup : "(id) - (status_id)"
    extras_secretsgroup |o--o{ dcim_deviceredundancygroup : "(id) - (secrets_group_id)"
    dcim_deviceredundancygroup ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_deviceredundancygroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_deviceredundancygroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_deviceredundancygroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_deviceredundancygroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_deviceredundancygroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_deviceredundancygroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_softwareversion ||--o{ dcim_softwareimagefile : "(id) - (software_version_id)"
    extras_externalintegration |o--o{ dcim_softwareimagefile : "(id) - (external_integration_id)"
    extras_status ||--o{ dcim_softwareimagefile : "(id) - (status_id)"
    dcim_softwareimagefile ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_softwareimagefile ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_softwareimagefile ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_softwareimagefile ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_softwareimagefile ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_softwareimagefile ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_softwareimagefile ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_platform ||--o{ dcim_softwareversion : "(id) - (platform_id)"
    extras_status ||--o{ dcim_softwareversion : "(id) - (status_id)"
    dcim_softwareversion ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_softwareversion ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_softwareversion ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_softwareversion ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_softwareversion ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_softwareversion ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_softwareversion ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ dcim_controller : "(id) - (status_id)"
    dcim_location ||--o{ dcim_controller : "(id) - (location_id)"
    dcim_platform |o--o{ dcim_controller : "(id) - (platform_id)"
    extras_role |o--o{ dcim_controller : "(id) - (role_id)"
    tenancy_tenant |o--o{ dcim_controller : "(id) - (tenant_id)"
    extras_externalintegration |o--o{ dcim_controller : "(id) - (external_integration_id)"
    dcim_device |o--o{ dcim_controller : "(id) - (controller_device_id)"
    dcim_deviceredundancygroup |o--o{ dcim_controller : "(id) - (controller_device_redundancy_group_id)"
    dcim_controller ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_controller ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_controller ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_controller ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_controller ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_controller ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_controller ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_controllermanageddevicegroup |o--o{ dcim_controllermanageddevicegroup : "(id) - (parent_id)"
    dcim_controller ||--o{ dcim_controllermanageddevicegroup : "(id) - (controller_id)"
    tenancy_tenant |o--o{ dcim_controllermanageddevicegroup : "(id) - (tenant_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_controllermanageddevicegroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_controllermanageddevicegroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_modulefamily ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_modulefamily ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_modulefamily ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_modulefamily ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_modulefamily ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_modulefamily ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_modulefamily ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_manufacturer ||--o{ dcim_moduletype : "(id) - (manufacturer_id)"
    dcim_modulefamily |o--o{ dcim_moduletype : "(id) - (module_family_id)"
    dcim_moduletype ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_moduletype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_moduletype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_moduletype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_moduletype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_moduletype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_moduletype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_moduletype ||--o{ dcim_module : "(id) - (module_type_id)"
    dcim_modulebay |o--o| dcim_module : "(id) - (parent_module_bay_id)"
    extras_status ||--o{ dcim_module : "(id) - (status_id)"
    extras_role |o--o{ dcim_module : "(id) - (role_id)"
    tenancy_tenant |o--o{ dcim_module : "(id) - (tenant_id)"
    dcim_location |o--o{ dcim_module : "(id) - (location_id)"
    dcim_module ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_module ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_module ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_module ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_module ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_module ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_module ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device ||--o{ dcim_virtualdevicecontext : "(id) - (device_id)"
    extras_status ||--o{ dcim_virtualdevicecontext : "(id) - (status_id)"
    extras_role |o--o{ dcim_virtualdevicecontext : "(id) - (role_id)"
    ipam_ipaddress |o--o{ dcim_virtualdevicecontext : "(id) - (primary_ip4_id)"
    ipam_ipaddress |o--o{ dcim_virtualdevicecontext : "(id) - (primary_ip6_id)"
    tenancy_tenant |o--o{ dcim_virtualdevicecontext : "(id) - (tenant_id)"
    dcim_controllermanageddevicegroup |o--o{ dcim_virtualdevicecontext : "(id) - (controller_managed_device_group_id)"
    dcim_virtualdevicecontext ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_virtualdevicecontext ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_virtualdevicecontext ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_virtualdevicecontext ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_virtualdevicecontext ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_virtualdevicecontext ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_virtualdevicecontext ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_virtualdevicecontext ||--o{ dcim_interfacevdcassignment : "(id) - (virtual_device_context_id)"
    dcim_interface ||--o{ dcim_interfacevdcassignment : "(id) - (interface_id)"
    dcim_interfacevdcassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_interfacevdcassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_locationtype |o--o{ dcim_locationtype : "(id) - (parent_id)"
    dcim_locationtype }o--o{ django_content_type : "(content_types) - (location_types)"
    dcim_locationtype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_locationtype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_locationtype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_locationtype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_locationtype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_locationtype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_location |o--o{ dcim_location : "(id) - (parent_id)"
    dcim_locationtype ||--o{ dcim_location : "(id) - (location_type_id)"
    extras_status ||--o{ dcim_location : "(id) - (status_id)"
    tenancy_tenant |o--o{ dcim_location : "(id) - (tenant_id)"
    dcim_location ||..o{ extras_imageattachment : "(id) - (object_id)"
    dcim_location ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_location ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_location ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_location ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_location ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_location ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_location ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_location ||--o{ dcim_powerpanel : "(id) - (location_id)"
    dcim_rackgroup |o--o{ dcim_powerpanel : "(id) - (rack_group_id)"
    dcim_powerpanel ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_powerpanel ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_powerpanel ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_powerpanel ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_powerpanel ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_powerpanel ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_powerpanel ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_powerpanel ||--o{ dcim_powerfeed : "(id) - (power_panel_id)"
    dcim_powerpanel |o--o{ dcim_powerfeed : "(id) - (destination_panel_id)"
    dcim_rack |o--o{ dcim_powerfeed : "(id) - (rack_id)"
    extras_status ||--o{ dcim_powerfeed : "(id) - (status_id)"
    dcim_powerfeed ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_powerfeed ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_powerfeed ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_powerfeed ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_powerfeed ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_powerfeed ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_powerfeed ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_powerfeed ||..o{ dcim_cablepath : "(id) - (origin_id)"
    dcim_rackgroup |o--o{ dcim_rackgroup : "(id) - (parent_id)"
    dcim_location ||--o{ dcim_rackgroup : "(id) - (location_id)"
    dcim_rackgroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_rackgroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_rackgroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_rackgroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_rackgroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_rackgroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ dcim_rack : "(id) - (status_id)"
    extras_role |o--o{ dcim_rack : "(id) - (role_id)"
    dcim_location ||--o{ dcim_rack : "(id) - (location_id)"
    dcim_rackgroup |o--o{ dcim_rack : "(id) - (rack_group_id)"
    tenancy_tenant |o--o{ dcim_rack : "(id) - (tenant_id)"
    dcim_rack ||..o{ extras_imageattachment : "(id) - (object_id)"
    dcim_rack ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_rack ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_rack ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_rack ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_rack ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_rack ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_rack ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_rack ||--o{ dcim_rackreservation : "(id) - (rack_id)"
    tenancy_tenant |o--o{ dcim_rackreservation : "(id) - (tenant_id)"
    auth_user ||--o{ dcim_rackreservation : "(id) - (user_id)"
    dcim_rackreservation ||..o{ extras_taggeditem : "(id) - (object_id)"
    dcim_rackreservation ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    dcim_rackreservation ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    dcim_rackreservation ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    dcim_rackreservation ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    dcim_rackreservation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    dcim_rackreservation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    auth_user |o--o{ extras_objectchange : "(id) - (user_id)"
    django_content_type |o--o{ extras_objectchange : "(id) - (changed_object_type_id)"
    django_content_type |o--o{ extras_objectchange : "(id) - (related_object_type_id)"
    extras_objectchange ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_objectchange ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_computedfield : "(id) - (content_type_id)"
    extras_computedfield ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_computedfield ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_computedfield ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_computedfield ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_customfield }o--o{ django_content_type : "(content_types) - (custom_fields)"
    extras_customfield ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_customfield ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_customfield ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_customfield ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_customfield ||--o{ extras_customfieldchoice : "(id) - (custom_field_id)"
    extras_customfieldchoice ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_customfieldchoice ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_relationship : "(id) - (source_type_id)"
    django_content_type ||--o{ extras_relationship : "(id) - (destination_type_id)"
    extras_relationship ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_relationship ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_relationship ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_relationship ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_relationship ||--o{ extras_relationshipassociation : "(id) - (relationship_id)"
    django_content_type ||--o{ extras_relationshipassociation : "(id) - (source_type_id)"
    django_content_type ||--o{ extras_relationshipassociation : "(id) - (destination_type_id)"
    extras_relationshipassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_relationshipassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_tag }o--o{ django_content_type : "(content_types) - (tags)"
    extras_tag ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_tag ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_tag ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_tag ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_tag ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_tag ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_taggeditem : "(id) - (content_type_id)"
    extras_tag ||--o{ extras_taggeditem : "(id) - (tag_id)"
    extras_taggeditem ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_taggeditem ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_approvalworkflowdefinition : "(id) - (model_content_type_id)"
    extras_approvalworkflowdefinition ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_approvalworkflowdefinition ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowdefinition ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowdefinition ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_approvalworkflowdefinition ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_approvalworkflowdefinition ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_approvalworkflowdefinition ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_approvalworkflowdefinition ||--o{ extras_approvalworkflowstagedefinition : "(id) - (approval_workflow_definition_id)"
    auth_group ||--o{ extras_approvalworkflowstagedefinition : "(id) - (approver_group_id)"
    extras_approvalworkflowstagedefinition ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowstagedefinition ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowstagedefinition ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_approvalworkflowstagedefinition ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_approvalworkflowstagedefinition ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_approvalworkflowstagedefinition ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_approvalworkflowdefinition |o--o{ extras_approvalworkflow : "(id) - (approval_workflow_definition_id)"
    django_content_type ||--o{ extras_approvalworkflow : "(id) - (object_under_review_content_type_id)"
    auth_user |o--o{ extras_approvalworkflow : "(id) - (user_id)"
    extras_approvalworkflow ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_approvalworkflow ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_approvalworkflow ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_approvalworkflow ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_approvalworkflow ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_approvalworkflow ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_approvalworkflow ||--o{ extras_approvalworkflowstage : "(id) - (approval_workflow_id)"
    extras_approvalworkflowstagedefinition |o--o{ extras_approvalworkflowstage : "(id) - (approval_workflow_stage_definition_id)"
    extras_approvalworkflowstage ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowstage ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_approvalworkflowstage ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_approvalworkflowstage ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_approvalworkflowstage ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_approvalworkflowstage ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_approvalworkflowstage ||--o{ extras_approvaluserresponse : "(id) - (approval_workflow_stage_id)"
    auth_user ||--o{ extras_approvaluserresponse : "(id) - (user_id)"
    extras_approvaluserresponse ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_approvaluserresponse ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_role }o--o{ django_content_type : "(content_types) - (roles)"
    extras_role ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_role ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_role ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_role ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_role ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_role ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status }o--o{ django_content_type : "(content_types) - (statuses)"
    extras_status ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_status ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_status ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_status ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_status ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_status ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_contact ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_contact ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_contact ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_contact ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_contact ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_contact ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_contact ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_team }o--o{ extras_contact : "(contacts) - (teams)"
    extras_team ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_team ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_team ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_team ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_team ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_team ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_team ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_contact |o--o{ extras_contactassociation : "(id) - (contact_id)"
    extras_team |o--o{ extras_contactassociation : "(id) - (team_id)"
    django_content_type |o--o{ extras_contactassociation : "(id) - (associated_object_type_id)"
    extras_role ||--o{ extras_contactassociation : "(id) - (role_id)"
    extras_status ||--o{ extras_contactassociation : "(id) - (status_id)"
    extras_contactassociation ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_contactassociation ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_contactassociation ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_contactassociation ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_contactassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_contactassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secretsgroup |o--o{ extras_gitrepository : "(id) - (secrets_group_id)"
    extras_gitrepository ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_gitrepository ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_gitrepository ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_gitrepository ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_gitrepository ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_gitrepository ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_gitrepository ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_dynamicgroup : "(id) - (content_type_id)"
    tenancy_tenant |o--o{ extras_dynamicgroup : "(id) - (tenant_id)"
    extras_dynamicgroup ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_dynamicgroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_dynamicgroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_dynamicgroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_dynamicgroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_dynamicgroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_dynamicgroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_dynamicgroup ||--o{ extras_dynamicgroupmembership : "(id) - (group_id)"
    extras_dynamicgroup ||--o{ extras_dynamicgroupmembership : "(id) - (parent_group_id)"
    extras_dynamicgroupmembership ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_dynamicgroupmembership ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_dynamicgroup ||--o{ extras_staticgroupassociation : "(id) - (dynamic_group_id)"
    django_content_type ||--o{ extras_staticgroupassociation : "(id) - (associated_object_type_id)"
    extras_staticgroupassociation ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_staticgroupassociation ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_staticgroupassociation ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_staticgroupassociation ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_staticgroupassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_staticgroupassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_jobqueue ||--o{ extras_job : "(id) - (default_job_queue_id)"
    extras_job ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_job ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_job ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_job ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_job ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_job ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_job ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_job ||--o{ extras_jobhook : "(id) - (job_id)"
    extras_jobhook }o--o{ django_content_type : "(content_types) - (job_hooks)"
    extras_jobhook ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_jobhook ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_jobhook ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_jobhook ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_jobhook ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobhook ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_jobresult ||--o{ extras_joblogentry : "(id) - (job_result_id)"
    extras_joblogentry ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_joblogentry ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ extras_jobqueue : "(id) - (tenant_id)"
    extras_jobqueue ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_jobqueue ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_jobqueue ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_jobqueue ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_jobqueue ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_jobqueue ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobqueue ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_job ||--o{ extras_jobqueueassignment : "(id) - (job_id)"
    extras_jobqueue ||--o{ extras_jobqueueassignment : "(id) - (job_queue_id)"
    extras_jobqueueassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobqueueassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_job |o--o{ extras_jobresult : "(id) - (job_model_id)"
    auth_user |o--o{ extras_jobresult : "(id) - (user_id)"
    extras_scheduledjob |o--o{ extras_jobresult : "(id) - (scheduled_job_id)"
    auth_user |o--o{ extras_jobresult : "(id) - (canceled_by_id)"
    extras_jobresult ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobresult ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_jobresult ||--o{ extras_jobconsoleentry : "(id) - (job_result_id)"
    extras_jobconsoleentry ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobconsoleentry ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_job ||--o{ extras_jobbutton : "(id) - (job_id)"
    extras_jobbutton }o--o{ django_content_type : "(content_types) - (job_buttons)"
    extras_jobbutton ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_jobbutton ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_jobbutton ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_jobbutton ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_celery_beat_solarschedule |o--o{ extras_scheduledjob : "(id) - (solar_id)"
    django_celery_beat_clockedschedule |o--o{ extras_scheduledjob : "(id) - (clocked_id)"
    extras_job |o--o{ extras_scheduledjob : "(id) - (job_model_id)"
    extras_jobqueue |o--o{ extras_scheduledjob : "(id) - (job_queue_id)"
    auth_user |o--o{ extras_scheduledjob : "(id) - (user_id)"
    extras_scheduledjob ||..o{ extras_approvalworkflow : "(id) - (object_under_review_object_id)"
    extras_scheduledjob ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_scheduledjob ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_metadatatype }o--o{ django_content_type : "(content_types) - (metadata_types)"
    extras_metadatatype ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_metadatatype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_metadatatype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_metadatatype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_metadatatype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_metadatatype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_metadatatype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_metadatatype ||--o{ extras_metadatachoice : "(id) - (metadata_type_id)"
    extras_metadatachoice ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_metadatachoice ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_metadatatype ||--o{ extras_objectmetadata : "(id) - (metadata_type_id)"
    extras_contact |o--o{ extras_objectmetadata : "(id) - (contact_id)"
    extras_team |o--o{ extras_objectmetadata : "(id) - (team_id)"
    django_content_type ||--o{ extras_objectmetadata : "(id) - (assigned_object_type_id)"
    extras_objectmetadata ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_objectmetadata ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type |o--o{ extras_configcontext : "(id) - (owner_content_type_id)"
    extras_configcontextschema |o--o{ extras_configcontext : "(id) - (config_context_schema_id)"
    extras_configcontext }o--o{ dcim_location : "(locations) - (_extras_configcontext_locations_+)"
    extras_configcontext }o--o{ extras_role : "(roles) - (_extras_configcontext_roles_+)"
    extras_configcontext }o--o{ dcim_devicetype : "(device_types) - (_extras_configcontext_device_types_+)"
    extras_configcontext }o--o{ dcim_deviceredundancygroup : "(device_redundancy_groups) - (_extras_configcontext_device_redundancy_groups_+)"
    extras_configcontext }o--o{ dcim_platform : "(platforms) - (_extras_configcontext_platforms_+)"
    extras_configcontext }o--o{ virtualization_clustergroup : "(cluster_groups) - (_extras_configcontext_cluster_groups_+)"
    extras_configcontext }o--o{ virtualization_cluster : "(clusters) - (_extras_configcontext_clusters_+)"
    extras_configcontext }o--o{ tenancy_tenantgroup : "(tenant_groups) - (_extras_configcontext_tenant_groups_+)"
    extras_configcontext }o--o{ tenancy_tenant : "(tenants) - (_extras_configcontext_tenants_+)"
    extras_configcontext }o--o{ extras_tag : "(tags) - (_extras_configcontext_tags_+)"
    extras_configcontext }o--o{ dcim_devicefamily : "(device_families) - (_extras_configcontext_device_families_+)"
    extras_configcontext }o--o{ extras_dynamicgroup : "(dynamic_groups) - (_extras_configcontext_dynamic_groups_+)"
    extras_configcontext ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_configcontext ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_configcontext ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type |o--o{ extras_configcontextschema : "(id) - (owner_content_type_id)"
    extras_configcontextschema ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_configcontextschema ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_configcontextschema ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_configcontextschema ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_configcontextschema ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_configcontextschema ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_customlink : "(id) - (content_type_id)"
    extras_customlink ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_customlink ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_customlink ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_customlink ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type |o--o{ extras_exporttemplate : "(id) - (owner_content_type_id)"
    django_content_type ||--o{ extras_exporttemplate : "(id) - (content_type_id)"
    extras_exporttemplate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_exporttemplate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_exporttemplate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_exporttemplate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_exporttemplate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_exporttemplate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secretsgroup |o--o{ extras_externalintegration : "(id) - (secrets_group_id)"
    extras_externalintegration ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_externalintegration ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_externalintegration ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_externalintegration ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_externalintegration ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_externalintegration ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_externalintegration ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_fileattachment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_fileattachment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_jobresult |o--o{ extras_fileproxy : "(id) - (job_result_id)"
    extras_fileproxy ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_fileproxy ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type |o--o{ extras_graphqlquery : "(id) - (owner_content_type_id)"
    extras_graphqlquery ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_graphqlquery ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_graphqlquery ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_graphqlquery ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_healthchecktestmodel ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_healthchecktestmodel ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_imageattachment : "(id) - (content_type_id)"
    extras_imageattachment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_imageattachment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    django_content_type ||--o{ extras_note : "(id) - (assigned_object_type_id)"
    auth_user |o--o{ extras_note : "(id) - (user_id)"
    extras_note ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_note ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    auth_user ||--o{ extras_savedview : "(id) - (owner_id)"
    extras_savedview ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_savedview ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_savedview ||--o{ extras_usersavedviewassociation : "(id) - (saved_view_id)"
    auth_user ||--o{ extras_usersavedviewassociation : "(id) - (user_id)"
    extras_usersavedviewassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_usersavedviewassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_webhook }o--o{ django_content_type : "(content_types) - (webhooks)"
    extras_webhook ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_webhook ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_webhook ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_webhook ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secret ||..o{ extras_taggeditem : "(id) - (object_id)"
    extras_secret ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_secret ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_secret ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_secret ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_secret ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_secret ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secretsgroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    extras_secretsgroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    extras_secretsgroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    extras_secretsgroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_secretsgroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_secretsgroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secretsgroup ||--o{ extras_secretsgroupassociation : "(id) - (secrets_group_id)"
    extras_secret ||--o{ extras_secretsgroupassociation : "(id) - (secret_id)"
    extras_secretsgroupassociation ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    extras_secretsgroupassociation ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_location |o--o{ ipam_namespace : "(id) - (location_id)"
    tenancy_tenant |o--o{ ipam_namespace : "(id) - (tenant_id)"
    ipam_namespace ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_namespace ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_namespace ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_namespace ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_namespace ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_namespace ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_namespace ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status |o--o{ ipam_vrf : "(id) - (status_id)"
    ipam_namespace ||--o{ ipam_vrf : "(id) - (namespace_id)"
    tenancy_tenant |o--o{ ipam_vrf : "(id) - (tenant_id)"
    ipam_vrf }o--o{ ipam_routetarget : "(import_targets) - (importing_vrfs)"
    ipam_vrf }o--o{ ipam_routetarget : "(export_targets) - (exporting_vrfs)"
    ipam_vrf ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_vrf ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_vrf ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_vrf ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_vrf ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_vrf ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vrf ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_vrf ||--o{ ipam_vrfdeviceassignment : "(id) - (vrf_id)"
    dcim_device |o--o{ ipam_vrfdeviceassignment : "(id) - (device_id)"
    virtualization_virtualmachine |o--o{ ipam_vrfdeviceassignment : "(id) - (virtual_machine_id)"
    dcim_virtualdevicecontext |o--o{ ipam_vrfdeviceassignment : "(id) - (virtual_device_context_id)"
    ipam_vrfdeviceassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vrfdeviceassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_vrf ||--o{ ipam_vrfprefixassignment : "(id) - (vrf_id)"
    ipam_prefix ||--o{ ipam_vrfprefixassignment : "(id) - (prefix_id)"
    ipam_vrfprefixassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vrfprefixassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ ipam_routetarget : "(id) - (tenant_id)"
    ipam_routetarget ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_routetarget ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_routetarget ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_routetarget ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_routetarget ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_routetarget ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_routetarget ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_rir ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_rir ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_rir ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_rir ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_rir ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_rir ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ ipam_prefix : "(id) - (status_id)"
    extras_role |o--o{ ipam_prefix : "(id) - (role_id)"
    ipam_prefix |o--o{ ipam_prefix : "(id) - (parent_id)"
    ipam_namespace ||--o{ ipam_prefix : "(id) - (namespace_id)"
    tenancy_tenant |o--o{ ipam_prefix : "(id) - (tenant_id)"
    ipam_vlan |o--o{ ipam_prefix : "(id) - (vlan_id)"
    ipam_rir |o--o{ ipam_prefix : "(id) - (rir_id)"
    ipam_prefix ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_prefix ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_prefix ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_prefix ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_prefix ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_prefix ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_prefix ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_prefix ||--o{ ipam_prefixlocationassignment : "(id) - (prefix_id)"
    dcim_location ||--o{ ipam_prefixlocationassignment : "(id) - (location_id)"
    ipam_prefixlocationassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_prefixlocationassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ ipam_ipaddress : "(id) - (status_id)"
    extras_role |o--o{ ipam_ipaddress : "(id) - (role_id)"
    ipam_prefix |o--o{ ipam_ipaddress : "(id) - (parent_id)"
    tenancy_tenant |o--o{ ipam_ipaddress : "(id) - (tenant_id)"
    ipam_ipaddress |o--o{ ipam_ipaddress : "(id) - (nat_inside_id)"
    ipam_ipaddress ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_ipaddress ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_ipaddress ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_ipaddress ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_ipaddress ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_ipaddress ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_ipaddress ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_ipaddress ||--o{ ipam_ipaddresstointerface : "(id) - (ip_address_id)"
    dcim_interface |o--o{ ipam_ipaddresstointerface : "(id) - (interface_id)"
    virtualization_vminterface |o--o{ ipam_ipaddresstointerface : "(id) - (vm_interface_id)"
    ipam_ipaddresstointerface ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_ipaddresstointerface ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_prefix ||--o{ ipam_ipaddressrange : "(id) - (parent_id)"
    extras_status ||--o{ ipam_ipaddressrange : "(id) - (status_id)"
    extras_role |o--o{ ipam_ipaddressrange : "(id) - (role_id)"
    tenancy_tenant |o--o{ ipam_ipaddressrange : "(id) - (tenant_id)"
    ipam_ipaddressrange ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_ipaddressrange ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_ipaddressrange ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_ipaddressrange ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_ipaddressrange ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_ipaddressrange ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_ipaddressrange ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_location |o--o{ ipam_vlangroup : "(id) - (location_id)"
    ipam_vlangroup ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_vlangroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_vlangroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_vlangroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_vlangroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_vlangroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vlangroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_vlangroup |o--o{ ipam_vlan : "(id) - (vlan_group_id)"
    extras_status ||--o{ ipam_vlan : "(id) - (status_id)"
    extras_role |o--o{ ipam_vlan : "(id) - (role_id)"
    tenancy_tenant |o--o{ ipam_vlan : "(id) - (tenant_id)"
    ipam_vlan ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_vlan ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_vlan ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_vlan ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_vlan ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_vlan ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vlan ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_vlan ||--o{ ipam_vlanlocationassignment : "(id) - (vlan_id)"
    dcim_location ||--o{ ipam_vlanlocationassignment : "(id) - (location_id)"
    ipam_vlanlocationassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_vlanlocationassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ ipam_service : "(id) - (device_id)"
    virtualization_virtualmachine |o--o{ ipam_service : "(id) - (virtual_machine_id)"
    ipam_service }o--o{ ipam_ipaddress : "(ip_addresses) - (services)"
    ipam_service ||..o{ extras_taggeditem : "(id) - (object_id)"
    ipam_service ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    ipam_service ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    ipam_service ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    ipam_service ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    ipam_service ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    ipam_service ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_ipaddress ||--o{ load_balancers_virtualserver : "(id) - (vip_id)"
    ipam_prefix |o--o{ load_balancers_virtualserver : "(id) - (source_nat_pool_id)"
    dcim_device |o--o{ load_balancers_virtualserver : "(id) - (device_id)"
    dcim_deviceredundancygroup |o--o{ load_balancers_virtualserver : "(id) - (device_redundancy_group_id)"
    cloud_cloudservice |o--o{ load_balancers_virtualserver : "(id) - (cloud_service_id)"
    dcim_virtualchassis |o--o{ load_balancers_virtualserver : "(id) - (virtual_chassis_id)"
    tenancy_tenant |o--o{ load_balancers_virtualserver : "(id) - (tenant_id)"
    load_balancers_loadbalancerpool |o--o{ load_balancers_virtualserver : "(id) - (load_balancer_pool_id)"
    load_balancers_healthcheckmonitor |o--o{ load_balancers_virtualserver : "(id) - (health_check_monitor_id)"
    load_balancers_virtualserver ||..o{ extras_taggeditem : "(id) - (object_id)"
    load_balancers_virtualserver ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    load_balancers_virtualserver ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    load_balancers_virtualserver ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    load_balancers_virtualserver ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    load_balancers_virtualserver ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_virtualserver ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    load_balancers_healthcheckmonitor |o--o{ load_balancers_loadbalancerpool : "(id) - (health_check_monitor_id)"
    tenancy_tenant |o--o{ load_balancers_loadbalancerpool : "(id) - (tenant_id)"
    load_balancers_loadbalancerpool ||..o{ extras_taggeditem : "(id) - (object_id)"
    load_balancers_loadbalancerpool ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    load_balancers_loadbalancerpool ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    load_balancers_loadbalancerpool ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    load_balancers_loadbalancerpool ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    load_balancers_loadbalancerpool ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_loadbalancerpool ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    ipam_ipaddress ||--o{ load_balancers_loadbalancerpoolmember : "(id) - (ip_address_id)"
    load_balancers_loadbalancerpool ||--o{ load_balancers_loadbalancerpoolmember : "(id) - (load_balancer_pool_id)"
    load_balancers_healthcheckmonitor |o--o{ load_balancers_loadbalancerpoolmember : "(id) - (health_check_monitor_id)"
    tenancy_tenant |o--o{ load_balancers_loadbalancerpoolmember : "(id) - (tenant_id)"
    extras_status ||--o{ load_balancers_loadbalancerpoolmember : "(id) - (status_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_taggeditem : "(id) - (object_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    load_balancers_loadbalancerpoolmember ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_loadbalancerpoolmember ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ load_balancers_healthcheckmonitor : "(id) - (tenant_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_taggeditem : "(id) - (object_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    load_balancers_healthcheckmonitor ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_healthcheckmonitor ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ load_balancers_certificateprofile : "(id) - (tenant_id)"
    load_balancers_certificateprofile ||..o{ extras_taggeditem : "(id) - (object_id)"
    load_balancers_certificateprofile ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    load_balancers_certificateprofile ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    load_balancers_certificateprofile ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    load_balancers_certificateprofile ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    load_balancers_certificateprofile ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_certificateprofile ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    load_balancers_virtualserver ||--o{ load_balancers_virtualservercertificateprofileassignment : "(id) - (virtual_server_id)"
    load_balancers_certificateprofile ||--o{ load_balancers_virtualservercertificateprofileassignment : "(id) - (certificate_profile_id)"
    load_balancers_virtualservercertificateprofileassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_virtualservercertificateprofileassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    load_balancers_loadbalancerpoolmember ||--o{ load_balancers_loadbalancerpoolmembercertificateprofileassi0862 : "(id) - (load_balancer_pool_member_id)"
    load_balancers_certificateprofile ||--o{ load_balancers_loadbalancerpoolmembercertificateprofileassi0862 : "(id) - (certificate_profile_id)"
    load_balancers_loadbalancerpoolmembercertificateprofileassi0862 ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    load_balancers_loadbalancerpoolmembercertificateprofileassi0862 ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenantgroup |o--o{ tenancy_tenantgroup : "(id) - (parent_id)"
    tenancy_tenantgroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    tenancy_tenantgroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    tenancy_tenantgroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    tenancy_tenantgroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    tenancy_tenantgroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    tenancy_tenantgroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenantgroup |o--o{ tenancy_tenant : "(id) - (tenant_group_id)"
    tenancy_tenant ||..o{ extras_taggeditem : "(id) - (object_id)"
    tenancy_tenant ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    tenancy_tenant ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    tenancy_tenant ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    tenancy_tenant ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    tenancy_tenant ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    tenancy_tenant ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    auth_user ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    auth_user ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    auth_user ||--o{ users_token : "(id) - (user_id)"
    users_token ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    users_token ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    users_objectpermission }o--o{ django_content_type : "(object_types) - (object_permissions)"
    users_objectpermission }o--o{ auth_group : "(groups) - (object_permissions)"
    users_objectpermission }o--o{ auth_user : "(users) - (object_permissions)"
    users_objectpermission ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    users_objectpermission ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    virtualization_clustertype ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    virtualization_clustertype ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    virtualization_clustertype ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_clustertype ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    virtualization_clustertype ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    virtualization_clustertype ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    virtualization_clustergroup ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    virtualization_clustergroup ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    virtualization_clustergroup ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_clustergroup ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    virtualization_clustergroup ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    virtualization_clustergroup ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    virtualization_clustertype ||--o{ virtualization_cluster : "(id) - (cluster_type_id)"
    virtualization_clustergroup |o--o{ virtualization_cluster : "(id) - (cluster_group_id)"
    tenancy_tenant |o--o{ virtualization_cluster : "(id) - (tenant_id)"
    dcim_location |o--o{ virtualization_cluster : "(id) - (location_id)"
    virtualization_cluster ||..o{ extras_taggeditem : "(id) - (object_id)"
    virtualization_cluster ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    virtualization_cluster ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    virtualization_cluster ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_cluster ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    virtualization_cluster ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    virtualization_cluster ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_configcontextschema |o--o{ virtualization_virtualmachine : "(id) - (local_config_context_schema_id)"
    django_content_type |o--o{ virtualization_virtualmachine : "(id) - (local_config_context_data_owner_content_type_id)"
    virtualization_cluster ||--o{ virtualization_virtualmachine : "(id) - (cluster_id)"
    tenancy_tenant |o--o{ virtualization_virtualmachine : "(id) - (tenant_id)"
    dcim_platform |o--o{ virtualization_virtualmachine : "(id) - (platform_id)"
    extras_status ||--o{ virtualization_virtualmachine : "(id) - (status_id)"
    extras_role |o--o{ virtualization_virtualmachine : "(id) - (role_id)"
    ipam_ipaddress |o--o{ virtualization_virtualmachine : "(id) - (primary_ip4_id)"
    ipam_ipaddress |o--o{ virtualization_virtualmachine : "(id) - (primary_ip6_id)"
    dcim_softwareversion |o--o{ virtualization_virtualmachine : "(id) - (software_version_id)"
    virtualization_virtualmachine }o--o{ dcim_softwareimagefile : "(software_image_files) - (virtual_machines)"
    virtualization_virtualmachine ||..o{ extras_taggeditem : "(id) - (object_id)"
    virtualization_virtualmachine ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    virtualization_virtualmachine ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    virtualization_virtualmachine ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_virtualmachine ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    virtualization_virtualmachine ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    virtualization_virtualmachine ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_status ||--o{ virtualization_vminterface : "(id) - (status_id)"
    extras_role |o--o{ virtualization_vminterface : "(id) - (role_id)"
    virtualization_vminterface |o--o{ virtualization_vminterface : "(id) - (parent_interface_id)"
    virtualization_vminterface |o--o{ virtualization_vminterface : "(id) - (bridge_id)"
    virtualization_virtualmachine ||--o{ virtualization_vminterface : "(id) - (virtual_machine_id)"
    ipam_vlan |o--o{ virtualization_vminterface : "(id) - (untagged_vlan_id)"
    ipam_vrf |o--o{ virtualization_vminterface : "(id) - (vrf_id)"
    virtualization_vminterface }o--o{ ipam_vlan : "(tagged_vlans) - (vminterfaces_as_tagged)"
    virtualization_vminterface ||..o{ extras_taggeditem : "(id) - (object_id)"
    virtualization_vminterface ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    virtualization_vminterface ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    virtualization_vminterface ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_vminterface ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    virtualization_vminterface ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    virtualization_vminterface ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    virtualization_vminterface ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    virtualization_vminterface ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    extras_role |o--o{ vpn_vpnprofile : "(id) - (role_id)"
    extras_secretsgroup |o--o{ vpn_vpnprofile : "(id) - (secrets_group_id)"
    tenancy_tenant |o--o{ vpn_vpnprofile : "(id) - (tenant_id)"
    vpn_vpnprofile ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpnprofile ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpnprofile ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpnprofile ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpnprofile ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpnprofile ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpnprofile ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ vpn_vpnphase1policy : "(id) - (tenant_id)"
    vpn_vpnphase1policy ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpnphase1policy ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpnphase1policy ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpnphase1policy ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpnphase1policy ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpnphase1policy ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpnphase1policy ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    tenancy_tenant |o--o{ vpn_vpnphase2policy : "(id) - (tenant_id)"
    vpn_vpnphase2policy ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpnphase2policy ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpnphase2policy ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpnphase2policy ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpnphase2policy ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpnphase2policy ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpnphase2policy ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    vpn_vpnprofile ||--o{ vpn_vpnprofilephase1policyassignment : "(id) - (vpn_profile_id)"
    vpn_vpnphase1policy ||--o{ vpn_vpnprofilephase1policyassignment : "(id) - (vpn_phase1_policy_id)"
    vpn_vpnprofilephase1policyassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpnprofilephase1policyassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    vpn_vpnprofile ||--o{ vpn_vpnprofilephase2policyassignment : "(id) - (vpn_profile_id)"
    vpn_vpnphase2policy ||--o{ vpn_vpnprofilephase2policyassignment : "(id) - (vpn_phase2_policy_id)"
    vpn_vpnprofilephase2policyassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpnprofilephase2policyassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    vpn_vpnprofile |o--o{ vpn_vpn : "(id) - (vpn_profile_id)"
    extras_role |o--o{ vpn_vpn : "(id) - (role_id)"
    tenancy_tenant |o--o{ vpn_vpn : "(id) - (tenant_id)"
    extras_status |o--o{ vpn_vpn : "(id) - (status_id)"
    vpn_vpn ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpn ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpn ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpn ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpn ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpn ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpn ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    vpn_vpnprofile |o--o{ vpn_vpntunnel : "(id) - (vpn_profile_id)"
    vpn_vpn |o--o{ vpn_vpntunnel : "(id) - (vpn_id)"
    extras_role |o--o{ vpn_vpntunnel : "(id) - (role_id)"
    extras_status ||--o{ vpn_vpntunnel : "(id) - (status_id)"
    extras_secretsgroup |o--o{ vpn_vpntunnel : "(id) - (secrets_group_id)"
    vpn_vpntunnelendpoint |o--o{ vpn_vpntunnel : "(id) - (endpoint_a_id)"
    vpn_vpntunnelendpoint |o--o{ vpn_vpntunnel : "(id) - (endpoint_z_id)"
    tenancy_tenant |o--o{ vpn_vpntunnel : "(id) - (tenant_id)"
    vpn_vpntunnel ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpntunnel ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpntunnel ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpntunnel ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpntunnel ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpntunnel ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpntunnel ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_device |o--o{ vpn_vpntunnelendpoint : "(id) - (device_id)"
    dcim_interface |o--o| vpn_vpntunnelendpoint : "(id) - (source_interface_id)"
    ipam_ipaddress |o--o{ vpn_vpntunnelendpoint : "(id) - (source_ipaddress_id)"
    dcim_interface |o--o| vpn_vpntunnelendpoint : "(id) - (tunnel_interface_id)"
    vpn_vpnprofile |o--o{ vpn_vpntunnelendpoint : "(id) - (vpn_profile_id)"
    extras_role |o--o{ vpn_vpntunnelendpoint : "(id) - (role_id)"
    tenancy_tenant |o--o{ vpn_vpntunnelendpoint : "(id) - (tenant_id)"
    vpn_vpntunnelendpoint }o--o{ ipam_prefix : "(protected_prefixes) - (vpn_tunnel_endpoints)"
    vpn_vpntunnelendpoint }o--o{ extras_dynamicgroup : "(protected_prefixes_dg) - (vpn_tunnel_endpoints)"
    vpn_vpntunnelendpoint ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpntunnelendpoint ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpntunnelendpoint ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpntunnelendpoint ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpntunnelendpoint ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpntunnelendpoint ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpntunnelendpoint ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    vpn_vpn ||--o{ vpn_vpntermination : "(id) - (vpn_id)"
    ipam_vlan |o--o{ vpn_vpntermination : "(id) - (vlan_id)"
    dcim_interface |o--o{ vpn_vpntermination : "(id) - (interface_id)"
    virtualization_vminterface |o--o{ vpn_vpntermination : "(id) - (vm_interface_id)"
    vpn_vpntermination ||..o{ extras_taggeditem : "(id) - (object_id)"
    vpn_vpntermination ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    vpn_vpntermination ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    vpn_vpntermination ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    vpn_vpntermination ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    vpn_vpntermination ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    vpn_vpntermination ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    wireless_supporteddatarate ||..o{ extras_taggeditem : "(id) - (object_id)"
    wireless_supporteddatarate ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    wireless_supporteddatarate ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    wireless_supporteddatarate ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    wireless_supporteddatarate ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    wireless_supporteddatarate ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    wireless_supporteddatarate ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    wireless_radioprofile }o--o{ wireless_supporteddatarate : "(supported_data_rates) - (radio_profiles)"
    wireless_radioprofile ||..o{ extras_taggeditem : "(id) - (object_id)"
    wireless_radioprofile ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    wireless_radioprofile ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    wireless_radioprofile ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    wireless_radioprofile ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    wireless_radioprofile ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    wireless_radioprofile ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    extras_secretsgroup |o--o{ wireless_wirelessnetwork : "(id) - (secrets_group_id)"
    tenancy_tenant |o--o{ wireless_wirelessnetwork : "(id) - (tenant_id)"
    wireless_wirelessnetwork ||..o{ extras_taggeditem : "(id) - (object_id)"
    wireless_wirelessnetwork ||..o{ extras_contactassociation : "(id) - (associated_object_id)"
    wireless_wirelessnetwork ||..o{ extras_staticgroupassociation : "(id) - (associated_object_id)"
    wireless_wirelessnetwork ||..o{ extras_relationshipassociation : "(id) - (source_id)"
    wireless_wirelessnetwork ||..o{ extras_relationshipassociation : "(id) - (destination_id)"
    wireless_wirelessnetwork ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    wireless_wirelessnetwork ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_controllermanageddevicegroup ||--o{ wireless_controllermanageddevicegroupwirelessnetworkassignment : "(id) - (controller_managed_device_group_id)"
    wireless_wirelessnetwork ||--o{ wireless_controllermanageddevicegroupwirelessnetworkassignment : "(id) - (wireless_network_id)"
    ipam_vlan |o--o{ wireless_controllermanageddevicegroupwirelessnetworkassignment : "(id) - (vlan_id)"
    wireless_controllermanageddevicegroupwirelessnetworkassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    wireless_controllermanageddevicegroupwirelessnetworkassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    dcim_controllermanageddevicegroup ||--o{ wireless_controllermanageddevicegroupradioprofileassignment : "(id) - (controller_managed_device_group_id)"
    wireless_radioprofile ||--o{ wireless_controllermanageddevicegroupradioprofileassignment : "(id) - (radio_profile_id)"
    wireless_controllermanageddevicegroupradioprofileassignment ||..o{ extras_objectmetadata : "(id) - (assigned_object_id)"
    wireless_controllermanageddevicegroupradioprofileassignment ||..o{ data_validation_datacompliance : "(id) - (object_id)"
    silk_request ||--o| silk_response : "(id) - (request_id)"
    silk_request |o--o{ silk_sqlquery : "(id) - (request_id)"
    silk_request |o--o{ silk_profile : "(id) - (request_id)"
    silk_profile }o--o{ silk_sqlquery : "(queries) - (profiles)"
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
    social_auth_usersocialauth {
        BigAutoField id PK
        ForeignKey user_id FK
        CharField provider
        CharField id_key
        CharField uid
        JSONField extra_data
        DateTimeField created
        DateTimeField modified
    }
    social_auth_nonce {
        BigAutoField id PK
        CharField server_url
        IntegerField timestamp
        CharField salt
    }
    social_auth_association {
        BigAutoField id PK
        CharField server_url
        CharField handle
        CharField secret
        IntegerField issued
        IntegerField lifetime
        CharField assoc_type
    }
    social_auth_code {
        BigAutoField id PK
        EmailField email
        CharField code
        BooleanField verified
        DateTimeField timestamp
    }
    social_auth_partial {
        BigAutoField id PK
        CharField token
        PositiveSmallIntegerField next_step
        CharField backend
        JSONField data
        DateTimeField timestamp
    }
    taggit_tag {
        AutoField id PK
        CharField name UK
        SlugField slug UK
    }
    taggit_taggeditem {
        AutoField id PK
        ForeignKey tag_id FK
        ForeignKey content_type_id FK
        IntegerField object_id
    }
    constance_constance {
        AutoField id PK
        CharField key UK
        TextField value
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
    django_celery_results_taskresult {
        AutoField id PK
        CharField task_id UK "Celery ID for the Task that was run"
        CharField periodic_task_name "Name of the Periodic Task which was run"
        CharField task_name "Name of the Task which was run"
        TextField task_args "JSON representation of the positional arguments used with the task"
        TextField task_kwargs "JSON representation of the named arguments used with the task"
        CharField status "Current state of the task being run"
        CharField worker "Worker that executes the task"
        CharField content_type "Content type of the result data"
        CharField content_encoding "The encoding used to save the task result data"
        TextField result "The data returned by the task.  Use content_encoding and content_type fields to read."
        DateTimeField date_created "Datetime field when the task result was created in UTC"
        DateTimeField date_started "Datetime field when the task was started in UTC"
        DateTimeField date_done "Datetime field when the task was completed in UTC"
        TextField traceback "Text of the traceback if the task generated one"
        TextField meta "JSON meta information about the task, such as information on child tasks"
    }
    django_celery_results_chordcounter {
        AutoField id PK
        CharField group_id UK "Celery ID for the Chord header group"
        TextField sub_tasks "JSON serialized list of task result tuples. use .group_result() to decode"
        PositiveIntegerField count "Starts at len(chord header) and decrements after each task is finished"
    }
    django_celery_results_groupresult {
        AutoField id PK
        CharField group_id UK "Celery ID for the Group that was run"
        DateTimeField date_created "Datetime field when the group result was created in UTC"
        DateTimeField date_done "Datetime field when the group was completed in UTC"
        CharField content_type "Content type of the result data"
        CharField content_encoding "The encoding used to save the task result data"
        TextField result "The data returned by the task.  Use content_encoding and content_type fields to read."
    }
    circuits_providernetwork {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        ForeignKey provider_id FK
        CharField description
        TextField comments
    }
    circuits_provider {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ASNField asn "32-bit autonomous system number"
        CharField account
        URLField portal_url
        TextField noc_contact
        TextField admin_contact
        TextField comments
    }
    circuits_circuittype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    circuits_circuit {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField cid
        StatusField status_id FK
        ForeignKey provider_id FK
        ForeignKey circuit_type_id FK
        ForeignKey tenant_id FK
        DateField install_date
        PositiveIntegerField commit_rate
        CharField description
        TextField comments
        ForeignKey circuit_termination_a_id FK
        ForeignKey circuit_termination_z_id FK
    }
    circuits_circuittermination {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey circuit_id FK
        CircuitTerminationTermSide term_side
        ForeignKey location_id FK
        ForeignKey provider_network_id FK
        ForeignKey cloud_network_id FK
        PositiveIntegerField port_speed
        PositiveIntegerField upstream_speed "Upstream speed, if different from port speed"
        CharField xconnect_id
        CharField pp_info
        CharField description
    }
    cloud_cloudaccount {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK "The name of this Cloud Account."
        CharField description
        CharField account_number "The account identifier of this Cloud Account."
        ForeignKey provider_id FK "The Manufacturer instance which represents the Cloud Provider"
        ForeignKey secrets_group_id FK
    }
    cloud_cloudresourcetype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK "Type of cloud objects"
        CharField description
        ForeignKey provider_id FK "The Manufacturer instance which represents the Cloud Provider"
        JSONField config_schema
    }
    cloud_cloudnetwork {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKeyLimitedByContentTypes cloud_resource_type_id FK
        JSONField extra_config
        CharField name UK
        CharField description
        ForeignKey cloud_account_id FK
        ForeignKey parent_id FK
    }
    cloud_cloudnetworkprefixassignment {
        UUIDField id PK
        ForeignKey cloud_network_id FK
        ForeignKey prefix_id FK
    }
    cloud_cloudservice {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKeyLimitedByContentTypes cloud_resource_type_id FK
        JSONField extra_config
        CharField name UK
        CharField description
        ForeignKey cloud_account_id FK
    }
    cloud_cloudservicenetworkassignment {
        UUIDField id PK
        ForeignKey cloud_network_id FK
        ForeignKey cloud_service_id FK
    }
    data_validation_regexrule {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField field
        ForeignKey content_type_id FK
        BooleanField enabled
        CharField error_message "Optional error message to display when validation fails."
        TextField regular_expression
        BooleanField context_processing "When enabled, the regular expression value is first processed as a Jinja2 template with access to the object being validated in a variable named <code>obj</code>."
    }
    data_validation_minmaxrule {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField field
        ForeignKey content_type_id FK
        BooleanField enabled
        CharField error_message "Optional error message to display when validation fails."
        FloatField min "When set, apply a minimum value contraint to the value of the model field."
        FloatField max "When set, apply a maximum value contraint to the value of the model field."
    }
    data_validation_requiredrule {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField field
        ForeignKey content_type_id FK
        BooleanField enabled
        CharField error_message "Optional error message to display when validation fails."
    }
    data_validation_uniquerule {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField field
        ForeignKey content_type_id FK
        BooleanField enabled
        CharField error_message "Optional error message to display when validation fails."
        PositiveIntegerField max_instances
    }
    data_validation_datacompliance {
        UUIDField id PK
        CharField compliance_class_name
        DateTimeField last_validation_date
        ForeignKey content_type_id FK
        UUIDField object_id
        CharField validated_object_str
        CharField validated_attribute
        CharField validated_attribute_value
        BooleanField valid
        TextField message
    }
    dcim_consoleport {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        ConsolePortType type "Physical port type"
    }
    dcim_consoleserverport {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        ConsoleServerPortType type "Physical port type"
    }
    dcim_powerport {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        PowerPortType type "Physical port type"
        PositiveSmallIntegerField maximum_draw "Maximum power draw (watts)"
        PositiveSmallIntegerField allocated_draw "Allocated power draw (watts)"
        DecimalField power_factor "Power factor (0.01-1.00) for converting between watts (W) and volt-amps (VA). Defaults to 0.95."
    }
    dcim_poweroutlet {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        PowerOutletType type "Physical port type"
        ForeignKey power_port_id FK
        PowerOutletFeedLeg feed_leg "Phase (for three-phase feeds)"
    }
    dcim_interface {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        StatusField status_id FK
        BooleanField enabled
        MACAddressCharField mac_address
        PositiveIntegerField mtu
        RoleField role_id FK
        InterfaceMode mode
        ForeignKey parent_interface_id FK "Assigned parent interface"
        ForeignKey bridge_id FK "Assigned bridge interface"
        NaturalOrderingField _name
        ForeignKey lag_id FK "Assigned LAG interface"
        InterfaceType type
        InterfacePortType port_type "Physical connector type"
        BooleanField mgmt_only "This interface is used only for out-of-band management"
        ForeignKey untagged_vlan_id FK
        ForeignKey vrf_id FK
        PositiveIntegerField speed
        InterfaceDuplex duplex
        PositiveSmallIntegerField breakout_position "For a child interface of a breakout-cable trunk, the position on the parent interface's trunk connector that this child interface maps to."
    }
    dcim_interfaceredundancygroup {
        UUIDField id PK
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        StatusField status_id FK
        DateTimeField created
        CharField description
        InterfaceRedundancyGroupProtocol protocol
        CharField protocol_group_id
        ForeignKey secrets_group_id FK
        ForeignKey virtual_ip_id FK
    }
    dcim_interfaceredundancygroupassociation {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey interface_redundancy_group_id FK
        ForeignKey interface_id FK
        PositiveIntegerField priority
    }
    dcim_frontport {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        FrontPortType type
        ForeignKey rear_port_id FK
        PositiveSmallIntegerField rear_port_position
    }
    dcim_rearport {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_id FK
        ForeignKeyWithAutoRelatedName module_id FK
        RearPortType type
        PositiveSmallIntegerField positions
    }
    dcim_devicebay {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKeyWithAutoRelatedName device_id FK
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        OneToOneField installed_device_id FK, UK
    }
    dcim_inventoryitem {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKeyWithAutoRelatedName device_id FK
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKey manufacturer_id FK
        CharField part_id "Manufacturer-assigned part identifier"
        CharField serial
        CharField asset_tag UK "A unique tag used to identify this item"
        BooleanField discovered "This item was automatically discovered"
        ForeignKey software_version_id FK "The software version installed on this inventory item"
    }
    dcim_modulebay {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey parent_device_id FK
        ForeignKey parent_module_id FK
        ForeignKey module_family_id FK "Module family that can be installed in this bay"
        BooleanField requires_first_party_modules "This bay will only accept modules from the same manufacturer as the parent device or module"
        CharField name
        NaturalOrderingField _name
        CharField position "The position of the module bay within the parent device/module"
        CharField label "Physical label"
        CharField description
    }
    dcim_cabletype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        ForeignKey manufacturer_id FK
        CharField part_number "Discrete part number (optional)"
        BooleanField has_embedded_transceivers "Indicates that this cable type has transceivers (e.g. SFP) built in."
        PositiveSmallIntegerField a_connectors "Number of physical connectors on the A side."
        PositiveSmallIntegerField b_connectors "Number of physical connectors on the B side."
        PositiveSmallIntegerField total_lanes "Total number of logical lanes in the breakout, distributed evenly across connectors on each side."
        JSONField mapping "A→B lane mapping as a JSON array of objects with keys: label, a_connector, a_position, b_connector, b_position. If empty, will be auto-populated."
        BooleanField is_shuffle "Indicates non-linear (polarity-shuffled) position mapping. Informational only."
        PositiveSmallIntegerField strands_per_lane "Number of physical strands per logical lane (e.g. 1 for copper, 2 for duplex fiber)."
        CableTypePolarityMethod polarity_method "Fiber polarity method. Informational only."
    }
    dcim_cable {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey cable_type_id FK "The cable type defining this cable's lane structure and other properties.May be null for standard point-to-point cables."
        CableType type
        StatusField status_id FK
        CharField label
        ColorField color
        PositiveSmallIntegerField length
        CableLengthUnit length_unit
        DecimalField _abs_length
    }
    dcim_cabletocabletermination {
        UUIDField id PK
        ForeignKey cable_id FK
        CableToCableTerminationCableEnd cable_end
        OneToOneField circuit_termination_id FK, UK
        OneToOneField console_port_id FK, UK
        OneToOneField console_server_port_id FK, UK
        OneToOneField front_port_id FK, UK
        OneToOneField interface_id FK, UK
        OneToOneField power_feed_id FK, UK
        OneToOneField power_outlet_id FK, UK
        OneToOneField power_port_id FK, UK
        OneToOneField rear_port_id FK, UK
        PositiveSmallIntegerField connector "The connector number on this cable end. Always 1 for standard cables."
    }
    dcim_cablepath {
        UUIDField id PK
        ForeignKey origin_type_id FK
        UUIDField origin_id
        ForeignKey destination_type_id FK
        UUIDField destination_id
        JSONPathField path
        BooleanField is_active
        BooleanField is_split
        PositiveSmallIntegerField peer_connector
        BooleanField origin_fans_out
        BooleanField destination_fans_out
    }
    dcim_consoleporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        ConsolePortTemplateType type
    }
    dcim_consoleserverporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        ConsoleServerPortTemplateType type
    }
    dcim_powerporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        PowerPortTemplateType type
        PositiveSmallIntegerField maximum_draw "Maximum power draw (watts)"
        PositiveSmallIntegerField allocated_draw "Allocated power draw (watts)"
        DecimalField power_factor "Power factor (0.01-1.00) for converting between watts (W) and volt-amps (VA). Defaults to 0.95."
    }
    dcim_poweroutlettemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        PowerOutletTemplateType type
        ForeignKey power_port_template_id FK
        PowerOutletTemplateFeedLeg feed_leg "Phase (for three-phase feeds)"
    }
    dcim_interfacetemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        NaturalOrderingField _name
        InterfaceTemplateType type
        InterfaceTemplatePortType port_type "Physical connector type"
        BooleanField mgmt_only
        PositiveIntegerField speed
        InterfaceTemplateDuplex duplex
    }
    dcim_frontporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        FrontPortTemplateType type
        ForeignKey rear_port_template_id FK
        PositiveSmallIntegerField rear_port_position
    }
    dcim_rearporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        RearPortTemplateType type
        PositiveSmallIntegerField positions
    }
    dcim_devicebaytemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKeyWithAutoRelatedName device_type_id FK
        CharField name
        NaturalOrderingField _name
        CharField label "Physical label"
        CharField description
    }
    dcim_modulebaytemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        ForeignKeyWithAutoRelatedName device_type_id FK
        ForeignKeyWithAutoRelatedName module_type_id FK
        CharField position "The position of the module bay within the device or module"
        CharField label "Physical label"
        CharField description
        ForeignKey module_family_id FK "Module family that can be installed in this bay. Leave blank for no restriction."
        BooleanField requires_first_party_modules "This bay will only accept modules from the same manufacturer as the parent device or module"
    }
    dcim_manufacturer {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    dcim_devicefamily {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    dcim_devicetypetosoftwareimagefile {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey device_type_id FK
        ForeignKey software_image_file_id FK
    }
    dcim_devicetype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey manufacturer_id FK
        ForeignKey device_family_id FK
        CharField model
        CharField part_number "Discrete part number (optional)"
        PositiveSmallIntegerField u_height
        BooleanField is_full_depth "Device consumes both front and rear rack faces"
        DeviceTypeSubdeviceRole subdevice_role "Parent devices house child devices in device bays. Leave blank if this device type is neither a parent nor a child."
        ImageField front_image
        ImageField rear_image
        TextField comments
    }
    dcim_platform {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ForeignKey manufacturer_id FK "Optionally limit this platform to devices of a certain manufacturer"
        CharField network_driver "The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate"
        CharField napalm_driver "The name of the NAPALM driver to use when Nautobot internals interact with devices"
        JSONField napalm_args "Additional arguments to pass when initiating the NAPALM driver (JSON format)"
        CharField description
    }
    dcim_device {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        JSONField local_config_context_data
        ForeignKeyWithAutoRelatedName local_config_context_schema_id FK "Optional schema to validate the structure of the data"
        ForeignKeyWithAutoRelatedName local_config_context_data_owner_content_type_id FK
        UUIDField local_config_context_data_owner_object_id
        ForeignKey device_type_id FK
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
        ForeignKey platform_id FK
        CharField name
        NaturalOrderingField _name
        CharField serial
        CharField asset_tag UK "A unique tag used to identify this device"
        ForeignKey location_id FK
        ForeignKey rack_id FK
        PositiveSmallIntegerField position "The lowest-numbered unit occupied by the device"
        DeviceFace face
        ForeignKey primary_ip4_id FK
        ForeignKey primary_ip6_id FK
        ForeignKey virtual_chassis_id FK
        ForeignKey device_redundancy_group_id FK
        PositiveIntegerField device_redundancy_group_priority "The priority the device has in the device redundancy group."
        ForeignKey software_version_id FK "The software version installed on this device"
        PositiveSmallIntegerField vc_position
        PositiveSmallIntegerField vc_priority
        TextField comments
        ForeignKey secrets_group_id FK
        ForeignKey controller_managed_device_group_id FK
    }
    dcim_deviceclusterassignment {
        UUIDField id PK
        ForeignKey device_id FK
        ForeignKey cluster_id FK
    }
    dcim_virtualchassis {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        OneToOneField master_id FK, UK
        CharField name UK
        CharField domain
    }
    dcim_deviceredundancygroup {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        StatusField status_id FK
        CharField description
        DeviceRedundancyGroupFailoverStrategy failover_strategy
        TextField comments
        ForeignKey secrets_group_id FK
    }
    dcim_softwareimagefile {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey software_version_id FK
        CharField image_file_name
        CharField image_file_checksum
        SoftwareImageFileHashingAlgorithm hashing_algorithm "Hashing algorithm for image file checksum"
        PositiveBigIntegerField image_file_size "Image file size in bytes"
        LaxURLField download_url
        ForeignKey external_integration_id FK
        BooleanField default_image "Is the default image for this software version"
        StatusField status_id FK
    }
    dcim_softwareversion {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey platform_id FK
        CharField version
        CharField alias "Optional alternative label for this version"
        DateField release_date
        DateField end_of_support_date
        URLField documentation_url
        BooleanField long_term_support "Is a Long Term Support version"
        BooleanField pre_release "Is a Pre-Release version"
        StatusField status_id FK
    }
    dcim_controller {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        StatusField status_id FK
        CharField description
        ForeignKey location_id FK
        ForeignKey platform_id FK
        RoleField role_id FK
        JSONArrayField capabilities "List of capabilities supported by the controller, these capabilities are used to enhance views in Nautobot."
        ForeignKey tenant_id FK
        ForeignKey external_integration_id FK
        ForeignKey controller_device_id FK
        ForeignKey controller_device_redundancy_group_id FK
    }
    dcim_controllermanageddevicegroup {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK "Name of the controller device group"
        CharField description
        PositiveIntegerField weight "Weight of the controller device group, used to sort the groups within its parent group"
        ForeignKey controller_id FK "Controller that manages the devices in this group"
        JSONArrayField capabilities "List of capabilities supported by the controller device group, these capabilities are used to enhance views in Nautobot."
        ForeignKey tenant_id FK
    }
    dcim_modulefamily {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    dcim_moduletype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey manufacturer_id FK
        ForeignKey module_family_id FK
        CharField model
        CharField part_number "Discrete part number (optional)"
        ImageField front_image
        ImageField rear_image
        TextField comments
    }
    dcim_module {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey module_type_id FK
        OneToOneField parent_module_bay_id FK, UK
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
        CharField serial
        CharField asset_tag UK "A unique tag used to identify this module"
        ForeignKey location_id FK
    }
    dcim_virtualdevicecontext {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        ForeignKey device_id FK
        PositiveSmallIntegerField identifier "Unique identifier provided by the platform being virtualized (Example: Nexus VDC Identifier)"
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey primary_ip4_id FK
        ForeignKey primary_ip6_id FK
        ForeignKey tenant_id FK
        CharField description
        ForeignKey controller_managed_device_group_id FK
    }
    dcim_interfacevdcassignment {
        UUIDField id PK
        ForeignKey virtual_device_context_id FK
        ForeignKey interface_id FK
    }
    dcim_locationtype {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        BooleanField nestable "Allow Locations of this type to be parents/children of other Locations of this same type"
    }
    dcim_location {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        ForeignKey location_type_id FK
        StatusField status_id FK
        ForeignKey tenant_id FK
        CharField description
        CharField facility "Local facility ID or description"
        ASNField asn "32-bit autonomous system number"
        LocationTimeZone time_zone
        TextField physical_address
        TextField shipping_address
        DecimalField latitude "GPS coordinate (latitude)"
        DecimalField longitude "GPS coordinate (longitude)"
        CharField contact_name
        CharField contact_phone
        EmailField contact_email
        TextField comments
    }
    dcim_powerpanel {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey location_id FK
        ForeignKey rack_group_id FK
        CharField name
        PowerPanelPanelType panel_type
        PositiveIntegerField breaker_position_count "Total number of breaker positions in the panel (e.g., 42)"
        PowerPanelPowerPath power_path "Physical power distribution redundancy path."
    }
    dcim_powerfeed {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey power_panel_id FK "Source panel that originates this power feed"
        ForeignKey destination_panel_id FK "Destination panel that receives power from this feed"
        ForeignKey rack_id FK
        CharField name
        StatusField status_id FK
        PowerFeedType type
        PowerFeedPowerPath power_path "Physical power distribution redundancy path."
        PowerFeedSupply supply
        PowerFeedPhase phase
        SmallIntegerField voltage
        PositiveSmallIntegerField amperage
        PositiveSmallIntegerField max_utilization "Maximum permissible draw (percentage)"
        PositiveIntegerField breaker_position "Starting circuit breaker position in panel"
        PowerFeedBreakerPoleCount breaker_pole_count "Number of breaker poles"
        PositiveIntegerField available_power
        TextField comments
    }
    dcim_rackgroup {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        ForeignKey location_id FK
        CharField description
    }
    dcim_rack {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        NaturalOrderingField _name
        StatusField status_id FK
        RoleField role_id FK
        CharField facility_id "Locally-assigned identifier"
        ForeignKey location_id FK
        ForeignKey rack_group_id FK "Assigned group"
        ForeignKey tenant_id FK
        CharField serial
        CharField asset_tag UK "A unique tag used to identify this rack"
        RackType type
        RackWidth width "Rail-to-rail width"
        PositiveSmallIntegerField u_height "Height in rack units"
        BooleanField desc_units "Units are numbered top-to-bottom"
        PositiveSmallIntegerField outer_width "Outer dimension of rack (width)"
        PositiveSmallIntegerField outer_depth "Outer dimension of rack (depth)"
        RackOuterUnit outer_unit
        TextField comments
    }
    dcim_rackreservation {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey rack_id FK
        JSONArrayField units
        ForeignKey tenant_id FK
        ForeignKey user_id FK
        CharField description
    }
    extras_objectchange {
        UUIDField id PK
        DateTimeField time
        ForeignKey user_id FK
        CharField user_name
        UUIDField request_id
        ObjectChangeAction action
        ForeignKey changed_object_type_id FK
        UUIDField changed_object_id
        ObjectChangeChangeContext change_context
        CharField change_context_detail
        ForeignKey related_object_type_id FK
        UUIDField related_object_id
        CharField object_repr
        JSONField object_data
        JSONField object_data_v2
    }
    extras_computedfield {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey content_type_id FK
        AutoSlugField key UK "Internal field name. Please use underscores rather than dashes in this key."
        ComputedFieldOutputType output_type "How the rendered result is to be output, by default it is as Plain Text"
        CharField grouping "Human-readable grouping that this computed field belongs to."
        CharField label "Name of the field as displayed to users"
        CharField description
        TextField template "Jinja2 template code for field value"
        CharField fallback_value "Fallback value (if any) to be output for the field in the case of a template rendering error."
        PositiveSmallIntegerField weight
        BooleanField advanced_ui "Hide this field from the object's primary information tab. It will appear in the 'Advanced' tab instead."
    }
    extras_customfield {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField grouping "Human-readable grouping that this custom field belongs to."
        CustomFieldType type "The type of value(s) allowed for this field."
        CharField label "Name of the field as displayed to users."
        AutoSlugField key UK "Internal field name. Please use underscores rather than dashes in this key."
        CharField description "A helpful description for this field."
        BooleanField required "If true, this field is required when creating new objects or editing an existing object."
        CustomFieldFilterLogic filter_logic "Loose matches any instance of a given string; Exact matches the entire field."
        JSONField default "Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. 'Foo')."
        PositiveSmallIntegerField weight "Fields with higher weights appear lower in a form."
        BigIntegerField validation_minimum "Minimum allowed value (for numeric fields) or length (for text fields)."
        BigIntegerField validation_maximum "Maximum allowed value (for numeric fields) or length (for text fields)."
        CharField validation_regex "Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters. Regular expression on select and multi-select will be applied at <code>Custom Field Choices</code> definition."
        BooleanField advanced_ui "Hide this field from the object's primary information tab. It will appear in the 'Advanced' tab instead."
        JSONField scope_filter "A JSON-encoded dictionary of filter parameters defining possible objects that can use this custom field."
    }
    extras_customfieldchoice {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey custom_field_id FK
        CharField value
        PositiveSmallIntegerField weight "Higher weights appear later in the list"
    }
    extras_relationship {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField label UK "Label of the relationship as displayed to users"
        AutoSlugField key UK "Internal relationship key. Please use underscores rather than dashes in this key."
        CharField description
        RelationshipType type "Cardinality of this relationship"
        RelationshipRequiredOn required_on "Objects on the specified side MUST implement this relationship. Not permitted for symmetric relationships."
        ForeignKey source_type_id FK "The source object type to which this relationship applies."
        CharField source_label "Label for related destination objects, as displayed on the source object."
        BooleanField source_hidden "Hide this relationship on the source object."
        JSONField source_filter "Filterset filter matching the applicable source objects of the selected type"
        ForeignKey destination_type_id FK "The destination object type to which this relationship applies."
        CharField destination_label "Label for related source objects, as displayed on the destination object."
        BooleanField destination_hidden "Hide this relationship on the destination object."
        JSONField destination_filter "Filterset filter matching the applicable destination objects of the selected type"
        BooleanField advanced_ui "Hide this field from the object's primary information tab. It will appear in the 'Advanced' tab instead."
    }
    extras_relationshipassociation {
        UUIDField id PK
        ForeignKey relationship_id FK
        ForeignKey source_type_id FK
        UUIDField source_id
        ForeignKey destination_type_id FK
        UUIDField destination_id
    }
    extras_tag {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ColorField color
        CharField description
    }
    extras_taggeditem {
        ForeignKey content_type_id FK
        UUIDField object_id
        UUIDField id PK
        ForeignKey tag_id FK
    }
    extras_approvalworkflowdefinition {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ForeignKey model_content_type_id FK
        JSONField model_constraints "Constraints to filter the objects that can be approved using this workflow."
        IntegerField weight "Determines workflow relevance when multiple apply. Higher weight wins."
    }
    extras_approvalworkflowstagedefinition {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey approval_workflow_definition_id FK "Approval workflow definition to which this stage belongs."
        PositiveIntegerField sequence "The sequence dictates the order in which this stage will need to be approved. The lower the number, the earlier it will be."
        CharField name
        PositiveIntegerField min_approvers "Minimum number of approvers required to approve this stage."
        CharField denial_message "Message to show when the stage is denied."
        ForeignKey approver_group_id FK "Group of users who are eligible to approve this stage. Only admin users can create new groups."
    }
    extras_approvalworkflow {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey approval_workflow_definition_id FK "Approval workflow definition to which this approval workflow belongs."
        ForeignKey object_under_review_content_type_id FK
        UUIDField object_under_review_object_id
        ApprovalWorkflowCurrentState current_state "Current state of the approval workflow. Eligible values are: Pending, Approved, Denied, Canceled."
        DateTimeField decision_date "Date and time when the decision of approval/denial was made."
        ForeignKey user_id FK
        CharField user_name
    }
    extras_approvalworkflowstage {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey approval_workflow_id FK "Approval workflow to which this stage belongs."
        ForeignKey approval_workflow_stage_definition_id FK "Approval workflow stage definition to which this stage belongs."
        ApprovalWorkflowStageState state "State of the approval workflow stage instance. Eligible values are: Pending, Approved, Denied."
        DateTimeField decision_date "Date and time when the decision of approval/denial was made."
    }
    extras_approvaluserresponse {
        UUIDField id PK
        ForeignKey approval_workflow_stage_id FK
        ForeignKey user_id FK
        CharField comments "User comments to explain the decision that he/she made"
        ApprovalWorkflowStageResponseState state "User response to this approval workflow stage instance. Eligible values are: Pending, Comment, Approved, Denied."
        DateTimeField last_updated
    }
    extras_role {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ColorField color
        CharField description
        PositiveSmallIntegerField weight
    }
    extras_status {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ColorField color
        CharField description
    }
    extras_contact {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        CharField phone
        EmailField email
        TextField address
        TextField comments
    }
    extras_team {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        CharField phone
        EmailField email
        TextField address
        TextField comments
    }
    extras_contactassociation {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey contact_id FK
        ForeignKey team_id FK
        ForeignKey associated_object_type_id FK
        UUIDField associated_object_id
        RoleField role_id FK
        StatusField status_id FK
    }
    extras_gitrepository {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        AutoSlugField slug UK "Internal field name. Please use underscores rather than dashes in this key."
        LaxURLField remote_url "Only HTTP and HTTPS URLs are presently supported"
        CharField branch "Branch, tag, or commit"
        CharField current_head "Commit hash of the most recent fetch from the selected branch. Used for syncing between workers."
        ForeignKey secrets_group_id FK
        JSONField provided_contents
    }
    extras_dynamicgroup {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        DynamicGroupGroupType group_type
        ForeignKey content_type_id FK "The type of object contained in this group."
        ForeignKey tenant_id FK
        JSONField filter "A JSON-encoded dictionary of filter parameters defining membership of this group"
    }
    extras_dynamicgroupmembership {
        UUIDField id PK
        ForeignKey group_id FK
        ForeignKey parent_group_id FK
        DynamicGroupMembershipOperator operator
        PositiveSmallIntegerField weight
    }
    extras_staticgroupassociation {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey dynamic_group_id FK
        ForeignKey associated_object_type_id FK
        UUIDField associated_object_id
    }
    extras_job {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField module_name "Dotted name of the Python module providing this job"
        CharField job_class_name "Name of the Python class providing this job"
        CharField grouping "Human-readable grouping that this job belongs to"
        CharField name UK "Human-readable name of this job"
        TextField description "Markdown formatting and a limited subset of HTML are supported"
        BooleanField installed "Whether the Python module and class providing this job are presently installed and loadable"
        BooleanField enabled "Whether this job can be executed by users"
        BooleanField is_job_hook_receiver "Whether this job is a job hook receiver"
        BooleanField is_job_button_receiver "Whether this job is a job button receiver"
        BooleanField has_sensitive_variables "Whether this job contains sensitive variables"
        BooleanField is_singleton "Whether this job should fail to run if another instance of this job is already running"
        BooleanField console_log_default "Whether the job defaults to running with console log argument set to true"
        BooleanField hidden "Whether the job defaults to not being shown in the UI"
        BooleanField dryrun_default "Whether the job defaults to running with dryrun argument set to true"
        BooleanField read_only "Set to true if the job does not make any changes to the environment"
        FloatField soft_time_limit "Maximum runtime in seconds before the job will receive a <code>SoftTimeLimitExceeded</code> exception.<br>Set to 0 to use Nautobot system default"
        FloatField time_limit "Maximum runtime in seconds before the job will be forcibly terminated.<br>Set to 0 to use Nautobot system default"
        BooleanField supports_dryrun "If supported, allows the job to bypass approval when running with dryrun argument set to true"
        ForeignKey default_job_queue_id FK
        BooleanField grouping_override "If set, the configured grouping will remain even if the underlying Job source code changes"
        BooleanField name_override "If set, the configured name will remain even if the underlying Job source code changes"
        BooleanField console_log_default_override "If set, the configured console log default will remain even if the underlying Job source code changes"
        BooleanField description_override "If set, the configured description will remain even if the underlying Job source code changes"
        BooleanField dryrun_default_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField hidden_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField soft_time_limit_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField time_limit_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField has_sensitive_variables_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField job_queues_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField default_job_queue_override "If set, the configured value will remain even if the underlying Job source code changes"
        BooleanField is_singleton_override "If set, the configured value will remain even if the underlying Job source code changes"
    }
    extras_jobhook {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        BooleanField enabled
        ForeignKey job_id FK "The job that this job hook will initiate"
        CharField name UK
        BooleanField type_create "Call this job hook when a matching object is created."
        BooleanField type_delete "Call this job hook when a matching object is deleted."
        BooleanField type_update "Call this job hook when a matching object is updated."
    }
    extras_joblogentry {
        UUIDField id PK
        ForeignKey job_result_id FK
        JobLogEntryLogLevel log_level
        CharField grouping
        TextField message
        DateTimeField created
        CharField log_object
        CharField absolute_url
    }
    extras_jobqueue {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        JobQueueQueueType queue_type
        ForeignKey tenant_id FK
    }
    extras_jobqueueassignment {
        UUIDField id PK
        ForeignKey job_id FK
        ForeignKey job_queue_id FK
    }
    extras_jobresult {
        UUIDField id PK
        JSONField _custom_field_data
        ForeignKey job_model_id FK
        CharField name
        CharField task_name "Registered name of the Celery task for this job. Internal use only."
        DateTimeField date_created
        DateTimeField date_started
        DateTimeField date_done
        ForeignKey user_id FK
        JobResultStatus status "Current state of the Job being run"
        JSONField result "The data returned by the task"
        CharField worker
        JSONField task_args
        JSONField task_kwargs
        JSONField celery_kwargs
        TextField traceback
        JSONField meta
        ForeignKey scheduled_job_id FK
        PositiveIntegerField debug_log_count
        PositiveIntegerField success_log_count
        PositiveIntegerField info_log_count
        PositiveIntegerField warning_log_count
        PositiveIntegerField error_log_count
        ForeignKey canceled_by_id FK "The user who initiated the cancel action."
        CharField canceled_by_user_name
        JobResultCancelType cancel_type "Cancel type of the Job being canceled"
        DateTimeField date_canceled "Timestamp at which the job was canceled"
    }
    extras_jobconsoleentry {
        UUIDField id PK
        ForeignKey job_result_id FK
        DateTimeField timestamp "Timestamp when this output has been produced / received"
        JobConsoleEntryOutputType output_type "Type of the output (e.g. stdout, stderr, output)"
        TextField text "Actual line of output data"
    }
    extras_jobbutton {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField name UK
        BooleanField enabled
        CharField text "Jinja2 template code for button text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Buttons which render as empty text will not be displayed."
        ForeignKey job_id FK "Job this button will run"
        PositiveSmallIntegerField weight
        CharField group_name "Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group."
        JobButtonButtonClass button_class
        BooleanField confirmation "Enable confirmation pop-up box. <span class='text-danger'>WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!</span>"
    }
    extras_scheduledjobs {
        SmallIntegerField ident PK
        DateTimeField last_update
    }
    extras_scheduledjob {
        UUIDField id PK
        CharField name UK "Human-readable description of this scheduled task"
        CharField task "The name of the Celery task that should be run. (Example: 'proj.tasks.import_contacts')"
        ScheduledJobInterval interval
        CharField crontab "Cronjob syntax string for custom scheduling"
        ForeignKey solar_id FK
        ForeignKey clocked_id FK
        JSONField args
        JSONField kwargs
        CharField exchange
        CharField routing_key
        TextField headers
        PositiveIntegerField priority
        DateTimeField expires
        PositiveIntegerField expire_seconds
        BooleanField one_off "If True, the schedule will only run the task a single time"
        DateTimeField start_time "Datetime when the schedule should begin triggering the task to run"
        ScheduledJobState state "Current state of the Scheduled Job"
        BooleanField enabled "Set to False to disable the schedule"
        DateTimeField last_run_at "Datetime that the schedule last triggered the task to run. Reset to None if enabled is set to False."
        PositiveIntegerField total_run_count "Running count of how many times the schedule has triggered the task"
        DateTimeField date_changed "Datetime that this scheduled job was last modified"
        TextField description "Detailed description about the details of this scheduled job"
        JSONField celery_kwargs
        ForeignKey job_model_id FK
        ForeignKey job_queue_id FK
        ScheduledJobTimeZone time_zone
        ForeignKey user_id FK "User that requested the schedule"
        DateTimeField decision_date "Datetime that the schedule was approved or denied"
    }
    extras_metadatatype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        MetadataTypeDataType data_type "The type of data allowed for any Metadata of this type."
    }
    extras_metadatachoice {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey metadata_type_id FK
        CharField value
        PositiveSmallIntegerField weight "Higher weights appear later in the list"
    }
    extras_objectmetadata {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey metadata_type_id FK
        ForeignKey contact_id FK
        ForeignKey team_id FK
        JSONArrayField scoped_fields "List of scoped fields, only direct fields on the model"
        JSONField _value "Relevant data value to an object field or a set of object fields"
        ForeignKey assigned_object_type_id FK
        UUIDField assigned_object_id
    }
    extras_configcontext {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField name UK
        ForeignKey owner_content_type_id FK
        UUIDField owner_object_id
        PositiveSmallIntegerField weight
        CharField description
        BooleanField is_active
        ForeignKey config_context_schema_id FK "Optional schema to validate the structure of the data"
        JSONField data
    }
    extras_configcontextschema {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        JSONField data_schema "A JSON Schema document which is used to validate a config context object."
        ForeignKey owner_content_type_id FK
        UUIDField owner_object_id
    }
    extras_customlink {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey content_type_id FK
        CharField name UK
        CharField text "Jinja2 template code for link text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Links which render as empty text will not be displayed."
        CharField target_url "Jinja2 template code for link URL. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>."
        PositiveSmallIntegerField weight
        CharField group_name "Links with the same group will appear as a dropdown menu"
        CustomLinkButtonClass button_class "The class of the first link in a group will be used for the dropdown button"
        BooleanField new_window "Force link to open in a new window"
    }
    extras_exporttemplate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey owner_content_type_id FK
        UUIDField owner_object_id
        ForeignKey content_type_id FK
        CharField name
        CharField description
        TextField template_code "The list of objects being exported is passed as a context variable named <code>queryset</code>."
        CharField mime_type "Defaults to <code>text/plain</code>"
        CharField file_extension "Extension to append to the rendered filename"
    }
    extras_externalintegration {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        LaxURLField remote_url
        ForeignKey secrets_group_id FK "Credentials used for authenticating with the remote system"
        BooleanField verify_ssl "Verify SSL certificates when connecting to the remote system"
        IntegerField timeout "Number of seconds to wait for a response"
        JSONField extra_config "Optional user-defined JSON data for this integration"
        ExternalIntegrationHttpMethod http_method
        JSONField headers "Headers for the HTTP request"
        CharField ca_file_path
    }
    extras_fileattachment {
        UUIDField id PK
        BinaryField bytes
        CharField filename
        CharField mimetype
    }
    extras_fileproxy {
        UUIDField id PK
        CharField name
        FileField file
        DateTimeField uploaded_at
        ForeignKey job_result_id FK
    }
    extras_graphqlquery {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey owner_content_type_id FK
        UUIDField owner_object_id
        CharField name UK
        TextField query
        JSONField variables
    }
    extras_healthchecktestmodel {
        UUIDField id PK
        CharField title
    }
    extras_imageattachment {
        UUIDField id PK
        ForeignKey content_type_id FK
        UUIDField object_id
        ImageField image
        PositiveSmallIntegerField image_height
        PositiveSmallIntegerField image_width
        CharField name
        DateTimeField created
    }
    extras_note {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey assigned_object_type_id FK
        UUIDField assigned_object_id
        ForeignKey user_id FK
        CharField user_name
        TextField note
    }
    extras_savedview {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        ForeignKey owner_id FK "The user that created this view"
        CharField name "The name of this view"
        CharField view "The name of the list view that the saved view is derived from, e.g. dcim:device_list"
        JSONField config "Saved Configuration on this view"
        BooleanField is_global_default
        BooleanField is_shared
    }
    extras_usersavedviewassociation {
        UUIDField id PK
        ForeignKey saved_view_id FK
        ForeignKey user_id FK
        CharField view_name
    }
    extras_webhook {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField name UK
        BooleanField type_create "Call this webhook when a matching object is created."
        BooleanField type_update "Call this webhook when a matching object is updated."
        BooleanField type_delete "Call this webhook when a matching object is deleted."
        CharField payload_url "A POST will be sent to this URL when the webhook is called."
        BooleanField enabled
        WebhookHttpMethod http_method
        CharField http_content_type "The complete list of official content types is available <a href='https://www.iana.org/assignments/media-types/media-types.xhtml'>here</a>."
        TextField additional_headers "User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below)."
        TextField body_template "Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>."
        CharField secret "When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request."
        BooleanField ssl_verification "Enable SSL certificate verification. Disable with caution!"
        CharField ca_file_path "The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults."
    }
    extras_secret {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        CharField provider
        JSONField parameters
    }
    extras_secretsgroup {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    extras_secretsgroupassociation {
        UUIDField id PK
        ForeignKey secrets_group_id FK
        ForeignKey secret_id FK
        SecretsGroupAssociationAccessType access_type
        SecretsGroupAssociationSecretType secret_type
    }
    ipam_namespace {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        ForeignKey location_id FK
        ForeignKey tenant_id FK
    }
    ipam_vrf {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        CharField rd "Unique route distinguisher (as defined in RFC 4364)"
        StatusField status_id FK
        ForeignKey namespace_id FK
        ForeignKey tenant_id FK
        CharField description
    }
    ipam_vrfdeviceassignment {
        UUIDField id PK
        ForeignKey vrf_id FK
        ForeignKey device_id FK
        ForeignKey virtual_machine_id FK
        ForeignKey virtual_device_context_id FK
        CharField rd "Unique route distinguisher (as defined in RFC 4364)"
        CharField name
    }
    ipam_vrfprefixassignment {
        UUIDField id PK
        ForeignKey vrf_id FK
        ForeignKey prefix_id FK
    }
    ipam_routetarget {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK "Route target value (formatted in accordance with RFC 4360)"
        CharField description
        ForeignKey tenant_id FK
    }
    ipam_rir {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        BooleanField is_private "IP space managed by this RIR is considered private"
        CharField description
    }
    ipam_prefix {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        VarbinaryIPField network "IPv4 or IPv6 network address"
        VarbinaryIPField broadcast "IPv4 or IPv6 broadcast address"
        IntegerField prefix_length "Length of the Network prefix, in bits."
        PrefixType type
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey parent_id FK "The parent Prefix of this Prefix."
        PrefixIpVersion ip_version
        ForeignKey namespace_id FK
        ForeignKey tenant_id FK
        ForeignKey vlan_id FK
        ForeignKey rir_id FK "Regional Internet Registry responsible for this prefix"
        DateTimeField date_allocated "Date this prefix was allocated to an RIR, reserved in IPAM, etc."
        CharField description
    }
    ipam_prefixlocationassignment {
        UUIDField id PK
        ForeignKey prefix_id FK
        ForeignKey location_id FK
    }
    ipam_ipaddress {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        VarbinaryIPField host "IPv4 or IPv6 host address"
        IntegerField mask_length "Length of the network mask, in bits."
        IPAddressType type
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey parent_id FK "The parent Prefix of this IPAddress."
        IPAddressIpVersion ip_version
        ForeignKey tenant_id FK
        ForeignKey nat_inside_id FK "The IP Addresses for which this address is the 'outside' IP"
        CharField dns_name "Hostname or FQDN (not case-sensitive)"
        CharField description
    }
    ipam_ipaddresstointerface {
        UUIDField id PK
        ForeignKey ip_address_id FK
        ForeignKey interface_id FK
        ForeignKey vm_interface_id FK
        BooleanField is_source "Is source address on interface"
        BooleanField is_destination "Is destination address on interface"
        BooleanField is_default "Is default address on interface"
        BooleanField is_preferred "Is preferred address on interface"
        BooleanField is_primary "Is primary address on interface"
        BooleanField is_secondary "Is secondary address on interface"
        BooleanField is_standby "Is standby address on interface"
    }
    ipam_ipaddressrange {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name "Name of the IP Address Range"
        VarbinaryIPField start_host "First IP host address in the range (inclusive)"
        VarbinaryIPField end_host "Last IP host address in the range (inclusive)"
        IPAddressRangeIpVersion ip_version
        ForeignKey parent_id FK "The parent Prefix of this IP Address Range. Auto-resolved from the start/end host."
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
        CharField description
        BooleanField count_as_utilized "Forces this range to count as fully utilized in prefix utilization calculations."
        BooleanField is_exclusive "Prevent individual IP Address objects from being created within this range."
    }
    ipam_vlangroup {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ForeignKey location_id FK
        CharField description
        PositiveRangeNumberTextField range "Permitted VID range(s) as comma-separated list, default '1-4094' if left blank."
    }
    ipam_vlan {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey vlan_group_id FK
        PositiveSmallIntegerField vid
        CharField name
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
        CharField description
    }
    ipam_vlanlocationassignment {
        UUIDField id PK
        ForeignKey vlan_id FK
        ForeignKey location_id FK
    }
    ipam_service {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey device_id FK
        ForeignKey virtual_machine_id FK
        CharField name
        ServiceProtocol protocol
        JSONArrayField ports
        CharField description
    }
    load_balancers_virtualserver {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey vip_id FK
        CharField name
        PositiveIntegerField port
        VirtualServerProtocol protocol
        ForeignKey source_nat_pool_id FK
        VirtualServerSourceNatType source_nat_type
        VirtualServerLoadBalancerType load_balancer_type
        BooleanField enabled
        BooleanField ssl_offload
        ForeignKey device_id FK
        ForeignKey device_redundancy_group_id FK
        ForeignKey cloud_service_id FK
        ForeignKey virtual_chassis_id FK
        ForeignKey tenant_id FK
        ForeignKey load_balancer_pool_id FK
        ForeignKey health_check_monitor_id FK
    }
    load_balancers_loadbalancerpool {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        LoadBalancerPoolLoadBalancingAlgorithm load_balancing_algorithm
        ForeignKey health_check_monitor_id FK
        ForeignKey tenant_id FK
    }
    load_balancers_loadbalancerpoolmember {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField label "Optional label for the load balancer pool member."
        ForeignKey ip_address_id FK
        ForeignKey load_balancer_pool_id FK
        PositiveIntegerField port
        BooleanField ssl_offload
        ForeignKey health_check_monitor_id FK
        ForeignKey tenant_id FK
        StatusField status_id FK
    }
    load_balancers_healthcheckmonitor {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        PositiveIntegerField interval
        PositiveIntegerField retry "Number of retries before marking as down"
        PositiveIntegerField timeout
        PositiveIntegerField port
        HealthCheckMonitorHealthCheckType health_check_type
        ForeignKey tenant_id FK
    }
    load_balancers_certificateprofile {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CertificateProfileCertificateType certificate_type
        CharField certificate_file_path
        CharField chain_file_path
        CharField key_file_path
        DateTimeField expiration_date
        CharField cipher
        ForeignKey tenant_id FK
    }
    load_balancers_virtualservercertificateprofileassignment {
        UUIDField id PK
        ForeignKey virtual_server_id FK
        ForeignKey certificate_profile_id FK
    }
    load_balancers_loadbalancerpoolmembercertificateprofileassi0862 {
        UUIDField id PK
        ForeignKey load_balancer_pool_member_id FK
        ForeignKey certificate_profile_id FK
    }
    tenancy_tenantgroup {
        UUIDField id PK
        TreeNodeForeignKey parent_id FK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    tenancy_tenant {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ForeignKey tenant_group_id FK
        CharField description
        TextField comments
    }
    auth_user {
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField username UK "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        CharField first_name
        CharField last_name
        EmailField email
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
        UUIDField id PK
        JSONField config_data
    }
    users_token {
        UUIDField id PK
        ForeignKey user_id FK
        DateTimeField created
        DateTimeField expires
        CharField key UK
        BooleanField write_enabled "Permit create/update/delete operations using this key"
        CharField description
    }
    users_objectpermission {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        CharField name UK
        CharField description
        BooleanField enabled
        JSONArrayField actions "The list of actions granted by this permission"
        JSONField constraints "Queryset filter matching the applicable objects of the selected type(s)"
    }
    virtualization_clustertype {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    virtualization_clustergroup {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
    }
    virtualization_cluster {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        ForeignKey cluster_type_id FK
        ForeignKey cluster_group_id FK
        ForeignKey tenant_id FK
        ForeignKey location_id FK
        TextField comments
    }
    virtualization_virtualmachine {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        JSONField local_config_context_data
        ForeignKeyWithAutoRelatedName local_config_context_schema_id FK "Optional schema to validate the structure of the data"
        ForeignKeyWithAutoRelatedName local_config_context_data_owner_content_type_id FK
        UUIDField local_config_context_data_owner_object_id
        ForeignKey cluster_id FK
        ForeignKey tenant_id FK
        ForeignKey platform_id FK
        CharField name
        StatusField status_id FK
        RoleField role_id FK
        ForeignKey primary_ip4_id FK
        ForeignKey primary_ip6_id FK
        PositiveSmallIntegerField vcpus
        PositiveIntegerField memory
        PositiveIntegerField disk
        TextField comments
        ForeignKey software_version_id FK "The software version installed on this virtual machine"
    }
    virtualization_vminterface {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        StatusField status_id FK
        BooleanField enabled
        MACAddressCharField mac_address
        PositiveIntegerField mtu
        RoleField role_id FK
        VMInterfaceMode mode
        ForeignKey parent_interface_id FK "Assigned parent interface"
        ForeignKey bridge_id FK "Assigned bridge interface"
        ForeignKey virtual_machine_id FK
        CharField name
        NaturalOrderingField _name
        CharField description
        ForeignKey untagged_vlan_id FK
        ForeignKey vrf_id FK
    }
    vpn_vpnprofile {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        RoleField role_id FK
        ForeignKey secrets_group_id FK
        BooleanField keepalive_enabled
        PositiveIntegerField keepalive_interval
        PositiveIntegerField keepalive_retries
        BooleanField nat_traversal
        JSONField extra_options "Additional options specific to the VPN technology and/or implementation"
        ForeignKey tenant_id FK
    }
    vpn_vpnphase1policy {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        VPNPhase1PolicyIkeVersion ike_version
        BooleanField aggressive_mode "Only applicable to IKEv1"
        JSONArrayField encryption_algorithm
        JSONArrayField integrity_algorithm
        JSONArrayField dh_group
        PositiveIntegerField lifetime_seconds
        PositiveIntegerField lifetime_kb
        VPNPhase1PolicyAuthenticationMethod authentication_method "PSK, RSA, ECDSA, Certificate"
        ForeignKey tenant_id FK
    }
    vpn_vpnphase2policy {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        JSONArrayField encryption_algorithm
        JSONArrayField integrity_algorithm
        JSONArrayField pfs_group
        PositiveIntegerField lifetime
        ForeignKey tenant_id FK
    }
    vpn_vpnprofilephase1policyassignment {
        UUIDField id PK
        ForeignKey vpn_profile_id FK
        ForeignKey vpn_phase1_policy_id FK
        PositiveIntegerField weight "Higher weights appear later in the list"
    }
    vpn_vpnprofilephase2policyassignment {
        UUIDField id PK
        ForeignKey vpn_profile_id FK
        ForeignKey vpn_phase2_policy_id FK
        PositiveIntegerField weight "Higher weights appear later in the list"
    }
    vpn_vpn {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        CharField vpn_id
        ForeignKey vpn_profile_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
        VPNServiceType service_type "Optional classification of this VPN service, for example IPSec or VXLAN-EVPN."
        StatusField status_id FK
        JSONField extra_attributes "Free-form scalar service metadata only; not for references to real Nautobot objects."
    }
    vpn_vpntunnel {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        CharField tunnel_id
        ForeignKey vpn_profile_id FK
        ForeignKey vpn_id FK "VPN to which this tunnel belongs"
        RoleField role_id FK
        StatusField status_id FK
        ForeignKey secrets_group_id FK
        VPNTunnelEncapsulation encapsulation
        ForeignKey endpoint_a_id FK "Tunnel termination A"
        ForeignKey endpoint_z_id FK "Tunnel termination Z"
        ForeignKey tenant_id FK
    }
    vpn_vpntunnelendpoint {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name
        ForeignKey device_id FK
        OneToOneField source_interface_id FK, UK
        ForeignKey source_ipaddress_id FK "Mutually Exclusive with Source FQDN."
        CharField source_fqdn "Mutually Exclusive with Source IP Address"
        OneToOneField tunnel_interface_id FK, UK
        ForeignKey vpn_profile_id FK
        RoleField role_id FK
        ForeignKey tenant_id FK
    }
    vpn_vpntermination {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        ForeignKey vpn_id FK
        ForeignKey vlan_id FK
        ForeignKey interface_id FK
        ForeignKey vm_interface_id FK
    }
    wireless_supporteddatarate {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        SupportedDataRateStandard standard
        PositiveIntegerField rate
        IntegerField mcs_index "The Modulation and Coding Scheme (MCS) index is a value used in wireless communications to define the modulation type, coding rate, and number of spatial streams used in a transmission."
    }
    wireless_radioprofile {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        RadioProfileFrequency frequency
        IntegerField tx_power_min
        IntegerField tx_power_max
        JSONArrayField channel_width
        JSONArrayField allowed_channel_list "Comma-separated list of allowed channel numbers for this radio profile."
        RadioProfileRegulatoryDomain regulatory_domain
        IntegerField rx_power_min
    }
    wireless_wirelessnetwork {
        UUIDField id PK
        DateTimeField created
        DateTimeField last_updated
        JSONField _custom_field_data
        CharField name UK
        CharField description
        CharField ssid
        WirelessNetworkMode mode
        BooleanField enabled
        WirelessNetworkAuthentication authentication
        ForeignKey secrets_group_id FK
        BooleanField hidden
        ForeignKey tenant_id FK
    }
    wireless_controllermanageddevicegroupwirelessnetworkassignment {
        UUIDField id PK
        ForeignKey controller_managed_device_group_id FK
        ForeignKey wireless_network_id FK
        ForeignKey vlan_id FK
    }
    wireless_controllermanageddevicegroupradioprofileassignment {
        UUIDField id PK
        ForeignKey controller_managed_device_group_id FK
        ForeignKey radio_profile_id FK
    }
    silk_request {
        CharField id PK
        CharField path
        TextField query_params
        TextField raw_body
        TextField body
        CharField method
        DateTimeField start_time
        CharField view_name
        DateTimeField end_time
        FloatField time_taken
        TextField encoded_headers
        FloatField meta_time
        IntegerField meta_num_queries
        FloatField meta_time_spent_queries
        TextField pyprofile
        FileField prof_file
        IntegerField num_sql_queries
    }
    silk_response {
        CharField id PK
        OneToOneField request_id FK, UK
        IntegerField status_code
        TextField raw_body
        TextField body
        TextField encoded_headers
    }
    silk_sqlquery {
        AutoField id PK
        TextField query
        DateTimeField start_time
        DateTimeField end_time
        FloatField time_taken
        IntegerField identifier
        ForeignKey request_id FK
        TextField traceback
        TextField analysis
    }
    silk_profile {
        AutoField id PK
        CharField name
        DateTimeField start_time
        DateTimeField end_time
        ForeignKey request_id FK
        FloatField time_taken
        CharField file_path
        IntegerField line_num
        IntegerField end_line_num
        CharField func_name
        BooleanField exception_raised
        BooleanField dynamic
    }
```
