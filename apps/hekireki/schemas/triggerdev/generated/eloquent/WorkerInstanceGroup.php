<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property string $name For example "us-east-1"
 * @property string $masterQueue If managed, it will default to the name, e.g. "us-east-1"
 * If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
 * @property string|null $description "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes
 * @property string|null $location "usa", "europe", etc. Used like a pseudo enum for things like flags
 * @property string|null $region Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs.
 * @property bool $enableFastPath When true, runs enqueued to this worker queue may skip the intermediate queue
 * and be pushed directly to the worker queue when concurrency is available.
 */
class WorkerInstanceGroup extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'WorkerInstanceGroup';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'type',
        'name',
        'masterQueue',
        'description',
        'hidden',
        'tokenId',
        'organizationId',
        'projectId',
        'cloudProvider',
        'location',
        'staticIPs',
        'workloadType',
        'region',
        'enableFastPath',
    ];

    protected $attributes = [
        'hidden' => false,
        'workloadType' => 'CONTAINER',
        'enableFastPath' => false,
    ];

    protected $casts = [
        'type' => WorkerInstanceGroupType::class,
        'hidden' => 'boolean',
        'workloadType' => WorkloadType::class,
        'enableFastPath' => 'boolean',
    ];

    public function token(): BelongsTo
    {
        return $this->belongsTo(WorkerGroupToken::class, 'tokenId');
    }

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function childTokens(): HasMany
    {
        return $this->hasMany(WorkerGroupToken::class, 'workerGroupId');
    }

    public function workers(): HasMany
    {
        return $this->hasMany(WorkerInstance::class, 'workerGroupId');
    }

    public function backgroundWorkers(): HasMany
    {
        return $this->hasMany(BackgroundWorker::class, 'workerGroupId');
    }

    public function defaultForProjects(): HasMany
    {
        return $this->hasMany(Project::class, 'defaultWorkerGroupId');
    }
}