<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * One transactional notification held for one recipient (or for the whole project) by a
 * send with `delivery.throttle`. While it is held, the same notification is not sent again.
 *
 * The claim is taken by inserting (or taking over an expired) row, so the unique constraint
 * decides the race between concurrent sends, like `IdempotencyKey`. It is held from the
 * moment the send is accepted, re-anchored to the real send time once SES has the message,
 * and released when the email never gets there (cancelled, or failed before sending).
 */
class ThrottleClaim extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'throttle_claims';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'projectId',
        'key',
        'recipientKey',
        'contactId',
        'emailId',
        'workflowId',
        'windowSeconds',
        'expiresAt',
    ];

    protected $casts = [
        'windowSeconds' => 'integer',
        'expiresAt' => 'datetime',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function contact(): BelongsTo
    {
        return $this->belongsTo(Contact::class, 'contactId');
    }

    public function workflow(): BelongsTo
    {
        return $this->belongsTo(Workflow::class, 'workflowId');
    }
}