"""Scaffold examples/<name>: uv run scripts/new_example.py NAME "dep1,dep2" "app1,app2" [extra settings]"""

import sys
from pathlib import Path

name, deps, apps = sys.argv[1:4]
extra = sys.argv[4] if len(sys.argv) > 4 else ""
d = Path(__file__).resolve().parent.parent / "examples" / name
d.mkdir(parents=True, exist_ok=True)
dep_list = ", ".join(f'"{x.strip()}"' for x in deps.split(",") if x.strip())
(d / "pyproject.toml").write_text(f"""[project]
name = "example-{name}"
version = "0.0.0"
requires-python = ">=3.11"
dependencies = ["djangoer>=0.5.0", {dep_list}]

[tool.uv]
package = false

[tool.example]
all = true
""")
base = [
    "django.contrib.admin",
    "django.contrib.auth",
    "django.contrib.contenttypes",
    "django.contrib.sessions",
    "django.contrib.messages",
    "django.contrib.staticfiles",
]
app_list = base + [a.strip() for a in apps.split(",") if a.strip()]
apps_src = "\n".join(f'    "{a}",' for a in dict.fromkeys(app_list))
(d / "settings.py").write_text(f"""SECRET_KEY = "example"
USE_TZ = True
SITE_ID = 1
STATIC_URL = "/static/"
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
DATABASES = {{"default": {{"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}}}}
TEMPLATES = [{{"BACKEND": "django.template.backends.django.DjangoTemplates", "APP_DIRS": True, "DIRS": [],
    "OPTIONS": {{"context_processors": ["django.template.context_processors.request",
    "django.contrib.auth.context_processors.auth", "django.contrib.messages.context_processors.messages"]}}}}]
INSTALLED_APPS = [
{apps_src}
]
{extra}
""")
(d / "BUILD.bazel").write_text('load("//apps/djangoer/tools:defs.bzl", "djangoer_example")\n\ndjangoer_example()\n')
print(
    f"wrote {d}; lock it with `uv lock --project {d}`, then `bazel run //apps/djangoer/examples/{name}:generate.write`"
)
