<?php

namespace App\Models;

enum OAuthClientKind: string
{
    case PREREGISTERED = 'PREREGISTERED';
    case CIMD = 'CIMD';
    case DCR = 'DCR';
}