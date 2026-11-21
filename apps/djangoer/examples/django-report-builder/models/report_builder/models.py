import datetime
import re
import time
from decimal import Decimal

from asgiref.sync import sync_to_async
from dateutil import parser
from django import forms
from django.conf import settings
from django.contrib.contenttypes.models import ContentType
from django.core.exceptions import FieldDoesNotExist, ObjectDoesNotExist, ValidationError
from django.core.files.base import ContentFile
from django.db import models
from django.db.models import Avg, Count, F, Max, Min, Sum
from django.urls import reverse
from django.utils.functional import cached_property
from django.utils.safestring import mark_safe

from report_builder.unique_slugify import unique_slugify

from .email import email_report
from .mixins import DataExportMixin, generate_filename
from .utils import get_model_from_path_string

AUTH_USER_MODEL = getattr(settings, "AUTH_USER_MODEL", "auth.User")


def get_allowed_models():
    models = ContentType.objects.all()
    if getattr(settings, "REPORT_BUILDER_INCLUDE", False):
        all_model_names = []
        additional_models = []
        for element in settings.REPORT_BUILDER_INCLUDE:
            split_element = element.split(".")
            if len(split_element) == 2:
                additional_models.append(
                    models.filter(app_label=split_element[0], model=split_element[1])
                )
            else:
                all_model_names.append(element)
        models = models.filter(model__in=all_model_names)
        for additional_model in additional_models:
            models = models | additional_model
    if getattr(settings, "REPORT_BUILDER_EXCLUDE", False):
        all_model_names = []
        for element in settings.REPORT_BUILDER_EXCLUDE:
            split_element = element.split(".")
            if len(split_element) == 2:
                models = models.exclude(app_label=split_element[0], model=split_element[1])
            else:
                all_model_names.append(element)
        models = models.exclude(model__in=all_model_names)
    return models


def get_limit_choices_to_callable():
    return {"pk__in": get_allowed_models()}


