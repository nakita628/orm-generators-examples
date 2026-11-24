import uuid
from datetime import datetime
from datetime import timezone as dt_timezone
from typing import Any, TypeVar

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


class UtcAutoNowDateTimeField(UtcDateTimeField):
    def stamp(self) -> datetime:
        return now_milliseconds()

    def pre_save(self, model_instance: models.Model, add: bool) -> Any:
        if not self.auto_now:
            return super().pre_save(model_instance, add)
        value = self.stamp()
        setattr(model_instance, self.attname, value)
        return value


_M = TypeVar("_M", bound=models.Model)


class AutoNowQuerySet(models.QuerySet[_M]):
    def update(self, **kwargs: Any) -> int:
        for field in self.model._meta.concrete_fields:
            if getattr(field, "auto_now", False):
                kwargs.setdefault(field.name, getattr(field, "stamp", timezone.now)())
        return super().update(**kwargs)


def uuid4_str() -> str:
    return str(uuid.uuid4())


def cuid_str() -> str:
    return str(cuid.cuid())


def now_milliseconds() -> datetime:
    now = timezone.now()
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class NotificationType(models.TextChoices):
    LIKE = "LIKE"
    REPLY = "REPLY"
    MENTION = "MENTION"
    MISC = "MISC"
    SYSTEM = "SYSTEM"


class Tags(models.TextChoices):
    POPULAR = "POPULAR"
    NEWEST = "NEWEST"


class Difficulty(models.TextChoices):
    BEGINNER = "BEGINNER"
    EASY = "EASY"
    MEDIUM = "MEDIUM"
    HARD = "HARD"
    EXTREME = "EXTREME"
    EVENT = "EVENT"


class RoleTypes(models.TextChoices):
    USER = "USER"
    ADMIN = "ADMIN"
    MODERATOR = "MODERATOR"
    CREATOR = "CREATOR"
    SUPPORTER = "SUPPORTER"
    CONTRIBUTOR = "CONTRIBUTOR"


class UserStatus(models.TextChoices):
    BANNED = "BANNED"
    ACTIVE = "ACTIVE"


class ReportStatus(models.TextChoices):
    PENDING = "PENDING"
    CLEARED = "CLEARED"
    DISMISSED = "DISMISSED"


class ChallengeStatus(models.TextChoices):
    DECLINED = "DECLINED"
    BANNED = "BANNED"
    PENDING = "PENDING"
    ACTIVE = "ACTIVE"


class VoteType(models.TextChoices):
    CHALLENGE = "CHALLENGE"
    COMMENT = "COMMENT"
    SUBMISSION = "SUBMISSION"
    SHAREDSOLUTION = "SHAREDSOLUTION"


class ReportType(models.TextChoices):
    CHALLENGE = "CHALLENGE"
    USER = "USER"
    COMMENT = "COMMENT"
    SOLUTION = "SOLUTION"


class IssueType(models.TextChoices):
    DEROGATORY = "DEROGATORY"
    OTHER = "OTHER"
    UNCLEAR = "UNCLEAR"
    BULLYING = "BULLYING"
    SPAM = "SPAM"
    HATESPEECH = "HATESPEECH"
    THREAT = "THREAT"


class CommentRoot(models.TextChoices):
    SOLUTION = "SOLUTION"
    CHALLENGE = "CHALLENGE"


