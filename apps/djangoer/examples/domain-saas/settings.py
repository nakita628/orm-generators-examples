SECRET_KEY = "example"
USE_TZ = True
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
AUTH_USER_MODEL = "accounts.User"
INSTALLED_APPS = [
    "django.contrib.contenttypes",
    "django.contrib.auth",
    "accounts",
    "billing",
]
