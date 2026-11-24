<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @zod.import(["import { ZAdminRuleDefinitionSchema, ZAdminRuleTriggerSchema } from '@documenso/lib/types/admin-rules';"])
 *
 * @property string $trigger [AdminRuleTrigger] @zod.custom.use(ZAdminRuleTriggerSchema)
 * @property mixed $definition [AdminRuleDefinition] @zod.custom.use(ZAdminRuleDefinitionSchema)
 * @property list<string> $exemptOrganisationIds Escape hatch: organisations this rule is skipped for. Only meaningful on triggers that have an
 * organisation
 */
class AdminRule extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'AdminRule';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'name',
        'description',
        'mode',
        'trigger',
        'definition',
        'priority',
        'stopOnMatch',
        'exemptOrganisationIds',
    ];

    protected $casts = [
        'mode' => AdminRuleMode::class,
        'definition' => 'array',
        'priority' => 'integer',
        'stopOnMatch' => 'boolean',
        'exemptOrganisationIds' => AsPrismaList::class,
    ];

    public function matches(): HasMany
    {
        return $this->hasMany(AdminRuleMatch::class, 'ruleId');
    }

    public function adminTags(): HasMany
    {
        return $this->hasMany(AdminTag::class, 'adminRuleId');
    }
}