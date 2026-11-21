"""Every field type Django ships, with the options that change the column."""

import uuid
from decimal import Decimal

from django.core.validators import MinValueValidator
from django.db import models
from django.db.models import F, Value
from django.db.models.functions import Concat, Lower, Now


class Status(models.TextChoices):
    DRAFT = "draft", "Draft"
    REVIEW = "in_review", "In review"
    PUBLISHED = "published", "Published"


class Priority(models.IntegerChoices):
    LOW = 1
    NORMAL = 2
    HIGH = 3


class Everything(models.Model):
    """A row with one column of each kind."""

    # Numbers
    small = models.SmallIntegerField()
    integer = models.IntegerField(validators=[MinValueValidator(0)])
    big = models.BigIntegerField()
    positive_small = models.PositiveSmallIntegerField()
    positive = models.PositiveIntegerField()
    positive_big = models.PositiveBigIntegerField()
    floating = models.FloatField()
    price = models.DecimalField(max_digits=10, decimal_places=2, default=Decimal("0.00"))
    # Text
    char = models.CharField(max_length=100, help_text="Up to 100 characters")
    text = models.TextField(blank=True)
    slug = models.SlugField(unique=True)
    email = models.EmailField()
    url = models.URLField(max_length=500)
    ip = models.GenericIPAddressField(protocol="both", unpack_ipv4=True)
    path = models.FilePathField(path="/tmp", match=r".*\.txt$", recursive=True, blank=True)
    # Dates and times
    date = models.DateField()
    time = models.TimeField(null=True)
    datetime = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
    duration = models.DurationField(null=True)
    # Others
    flag = models.BooleanField(default=False)
    maybe = models.BooleanField(null=True)
    uuid = models.UUIDField(default=uuid.uuid4, editable=False, unique=True)
    data = models.JSONField(default=dict, blank=True)
    blob = models.BinaryField(max_length=1024, null=True)
    file = models.FileField(upload_to="files/%Y/%m/", max_length=255, blank=True)
    image = models.ImageField(upload_to="images/", height_field="image_height", width_field="image_width", blank=True)
    image_height = models.PositiveIntegerField(null=True, editable=False)
    image_width = models.PositiveIntegerField(null=True, editable=False)
    # Choices
    status = models.CharField(max_length=20, choices=Status, default=Status.DRAFT)
    priority = models.IntegerField(choices=Priority, default=Priority.NORMAL)
    grouped = models.CharField(
        max_length=10,
        choices={"Audio": {"vinyl": "Vinyl", "cd": "CD"}, "Video": {"vhs": "VHS Tape", "dvd": "DVD"}, "unknown": "Unknown"},
        blank=True,
    )
    # What the database computes, and what it stores for the column
    full = models.GeneratedField(
        expression=Concat(F("char"), Value(" "), F("slug")),
        output_field=models.CharField(max_length=201),
        db_persist=True,
    )
    lower_email = models.GeneratedField(expression=Lower("email"), output_field=models.EmailField(), db_persist=False)
    created_in_db = models.DateTimeField(db_default=Now())
    counter = models.IntegerField(db_default=0, db_comment="How many times the row was seen")
    collated = models.CharField(max_length=50, db_collation="NOCASE", blank=True)
    renamed = models.CharField(max_length=10, db_column="RenamedColumn", blank=True)
    indexed = models.CharField(max_length=10, db_index=True, blank=True)

    class Meta:
        db_table_comment = "One column of every kind"
        verbose_name_plural = "everything"


class AutoKeys(models.Model):
    id = models.SmallAutoField(primary_key=True)
    name = models.CharField(max_length=10)


class BigKeys(models.Model):
    id = models.BigAutoField(primary_key=True)


class NaturalKey(models.Model):
    code = models.CharField(max_length=3, primary_key=True)
    name = models.CharField(max_length=50)


class UuidKey(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4)
    natural = models.ForeignKey(NaturalKey, on_delete=models.PROTECT)
