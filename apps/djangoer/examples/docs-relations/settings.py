SECRET_KEY = "example"
USE_TZ = True
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
INSTALLED_APPS = [
    "django.contrib.contenttypes",
    "django.contrib.auth",
    "many_to_one",
    "many_to_many",
    "one_to_one",
]
