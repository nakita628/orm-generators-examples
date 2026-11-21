```mermaid
erDiagram
    user ||--o{ account : "(id) - (user_id)"
    user ||--o{ organization : "(id) - (owner_id)"
    organization ||--o{ organization_role : "(id) - (organization_id)"
    organization ||--o{ member : "(id) - (organization_id)"
    user ||--o{ member : "(id) - (user_id)"
    organization ||--o{ invitation : "(id) - (organization_id)"
    user ||--o{ invitation : "(id) - (inviter_id)"
    user ||--o{ two_factor : "(id) - (user_id)"
    user ||--o{ passkey : "(id) - (user_id)"
    user ||--o{ apikey : "(id) - (reference_id)"
    organization ||--o{ ai : "(id) - (organizationId)"
    "ssh-key" |o--o{ application : "(sshKeyId) - (customGitSSHKeyId)"
    registry |o--o{ application : "(registryId) - (registryId)"
    registry |o--o{ application : "(registryId) - (rollbackRegistryId)"
    environment ||--o{ application : "(environmentId) - (environmentId)"
    github |o--o{ application : "(githubId) - (githubId)"
    gitlab |o--o{ application : "(gitlabId) - (gitlabId)"
    gitea |o--o{ application : "(giteaId) - (giteaId)"
    bitbucket |o--o{ application : "(bitbucketId) - (bitbucketId)"
    server |o--o{ application : "(serverId) - (serverId)"
    server |o--o{ application : "(serverId) - (buildServerId)"
    registry |o--o{ application : "(registryId) - (buildRegistryId)"
    organization |o--o{ audit_log : "(id) - (organization_id)"
    user |o--o{ audit_log : "(id) - (user_id)"
    destination ||--o{ backup : "(destinationId) - (destinationId)"
    compose |o--o{ backup : "(composeId) - (composeId)"
    postgres |o--o{ backup : "(postgresId) - (postgresId)"
    mariadb |o--o{ backup : "(mariadbId) - (mariadbId)"
    mysql |o--o{ backup : "(mysqlId) - (mysqlId)"
    mongo |o--o{ backup : "(mongoId) - (mongoId)"
    libsql |o--o{ backup : "(libsqlId) - (libsqlId)"
    user |o--o{ backup : "(id) - (userId)"
    git_provider ||--o{ bitbucket : "(gitProviderId) - (gitProviderId)"
    organization ||--o{ certificate : "(id) - (organizationId)"
    server |o--o{ certificate : "(serverId) - (serverId)"
    "ssh-key" |o--o{ compose : "(sshKeyId) - (customGitSSHKeyId)"
    environment ||--o{ compose : "(environmentId) - (environmentId)"
    github |o--o{ compose : "(githubId) - (githubId)"
    gitlab |o--o{ compose : "(gitlabId) - (gitlabId)"
    bitbucket |o--o{ compose : "(bitbucketId) - (bitbucketId)"
    gitea |o--o{ compose : "(giteaId) - (giteaId)"
    server |o--o{ compose : "(serverId) - (serverId)"
    application |o--o{ deployment : "(applicationId) - (applicationId)"
    compose |o--o{ deployment : "(composeId) - (composeId)"
    server |o--o{ deployment : "(serverId) - (serverId)"
    preview_deployments |o--o{ deployment : "(previewDeploymentId) - (previewDeploymentId)"
    schedule |o--o{ deployment : "(scheduleId) - (scheduleId)"
    backup |o--o{ deployment : "(backupId) - (backupId)"
    rollback |o--o{ deployment : "(rollbackId) - (rollbackId)"
    volume_backup |o--o{ deployment : "(volumeBackupId) - (volumeBackupId)"
    server |o--o{ deployment : "(serverId) - (buildServerId)"
    organization ||--o{ destination : "(id) - (organizationId)"
    organization ||--o{ dns_provider : "(id) - (organizationId)"
    compose |o--o{ domain : "(composeId) - (composeId)"
    application |o--o{ domain : "(applicationId) - (applicationId)"
    preview_deployments |o--o{ domain : "(previewDeploymentId) - (previewDeploymentId)"
    project ||--o{ environment : "(projectId) - (projectId)"
    sso_provider |o--o{ forward_auth_settings : "(provider_id) - (providerId)"
    server |o--o| forward_auth_settings : "(serverId) - (serverId)"
    organization ||--o{ git_provider : "(id) - (organizationId)"
    user ||--o{ git_provider : "(id) - (userId)"
    git_provider ||--o{ gitea : "(gitProviderId) - (gitProviderId)"
    git_provider ||--o{ github : "(gitProviderId) - (gitProviderId)"
    git_provider ||--o{ gitlab : "(gitProviderId) - (gitProviderId)"
    environment ||--o{ libsql : "(environmentId) - (environmentId)"
    server |o--o{ libsql : "(serverId) - (serverId)"
    environment ||--o{ mariadb : "(environmentId) - (environmentId)"
    server |o--o{ mariadb : "(serverId) - (serverId)"
    environment ||--o{ mongo : "(environmentId) - (environmentId)"
    server |o--o{ mongo : "(serverId) - (serverId)"
    application |o--o{ mount : "(applicationId) - (applicationId)"
    compose |o--o{ mount : "(composeId) - (composeId)"
    libsql |o--o{ mount : "(libsqlId) - (libsqlId)"
    mariadb |o--o{ mount : "(mariadbId) - (mariadbId)"
    mongo |o--o{ mount : "(mongoId) - (mongoId)"
    mysql |o--o{ mount : "(mysqlId) - (mysqlId)"
    postgres |o--o{ mount : "(postgresId) - (postgresId)"
    redis |o--o{ mount : "(redisId) - (redisId)"
    environment ||--o{ mysql : "(environmentId) - (environmentId)"
    server |o--o{ mysql : "(serverId) - (serverId)"
    organization ||--o{ network : "(id) - (organizationId)"
    server |o--o{ network : "(serverId) - (serverId)"
    slack |o--o{ notification : "(slackId) - (slackId)"
    telegram |o--o{ notification : "(telegramId) - (telegramId)"
    discord |o--o{ notification : "(discordId) - (discordId)"
    email |o--o{ notification : "(emailId) - (emailId)"
    resend |o--o{ notification : "(resendId) - (resendId)"
    gotify |o--o{ notification : "(gotifyId) - (gotifyId)"
    ntfy |o--o{ notification : "(ntfyId) - (ntfyId)"
    mattermost |o--o{ notification : "(mattermostId) - (mattermostId)"
    custom |o--o{ notification : "(customId) - (customId)"
    lark |o--o{ notification : "(larkId) - (larkId)"
    pushover |o--o{ notification : "(pushoverId) - (pushoverId)"
    teams |o--o{ notification : "(teamsId) - (teamsId)"
    organization ||--o{ notification : "(id) - (organizationId)"
    application |o--o{ patch : "(applicationId) - (applicationId)"
    compose |o--o{ patch : "(composeId) - (composeId)"
    application ||--o{ port : "(applicationId) - (applicationId)"
    environment ||--o{ postgres : "(environmentId) - (environmentId)"
    server |o--o{ postgres : "(serverId) - (serverId)"
    application ||--o{ preview_deployments : "(applicationId) - (applicationId)"
    domain |o--o{ preview_deployments : "(domainId) - (domainId)"
    organization ||--o{ project : "(id) - (organizationId)"
    application ||--o{ redirect : "(applicationId) - (applicationId)"
    environment ||--o{ redis : "(environmentId) - (environmentId)"
    server |o--o{ redis : "(serverId) - (serverId)"
    organization ||--o{ registry : "(id) - (organizationId)"
    deployment ||--o{ rollback : "(deploymentId) - (deploymentId)"
    organization ||--o{ sandbox : "(id) - (organizationId)"
    server |o--o{ sandbox : "(serverId) - (serverId)"
    application |o--o{ schedule : "(applicationId) - (applicationId)"
    compose |o--o{ schedule : "(composeId) - (composeId)"
    server |o--o{ schedule : "(serverId) - (serverId)"
    organization |o--o{ schedule : "(id) - (organizationId)"
    organization |o--o{ scim_provider : "(id) - (organization_id)"
    application ||--o{ security : "(applicationId) - (applicationId)"
    organization ||--o{ server : "(id) - (organizationId)"
    "ssh-key" |o--o{ server : "(sshKeyId) - (sshKeyId)"
    user ||--o{ session : "(id) - (user_id)"
    organization ||--o{ "ssh-key" : "(id) - (organizationId)"
    user |o--o{ sso_provider : "(id) - (user_id)"
    organization |o--o{ sso_provider : "(id) - (organization_id)"
    organization ||--o{ tag : "(id) - (organizationId)"
    project ||--o{ project_tag : "(projectId) - (projectId)"
    tag ||--o{ project_tag : "(tagId) - (tagId)"
    organization ||--o{ telemetryProvider : "(id) - (organizationId)"
    organization ||--o{ vault_provider : "(id) - (organizationId)"
    application |o--o{ volume_backup : "(applicationId) - (applicationId)"
    postgres |o--o{ volume_backup : "(postgresId) - (postgresId)"
    mariadb |o--o{ volume_backup : "(mariadbId) - (mariadbId)"
    mongo |o--o{ volume_backup : "(mongoId) - (mongoId)"
    mysql |o--o{ volume_backup : "(mysqlId) - (mysqlId)"
    redis |o--o{ volume_backup : "(redisId) - (redisId)"
    libsql |o--o{ volume_backup : "(libsqlId) - (libsqlId)"
    compose |o--o{ volume_backup : "(composeId) - (composeId)"
    destination ||--o{ volume_backup : "(destinationId) - (destinationId)"
    organization |o--o{ webServerSettings : "(id) - (vectorAgentOrganizationId)"
    rollback ||..o| deployment : "(deploymentId) - (deploymentId)"
    github ||..o| git_provider : "(gitProviderId) - (gitProviderId)"
    gitlab ||..o| git_provider : "(gitProviderId) - (gitProviderId)"
    bitbucket ||..o| git_provider : "(gitProviderId) - (gitProviderId)"
    gitea ||..o| git_provider : "(gitProviderId) - (gitProviderId)"
    account ||..o| user : "(user_id) - (id)"
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
        boolean is2FAEnabled
        timestamp created_at
        timestamp updated_at
        text resetPasswordToken
        text resetPasswordExpiresAt
        text confirmationToken
        text confirmationExpiresAt
    }
    verification {
        text id PK
        text identifier
        text value
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
    }
    organization {
        text id PK
        text name
        text slug
        text logo
        timestamp created_at
        text metadata
        text default_role
        text owner_id FK
    }
    organization_role {
        text id PK
        text organization_id FK
        text role
        text permission
        timestamp created_at
        timestamp updated_at
    }
    member {
        text id PK
        text organization_id FK
        text user_id FK
        text role
        timestamp created_at
        text team_id
        boolean is_default
        boolean canCreateProjects "Permissions"
        boolean canAccessToSSHKeys
        boolean canCreateServices
        boolean canDeleteProjects
        boolean canDeleteServices
        boolean canAccessToDocker
        boolean canAccessToAPI
        boolean canAccessToGitProviders
        boolean canAccessToTraefikFiles
        boolean canDeleteEnvironments
        boolean canCreateEnvironments
        text accesedProjects
        text accessedEnvironments
        text accesedServices
        text accessedGitProviders
        text accessedServers
    }
    invitation {
        text id PK
        text organization_id FK
        text email
        text role
        text status
        timestamp expires_at
        text inviter_id FK
        text team_id
        timestamp created_at
    }
    two_factor {
        text id PK
        text secret
        text backup_codes
        text user_id FK
        boolean verified
        integer failed_verification_count
        timestamp locked_until
    }
    passkey {
        text id PK
        text name
        text public_key
        text user_id FK
        text credential_id
        integer counter
        text device_type
        boolean backed_up
        text transports
        timestamp created_at
        text aaguid
    }
    apikey {
        text id PK
        text name
        text start
        text prefix
        text key
        text config_id
        text reference_id FK
        integer refill_interval
        integer refill_amount
        timestamp last_refill_at
        boolean enabled
        boolean rate_limit_enabled
        integer rate_limit_time_window
        integer rate_limit_max
        integer request_count
        integer remaining
        timestamp last_request
        timestamp expires_at
        timestamp created_at
        timestamp updated_at
        text permissions
        text metadata
    }
    ai {
        text aiId PK
        text name
        text apiUrl
        text apiKey
        text model
        boolean isEnabled
        text organizationId FK
        text createdAt
    }
    application {
        text applicationId PK
        text name
        text appName
        text description
        text env
        text previewEnv
        text watchPaths
        text previewBuildArgs
        text previewBuildSecrets
        text previewLabels
        text previewWildcard
        integer previewPort
        boolean previewHttps
        text previewPath
        certificateType certificateType
        text previewCustomCertResolver
        integer previewLimit
        boolean previewForwardAuthEnabled
        boolean isPreviewDeploymentsActive
        boolean previewRequireCollaboratorPermissions "Security: Require collaborator permissions for preview deployments"
        boolean rollbackActive
        text buildArgs
        text buildSecrets
        text memoryReservation
        text memoryLimit
        text cpuReservation
        text cpuLimit
        text title
        boolean enabled
        text subtitle
        text command
        text args
        text icon
        text refreshToken
        sourceType sourceType
        boolean cleanCache
        text repository "Github"
        text owner
        text branch
        text buildPath
        triggerType triggerType
        boolean autoDeploy
        integer gitlabProjectId "Gitlab"
        text gitlabRepository
        text gitlabOwner
        text gitlabBranch
        text gitlabBuildPath
        text gitlabPathNamespace
        text giteaRepository "Gitea"
        text giteaOwner
        text giteaBranch
        text giteaBuildPath
        text bitbucketRepository "Bitbucket"
        text bitbucketRepositorySlug
        text bitbucketOwner
        text bitbucketBranch
        text bitbucketBuildPath
        text username "Docker"
        text password
        text dockerImage
        text registryUrl
        text customGitUrl "Git"
        text customGitBranch
        text customGitBuildPath
        text customGitSSHKeyId FK
        boolean enableSubmodules
        text dockerfile
        text dockerContextPath
        text dockerBuildStage
        text dropBuildPath "Drop"
        json healthCheckSwarm "Docker swarm json"
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        applicationStatus applicationStatus
        buildType buildType
        text railpackVersion
        text herokuVersion
        text publishDirectory
        boolean isStaticSpa
        boolean createEnvFile
        text createdAt
        text registryId FK
        text rollbackRegistryId FK
        text environmentId FK
        text githubId FK
        text gitlabId FK
        text giteaId FK
        text bitbucketId FK
        text serverId FK
        text buildServerId FK
        text buildRegistryId FK
        text networkIds
        boolean detachDokployNetwork
    }
    audit_log {
        text id PK
        text organization_id FK
        text user_id FK
        text user_email
        text user_role
        text action
        text resource_type
        text resource_id
        text resource_name
        text metadata
        timestamp created_at
    }
    backup {
        text backupId PK
        text appName
        text schedule
        boolean enabled
        text database
        text prefix
        text customName
        text serviceName
        text destinationId FK
        integer keepLatestCount
        boolean includeEncryptionKey
        backupType backupType
        databaseType databaseType
        text composeId FK
        text postgresId FK
        text mariadbId FK
        text mysqlId FK
        text mongoId FK
        text libsqlId FK
        text userId FK
        jsonb metadata "Only for compose backups"
    }
    bitbucket {
        text bitbucketId PK
        text bitbucketUsername
        text bitbucketEmail
        text appPassword
        text apiToken
        text bitbucketWorkspaceName
        text gitProviderId FK
    }
    certificate {
        text certificateId PK
        text name
        text certificateData
        text privateKey
        text certificatePath
        boolean autoRenew
        text organizationId FK
        text serverId FK
    }
    compose {
        text composeId PK
        text name
        text appName
        text description
        text env
        text composeFile
        text refreshToken
        sourceTypeCompose sourceType
        composeType composeType
        text repository "Github"
        text owner
        text branch
        boolean autoDeploy
        integer gitlabProjectId "Gitlab"
        text gitlabRepository
        text gitlabOwner
        text gitlabBranch
        text gitlabPathNamespace
        text bitbucketRepository "Bitbucket"
        text bitbucketRepositorySlug
        text bitbucketOwner
        text bitbucketBranch
        text giteaRepository "Gitea"
        text giteaOwner
        text giteaBranch
        text customGitUrl "Git"
        text customGitBranch
        text customGitSSHKeyId FK
        text command
        boolean createEnvFile
        boolean enableSubmodules
        text composePath
        text suffix
        boolean randomize
        boolean isolatedDeployment
        boolean isolatedDeploymentsVolume "Keep this for backward compatibility since we will not add the prefix anymore to volumes"
        boolean pullImages
        triggerType triggerType
        applicationStatus composeStatus
        text icon
        text environmentId FK
        text createdAt
        text watchPaths
        text githubId FK
        text gitlabId FK
        text bitbucketId FK
        text giteaId FK
        text serverId FK
        jsonb serviceNetworks
    }
    deployment {
        text deploymentId PK, FK
        text title
        text description
        deploymentStatus status
        text logPath
        text pid
        text applicationId FK
        text composeId FK
        text serverId FK
        boolean isPreviewDeployment
        text previewDeploymentId FK
        text createdAt
        text startedAt
        text finishedAt
        text errorMessage
        text scheduleId FK
        text backupId FK
        text rollbackId FK
        text volumeBackupId FK
        text buildServerId FK
    }
    destination {
        text destinationId PK
        text name
        text provider
        text accessKey
        text secretAccessKey
        text bucket
        text region
        text endpoint
        text additionalFlags
        text organizationId FK
        timestamp createdAt
    }
    dns_provider {
        text dnsProviderId PK
        text name
        DnsProviderType providerType
        jsonb config
        text organizationId FK
        text createdAt
    }
    domain {
        text domainId PK
        text host
        boolean https
        integer port
        text customEntrypoint
        text path
        text serviceName
        domainType domainType
        serial uniqueConfigKey
        text createdAt
        text composeId FK
        text customCertResolver
        text applicationId FK
        text previewDeploymentId FK
        certificateType certificateType
        text internalPath
        boolean stripPath
        text middlewares
        boolean forwardAuthEnabled
        boolean enabled
    }
    environment {
        text environmentId PK
        text name
        text description
        text createdAt
        text env
        text projectId FK
        boolean isDefault
    }
    forward_auth_settings {
        text forwardAuthSettingsId PK
        text authDomain
        text baseDomain
        boolean https
        certificateType certificateType
        text customCertResolver
        text providerId FK
        text serverId FK
        text createdAt
    }
    git_provider {
        text gitProviderId PK, FK
        text name
        gitProviderType providerType
        text createdAt
        text organizationId FK
        text userId FK
        boolean sharedWithOrganization
    }
    gitea {
        text giteaId PK
        text giteaUrl
        text giteaInternalUrl
        text redirect_uri
        text client_id
        text client_secret
        text gitProviderId FK
        text access_token
        text refresh_token
        integer expires_at
        text scopes
        integer last_authenticated_at
    }
    github {
        text githubId PK
        text githubAppName
        integer githubAppId
        text githubClientId
        text githubClientSecret
        text githubInstallationId
        text githubPrivateKey
        text githubWebhookSecret
        text githubUrl
        text gitProviderId FK
    }
    gitlab {
        text gitlabId PK
        text gitlabUrl
        text gitlabInternalUrl
        text application_id
        text redirect_uri
        text secret
        text access_token
        text refresh_token
        text group_name
        integer expires_at
        text gitProviderId FK
    }
    libsql {
        text libsqlId PK
        text name
        text appName
        text description
        text databaseUser
        text databasePassword
        sqldNode sqldNode
        text sqldPrimaryUrl
        boolean enableNamespaces
        text dockerImage
        text command
        text env
        text memoryReservation "RESOURCES"
        text memoryLimit
        text cpuReservation
        text cpuLimit
        integer externalPort
        integer externalGRPCPort
        integer externalAdminPort
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        integer replicas
        text networkIds
        boolean detachDokployNetwork
        text createdAt
        text environmentId FK
        text serverId FK
    }
    mariadb {
        text mariadbId PK
        text name
        text appName
        text description
        text databaseName
        text databaseUser
        text databasePassword
        text rootPassword
        text dockerImage
        text command
        text args
        text env
        text memoryReservation "RESOURCES"
        text memoryLimit
        text cpuReservation
        text cpuLimit
        integer externalPort
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        text createdAt
        text environmentId FK
        text serverId FK
        text networkIds
        boolean detachDokployNetwork
    }
    mongo {
        text mongoId PK
        text name
        text appName
        text description
        text databaseUser
        text databasePassword
        text dockerImage
        text command
        text args
        text env
        text memoryReservation
        text memoryLimit
        text cpuReservation
        text cpuLimit
        integer externalPort
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        text createdAt
        text environmentId FK
        text serverId FK
        boolean replicaSets
        text networkIds
        boolean detachDokployNetwork
    }
    mount {
        text mountId PK
        mountType type
        text hostPath
        text volumeName
        text filePath
        text content
        serviceType serviceType
        text mountPath
        text applicationId FK
        text composeId FK
        text libsqlId FK
        text mariadbId FK
        text mongoId FK
        text mysqlId FK
        text postgresId FK
        text redisId FK
    }
    mysql {
        text mysqlId PK
        text name
        text appName
        text description
        text databaseName
        text databaseUser
        text databasePassword
        text rootPassword
        text dockerImage
        text command
        text args
        text env
        text memoryReservation
        text memoryLimit
        text cpuReservation
        text cpuLimit
        integer externalPort
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        text createdAt
        text environmentId FK
        text serverId FK
        text networkIds
        boolean detachDokployNetwork
    }
    network {
        text networkId PK
        text name
        text dockerId
        networkDriver driver
        boolean internal
        boolean attachable
        boolean enableIPv4
        boolean enableIPv6
        integer mtu
        jsonb ipam
        text createdAt
        text organizationId FK
        text serverId FK
    }
    notification {
        text notificationId PK
        text name
        boolean appDeploy
        boolean appBuildError
        boolean databaseBackup
        boolean volumeBackup
        boolean dokployRestart
        boolean dokployBackup
        boolean dockerCleanup
        boolean serverThreshold
        notificationType notificationType
        text createdAt
        text slackId FK
        text telegramId FK
        text discordId FK
        text emailId FK
        text resendId FK
        text gotifyId FK
        text ntfyId FK
        text mattermostId FK
        text customId FK
        text larkId FK
        text pushoverId FK
        text teamsId FK
        text organizationId FK
    }
    slack {
        text slackId PK
        text webhookUrl
        text channel
    }
    telegram {
        text telegramId PK
        text botToken
        text chatId
        text messageThreadId
    }
    discord {
        text discordId PK
        text webhookUrl
        boolean decoration
    }
    email {
        text emailId PK
        text smtpServer
        integer smtpPort
        text username
        text password
        text fromAddress
        text toAddress
    }
    resend {
        text resendId PK
        text apiKey
        text fromAddress
        text toAddress
    }
    gotify {
        text gotifyId PK
        text serverUrl
        text appToken
        integer priority
        boolean decoration
    }
    ntfy {
        text ntfyId PK
        text serverUrl
        text topic
        text accessToken
        integer priority
    }
    mattermost {
        text mattermostId PK
        text webhookUrl
        text channel
        text username
    }
    custom {
        text customId PK
        text endpoint
        jsonb headers
    }
    lark {
        text larkId PK
        text webhookUrl
    }
    pushover {
        text pushoverId PK
        text userKey
        text apiToken
        integer priority
        integer retry
        integer expire
    }
    teams {
        text teamsId PK
        text webhookUrl
    }
    patch {
        text patchId PK
        patchType type
        text filePath
        boolean enabled
        text content
        text createdAt
        text updatedAt
        text applicationId FK "Relations - one of these must be set"
        text composeId FK
    }
    port {
        text portId PK
        integer publishedPort
        publishModeType publishMode
        integer targetPort
        protocolType protocol
        text applicationId FK
    }
    postgres {
        text postgresId PK
        text name
        text appName
        text databaseName
        text databaseUser
        text databasePassword
        text description
        text dockerImage
        text command
        text args
        text env
        text memoryReservation
        integer externalPort
        text memoryLimit
        text cpuReservation
        text cpuLimit
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        text createdAt
        text environmentId FK
        text serverId FK
        text networkIds
        boolean detachDokployNetwork
    }
    preview_deployments {
        text previewDeploymentId PK
        text branch
        text pullRequestId
        text pullRequestNumber
        text pullRequestURL
        text pullRequestTitle
        text pullRequestCommentId
        applicationStatus previewStatus
        text appName
        text applicationId FK
        text domainId FK
        text createdAt
        text expiresAt
    }
    project {
        text projectId PK
        text name
        text description
        text createdAt
        text organizationId FK
        text env
    }
    redirect {
        text redirectId PK
        text regex
        text replacement
        boolean permanent
        serial uniqueConfigKey
        text createdAt
        text applicationId FK
    }
    redis {
        text redisId PK
        text name
        text appName
        text description
        text password
        text dockerImage
        text command
        text args
        text env
        text memoryReservation
        text memoryLimit
        text cpuReservation
        text cpuLimit
        integer externalPort
        text createdAt
        applicationStatus applicationStatus
        json healthCheckSwarm
        json restartPolicySwarm
        json placementSwarm
        json updateConfigSwarm
        json rollbackConfigSwarm
        json modeSwarm
        json labelsSwarm
        json networkSwarm
        bigint stopGracePeriodSwarm
        json endpointSpecSwarm
        json ulimitsSwarm
        integer replicas
        text environmentId FK
        text serverId FK
        text networkIds
        boolean detachDokployNetwork
    }
    registry {
        text registryId PK
        text registryName
        text imagePrefix
        text username
        text password
        text registryUrl
        text createdAt
        RegistryType selfHosted
        text organizationId FK
    }
    rollback {
        text rollbackId PK
        text deploymentId FK
        serial version
        text image
        text createdAt
        jsonb fullContext
    }
    sandbox {
        text sandboxId PK
        text name
        text organizationId FK
        text serverId FK
        text image
        text template
        text containerId
        sandboxStatus status
        real cpu
        integer memoryMb
        integer pidsLimit
        integer timeoutMs
        timestamptz expiresAt
        timestamptz deadlineAt
        timestamptz lastActivityAt
        sandboxNetworkMode networkMode
        text envVars
        text workdir
        text user
        text runtime
        text createdAt
        timestamptz killedAt
    }
    schedule {
        text scheduleId PK
        text name
        text description
        text cronExpression
        text appName
        text serviceName
        shellType shellType
        scheduleType scheduleType
        text command
        text script
        text applicationId FK
        text composeId FK
        text serverId FK
        text organizationId FK
        boolean enabled
        text timezone
        text createdAt
    }
    scim_provider {
        text id PK
        text provider_id
        text scim_token
        text organization_id FK
    }
    security {
        text securityId PK
        text username
        text password
        text createdAt
        text applicationId FK
    }
    server {
        text serverId PK
        text name
        text description
        text ipAddress
        integer port
        text username
        text appName
        boolean enableDockerCleanup
        text telemetryProviderIds
        integer buildsConcurrency
        text createdAt
        text organizationId FK
        serverStatus serverStatus
        serverType serverType
        text command
        text sshKeyId FK
        jsonb metricsConfig
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
        text impersonated_by
        text active_organization_id
    }
    "ssh-key" {
        text sshKeyId PK
        text privateKey
        text publicKey
        text name
        text description
        text createdAt
        text lastUsedAt
        text organizationId FK
    }
    sso_provider {
        text id PK
        text issuer
        text oidc_config
        text saml_config
        text provider_id
        text user_id FK
        text organization_id FK
        text domain
        boolean domain_verified
        timestamp created_at
    }
    tag {
        text tagId PK
        text name
        text color
        text createdAt
        text organizationId FK
    }
    project_tag {
        text id PK
        text projectId FK
        text tagId FK
    }
    telemetryProvider {
        text telemetryProviderId PK
        text name
        TelemetryProviderType providerType
        text signals
        text endpoint
        text apiKey
        text apiSecret
        jsonb extraConfig
        boolean enabled
        text createdAt
        text organizationId FK
    }
    user {
        text id PK, FK
        text firstName
        text lastName
        boolean isRegistered
        text expirationDate
        text createdAt
        timestamp created_at
        boolean two_factor_enabled "Auth"
        text email
        boolean email_verified
        text image
        boolean banned
        text ban_reason
        timestamp ban_expires
        timestamp updated_at
        text role "Admin"
        boolean enablePaidFeatures "Metrics"
        boolean allowImpersonation
        boolean enableEnterpriseFeatures "Enterprise / proprietary features"
        text licenseKey
        boolean isValidEnterpriseLicense
        text stripeCustomerId
        text stripeSubscriptionId
        integer serversQuantity
        boolean sendInvoiceNotifications
        boolean isEnterpriseCloud
        text trustedOrigins
        text bookmarkedTemplates
        timestamp onboardingCompletedAt
        text onboardingStep
        integer onboardingReturnEmails
        text trialPlanPromptSubscriptionId
        timestamp serverSetupFailedEmailAt
        text serverConnectFailure
    }
    vault_provider {
        text vaultProviderId PK
        text name
        VaultProviderType providerType
        jsonb config
        jsonb assignments
        text organizationId FK
        text createdAt
    }
    volume_backup {
        text volumeBackupId PK
        text name
        text volumeName
        text prefix
        serviceType serviceType
        text appName
        text serviceName
        boolean turnOff
        text cronExpression
        integer keepLatestCount
        boolean enabled
        text applicationId FK
        text postgresId FK
        text mariadbId FK
        text mongoId FK
        text mysqlId FK
        text redisId FK
        text libsqlId FK
        text composeId FK
        text createdAt
        text destinationId FK
    }
    webServerSettings {
        text id PK
        text serverIp "Web Server Configuration"
        certificateType certificateType
        boolean https
        text host
        text letsEncryptEmail
        text sshPrivateKey
        boolean enableDockerCleanup
        text logCleanupCron
        text telemetryProviderIds
        text vectorAgentOrganizationId FK
        jsonb metricsConfig "Metrics Configuration"
        jsonb whitelabelingConfig "Whitelabeling Configuration (Enterprise / Proprietary)"
        boolean remoteServersOnly "Deployment Configuration (self-hosted only)"
        integer buildsConcurrency "Concurrent builds on the local web server"
        boolean enforceSSO "Auth Configuration (self-hosted only)"
        boolean cleanupCacheApplications "Cache Cleanup Configuration"
        boolean cleanupCacheOnPreviews
        boolean cleanupCacheOnCompose
        timestamp created_at
        timestamp updated_at
    }
```
