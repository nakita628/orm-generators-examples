<?php

namespace App\Models;

enum GatewayGuardrailAction: string
{
    case BLOCK = 'BLOCK';
    case LOG = 'LOG';
    case REDACT = 'REDACT';
}