"""The same kinds of models on MySQL: its column types (unsigned integers, char(32) UUIDs,
longtext, json), collations, comments, prefix-free indexes and functional constraints."""

import uuid

from django.db import models
from django.db.models.functions import Lower


class Customer(models.Model):
    uuid = models.UUIDField(default=uuid.uuid4, unique=True, editable=False)
    email = models.EmailField(max_length=191, db_comment="191 characters fit a utf8mb4 index of 767 bytes")
    name = models.CharField(max_length=100, db_collation="utf8mb4_unicode_ci")
    code = models.CharField(max_length=20, unique=True, db_collation="utf8mb4_bin")
    points = models.PositiveIntegerField(default=0)
    age = models.PositiveSmallIntegerField(null=True, blank=True)
    settings = models.JSONField(default=dict)
    notes = models.TextField(blank=True)
    signed_up = models.DateTimeField(auto_now_add=True)
    last_ip = models.GenericIPAddressField(null=True, blank=True)

    class Meta:
        db_table_comment = "People who buy"
        constraints = [models.UniqueConstraint(Lower("email"), name="customer_email_ci_unique")]


class Product(models.Model):
    class Kind(models.TextChoices):
        PHYSICAL = "physical"
        DIGITAL = "digital"

    name = models.CharField(max_length=200)
    kind = models.CharField(max_length=8, choices=Kind, default=Kind.PHYSICAL)
    price = models.DecimalField(max_digits=10, decimal_places=2)
    stock = models.PositiveBigIntegerField(default=0)
    weight = models.FloatField(null=True, blank=True)
    tags = models.JSONField(default=list, blank=True)
    image = models.BinaryField(null=True, blank=True, editable=True)
    active = models.BooleanField(default=True, db_default=True)
    released = models.DateField(null=True, blank=True)
    duration = models.DurationField(null=True, blank=True)

    class Meta:
        indexes = [models.Index(fields=["kind", "-price"], name="product_kind_price")]
        constraints = [models.CheckConstraint(condition=models.Q(price__gte=0), name="product_price_not_negative")]


class Order(models.Model):
    customer = models.ForeignKey(Customer, on_delete=models.PROTECT, related_name="orders")
    products = models.ManyToManyField(Product, through="OrderItem")
    placed_at = models.DateTimeField(auto_now_add=True)
    total = models.DecimalField(max_digits=12, decimal_places=2)
    total_with_tax = models.GeneratedField(
        expression=models.F("total") * 1.1,
        output_field=models.DecimalField(max_digits=14, decimal_places=2),
        db_persist=False,
    )


class OrderItem(models.Model):
    order = models.ForeignKey(Order, on_delete=models.CASCADE)
    product = models.ForeignKey(Product, on_delete=models.PROTECT)
    quantity = models.PositiveSmallIntegerField(default=1)
    unit_price = models.DecimalField(max_digits=10, decimal_places=2)

    class Meta:
        unique_together = [("order", "product")]
