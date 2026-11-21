"""A learning platform: courses made of modules and lessons, enrolments, progress, quizzes and
certificates."""

from django.conf import settings
from django.db import models


class Instructor(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="instructor")
    bio = models.TextField(blank=True)
    website = models.URLField(blank=True)


class Course(models.Model):
    class Level(models.IntegerChoices):
        BEGINNER = 1
        INTERMEDIATE = 2
        ADVANCED = 3

    title = models.CharField(max_length=200)
    slug = models.SlugField(unique=True)
    summary = models.TextField()
    level = models.PositiveSmallIntegerField(choices=Level, default=Level.BEGINNER)
    instructors = models.ManyToManyField(Instructor, related_name="courses")
    prerequisites = models.ManyToManyField("self", symmetrical=False, blank=True, related_name="unlocks")
    price = models.DecimalField(max_digits=8, decimal_places=2, default=0)
    published_at = models.DateTimeField(null=True, blank=True)


class Module(models.Model):
    course = models.ForeignKey(Course, on_delete=models.CASCADE, related_name="modules")
    title = models.CharField(max_length=200)

    class Meta:
        order_with_respect_to = "course"


class Lesson(models.Model):
    class Kind(models.TextChoices):
        VIDEO = "video", "Video"
        TEXT = "text", "Text"
        QUIZ = "quiz", "Quiz"

    module = models.ForeignKey(Module, on_delete=models.CASCADE, related_name="lessons")
    title = models.CharField(max_length=200)
    kind = models.CharField(max_length=5, choices=Kind)
    body = models.TextField(blank=True)
    video_url = models.URLField(blank=True)
    duration = models.DurationField(null=True, blank=True)
    is_preview = models.BooleanField(default=False, help_text="Open to visitors who have not enrolled")

    class Meta:
        order_with_respect_to = "module"


class Enrollment(models.Model):
    student = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="enrollments")
    course = models.ForeignKey(Course, on_delete=models.CASCADE, related_name="enrollments")
    enrolled_at = models.DateTimeField(auto_now_add=True)
    completed_at = models.DateTimeField(null=True, blank=True)
    completed_lessons = models.ManyToManyField(Lesson, through="LessonProgress", related_name="+")

    class Meta:
        constraints = [models.UniqueConstraint(fields=["student", "course"], name="enroll_once")]


class LessonProgress(models.Model):
    enrollment = models.ForeignKey(Enrollment, on_delete=models.CASCADE)
    lesson = models.ForeignKey(Lesson, on_delete=models.CASCADE)
    completed_at = models.DateTimeField(auto_now_add=True)
    seconds_watched = models.PositiveIntegerField(default=0)


class Question(models.Model):
    lesson = models.ForeignKey(Lesson, on_delete=models.CASCADE, related_name="questions", limit_choices_to={"kind": "quiz"})
    text = models.TextField()
    allows_many = models.BooleanField(default=False)
    points = models.PositiveSmallIntegerField(default=1)


class Choice(models.Model):
    question = models.ForeignKey(Question, on_delete=models.CASCADE, related_name="choices")
    text = models.CharField(max_length=300)
    is_correct = models.BooleanField(default=False)


class Attempt(models.Model):
    enrollment = models.ForeignKey(Enrollment, on_delete=models.CASCADE, related_name="attempts")
    lesson = models.ForeignKey(Lesson, on_delete=models.CASCADE, related_name="attempts")
    started_at = models.DateTimeField(auto_now_add=True)
    submitted_at = models.DateTimeField(null=True, blank=True)
    score = models.DecimalField(max_digits=5, decimal_places=2, null=True, blank=True)


class Answer(models.Model):
    attempt = models.ForeignKey(Attempt, on_delete=models.CASCADE, related_name="answers")
    question = models.ForeignKey(Question, on_delete=models.CASCADE)
    choices = models.ManyToManyField(Choice)

    class Meta:
        unique_together = [("attempt", "question")]


class Certificate(models.Model):
    enrollment = models.OneToOneField(Enrollment, on_delete=models.PROTECT, related_name="certificate")
    code = models.CharField(max_length=20, unique=True)
    issued_at = models.DateTimeField(auto_now_add=True)