class Account(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="accounts", db_column="userId", db_index=False)
    type = models.CharField(max_length=191)
    provider = models.CharField(max_length=191)
    provider_account_id = models.CharField(max_length=191, db_column="providerAccountId")
    refresh_token = models.CharField(max_length=191, null=True)
    access_token = models.CharField(max_length=191, null=True)
    expires_at = models.IntegerField(null=True)
    token_type = models.CharField(max_length=191, null=True)
    scope = models.CharField(max_length=191, null=True)
    id_token = models.CharField(max_length=191, null=True)
    session_state = models.CharField(max_length=191, null=True)

    class Meta:
        db_table = "Account"
        constraints = [
            models.UniqueConstraint(fields=["provider", "provider_account_id"], name="Account_provider_providerAccountId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class Session(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=cuid_str)
    session_token = models.CharField(max_length=191, unique=True, db_column="sessionToken")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sessions", db_column="userId", db_index=False)
    expires = UtcDateTimeField()

    class Meta:
        db_table = "Session"
        indexes = [
            models.Index(fields=["user"]),
        ]


class Role(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=uuid4_str)
    user: "models.ManyToManyField[User, RoleToUser]" = models.ManyToManyField("User", through="RoleToUser", related_name="roles")
    role = models.TextField(choices=RoleTypes.choices, unique=True)

    class Meta:
        db_table = "Role"


class UserLink(models.Model):
    """This is a table that contains all the user links they can add to profile"""

    id = models.CharField(max_length=191, primary_key=True, default=uuid4_str)
    url = models.CharField(max_length=256)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "UserLink"


class User(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=uuid4_str)
    name = models.CharField(max_length=191)
    email = models.CharField(max_length=191, unique=True)
    email_verified = UtcDateTimeField(null=True, db_column="emailVerified")
    image = models.CharField(max_length=191, null=True)
    bio = models.CharField(max_length=1000, default="")
    """this is the bio that shows up on the profile page"""
    user_links: "models.ManyToManyField[UserLink, UserToUserLink]" = models.ManyToManyField("UserLink", through="UserToUserLink", related_name="user")
    """this is the twitter URL that shows up on the profile page"""
    status = models.TextField(choices=UserStatus.choices, default=UserStatus.ACTIVE)
    ban_reason = models.TextField(null=True, db_column="banReason")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "User"


class VerificationToken(models.Model):
    identifier = models.CharField(max_length=191)
    token = models.CharField(max_length=191, primary_key=True, default=None)
    expires = UtcDateTimeField()

    class Meta:
        db_table = "VerificationToken"
        constraints = [
            models.UniqueConstraint(fields=["identifier", "token"], name="VerificationToken_identifier_token_key"),
        ]


class Tag(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=uuid4_str)
    tag = models.TextField(choices=Tags.choices, unique=True)

    class Meta:
        db_table = "Tag"


class Challenge(models.Model):
    id = models.AutoField(primary_key=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    difficulty = models.TextField(choices=Difficulty.choices)
    name = models.CharField(max_length=191, unique=True)
    slug = models.CharField(max_length=191, unique=True)
    description = models.TextField()
    """this is the description to describe the challenge on description panel"""
    short_description = models.CharField(max_length=191, db_column="shortDescription")
    """this is a short blue thats shown on the explore card"""
    code = models.TextField()
    """Starter code for the challenge"""
    tests = models.TextField()
    """Test cases that need to be passed in order have the submission be accepted."""
    status = models.TextField(choices=ChallengeStatus.choices, default=ChallengeStatus.PENDING)
    tsconfig = models.JSONField(null=True)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="challenge", db_column="userId", db_index=False)
    tags: "models.ManyToManyField[Tag, ChallengeToTag]" = models.ManyToManyField("Tag", through="ChallengeToTag", related_name="challenges")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Challenge"
        indexes = [
            models.Index(fields=["user"]),
        ]


class Vote(models.Model):
    id = models.AutoField(primary_key=True)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="vote", db_column="userId", db_index=False)
    root_type = models.TextField(choices=VoteType.choices, db_column="rootType")
    challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="vote", db_column="challengeId", null=True, db_index=False)
    comment = models.ForeignKey("Comment", on_delete=models.SET_NULL, related_name="vote", db_column="commentId", null=True, db_index=False)
    submission = models.ForeignKey("Submission", on_delete=models.SET_NULL, related_name="vote", db_column="submissionId", null=True, db_index=False)
    shared_solution = models.ForeignKey("SharedSolution", on_delete=models.SET_NULL, related_name="vote", db_column="sharedSolutionId", null=True, db_index=False)

    class Meta:
        db_table = "Vote"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["challenge"]),
            models.Index(fields=["comment"]),
            models.Index(fields=["submission"]),
            models.Index(fields=["shared_solution"]),
        ]


class SharedSolution(models.Model):
    """This is a public solution created by a user for a challenge"""

    id = models.AutoField(primary_key=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    is_pinned = models.BooleanField(db_column="isPinned", default=False)
    title = models.CharField(max_length=191)
    description = models.TextField()
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="shared_solution", db_column="userId", null=True, db_index=False)
    challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="shared_solution", db_column="challengeId", null=True, db_index=False)

    class Meta:
        db_table = "SharedSolution"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["challenge"]),
        ]


