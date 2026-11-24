<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"])
 *
 * @property string $type [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema)
 */
class AdminTag extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'AdminTag';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'type',
        'description',
        'userId',
        'organisationId',
        'createdByUserId',
        'adminRuleId',
    ];

    protected $casts = [
        'userId' => 'integer',
        'createdByUserId' => 'integer',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }

    public function organisation(): BelongsTo
    {
        return $this->belongsTo(Organisation::class, 'organisationId');
    }

    public function createdByUser(): BelongsTo
    {
        return $this->belongsTo(User::class, 'createdByUserId');
    }

    public function adminRule(): BelongsTo
    {
        return $this->belongsTo(AdminRule::class, 'adminRuleId');
    }
}