```mermaid
erDiagram
    openwisp_users_user ||--o{ django_admin_log : "(id) - (user_id)"
    django_content_type |o--o{ django_admin_log : "(id) - (content_type_id)"
    django_content_type ||--o{ auth_permission : "(id) - (content_type_id)"
    auth_group }o--o{ auth_permission : "(permissions) - (group_set)"
    openwisp_users_user ||--o{ account_emailaddress : "(id) - (user_id)"
    account_emailaddress ||--o{ account_emailconfirmation : "(id) - (email_address_id)"
    socialaccount_socialapp }o--o{ django_site : "(sites) - (socialapp_set)"
    openwisp_users_user ||--o{ socialaccount_socialaccount : "(id) - (user_id)"
    socialaccount_socialapp |o--o{ socialaccount_socialtoken : "(id) - (app_id)"
    socialaccount_socialaccount ||--o{ socialaccount_socialtoken : "(id) - (account_id)"
    openwisp_users_user }o--o{ auth_group : "(groups) - (user_set)"
    openwisp_users_user }o--o{ auth_permission : "(user_permissions) - (user_set)"
    openwisp_users_user ||--o{ openwisp_users_organizationuser : "(id) - (user_id)"
    openwisp_users_organization ||--o{ openwisp_users_organizationuser : "(id) - (organization_id)"
    openwisp_users_organizationuser ||--o| openwisp_users_organizationowner : "(id) - (organization_user_id)"
    openwisp_users_organization ||--o| openwisp_users_organizationowner : "(id) - (organization_id)"
    openwisp_users_user ||--o{ openwisp_users_organizationinvitation : "(id) - (invited_by_id)"
    openwisp_users_user |o--o{ openwisp_users_organizationinvitation : "(id) - (invitee_id)"
    openwisp_users_organization ||--o{ openwisp_users_organizationinvitation : "(id) - (organization_id)"
    openwisp_users_user ||--o| authtoken_token : "(id) - (user_id)"
    openwisp_users_organization ||--o{ radcheck : "(id) - (organization_id)"
    openwisp_users_user |o--o{ radcheck : "(id) - (user_id)"
    openwisp_users_organization ||--o{ radreply : "(id) - (organization_id)"
    openwisp_users_user |o--o{ radreply : "(id) - (user_id)"
    openwisp_users_organization ||--o{ radacct : "(id) - (organization_id)"
    openwisp_users_organization ||--o{ openwisp_radius_radiusgroup : "(id) - (organization_id)"
    openwisp_radius_radiusgroup |o--o{ radgroupcheck : "(id) - (group_id)"
    openwisp_radius_radiusgroup |o--o{ radgroupreply : "(id) - (group_id)"
    openwisp_users_user |o--o{ radusergroup : "(id) - (user_id)"
    openwisp_radius_radiusgroup |o--o{ radusergroup : "(id) - (group_id)"
    openwisp_users_organization ||--o{ radpostauth : "(id) - (organization_id)"
    openwisp_users_organization ||--o{ nas : "(id) - (organization_id)"
    openwisp_users_organization ||--o{ radbatch : "(id) - (organization_id)"
    openwisp_radius_radiusgroup |o--o{ radbatch : "(id) - (group_id)"
    radbatch }o--o{ openwisp_users_user : "(users) - (radius_batch)"
    openwisp_users_organization ||--o{ radiustoken : "(id) - (organization_id)"
    openwisp_users_user ||--o| radiustoken : "(id) - (user_id)"
    openwisp_users_organization ||--o| openwisp_radius_organizationradiussettings : "(id) - (organization_id)"
    openwisp_users_organization ||--o{ openwisp_radius_phonetoken : "(id) - (organization_id)"
    openwisp_users_user ||--o{ openwisp_radius_phonetoken : "(id) - (user_id)"
    openwisp_users_organization ||--o{ openwisp_radius_registereduser : "(id) - (organization_id)"
    openwisp_users_user ||--o{ openwisp_radius_registereduser : "(id) - (user_id)"
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
    django_site {
        AutoField id PK
        CharField domain UK
        CharField name
    }
    account_emailaddress {
        AutoField id PK
        ForeignKey user_id FK
        EmailField email
        BooleanField verified
        BooleanField primary
    }
    account_emailconfirmation {
        AutoField id PK
        ForeignKey email_address_id FK
        DateTimeField created
        DateTimeField sent
        CharField key UK
    }
    socialaccount_socialapp {
        AutoField id PK
        CharField provider
        CharField provider_id
        CharField name
        CharField client_id "App ID, or consumer key"
        CharField secret "API secret, client secret, or consumer secret"
        CharField key "Key"
        JSONField settings
    }
    socialaccount_socialaccount {
        AutoField id PK
        ForeignKey user_id FK
        CharField provider
        CharField uid
        DateTimeField last_login
        DateTimeField date_joined
        JSONField extra_data
    }
    socialaccount_socialtoken {
        AutoField id PK
        ForeignKey app_id FK
        ForeignKey account_id FK
        TextField token "'oauth_token' (OAuth1) or access token (OAuth2)"
        TextField token_secret "'oauth_token_secret' (OAuth1) or refresh token (OAuth2)"
        DateTimeField expires_at
    }
    openwisp_users_user {
        CharField password
        DateTimeField last_login
        BooleanField is_superuser "Designates that this user has all permissions without explicitly assigning them."
        CharField username UK "Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only."
        CharField first_name
        CharField last_name
        BooleanField is_staff "Designates whether the user can log into this admin site."
        BooleanField is_active "Designates whether this user should be treated as active. Unselect this instead of deleting accounts."
        DateTimeField date_joined
        UUIDField id PK
        EmailField email UK
        TextField bio
        URLField url
        CharField company
        CharField location
        PhoneNumberField phone_number UK
        DateField birth_date
        TextField notes "notes for internal usage"
        UserLanguage language
        DateField password_updated
        BooleanField password_based_token "Indicates whether the last authentication token was obtained using the local password. When false, the token came from an external method (eg: SSO, SAML) and password expiration is not enforced for it. None means no token has been issued for this user since this feature was introduced."
        DateField expiration_date "If set, the account will be deactivated after this date passes and the user will no longer be able to log in."
    }
    openwisp_users_organization {
        CharField name "The name of the organization"
        BooleanField is_active
        AutoCreatedField created
        AutoLastModifiedField modified
        SlugField slug UK "The name in all lowercase, suitable for URL identification"
        UUIDField id PK
        TextField description
        EmailField email
        URLField url
    }
    openwisp_users_organizationuser {
        AutoCreatedField created
        AutoLastModifiedField modified
        UUIDField id PK
        BooleanField is_admin
        ForeignKey user_id FK
        ForeignKey organization_id FK
    }
    openwisp_users_organizationowner {
        AutoCreatedField created
        AutoLastModifiedField modified
        UUIDField id PK
        OneToOneField organization_user_id FK, UK
        OneToOneField organization_id FK, UK
    }
    openwisp_users_organizationinvitation {
        AutoField id PK
        UUIDField guid
        CharField invitee_identifier "The contact identifier for the invitee, email, phone number, social media handle, etc."
        AutoCreatedField created
        AutoLastModifiedField modified
        ForeignKey invited_by_id FK
        ForeignKey invitee_id FK
        ForeignKey organization_id FK
    }
    authtoken_token {
        CharField key PK
        OneToOneField user_id FK, UK
        DateTimeField created
    }
    radcheck {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField username
        CharField value
        RadiusCheckOp op
        CharField attribute
        ForeignKey user_id FK
    }
    radreply {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField username
        CharField value
        RadiusReplyOp op
        CharField attribute
        ForeignKey user_id FK
    }
    radacct {
        ForeignKey organization_id FK
        CharField acctsessionid
        CharField acctuniqueid PK
        CharField username
        CharField groupname
        CharField realm
        GenericIPAddressField nasipaddress
        CharField nasportid
        CharField nasporttype
        DateTimeField acctstarttime
        DateTimeField acctupdatetime
        DateTimeField acctstoptime
        IntegerField acctinterval
        PositiveIntegerField acctsessiontime
        CharField acctauthentic
        CharField connectinfo_start
        CharField connectinfo_stop
        BigIntegerField acctinputoctets
        BigIntegerField acctoutputoctets
        CharField calledstationid
        CharField callingstationid
        CharField acctterminatecause
        CharField servicetype
        CharField framedprotocol
        GenericIPAddressField framedipaddress
        GenericIPAddressField framedipv6address
        CharField framedipv6prefix
        CharField framedinterfaceid
        CharField delegatedipv6prefix
    }
    openwisp_radius_radiusgroup {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField name UK
        CharField description
        BooleanField default "The default group is automatically assigned to new users; changing the default group has only effect on new users (existing users will keep being members of their current group)"
    }
    radgroupcheck {
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField groupname
        CharField attribute
        RadiusGroupCheckOp op
        CharField value
        ForeignKey group_id FK
    }
    radgroupreply {
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField groupname
        CharField attribute
        RadiusGroupReplyOp op
        CharField value
        ForeignKey group_id FK
    }
    radusergroup {
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField username
        CharField groupname
        IntegerField priority
        ForeignKey user_id FK
        ForeignKey group_id FK
    }
    radpostauth {
        ForeignKey organization_id FK
        UUIDField id PK
        CharField username
        CharField pass
        CharField reply
        CharField calledstationid
        CharField callingstationid
        DateTimeField authdate
    }
    nas {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        CharField nasname "NAS Name (or IP address)"
        CharField shortname
        NasType type
        PositiveIntegerField ports
        CharField secret "Shared Secret"
        CharField server
        CharField community
        CharField description
    }
    radbatch {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        RadiusBatchStrategy strategy "Import users from a CSV or generate using a prefix"
        RadiusBatchStatus status
        CharField name "A unique batch name"
        PrivateFileField csvfile "The csv file containing the user details to be uploaded"
        CharField prefix "Usernames generated will be of the format [prefix][number]"
        JSONField user_credentials
        DateField expiration_date "If left blank users will never expire"
        ForeignKey group_id FK
        TextField notes "internal notes"
    }
    radiustoken {
        ForeignKey organization_id FK
        AutoCreatedField created
        CharField key PK
        OneToOneField user_id FK, UK
        BooleanField can_auth "Enable the radius token to be used for freeradius authorization request"
        BooleanField password_based "Indicates whether the user's most recent authentication used the local password. When false, the radius token is treated as externally authenticated (eg: SSO, SAML), so the local password expiration policy does not apply to it."
    }
    openwisp_radius_organizationradiussettings {
        UUIDField id PK
        OneToOneField organization_id FK, UK
        KeyField token
        FallbackBooleanChoiceField sms_verification "Whether users who sign up should be required to verify their mobile phone number via SMS"
        FallbackBooleanChoiceField needs_identity_verification "Whether identity verification is required at the time of user registration"
        CharField sms_sender "alpha numeric identifier used as sender for SMS sent by this organization"
        FallbackTextField sms_message "SMS message template used for sending verification code. Must contain '{code}' placeholder for OTP value."
        FallbackPositiveIntegerField sms_cooldown "Time period a user will have to wait before requesting another SMS token (in seconds)."
        JSONField sms_meta_data "Additional configuration for SMS backend in JSON format (optional, leave blank if unsure)"
        FallbackTextField freeradius_allowed_hosts "Comma separated list of IP addresses allowed to access freeradius API"
        FallbackBooleanChoiceField coa_enabled "Whether RADIUS Change Of Authoization (CoA) is enabled"
        FallbackTextField allowed_mobile_prefixes "Comma separated list of international mobile prefixes allowed to register via the user registration API."
        OrganizationRadiusSettingsFirstName first_name "Whether this field should be disabled, allowed or mandatory in the user registration API."
        OrganizationRadiusSettingsLastName last_name "Whether this field should be disabled, allowed or mandatory in the user registration API."
        OrganizationRadiusSettingsLocation location "Whether this field should be disabled, allowed or mandatory in the user registration API."
        OrganizationRadiusSettingsBirthDate birth_date "Whether this field should be disabled, allowed or mandatory in the user registration API."
        FallbackBooleanChoiceField registration_enabled "Whether the registration API endpoint should be enabled or not"
        FallbackBooleanChoiceField saml_registration_enabled "Whether the registration using SAML should be enabled or not"
        FallbackBooleanChoiceField mac_addr_roaming_enabled "Whether the MAC address roaming should be enabled or not."
        FallbackBooleanChoiceField social_registration_enabled "Whether the registration using social applications should be enabled or not"
        URLField login_url "Enter the URL where users can log in to the wifi service"
        URLField status_url "Enter the URL where users can log out from the wifi service"
        FallbackCharField password_reset_url "Enter the URL where users can reset their password"
    }
    openwisp_radius_phonetoken {
        ForeignKey organization_id FK
        UUIDField id PK
        AutoCreatedField created
        AutoLastModifiedField modified
        ForeignKey user_id FK
        DateTimeField valid_until
        PositiveIntegerField attempts
        BooleanField verified
        CharField token
        GenericIPAddressField ip
        PhoneNumberField phone_number
    }
    openwisp_radius_registereduser {
        ForeignKey organization_id FK
        UUIDField id PK
        ForeignKey user_id FK
        RegisteredUserMethod method "users can sign up in different ways, some methods are valid as indirect identity verification (eg: mobile phone SIM card in most countries)"
        BooleanField is_verified "whether the user has completed any identity verification process sucessfully"
        AutoLastModifiedField modified
    }
```
