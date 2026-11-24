from datetime import datetime
from datetime import timezone as dt_timezone
from typing import Any

import cuid  # type: ignore[import-untyped, unused-ignore]
from django.conf import settings
from django.db import models
from django.db.backends.base.base import BaseDatabaseWrapper
from django.utils import timezone


class UtcDateTimeField(models.DateTimeField):  # type: ignore[type-arg]
    def get_db_prep_value(self, value: Any, connection: BaseDatabaseWrapper, prepared: bool = False) -> Any:
        if not prepared:
            value = self.get_prep_value(value)
        if not isinstance(value, datetime):
            return super().get_db_prep_value(value, connection, prepared=True)
        if timezone.is_naive(value):
            value = timezone.make_aware(value, timezone.get_default_timezone())
        value = value.astimezone(dt_timezone.utc)
        if connection.vendor == "sqlite":
            return value.isoformat(timespec="milliseconds")
        return connection.ops.adapt_datetimefield_value(value.replace(tzinfo=None))

    def from_db_value(self, value: datetime | None, expression: Any, connection: BaseDatabaseWrapper) -> datetime | None:
        if value is None:
            return None
        if timezone.is_naive(value):
            value = value.replace(tzinfo=dt_timezone.utc)
        return value if settings.USE_TZ else timezone.make_naive(value, timezone.get_default_timezone())


def cuid_str() -> str:
    return str(cuid.cuid())


def now_milliseconds() -> datetime:
    now = timezone.now()
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class Account(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="accounts", db_column="userId", db_index=False)
    type = models.CharField(max_length=191)
    provider = models.CharField(max_length=191)
    provider_account_id = models.CharField(max_length=191, db_column="providerAccountId")
    refresh_token = models.TextField(null=True)
    access_token = models.TextField(null=True)
    expires_at = models.IntegerField(null=True)
    token_type = models.CharField(max_length=191, null=True)
    scope = models.CharField(max_length=191, null=True)
    id_token = models.TextField(null=True)
    session_state = models.CharField(max_length=191, null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcDateTimeField(default=now_milliseconds)

    class Meta:
        db_table = "accounts"
        constraints = [
            models.UniqueConstraint(fields=["provider", "provider_account_id"], name="accounts_provider_providerAccountId_key"),
        ]


class Session(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    session_token = models.CharField(max_length=191, unique=True, db_column="sessionToken")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sessions", db_column="userId", db_index=False)
    expires = UtcDateTimeField()

    class Meta:
        db_table = "sessions"


class User(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    name = models.CharField(max_length=191, null=True)
    email = models.CharField(max_length=191, null=True, unique=True)
    email_verified = UtcDateTimeField(null=True, db_column="emailVerified")
    image = models.CharField(max_length=191, null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcDateTimeField(default=now_milliseconds)
    stripe_customer_id = models.CharField(max_length=191, null=True, unique=True)
    stripe_subscription_id = models.CharField(max_length=191, null=True, unique=True)
    stripe_price_id = models.CharField(max_length=191, null=True)
    stripe_current_period_end = UtcDateTimeField(null=True)

    class Meta:
        db_table = "users"


class VerificationToken(models.Model):
    identifier = models.CharField(max_length=191)
    token = models.CharField(max_length=191, primary_key=True, default=None)
    expires = UtcDateTimeField()

    class Meta:
        db_table = "verification_tokens"
        constraints = [
            models.UniqueConstraint(fields=["identifier", "token"], name="verification_tokens_identifier_token_key"),
        ]


class Post(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    title = models.CharField(max_length=191)
    content = models.JSONField(null=True)
    published = models.BooleanField(default=False)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcDateTimeField(default=now_milliseconds)
    author = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="post", db_column="authorId", db_index=False)

    class Meta:
        db_table = "posts"
