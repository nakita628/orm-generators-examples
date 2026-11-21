"""Constraints, indexes and table options SQLite can carry."""

from django.db import models
from django.db.models import F, Q
from django.db.models.functions import Lower, Upper


class Product(models.Model):
    sku = models.CharField(max_length=20)
    name = models.CharField(max_length=100)
    price = models.DecimalField(max_digits=10, decimal_places=2)
    discounted = models.DecimalField(max_digits=10, decimal_places=2, null=True)
    stock = models.IntegerField(default=0)
    active = models.BooleanField(default=True)
    category = models.CharField(max_length=30)

    class Meta:
        db_table = "catalog_products"
        ordering = ["name", "-price"]
        get_latest_by = "id"
        unique_together = [("category", "name")]
        constraints = [
            models.CheckConstraint(condition=Q(price__gte=0), name="price_not_negative"),
            models.CheckConstraint(
                condition=Q(discounted__isnull=True) | Q(discounted__lt=F("price")),
                name="discount_below_price",
                violation_error_message="A discount is below the price.",
            ),
            models.UniqueConstraint(Lower("sku"), name="sku_unique_case_insensitive"),
            models.UniqueConstraint(fields=["name"], condition=Q(active=True), name="one_active_name"),
        ]
        indexes = [
            models.Index(fields=["category", "-price"], name="category_price"),
            models.Index(Upper("name"), name="name_upper"),
            models.Index(fields=["stock"], condition=Q(stock__lte=5), name="low_stock"),
        ]


class Legacy(models.Model):
    """A table another system owns."""

    code = models.CharField(max_length=10, primary_key=True)
    label = models.CharField(max_length=50)

    class Meta:
        managed = False
        db_table = "legacy_codes"


class LegacyView(models.Model):
    total = models.IntegerField()

    class Meta:
        managed = False
        db_table = "v_totals"


class Booking(models.Model):
    pk = models.CompositePrimaryKey("room", "day")
    room = models.CharField(max_length=10)
    day = models.DateField()
    guest = models.CharField(max_length=50)
