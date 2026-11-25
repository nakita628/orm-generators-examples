<?php

namespace App\Models;

enum GatewayGuardrailType: string
{
    case REGEX = 'REGEX';
    case BEDROCK = 'BEDROCK';
}