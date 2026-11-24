<?php

namespace App\Models;

enum DeploymentInitiatedByCredential: string
{
    case ROOT_API_KEY = 'ROOT_API_KEY';
    case SCOPED_API_KEY = 'SCOPED_API_KEY';
    case PERSONAL_ACCESS_TOKEN = 'PERSONAL_ACCESS_TOKEN';
    case ORGANIZATION_ACCESS_TOKEN = 'ORGANIZATION_ACCESS_TOKEN';
    case TRUSTED_PUBLISHING = 'TRUSTED_PUBLISHING';
}