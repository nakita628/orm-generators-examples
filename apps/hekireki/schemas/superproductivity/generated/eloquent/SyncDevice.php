<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string|null $app_version Bare semver the client last reported on download (`18.22.0`), NULL until
 * it does. Read by the checkpoint gate (#9962); never shown to users.
 * @property int $last_acked_seq Reserved: written once at device registration, never read or updated.
 */
class SyncDevice extends Model
{
    const KEY_COLUMNS = ['user_id', 'client_id'];

    protected $table = 'sync_devices';

    protected $primaryKey = null;

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'client_id',
        'user_id',
        'device_name',
        'user_agent',
        'app_version',
        'last_seen_at',
        'last_acked_seq',
        'created_at',
    ];

    protected $attributes = [
        'last_acked_seq' => 0,
    ];

    protected $casts = [
        'user_id' => 'integer',
        'last_seen_at' => 'integer',
        'last_acked_seq' => 'integer',
        'created_at' => 'integer',
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

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'user_id');
    }
}