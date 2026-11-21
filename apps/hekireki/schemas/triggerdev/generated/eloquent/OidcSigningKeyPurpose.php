<?php

namespace App\Models;

/** Each purpose is a separate key set with its own issuer, rotation and publication. */
enum OidcSigningKeyPurpose: string
{
    case EXTERNAL_OIDC = 'EXTERNAL_OIDC';
    case INTERNAL_RUN = 'INTERNAL_RUN';
}