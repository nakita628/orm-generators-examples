<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
 * Multiple organizations can share a single data store row via organizationIds.
 *
 * @property string $key Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
 * @property mixed $organizationIds Organization IDs that use this data store
 * @property mixed $config Versioned config JSON. Structure is discriminated by the top-level `version` field.
 */
class OrganizationDataStore extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'OrganizationDataStore';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'key',
        'organizationIds',
        'kind',
        'config',
    ];

    protected $casts = [
        'kind' => DataStoreKind::class,
        'config' => 'array',
    ];
}