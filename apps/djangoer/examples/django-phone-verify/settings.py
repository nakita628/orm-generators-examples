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
    "rest_framework",
    "phone_verify",
]
PHONE_VERIFICATION = {"BACKEND": "phone_verify.backends.twilio.TwilioBackend", "OPTIONS": {"SID": "x", "SECRET": "x", "FROM": "+10000000000"}, "TOKEN_LENGTH": 6, "MESSAGE": "{otp}", "APP_NAME": "example", "SECURITY_CODE_EXPIRATION_TIME": 3600, "VERIFY_SECURITY_CODE_ONLY_ONCE": False}
