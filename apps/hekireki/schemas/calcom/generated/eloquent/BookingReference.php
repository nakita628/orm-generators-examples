<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $type @zod.string.min(1)
 * @property string $uid @zod.string.min(1)
 */
class BookingReference extends Model
{
    protected $table = 'BookingReference';

    public $timestamps = false;

    protected $fillable = [
        'type',
        'uid',
        'meetingId',
        'thirdPartyRecurringEventId',
        'meetingPassword',
        'meetingUrl',
        'bookingId',
        'externalCalendarId',
        'deleted',
        'credentialId',
        'delegationCredentialId',
    ];

    protected $casts = [
        'bookingId' => 'integer',
        'deleted' => 'boolean',
        'credentialId' => 'integer',
    ];

    public function booking(): BelongsTo
    {
        return $this->belongsTo(Booking::class, 'bookingId');
    }

    public function credential(): BelongsTo
    {
        return $this->belongsTo(Credential::class, 'credentialId');
    }

    public function delegationCredential(): BelongsTo
    {
        return $this->belongsTo(DelegationCredential::class, 'delegationCredentialId');
    }
}