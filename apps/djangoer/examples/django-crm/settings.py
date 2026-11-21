from crm.settings import *  # noqa: F403

SECRET_KEY = "example"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
