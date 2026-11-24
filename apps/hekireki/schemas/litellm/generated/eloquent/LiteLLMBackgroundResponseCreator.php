<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class LiteLLMBackgroundResponseCreator extends Model
{
    use PrismaDates;

    const UPDATED_AT = null;

    protected $table = 'LiteLLM_BackgroundResponseCreator';

    protected $primaryKey = 'response_id';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'response_id',
        'upstream_url',
        'create_context',
        'poll_lease_expires_at',
    ];

    protected $casts = [
        'create_context' => 'array',
        'poll_lease_expires_at' => 'datetime',
    ];
}