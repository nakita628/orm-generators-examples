"""A multi-tenant SaaS's accounts: users, organizations, roles, invitations and API keys."""

import uuid

from django.contrib.auth.models import AbstractBaseUser, PermissionsMixin, UserManager
from django.db import models
from django.db.models.functions import Lower


class User(AbstractBaseUser, PermissionsMixin):
    """Signs in with an e-mail address, not a username."""

    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    email = models.EmailField(unique=True)
    full_name = models.CharField(max_length=150, blank=True)
    is_staff = models.BooleanField(default=False)
    is_active = models.BooleanField(default=True)
    date_joined = models.DateTimeField(auto_now_add=True)

    USERNAME_FIELD = "email"
    objects = UserManager()

    class Meta:
        constraints = [models.UniqueConstraint(Lower("email"), name="user_email_ci_unique")]


class Organization(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    name = models.CharField(max_length=100)
    slug = models.SlugField(unique=True)
    owner = models.ForeignKey(User, on_delete=models.PROTECT, related_name="owned_organizations")
    members = models.ManyToManyField(User, through="Membership", related_name="organizations")
    created_at = models.DateTimeField(auto_now_add=True)


class Membership(models.Model):
    class Role(models.TextChoices):
        OWNER = "owner", "Owner"
        ADMIN = "admin", "Admin"
        MEMBER = "member", "Member"
        VIEWER = "viewer", "Viewer"

    organization = models.ForeignKey(Organization, on_delete=models.CASCADE)
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    role = models.CharField(max_length=10, choices=Role, default=Role.MEMBER)
    joined_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        constraints = [models.UniqueConstraint(fields=["organization", "user"], name="membership_unique")]


class Team(models.Model):
    organization = models.ForeignKey(Organization, on_delete=models.CASCADE, related_name="teams")
    name = models.CharField(max_length=100)
    members = models.ManyToManyField(Membership, related_name="teams", blank=True)

    class Meta:
        unique_together = [("organization", "name")]


class Invitation(models.Model):
    organization = models.ForeignKey(Organization, on_delete=models.CASCADE, related_name="invitations")
    email = models.EmailField()
    role = models.CharField(max_length=10, choices=Membership.Role, default=Membership.Role.MEMBER)
    token = models.CharField(max_length=64, unique=True)
    invited_by = models.ForeignKey(User, on_delete=models.SET_NULL, null=True, related_name="+")
    expires_at = models.DateTimeField()
    accepted_at = models.DateTimeField(null=True, blank=True)


class ApiKey(models.Model):
    organization = models.ForeignKey(Organization, on_delete=models.CASCADE, related_name="api_keys")
    name = models.CharField(max_length=50)
    prefix = models.CharField(max_length=8, unique=True)
    hashed_key = models.CharField(max_length=128)
    scopes = models.JSONField(default=list)
    created_by = models.ForeignKey(User, on_delete=models.SET_NULL, null=True, related_name="+")
    last_used_at = models.DateTimeField(null=True, blank=True)
    revoked = models.BooleanField(default=False)


class AuditEvent(models.Model):
    organization = models.ForeignKey(Organization, on_delete=models.CASCADE, related_name="audit_events")
    # Kept when the user is deleted, so the log stays whole
    actor = models.ForeignKey(User, on_delete=models.DO_NOTHING, null=True, db_constraint=False, related_name="+")
    action = models.CharField(max_length=50)
    target = models.CharField(max_length=200, blank=True)
    ip_address = models.GenericIPAddressField(null=True, blank=True)
    data = models.JSONField(default=dict)
    at = models.DateTimeField(auto_now_add=True, db_index=True)
