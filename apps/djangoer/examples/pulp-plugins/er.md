```mermaid
erDiagram
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    core_accesspolicy ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_accesspolicy ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_contentguard |o--o{ core_domain : "(pulp_id) - (default_content_guard_id)"
    core_domain ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_domain ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_remote |o--o{ core_alternatecontentsource : "(pulp_id) - (remote_id)"
    core_domain ||--o{ core_alternatecontentsource : "(pulp_id) - (pulp_domain_id)"
    core_alternatecontentsource ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_alternatecontentsource ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_alternatecontentsource |o--o{ core_alternatecontentsourcepath : "(pulp_id) - (alternate_content_source_id)"
    core_repository |o--o{ core_alternatecontentsourcepath : "(pulp_id) - (repository_id)"
    core_alternatecontentsourcepath ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_alternatecontentsourcepath ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_artifact : "(pulp_id) - (pulp_domain_id)"
    core_artifact ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_artifact ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_pulptemporaryfile : "(pulp_id) - (pulp_domain_id)"
    core_pulptemporaryfile ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_pulptemporaryfile ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_content : "(pulp_id) - (pulp_domain_id)"
    core_content ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_content ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_artifact |o--o{ core_contentartifact : "(pulp_id) - (artifact_id)"
    core_content ||--o{ core_contentartifact : "(pulp_id) - (content_id)"
    core_contentartifact ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_contentartifact ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_contentartifact ||--o{ core_remoteartifact : "(pulp_id) - (content_artifact_id)"
    core_remote ||--o{ core_remoteartifact : "(pulp_id) - (remote_id)"
    core_domain ||--o{ core_remoteartifact : "(pulp_id) - (pulp_domain_id)"
    core_remoteartifact ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_remoteartifact ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_signingservice ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_signingservice ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_signingservice ||--o| core_asciiarmoreddetachedsigningservice : "is a, by signingservice_ptr_id"
    core_asciiarmoreddetachedsigningservice ||..o{ core_userrole : "(signingservice_ptr_id) - (object_id)"
    core_asciiarmoreddetachedsigningservice ||..o{ core_grouprole : "(signingservice_ptr_id) - (object_id)"
    core_appstatus ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_appstatus ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_role }o--o{ auth_permission : "(permissions) - (role_set)"
    core_role ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_role ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    auth_user ||--o{ core_userrole : "(id) - (user_id)"
    core_role ||--o{ core_userrole : "(pulp_id) - (role_id)"
    django_content_type |o--o{ core_userrole : "(id) - (content_type_id)"
    core_domain |o--o{ core_userrole : "(pulp_id) - (domain_id)"
    core_userrole ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_userrole ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_role ||--o{ core_grouprole : "(pulp_id) - (role_id)"
    django_content_type |o--o{ core_grouprole : "(id) - (content_type_id)"
    core_domain |o--o{ core_grouprole : "(pulp_id) - (domain_id)"
    core_grouprole ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_grouprole ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_artifact ||--o{ core_profileartifact : "(pulp_id) - (artifact_id)"
    core_task ||--o{ core_profileartifact : "(pulp_id) - (task_id)"
    core_profileartifact ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_profileartifact ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_appstatus |o--o{ core_task : "(pulp_id) - (app_lock_id)"
    core_task |o--o{ core_task : "(pulp_id) - (parent_task_id)"
    core_taskgroup |o--o{ core_task : "(pulp_id) - (task_group_id)"
    core_domain ||--o{ core_task : "(pulp_id) - (pulp_domain_id)"
    core_task ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_task ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_taskgroup : "(pulp_id) - (pulp_domain_id)"
    core_taskgroup ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_taskgroup ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    django_content_type ||--o{ core_createdresource : "(id) - (content_type_id)"
    core_task ||--o{ core_createdresource : "(pulp_id) - (task_id)"
    core_createdresource ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_createdresource ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_task |o--o{ core_taskschedule : "(pulp_id) - (last_task_id)"
    core_domain ||--o{ core_taskschedule : "(pulp_id) - (pulp_domain_id)"
    core_taskschedule ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_taskschedule ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_remote |o--o{ core_repository : "(pulp_id) - (remote_id)"
    core_domain ||--o{ core_repository : "(pulp_id) - (pulp_domain_id)"
    core_repository ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_repository ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_remote : "(pulp_id) - (pulp_domain_id)"
    core_remote ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_remote ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_content ||--o{ core_repositorycontent : "(pulp_id) - (content_id)"
    core_repository ||--o{ core_repositorycontent : "(pulp_id) - (repository_id)"
    core_repositoryversion ||--o{ core_repositorycontent : "(pulp_id) - (version_added_id)"
    core_repositoryversion |o--o{ core_repositorycontent : "(pulp_id) - (version_removed_id)"
    core_repositorycontent ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_repositorycontent ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_repository ||--o{ core_repositoryversion : "(pulp_id) - (repository_id)"
    core_repositoryversion |o--o{ core_repositoryversion : "(pulp_id) - (base_version_id)"
    core_repositoryversion ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_repositoryversion ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_repositoryversion ||--o{ core_repositoryversioncontentdetails : "(pulp_id) - (repository_version_id)"
    core_task |o--o{ core_export : "(pulp_id) - (task_id)"
    core_exporter ||--o{ core_export : "(pulp_id) - (exporter_id)"
    core_domain ||--o{ core_export : "(pulp_id) - (pulp_domain_id)"
    core_export ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_export ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    django_content_type ||--o{ core_exportedresource : "(id) - (content_type_id)"
    core_export ||--o{ core_exportedresource : "(pulp_id) - (export_id)"
    core_exportedresource ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_exportedresource ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_exporter : "(pulp_id) - (pulp_domain_id)"
    core_exporter ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_exporter ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_export ||--o| core_filesystemexport : "is a, by export_ptr_id"
    core_filesystemexport ||..o{ core_userrole : "(export_ptr_id) - (object_id)"
    core_filesystemexport ||..o{ core_grouprole : "(export_ptr_id) - (object_id)"
    core_exporter ||--o| core_filesystemexporter : "is a, by exporter_ptr_id"
    core_filesystemexporter ||..o{ core_userrole : "(exporter_ptr_id) - (object_id)"
    core_filesystemexporter ||..o{ core_grouprole : "(exporter_ptr_id) - (object_id)"
    core_export ||--o| core_pulpexport : "is a, by export_ptr_id"
    core_pulpexport ||..o{ core_userrole : "(export_ptr_id) - (object_id)"
    core_pulpexport ||..o{ core_grouprole : "(export_ptr_id) - (object_id)"
    core_exporter ||--o| core_pulpexporter : "is a, by exporter_ptr_id"
    core_pulpexport |o--o{ core_pulpexporter : "(export_ptr_id) - (last_export_id)"
    core_pulpexporter }o--o{ core_repository : "(repositories) - (core_pulp_exporter)"
    core_pulpexporter ||..o{ core_userrole : "(exporter_ptr_id) - (object_id)"
    core_pulpexporter ||..o{ core_grouprole : "(exporter_ptr_id) - (object_id)"
    core_task ||--o{ core_import : "(pulp_id) - (task_id)"
    core_importer ||--o{ core_import : "(pulp_id) - (importer_id)"
    core_domain ||--o{ core_import : "(pulp_id) - (pulp_domain_id)"
    core_import ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_import ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_importer : "(pulp_id) - (pulp_domain_id)"
    core_importer ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_importer ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_importer ||--o| core_pulpimporter : "is a, by importer_ptr_id"
    core_pulpimporter ||..o{ core_userrole : "(importer_ptr_id) - (object_id)"
    core_pulpimporter ||..o{ core_grouprole : "(importer_ptr_id) - (object_id)"
    core_pulpimporter ||--o{ core_pulpimporterrepository : "(importer_ptr_id) - (pulp_importer_id)"
    core_repository ||--o{ core_pulpimporterrepository : "(pulp_id) - (repository_id)"
    core_pulpimporterrepository ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_pulpimporterrepository ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_import ||--o| core_pulpimport : "is a, by import_ptr_id"
    core_pulpimport ||..o{ core_userrole : "(import_ptr_id) - (object_id)"
    core_pulpimport ||..o{ core_grouprole : "(import_ptr_id) - (object_id)"
    core_repositoryversion ||--o{ core_publication : "(pulp_id) - (repository_version_id)"
    core_domain ||--o{ core_publication : "(pulp_id) - (pulp_domain_id)"
    core_publication ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_publication ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_contentartifact ||--o{ core_publishedartifact : "(pulp_id) - (content_artifact_id)"
    core_publication ||--o{ core_publishedartifact : "(pulp_id) - (publication_id)"
    core_publishedartifact ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_publishedartifact ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_content ||--o| core_publishedmetadata : "is a, by content_ptr_id"
    core_publication ||--o{ core_publishedmetadata : "(pulp_id) - (publication_id)"
    core_publishedmetadata ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_publishedmetadata ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_domain ||--o{ core_contentguard : "(pulp_id) - (pulp_domain_id)"
    core_contentguard ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_contentguard ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_contentguard ||--o| core_rbaccontentguard : "is a, by contentguard_ptr_id"
    core_rbaccontentguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    core_rbaccontentguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_contentguard ||--o| core_contentredirectcontentguard : "is a, by contentguard_ptr_id"
    core_contentredirectcontentguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    core_contentredirectcontentguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_contentguard ||--o| core_headercontentguard : "is a, by contentguard_ptr_id"
    core_headercontentguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    core_headercontentguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_contentguard ||--o| core_envvarheadercontentguard : "is a, by contentguard_ptr_id"
    core_envvarheadercontentguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    core_envvarheadercontentguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_contentguard ||--o| core_compositecontentguard : "is a, by contentguard_ptr_id"
    core_compositecontentguard }o--o{ core_contentguard : "(guards) - (_core_compositecontentguard_guards_+)"
    core_compositecontentguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    core_compositecontentguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_domain ||--o{ core_distribution : "(pulp_id) - (pulp_domain_id)"
    core_contentguard |o--o{ core_distribution : "(pulp_id) - (content_guard_id)"
    core_publication |o--o{ core_distribution : "(pulp_id) - (publication_id)"
    core_remote |o--o{ core_distribution : "(pulp_id) - (remote_id)"
    core_repository |o--o{ core_distribution : "(pulp_id) - (repository_id)"
    core_repositoryversion |o--o{ core_distribution : "(pulp_id) - (repository_version_id)"
    core_distribution ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_distribution ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_distribution ||--o| core_artifactdistribution : "is a, by distribution_ptr_id"
    core_artifactdistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    core_artifactdistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_distribution ||--o{ core_distributedpublication : "(pulp_id) - (distribution_id)"
    core_publication ||--o{ core_distributedpublication : "(pulp_id) - (publication_id)"
    core_distributedpublication ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_distributedpublication ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_upload : "(pulp_id) - (pulp_domain_id)"
    core_upload ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_upload ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_upload ||--o{ core_uploadchunk : "(pulp_id) - (upload_id)"
    core_uploadchunk ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_uploadchunk ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_content ||--o| core_vulnerabilityreport : "(pulp_id) - (content_id)"
    core_domain ||--o{ core_vulnerabilityreport : "(pulp_id) - (pulp_domain_id)"
    core_vulnerabilityreport }o--o{ core_repositoryversion : "(repo_versions) - (core_vulnerabilityreport)"
    core_vulnerabilityreport ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_vulnerabilityreport ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_task ||--o{ core_progressreport : "(pulp_id) - (task_id)"
    core_progressreport ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_progressreport ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_taskgroup ||--o{ core_groupprogressreport : "(pulp_id) - (task_group_id)"
    core_groupprogressreport ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_groupprogressreport ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ core_upstreampulp : "(pulp_id) - (pulp_domain_id)"
    core_upstreampulp ||..o{ core_userrole : "(pulp_id) - (object_id)"
    core_upstreampulp ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_content ||--o| core_openpgppublickey : "is a, by content_ptr_id"
    core_domain ||--o{ core_openpgppublickey : "(pulp_id) - (_pulp_domain_id)"
    core_openpgppublickey ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_openpgppublickey ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| core_openpgppublicsubkey : "is a, by content_ptr_id"
    core_openpgppublickey ||--o{ core_openpgppublicsubkey : "(content_ptr_id) - (public_key_id)"
    core_openpgppublicsubkey ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_openpgppublicsubkey ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| core_openpgpuserid : "is a, by content_ptr_id"
    core_openpgppublickey ||--o{ core_openpgpuserid : "(content_ptr_id) - (public_key_id)"
    core_openpgpuserid ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_openpgpuserid ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| core_openpgpuserattribute : "is a, by content_ptr_id"
    core_openpgppublickey ||--o{ core_openpgpuserattribute : "(content_ptr_id) - (public_key_id)"
    core_openpgpuserattribute ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_openpgpuserattribute ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| core_openpgpsignature : "is a, by content_ptr_id"
    core_content ||--o{ core_openpgpsignature : "(pulp_id) - (signed_content_id)"
    core_openpgpsignature ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    core_openpgpsignature ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_repository ||--o| core_openpgpkeyring : "is a, by repository_ptr_id"
    core_openpgpkeyring ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    core_openpgpkeyring ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    core_distribution ||--o| core_openpgpdistribution : "is a, by distribution_ptr_id"
    core_openpgpdistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    core_openpgpdistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_contentguard ||--o| certguard_rhsmcertguard : "is a, by contentguard_ptr_id"
    certguard_rhsmcertguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    certguard_rhsmcertguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_contentguard ||--o| certguard_x509certguard : "is a, by contentguard_ptr_id"
    certguard_x509certguard ||..o{ core_userrole : "(contentguard_ptr_id) - (object_id)"
    certguard_x509certguard ||..o{ core_grouprole : "(contentguard_ptr_id) - (object_id)"
    core_content ||--o| file_filecontent : "is a, by content_ptr_id"
    core_domain ||--o{ file_filecontent : "(pulp_id) - (_pulp_domain_id)"
    file_filecontent ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    file_filecontent ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_remote ||--o| file_fileremote : "is a, by remote_ptr_id"
    file_fileremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    file_fileremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_remote ||--o| file_filegitremote : "is a, by remote_ptr_id"
    file_filegitremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    file_filegitremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_repository ||--o| file_filerepository : "is a, by repository_ptr_id"
    file_filerepository ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    file_filerepository ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    core_publication ||--o| file_filepublication : "is a, by publication_ptr_id"
    file_filepublication ||..o{ core_userrole : "(publication_ptr_id) - (object_id)"
    file_filepublication ||..o{ core_grouprole : "(publication_ptr_id) - (object_id)"
    core_distribution ||--o| file_filedistribution : "is a, by distribution_ptr_id"
    file_filedistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    file_filedistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_alternatecontentsource ||--o| file_filealternatecontentsource : "is a, by alternatecontentsource_ptr_id"
    file_filealternatecontentsource ||..o{ core_userrole : "(alternatecontentsource_ptr_id) - (object_id)"
    file_filealternatecontentsource ||..o{ core_grouprole : "(alternatecontentsource_ptr_id) - (object_id)"
    core_content ||--o| container_blob : "is a, by content_ptr_id"
    core_domain ||--o{ container_blob : "(pulp_id) - (_pulp_domain_id)"
    container_blob ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    container_blob ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| container_manifest : "is a, by content_ptr_id"
    container_blob |o--o{ container_manifest : "(content_ptr_id) - (config_blob_id)"
    core_domain ||--o{ container_manifest : "(pulp_id) - (_pulp_domain_id)"
    container_manifest ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    container_manifest ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    container_manifest ||--o{ container_blobmanifest : "(content_ptr_id) - (manifest_id)"
    container_blob ||--o{ container_blobmanifest : "(content_ptr_id) - (manifest_blob_id)"
    container_manifest ||--o{ container_manifestlistmanifest : "(content_ptr_id) - (image_manifest_id)"
    container_manifest ||--o{ container_manifestlistmanifest : "(content_ptr_id) - (manifest_list_id)"
    core_content ||--o| container_tag : "is a, by content_ptr_id"
    container_manifest ||--o{ container_tag : "(content_ptr_id) - (tagged_manifest_id)"
    core_domain ||--o{ container_tag : "(pulp_id) - (_pulp_domain_id)"
    container_tag ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    container_tag ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| container_manifestsignature : "is a, by content_ptr_id"
    container_manifest ||--o{ container_manifestsignature : "(content_ptr_id) - (signed_manifest_id)"
    core_domain ||--o{ container_manifestsignature : "(pulp_id) - (_pulp_domain_id)"
    container_manifestsignature ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    container_manifestsignature ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_domain ||--o{ container_containernamespace : "(pulp_id) - (pulp_domain_id)"
    container_containernamespace ||..o{ core_userrole : "(pulp_id) - (object_id)"
    container_containernamespace ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_remote ||--o| container_containerremote : "is a, by remote_ptr_id"
    container_containerremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    container_containerremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_remote ||--o| container_containerpullthroughremote : "is a, by remote_ptr_id"
    container_containerpullthroughremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    container_containerpullthroughremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_signingservice ||--o| container_manifestsigningservice : "is a, by signingservice_ptr_id"
    container_manifestsigningservice ||..o{ core_userrole : "(signingservice_ptr_id) - (object_id)"
    container_manifestsigningservice ||..o{ core_grouprole : "(signingservice_ptr_id) - (object_id)"
    core_repository ||--o| container_containerrepository : "is a, by repository_ptr_id"
    container_manifestsigningservice |o--o{ container_containerrepository : "(signingservice_ptr_id) - (manifest_signing_service_id)"
    container_containerrepository }o--o{ container_blob : "(pending_blobs) - (container_containerrepository)"
    container_containerrepository }o--o{ container_manifest : "(pending_manifests) - (container_containerrepository)"
    container_containerrepository ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    container_containerrepository ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    core_repository ||--o| container_containerpushrepository : "is a, by repository_ptr_id"
    container_manifestsigningservice |o--o{ container_containerpushrepository : "(signingservice_ptr_id) - (manifest_signing_service_id)"
    container_containerpushrepository }o--o{ container_blob : "(pending_blobs) - (container_containerpushrepository)"
    container_containerpushrepository }o--o{ container_manifest : "(pending_manifests) - (container_containerpushrepository)"
    container_containerpushrepository ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    container_containerpushrepository ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    core_distribution ||--o| container_containerpullthroughdistribution : "is a, by distribution_ptr_id"
    container_containernamespace |o--o{ container_containerpullthroughdistribution : "(pulp_id) - (namespace_id)"
    container_containerpullthroughdistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    container_containerpullthroughdistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_distribution ||--o| container_containerdistribution : "is a, by distribution_ptr_id"
    container_containernamespace |o--o{ container_containerdistribution : "(pulp_id) - (namespace_id)"
    container_containerpullthroughdistribution |o--o{ container_containerdistribution : "(distribution_ptr_id) - (pull_through_distribution_id)"
    container_containerdistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    container_containerdistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_upload ||--o| container_upload : "is a, by upload_ptr_id"
    core_repository ||--o{ container_upload : "(pulp_id) - (repository_id)"
    core_artifact |o--o{ container_upload : "(pulp_id) - (artifact_id)"
    container_upload ||..o{ core_userrole : "(upload_ptr_id) - (object_id)"
    container_upload ||..o{ core_grouprole : "(upload_ptr_id) - (object_id)"
    core_content ||--o| ansible_role : "is a, by content_ptr_id"
    core_domain ||--o{ ansible_role : "(pulp_id) - (_pulp_domain_id)"
    ansible_role ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_role ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_domain ||--o{ ansible_collection : "(pulp_id) - (pulp_domain_id)"
    ansible_collection ||..o{ core_userrole : "(pulp_id) - (object_id)"
    ansible_collection ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_task ||--o| ansible_collectionimport : "(pulp_id) - (task_id)"
    core_domain ||--o{ ansible_ansiblenamespace : "(pulp_id) - (pulp_domain_id)"
    ansible_ansiblenamespace ||..o{ core_userrole : "(pulp_id) - (object_id)"
    ansible_ansiblenamespace ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_content ||--o| ansible_collectionversion : "is a, by content_ptr_id"
    ansible_collection ||--o{ ansible_collectionversion : "(pulp_id) - (collection_id)"
    core_domain ||--o{ ansible_collectionversion : "(pulp_id) - (_pulp_domain_id)"
    ansible_collectionversion ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_collectionversion ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| ansible_collectionversionmark : "is a, by content_ptr_id"
    ansible_collectionversion ||--o{ ansible_collectionversionmark : "(content_ptr_id) - (marked_collection_id)"
    core_domain ||--o{ ansible_collectionversionmark : "(pulp_id) - (_pulp_domain_id)"
    ansible_collectionversionmark ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_collectionversionmark ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| ansible_collectionversionsignature : "is a, by content_ptr_id"
    ansible_collectionversion ||--o{ ansible_collectionversionsignature : "(content_ptr_id) - (signed_collection_id)"
    core_signingservice |o--o{ ansible_collectionversionsignature : "(pulp_id) - (signing_service_id)"
    core_domain ||--o{ ansible_collectionversionsignature : "(pulp_id) - (_pulp_domain_id)"
    ansible_collectionversionsignature ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_collectionversionsignature ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| ansible_ansiblenamespacemetadata : "is a, by content_ptr_id"
    ansible_ansiblenamespace ||--o{ ansible_ansiblenamespacemetadata : "(pulp_id) - (namespace_id)"
    core_domain ||--o{ ansible_ansiblenamespacemetadata : "(pulp_id) - (_pulp_domain_id)"
    ansible_ansiblenamespacemetadata ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_ansiblenamespacemetadata ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o{ ansible_downloadlog : "(pulp_id) - (content_unit_id)"
    auth_user |o--o{ ansible_downloadlog : "(id) - (user_id)"
    core_repository ||--o{ ansible_downloadlog : "(pulp_id) - (repository_id)"
    core_repositoryversion |o--o{ ansible_downloadlog : "(pulp_id) - (repository_version_id)"
    ansible_downloadlog ||..o{ core_userrole : "(pulp_id) - (object_id)"
    ansible_downloadlog ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_domain ||--o{ ansible_collectiondownloadcount : "(pulp_id) - (pulp_domain_id)"
    ansible_collectiondownloadcount ||..o{ core_userrole : "(pulp_id) - (object_id)"
    ansible_collectiondownloadcount ||..o{ core_grouprole : "(pulp_id) - (object_id)"
    core_remote ||--o| ansible_roleremote : "is a, by remote_ptr_id"
    ansible_roleremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    ansible_roleremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_remote ||--o| ansible_collectionremote : "is a, by remote_ptr_id"
    ansible_collectionremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    ansible_collectionremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_remote ||--o| ansible_gitremote : "is a, by remote_ptr_id"
    ansible_gitremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    ansible_gitremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_content ||--o| ansible_ansiblecollectiondeprecated : "is a, by content_ptr_id"
    core_domain ||--o{ ansible_ansiblecollectiondeprecated : "(pulp_id) - (_pulp_domain_id)"
    ansible_ansiblecollectiondeprecated ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    ansible_ansiblecollectiondeprecated ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_repository ||--o| ansible_ansiblerepository : "is a, by repository_ptr_id"
    ansible_ansiblerepository ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    ansible_ansiblerepository ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    core_distribution ||--o| ansible_ansibledistribution : "is a, by distribution_ptr_id"
    ansible_ansibledistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    ansible_ansibledistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    ansible_ansiblerepository ||--o{ ansible_crossrepositorycollectionversionindex : "(repository_ptr_id) - (repository_id)"
    core_repositoryversion |o--o{ ansible_crossrepositorycollectionversionindex : "(pulp_id) - (repository_version_id)"
    ansible_collectionversion ||--o{ ansible_crossrepositorycollectionversionindex : "(content_ptr_id) - (collection_version_id)"
    ansible_ansiblenamespacemetadata |o--o{ ansible_crossrepositorycollectionversionindex : "(content_ptr_id) - (namespace_metadata_id)"
    core_distribution ||--o| python_pythondistribution : "is a, by distribution_ptr_id"
    python_pythondistribution ||..o{ core_userrole : "(distribution_ptr_id) - (object_id)"
    python_pythondistribution ||..o{ core_grouprole : "(distribution_ptr_id) - (object_id)"
    core_content ||--o| python_pythonpackagecontent : "is a, by content_ptr_id"
    core_domain ||--o{ python_pythonpackagecontent : "(pulp_id) - (_pulp_domain_id)"
    python_pythonpackagecontent ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    python_pythonpackagecontent ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| python_packageprovenance : "is a, by content_ptr_id"
    python_pythonpackagecontent ||--o{ python_packageprovenance : "(content_ptr_id) - (package_id)"
    core_domain ||--o{ python_packageprovenance : "(pulp_id) - (_pulp_domain_id)"
    python_packageprovenance ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    python_packageprovenance ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_content ||--o| python_packageyank : "is a, by content_ptr_id"
    core_domain ||--o{ python_packageyank : "(pulp_id) - (_pulp_domain_id)"
    python_packageyank ||..o{ core_userrole : "(content_ptr_id) - (object_id)"
    python_packageyank ||..o{ core_grouprole : "(content_ptr_id) - (object_id)"
    core_publication ||--o| python_pythonpublication : "is a, by publication_ptr_id"
    python_pythonpublication ||..o{ core_userrole : "(publication_ptr_id) - (object_id)"
    python_pythonpublication ||..o{ core_grouprole : "(publication_ptr_id) - (object_id)"
    core_remote ||--o| python_pythonremote : "is a, by remote_ptr_id"
    python_pythonremote ||..o{ core_userrole : "(remote_ptr_id) - (object_id)"
    python_pythonremote ||..o{ core_grouprole : "(remote_ptr_id) - (object_id)"
    core_repository ||--o| python_pythonrepository : "is a, by repository_ptr_id"
    python_pythonrepository ||..o{ core_userrole : "(repository_ptr_id) - (object_id)"
    python_pythonrepository ||..o{ core_grouprole : "(repository_ptr_id) - (object_id)"
    python_pythonrepository ||--o{ python_pythonblocklistentry : "(repository_ptr_id) - (repository_id)"
    python_pythonblocklistentry ||..o{ core_userrole : "(pulp_id) - (object_id)"
    python_pythonblocklistentry ||..o{ core_grouprole : "(pulp_id) - (object_id)"
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
        CharField first_name
        CharField last_name
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
    core_accesspolicy {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        JSONField creation_hooks
        JSONField statements
        TextField viewset_name UK
        BooleanField customized
        JSONField queryset_scoping
    }
    core_domain {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        SlugField name UK
        HStoreField pulp_labels
        TextField description
        TextField storage_class
        EncryptedJSONField storage_settings
        BooleanField redirect_to_object_storage
        BooleanField hide_guarded_distributions
        ForeignKey default_content_guard_id FK
    }
    core_alternatecontentsource {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        DateTimeField last_refreshed
        ForeignKey remote_id FK
        ForeignKey pulp_domain_id FK
    }
    core_alternatecontentsourcepath {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey alternate_content_source_id FK
        TextField path
        ForeignKey repository_id FK
    }
    core_artifact {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ArtifactFileField file
        BigIntegerField size
        CharField md5
        CharField sha1
        CharField sha224
        CharField sha256
        CharField sha384
        CharField sha512
        DateTimeField timestamp_of_interest
        ForeignKey pulp_domain_id FK
    }
    core_pulptemporaryfile {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ArtifactFileField file
        ForeignKey pulp_domain_id FK
    }
    core_content {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        UUIDField upstream_id
        HStoreField pulp_labels
        DateTimeField timestamp_of_interest
        ForeignKey pulp_domain_id FK
    }
    core_contentartifact {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey artifact_id FK
        ForeignKey content_id FK
        TextField relative_path
    }
    core_remoteartifact {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField url
        BigIntegerField size
        CharField md5
        CharField sha1
        CharField sha224
        CharField sha256
        CharField sha384
        CharField sha512
        DateTimeField failed_at
        ForeignKey content_artifact_id FK
        ForeignKey remote_id FK
        ForeignKey pulp_domain_id FK
    }
    core_signingservice {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name UK
        TextField public_key
        TextField pubkey_fingerprint
        TextField script
    }
    core_asciiarmoreddetachedsigningservice {
        OneToOneField signingservice_ptr_id PK
    }
    core_appstatus {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        AppStatusAppType app_type
        TextField name
        HStoreField versions
        DurationField ttl
        DateTimeField last_heartbeat
    }
    core_role {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name UK
        TextField description
        BooleanField locked
    }
    core_userrole {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey user_id FK
        ForeignKey role_id FK
        ForeignKey content_type_id FK
        TextField object_id
        ForeignKey domain_id FK
    }
    core_grouprole {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey group_id FK
        ForeignKey role_id FK
        ForeignKey content_type_id FK
        TextField object_id
        ForeignKey domain_id FK
    }
    core_profileartifact {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey artifact_id FK
        ForeignKey task_id FK
        TextField name
    }
    core_task {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TaskState state
        TextField name
        TextField logging_cid
        DateTimeField unblocked_at
        DateTimeField started_at
        DateTimeField finished_at
        JSONField error
        EncryptedJSONField enc_args
        EncryptedJSONField enc_kwargs
        ForeignKey app_lock_id FK
        ForeignKey parent_task_id FK
        ForeignKey task_group_id FK
        ArrayField reserved_resources_record
        ForeignKey pulp_domain_id FK
        HStoreField versions
        ArrayField profile_options
        BooleanField immediate
        BooleanField deferred
        JSONField result
        TextField pulp_api_version
    }
    core_taskgroup {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField description
        BooleanField all_tasks_dispatched
        ForeignKey pulp_domain_id FK
    }
    core_createdresource {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey content_type_id FK
        UUIDField object_id
        ForeignKey task_id FK
    }
    core_taskschedule {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name
        DateTimeField next_dispatch
        DurationField dispatch_interval
        TextField task_name
        EncryptedJSONField task_args
        EncryptedJSONField task_kwargs
        ForeignKey last_task_id FK
        ForeignKey pulp_domain_id FK
    }
    core_repository {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        HStoreField pulp_labels
        TextField description
        PositiveIntegerField next_version
        PositiveIntegerField retain_repo_versions
        PositiveIntegerField retain_checkpoints
        BooleanField user_hidden
        ForeignKey remote_id FK
        ForeignKey pulp_domain_id FK
    }
    core_remote {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        HStoreField pulp_labels
        TextField url
        TextField ca_cert
        TextField client_cert
        EncryptedTextField client_key
        BooleanField tls_validation
        EncryptedTextField username
        EncryptedTextField password
        TextField proxy_url
        EncryptedTextField proxy_username
        EncryptedTextField proxy_password
        PositiveIntegerField download_concurrency
        PositiveIntegerField max_retries
        RemotePolicy policy
        FloatField total_timeout
        FloatField connect_timeout
        FloatField sock_connect_timeout
        FloatField sock_read_timeout
        JSONField headers
        IntegerField rate_limit
        ForeignKey pulp_domain_id FK
    }
    core_repositorycontent {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey content_id FK
        ForeignKey repository_id FK
        ForeignKey version_added_id FK
        ForeignKey version_removed_id FK
    }
    core_repositoryversion {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey repository_id FK
        PositiveIntegerField number
        BooleanField complete
        ForeignKey base_version_id FK
        JSONField info
        ArrayField content_ids
    }
    core_repositoryversioncontentdetails {
        AutoField id PK
        RepositoryVersionContentDetailsCountType count_type
        TextField content_type
        ForeignKey repository_version_id FK
        IntegerField count
    }
    core_export {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        JSONField params
        ForeignKey task_id FK
        ForeignKey exporter_id FK
        ForeignKey pulp_domain_id FK
    }
    core_exportedresource {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey content_type_id FK
        UUIDField object_id
        ForeignKey export_id FK
    }
    core_exporter {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        ForeignKey pulp_domain_id FK
    }
    core_filesystemexport {
        OneToOneField export_ptr_id PK
    }
    core_filesystemexporter {
        OneToOneField exporter_ptr_id PK
        TextField path
        FilesystemExporterMethod method
    }
    core_pulpexport {
        OneToOneField export_ptr_id PK
        JSONField output_file_info
        JSONField toc_info
    }
    core_pulpexporter {
        OneToOneField exporter_ptr_id PK
        TextField path
        ForeignKey last_export_id FK
    }
    core_import {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        JSONField params
        ForeignKey task_id FK
        ForeignKey importer_id FK
        ForeignKey pulp_domain_id FK
    }
    core_importer {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        ForeignKey pulp_domain_id FK
    }
    core_pulpimporter {
        OneToOneField importer_ptr_id PK
    }
    core_pulpimporterrepository {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField source_repo
        ForeignKey pulp_importer_id FK
        ForeignKey repository_id FK
    }
    core_pulpimport {
        OneToOneField import_ptr_id PK
    }
    core_publication {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        BooleanField complete
        BooleanField pass_through
        BooleanField checkpoint
        ForeignKey repository_version_id FK
        ForeignKey pulp_domain_id FK
    }
    core_publishedartifact {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField relative_path
        ForeignKey content_artifact_id FK
        ForeignKey publication_id FK
    }
    core_publishedmetadata {
        OneToOneField content_ptr_id PK
        TextField relative_path
        ForeignKey publication_id FK
    }
    core_contentguard {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        TextField description
        ForeignKey pulp_domain_id FK
    }
    core_rbaccontentguard {
        OneToOneField contentguard_ptr_id PK
    }
    core_contentredirectcontentguard {
        OneToOneField contentguard_ptr_id PK
        BinaryField shared_secret
    }
    core_headercontentguard {
        OneToOneField contentguard_ptr_id PK
        TextField header_name
        TextField header_value
        TextField jq_filter
    }
    core_envvarheadercontentguard {
        OneToOneField contentguard_ptr_id PK
        TextField header_name
        TextField env_var
    }
    core_compositecontentguard {
        OneToOneField contentguard_ptr_id PK
    }
    core_distribution {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField pulp_type
        TextField name
        HStoreField pulp_labels
        RelativePathField base_path
        ForeignKey pulp_domain_id FK
        BooleanField hidden
        BooleanField checkpoint
        ForeignKey content_guard_id FK
        ForeignKey publication_id FK
        ForeignKey remote_id FK
        ForeignKey repository_id FK
        ForeignKey repository_version_id FK
    }
    core_artifactdistribution {
        OneToOneField distribution_ptr_id PK
    }
    core_distributedpublication {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey distribution_id FK
        ForeignKey publication_id FK
        DateTimeField expires_at
    }
    core_systemid {
        UUIDField pulp_id PK
    }
    core_upload {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        BigIntegerField size
        ForeignKey pulp_domain_id FK
    }
    core_uploadchunk {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        FileField file
        ForeignKey upload_id FK
        BigIntegerField offset
        BigIntegerField size
    }
    core_vulnerabilityreport {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        OneToOneField content_id FK, UK
        JSONField vulns
        ForeignKey pulp_domain_id FK
    }
    core_progressreport {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField message
        TextField code
        ProgressReportState state
        IntegerField total
        IntegerField done
        ForeignKey task_id FK
        TextField suffix
    }
    core_groupprogressreport {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField message
        TextField code
        IntegerField total
        IntegerField done
        ForeignKey task_group_id FK
        TextField suffix
    }
    core_upstreampulp {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name
        ForeignKey pulp_domain_id FK
        TextField base_url
        TextField api_root
        TextField domain
        TextField ca_cert
        TextField client_cert
        EncryptedTextField client_key
        BooleanField tls_validation
        EncryptedTextField username
        EncryptedTextField password
        PositiveIntegerField download_concurrency
        PositiveIntegerField max_retries
        FloatField total_timeout
        FloatField connect_timeout
        FloatField sock_connect_timeout
        FloatField sock_read_timeout
        UpstreamPulpRemotePolicy remote_policy
        TextField q_select
        UpstreamPulpPolicy policy
        DateTimeField last_replication
    }
    core_openpgppublickey {
        OneToOneField content_ptr_id PK
        BinaryField raw_data
        CharField fingerprint
        DateTimeField created
        ForeignKey _pulp_domain_id FK
    }
    core_openpgppublicsubkey {
        OneToOneField content_ptr_id PK
        BinaryField raw_data
        ForeignKey public_key_id FK
        CharField fingerprint
        DateTimeField created
    }
    core_openpgpuserid {
        OneToOneField content_ptr_id PK
        BinaryField raw_data
        ForeignKey public_key_id FK
        CharField user_id
    }
    core_openpgpuserattribute {
        OneToOneField content_ptr_id PK
        BinaryField raw_data
        CharField sha256
        ForeignKey public_key_id FK
    }
    core_openpgpsignature {
        OneToOneField content_ptr_id PK
        BinaryField raw_data
        CharField sha256
        PositiveSmallIntegerField signature_type
        DateTimeField created
        DurationField expiration_time
        DurationField key_expiration_time
        CharField issuer
        CharField signers_user_id
        ForeignKey signed_content_id FK
    }
    core_openpgpkeyring {
        OneToOneField repository_ptr_id PK
    }
    core_openpgpdistribution {
        OneToOneField distribution_ptr_id PK
    }
    certguard_rhsmcertguard {
        OneToOneField contentguard_ptr_id PK
        TextField ca_certificate
    }
    certguard_x509certguard {
        OneToOneField contentguard_ptr_id PK
        TextField ca_certificate
    }
    file_filecontent {
        OneToOneField content_ptr_id PK
        TextField relative_path
        CharField digest
        ForeignKey _pulp_domain_id FK
    }
    file_fileremote {
        OneToOneField remote_ptr_id PK
    }
    file_filegitremote {
        OneToOneField remote_ptr_id PK
        TextField git_ref
    }
    file_filerepository {
        OneToOneField repository_ptr_id PK
        TextField manifest
        FileRepositorySha256Sums sha256sums
        BooleanField autopublish
        JSONField last_sync_details
    }
    file_filepublication {
        OneToOneField publication_ptr_id PK
        TextField manifest
        FilePublicationSha256Sums sha256sums
    }
    file_filedistribution {
        OneToOneField distribution_ptr_id PK
    }
    file_filealternatecontentsource {
        OneToOneField alternatecontentsource_ptr_id PK
    }
    container_blob {
        OneToOneField content_ptr_id PK
        TextField digest
        ForeignKey _pulp_domain_id FK
    }
    container_manifest {
        OneToOneField content_ptr_id PK
        TextField digest
        IntegerField schema_version
        ManifestMediaType media_type
        CharField type
        TextField data
        JSONField annotations
        JSONField labels
        TextField architecture
        TextField os
        BigIntegerField compressed_image_size
        BooleanField is_bootable
        BooleanField is_flatpak
        ForeignKey config_blob_id FK
        ForeignKey _pulp_domain_id FK
    }
    container_blobmanifest {
        AutoField id PK
        ForeignKey manifest_id FK
        ForeignKey manifest_blob_id FK
    }
    container_manifestlistmanifest {
        AutoField id PK
        TextField architecture
        TextField os
        TextField os_version
        TextField os_features
        TextField features
        TextField variant
        ForeignKey image_manifest_id FK
        ForeignKey manifest_list_id FK
    }
    container_tag {
        OneToOneField content_ptr_id PK
        TextField name
        ForeignKey tagged_manifest_id FK
        ForeignKey _pulp_domain_id FK
    }
    container_manifestsignature {
        OneToOneField content_ptr_id PK
        TextField name
        TextField digest
        ManifestSignatureType type
        TextField key_id
        TextField fingerprint
        PositiveIntegerField timestamp
        TextField creator
        TextField data
        ForeignKey signed_manifest_id FK
        ForeignKey _pulp_domain_id FK
    }
    container_containernamespace {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name
        ForeignKey pulp_domain_id FK
        HStoreField pulp_labels
    }
    container_containerremote {
        OneToOneField remote_ptr_id PK
        TextField upstream_name
        BooleanField include_foreign_layers
        ArrayField includes
        ArrayField excludes
        TextField sigstore
        ArrayField include_tags
        ArrayField exclude_tags
    }
    container_containerpullthroughremote {
        OneToOneField remote_ptr_id PK
        ArrayField includes
        ArrayField excludes
    }
    container_manifestsigningservice {
        OneToOneField signingservice_ptr_id PK
    }
    container_containerrepository {
        OneToOneField repository_ptr_id PK
        ForeignKey manifest_signing_service_id FK
    }
    container_containerpushrepository {
        OneToOneField repository_ptr_id PK
        ForeignKey manifest_signing_service_id FK
    }
    container_containerpullthroughdistribution {
        OneToOneField distribution_ptr_id PK
        ForeignKey namespace_id FK
        BooleanField private "Restrict pull access to explicitly authorized users. Related distributions inherit this value. Defaults to unrestricted pull access."
        TextField description
    }
    container_containerdistribution {
        OneToOneField distribution_ptr_id PK
        ForeignKey namespace_id FK
        BooleanField private "Restrict pull access to explicitly authorized users. Defaults to unrestricted pull access."
        TextField description
        ForeignKey pull_through_distribution_id FK
    }
    container_upload {
        OneToOneField upload_ptr_id PK
        ForeignKey repository_id FK
        ForeignKey artifact_id FK
    }
    ansible_role {
        OneToOneField content_ptr_id PK
        CharField namespace
        CharField name
        CharField version
        ForeignKey _pulp_domain_id FK
    }
    ansible_collection {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        CharField namespace
        CharField name
        ForeignKey pulp_domain_id FK
    }
    ansible_collectionimport {
        OneToOneField task_id PK
        JSONField messages
    }
    ansible_ansiblenamespace {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        CharField name
        ForeignKey pulp_domain_id FK
    }
    ansible_collectionversion {
        OneToOneField content_ptr_id PK
        ArrayField authors
        JSONField contents
        JSONField dependencies
        TextField description
        JSONField docs_blob
        JSONField manifest
        JSONField files
        CharField documentation
        CharField homepage
        CharField issues
        ArrayField license
        CharField name
        CharField namespace
        CharField repository
        CharField requires_ansible
        CharField sha256
        CharField version
        IntegerField version_major
        IntegerField version_minor
        IntegerField version_patch
        CharField version_prerelease
        ArrayField tags
        ForeignKey collection_id FK
        SearchVectorField search_vector
        ForeignKey _pulp_domain_id FK
    }
    ansible_collectionversionmark {
        OneToOneField content_ptr_id PK
        SlugField value
        ForeignKey marked_collection_id FK
        ForeignKey _pulp_domain_id FK
    }
    ansible_collectionversionsignature {
        OneToOneField content_ptr_id PK
        ForeignKey signed_collection_id FK
        TextField data
        CharField digest
        CharField pubkey_fingerprint
        ForeignKey signing_service_id FK
        ForeignKey _pulp_domain_id FK
    }
    ansible_ansiblenamespacemetadata {
        OneToOneField content_ptr_id PK
        CharField name
        CharField company
        CharField email
        CharField description
        TextField resources
        HStoreField links
        CharField avatar_sha256
        CharField metadata_sha256
        ForeignKey namespace_id FK
        ForeignKey _pulp_domain_id FK
    }
    ansible_downloadlog {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey content_unit_id FK
        ForeignKey user_id FK
        GenericIPAddressField ip
        JSONField extra_data
        TextField user_agent
        ForeignKey repository_id FK
        ForeignKey repository_version_id FK
    }
    ansible_collectiondownloadcount {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        ForeignKey pulp_domain_id FK
        CharField namespace
        CharField name
        BigIntegerField download_count
    }
    ansible_roleremote {
        OneToOneField remote_ptr_id PK
    }
    ansible_collectionremote {
        OneToOneField remote_ptr_id PK
        TextField requirements_file
        CharField auth_url
        EncryptedTextField token
        BooleanField sync_dependencies
        BooleanField signed_only
        PositiveIntegerField sync_highest_versions
    }
    ansible_gitremote {
        OneToOneField remote_ptr_id PK
        BooleanField metadata_only
        TextField git_ref
    }
    ansible_ansiblecollectiondeprecated {
        OneToOneField content_ptr_id PK
        CharField namespace
        CharField name
        ForeignKey _pulp_domain_id FK
    }
    ansible_ansiblerepository {
        OneToOneField repository_ptr_id PK
        DateTimeField last_synced_metadata_time
        TextField gpgkey
        BooleanField private
    }
    ansible_ansibledistribution {
        OneToOneField distribution_ptr_id PK
    }
    ansible_crossrepositorycollectionversionindex {
        AutoField id PK
        ForeignKey repository_id FK
        ForeignKey repository_version_id FK
        ForeignKey collection_version_id FK
        ForeignKey namespace_metadata_id FK
        BooleanField is_deprecated
        BooleanField is_signed
        BooleanField is_highest
    }
    python_pythondistribution {
        OneToOneField distribution_ptr_id PK
        BooleanField allow_uploads
    }
    python_pythonpackagecontent {
        OneToOneField content_ptr_id PK
        TextField author
        TextField author_email
        TextField description
        TextField home_page
        TextField keywords
        TextField license
        TextField metadata_version
        TextField name
        TextField platform
        TextField summary
        TextField version
        JSONField classifiers
        TextField download_url
        TextField supported_platform
        TextField maintainer
        TextField maintainer_email
        JSONField obsoletes_dist
        TextField project_url
        JSONField project_urls
        JSONField provides_dist
        JSONField requires_external
        JSONField requires_dist
        TextField requires_python
        TextField description_content_type
        JSONField provides_extras
        JSONField dynamic
        TextField license_expression
        JSONField license_file
        TextField name_normalized
        TextField filename
        PythonPackageContentPackagetype packagetype
        TextField python_version
        CharField sha256
        CharField metadata_sha256
        BigIntegerField size
        ForeignKey _pulp_domain_id FK
    }
    python_packageprovenance {
        OneToOneField content_ptr_id PK
        ForeignKey package_id FK
        JSONField provenance
        CharField sha256
        ForeignKey _pulp_domain_id FK
    }
    python_packageyank {
        OneToOneField content_ptr_id PK
        TextField name_normalized
        TextField version
        TextField yanked_reason
        ForeignKey _pulp_domain_id FK
    }
    python_pythonpublication {
        OneToOneField publication_ptr_id PK
    }
    python_pythonremote {
        OneToOneField remote_ptr_id PK
        BooleanField prereleases
        JSONField includes
        JSONField excludes
        PythonRemotePackageTypes package_types
        IntegerField keep_latest_packages
        PythonRemoteExcludePlatforms exclude_platforms
        BooleanField provenance
        BooleanField vulnerabilities
    }
    python_pythonrepository {
        OneToOneField repository_ptr_id PK
        BooleanField autopublish
        BooleanField allow_package_substitution
        BooleanField error_on_reject
    }
    python_pythonblocklistentry {
        UUIDField pulp_id PK
        DateTimeField pulp_created
        DateTimeField pulp_last_updated
        TextField name
        TextField version
        TextField filename
        TextField added_by
        ForeignKey repository_id FK
    }
```
