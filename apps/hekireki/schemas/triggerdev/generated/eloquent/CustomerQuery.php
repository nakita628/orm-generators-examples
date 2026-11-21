<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $query The TSQL query text that was executed
 * @property CustomerQueryScope $scope The scope of the query (determines which tenant IDs were used)
 * @property mixed $stats Query execution statistics from ClickHouse
 * @property string|null $title AI-generated title summarizing the query
 * @property CustomerQuerySource $source Where the query originated from
 * @property string|null $filterPeriod Time filter settings used when the query was executed
 * Period like "7d", "24h", etc.
 * @property \Carbon\CarbonInterface|null $filterFrom Custom start date (ISO 8601)
 * @property \Carbon\CarbonInterface|null $filterTo Custom end date (ISO 8601)
 */
class CustomerQuery extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'CustomerQuery';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'query',
        'scope',
        'stats',
        'title',
        'source',
        'organizationId',
        'projectId',
        'environmentId',
        'userId',
        'filterPeriod',
        'filterFrom',
        'filterTo',
    ];

    protected $attributes = [
        'source' => 'DASHBOARD',
    ];

    protected $casts = [
        'scope' => CustomerQueryScope::class,
        'stats' => 'array',
        'source' => CustomerQuerySource::class,
        'filterFrom' => 'datetime',
        'filterTo' => 'datetime',
    ];

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function environment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'environmentId');
    }

    /** Optional user who executed the query (null for API calls) */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}