<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

/** @property mixed $data @zod.import(["import { bookingSeatDataSchema } from '../../zod-utils'"]).custom.use(bookingSeatDataSchema) */
class BookingSeat extends Model
{
    protected $table = 'BookingSeat';

    public $timestamps = false;

    protected $fillable = [
        'referenceUid',
        'bookingId',
        'attendeeId',
        'data',
        'metadata',
    ];

    protected $casts = [
        'bookingId' => 'integer',
        'attendeeId' => 'integer',
        'data' => 'array',
        'metadata' => 'array',
    ];

    public function booking(): BelongsTo
    {
        return $this->belongsTo(Booking::class, 'bookingId');
    }

    public function attendee(): BelongsTo
    {
        return $this->belongsTo(Attendee::class, 'attendeeId');
    }
}