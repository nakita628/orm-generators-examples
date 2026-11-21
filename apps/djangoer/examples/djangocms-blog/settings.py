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
    "django.contrib.sites",
    "cms",
    "menus",
    "treebeard",
    "sekizai",
    "easy_thumbnails",
    "filer",
    "djangocms_text",
    "parler",
    "taggit",
    "taggit_autosuggest",
    "meta",
    "sortedm2m",
    "djangocms_blog",
]
LANGUAGES = [("en", "English")]
PARLER_LANGUAGES = {1: ({"code": "en"},), "default": {"fallbacks": ["en"]}}
CMS_CONFIRM_VERSION4 = True
META_SITE_PROTOCOL = "https"
META_USE_SITES = True
