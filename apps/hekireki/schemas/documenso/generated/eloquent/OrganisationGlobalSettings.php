<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;
use Illuminate\Database\Eloquent\Relations\HasOne;

/**
 * @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"])
 *
 * @property mixed $defaultRecipients [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema)
 * @property mixed $emailDocumentSettings [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema)
 * @property mixed $brandingColors [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema)
 * @property mixed $envelopeExpirationPeriod [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema)
 * @property mixed $reminderSettings [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema)
 */
class OrganisationGlobalSettings extends Model
{
    protected $table = 'OrganisationGlobalSettings';

    protected $keyType = 'string';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'documentVisibility',
        'documentLanguage',
        'includeSenderDetails',
        'includeSigningCertificate',
        'includeAuditLog',
        'documentTimezone',
        'documentDateFormat',
        'delegateDocumentOwnership',
        'typedSignatureEnabled',
        'uploadSignatureEnabled',
        'drawSignatureEnabled',
        'defaultRecipients',
        'emailId',
        'emailReplyTo',
        'emailDocumentSettings',
        'brandingEnabled',
        'brandingLogo',
        'brandingUrl',
        'brandingCompanyDetails',
        'brandingColors',
        'brandingCss',
        'envelopeExpirationPeriod',
        'reminderSettings',
        'aiFeaturesEnabled',
    ];

    protected $attributes = [
        'documentVisibility' => 'EVERYONE',
        'documentLanguage' => 'en',
        'includeSenderDetails' => true,
        'includeSigningCertificate' => true,
        'includeAuditLog' => false,
        'documentDateFormat' => 'yyyy-MM-dd hh:mm a',
        'delegateDocumentOwnership' => false,
        'typedSignatureEnabled' => true,
        'uploadSignatureEnabled' => true,
        'drawSignatureEnabled' => true,
        'brandingEnabled' => false,
        'brandingLogo' => '',
        'brandingUrl' => '',
        'brandingCompanyDetails' => '',
        'brandingCss' => '',
        'aiFeaturesEnabled' => false,
    ];

    protected $casts = [
        'documentVisibility' => DocumentVisibility::class,
        'includeSenderDetails' => 'boolean',
        'includeSigningCertificate' => 'boolean',
        'includeAuditLog' => 'boolean',
        'delegateDocumentOwnership' => 'boolean',
        'typedSignatureEnabled' => 'boolean',
        'uploadSignatureEnabled' => 'boolean',
        'drawSignatureEnabled' => 'boolean',
        'defaultRecipients' => 'array',
        'emailDocumentSettings' => 'array',
        'brandingEnabled' => 'boolean',
        'brandingColors' => 'array',
        'envelopeExpirationPeriod' => 'array',
        'reminderSettings' => 'array',
        'aiFeaturesEnabled' => 'boolean',
    ];

    public function email(): BelongsTo
    {
        return $this->belongsTo(OrganisationEmail::class, 'emailId');
    }

    public function organisation(): HasOne
    {
        return $this->hasOne(Organisation::class, 'organisationGlobalSettingsId');
    }
}