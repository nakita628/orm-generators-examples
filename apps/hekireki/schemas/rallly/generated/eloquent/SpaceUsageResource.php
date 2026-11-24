<?php

namespace App\Models;

enum SpaceUsageResource: string
{
    case api_poll = 'api_poll';
    case email_invite = 'email_invite';
}