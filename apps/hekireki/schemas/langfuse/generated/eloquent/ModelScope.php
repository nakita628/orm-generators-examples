<?php

namespace App\Models;

enum ModelScope: string
{
    case LANGFUSE = 'LANGFUSE';
    case ORGANIZATION = 'ORGANIZATION';
    case PROJECT = 'PROJECT';
}