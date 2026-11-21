<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * `key` is sha256(organisationId:lower(recipientEmail):clientUserId) so neither
 * the recipient email nor the embedder's identifier is stored in plaintext.
 *
 * @property int $recipientId Last writer. Not a relation so the row survives document deletion and
 * remains useful for abuse tracing.
 */
class EmbedSignature extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'EmbedSignature';

    protected $primaryKey = 'key';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'key',
        'organisationId',
        'recipientId',
        'signature',
    ];

    protected $casts = [
        'recipientId' => 'integer',
    ];

    public function organisation(): BelongsTo
    {
        return $this->belongsTo(Organisation::class, 'organisationId');
    }
}