<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * A stock split of an asset profile. The numerator is the number of shares
 * held after the split, the denominator the number of shares held before, for
 * example 4 and 1 for a 4:1 split or 1 and 10 for a 1:10 reverse split.
 *
 * The ratio is stored as two integers instead of the resulting factor, so that
 * it stays exact, for example 1/3 for a 1:3 reverse split.
 *
 * Activities before the date are adjusted by the ratio: the quantity is
 * multiplied by it and the unit price divided by it, so that the total value
 * of an activity stays the same.
 *
 * Market data is already split-adjusted by the data providers and must not be
 * adjusted again.
 */
class AssetProfileSplit extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'AssetProfileSplit';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'date',
        'denominator',
        'numerator',
        'symbolProfileId',
    ];

    protected $casts = [
        'date' => 'datetime',
        'denominator' => 'integer',
        'numerator' => 'integer',
    ];

    public function symbolProfile(): BelongsTo
    {
        return $this->belongsTo(SymbolProfile::class, 'symbolProfileId');
    }
}