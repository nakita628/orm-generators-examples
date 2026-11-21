"""Hotel bookings on PostgreSQL: date ranges that must not overlap, rates by season, guests."""

from django.conf import settings
from django.contrib.postgres.constraints import ExclusionConstraint
from django.contrib.postgres.fields import ArrayField, DateRangeField, RangeOperators
from django.contrib.postgres.indexes import GistIndex
from django.db import models


class Hotel(models.Model):
    name = models.CharField(max_length=100)
    city = models.CharField(max_length=100)
    stars = models.PositiveSmallIntegerField(null=True, blank=True)
    amenities = ArrayField(models.CharField(max_length=30), default=list, blank=True)
    timezone = models.CharField(max_length=50, default="UTC")


class RoomType(models.Model):
    hotel = models.ForeignKey(Hotel, on_delete=models.CASCADE, related_name="room_types")
    name = models.CharField(max_length=50)
    capacity = models.PositiveSmallIntegerField()
    beds = models.JSONField(default=dict, help_text='e.g. {"double": 1, "single": 1}')

    class Meta:
        unique_together = [("hotel", "name")]


class Room(models.Model):
    room_type = models.ForeignKey(RoomType, on_delete=models.PROTECT, related_name="rooms")
    number = models.CharField(max_length=10)
    floor = models.SmallIntegerField()
    out_of_service = models.BooleanField(default=False)


class Season(models.Model):
    hotel = models.ForeignKey(Hotel, on_delete=models.CASCADE, related_name="seasons")
    name = models.CharField(max_length=50)
    dates = DateRangeField()

    class Meta:
        constraints = [
            ExclusionConstraint(
                name="season_no_overlap",
                expressions=[("dates", RangeOperators.OVERLAPS), ("hotel", RangeOperators.EQUAL)],
            )
        ]


class Rate(models.Model):
    room_type = models.ForeignKey(RoomType, on_delete=models.CASCADE, related_name="rates")
    season = models.ForeignKey(Season, on_delete=models.CASCADE, related_name="rates")
    nightly_price = models.DecimalField(max_digits=8, decimal_places=2)

    class Meta:
        unique_together = [("room_type", "season")]


class Guest(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, blank=True)
    full_name = models.CharField(max_length=150)
    email = models.EmailField()
    phone = models.CharField(max_length=30, blank=True)
    preferences = models.JSONField(default=dict, blank=True)


class Reservation(models.Model):
    class Status(models.TextChoices):
        HELD = "held", "Held"
        CONFIRMED = "confirmed", "Confirmed"
        CHECKED_IN = "checked_in", "Checked in"
        CHECKED_OUT = "checked_out", "Checked out"
        CANCELLED = "cancelled", "Cancelled"

    room = models.ForeignKey(Room, on_delete=models.PROTECT, related_name="reservations")
    guest = models.ForeignKey(Guest, on_delete=models.PROTECT, related_name="reservations")
    stay = DateRangeField()
    adults = models.PositiveSmallIntegerField(default=1)
    children = models.PositiveSmallIntegerField(default=0)
    status = models.CharField(max_length=12, choices=Status, default=Status.HELD)
    total_price = models.DecimalField(max_digits=10, decimal_places=2)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        indexes = [GistIndex(fields=["stay"], name="reservation_stay_gist")]
        constraints = [
            ExclusionConstraint(
                name="reservation_no_double_booking",
                expressions=[("stay", RangeOperators.OVERLAPS), ("room", RangeOperators.EQUAL)],
                condition=~models.Q(status="cancelled"),
            ),
            models.CheckConstraint(condition=models.Q(adults__gte=1), name="reservation_at_least_one_adult"),
        ]


class Payment(models.Model):
    reservation = models.ForeignKey(Reservation, on_delete=models.CASCADE, related_name="payments")
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    is_deposit = models.BooleanField(default=False)
    paid_at = models.DateTimeField()
