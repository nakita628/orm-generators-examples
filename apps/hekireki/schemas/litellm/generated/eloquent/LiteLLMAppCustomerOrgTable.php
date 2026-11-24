<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class LiteLLMAppCustomerOrgTable extends Model
{
    use PrismaDates;

    const KEY_COLUMNS = ['app_id', 'customer_org_id'];

    protected $table = 'LiteLLM_AppCustomerOrgTable';

    protected $primaryKey = null;

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'app_id',
        'customer_org_id',
        'spend',
        'spend_reset_window',
        'tier_budget_id',
        'budget_id',
        'budget_override_id',
        'budget_override_expires_at',
        'customer_member_budget_id',
    ];

    protected $attributes = [
        'spend' => 0.0,
    ];

    protected $casts = [
        'spend' => 'float',
        'budget_override_expires_at' => 'datetime',
    ];

    public function getKey()
    {
        $key = [];
        foreach (static::KEY_COLUMNS as $column) {
            $key[$column] = $this->getAttribute($column);
        }

        return $key;
    }

    public function newEloquentBuilder($query)
    {
        return new CompositeKeyBuilder($query);
    }

    public function newCollection(array $models = [])
    {
        return new CompositeKeyCollection($models);
    }

    public static function destroy($ids)
    {
        $count = 0;
        foreach ((new static())->newQuery()->whereKey(func_num_args() > 1 ? func_get_args() : $ids)->get() as $model) {
            if ($model->delete()) {
                $count++;
            }
        }

        return $count;
    }

    protected function setKeysForSaveQuery($query)
    {
        foreach (static::KEY_COLUMNS as $column) {
            $query->where($column, '=', $this->original[$column] ?? $this->getAttribute($column));
        }

        return $query;
    }

    protected function setKeysForSelectQuery($query)
    {
        foreach (static::KEY_COLUMNS as $column) {
            $query->where($column, '=', $this->original[$column] ?? $this->getAttribute($column));
        }

        return $query;
    }

    public function delete()
    {
        $this->mergeAttributesFromCachedCasts();

        if (! $this->exists) {
            return null;
        }

        if ($this->fireModelEvent('deleting') === false) {
            return false;
        }

        $this->touchOwners();
        $this->performDeleteOnModel();
        $this->fireModelEvent('deleted', false);

        return true;
    }

    public function budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'budget_id');
    }

    public function app(): BelongsTo
    {
        return $this->belongsTo(LiteLLMAppTable::class, 'app_id');
    }

    public function customer_org(): BelongsTo
    {
        return $this->belongsTo(LiteLLMCustomerOrganizationTable::class, 'customer_org_id');
    }

    public function tier_budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'tier_budget_id');
    }

    public function budget_override(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'budget_override_id');
    }

    public function customer_member_budget(): BelongsTo
    {
        return $this->belongsTo(LiteLLMBudgetTable::class, 'customer_member_budget_id');
    }
}