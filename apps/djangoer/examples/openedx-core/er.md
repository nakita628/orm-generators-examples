```mermaid
erDiagram
    auth_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    auth_user }o--o{ auth_group : "(groups) - (user_set)"
    auth_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    openedx_content_learningpackage ||--o{ openedx_content_publishableentity : "(id) - (learning_package_id)"
    auth_user |o--o{ openedx_content_publishableentity : "(id) - (created_by_id)"
    openedx_content_publishableentity ||--o{ openedx_content_publishableentityversion : "(id) - (entity_id)"
    auth_user |o--o{ openedx_content_publishableentityversion : "(id) - (created_by_id)"
    openedx_content_publishableentityversion ||--o{ openedx_content_publishableentityversiondependency : "(id) - (referring_version_id)"
    openedx_content_publishableentity ||--o{ openedx_content_publishableentityversiondependency : "(id) - (referenced_entity_id)"
    openedx_content_publishableentity ||--o| openedx_content_draft : "(id) - (entity_id)"
    openedx_content_publishableentityversion |o--o| openedx_content_draft : "(id) - (version_id)"
    openedx_content_draftchangelogrecord |o--o{ openedx_content_draft : "(id) - (draft_log_record_id)"
    openedx_content_learningpackage ||--o{ openedx_content_draftchangelog : "(id) - (learning_package_id)"
    auth_user |o--o{ openedx_content_draftchangelog : "(id) - (changed_by_id)"
    openedx_content_draftchangelog ||--o{ openedx_content_draftchangelogrecord : "(id) - (draft_change_log_id)"
    openedx_content_publishableentity ||--o{ openedx_content_draftchangelogrecord : "(id) - (entity_id)"
    openedx_content_publishableentityversion |o--o{ openedx_content_draftchangelogrecord : "(id) - (old_version_id)"
    openedx_content_publishableentityversion |o--o{ openedx_content_draftchangelogrecord : "(id) - (new_version_id)"
    openedx_content_draftchangelogrecord ||--o{ openedx_content_draftsideeffect : "(id) - (cause_id)"
    openedx_content_draftchangelogrecord ||--o{ openedx_content_draftsideeffect : "(id) - (effect_id)"
    openedx_content_learningpackage ||--o{ openedx_content_publishlog : "(id) - (learning_package_id)"
    auth_user |o--o{ openedx_content_publishlog : "(id) - (published_by_id)"
    openedx_content_publishlog ||--o{ openedx_content_publishlogrecord : "(id) - (publish_log_id)"
    openedx_content_publishableentity ||--o{ openedx_content_publishlogrecord : "(id) - (entity_id)"
    openedx_content_publishableentityversion |o--o{ openedx_content_publishlogrecord : "(id) - (old_version_id)"
    openedx_content_publishableentityversion |o--o{ openedx_content_publishlogrecord : "(id) - (new_version_id)"
    openedx_content_publishableentity ||--o| openedx_content_published : "(id) - (entity_id)"
    openedx_content_publishableentityversion |o--o| openedx_content_published : "(id) - (version_id)"
    openedx_content_publishlogrecord ||--o{ openedx_content_published : "(id) - (publish_log_record_id)"
    openedx_content_publishlogrecord ||--o{ openedx_content_publishsideeffect : "(id) - (cause_id)"
    openedx_content_publishlogrecord ||--o{ openedx_content_publishsideeffect : "(id) - (effect_id)"
    openedx_content_learningpackage ||--o{ openedx_content_collection : "(id) - (learning_package_id)"
    auth_user |o--o{ openedx_content_collection : "(id) - (created_by_id)"
    openedx_content_collection ||--o{ openedx_content_collectionpublishableentity : "(id) - (collection_id)"
    openedx_content_publishableentity ||--o{ openedx_content_collectionpublishableentity : "(id) - (entity_id)"
    auth_user |o--o{ openedx_content_collectionpublishableentity : "(id) - (created_by_id)"
    openedx_content_learningpackage ||--o{ openedx_content_media : "(id) - (learning_package_id)"
    openedx_content_mediatype ||--o{ openedx_content_media : "(id) - (media_type_id)"
    openedx_content_publishableentity ||--o| openedx_content_component : "(id) - (publishable_entity_id)"
    openedx_content_learningpackage ||--o{ openedx_content_component : "(id) - (learning_package_id)"
    openedx_content_componenttype ||--o{ openedx_content_component : "(id) - (component_type_id)"
    openedx_content_publishableentityversion ||--o| openedx_content_componentversion : "(id) - (publishable_entity_version_id)"
    openedx_content_component ||--o{ openedx_content_componentversion : "(publishable_entity_id) - (component_id)"
    openedx_content_componentversion ||--o{ openedx_content_componentversionmedia : "(publishable_entity_version_id) - (component_version_id)"
    openedx_content_media ||--o{ openedx_content_componentversionmedia : "(id) - (media_id)"
    openedx_content_entitylist ||--o{ openedx_content_entitylistrow : "(id) - (entity_list_id)"
    openedx_content_publishableentity ||--o{ openedx_content_entitylistrow : "(id) - (entity_id)"
    openedx_content_publishableentityversion |o--o{ openedx_content_entitylistrow : "(id) - (entity_version_id)"
    openedx_content_publishableentity ||--o| openedx_content_container : "(id) - (publishable_entity_id)"
    openedx_content_learningpackage ||--o{ openedx_content_container : "(id) - (learning_package_id)"
    openedx_content_containertype ||--o{ openedx_content_container : "(id) - (container_type_id)"
    openedx_content_publishableentityversion ||--o| openedx_content_containerversion : "(id) - (publishable_entity_version_id)"
    openedx_content_container ||--o{ openedx_content_containerversion : "(publishable_entity_id) - (container_id)"
    openedx_content_entitylist ||--o{ openedx_content_containerversion : "(id) - (entity_list_id)"
    openedx_content_container ||--o| openedx_content_unit : "is a, by container_id"
    openedx_content_containerversion ||--o| openedx_content_unitversion : "is a, by container_version_id"
    openedx_content_container ||--o| openedx_content_subsection : "is a, by container_id"
    openedx_content_containerversion ||--o| openedx_content_subsectionversion : "is a, by container_version_id"
    openedx_content_container ||--o| openedx_content_section : "is a, by container_id"
    openedx_content_containerversion ||--o| openedx_content_sectionversion : "is a, by container_version_id"
    oel_publishing_publishableentity ||--o| oel_components_component : "(id) - (publishable_entity_id)"
    oel_publishing_publishableentity ||--o| oel_publishing_container : "(id) - (publishable_entity_id)"
    oel_tagging_taxonomy |o--o{ oel_tagging_tag : "(id) - (taxonomy_id)"
    oel_tagging_tag |o--o{ oel_tagging_tag : "(id) - (parent_id)"
    oel_tagging_taxonomy |o--o{ oel_tagging_objecttag : "(id) - (taxonomy_id)"
    oel_tagging_tag |o--o{ oel_tagging_objecttag : "(id) - (tag_id)"
    oel_tagging_taxonomy ||--o{ oel_tagging_tagimporttask : "(id) - (taxonomy_id)"
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
    openedx_content_learningpackage {
        IDField id PK
        UUIDField uuid UK
        MultiCollationCharField package_ref
        MultiCollationCharField title
        MultiCollationTextField description
        DateTimeField created
        DateTimeField updated
    }
    openedx_content_publishableentity {
        IDField id PK
        UUIDField uuid UK
        ForeignKey learning_package_id FK
        MultiCollationCharField entity_ref
        DateTimeField created
        ForeignKey created_by_id FK
        BooleanField can_stand_alone "Set to True when created independently, False when created as part of a container."
    }
    openedx_content_publishableentityversion {
        BigAutoField id PK
        UUIDField uuid UK
        ForeignKey entity_id FK
        MultiCollationCharField title
        PositiveIntegerField version_num
        DateTimeField created
        ForeignKey created_by_id FK
    }
    openedx_content_publishableentityversiondependency {
        BigAutoField id PK
        ForeignKey referring_version_id FK
        ForeignKey referenced_entity_id FK
    }
    openedx_content_draft {
        OneToOneField entity_id PK
        OneToOneField version_id FK, UK
        ForeignKey draft_log_record_id FK
    }
    openedx_content_draftchangelog {
        BigAutoField id PK
        UUIDField uuid UK
        ForeignKey learning_package_id FK
        DateTimeField changed_at
        ForeignKey changed_by_id FK
    }
    openedx_content_draftchangelogrecord {
        BigAutoField id PK
        ForeignKey draft_change_log_id FK
        ForeignKey entity_id FK
        ForeignKey old_version_id FK
        ForeignKey new_version_id FK
        CharField dependencies_hash_digest
    }
    openedx_content_draftsideeffect {
        BigAutoField id PK
        ForeignKey cause_id FK
        ForeignKey effect_id FK
    }
    openedx_content_publishlog {
        BigAutoField id PK
        UUIDField uuid UK
        ForeignKey learning_package_id FK
        MultiCollationCharField message
        DateTimeField published_at
        ForeignKey published_by_id FK
    }
    openedx_content_publishlogrecord {
        BigAutoField id PK
        ForeignKey publish_log_id FK
        ForeignKey entity_id FK
        ForeignKey old_version_id FK
        ForeignKey new_version_id FK
        CharField dependencies_hash_digest
        BooleanField direct
    }
    openedx_content_published {
        OneToOneField entity_id PK
        OneToOneField version_id FK, UK
        ForeignKey publish_log_record_id FK
    }
    openedx_content_publishsideeffect {
        BigAutoField id PK
        ForeignKey cause_id FK
        ForeignKey effect_id FK
    }
    openedx_content_collection {
        AutoField id PK
        ForeignKey learning_package_id FK
        MultiCollationCharField collection_code
        MultiCollationCharField title "The title of the collection."
        MultiCollationTextField description "Provides extra information for the user about this collection."
        BooleanField enabled "Disabled collections are 'soft deleted', and should be re-enabled before use, or be deleted."
        ForeignKey created_by_id FK
        DateTimeField created
        DateTimeField modified
    }
    openedx_content_collectionpublishableentity {
        BigAutoField id PK
        ForeignKey collection_id FK
        ForeignKey entity_id FK
        ForeignKey created_by_id FK
        DateTimeField created
    }
    openedx_content_mediatype {
        AutoField id PK
        MultiCollationCharField type
        MultiCollationCharField sub_type
        MultiCollationCharField suffix
    }
    openedx_content_media {
        IDField id PK
        ForeignKey learning_package_id FK
        ForeignKey media_type_id FK
        PositiveBigIntegerField size
        CharField hash_digest
        BooleanField has_file
        MultiCollationTextField text
        DateTimeField created
    }
    openedx_content_componenttype {
        AutoField id PK
        MultiCollationCharField namespace
        MultiCollationCharField name
    }
    openedx_content_component {
        OneToOneField publishable_entity_id PK
        ForeignKey learning_package_id FK
        ForeignKey component_type_id FK
        MultiCollationCharField component_code
    }
    openedx_content_componentversion {
        OneToOneField publishable_entity_version_id PK
        ForeignKey component_id FK
    }
    openedx_content_componentversionmedia {
        BigAutoField id PK
        ForeignKey component_version_id FK
        ForeignKey media_id FK
        MultiCollationCharField path
    }
    openedx_content_entitylist {
        BigAutoField id PK
    }
    openedx_content_entitylistrow {
        BigAutoField id PK
        ForeignKey entity_list_id FK
        PositiveIntegerField order_num
        ForeignKey entity_id FK
        ForeignKey entity_version_id FK
    }
    openedx_content_containertype {
        AutoField id PK
        MultiCollationCharField type_code UK
    }
    openedx_content_container {
        OneToOneField publishable_entity_id PK
        ForeignKey learning_package_id FK
        ForeignKey container_type_id FK
        MultiCollationCharField container_code
    }
    openedx_content_containerversion {
        OneToOneField publishable_entity_version_id PK
        ForeignKey container_id FK
        ForeignKey entity_list_id FK
    }
    openedx_content_unit {
        OneToOneField container_id PK
    }
    openedx_content_unitversion {
        OneToOneField container_version_id PK
    }
    openedx_content_subsection {
        OneToOneField container_id PK
    }
    openedx_content_subsectionversion {
        OneToOneField container_version_id PK
    }
    openedx_content_section {
        OneToOneField container_id PK
    }
    openedx_content_sectionversion {
        OneToOneField container_version_id PK
    }
    oel_collections_collection {
        AutoField id PK
    }
    oel_components_component {
        OneToOneField publishable_entity_id PK
    }
    oel_publishing_learningpackage {
        AutoField id PK
    }
    oel_publishing_publishableentity {
        BigAutoField id PK
    }
    oel_publishing_draftchangelog {
        BigAutoField id PK
    }
    oel_publishing_draftchangelogrecord {
        BigAutoField id PK
    }
    oel_publishing_container {
        OneToOneField publishable_entity_id PK
    }
    oel_tagging_tag {
        BigAutoField id PK
        ForeignKey taxonomy_id FK "Namespace and rules for using a given set of tags."
        ForeignKey parent_id FK "Tag that lives one level up from the current tag, forming a hierarchy."
        MultiCollationCharField value "Content of a given tag, occupying the 'value' part of the key:value pair."
        MultiCollationCharField external_id "Used to link an Open edX Tag with a tag in an externally-defined taxonomy."
        IntegerField depth "Number of ancestors this tag has. Zero for root tags, one for their children, and so on. Set automatically by save(); do not set manually."
        MultiCollationCharField lineage "Tab-separated ancestor path including this tag: 'Root\tParent\t...\tThisValue\t'. Used for depth-first tree ordering and descendant prefix matching. Set automatically by save(); do not set manually."
        BooleanField archived "Hides this tag from active use by default, without removing it."
    }
    oel_tagging_taxonomy {
        BigAutoField id PK
        MultiCollationCharField name "User-facing label used when applying tags from this taxonomy to Open edX objects."
        MultiCollationTextField description "Provides extra information for the user when applying tags from this taxonomy to an object."
        BooleanField enabled "Only enabled taxonomies will be shown to authors."
        BooleanField allow_multiple "Indicates that multiple tags from this taxonomy may be added to an object."
        BooleanField allow_free_text "Indicates that tags in this taxonomy need not be predefined; authors may enter their own tag values."
        BooleanField visible_to_authors "Indicates whether this taxonomy should be visible to object authors."
        CharField export_id UK "User-facing ID that is used on import/export. Should only contain alphanumeric characters or '_' '-' '.'"
        BooleanField read_only "Indicates that the tags in this taxonomy are maintained by the system or an external integration; taxonomy admins will not be permitted to add, edit, or delete its tags."
        BooleanField archived "Hides this taxonomy from active use by default, without removing it. Unlike disabling it with 'enabled', archiving marks the taxonomy as retired."
    }
    oel_tagging_objecttag {
        BigAutoField id PK
        MultiCollationCharField object_id "Identifier for the object being tagged"
        ForeignKey taxonomy_id FK "Taxonomy that this object tag belongs to. Used for validating the tag and provides the tag's 'name' if set."
        ForeignKey tag_id FK "Tag associated with this object tag. Provides the tag's 'value' if set."
        BooleanField is_copied "True if this object tag has been copied from one object to another using 'copy_tags' api function"
        MultiCollationCharField _export_id "User-facing label used for this tag, stored in case taxonomy is (or becomes) null. If the taxonomy field is set, then taxonomy.export_id takes precedence over this field."
        MultiCollationCharField _value "User-facing value used for this tag, stored in case tag is null, e.g if taxonomy is free text, or if it becomes null (e.g. if the Tag is deleted). If the tag field is set, then tag.value takes precedence over this field."
        BooleanField archived "Hides this association from active use by default, without removing it."
        BooleanField deletion_locked "Set when something outside the tagging app depends on this association; a delete archives it instead of removing it."
    }
    oel_tagging_tagimporttask {
        BigAutoField id PK
        ForeignKey taxonomy_id FK "Taxonomy associated with this import"
        TextField log "Action execution logs"
        TagImportTaskState status "Task status"
        DateTimeField creation_date
    }
```
