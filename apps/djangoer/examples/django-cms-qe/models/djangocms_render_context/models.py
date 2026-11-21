import logging
from typing import Any, Optional, Union

from cms.models.pluginmodel import CMSPlugin
from django.db import models
from django.utils.translation import gettext_lazy as _
from filer.fields.file import FilerFileField

from .encoders import PrettyJsonEncoder
from .exceptions import SourceParseFailure
from .loaders import LOADERS, get_cached_data_from_url, get_data_from_file, value_to_bytes
from .utils import get_context_from_path

contentType = Optional[Union[dict[Any, Any], list[Any]]]  # noqa: UP045, UP007

LOGGER = logging.getLogger(__name__)


class RenderContext(CMSPlugin):
    mimetype = models.CharField(
        verbose_name=_("Type of Data for context"),
        max_length=255,
        default="application/json",
        help_text=_("The format must be consistent with the data in the context."),
    )
    context = models.TextField(
        verbose_name=_("Data for context"),
        null=True,
        blank=True,
        help_text=_("Context data in selected format. They take precedence over the source file and source URL."),
    )
    config = models.JSONField(
        verbose_name=_("Configuration"),
        null=True,
        blank=True,
        encoder=PrettyJsonEncoder,
        help_text=_(
            "Additional configuration for spreadsheet with source data. "
            "The configuration is automatically added or updated when the spreadsheet is saved. "
            'To edit the configuration "manually", you must have the Data for the context set to Source.'
        ),
    )
    file = FilerFileField(
        verbose_name=_("Source file"),
        null=True,
        blank=True,
        on_delete=models.SET_NULL,
        help_text=_(
            "Context data file. If Data is specified for the context, the source file is not used even "
            "though it is specified."
        ),
    )
    source = models.URLField(
        verbose_name=_("Source URL"),
        null=True,
        blank=True,
        help_text=_(
            "The URL of the source from which the data will be downloaded. If Data for Context or Source File"
            " is specified, the source URL is not used even though it is specified."
        ),
    )
    cached = models.PositiveSmallIntegerField(
        verbose_name=_("Download period from source URL"),
        default=5,
        help_text=_(
            "The time in minutes during which data will not be retrieved from the Source URL, but will remain "
            "in the cache. If set to zero, data from the source URL is not cached, but is loaded every time."
        ),
    )

    template = models.TextField(
        verbose_name=_("Template"),
        null=True,
        blank=True,
        help_text=_(
            "Django template for the context. It takes precedence over the template selected from the list."
            ' The context is available through the "data" value.'
        ),
    )
    template_list = models.CharField(
        verbose_name=_("Template list"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_("List of templates. If Template is specified, the template set in the list will not be used."),
    )
    path = models.CharField(
        verbose_name=_("Data path"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_(
            'The path to the data in the form "name.name.name".'
            ' The "data" context in the template starts from this key.'
        ),
    )
    processor = models.CharField(
        verbose_name=_("Processor"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_("List of functions that process the context."),
    )
    detail_extends = models.CharField(
        verbose_name=_("Detail extends"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_("The detail uses a template. If not specified, the default template is used."),
    )
    detail_template = models.TextField(
        verbose_name=_("Detail template"),
        null=True,
        blank=True,
        help_text=_(
            "Django template for the context. It takes precedence over the template selected from the list."
            ' The context is available through the "detail" value. The entire context is still available via the'
            ' "data" key.'
        ),
    )
    detail_template_list = models.CharField(
        verbose_name=_("Detail template list"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_("List of templates. If Template is specified, the template set in the list will not be used."),
    )
    detail_processor = models.CharField(
        verbose_name=_("Detail processor"),
        null=True,
        blank=True,
        max_length=255,
        help_text=_("List of functions that process the context."),
    )

    def __str__(self):
        text = []
        if self.context:
            text.append(self._meta.get_field("context").verbose_name)
        elif self.file:
            text.append(self._meta.get_field("file").verbose_name)
        elif self.source:
            text.append(self._meta.get_field("source").verbose_name)
        else:
            text.append(_("Data not entered."))
        if self.template:
            text.append(self._meta.get_field("template").verbose_name)
        elif self.template_list:
            text.append(self._meta.get_field("template_list").verbose_name)
        else:
            text.append(_("No template."))
        return " + ".join([str(t) for t in text])

    def get_data_by_path(self, data: contentType) -> contentType:
        if self.path:
            try:
                data = get_context_from_path(data, self.path)
            except (IndexError, KeyError, ValueError, TypeError) as err:
                LOGGER.error(err)
                return {}
        return data

    def get_data(self) -> contentType:
        if self.context:
            return self.get_data_by_path(LOADERS[self.mimetype](value_to_bytes(self.context, self.mimetype)))
        elif self.file:
            try:
                return self.get_data_by_path(get_data_from_file(self.file))
            except SourceParseFailure as err:
                LOGGER.error(err)
        elif self.source:
            try:
                return self.get_data_by_path(get_cached_data_from_url(self.pk, self.source, self.cached))
            except SourceParseFailure as err:
                LOGGER.error(err)
        return None


class Default(CMSPlugin):
    code = models.CharField(
        verbose_name=_("Code"),
        default="D",
        max_length=255,
        help_text=_("The value name for the default type."),
    )

    def __str__(self):
        return self.code


class Stack(CMSPlugin):
    selector = models.CharField(
        verbose_name=_("The path to data"),
        max_length=255,
        help_text=_("Keys or indexes separated by a dot."),
    )

    def __str__(self):
        return self.selector
