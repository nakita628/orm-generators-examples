<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/**
 * @property string $originalUrl the original full url that we want to shorten e.g. https://www.url.com
 * @property string $shortUrlSlug slug for the short url e.g. url.com/[shortUrlSlug]
 * @property \Carbon\CarbonInterface|null $expiresAt the time the short url will expire
 * @property int $clicks the number of times the short url has redirected to the original url
 */
class ShortURL extends Model
{
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = null;

    protected $table = 'ShortURL';

    protected $fillable = [
        'originalUrl',
        'shortUrlSlug',
        'expiresAt',
        'clicks',
        'userId',
    ];

    protected $attributes = [
        'clicks' => 0,
    ];

    protected $casts = [
        'expiresAt' => 'datetime',
        'clicks' => 'integer',
    ];

    public function user(): BelongsTo
    {
        return $this->belongsTo(User::class, 'userId');
    }
}