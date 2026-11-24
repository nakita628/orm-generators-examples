<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class LiteLLMAppTable extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    protected $table = 'LiteLLM_AppTable';

    protected $primaryKey = 'app_id';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'app_group_id',
        'resource_type',
        'resource_id',
        'alias',
        'spend',
        'spend_reset_window',
        'budget_id',
        'default_tier_budget_id',
        'default_customer_member_budget_id',
        'auto_create_customers',
        'customer_margin',
        'created_by',
        'updated_by',
    ];

    protected $attributes = [
        'spend' => 0.0,
        'auto_create_customers' => true,
    ];

    protected $casts = [
        'spend' => 'float',
        'auto_create_customers' => 'boolean',
        'customer_margin' => 'float',
    ];

    public function app_group(): BelongsTo
    {
        return $this->belongsTo(LiteLLMAppGroupTable::class, 'app_group_id');
    }

    public function budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'budget_id');
    }

    public function default_tier_budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'default_tier_budget_id');
    }

    public function default_customer_member_budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'default_customer_member_budget_id');
    }

    public function customer_organizations(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerOrgTable::class, 'app_id');
    }

    public function customer_members(): HasMany
    {
        return $this->hasMany(LiteLLMAppCustomerMemberTable::class, 'app_id');
    }
}