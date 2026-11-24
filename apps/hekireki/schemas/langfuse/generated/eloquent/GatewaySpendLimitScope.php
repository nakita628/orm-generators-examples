<?php

namespace App\Models;

enum GatewaySpendLimitScope: string
{
    case API_KEY = 'API_KEY';
    case USER = 'USER';
    case API_KEY_DEFAULT = 'API_KEY_DEFAULT';
    case USER_DEFAULT = 'USER_DEFAULT';
    case ORGANIZATION = 'ORGANIZATION';
}