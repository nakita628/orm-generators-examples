<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @zod.import(["import { ZAdminRuleContextSchema, ZAdminRuleActionRunsSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
 *
 * @property string $trigger [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
 * @property mixed $eventSnapshot [AdminRuleContext] @zod.custom.use(ZAdminRuleContextSchema)
 * @property mixed $actionsRun [AdminRuleActionRuns] @zod.custom.use(ZAdminRuleActionRunsSchema)
 */
class AdminRuleMatch extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'AdminRuleMatch';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'mode',
        'trigger',
        'organisationId',
        'envelopeId',
        'eventSnapshot',
        'actionsRun',
        'ruleId',
        'actorUserId',
    ];

    protected $casts = [
        'mode' => AdminRuleMode::class,
        'eventSnapshot' => 'array',
        'actionsRun' => 'array',
        'actorUserId' => 'integer',
    ];

    public function rule(): BelongsTo
    {
        return $this->belongsTo(AdminRule::class, 'ruleId');
    }

    public function actorUser(): BelongsTo
    {
        return $this->belongsTo(User::class, 'actorUserId');
    }
}