"""Runs the plan through the SQLAlchemy models hekireki wrote, against the database Prisma created.

    harness.py crud <plan.json> <models.py> <url>   insert, read, update, relations, delete
    harness.py ddl <plan.json> <models.py> <url>    metadata.create_all() into an empty database

Columns are found through the mappers (``mapper.get_property_by_column``), so an attribute named
anything at all still lines up with the plan's column; a column no mapper maps is a finding.
"""

from __future__ import annotations

import base64
import enum
import importlib.util
import json
import sys
import uuid
from datetime import date, datetime, time, timezone
from decimal import Decimal
from typing import Any

from sqlalchemy import Date, Enum, Time, create_engine, inspect, select
from sqlalchemy.orm import DeclarativeBase, Mapper, Session


def main(mode: str, plan_path: str, models_path: str, url: str) -> None:
    plan = json.loads(open(plan_path).read())
    spec = importlib.util.spec_from_file_location("hekireki_models", models_path)
    assert spec and spec.loader
    module = importlib.util.module_from_spec(spec)
    sys.modules["hekireki_models"] = module
    spec.loader.exec_module(module)
    base = next(
        v
        for v in vars(module).values()
        if isinstance(v, type) and issubclass(v, DeclarativeBase) and v is not DeclarativeBase and v.__dict__.get("registry") is not None
    )
    engine = create_engine(sqlalchemy_url(url))

    if mode == "ddl":
        base.metadata.create_all(engine)
        return

    mappers: dict[str, Mapper[Any]] = {m.local_table.name: m for m in base.registry.mappers}  # type: ignore[attr-defined]
    out: dict[str, Any] = {"orm": "sqlalchemy", "models": []}
    rows: dict[str, dict[str, Any]] = {}
    state: dict[str, dict[str, Any]] = {}

    for m in plan["models"]:
        r: dict[str, Any] = {"model": m["name"]}
        out["models"].append(r)
        mapper = mappers.get(m["table"])
        if mapper is None:
            r["missing"] = True
            continue
        columns = column_attrs(mapper)
        r["unmapped"] = [f["column"] for f in m["fields"] if f["column"] not in columns]
        if not m["insertable"]:
            r["skipped"] = "a required key closes a cycle"
            continue
        cls = mapper.class_
        obj = cls()
        try:
            for f in m["fields"]:
                if f["column"] not in columns:
                    continue
                attr, col = columns[f["column"]]
                if f["write"] == "value":
                    setattr(obj, attr, to_py(f, f["value"], col))
                elif f["write"] == "parent":
                    parent = rows.get(f["fk"]["model"]) if f["fk"]["available"] else None
                    setattr(obj, attr, to_py(f, parent[f["fk"]["column"]], col) if parent else None)
            with Session(engine, expire_on_commit=False) as session:
                session.add(obj)
                session.commit()
                identity = inspect(obj).identity
            r["insert"] = {"ok": True}
        except Exception as e:  # noqa: BLE001
            r["insert"] = {"ok": False, "error": message(e)}
            continue
        r["pk"] = {col: normal(v) for col, v in zip(m["pk"], identity or ())}
        try:
            with Session(engine) as session:
                loaded = session.get(cls, identity)
                if loaded is None:
                    raise LookupError("no row came back for the key the insert returned")
                read = {col: normal(getattr(loaded, attr)) for col, (attr, _) in columns.items()}
            r["read"] = {"ok": True, "values": read}
            rows[m["name"]] = read
        except Exception as e:  # noqa: BLE001
            r["read"] = {"ok": False, "error": message(e)}
            continue
        if m.get("update"):
            try:
                if m["update"]["column"] not in columns:
                    raise LookupError(f"no attribute maps {m['update']['column']}")
                attr, _ = columns[m["update"]["column"]]
                with Session(engine) as session:
                    loaded = session.get(cls, identity)
                    setattr(loaded, attr, m["update"]["value"])
                    session.commit()
                with Session(engine) as session:
                    r["update"] = {"ok": True, "value": normal(getattr(session.get(cls, identity), attr))}
            except Exception as e:  # noqa: BLE001
                r["update"] = {"ok": False, "error": message(e)}
        state[m["name"]] = {"cls": cls, "identity": identity, "mapper": mapper}

    by_model = {r["model"]: r for r in out["models"]}
    for m in plan["models"]:
        if m["name"] not in state:
            continue
        s = state[m["name"]]
        relations = []
        by_model[m["name"]]["relations"] = relations
        for rel in m["relations"]:
            result: dict[str, Any] = {"name": rel["name"], "kind": rel["kind"]}
            relations.append(result)
            if rel["model"] not in state:
                result["skipped"] = "no row on the other side"
                continue
            key = find_relation(s["mapper"], rel["name"])
            if key is None:
                result["found"] = False
                continue
            result["found"] = True
            try:
                with Session(engine) as session:
                    loaded = getattr(session.get(s["cls"], s["identity"]), key)
                count = 0 if loaded is None else len(loaded) if isinstance(loaded, (list, set, tuple)) else 1
                read = by_model[m["name"]]["read"]["values"]
                expected = all(read.get(c) is not None for c in rel["columns"]) if rel["kind"] == "toOne" else True
                result["ok"] = count > 0 if expected else count == 0
                result["loaded"] = count
            except Exception as e:  # noqa: BLE001
                result["ok"] = False
                result["error"] = message(e)

    for m in reversed(plan["models"]):
        if m["name"] not in state:
            continue
        s = state[m["name"]]
        r = by_model[m["name"]]
        try:
            with Session(engine) as session:
                session.delete(session.get(s["cls"], s["identity"]))
                session.commit()
            with Session(engine) as session:
                gone = session.get(s["cls"], s["identity"]) is None
            r["delete"] = {"ok": gone} if gone else {"ok": False, "error": "the row is still there"}
        except Exception as e:  # noqa: BLE001
            r["delete"] = {"ok": False, "error": message(e)}

    sys.stdout.write(json.dumps(out))


