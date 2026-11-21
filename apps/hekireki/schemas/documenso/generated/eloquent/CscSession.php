<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Per-recipient transient sign-time state held across the credential-scope
 * OAuth round-trip. The `recipientId @unique` constraint caps each recipient
 * to at most one in-flight session — re-clicking Sign UPSERTs the row.
 * Cleaned up transitively when the Recipient is deleted.
 * @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
 *
 * @property string $envelopeId Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
 * flows through the Recipient cascade.
 * @property \Carbon\CarbonInterface $signingTime Pinned at prep time; stable across overlay re-renders so the captured
 * signedAttrs digest matches the bytes the TSP signs.
 * @property mixed $itemsJson Contract between prep and sign: an array of
 * `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
 * `documentDataId` pins the immutable bytes that produced the captured
 * digest so sign-time can re-derive the same hash against the same content.
 * [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
 * @property string|null $encryptedSad Populated by the credential-scope OAuth callback. Both nullable at row
 * creation; set together when the SAD is exchanged.
 */
class CscSession extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'CscSession';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'envelopeId',
        'signingTime',
        'itemsJson',
        'encryptedSad',
        'sadExpiresAt',
        'recipientId',
    ];

    protected $casts = [
        'signingTime' => 'datetime',
        'itemsJson' => 'array',
        'encryptedSad' => AsBytes::class,
        'sadExpiresAt' => 'datetime',
        'recipientId' => 'integer',
    ];

    public function recipient(): BelongsTo
    {
        return $this->belongsTo(Recipient::class, 'recipientId');
    }
}