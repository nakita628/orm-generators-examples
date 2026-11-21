"""Warehouse inventory: locations in warehouses, lots, stock levels, movements, suppliers and
purchase orders."""

from django.conf import settings
from django.db import models


class Warehouse(models.Model):
    code = models.CharField(max_length=10, unique=True)
    name = models.CharField(max_length=100)
    address = models.TextField()


class Location(models.Model):
    """A shelf or bin, inside another location or at the top of a warehouse."""

    warehouse = models.ForeignKey(Warehouse, on_delete=models.CASCADE, related_name="locations")
    parent = models.ForeignKey("self", on_delete=models.CASCADE, null=True, blank=True, related_name="children")
    code = models.CharField(max_length=30)
    is_pickable = models.BooleanField(default=True)

    class Meta:
        unique_together = [("warehouse", "code")]


class UnitOfMeasure(models.Model):
    code = models.CharField(max_length=10, primary_key=True)
    name = models.CharField(max_length=30)


class Item(models.Model):
    sku = models.CharField(max_length=40, unique=True)
    name = models.CharField(max_length=200)
    unit = models.ForeignKey(UnitOfMeasure, on_delete=models.PROTECT, to_field="code")
    barcode = models.CharField(max_length=50, unique=True, null=True, blank=True)
    tracks_lots = models.BooleanField(default=False)
    reorder_point = models.DecimalField(max_digits=12, decimal_places=3, default=0)


class Lot(models.Model):
    item = models.ForeignKey(Item, on_delete=models.CASCADE, related_name="lots")
    number = models.CharField(max_length=40)
    expires_on = models.DateField(null=True, blank=True)

    class Meta:
        unique_together = [("item", "number")]


class Stock(models.Model):
    item = models.ForeignKey(Item, on_delete=models.CASCADE, related_name="stock")
    location = models.ForeignKey(Location, on_delete=models.CASCADE, related_name="stock")
    lot = models.ForeignKey(Lot, on_delete=models.CASCADE, null=True, blank=True)
    quantity = models.DecimalField(max_digits=12, decimal_places=3, default=0)
    reserved = models.DecimalField(max_digits=12, decimal_places=3, default=0)

    class Meta:
        verbose_name_plural = "stock"
        constraints = [
            # One row per item, location and lot, and one per item and location without a lot
            models.UniqueConstraint(fields=["item", "location", "lot"], name="stock_one_row_per_lot"),
            models.UniqueConstraint(fields=["item", "location"], condition=models.Q(lot__isnull=True), name="stock_one_row_without_lot"),
            models.CheckConstraint(condition=models.Q(quantity__gte=0), name="stock_not_negative"),
            models.CheckConstraint(condition=models.Q(reserved__lte=models.F("quantity")), name="stock_reserved_within"),
        ]


class Movement(models.Model):
    class Reason(models.TextChoices):
        RECEIPT = "receipt", "Receipt"
        SHIPMENT = "shipment", "Shipment"
        TRANSFER = "transfer", "Transfer"
        ADJUSTMENT = "adjustment", "Adjustment"
        RETURN = "return", "Return"

    item = models.ForeignKey(Item, on_delete=models.PROTECT, related_name="movements")
    lot = models.ForeignKey(Lot, on_delete=models.PROTECT, null=True, blank=True)
    from_location = models.ForeignKey(Location, on_delete=models.PROTECT, null=True, blank=True, related_name="moves_out")
    to_location = models.ForeignKey(Location, on_delete=models.PROTECT, null=True, blank=True, related_name="moves_in")
    quantity = models.DecimalField(max_digits=12, decimal_places=3)
    reason = models.CharField(max_length=10, choices=Reason)
    reference = models.CharField(max_length=50, blank=True)
    by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.PROTECT)
    at = models.DateTimeField(auto_now_add=True)

    class Meta:
        constraints = [
            models.CheckConstraint(condition=models.Q(quantity__gt=0), name="movement_positive"),
            models.CheckConstraint(
                condition=models.Q(from_location__isnull=False) | models.Q(to_location__isnull=False),
                name="movement_has_a_location",
            ),
        ]


class Supplier(models.Model):
    name = models.CharField(max_length=100)
    email = models.EmailField(blank=True)
    items = models.ManyToManyField(Item, through="SupplierItem", related_name="suppliers")


class SupplierItem(models.Model):
    supplier = models.ForeignKey(Supplier, on_delete=models.CASCADE)
    item = models.ForeignKey(Item, on_delete=models.CASCADE)
    supplier_sku = models.CharField(max_length=40, blank=True)
    unit_cost = models.DecimalField(max_digits=10, decimal_places=4)
    lead_time_days = models.PositiveSmallIntegerField(default=7)


class PurchaseOrder(models.Model):
    class Status(models.TextChoices):
        DRAFT = "draft"
        SENT = "sent"
        PARTIAL = "partial", "Partly received"
        RECEIVED = "received"
        CANCELLED = "cancelled"

    supplier = models.ForeignKey(Supplier, on_delete=models.PROTECT, related_name="purchase_orders")
    warehouse = models.ForeignKey(Warehouse, on_delete=models.PROTECT)
    status = models.CharField(max_length=10, choices=Status, default=Status.DRAFT)
    ordered_on = models.DateField(null=True, blank=True)
    expected_on = models.DateField(null=True, blank=True)


class PurchaseOrderLine(models.Model):
    order = models.ForeignKey(PurchaseOrder, on_delete=models.CASCADE, related_name="lines")
    item = models.ForeignKey(Item, on_delete=models.PROTECT)
    quantity = models.DecimalField(max_digits=12, decimal_places=3)
    received = models.DecimalField(max_digits=12, decimal_places=3, default=0)
    unit_cost = models.DecimalField(max_digits=10, decimal_places=4)
