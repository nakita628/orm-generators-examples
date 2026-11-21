SECRET_KEY = "example"
USE_TZ = True
SITE_ID = 1
STATIC_URL = "/static/"
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
TEMPLATES = [{"BACKEND": "django.template.backends.django.DjangoTemplates", "APP_DIRS": True, "DIRS": [],
    "OPTIONS": {"context_processors": ["django.template.context_processors.request",
    "django.contrib.auth.context_processors.auth", "django.contrib.messages.context_processors.messages"]}}]
INSTALLED_APPS = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
    "openedx_content",
    "openedx_content.backcompat.backup_restore",
    "openedx_content.backcompat.collections",
    "openedx_content.backcompat.components",
    "openedx_content.backcompat.contents",
    "openedx_content.backcompat.publishing",
    "openedx_content.backcompat.sections",
    "openedx_content.backcompat.subsections",
    "openedx_content.backcompat.units",
    "openedx_tagging",
    "rules.apps.AutodiscoverRulesConfig",
]

