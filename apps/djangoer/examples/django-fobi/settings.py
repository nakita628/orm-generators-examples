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
    "fobi",
    "fobi.contrib.themes.bootstrap3",
    "fobi.contrib.plugins.form_elements.fields.text",
    "fobi.contrib.plugins.form_handlers.db_store",
    "fobi.contrib.plugins.form_handlers.mail",
    "fobi.contrib.plugins.form_handlers.http_repost",
]
FOBI_DEFAULT_THEME = "bootstrap3"
