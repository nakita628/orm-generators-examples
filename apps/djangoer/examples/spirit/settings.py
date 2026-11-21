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
    "django.contrib.humanize",
    "spirit.core",
    "spirit.admin",
    "spirit.search",
    "spirit.user",
    "spirit.user.admin",
    "spirit.user.auth",
    "spirit.category",
    "spirit.category.admin",
    "spirit.topic",
    "spirit.topic.admin",
    "spirit.topic.favorite",
    "spirit.topic.moderate",
    "spirit.topic.notification",
    "spirit.topic.private",
    "spirit.topic.unread",
    "spirit.comment",
    "spirit.comment.bookmark",
    "spirit.comment.flag",
    "spirit.comment.flag.admin",
    "spirit.comment.history",
    "spirit.comment.like",
    "spirit.comment.poll",
    "djconfig",
    "haystack",
]
HAYSTACK_CONNECTIONS = {"default": {"ENGINE": "haystack.backends.simple_backend.SimpleEngine"}}
ST_SITE_URL = "http://127.0.0.1:8000/"
MIDDLEWARE = ["django.contrib.sessions.middleware.SessionMiddleware", "django.contrib.auth.middleware.AuthenticationMiddleware", "djconfig.middleware.DjConfigMiddleware"]