class Report(models.Model):
    """A saved report with queryset and descriptive fields"""

    def _get_model_manager(self):
        """Get  manager from settings else use objects"""
        model_manager = "objects"
        if getattr(settings, "REPORT_BUILDER_MODEL_MANAGER", False):
            model_manager = settings.REPORT_BUILDER_MODEL_MANAGER
        return model_manager

    @staticmethod
    def allowed_models():
        return get_allowed_models()

    name = models.CharField(max_length=255)
    slug = models.SlugField(verbose_name="Short Name")
    description = models.TextField(blank=True)
    root_model = models.ForeignKey(
        ContentType, limit_choices_to=get_limit_choices_to_callable, on_delete=models.CASCADE
    )
    created = models.DateField(auto_now_add=True)
    modified = models.DateField(auto_now=True)
    user_created = models.ForeignKey(
        AUTH_USER_MODEL, editable=False, blank=True, null=True, on_delete=models.SET_NULL
    )
    user_modified = models.ForeignKey(
        AUTH_USER_MODEL,
        editable=False,
        blank=True,
        null=True,
        related_name="report_modified_set",
        on_delete=models.SET_NULL,
    )
    distinct = models.BooleanField(default=False)
    row_limit = models.PositiveIntegerField(
        null=True,
        blank=True,
        help_text="Maximum number of rows to include in exports. Leave blank for no limit.",
    )
    report_file = models.FileField(upload_to="report_files", blank=True)
    report_file_creation = models.DateTimeField(blank=True, null=True)
    # Async-export correlation: ID of the task that wrote (or is writing)
    # ``report_file``. ``check_status`` only returns the URL when the
    # polled task ID matches this — keeps stale files from leaking through
    # if a later export was already started.
    report_file_task_id = models.CharField(max_length=64, blank=True, default="")
    starred = models.ManyToManyField(
        AUTH_USER_MODEL,
        blank=True,
        help_text="These users have starred this report for easy reference.",
        related_name="report_starred_set",
    )

    def __str__(self):
        return self.name

    def save(self, *args, **kwargs):
        if not self.id:
            unique_slugify(self, self.name)
        super().save(*args, **kwargs)

    async def add_aggregates(self, queryset, display_fields=None):
        agg_funcs = {
            "Avg": Avg,
            "Min": Min,
            "Max": Max,
            "Count": Count,
            "Sum": Sum,
        }
        if display_fields is None:
            display_fields = [
                df async for df in self.displayfield_set.filter(aggregate__isnull=False)
            ]
        for display_field in display_fields:
            if display_field.aggregate:
                func = agg_funcs[display_field.aggregate]
                full_name = display_field.path + display_field.field
                queryset = queryset.annotate(func(full_name))
        return queryset

    @property
    def root_model_class(self):
        return self.root_model.model_class()

    def get_field_type(self, field_name, path=""):
        """Get field type for given field name.
        field_name is the full path of the field
        path is optional
        """
        model = get_model_from_path_string(self.root_model_class, path + field_name)

        # Is it a ORM field?
        try:
            return model._meta.get_field(field_name).get_internal_type()
        except FieldDoesNotExist:
            pass
        # Is it a property?
        field_attr = getattr(model, field_name, None)
        if isinstance(field_attr, (property, cached_property)):
            return "Property"
        # Is it a custom field?
        try:
            model().get_custom_field(field_name)
            return "Custom Field"
        except (ObjectDoesNotExist, AttributeError):
            pass
        return "Invalid"

    async def get_good_display_fields(self):
        """Return a queryset of display fields whose field_type isn't Invalid.

        ``select_related('display_format')`` so that downstream iteration
        accessing ``df.display_format`` doesn't fire a sync DB query.
        """
        display_fields = self.displayfield_set.select_related("display_format")
        bad_ids = [df.id async for df in display_fields if df.field_type == "Invalid"]
        return display_fields.exclude(id__in=bad_ids)

    async def get_query(self):
        model_class = self.root_model_class

        if getattr(model_class, "report_builder_model_manager", False):
            objects = model_class.report_builder_model_manager.all()
        else:
            manager = self._get_model_manager()
            objects = getattr(model_class, manager).all()

        # Materialize once; we iterate it twice below.
        filter_fields = [ff async for ff in self.filterfield_set.order_by("position")]

        filters = {}
        excludes = {}
        for filter_field in filter_fields:
            if filter_field.field_type in ("Property", "Custom Field"):
                continue
            if filter_field.filter_type in ("max", "min"):
                # Annotation-filters are handled below.
                continue

            filter_string = str(filter_field.path + filter_field.field)

            if filter_field.filter_type:
                ft = filter_field.filter_type
                fs = ft if ft != "relative_range" else "range"
                filter_string += "__" + fs

            if filter_field.filter_type == "isnull" and filter_field.filter_value in ["0", "False"]:
                filter_ = {filter_string: False}
            elif filter_field.filter_type == "in":
                filter_ = {filter_string: filter_field.filter_value.split(",")}
            else:
                filter_value = filter_field.filter_value
                if filter_field.filter_type == "range":
                    filter_value = sorted([filter_value, filter_field.filter_value2])
                elif filter_field.filter_type == "relative_range":
                    filter_value = filter_field.get_relative_range()
                filter_ = {filter_string: filter_value}

            if not filter_field.exclude:
                filters.update(filter_)
            else:
                excludes.update(filter_)

        if filters:
            objects = objects.filter(**filters)
        if excludes:
            objects = objects.exclude(**excludes)

        for filter_field in filter_fields:
            if filter_field.filter_type in ("max", "min"):
                func = {"max": Max, "min": Min}[filter_field.filter_type]
                column_name = f"{filter_field.path}{filter_field.field}__{filter_field.field_type}"
                filter_string = filter_field.path + filter_field.field
                annotate_args = {column_name: func(filter_string)}
                filter_args = {column_name: F(filter_field.field)}
                objects = objects.annotate(**annotate_args).filter(**filter_args)

        objects = await self.add_aggregates(objects)

        if self.distinct:
            objects = objects.distinct()

        return objects

    def get_absolute_url(self):
        return reverse("report_update_view", args=(self.id,))

    def edit(self):
        return mark_safe(
            '<a href="{}"><img style="width: 26px; margin: -6px" src="{}report_builder/img/edit.svg"/></a>'.format(
                self.get_absolute_url(), getattr(settings, "STATIC_URL", "/static/")
            )
        )

    def download_xlsx(self):
        if getattr(settings, "REPORT_BUILDER_ASYNC_REPORT", False):
            return mark_safe(
                '<a href="javascript:void(0)" onclick="get_async_report({})"><img style="width: 26px; margin: -6px" src="{}report_builder/img/download.svg"/></a>'.format(
                    self.id,
                    getattr(settings, "STATIC_URL", "/static/"),
                )
            )
        else:
            return mark_safe(
                '<a href="{}"><img style="width: 26px; margin: -6px" src="{}report_builder/img/download.svg"/></a>'.format(
                    reverse("report_download_file", args=[self.id]),
                    getattr(settings, "STATIC_URL", "/static/"),
                )
            )

    download_xlsx.short_description = "Download"

    def copy_report(self):
        return mark_safe(
            '<a href="{}"><img style="width: 26px; margin: -6px" src="{}report_builder/img/copy.svg"/></a>'.format(
                reverse("report_builder_create_copy", args=[self.id]),
                getattr(settings, "STATIC_URL", "/static/"),
            )
        )

    copy_report.short_description = "Copy"

    def check_report_display_field_positions(self):
        """After report is saved, make sure positions are sane"""
        for i, display_field in enumerate(self.displayfield_set.all()):
            if display_field.position != i + 1:
                display_field.position = i + 1
                display_field.save()

    async def email_report(self, user=None, email=None):
        report_url = self.report_file.url
        await email_report(report_url, user=user, email=email)

    async def async_report_save(
        self,
        objects_list,
        title,
        header,
        widths,
        user=None,
        file_type="xlsx",
        email_to: list[str] | None = None,
    ):
        data_export = DataExportMixin()
        if file_type == "csv":
            csv_file = data_export.list_to_csv_file(objects_list, title, header, widths)
            title = generate_filename(title, ".csv")
            await sync_to_async(self.report_file.save)(
                title,
                ContentFile(csv_file.getvalue().encode()),
            )
        else:
            xlsx_file = data_export.list_to_xlsx_file(objects_list, title, header, widths)
            title = generate_filename(title, ".xlsx")
            await sync_to_async(self.report_file.save)(
                title,
                ContentFile(xlsx_file.getvalue()),
            )
        self.report_file_creation = datetime.datetime.today()
        await self.asave()
        if email_to:
            for addr in email_to:
                await self.email_report(email=addr)
        elif getattr(settings, "REPORT_BUILDER_EMAIL_NOTIFICATION", False):
            if user and user.email:
                await self.email_report(user=user)

    async def run_report(
        self,
        file_type,
        user=None,
        queryset=None,
        asynchronous=False,
        scheduled=False,
        email_to: list[str] | None = None,
    ):
        """Generate this report's output (response or saved file)."""
        if not queryset:
            queryset = await self.get_query()

        display_fields_qs = await self.get_good_display_fields()
        # Materialize once to read header/width here, but pass the
        # queryset to the (still-sync) mixin so it can use .filter()
        # for sort handling.
        display_fields = [df async for df in display_fields_qs]

        data_export = DataExportMixin()
        objects_list, _message = await data_export.report_to_list(
            queryset,
            display_fields_qs,
            user,
            preview=False,
        )
        if self.row_limit:
            objects_list = objects_list[: self.row_limit]
        title = re.sub(r"\W+", "", self.name)[:30]
        header = [field.name for field in display_fields]
        widths = [field.width for field in display_fields]

        if scheduled:
            await self.async_report_save(
                objects_list, title, header, widths, file_type=file_type, email_to=email_to
            )
        elif asynchronous:
            if user is None:
                raise Exception("Cannot run async report without a user")
            await self.async_report_save(
                objects_list, title, header, widths, user=user, file_type=file_type
            )
        else:
            if file_type == "csv":
                return data_export.list_to_csv_response(objects_list, title, header, widths)
            else:
                return data_export.list_to_xlsx_response(objects_list, title, header, widths)


