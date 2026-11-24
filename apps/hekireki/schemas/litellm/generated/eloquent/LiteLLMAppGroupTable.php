<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class LiteLLMAppGroupTable extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    protected $table = 'LiteLLM_AppGroupTable';

    protected $primaryKey = 'app_group_id';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'name',
        'description',
        'metadata',
        'budget_id',
        'spend',
        'spend_reset_window',
        'created_by',
        'updated_by',
    ];

    protected $attributes = [
        'spend' => 0.0,
    ];

    protected $casts = [
        'metadata' => 'array',
        'spend' => 'float',
    ];

    public function budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'budget_id');
    }

    public function apps(): HasMany
    {
        return $this->hasMany(LiteLLMAppTable::class, 'app_group_id');
    }

    public function customer_organizations(): HasMany
    {
        return $this->hasMany(LiteLLMAppGroupCustomerOrgTable::class, 'app_group_id');
    }
}