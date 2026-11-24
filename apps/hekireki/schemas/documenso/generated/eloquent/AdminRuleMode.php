<?php

namespace App\Models;

enum AdminRuleMode: string
{
    case ENFORCE = 'ENFORCE';
    case MONITOR = 'MONITOR';
    case DISABLED = 'DISABLED';
}