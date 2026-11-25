<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property mixed $parts [IPrismaUIMessageParts]
 * @property mixed $metadata [IPrismaChatMessageMetadata]
 */
class ChatMessage extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'chat_messages';

    protected $keyType = 'string';

    protected $fillable = [
        'id',
        'conversationId',
        'role',
        'parts',
        'metadata',
    ];

    protected $casts = [
        'parts' => 'array',
        'metadata' => 'array',
    ];

    protected function insertAndSetId(Builder $query, $attributes)
    {
        $base = $query->getQuery();
        $base->applyBeforeQueryCallbacks();
        $connection = $base->getConnection();
        $connection->recordsHaveBeenModified();
        $keyName = $this->getKeyName();
        $row = $connection->selectFromWriteConnection(
            $base->getGrammar()->compileInsertGetId($base, $attributes, $keyName),
            $base->cleanBindings($attributes),
        )[0];
        $this->setAttribute($keyName, is_object($row) ? $row->{$keyName} : $row[$keyName]);
    }

    public function conversation(): BelongsTo
    {
        return $this->belongsTo(Conversation::class, 'conversationId');
    }
}