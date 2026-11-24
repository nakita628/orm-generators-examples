<?php

namespace App\Models;

enum EnvironmentVariableIgnoreReason: string
{
    case EMPTY = 'EMPTY';
    case MISSING = 'MISSING';
    case OIDC = 'OIDC';
}