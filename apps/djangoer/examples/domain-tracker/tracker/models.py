"""An issue tracker: projects, workflows of states, issues with labels and links, sprints,
comments, attachments and time logs."""

from django.conf import settings
from django.db import models


class Project(models.Model):
    key = models.CharField(max_length=10, unique=True, help_text="Prefix of issue keys, e.g. WEB")
    name = models.CharField(max_length=100)
    lead = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, related_name="led_projects")
    members = models.ManyToManyField(settings.AUTH_USER_MODEL, related_name="projects", blank=True)
    next_number = models.PositiveIntegerField(default=1)
    archived = models.BooleanField(default=False)


class State(models.Model):
    class Category(models.TextChoices):
        TODO = "todo", "To do"
        IN_PROGRESS = "in_progress", "In progress"
        DONE = "done", "Done"

    project = models.ForeignKey(Project, on_delete=models.CASCADE, related_name="states")
    name = models.CharField(max_length=30)
    category = models.CharField(max_length=12, choices=Category)
    position = models.PositiveSmallIntegerField()
    transitions = models.ManyToManyField("self", symmetrical=False, blank=True, related_name="+")

    class Meta:
        unique_together = [("project", "name")]
        ordering = ["project", "position"]


class Label(models.Model):
    project = models.ForeignKey(Project, on_delete=models.CASCADE, related_name="labels")
    name = models.CharField(max_length=30)
    color = models.CharField(max_length=7, default="#888888")

    class Meta:
        unique_together = [("project", "name")]


class Sprint(models.Model):
    project = models.ForeignKey(Project, on_delete=models.CASCADE, related_name="sprints")
    name = models.CharField(max_length=50)
    goal = models.TextField(blank=True)
    starts_on = models.DateField()
    ends_on = models.DateField()

    class Meta:
        constraints = [models.CheckConstraint(condition=models.Q(ends_on__gt=models.F("starts_on")), name="sprint_ends_after_start")]


class Issue(models.Model):
    class Type(models.TextChoices):
        BUG = "bug", "Bug"
        TASK = "task", "Task"
        STORY = "story", "Story"
        EPIC = "epic", "Epic"

    class Priority(models.IntegerChoices):
        LOWEST = 1
        LOW = 2
        MEDIUM = 3
        HIGH = 4
        HIGHEST = 5

    project = models.ForeignKey(Project, on_delete=models.CASCADE, related_name="issues")
    number = models.PositiveIntegerField()
    type = models.CharField(max_length=5, choices=Type, default=Type.TASK)
    title = models.CharField(max_length=255)
    description = models.TextField(blank=True)
    state = models.ForeignKey(State, on_delete=models.PROTECT, related_name="issues")
    priority = models.PositiveSmallIntegerField(choices=Priority, default=Priority.MEDIUM)
    reporter = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.PROTECT, related_name="reported_issues")
    assignee = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, blank=True, related_name="assigned_issues")
    watchers = models.ManyToManyField(settings.AUTH_USER_MODEL, related_name="watched_issues", blank=True)
    epic = models.ForeignKey("self", on_delete=models.SET_NULL, null=True, blank=True, related_name="children", limit_choices_to={"type": "epic"})
    labels = models.ManyToManyField(Label, related_name="issues", blank=True)
    sprint = models.ForeignKey(Sprint, on_delete=models.SET_NULL, null=True, blank=True, related_name="issues")
    story_points = models.PositiveSmallIntegerField(null=True, blank=True)
    due_on = models.DateField(null=True, blank=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    links = models.ManyToManyField("self", through="IssueLink", symmetrical=False, related_name="linked_from")

    class Meta:
        constraints = [models.UniqueConstraint(fields=["project", "number"], name="issue_key_unique")]
        indexes = [
            models.Index(fields=["assignee", "state"], name="issue_assignee_state"),
            models.Index(fields=["project", "-updated_at"], name="issue_recent"),
        ]


class IssueLink(models.Model):
    class Kind(models.TextChoices):
        BLOCKS = "blocks", "blocks"
        DUPLICATES = "duplicates", "duplicates"
        RELATES = "relates", "relates to"

    source = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="outgoing_links")
    target = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="incoming_links")
    kind = models.CharField(max_length=10, choices=Kind)

    class Meta:
        unique_together = [("source", "target", "kind")]


class Comment(models.Model):
    issue = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="comments")
    author = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    body = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
    edited_at = models.DateTimeField(null=True, blank=True)


class Attachment(models.Model):
    issue = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="attachments")
    comment = models.ForeignKey(Comment, on_delete=models.SET_NULL, null=True, blank=True, related_name="attachments")
    file = models.FileField(upload_to="attachments/")
    size = models.PositiveBigIntegerField()
    uploaded_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True)


class WorkLog(models.Model):
    issue = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="work_logs")
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE)
    spent = models.DurationField()
    started_at = models.DateTimeField()
    note = models.CharField(max_length=200, blank=True)


class Change(models.Model):
    """One field of an issue changed, for its history."""

    issue = models.ForeignKey(Issue, on_delete=models.CASCADE, related_name="changes")
    user = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True)
    field = models.CharField(max_length=50)
    old_value = models.TextField(blank=True)
    new_value = models.TextField(blank=True)
    at = models.DateTimeField(auto_now_add=True)
