<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string|null $rbacRoleId Optional RBAC role to assign on invite acceptance. When set, the
 * accept-invite flow calls the loaded RBAC plugin's setUserRole with
 * this id after creating the OrgMember. Null = legacy behaviour, the
 * runtime fallback derives the role from `role` above.
 *
 * Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
 * a separate schema (Drizzle, not Prisma) so we can't model the FK
 * here. Validation happens at write time (action) and read time
 * (acceptInvite).
 */
class OrgMemberInvite extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'OrgMemberInvite';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'token',
        'email',
        'role',
        'rbacRoleId',
        'organizationId',
        'inviterId',
    ];

    protected $attributes = [
        'role' => 'MEMBER',
    ];

    protected $casts = [
        'role' => OrgMemberRole::class,
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    public function inviter(): BelongsTo
    {
        return $this->belongsTo(User::class, 'inviterId');
    }
}