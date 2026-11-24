<?php

namespace App\Models;

enum GatewaySpendLimitAlertType: string
{
    case WARNING = 'WARNING';
    case EXCEEDED = 'EXCEEDED';
}