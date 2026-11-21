from allianceauth.project_template.project_name.settings.base import *  # noqa: F403

SECRET_KEY = "example"
DATABASES = {"default": {"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}
# Alliance Auth talks to Redis while its apps load: an in-memory one stands in for it.
import fakeredis  # noqa: E402

CACHES = {
    "default": {
        "BACKEND": "django_redis.cache.RedisCache",
        "LOCATION": "redis://localhost:6379/0",
        "OPTIONS": {"CONNECTION_POOL_KWARGS": {"connection_class": fakeredis.FakeConnection}},
    }
}
