import pymysql

# PyMySQL in place of mysqlclient, which needs the MySQL client library to build
pymysql.version_info = (2, 2, 1, "final", 0)
pymysql.install_as_MySQLdb()

SECRET_KEY = "example"
USE_TZ = True
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
# No server needs to be running: Djangoer assumes a MySQL version when it cannot ask one.
DATABASES = {"default": {"ENGINE": "django.db.backends.mysql", "NAME": "sample", "HOST": "127.0.0.1"}}
INSTALLED_APPS = [
    "django.contrib.contenttypes",
    "django.contrib.auth",
    "store",
]
