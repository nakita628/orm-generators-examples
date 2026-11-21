<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Concerns\HasVersion4Uuids;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasMany;

/**
 * @property \Carbon\CarbonInterface|null $scheduledFor When a transactional email sent with `delivery.at` is due. Only meaningful while the
 * status is SCHEDULED; the delayed queue job is the timer and the scheduled-email sweep
 * re-queues any row still SCHEDULED past this time. The partial index on SCHEDULED rows
 * is built by the background index builder, not a migration (see background-index-builder.ts).
 * @property \Carbon\CarbonInterface|null $unsubscribedAt When the recipient opted out from this email (unsubscribe or snooze). Only set for
 * opt-outs attributed to a campaign email and counted in `Campaign.unsubscribedCount`,
 * in the same transaction, so the stamped rows are the list behind that figure. Indexed
 * by a partial index built in the background (see background-index-builder.ts).
 * @property bool $simulated The address handed to SES belonged to the mailbox simulator
 * (`@simulator.amazonses.com`), stamped at send time from the address actually used --
 * which is the recipient override when one is set, not necessarily `contact.email`.
 *
 * AWS itself excludes simulator mail from bounce rates, complaint rates, sending quotas
 * and reputation metrics, so Plunk's security rates exclude it from both the numerator
 * and the denominator. The events themselves are untouched: the webhook fires, the
 * activity feed shows the bounce, the contact is still suppressed.
 */
class Email extends Model
{
    use HasVersion4Uuids;
    use PrismaDates;

    const CREATED_AT = 'createdAt';
    const UPDATED_AT = 'updatedAt';

    protected $table = 'emails';

    protected $keyType = 'string';

    public $incrementing = false;

    protected $fillable = [
        'contactId',
        'toName',
        'subject',
        'body',
        'from',
        'fromName',
        'replyTo',
        'headers',
        'attachments',
        'messageId',
        'sourceType',
        'templateId',
        'campaignId',
        'workflowExecutionId',
        'workflowStepExecutionId',
        'status',
        'scheduledFor',
        'sentAt',
        'deliveredAt',
        'openedAt',
        'clickedAt',
        'bouncedAt',
        'complainedAt',
        'unsubscribedAt',
        'opens',
        'clicks',
        'simulated',
        'error',
        'projectId',
    ];

    protected $attributes = [
        'status' => 'PENDING',
        'opens' => 0,
        'clicks' => 0,
        'simulated' => false,
    ];

    protected $casts = [
        'headers' => 'array',
        'attachments' => 'array',
        'sourceType' => EmailSourceType::class,
        'status' => EmailStatus::class,
        'scheduledFor' => 'datetime',
        'sentAt' => 'datetime',
        'deliveredAt' => 'datetime',
        'openedAt' => 'datetime',
        'clickedAt' => 'datetime',
        'bouncedAt' => 'datetime',
        'complainedAt' => 'datetime',
        'unsubscribedAt' => 'datetime',
        'opens' => 'integer',
        'clicks' => 'integer',
        'simulated' => 'boolean',
    ];

    public function contact(): BelongsTo
    {
        return $this->belongsTo(Contact::class, 'contactId');
    }

    public function template(): BelongsTo
    {
        return $this->belongsTo(Template::class, 'templateId');
    }

    public function campaign(): BelongsTo
    {
        return $this->belongsTo(Campaign::class, 'campaignId');
    }

    public function workflowExecution(): BelongsTo
    {
        return $this->belongsTo(WorkflowExecution::class, 'workflowExecutionId');
    }

    public function workflowStepExecution(): BelongsTo
    {
        return $this->belongsTo(WorkflowStepExecution::class, 'workflowStepExecutionId');
    }

    public function project(): BelongsTo
    {
        return $this->belongsTo(Project::class, 'projectId');
    }

    public function events(): HasMany
    {
        return $this->hasMany(Event::class, 'emailId');
    }
}