<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Relations\BelongsTo;

class SessionView extends \Illuminate\Database\Eloquent\Model
{
    use PrismaDates;

    protected $table = 'session_views';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'id',
        'ownership',
        'is_pinned_by_default_for_project',
        'project_id',
        'name',
        'filters',
        'view_mode',
        'unmatched_traces',
        'created_by',
    ];

    protected $attributes = [
        'ownership' => 'PROJECT',
        'is_pinned_by_default_for_project' => false,
    ];

    protected $casts = [
        'ownership' => SessionViewOwnership::class,
        'is_pinned_by_default_for_project' => 'boolean',
        'filters' => 'array',
        'view_mode' => SessionViewMode::class,
        'unmatched_traces' => SessionUnmatchedTraces::class,
    ];

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'project_id');
    }

    public function createdByUser(): BelongsTo
    {
        return $this->belongsTo(User::class, 'created_by');
    }
}