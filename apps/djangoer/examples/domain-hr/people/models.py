"""People and HR: departments, positions, employees with managers, employment and salary
history, leave and reviews."""

from django.conf import settings
from django.db import models


class Department(models.Model):
    name = models.CharField(max_length=100, unique=True)
    parent = models.ForeignKey("self", on_delete=models.PROTECT, null=True, blank=True, related_name="sub_departments")
    head = models.OneToOneField("Employee", on_delete=models.SET_NULL, null=True, blank=True, related_name="heads")
    cost_center = models.CharField(max_length=20, blank=True)


class Position(models.Model):
    title = models.CharField(max_length=100)
    department = models.ForeignKey(Department, on_delete=models.PROTECT, related_name="positions")
    grade = models.PositiveSmallIntegerField()
    is_manager = models.BooleanField(default=False)


class Employee(models.Model):
    user = models.OneToOneField(settings.AUTH_USER_MODEL, on_delete=models.SET_NULL, null=True, blank=True, related_name="employee")
    employee_number = models.CharField(max_length=20, unique=True)
    first_name = models.CharField(max_length=50)
    last_name = models.CharField(max_length=50)
    birth_date = models.DateField(null=True, blank=True)
    work_email = models.EmailField(unique=True)
    manager = models.ForeignKey("self", on_delete=models.SET_NULL, null=True, blank=True, related_name="reports")
    position = models.ForeignKey(Position, on_delete=models.PROTECT, related_name="employees")
    hired_on = models.DateField()
    left_on = models.DateField(null=True, blank=True)
    skills = models.ManyToManyField("Skill", through="EmployeeSkill", related_name="employees", blank=True)

    class Meta:
        ordering = ["last_name", "first_name"]
        constraints = [
            models.CheckConstraint(condition=~models.Q(manager=models.F("id")), name="employee_not_own_manager"),
            models.CheckConstraint(
                condition=models.Q(left_on__isnull=True) | models.Q(left_on__gte=models.F("hired_on")),
                name="employee_left_after_hire",
            ),
        ]


class Contract(models.Model):
    class Kind(models.TextChoices):
        PERMANENT = "permanent"
        FIXED_TERM = "fixed", "Fixed term"
        CONTRACTOR = "contractor"
        INTERN = "intern"

    employee = models.ForeignKey(Employee, on_delete=models.CASCADE, related_name="contracts")
    kind = models.CharField(max_length=10, choices=Kind)
    starts_on = models.DateField()
    ends_on = models.DateField(null=True, blank=True)
    hours_per_week = models.DecimalField(max_digits=4, decimal_places=1, default=40)
    document = models.FileField(upload_to="contracts/", blank=True)


class Salary(models.Model):
    employee = models.ForeignKey(Employee, on_delete=models.CASCADE, related_name="salaries")
    amount = models.DecimalField(max_digits=12, decimal_places=2)
    currency = models.CharField(max_length=3, default="EUR")
    effective_from = models.DateField()
    reason = models.CharField(max_length=200, blank=True)

    class Meta:
        verbose_name_plural = "salaries"
        get_latest_by = "effective_from"
        unique_together = [("employee", "effective_from")]


class Skill(models.Model):
    name = models.CharField(max_length=50, unique=True)


class EmployeeSkill(models.Model):
    employee = models.ForeignKey(Employee, on_delete=models.CASCADE)
    skill = models.ForeignKey(Skill, on_delete=models.CASCADE)
    level = models.PositiveSmallIntegerField(choices=[(1, "Basic"), (2, "Working"), (3, "Expert")])


class LeaveType(models.Model):
    name = models.CharField(max_length=50, unique=True)
    days_per_year = models.DecimalField(max_digits=4, decimal_places=1)
    is_paid = models.BooleanField(default=True)


class LeaveRequest(models.Model):
    class Status(models.TextChoices):
        PENDING = "pending"
        APPROVED = "approved"
        REJECTED = "rejected"
        CANCELLED = "cancelled"

    employee = models.ForeignKey(Employee, on_delete=models.CASCADE, related_name="leave_requests")
    leave_type = models.ForeignKey(LeaveType, on_delete=models.PROTECT)
    starts_on = models.DateField()
    ends_on = models.DateField()
    half_day = models.BooleanField(default=False)
    status = models.CharField(max_length=10, choices=Status, default=Status.PENDING)
    approver = models.ForeignKey(Employee, on_delete=models.SET_NULL, null=True, blank=True, related_name="leave_to_approve")
    note = models.TextField(blank=True)

    class Meta:
        constraints = [models.CheckConstraint(condition=models.Q(ends_on__gte=models.F("starts_on")), name="leave_ends_after_start")]


class Review(models.Model):
    employee = models.ForeignKey(Employee, on_delete=models.CASCADE, related_name="reviews")
    reviewer = models.ForeignKey(Employee, on_delete=models.SET_NULL, null=True, related_name="reviews_given")
    period = models.CharField(max_length=7, help_text="e.g. 2026-H1")
    rating = models.PositiveSmallIntegerField(choices=[(i, str(i)) for i in range(1, 6)])
    summary = models.TextField()
    goals = models.JSONField(default=list, blank=True)

    class Meta:
        unique_together = [("employee", "period")]
