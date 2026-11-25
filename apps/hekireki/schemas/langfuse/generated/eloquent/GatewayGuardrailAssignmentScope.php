<?php

namespace App\Models;

enum GatewayGuardrailAssignmentScope: string
{
    case ORGANIZATION = 'ORGANIZATION';
    case API_KEY = 'API_KEY';
    case USER = 'USER';
}