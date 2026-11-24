<?php

namespace App\Models;

enum SessionViewOwnership: string
{
    case PRIVATE = 'PRIVATE';
    case PROJECT = 'PROJECT';
}