class Format(models.Model):
    """A specifies a Python string format for e.g. `DisplayField`s."""

    name = models.CharField(max_length=50, blank=True, default="")
    string = models.CharField(
        max_length=300,
        blank=True,
        default="",
        help_text="Python string format. Ex ${} would place a $ in front of the result.",
    )

    def __str__(self):
        return self.name


class AbstractField(models.Model):
    report = models.ForeignKey(Report, on_delete=models.CASCADE)
    path = models.CharField(max_length=2000, blank=True)
    path_verbose = models.CharField(max_length=2000, blank=True)
    field = models.CharField(max_length=2000)
    field_verbose = models.CharField(max_length=2000)
    position = models.PositiveSmallIntegerField(blank=True, null=True)

    class Meta:
        abstract = True
        ordering = ["position"]

    @property
    def field_type(self):
        return self.report.get_field_type(self.field, self.path)

    @property
    def field_key(self):
        """This key can be passed to a Django ORM values_list"""
        return self.path + self.field

    @property
    def choices(self):
        if self.pk:
            model = get_model_from_path_string(self.report.root_model.model_class(), self.path)
            return self.get_choices(model, self.field)


class DisplayField(AbstractField):
    """A display field to show in a report. Always belongs to a Report"""

    name = models.CharField(max_length=2000)
    sort = models.IntegerField(blank=True, null=True)
    sort_reverse = models.BooleanField(verbose_name="Reverse", default=False)
    width = models.IntegerField(default=15)
    aggregate = models.CharField(
        max_length=5,
        choices=(
            ("Sum", "Sum"),
            ("Count", "Count"),
            ("Avg", "Avg"),
            ("Max", "Max"),
            ("Min", "Min"),
        ),
        blank=True,
    )
    total = models.BooleanField(default=False)
    group = models.BooleanField(default=False)
    display_format = models.ForeignKey(Format, blank=True, null=True, on_delete=models.SET_NULL)

    def get_choices(self, model, field_name):
        try:
            model_field = model._meta.get_field(field_name)
        except FieldDoesNotExist:
            model_field = None
        if model_field and model_field.choices:
            return ((model_field.get_prep_value(key), val) for key, val in model_field.choices)

    @property
    def choices_dict(self):
        choices = self.choices
        choices_dict = {}
        if choices:
            for choice in choices:
                choices_dict.update({choice[0]: choice[1]})
        return choices_dict

    def __str__(self):
        return self.name


