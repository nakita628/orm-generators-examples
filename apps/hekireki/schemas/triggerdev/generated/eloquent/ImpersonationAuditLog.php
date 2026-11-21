<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class ImpersonationAuditLog extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'ImpersonationAuditLog';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'action',
        'adminId',
        'targetId',
        'ipAddress',
    ];

    protected $casts = [
        'action' => ImpersonationAuditLogAction::class,
    ];

    /** The admin user who initiated/ended the impersonation */
    public function admin(): BelongsTo
    {
        return $this->belongsTo(User::class, 'adminId');
    }

    /** The user being impersonated */
    public function target(): BelongsTo
    {
        return $this->belongsTo(User::class, 'targetId');
    }
}