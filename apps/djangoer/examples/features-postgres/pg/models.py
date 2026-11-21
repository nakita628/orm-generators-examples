"""What django.contrib.postgres adds: arrays, hstore, ranges, search, and the indexes and
constraints only PostgreSQL has."""

from django.contrib.postgres.constraints import ExclusionConstraint
from django.contrib.postgres.fields import (
    ArrayField,
    BigIntegerRangeField,
    DateRangeField,
    DateTimeRangeField,
    DecimalRangeField,
    HStoreField,
    IntegerRangeField,
    RangeOperators,
)
from django.contrib.postgres.indexes import (
    BloomIndex,
    BrinIndex,
    BTreeIndex,
    GinIndex,
    GistIndex,
    HashIndex,
    OpClass,
    SpGistIndex,
)
from django.contrib.postgres.search import SearchVectorField
from django.db import models
from django.db.models import Q
from django.db.models.functions import Lower


class Document(models.Model):
    title = models.CharField(max_length=200, db_collation="und-x-icu")
    tags = ArrayField(models.CharField(max_length=30), default=list, blank=True)
    matrix = ArrayField(ArrayField(models.IntegerField(), size=3), size=3, null=True)
    attributes = HStoreField(default=dict)
    search = SearchVectorField(null=True)
    data = models.JSONField(default=dict)
    body = models.TextField()

    class Meta:
        indexes = [
            GinIndex(fields=["search"], name="document_search"),
            GinIndex(fields=["tags"], name="document_tags"),
            GinIndex(fields=["data"], name="document_data", opclasses=["jsonb_path_ops"]),
            GinIndex(OpClass(Lower("title"), name="gin_trgm_ops"), name="document_title_trgm"),
            HashIndex(fields=["body"], name="document_body_hash"),
            BTreeIndex(fields=["title"], name="document_title", fillfactor=70),
            BloomIndex(fields=["title", "body"], name="document_bloom", length=80),
        ]


class Reading(models.Model):
    sensor = models.CharField(max_length=20)
    taken = models.DateTimeField()
    value = models.FloatField()

    class Meta:
        indexes = [
            BrinIndex(fields=["taken"], name="reading_taken", autosummarize=True, pages_per_range=16),
            models.Index(fields=["sensor"], include=["value"], name="reading_sensor_covering"),
        ]


class Room(models.Model):
    number = models.CharField(max_length=10, unique=True)


class Reservation(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE)
    during = DateTimeRangeField()
    nights = DateRangeField()
    guests = IntegerRangeField()
    budget = DecimalRangeField(null=True)
    ids = BigIntegerRangeField(null=True)
    cancelled = models.BooleanField(default=False)
    code = models.CharField(max_length=10)
    parent = models.ForeignKey("self", null=True, on_delete=models.CASCADE)

    class Meta:
        indexes = [GistIndex(fields=["during"], name="reservation_during"), SpGistIndex(fields=["nights"], name="reservation_nights")]
        constraints = [
            ExclusionConstraint(
                name="exclude_overlapping_reservations",
                expressions=[("during", RangeOperators.OVERLAPS), ("room", RangeOperators.EQUAL)],
                condition=Q(cancelled=False),
            ),
            models.UniqueConstraint(fields=["code"], name="code_unique_deferred", deferrable=models.Deferrable.DEFERRED),
            models.UniqueConstraint(fields=["code", "room"], name="code_room_covering", include=["guests"]),
            models.UniqueConstraint(fields=["room", "parent"], name="room_parent_unique", nulls_distinct=False),
        ]