class FilterField(AbstractField):
    """
    A filter model used to filter DisplayFields using Django ORM
    filter arguments or custom filter values.
    """

    filter_type = models.CharField(
        max_length=20,
        choices=(
            ("exact", "Equals"),
            ("iexact", "Equals (case-insensitive)"),
            ("contains", "Contains"),
            ("icontains", "Contains (case-insensitive)"),
            ("in", "in (comma seperated 1,2,3)"),
            ("gt", "Greater than"),
            ("gte", "Greater than equals"),
            ("lt", "Less than"),
            ("lte", "Less than equals"),
            ("startswith", "Starts with"),
            ("istartswith", "Starts with (case-insensitive)"),
            ("endswith", "Ends with"),
            ("iendswith", "Ends with  (case-insensitive)"),
            ("range", "range"),
            ("relative_range", "relative_range"),
            ("week_day", "Week day"),
            ("isnull", "Is null"),
            ("regex", "Regular Expression"),
            ("iregex", "Reg. Exp. (case-insensitive)"),
            ("max", "Max (annotation-filter)"),
            ("min", "Min (annotation-filter)"),
        ),
        blank=True,
        default="icontains",
    )
    filter_delta = models.BigIntegerField(null=True, blank=True)
    filter_value = models.CharField(max_length=2000, blank=True)
    filter_value2 = models.CharField(max_length=2000, blank=True)
    exclude = models.BooleanField(default=False)

    def clean(self):
        dt_types = ["DateField", "DateTimeField", "TimeField"]

        if self.filter_type == "range" and self.filter_value2 in [None, ""]:
            raise ValidationError("Range filters must have two values")

        # Raise error if 'relative_range' filter is applied to a non-supported
        # field type
        if self.filter_type == "relative_range" and self.field_type not in dt_types:
            raise ValidationError(
                "Relative Range filtering is only currently supported for"
                f" the following field types: {dt_types}."
            )

        # Check for required relative range filter_delta
        if self.filter_type == "relative_range" and self.filter_delta is None:
            raise ValidationError("Relative Range filters must have value and delta inputs.")

        if self.filter_type in ("max", "min"):
            # These filter types ignore their value.
            pass

        # clean DateTimeField inputs
        elif self.field_type == "DateField" and self.filter_type != "isnull":
            self.filter_value = str(self.parse_datetime_fields(self.filter_value))

        return super().clean()

    def parse_datetime_fields(self, dt_type):
        """Clean and parse datetime filter_value inputs."""
        date_value = parser.parse(dt_type)
        date_form = forms.DateTimeField()

        if self.field_type == "DateField":
            date_value = date_value.date()
            date_form = forms.DateField()
        if self.field_type == "TimeField":
            date_value = date_value.time()
            date_form = forms.TimeField()

        return date_form.clean(date_value)

    def get_choices(self, model, field_name):
        try:
            model_field = model._meta.get_field(field_name)
        except FieldDoesNotExist:
            model_field = None
        if model_field and model_field.choices:
            return model_field.choices

    def filter_property(self, value):
        """Determine if passed value should be filtered or not"""
        filter_field = self
        filter_type = filter_field.filter_type
        filter_value = filter_field.filter_value
        filtered = True
        WEEKDAY_INTS = {
            "monday": 0,
            "tuesday": 1,
            "wednesday": 2,
            "thursday": 3,
            "friday": 4,
            "saturday": 5,
            "sunday": 6,
        }
        if filter_type == "exact" and str(value) == filter_value:
            filtered = False
        if filter_type == "iexact" and str(value).lower() == str(filter_value).lower():
            filtered = False
        if filter_type == "contains" and filter_value in value:
            filtered = False
        if filter_type == "icontains" and str(filter_value).lower() in str(value).lower():
            filtered = False
        if filter_type == "in" and value in filter_value:
            filtered = False
        # convert dates and datetimes to timestamps in order to compare digits and date/times the same
        if isinstance(value, datetime.datetime) or isinstance(value, datetime.date):
            value = str(time.mktime(value.timetuple()))
            try:
                filter_value_dt = parser.parse(filter_value)
                filter_value = str(time.mktime(filter_value_dt.timetuple()))
            except ValueError:
                pass
        if filter_type == "gt" and Decimal(value) > Decimal(filter_value):
            filtered = False
        if filter_type == "gte" and Decimal(value) >= Decimal(filter_value):
            filtered = False
        if filter_type == "lt" and Decimal(value) < Decimal(filter_value):
            filtered = False
        if filter_type == "lte" and Decimal(value) <= Decimal(filter_value):
            filtered = False
        if filter_type == "startswith" and str(value).startswith(str(filter_value)):
            filtered = False
        if filter_type == "istartswith" and str(value).lower().startswith(str(filter_value)):
            filtered = False
        if filter_type == "endswith" and str(value).endswith(str(filter_value)):
            filtered = False
        if filter_type == "iendswith" and str(value).lower().endswith(str(filter_value)):
            filtered = False
        if filter_type == "range" and value in [int(x) for x in str(filter_value).split(",")]:
            filtered = False
        if (
            filter_type == "week_day"
            and WEEKDAY_INTS.get(str(filter_value).lower()) == value.weekday()
        ):
            filtered = False
        if filter_type == "isnull" and value is None:
            filtered = False
        if filter_type == "regex" and re.search(filter_value, value):
            filtered = False
        if filter_type == "iregex" and re.search(filter_value, value, re.I):
            filtered = False

        if filter_field.exclude:
            return not filtered
        return filtered

    def get_relative_range(self):
        """
        Generate a 'filterable' range from the current date and filter delta.
        Ex.
            With:
                self.filter_type = 'relative_range'
                self.filter_delta = -60 * 60 * 24 * 2 (i.e. -2 days)
            Return:
                # a 'negative' two day range from filter_value
                ["2017-01-01", "2017-01-03"]
        """
        day = 60 * 60 * 24

        if self.field_type == "DateField":
            if abs(self.filter_delta) < day:
                raise ValidationError("DateField delta must be at least 1 day.")
            first = datetime.date.today()
            second = first + datetime.timedelta(seconds=self.filter_delta)
            output_range = sorted([first, second])
            output = [date.strftime("%Y-%m-%d") for date in output_range]

        elif self.field_type == "DateTimeField":
            first = datetime.datetime.now()
            second = first + datetime.timedelta(seconds=self.filter_delta)
            output_range = sorted([first, second])
            output = [date.strftime("%Y-%m-%d %H:%M:%S") for date in output_range]

        elif self.field_type == "TimeField":
            day_const = datetime.datetime(1, 1, 1)
            first = datetime.datetime.now().time()
            delta = datetime.timedelta(seconds=self.filter_delta)
            diff = datetime.datetime.combine(day_const, first) + delta
            second = diff.time()
            output_range = sorted([first, second])
            output = [date.strftime("%H:%M") for date in output_range]

        return output

    @property
    def field_type(self):
        return self.report.get_field_type(self.field, self.path)

    @property
    def choices(self):
        if self.pk:
            model = get_model_from_path_string(self.report.root_model.model_class(), self.path)
            return self.get_choices(model, self.field)

    def __str__(self):
        return self.field
