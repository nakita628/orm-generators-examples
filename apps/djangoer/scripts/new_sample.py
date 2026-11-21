"""Scaffold a sample: uv run scripts/new_sample.py NAME "app1,app2" [postgres] [extra deps]

The example gets a pyproject.toml, a settings.py that installs the apps after Django's own, and
an empty package per app. Its models.py is written by hand.
"""

import sys
from pathlib import Path

name, apps = sys.argv[1], sys.argv[2].split(",")
postgres = len(sys.argv) > 3 and sys.argv[3] == "postgres"
extra = [d for d in (sys.argv[4].split(",") if len(sys.argv) > 4 else []) if d]
d = Path(__file__).resolve().parent.parent / "examples" / name
deps = ["djangoer>=0.5.0", "django>=6.0", *(["psycopg[binary]"] if postgres else []), *extra]
(d).mkdir(parents=True, exist_ok=True)
(d / "pyproject.toml").write_text(f"""[project]
name = "example-{name}"
version = "0.0.0"
requires-python = ">=3.12"
dependencies = [{", ".join(f'"{x}"' for x in deps)}]

[tool.uv]
package = false

[tool.example]
# A sample written here: Djangoer runs without --all, and its apps are the project's own.
""")
engine = (
    '{"ENGINE": "django.db.backends.postgresql", "NAME": "sample"}'
    if postgres
    else '{"ENGINE": "django.db.backends.sqlite3", "NAME": ":memory:"}'
)
installed = [
    "django.contrib.contenttypes",
    "django.contrib.auth",
    *(["django.contrib.postgres"] if postgres else []),
    *apps,
]
(d / "settings.py").write_text(f"""SECRET_KEY = "example"
USE_TZ = True
DEFAULT_AUTO_FIELD = "django.db.models.BigAutoField"
DATABASES = {{"default": {engine}}}
INSTALLED_APPS = [
{chr(10).join(f'    "{a}",' for a in installed)}
]
""")
for app in apps:
    (d / app).mkdir(exist_ok=True)
    (d / app / "__init__.py").touch()
(d / "BUILD.bazel").write_text('load("//apps/djangoer/tools:defs.bzl", "djangoer_example")\n\ndjangoer_example()\n')
print(
    f"wrote {d}; lock it with `uv lock --project {d}`, then `bazel run //apps/djangoer/examples/{name}:generate.write`"
)
