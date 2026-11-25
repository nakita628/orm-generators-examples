<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

class Conversation extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'conversations';

    protected $keyType = 'string';

    protected $fillable = [
        'id',
        'projectId',
        'organizationId',
        'userId',
        'title',
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

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function organization(): BelongsTo
    {
        return $this->belongsTo(Organization::class, 'organizationId');
    }

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }

    public function messages(): HasMany
    {
        return $this->hasMany(ChatMessage::class, 'conversationId');
    }
}