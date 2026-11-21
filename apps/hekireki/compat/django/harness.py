"""Runs the plan through the Django models hekireki wrote, against the database Prisma created.

    harness.py crud <plan.json> <project dir> <url>   insert, read, update, relations, delete
    harness.py ddl <plan.json> <project dir> <url>    makemigrations + migrate into an empty database

<project dir> holds app/models.py, the generated file (run.mjs copies it there with an empty
app/migrations/). Columns are found through ``Field.column``, so the plan's column lines up with
whatever the attribute is called; a column no field maps is a finding.
"""

from __future__ import annotations

import base64
import enum
import json
import sys
import uuid
from datetime import date, datetime, time, timezone
from decimal import Decimal
from typing import Any
from urllib.parse import unquote, urlparse


def main(mode: str, plan_path: str, project: str, url: str) -> None:
    plan = json.loads(open(plan_path).read())
    sys.path.insert(0, project)
    import django
    from django.conf import settings

    settings.configure(
        SECRET_KEY="hekireki-compat",
        INSTALLED_APPS=["app"],
        DATABASES={"default": database(url)},
        USE_TZ=True,
        TIME_ZONE="UTC",
        DEFAULT_AUTO_FIELD="django.db.models.BigAutoField",
    )
    django.setup()
    from django.apps import apps
    from django.core.management import call_command
    from django.db import connection

    if mode == "ddl":
        call_command("makemigrations", "app", verbosity=0)
        call_command("migrate", "app", verbosity=0)
        with connection.cursor() as cursor:
            cursor.execute("DROP TABLE django_migrations")
        return

    models = {m._meta.db_table: m for m in apps.get_app_config("app").get_models()}
    out: dict[str, Any] = {"orm": "django", "models": []}
    rows: dict[str, dict[str, Any]] = {}
    state: dict[str, Any] = {}

    for m in plan["models"]:
        r: dict[str, Any] = {"model": m["name"]}
        out["models"].append(r)
        model = models.get(m["table"])
        if model is None:
            r["missing"] = True
            continue
        columns = {f.column: f for f in model._meta.concrete_fields if getattr(f, "column", None)}
        r["unmapped"] = [f["column"] for f in m["fields"] if f["column"] not in columns]
        if not m["insertable"]:
            r["skipped"] = "a required key closes a cycle"
            continue
        obj = model()
        try:
            for f in m["fields"]:
                field = columns.get(f["column"])
                if field is None:
                    continue
                if f["write"] == "value":
                    setattr(obj, field.attname, to_py(f, f["value"], field))
                elif f["write"] == "parent":
                    parent = rows.get(f["fk"]["model"]) if f["fk"]["available"] else None
                    setattr(obj, field.attname, to_py(f, parent[f["fk"]["column"]], field) if parent else None)
            obj.save(force_insert=True)
            r["insert"] = {"ok": True}
        except Exception as e:  # noqa: BLE001
            r["insert"] = {"ok": False, "error": message(e)}
            continue
        pk = obj.pk
        r["pk"] = {col: normal(getattr(obj, columns[col].attname)) for col in m["pk"] if col in columns}
        try:
            loaded = model.objects.get(pk=pk)
            read = {col: normal(getattr(loaded, field.attname)) for col, field in columns.items()}
            r["read"] = {"ok": True, "values": read}
            rows[m["name"]] = read
        except Exception as e:  # noqa: BLE001
            r["read"] = {"ok": False, "error": message(e)}
            continue
        if m.get("update"):
            try:
                field = columns.get(m["update"]["column"])
                if field is None:
                    raise LookupError(f"no field maps {m['update']['column']}")
                loaded = model.objects.get(pk=pk)
                setattr(loaded, field.attname, m["update"]["value"])
                loaded.save()
                r["update"] = {"ok": True, "value": normal(getattr(model.objects.get(pk=pk), field.attname))}
            except Exception as e:  # noqa: BLE001
                r["update"] = {"ok": False, "error": message(e)}
        state[m["name"]] = (model, pk)

    by_model = {r["model"]: r for r in out["models"]}
    for m in plan["models"]:
        if m["name"] not in state:
            continue
        model, pk = state[m["name"]]
        relations = []
        by_model[m["name"]]["relations"] = relations
        for rel in m["relations"]:
            result: dict[str, Any] = {"name": rel["name"], "kind": rel["kind"]}
            relations.append(result)
            if rel["model"] not in state:
                result["skipped"] = "no row on the other side"
                continue
            accessor = find_relation(model, rel["name"])
            if accessor is None:
                result["found"] = False
                continue
            result["found"] = True
            try:
                value = getattr(model.objects.get(pk=pk), accessor, None)
                if value is not None and hasattr(value, "all"):
                    count = value.all().count()
                else:
                    count = 0 if value is None else 1
                read = by_model[m["name"]]["read"]["values"]
                expected = all(read.get(c) is not None for c in rel["columns"]) if rel["kind"] == "toOne" else True
                result["ok"] = count > 0 if expected else count == 0
                result["loaded"] = count
            except Exception as e:  # noqa: BLE001
                if type(e).__name__ == "RelatedObjectDoesNotExist":
                    result["ok"] = rel["kind"] != "toOne"
                    result["loaded"] = 0
                else:
                    result["ok"] = False
                    result["error"] = message(e)

    for m in reversed(plan["models"]):
        if m["name"] not in state:
            continue
        model, pk = state[m["name"]]
        r = by_model[m["name"]]
        try:
            model.objects.filter(pk=pk).delete()
            gone = not model.objects.filter(pk=pk).exists()
            r["delete"] = {"ok": True} if gone else {"ok": False, "error": "the row is still there"}
        except Exception as e:  # noqa: BLE001
            r["delete"] = {"ok": False, "error": message(e)}

    sys.stdout.write(json.dumps(out))


