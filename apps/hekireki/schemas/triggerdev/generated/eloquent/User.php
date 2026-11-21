<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property string|null $pendingEmail Address awaiting confirmation; becomes `email` only after a link sent to it is opened
 * @property mixed $dashboardPreferences Preferences for the dashboard
 * @property bool $isOnCloudWaitlist @deprecated
 * @property bool $featureCloud @deprecated
 * @property bool $isOnHostedRepoWaitlist @deprecated
 * @property string|null $mfaLastUsedCode Hash of the last used code to prevent replay attacks
 * @property int $sessionDuration Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
 * May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
 * @property \Carbon\CarbonInterface|null $nextSessionEnd Absolute timestamp at which the user's current session must end. Written
 * at login (and any time the effective duration is recomputed) and shortened
 * when an admin lowers an org cap. Read on requireUser / getUser; null is
 * treated as "no enforced deadline" (legacy sessions from before this
 * shipped). Read-on-demand has zero per-request DB cost — the value
 * piggybacks on the User fetch that requireUser already does.
 */
class User extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'User';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'email',
        'pendingEmail',
        'authenticationMethod',
        'authenticationProfile',
        'authenticationExtraParams',
        'authIdentifier',
        'displayName',
        'name',
        'avatarUrl',
        'admin',
        'dashboardPreferences',
        'isOnCloudWaitlist',
        'featureCloud',
        'isOnHostedRepoWaitlist',
        'marketingEmails',
        'confirmedBasicDetails',
        'referralSource',
        'onboardingData',
        'mfaEnabledAt',
        'mfaSecretReferenceId',
        'mfaLastUsedCode',
        'sessionDuration',
        'nextSessionEnd',
        'invitationCodeId',
    ];

    protected $attributes = [
        'admin' => false,
        'isOnCloudWaitlist' => false,
        'featureCloud' => false,
        'isOnHostedRepoWaitlist' => false,
        'marketingEmails' => true,
        'confirmedBasicDetails' => false,
        'sessionDuration' => 31556952,
    ];

    protected $casts = [
        'authenticationMethod' => AuthenticationMethod::class,
        'authenticationProfile' => 'array',
        'authenticationExtraParams' => 'array',
        'admin' => 'boolean',
        'dashboardPreferences' => 'array',
        'isOnCloudWaitlist' => 'boolean',
        'featureCloud' => 'boolean',
        'isOnHostedRepoWaitlist' => 'boolean',
        'marketingEmails' => 'boolean',
        'confirmedBasicDetails' => 'boolean',
        'onboardingData' => 'array',
        'mfaEnabledAt' => 'datetime',
        'sessionDuration' => 'integer',
        'nextSessionEnd' => 'datetime',
    ];

    public function mfaSecretReference(): BelongsTo
    {
        return $this->belongsTo(SecretReference::class, 'mfaSecretReferenceId');
    }

    public function invitationCode(): BelongsTo
    {
        return $this->belongsTo(InvitationCode::class, 'invitationCodeId');
    }

    public function orgMemberships(): HasMany
    {
        return $this->hasMany(OrgMember::class, 'userId');
    }

    public function sentInvites(): HasMany
    {
        return $this->hasMany(OrgMemberInvite::class, 'inviterId');
    }

    public function personalAccessTokens(): HasMany
    {
        return $this->hasMany(PersonalAccessToken::class, 'userId');
    }

    public function createdApiKeys(): HasMany
    {
        return $this->hasMany(ApiKey::class, 'created_by_user_id');
    }

    public function deployments(): HasMany
    {
        return $this->hasMany(WorkerDeployment::class, 'triggeredById');
    }

    public function backupCodes(): HasMany
    {
        return $this->hasMany(MfaBackupCode::class, 'userId');
    }

    public function bulkActions(): HasMany
    {
        return $this->hasMany(BulkActionGroup::class, 'userId');
    }

    public function impersonationsPerformed(): HasMany
    {
        return $this->hasMany(ImpersonationAuditLog::class, 'adminId');
    }

    public function impersonationsReceived(): HasMany
    {
        return $this->hasMany(ImpersonationAuditLog::class, 'targetId');
    }

    public function supportAccessRequestsMade(): HasMany
    {
        return $this->hasMany(SupportAccessRequest::class, 'requestedById');
    }

    public function supportAccessRequestsApproved(): HasMany
    {
        return $this->hasMany(SupportAccessRequest::class, 'approvedById');
    }

    public function customerQueries(): HasMany
    {
        return $this->hasMany(CustomerQuery::class, 'userId');
    }

    public function metricsDashboards(): HasMany
    {
        return $this->hasMany(MetricsDashboard::class, 'ownerId');
    }

    public function platformNotifications(): HasMany
    {
        return $this->hasMany(PlatformNotification::class, 'userId');
    }

    public function platformNotificationInteractions(): HasMany
    {
        return $this->hasMany(PlatformNotificationInteraction::class, 'userId');
    }
}