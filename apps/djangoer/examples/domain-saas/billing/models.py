"""A SaaS's billing: plans and their features, subscriptions, usage, invoices."""

from django.db import models


class Plan(models.Model):
    class Interval(models.TextChoices):
        MONTH = "month", "Monthly"
        YEAR = "year", "Yearly"

    code = models.SlugField(unique=True)
    name = models.CharField(max_length=50)
    price = models.DecimalField(max_digits=8, decimal_places=2)
    currency = models.CharField(max_length=3, default="USD")
    interval = models.CharField(max_length=5, choices=Interval, default=Interval.MONTH)
    seats_included = models.PositiveIntegerField(default=1)
    is_public = models.BooleanField(default=True)


class Feature(models.Model):
    code = models.SlugField(unique=True)
    description = models.CharField(max_length=200)
    plans = models.ManyToManyField(Plan, through="PlanFeature", related_name="features")


class PlanFeature(models.Model):
    plan = models.ForeignKey(Plan, on_delete=models.CASCADE)
    feature = models.ForeignKey(Feature, on_delete=models.CASCADE)
    limit = models.PositiveIntegerField(null=True, blank=True, help_text="Empty for no limit")

    class Meta:
        unique_together = [("plan", "feature")]


class Subscription(models.Model):
    class Status(models.TextChoices):
        TRIALING = "trialing", "Trialing"
        ACTIVE = "active", "Active"
        PAST_DUE = "past_due", "Past due"
        CANCELED = "canceled", "Canceled"

    organization = models.OneToOneField("accounts.Organization", on_delete=models.CASCADE, related_name="subscription")
    plan = models.ForeignKey(Plan, on_delete=models.PROTECT, related_name="subscriptions")
    status = models.CharField(max_length=10, choices=Status, default=Status.TRIALING)
    seats = models.PositiveIntegerField(default=1)
    current_period_start = models.DateTimeField()
    current_period_end = models.DateTimeField()
    cancel_at_period_end = models.BooleanField(default=False)
    provider_id = models.CharField(max_length=100, unique=True, null=True, blank=True)


class UsageRecord(models.Model):
    subscription = models.ForeignKey(Subscription, on_delete=models.CASCADE, related_name="usage")
    feature = models.ForeignKey(Feature, on_delete=models.CASCADE)
    quantity = models.PositiveBigIntegerField()
    period = models.DateField()

    class Meta:
        constraints = [models.UniqueConstraint(fields=["subscription", "feature", "period"], name="usage_once_per_period")]


class Invoice(models.Model):
    class Status(models.TextChoices):
        DRAFT = "draft", "Draft"
        OPEN = "open", "Open"
        PAID = "paid", "Paid"
        VOID = "void", "Void"

    subscription = models.ForeignKey(Subscription, on_delete=models.PROTECT, related_name="invoices")
    number = models.CharField(max_length=20, unique=True)
    status = models.CharField(max_length=5, choices=Status, default=Status.DRAFT)
    issued_on = models.DateField()
    due_on = models.DateField()
    paid_at = models.DateTimeField(null=True, blank=True)
    pdf = models.FileField(upload_to="invoices/", blank=True)

    class Meta:
        constraints = [models.CheckConstraint(condition=models.Q(due_on__gte=models.F("issued_on")), name="invoice_due_after_issue")]


class InvoiceLine(models.Model):
    invoice = models.ForeignKey(Invoice, on_delete=models.CASCADE, related_name="lines")
    description = models.CharField(max_length=200)
    quantity = models.PositiveIntegerField(default=1)
    unit_amount = models.DecimalField(max_digits=10, decimal_places=2)
    amount = models.GeneratedField(
        expression=models.F("quantity") * models.F("unit_amount"),
        output_field=models.DecimalField(max_digits=12, decimal_places=2),
        db_persist=True,
    )
