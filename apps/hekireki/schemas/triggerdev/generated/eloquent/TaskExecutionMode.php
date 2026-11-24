<?php

namespace App\Models;

enum TaskExecutionMode: string
{
    case STANDARD = 'STANDARD';
    case SANDBOX = 'SANDBOX';
}