def column_attrs(mapper: Mapper[Any]) -> dict[str, tuple[str, Any]]:
    out = {}
    for col in mapper.local_table.columns:
        try:
            prop = mapper.get_property_by_column(col)
        except Exception:  # noqa: BLE001
            continue
        out[col.name] = (prop.key, col)
    return out


def find_relation(mapper: Mapper[Any], prisma_name: str) -> str | None:
    squash = lambda s: s.lower().replace("_", "")  # noqa: E731
    for rel in mapper.relationships:
        if rel.key == prisma_name or squash(rel.key) == squash(prisma_name):
            return rel.key
    return None


def to_py(f: dict[str, Any], value: Any, col: Any) -> Any:
    if value is None:
        return None
    if f["list"]:
        return [to_py({**f, "list": False}, v, col) for v in value]
    t = f["type"]
    if t == "DateTime":
        dt = datetime.fromisoformat(value.replace("Z", "+00:00"))
        if isinstance(col.type, Date):
            return dt.date()
        if isinstance(col.type, Time):
            return dt.time()
        return dt
    if t == "BigInt":
        return int(value)
    if t == "Decimal":
        return Decimal(value)
    if t == "Bytes":
        return base64.b64decode(value)
    if t == "Enum" and isinstance(col.type, Enum) and col.type.enum_class is not None:
        cls = col.type.enum_class
        for member in cls:
            if member.value == value or member.name == value:
                return member
    return value


def normal(value: Any) -> Any:
    if isinstance(value, datetime):
        if value.tzinfo is None:
            value = value.replace(tzinfo=timezone.utc)
        return value.astimezone(timezone.utc).strftime("%Y-%m-%dT%H:%M:%S.") + f"{value.microsecond // 1000:03d}Z"
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


def sqlalchemy_url(url: str) -> str:
    if url.startswith("postgresql://"):
        return "postgresql+psycopg://" + url[len("postgresql://") :]
    if url.startswith("mysql://"):
        return "mysql+pymysql://" + url[len("mysql://") :]
    if url.startswith("file:"):
        return "sqlite:///" + url[len("file:") :]
    raise ValueError(url)


def message(e: BaseException) -> str:
    text = str(getattr(e, "orig", None) or e).strip()
    return f"{type(e).__name__}: {text}"[:600]


if __name__ == "__main__":
    main(*sys.argv[1:])
