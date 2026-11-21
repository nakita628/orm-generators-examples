<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/** @property \Carbon\CarbonInterface|null $expires_at @zod.custom(z.coerce.date()) */
class KvStoreEntry extends Model
{
    use PrismaDates;

    protected $table = 'kv';

    protected $primaryKey = 'key';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'key',
        'value',
        'expires_at',
    ];

    protected $casts = [
        'value' => 'array',
        'expires_at' => 'datetime',
    ];
}