def find_relation(model: Any, prisma_name: str) -> str | None:
    squash = lambda s: s.lower().replace("_", "")  # noqa: E731
    for field in model._meta.get_fields():
        if not field.is_relation:
            continue
        names = [field.name]
        if hasattr(field, "get_accessor_name"):
            names.append(field.get_accessor_name())
        for name in names:
            if name and squash(name) == squash(prisma_name):
                return name
    return None


def to_py(f: dict[str, Any], value: Any, field: Any) -> Any:
    from django.db import models

    if value is None:
        return None
    if f["list"]:
        return [to_py({**f, "list": False}, v, getattr(field, "base_field", field)) for v in value]
    t = f["type"]
    if t == "DateTime":
        dt = datetime.fromisoformat(value.replace("Z", "+00:00"))
        if isinstance(field, models.DateTimeField):
            return dt
        if isinstance(field, models.DateField):
            return dt.date()
        if isinstance(field, models.TimeField):
            return dt.time()
        return dt
    if t == "BigInt":
        return int(value)
    if t == "Decimal":
        return Decimal(value)
    if t == "Bytes":
        return base64.b64decode(value)
    if isinstance(field, models.UUIDField):
        return uuid.UUID(value)
    return value


def normal(value: Any) -> Any:
    if isinstance(value, datetime):
        if value.tzinfo is None:
            value = value.replace(tzinfo=timezone.utc)
        value = value.astimezone(timezone.utc)
        return value.strftime("%Y-%m-%dT%H:%M:%S.") + f"{value.microsecond // 1000:03d}Z"
    if isinstance(value, date):
        return value.strftime("%Y-%m-%dT00:00:00.000Z")
    if isinstance(value, time):
        return value.strftime("1970-01-01T%H:%M:%S.") + f"{value.microsecond // 1000:03d}Z"
    if isinstance(value, Decimal):
        return str(value)
    if isinstance(value, (bytes, bytearray, memoryview)):
        return base64.b64encode(bytes(value)).decode()
    if isinstance(value, enum.Enum):
        return value.value
    if isinstance(value, uuid.UUID):
        return str(value)
    if isinstance(value, (list, tuple)):
        return [normal(v) for v in value]
    if isinstance(value, int) and not isinstance(value, bool) and abs(value) > 2**53:
        return str(value)
    return value


def database(url: str) -> dict[str, Any]:
    if url.startswith("file:"):
        return {"ENGINE": "django.db.backends.sqlite3", "NAME": url[len("file:") :]}
    u = urlparse(url)
    engine = {"postgresql": "django.db.backends.postgresql", "mysql": "django.db.backends.mysql"}[u.scheme]
    return {
        "ENGINE": engine,
        "NAME": u.path.lstrip("/"),
        "USER": unquote(u.username or ""),
        "PASSWORD": unquote(u.password or ""),
        "HOST": u.hostname,
        "PORT": str(u.port),
        **({"OPTIONS": {"charset": "utf8mb4"}} if u.scheme == "mysql" else {}),
    }


def message(e: BaseException) -> str:
    return f"{type(e).__name__}: {str(e).strip()}"[:600]


if __name__ == "__main__":
    main(*sys.argv[1:])
