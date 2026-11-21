"""An online shop's catalog: a category tree, products with variants and options, prices and
stock."""

from django.db import models
from django.db.models import Q


class Category(models.Model):
    parent = models.ForeignKey("self", on_delete=models.CASCADE, null=True, blank=True, related_name="children")
    name = models.CharField(max_length=100)
    slug = models.SlugField(max_length=100)
    position = models.PositiveSmallIntegerField(default=0)

    class Meta:
        verbose_name_plural = "categories"
        ordering = ["position", "name"]
        constraints = [models.UniqueConstraint(fields=["parent", "slug"], name="category_slug_unique_in_parent")]


class Brand(models.Model):
    name = models.CharField(max_length=100, unique=True)
    logo = models.ImageField(upload_to="brands/", blank=True)


class Product(models.Model):
    class Status(models.TextChoices):
        DRAFT = "draft", "Draft"
        ACTIVE = "active", "Active"
        ARCHIVED = "archived", "Archived"

    name = models.CharField(max_length=200)
    slug = models.SlugField(max_length=200, unique=True)
    description = models.TextField(blank=True)
    brand = models.ForeignKey(Brand, on_delete=models.SET_NULL, null=True, blank=True, related_name="products")
    categories = models.ManyToManyField(Category, related_name="products")
    status = models.CharField(max_length=10, choices=Status, default=Status.DRAFT, db_index=True)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)


class Option(models.Model):
    """A way variants differ: size, colour."""

    name = models.CharField(max_length=50, unique=True)


class OptionValue(models.Model):
    option = models.ForeignKey(Option, on_delete=models.CASCADE, related_name="values")
    value = models.CharField(max_length=50)

    class Meta:
        unique_together = [("option", "value")]


class Variant(models.Model):
    product = models.ForeignKey(Product, on_delete=models.CASCADE, related_name="variants")
    sku = models.CharField("SKU", max_length=40, unique=True)
    values = models.ManyToManyField(OptionValue, blank=True)
    price = models.DecimalField(max_digits=10, decimal_places=2)
    compare_at_price = models.DecimalField(max_digits=10, decimal_places=2, null=True, blank=True)
    weight_grams = models.PositiveIntegerField(null=True, blank=True)
    stock = models.IntegerField(default=0)

    class Meta:
        constraints = [
            models.CheckConstraint(condition=Q(price__gte=0), name="variant_price_not_negative"),
            models.CheckConstraint(
                condition=Q(compare_at_price__isnull=True) | Q(compare_at_price__gt=models.F("price")),
                name="variant_compare_at_above_price",
            ),
        ]


class ProductImage(models.Model):
    product = models.ForeignKey(Product, on_delete=models.CASCADE, related_name="images")
    variant = models.ForeignKey(Variant, on_delete=models.SET_NULL, null=True, blank=True, related_name="images")
    image = models.ImageField(upload_to="products/")
    alt_text = models.CharField(max_length=200, blank=True)

    class Meta:
        order_with_respect_to = "product"
