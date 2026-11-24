<?php

namespace App\Models;

enum GatewaySpendLimitInterval: string
{
    case DAILY = 'DAILY';
    case WEEKLY = 'WEEKLY';
    case MONTHLY = 'MONTHLY';
    case TOTAL = 'TOTAL';
}