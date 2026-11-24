<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * A reason the environment variables "Needs attention" page shows a key in one environment, which
 * the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
 */
class EnvironmentVariableIgnore extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'EnvironmentVariableIgnore';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'projectId',
        'environmentId',
        'variableId',
        'reason',
        'createdById',
    ];

    protected $casts = [
        'reason' => EnvironmentVariableIgnoreReason::class,
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function environment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'environmentId');
    }

    /** For MISSING, a variable of the project with no value in this environment. */
    public function variable(): BelongsTo
    {
        return $this->belongsTo(EnvironmentVariable::class, 'variableId');
    }

    public function createdBy(): BelongsTo
    {
        return $this->belongsTo(User::class, 'createdById');
    }
}