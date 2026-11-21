<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Track individual run creation failures in batch processing (Run Engine v2)
 *
 * @property int $index Which item in the batch (0-based index)
 * @property string $taskIdentifier The task identifier that was being triggered
 * @property string|null $payload The payload that failed (JSON, may be truncated)
 * @property mixed $options The options that were used
 * @property string $error Error message
 * @property string|null $errorCode Error code if available
 */
class BatchTaskRunError extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'BatchTaskRunError';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'batchTaskRunId',
        'index',
        'taskIdentifier',
        'payload',
        'options',
        'error',
        'errorCode',
    ];

    protected $casts = [
        'index' => 'integer',
        'options' => 'array',
    ];

    public function batchTaskRun(): BelongsTo
    {
        return $this->belongsTo(BatchTaskRun::class, 'batchTaskRunId');
    }
}