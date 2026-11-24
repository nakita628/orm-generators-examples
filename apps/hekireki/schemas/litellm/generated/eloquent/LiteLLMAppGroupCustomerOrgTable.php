<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class LiteLLMAppGroupCustomerOrgTable extends Model
{
    const KEY_COLUMNS = ['app_group_id', 'customer_org_id'];

    protected $table = 'LiteLLM_AppGroupCustomerOrgTable';

    protected $primaryKey = null;

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'app_group_id',
        'customer_org_id',
        'budget_id',
        'spend',
        'spend_reset_window',
    ];

    protected $attributes = [
        'spend' => 0.0,
    ];

    protected $casts = [
        'spend' => 'float',
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

    public function app_group(): BelongsTo
    {
        return $this->belongsTo(LiteLLMAppGroupTable::class, 'app_group_id');
    }

    public function customer_org(): BelongsTo
    {
        return $this->belongsTo(LiteLLMCustomerOrganizationTable::class, 'customer_org_id');
    }
}