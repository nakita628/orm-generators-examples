<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @zod.import(["import { ZEnvelopeAttachmentTypeSchema } from '@documenso/lib/types/envelope-attachment';"])
 *
 * @property string $type [EnvelopeAttachmentType] @zod.custom.use(ZEnvelopeAttachmentTypeSchema)
 */
class EnvelopeAttachment extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'EnvelopeAttachment';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'type',
        'label',
        'data',
        'envelopeId',
    ];

    public function envelope(): BelongsTo
    {
        return $this->belongsTo(Envelope::class, 'envelopeId');
    }
}