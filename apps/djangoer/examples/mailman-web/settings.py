from mailman_web.settings.base import *  # noqa: F403
from mailman_web.settings.mailman import *  # noqa: F403

SECRET_KEY = "example"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
LOGGING = {"version": 1, "disable_existing_loggers": False}
