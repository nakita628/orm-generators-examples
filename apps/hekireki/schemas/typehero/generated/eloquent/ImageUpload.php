<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $id the id of the db row
 * @property string $userId the user id fk
 * @property string $url the url to the s3 uploaded assets
 */
class ImageUpload extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'ImageUpload';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'userId',
        'url',
    ];

    /** The user that this is related to */
    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}