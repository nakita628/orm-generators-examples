<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class ProgramSuggestedSet extends Model
{
    protected $table = 'program_suggested_sets';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'programSessionExerciseId',
        'setIndex',
        'types',
        'valuesInt',
        'valuesSec',
        'units',
    ];

    protected $casts = [
        'setIndex' => 'integer',
        'types' => AsPrismaList::class . ':' . WorkoutSetType::class,
        'valuesInt' => AsPrismaList::class . ':int',
        'valuesSec' => AsPrismaList::class . ':int',
        'units' => AsPrismaList::class . ':' . WorkoutSetUnit::class,
    ];

    public function programSessionExercise(): BelongsTo
    {
        return $this->belongsTo(ProgramSessionExercise::class, 'programSessionExerciseId');
    }
}