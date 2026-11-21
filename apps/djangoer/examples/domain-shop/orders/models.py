"""An online shop's customers, carts, orders, payments and shipments."""

import uuid

from django.conf import settings
from django.db import models


class Address(models.Model):
    customer = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="addresses")
    full_name = models.CharField(max_length=100)
    line1 = models.CharField(max_length=200)
    line2 = models.CharField(max_length=200, blank=True)
    city = models.CharField(max_length=100)
    postal_code = models.CharField(max_length=20)
    country = models.CharField(max_length=2)
    is_default = models.BooleanField(default=False)

    class Meta:
        verbose_name_plural = "addresses"
        constraints = [
            models.UniqueConstraint(
                fields=["customer"], condition=models.Q(is_default=True), name="one_default_address_per_customer"
            )
        ]


class Cart(models.Model):
    customer = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, null=True, blank=True)
    session_key = models.CharField(max_length=40, blank=True, db_index=True)
    created_at = models.DateTimeField(auto_now_add=True)


class CartItem(models.Model):
    cart = models.ForeignKey(Cart, on_delete=models.CASCADE, related_name="items")
    variant = models.ForeignKey("catalog.Variant", on_delete=models.CASCADE)
    quantity = models.PositiveSmallIntegerField(default=1)

    class Meta:
        unique_together = [("cart", "variant")]


class Coupon(models.Model):
    code = models.CharField(max_length=30, unique=True)
    percent_off = models.PositiveSmallIntegerField(null=True, blank=True)
    amount_off = models.DecimalField(max_digits=10, decimal_places=2, null=True, blank=True)
    valid_from = models.DateTimeField()
    valid_until = models.DateTimeField(null=True, blank=True)
    max_uses = models.PositiveIntegerField(null=True, blank=True)

    class Meta:
        constraints = [
            models.CheckConstraint(
                condition=models.Q(percent_off__isnull=True) ^ models.Q(amount_off__isnull=True),
                name="coupon_percent_xor_amount",
            ),
            models.CheckConstraint(condition=models.Q(percent_off__lte=100), name="coupon_percent_at_most_100"),
        ]


class Order(models.Model):
    class Status(models.TextChoices):
        PENDING = "pending", "Pending payment"
        PAID = "paid", "Paid"
        SHIPPED = "shipped", "Shipped"
        DELIVERED = "delivered", "Delivered"
        CANCELLED = "cancelled", "Cancelled"
        REFUNDED = "refunded", "Refunded"

    number = models.UUIDField(default=uuid.uuid4, unique=True, editable=False)
    customer = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.PROTECT, related_name="orders")
    status = models.CharField(max_length=10, choices=Status, default=Status.PENDING)
    shipping_address = models.ForeignKey(Address, on_delete=models.PROTECT, related_name="+")
    billing_address = models.ForeignKey(Address, on_delete=models.PROTECT, related_name="+")
    coupon = models.ForeignKey(Coupon, on_delete=models.SET_NULL, null=True, blank=True, related_name="orders")
    subtotal = models.DecimalField(max_digits=12, decimal_places=2)
    discount = models.DecimalField(max_digits=12, decimal_places=2, default=0)
    shipping_cost = models.DecimalField(max_digits=10, decimal_places=2, default=0)
    total = models.GeneratedField(
        expression=models.F("subtotal") - models.F("discount") + models.F("shipping_cost"),
        output_field=models.DecimalField(max_digits=12, decimal_places=2),
        db_persist=True,
    )
    placed_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        ordering = ["-placed_at"]
        indexes = [models.Index(fields=["customer", "-placed_at"], name="order_customer_recent")]


class OrderLine(models.Model):
    order = models.ForeignKey(Order, on_delete=models.CASCADE, related_name="lines")
    variant = models.ForeignKey("catalog.Variant", on_delete=models.SET_NULL, null=True)
    # Kept as sold, so the order still reads right when the product changes
    product_name = models.CharField(max_length=200)
    sku = models.CharField(max_length=40)
    unit_price = models.DecimalField(max_digits=10, decimal_places=2)
    quantity = models.PositiveSmallIntegerField()


class Payment(models.Model):
    class Method(models.TextChoices):
        CARD = "card", "Card"
        BANK = "bank", "Bank transfer"
        WALLET = "wallet", "Wallet"

    order = models.ForeignKey(Order, on_delete=models.CASCADE, related_name="payments")
    method = models.CharField(max_length=10, choices=Method)
    amount = models.DecimalField(max_digits=12, decimal_places=2)
    provider_reference = models.CharField(max_length=100, unique=True)
    succeeded = models.BooleanField(default=False)
    raw_response = models.JSONField(default=dict, blank=True)
    created_at = models.DateTimeField(auto_now_add=True)


class Shipment(models.Model):
    order = models.ForeignKey(Order, on_delete=models.CASCADE, related_name="shipments")
    lines = models.ManyToManyField(OrderLine, related_name="shipments")
    carrier = models.CharField(max_length=50)
    tracking_number = models.CharField(max_length=100, blank=True)
    shipped_at = models.DateTimeField(null=True, blank=True)
    delivered_at = models.DateTimeField(null=True, blank=True)


class Review(models.Model):
    product = models.ForeignKey("catalog.Product", on_delete=models.CASCADE, related_name="reviews")
    customer = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="reviews")
    rating = models.PositiveSmallIntegerField(choices=[(i, str(i)) for i in range(1, 6)])
    title = models.CharField(max_length=100)
    body = models.TextField(blank=True)
    created_at = models.DateTimeField(auto_now_add=True)

    class Meta:
        constraints = [models.UniqueConstraint(fields=["product", "customer"], name="one_review_per_customer")]
