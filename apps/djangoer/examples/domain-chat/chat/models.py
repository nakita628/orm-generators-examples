"""A chat app: workspaces, channels and direct messages, threads, reactions, read state.
UUID keys throughout, as for ids that go to clients."""

import uuid

from django.conf import settings
from django.db import models


class UUIDModel(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        abstract = True


class Workspace(UUIDModel):
    name = models.CharField(max_length=80)
    domain = models.SlugField(unique=True)
    members = models.ManyToManyField(settings.AUTH_USER_MODEL, through="WorkspaceMember", related_name="workspaces")


class WorkspaceMember(models.Model):
    workspace = models.ForeignKey(Workspace, on_delete=models.CASCADE)
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    display_name = models.CharField(max_length=80, blank=True)
    is_admin = models.BooleanField(default=False)
    status_text = models.CharField(max_length=100, blank=True)
    status_emoji = models.CharField(max_length=32, blank=True)

    class Meta:
        unique_together = [("workspace", "user")]


class Conversation(UUIDModel):
    class Kind(models.TextChoices):
        CHANNEL = "channel", "Public channel"
        PRIVATE = "private", "Private channel"
        DIRECT = "direct", "Direct message"

    workspace = models.ForeignKey(Workspace, on_delete=models.CASCADE, related_name="conversations")
    kind = models.CharField(max_length=7, choices=Kind)
    name = models.CharField(max_length=80, blank=True, help_text="Empty for direct messages")
    topic = models.CharField(max_length=250, blank=True)
    created_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, related_name="+")
    participants = models.ManyToManyField(settings.AUTH_USER_MODEL, through="Participant", related_name="conversations")
    archived_at = models.DateTimeField(null=True, blank=True)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=["workspace", "name"], condition=~models.Q(kind="direct"), name="channel_name_unique"),
            models.CheckConstraint(condition=models.Q(kind="direct") | ~models.Q(name=""), name="channel_has_name"),
        ]


class Participant(models.Model):
    conversation = models.ForeignKey(Conversation, on_delete=models.CASCADE)
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    joined_at = models.DateTimeField(auto_now_add=True)
    last_read = models.ForeignKey("Message", on_delete=models.SET_NULL, null=True, blank=True, related_name="+")
    muted = models.BooleanField(default=False)

    class Meta:
        unique_together = [("conversation", "user")]


class Message(UUIDModel):
    conversation = models.ForeignKey(Conversation, on_delete=models.CASCADE, related_name="messages")
    author = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, related_name="messages")
    thread_root = models.ForeignKey("self", on_delete=models.CASCADE, null=True, blank=True, related_name="replies")
    text = models.TextField()
    edited_at = models.DateTimeField(null=True, blank=True)
    deleted_at = models.DateTimeField(null=True, blank=True)
    pinned = models.BooleanField(default=False)

    class Meta:
        indexes = [models.Index(fields=["conversation", "created_at"], name="message_history")]


class Reaction(models.Model):
    message = models.ForeignKey(Message, on_delete=models.CASCADE, related_name="reactions")
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    emoji = models.CharField(max_length=32)

    class Meta:
        constraints = [models.UniqueConstraint(fields=["message", "user", "emoji"], name="react_once")]


class File(UUIDModel):
    message = models.ForeignKey(Message, on_delete=models.CASCADE, related_name="files")
    name = models.CharField(max_length=255)
    content_type = models.CharField(max_length=100)
    size = models.PositiveBigIntegerField()
    storage_key = models.CharField(max_length=300, unique=True)
