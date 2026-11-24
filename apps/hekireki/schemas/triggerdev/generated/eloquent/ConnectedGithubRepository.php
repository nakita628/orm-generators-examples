<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/** @property bool $supersedeProductionBuilds Skip queued builds once a newer commit is queued for the same environment or preview branch */
class ConnectedGithubRepository extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'ConnectedGithubRepository';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'projectId',
        'repositoryId',
        'branchTracking',
        'previewDeploymentsEnabled',
        'supersedeProductionBuilds',
        'supersedeStagingBuilds',
        'supersedePreviewBuilds',
        'buildFilters',
    ];

    protected $attributes = [
        'previewDeploymentsEnabled' => true,
        'supersedeProductionBuilds' => false,
        'supersedeStagingBuilds' => true,
        'supersedePreviewBuilds' => true,
    ];

    protected $casts = [
        'branchTracking' => 'array',
        'previewDeploymentsEnabled' => 'boolean',
        'supersedeProductionBuilds' => 'boolean',
        'supersedeStagingBuilds' => 'boolean',
        'supersedePreviewBuilds' => 'boolean',
        'buildFilters' => 'array',
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function repository(): BelongsTo
    {
        return $this->belongsTo(GithubRepository::class, 'repositoryId');
    }
}