class Submission(models.Model):
    """Every attempt a user makes to a challenge is stored here"""

    id = models.AutoField(primary_key=True)
    code = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    is_successful = models.BooleanField(db_column="isSuccessful")
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="submission", db_column="userId", db_index=False)
    challenge = models.ForeignKey("Challenge", on_delete=models.RESTRICT, related_name="submission", db_column="challengeId", db_index=False)

    class Meta:
        db_table = "Submission"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["challenge"]),
        ]


class Bookmark(models.Model):
    id = models.AutoField(primary_key=True)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="bookmark", db_column="userId", db_index=False)
    challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="bookmark", db_column="challengeId", null=True, db_index=False)
    created_at = UtcDateTimeField(null=True, db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "Bookmark"
        constraints = [
            models.UniqueConstraint(fields=["user", "challenge"], name="Bookmark_userId_challengeId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["challenge"]),
        ]


class ReportIssue(models.Model):
    id = models.AutoField(primary_key=True)
    report = models.ForeignKey("Report", on_delete=models.RESTRICT, related_name="issues", db_column="reportId", db_index=False)
    type = models.TextField(choices=IssueType.choices)

    class Meta:
        db_table = "ReportIssue"
        constraints = [
            models.UniqueConstraint(fields=["id"], name="ReportIssue_id_key"),
        ]
        indexes = [
            models.Index(fields=["report"]),
        ]


class Report(models.Model):
    id = models.AutoField(primary_key=True)
    text = models.TextField()
    type = models.TextField(choices=ReportType.choices)
    status = models.TextField(choices=ReportStatus.choices)
    reporter = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="reports", db_column="reporterId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="reports_about", db_column="userId", null=True, db_index=False)
    challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="report", db_column="challengeId", null=True, db_index=False)
    comment = models.ForeignKey("Comment", on_delete=models.SET_NULL, related_name="report", db_column="commentId", null=True, db_index=False)
    solution = models.ForeignKey("SharedSolution", on_delete=models.SET_NULL, related_name="report", db_column="solutionId", null=True, db_index=False)
    moderator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="moderated", db_column="moderatorId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Report"
        constraints = [
            models.UniqueConstraint(fields=["id"], name="Report_id_key"),
        ]
        indexes = [
            models.Index(fields=["comment"]),
            models.Index(fields=["reporter"]),
            models.Index(fields=["user"]),
            models.Index(fields=["challenge"]),
            models.Index(fields=["solution"]),
            models.Index(fields=["moderator"]),
        ]


class Comment(models.Model):
    id = models.AutoField(primary_key=True)
    text = models.TextField()
    visible = models.BooleanField(default=True)
    parent_comment = models.ForeignKey("self", on_delete=models.DO_NOTHING, related_name="replies", db_column="parentId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="comment", db_column="userId", db_index=False)
    root_type = models.TextField(choices=CommentRoot.choices, db_column="rootType", default=CommentRoot.CHALLENGE)
    root_challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="comment", db_column="rootChallengeId", null=True, db_index=False)
    root_solution = models.ForeignKey("SharedSolution", on_delete=models.SET_NULL, related_name="solution_comment", db_column="rootSolutionId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Comment"
        indexes = [
            models.Index(fields=["root_solution"]),
            models.Index(fields=["root_challenge"]),
            models.Index(fields=["parent_comment"]),
            models.Index(fields=["user"]),
        ]


class ImageUpload(models.Model):
    id = models.CharField(max_length=191, primary_key=True, default=uuid4_str)
    """the id of the db row"""
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="image_upload", db_column="userId", db_index=False)
    """the user id fk"""
    url = models.CharField(max_length=256)
    """the url to the s3 uploaded assets"""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ImageUpload"
        indexes = [
            models.Index(fields=["user"]),
        ]


class TrackChallenge(models.Model):
    id = models.AutoField(primary_key=True)
    order_id = models.IntegerField(db_column="orderId")
    challenge = models.ForeignKey("Challenge", on_delete=models.RESTRICT, related_name="track_challenge", db_column="challengeId", db_index=False)
    track = models.ForeignKey("Track", on_delete=models.SET_NULL, related_name="track_challenges", db_column="trackId", null=True, db_index=False)

    class Meta:
        db_table = "TrackChallenge"
        indexes = [
            models.Index(fields=["challenge"]),
            models.Index(fields=["track"]),
        ]


class Track(models.Model):
    id = models.AutoField(primary_key=True)
    name = models.CharField(max_length=191, unique=True)
    slug = models.CharField(max_length=191, unique=True)
    description = models.TextField()
    visible = models.BooleanField(default=True)
    """determines if a track can be seen on tracks page"""
    is_coming_soon = models.BooleanField(db_column="isComingSoon", default=False)
    """use this to display coming soon UI"""
    enrolled_users: "models.ManyToManyField[User, TrackToUser]" = models.ManyToManyField("User", through="TrackToUser", related_name="tracks")

    class Meta:
        db_table = "Track"


class BetaTokens(models.Model):
    id = models.AutoField(primary_key=True)
    token = models.CharField(max_length=191, unique=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="beta_tokens", db_column="userId", null=True, db_index=False)

    class Meta:
        db_table = "BetaTokens"
        indexes = [
            models.Index(fields=["user"]),
        ]


class Donation(models.Model):
    id = models.AutoField(primary_key=True)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="donations", db_column="userId", db_index=False)
    amount = models.IntegerField()
    stripe_checkout_session_id = models.CharField(max_length=191, unique=True, db_column="stripeCheckoutSessionId")
    created_at = UtcDateTimeField(null=True, db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "Donation"
        indexes = [
            models.Index(fields=["user"]),
        ]


class Notification(models.Model):
    id = models.AutoField(primary_key=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    is_read = models.BooleanField(db_column="isRead", default=False)
    type = models.TextField(choices=NotificationType.choices)
    to_user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="to_user_notification", db_column="toUserId", db_index=False)
    from_user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="from_user_notification", db_column="fromUserId", db_index=False)
    comment = models.ForeignKey("Comment", on_delete=models.SET_NULL, related_name="notification", db_column="commentId", null=True, db_index=False)
    content = models.CharField(max_length=191, null=True)
    """could be comment text or solution title to avoid having to join to grab this data"""
    blurb = models.CharField(max_length=191, null=True)
    """can be used to replace the blurb (eg: mention you in a comment)"""
    url = models.CharField(max_length=191)
    """use this instead of building the url on the fly (avoids a lot of joins and perf hits to build a string)"""
    challenge = models.ForeignKey("Challenge", on_delete=models.SET_NULL, related_name="notification", db_column="challengeId", null=True, db_index=False)
    shared_solution = models.ForeignKey("SharedSolution", on_delete=models.SET_NULL, related_name="notification", db_column="sharedSolutionId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Notification"
        indexes = [
            models.Index(fields=["to_user"]),
            models.Index(fields=["from_user"]),
            models.Index(fields=["comment"]),
            models.Index(fields=["challenge"]),
            models.Index(fields=["shared_solution"]),
        ]


class ShortURL(models.Model):
    id = models.AutoField(primary_key=True)
    original_url = models.TextField(db_column="originalUrl")
    """the original full url that we want to shorten e.g. https://www.url.com"""
    short_url_slug = models.CharField(max_length=191, unique=True, db_column="shortUrlSlug")
    """slug for the short url e.g. url.com/[shortUrlSlug]"""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    """the time the short url will expire"""
    clicks = models.IntegerField(default=0)
    """the number of times the short url has redirected to the original url"""
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="short_urls", db_column="userId", null=True, db_index=False)

    class Meta:
        db_table = "ShortURL"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["short_url_slug"]),
        ]


class RoleToUser(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Role", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("User", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_RoleToUser"
        indexes = [
            models.Index(fields=["b"]),
        ]


class UserToUserLink(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("User", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("UserLink", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_UserToUserLink"
        indexes = [
            models.Index(fields=["b"]),
        ]


class TrackToUser(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Track", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("User", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_TrackToUser"
        indexes = [
            models.Index(fields=["b"]),
        ]


class ChallengeToTag(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Challenge", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("Tag", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_ChallengeToTag"
        indexes = [
            models.Index(fields=["b"]),
        ]
