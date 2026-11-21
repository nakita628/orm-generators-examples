"""A social network: profiles, follows, posts with media, likes and comments on anything,
hashtags, blocks and notifications."""

from django.conf import settings
from django.contrib.contenttypes.fields import GenericForeignKey, GenericRelation
from django.contrib.contenttypes.models import ContentType
from django.db import models


class Profile(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, primary_key=True)
    handle = models.CharField(max_length=30, unique=True)
    display_name = models.CharField(max_length=50)
    bio = models.CharField(max_length=160, blank=True)
    avatar = models.URLField(blank=True)
    is_private = models.BooleanField(default=False)
    following = models.ManyToManyField("self", through="Follow", symmetrical=False, related_name="followers")
    blocked = models.ManyToManyField("self", symmetrical=False, related_name="blocked_by", blank=True)


class Follow(models.Model):
    follower = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="+")
    followee = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="+")
    approved = models.BooleanField(default=True, help_text="False while a private profile has not accepted")
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=["follower", "followee"], name="follow_once"),
            models.CheckConstraint(condition=~models.Q(follower=models.F("followee")), name="no_self_follow"),
        ]


class Hashtag(models.Model):
    name = models.CharField(max_length=100, unique=True)


class Like(models.Model):
    profile = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="likes")
    content_type = models.ForeignKey(ContentType, on_delete=models.CASCADE)
    object_id = models.PositiveBigIntegerField()
    target = GenericForeignKey("content_type", "object_id")
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        constraints = [models.UniqueConstraint(fields=["profile", "content_type", "object_id"], name="like_once")]


class Post(models.Model):
    class Visibility(models.TextChoices):
        PUBLIC = "public", "Public"
        FOLLOWERS = "followers", "Followers only"
        MENTIONED = "mentioned", "Mentioned only"

    author = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="posts")
    text = models.CharField(max_length=500)
    reply_to = models.ForeignKey("self", on_delete=models.SET_NULL, null=True, blank=True, related_name="replies")
    repost_of = models.ForeignKey("self", on_delete=models.CASCADE, null=True, blank=True, related_name="reposts")
    visibility = models.CharField(max_length=10, choices=Visibility, default=Visibility.PUBLIC)
    hashtags = models.ManyToManyField(Hashtag, related_name="posts", blank=True)
    mentions = models.ManyToManyField(Profile, related_name="mentioned_in", blank=True)
    likes = GenericRelation(Like, related_query_name="post")
    created_at = models.DateTimeField(auto_now_add=True)
    edited_at = models.DateTimeField(null=True, blank=True)

    class Meta:
        ordering = ["-created_at"]
        indexes = [models.Index(fields=["author", "-created_at"], name="post_timeline")]


class Media(models.Model):
    class Kind(models.TextChoices):
        IMAGE = "image", "Image"
        VIDEO = "video", "Video"
        GIF = "gif", "GIF"

    post = models.ForeignKey(Post, on_delete=models.CASCADE, related_name="media")
    kind = models.CharField(max_length=5, choices=Kind)
    file = models.FileField(upload_to="media/%Y/%m/")
    alt_text = models.CharField(max_length=400, blank=True)
    width = models.PositiveIntegerField(null=True)
    height = models.PositiveIntegerField(null=True)

    class Meta:
        verbose_name_plural = "media"
        order_with_respect_to = "post"


class Comment(models.Model):
    post = models.ForeignKey(Post, on_delete=models.CASCADE, related_name="comments")
    author = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="comments")
    parent = models.ForeignKey("self", on_delete=models.CASCADE, null=True, blank=True, related_name="children")
    text = models.TextField()
    likes = GenericRelation(Like, related_query_name="comment")
    created_at = models.DateTimeField(auto_now_add=True)


class Notification(models.Model):
    class Verb(models.TextChoices):
        FOLLOWED = "followed"
        LIKED = "liked"
        REPLIED = "replied"
        MENTIONED = "mentioned"
        REPOSTED = "reposted"

    recipient = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="notifications")
    actor = models.ForeignKey(Profile, on_delete=models.CASCADE, related_name="+")
    verb = models.CharField(max_length=10, choices=Verb)
    target_type = models.ForeignKey(ContentType, on_delete=models.CASCADE, null=True, blank=True)
    target_id = models.PositiveBigIntegerField(null=True, blank=True)
    target = GenericForeignKey("target_type", "target_id")
    read_at = models.DateTimeField(null=True, blank=True)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        indexes = [
            models.Index(fields=["recipient", "-created_at"], name="notification_inbox"),
            models.Index(fields=["recipient"], condition=models.Q(read_at__isnull=True), name="notification_unread"),
        ]
