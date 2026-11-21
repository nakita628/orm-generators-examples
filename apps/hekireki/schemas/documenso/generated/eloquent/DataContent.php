<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * Generic stored data (e.g. an uploaded image) which other models reference.
 * The kind of data and its kind specific details live in `metadata`.
 * @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
 *
 * @property DocumentDataType $type How `data` is stored, e.g. an S3 path or inline bytes.
 * @property mixed $metadata [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema)
 */
class DataContent extends Model
{
    protected $table = 'DataContent';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'type',
        'data',
        'metadata',
    ];

    protected $casts = [
        'type' => DocumentDataType::class,
        'metadata' => 'array',
    ];

    public function envelopeContents(): HasMany
    {
        return $this->hasMany(EnvelopeContent::class, 'dataContentId');
    }
}