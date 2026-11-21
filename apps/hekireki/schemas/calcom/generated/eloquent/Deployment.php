<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

/**
 * @property int $id This is a single row table, so we use a fixed id
 * @property mixed $theme @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
 */
class Deployment extends Model
{
    use PrismaDates;

    protected $table = 'Deployment';

    public $incrementing = false;

    public $timestamps = false;

    protected $fillable = [
        'id',
        'logo',
        'theme',
        'licenseKey',
        'signatureTokenEncrypted',
        'agreedLicenseAt',
    ];

    protected $casts = [
        'theme' => 'array',
        'agreedLicenseAt' => 'datetime',
    ];
}