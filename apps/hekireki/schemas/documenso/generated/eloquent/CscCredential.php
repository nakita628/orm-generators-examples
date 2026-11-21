<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * Per-recipient cached credential metadata for the configured Cloud Signature
 * Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
 * symmetric-encrypted service-scope access token. One row per recipient on a
 * CSC envelope; absent for SES recipients.
 *
 * @property string|null $certCache TSP-validated certificate chain (length-prefixed). Nullable until the
 * service-scope OAuth callback successfully retrieves and validates it.
 * @property string $signatureAlgorithm Algorithm metadata derived from the CSC `credentials/info` response.
 * Held as flat scalars so each is independently queryable.
 * @property string|null $serviceTokenCiphertext Service-scope access token, symmetric-encrypted at the application layer
 * before persistence. Nullable until the service-scope OAuth callback runs.
 */
class CscCredential extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'CscCredential';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'providerId',
        'credentialId',
        'certCache',
        'signatureAlgorithm',
        'keyType',
        'digestAlgorithm',
        'keyLenBits',
        'signAlgoParams',
        'serviceTokenCiphertext',
        'serviceTokenExpiresAt',
        'recipientId',
    ];

    protected $casts = [
        'certCache' => AsBytes::class,
        'keyLenBits' => 'integer',
        'serviceTokenCiphertext' => AsBytes::class,
        'serviceTokenExpiresAt' => 'datetime',
        'recipientId' => 'integer',
    ];

    public function recipient(): BelongsTo
    {
        return $this->belongsTo(Recipient::class, 'recipientId');
    }
}