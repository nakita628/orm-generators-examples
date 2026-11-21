<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $chatId The chat session ID used by the transport
 * @property string $title User-editable conversation title (auto-generated from first message)
 * @property string $agentSlug Which agent this conversation is with
 * @property string|null $runId The current active run backing this conversation (null if no run yet)
 * @property mixed $clientData The client data JSON used for this conversation
 * @property mixed $messages Accumulated UIMessages from completed turns (for resume without stream replay)
 * @property string|null $lastEventId Last SSE event ID — resume from this position to avoid replaying old turns
 * @property string $userId The user who started this conversation
 */
class PlaygroundConversation extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'PlaygroundConversation';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'chatId',
        'title',
        'agentSlug',
        'runId',
        'clientData',
        'messages',
        'lastEventId',
        'projectId',
        'runtimeEnvironmentId',
        'userId',
    ];

    protected $attributes = [
        'title' => 'New conversation',
    ];

    protected $casts = [
        'clientData' => 'array',
        'messages' => 'array',
    ];

    public function run(): BelongsTo
    {
        return $this->belongsTo(TaskRun::class, 'runId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function runtimeEnvironment(): BelongsTo
    {
        return $this->belongsTo(RuntimeEnvironment::class, 'runtimeEnvironmentId');
    }
}