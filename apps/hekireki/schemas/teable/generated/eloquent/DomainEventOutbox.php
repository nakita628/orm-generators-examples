<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class DomainEventOutbox extends Model
{
    use PrismaDates;

    const UPDATED_AT = null;

    protected $table = 'domain_event_outbox';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'base_id',
        'table_id',
        'message_name',
        'schema_version',
        'aggregate_id',
        'payload',
        'payload_bytes',
        'catalog_generation',
        'required_consumers',
        'binding_id',
        'storage_epoch',
        'unpublished',
        'settled',
        'settled_at',
    ];

    protected $attributes = [
        'unpublished' => true,
    ];

    protected $casts = [
        'schema_version' => 'integer',
        'payload' => 'array',
        'payload_bytes' => 'integer',
        'catalog_generation' => 'integer',
        'required_consumers' => 'array',
        'storage_epoch' => 'integer',
        'unpublished' => 'boolean',
        'settled_at' => 'datetime',
    ];
}