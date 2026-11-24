// Runs the plan through the entities and DbContext hekireki-efcore wrote, against the database
// Prisma created.
//
//   dotnet run -- crud <plan.json> <url>   insert, read, update, relations, delete
//   dotnet run -- ddl <plan.json> <url>    context.Database.EnsureCreated() into an empty database
//
// Columns are found through EF Core's model (IProperty.GetColumnName), so the plan's column lines
// up with whatever the property is called; a column no property maps is a finding.
#nullable disable
using System;
using System.Collections;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text.Json;
using System.Text.Json.Nodes;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata;

public static class Harness
{
    static JsonElement Plan;
    static Type ContextType;
    static string Url;
#if PG
    static Npgsql.NpgsqlDataSource DataSource;
#endif

    public static int Main(string[] args)
    {
        // An unhandled exception leaves the process spinning in the runtime's crash handling
        // instead of exiting: report it and exit.
        try
        {
            return Run(args);
        }
        catch (Exception e)
        {
            Console.Error.WriteLine(Message(e));
            Console.Error.Flush();
            Environment.Exit(1);
            return 1;
        }
    }

    static int Run(string[] args)
    {
        var mode = args[0];
        Plan = JsonDocument.Parse(File.ReadAllText(args[1])).RootElement;
        Url = args[2];
        ContextType = typeof(Harness).Assembly.GetTypes().First(t => t.IsSubclassOf(typeof(DbContext)) && !t.IsAbstract);
        if (mode == "ddl")
        {
            using var ctx = NewContext();
            ctx.Database.EnsureCreated();
            return 0;
        }
        Console.Out.Write(Crud().ToJsonString());
        return 0;
    }

    static DbContext NewContext()
    {
        var builderType = typeof(DbContextOptionsBuilder<>).MakeGenericType(ContextType);
        var builder = (DbContextOptionsBuilder)Activator.CreateInstance(builderType);
        var u = new Uri(Url.StartsWith("file:") ? "file:///x" : Url);
        var user = Uri.UnescapeDataString(u.UserInfo.Split(':')[0]);
        var pass = u.UserInfo.Contains(':') ? Uri.UnescapeDataString(u.UserInfo.Split(':')[1]) : "";
        var db = u.AbsolutePath.TrimStart('/');
#if PG
        // The context's own enum mapping (MapEnums), on the data source and the options, as an
        // application wires it.
        // One data source (one connection pool) for every context, as an application has.
        if (DataSource == null)
        {
            var dsb = new Npgsql.NpgsqlDataSourceBuilder($"Host={u.Host};Port={u.Port};Database={db};Username={user};Password={pass};Maximum Pool Size=5");
            ContextType.GetMethod("MapEnums", new[] { typeof(Npgsql.NpgsqlDataSourceBuilder) })?.Invoke(null, new object[] { dsb });
            DataSource = dsb.Build();
        }
        var dataSource = DataSource;
        builder.UseNpgsql(dataSource, o =>
        {
            ContextType.GetMethod("MapEnums", new[] { typeof(Npgsql.EntityFrameworkCore.PostgreSQL.Infrastructure.NpgsqlDbContextOptionsBuilder) })?.Invoke(null, new object[] { o });
        });
#elif MY
        var cs = $"Server={u.Host};Port={u.Port};Database={db};User={user};Password={pass}";
        builder.UseMySql(cs, new MySqlServerVersion(new Version(8, 4, 0)));
#else
        builder.UseSqlite($"Data Source={Url.Substring("file:".Length)}");
#endif
        return (DbContext)Activator.CreateInstance(ContextType, builder.Options);
    }

    static JsonObject Crud()
    {
        var results = new JsonArray();
        var rows = new Dictionary<string, JsonObject>();
        var keys = new Dictionary<string, (IEntityType et, object[] key)>();
        IModel model;
        using (var probe = NewContext()) model = probe.Model;
        var byTable = new Dictionary<string, IEntityType>();
        foreach (var et in model.GetEntityTypes())
        {
            var t = et.GetTableName();
            if (t != null && !et.IsOwned()) byTable.TryAdd(t, et);
        }

        foreach (var m in Plan.GetProperty("models").EnumerateArray())
        {
            var r = new JsonObject { ["model"] = m.GetProperty("name").GetString() };
            results.Add(r);
            if (!byTable.TryGetValue(m.GetProperty("table").GetString(), out var et))
            {
                r["missing"] = true;
                continue;
            }
            var cols = Columns(et);
            var unmapped = new JsonArray();
            foreach (var f in m.GetProperty("fields").EnumerateArray())
                if (!cols.ContainsKey(f.GetProperty("column").GetString())) unmapped.Add(f.GetProperty("column").GetString());
            r["unmapped"] = unmapped;
            if (!m.GetProperty("insertable").GetBoolean())
            {
                r["skipped"] = "a required key closes a cycle";
                continue;
            }
            object entity;
            try
            {
                entity = Activator.CreateInstance(et.ClrType);
                foreach (var f in m.GetProperty("fields").EnumerateArray())
                {
                    if (!cols.TryGetValue(f.GetProperty("column").GetString(), out var p) || p.PropertyInfo == null) continue;
                    var write = f.GetProperty("write").GetString();
                    JsonNode value = null;
                    if (write == "value") value = JsonNode.Parse(f.GetProperty("value").GetRawText());
                    else if (write == "parent")
                    {
                        var fk = f.GetProperty("fk");
                        if (fk.GetProperty("available").GetBoolean() && rows.TryGetValue(fk.GetProperty("model").GetString(), out var parent))
                            value = parent[fk.GetProperty("column").GetString()]?.DeepClone();
                    }
                    else continue;
                    p.PropertyInfo.SetValue(entity, ToClr(value, p.ClrType, f, p));
                }
                using var ctx = NewContext();
                ctx.Add(entity);
                ctx.SaveChanges();
                r["insert"] = new JsonObject { ["ok"] = true };
            }
            catch (Exception e)
            {
                r["insert"] = Error(e);
                continue;
            }
            var pkProps = et.FindPrimaryKey().Properties;
            var key = pkProps.Select(p => p.PropertyInfo.GetValue(entity)).ToArray();
            var pkJson = new JsonObject();
            foreach (var p in pkProps) pkJson[Column(p)] = Normal(p.PropertyInfo.GetValue(entity), p);
            r["pk"] = pkJson;
            try
            {
                using var ctx = NewContext();
                var loaded = ctx.Find(et.ClrType, key) ?? throw new Exception("no row came back for the key the insert returned");
                var values = new JsonObject();
                foreach (var (col, p) in cols) if (p.PropertyInfo != null) values[col] = Normal(p.PropertyInfo.GetValue(loaded), p);
                r["read"] = new JsonObject { ["ok"] = true, ["values"] = values };
                rows[m.GetProperty("name").GetString()] = values;
            }
            catch (Exception e)
            {
                r["read"] = Error(e);
                continue;
            }
            if (m.TryGetProperty("update", out var upd) && upd.ValueKind == JsonValueKind.Object)
            {
                try
                {
                    var col = upd.GetProperty("column").GetString();
                    var p = cols[col];
                    var f = m.GetProperty("fields").EnumerateArray().First(x => x.GetProperty("column").GetString() == col);
                    using (var ctx = NewContext())
                    {
                        var row = ctx.Find(et.ClrType, key);
                        p.PropertyInfo.SetValue(row, ToClr(JsonNode.Parse(upd.GetProperty("value").GetRawText()), p.ClrType, f, p));
                        ctx.SaveChanges();
                    }
                    using (var ctx = NewContext())
                        r["update"] = new JsonObject { ["ok"] = true, ["value"] = Normal(p.PropertyInfo.GetValue(ctx.Find(et.ClrType, key)), p) };
                }
                catch (Exception e)
                {
                    r["update"] = Error(e);
                }
            }
            keys[m.GetProperty("name").GetString()] = (et, key);
        }

        var models = Plan.GetProperty("models").EnumerateArray().ToList();
        for (var i = 0; i < models.Count; i++)
        {
            var m = models[i];
            var name = m.GetProperty("name").GetString();
            if (!keys.TryGetValue(name, out var k)) continue;
            var rels = new JsonArray();
            foreach (var rel in m.GetProperty("relations").EnumerateArray())
            {
                var res = new JsonObject { ["name"] = rel.GetProperty("name").GetString(), ["kind"] = rel.GetProperty("kind").GetString() };
                rels.Add(res);
                if (!keys.ContainsKey(rel.GetProperty("model").GetString()))
                {
                    res["skipped"] = "no row on the other side";
                    continue;
                }
                // A member cannot have its class's name (CS0542): a relation named as its own model
                // (categories_category on categories_category) is CategoriesCategory1, as EF Core's
                // scaffolding names it.
                var relName = Squash(rel.GetProperty("name").GetString());
                if (relName == Squash(k.et.ClrType.Name)) relName += "1";
                var nav = k.et.GetNavigations().Cast<INavigationBase>().Concat(k.et.GetSkipNavigations())
                    .FirstOrDefault(n => Squash(n.Name) == relName);
                if (nav == null)
                {
                    res["found"] = false;
                    continue;
                }
                res["found"] = true;
                try
                {
                    using var ctx = NewContext();
                    var row = ctx.Find(k.et.ClrType, k.key);
                    ctx.Entry(row).Navigation(nav.Name).Load();
                    var v = nav.PropertyInfo.GetValue(row);
                    res["loaded"] = v == null ? 0 : v is IEnumerable e && v is not string ? e.Cast<object>().Count() : 1;
                }
                catch (Exception e)
                {
                    res["error"] = Message(e);
                }
            }
            ((JsonObject)results[i])["relations"] = rels;
        }

        for (var i = models.Count - 1; i >= 0; i--)
        {
            var name = models[i].GetProperty("name").GetString();
            if (!keys.TryGetValue(name, out var k)) continue;
            var r = (JsonObject)results[i];
            try
            {
                using (var ctx = NewContext())
                {
                    ctx.Remove(ctx.Find(k.et.ClrType, k.key));
                    ctx.SaveChanges();
                }
                using (var ctx = NewContext())
                    r["delete"] = ctx.Find(k.et.ClrType, k.key) == null ? new JsonObject { ["ok"] = true } : new JsonObject { ["ok"] = false, ["error"] = "the row is still there" };
            }
            catch (Exception e)
            {
                r["delete"] = Error(e);
            }
        }
        return new JsonObject { ["orm"] = "efcore", ["models"] = results };
    }

    static Dictionary<string, IProperty> Columns(IEntityType et)
    {
        var d = new Dictionary<string, IProperty>();
        foreach (var p in et.GetProperties())
        {
            var c = Column(p);
            if (c != null) d.TryAdd(c, p);
        }
        return d;
    }

    static string Column(IProperty p)
    {
        var et = p.DeclaringType as IEntityType;
        var store = StoreObjectIdentifier.Table(et.GetTableName(), et.GetSchema());
        return p.GetColumnName(store);
    }

    static string Squash(string s) => s.ToLowerInvariant().Replace("_", "");

    static object ToClr(JsonNode node, Type type, JsonElement f, IProperty p)
    {
        if (node == null) return null;
        var inner = Nullable.GetUnderlyingType(type) ?? type;
        if (inner != typeof(string) && inner != typeof(byte[]) && typeof(IEnumerable).IsAssignableFrom(inner))
        {
            var elem = inner.IsArray ? inner.GetElementType() : inner.GetGenericArguments()[0];
            var items = node.AsArray().Select(x => ToClr(x, elem, f, p)).ToList();
            if (inner.IsArray)
            {
                var arr = Array.CreateInstance(elem, items.Count);
                for (var i = 0; i < items.Count; i++) arr.SetValue(items[i], i);
                return arr;
            }
            var list = (IList)Activator.CreateInstance(typeof(List<>).MakeGenericType(elem));
            foreach (var x in items) list.Add(x);
            return list;
        }
        var type0 = f.GetProperty("type").GetString();
        if (inner == typeof(string))
            return type0 == "Json" ? node.ToJsonString() : node.GetValueKind() == JsonValueKind.String ? node.GetValue<string>() : node.ToJsonString();
        var text = node.GetValueKind() == JsonValueKind.String ? node.GetValue<string>() : node.ToJsonString();
        if (inner == typeof(DateTime))
        {
            var dt = DateTime.Parse(text, CultureInfo.InvariantCulture, DateTimeStyles.AdjustToUniversal | DateTimeStyles.AssumeUniversal);
            var colType = (p.GetColumnType() ?? "").ToLowerInvariant();
            var utc = colType.Contains("with time zone") || colType.Contains("timestamptz");
            return DateTime.SpecifyKind(dt, utc ? DateTimeKind.Utc : DateTimeKind.Unspecified);
        }
        if (inner == typeof(DateTimeOffset)) return DateTimeOffset.Parse(text, CultureInfo.InvariantCulture).ToUniversalTime();
        if (inner == typeof(DateOnly)) return DateOnly.FromDateTime(DateTime.Parse(text, CultureInfo.InvariantCulture, DateTimeStyles.AdjustToUniversal));
        if (inner == typeof(TimeOnly)) return TimeOnly.FromDateTime(DateTime.Parse(text, CultureInfo.InvariantCulture, DateTimeStyles.AdjustToUniversal));
        if (inner == typeof(TimeSpan)) return DateTime.Parse(text, CultureInfo.InvariantCulture, DateTimeStyles.AdjustToUniversal).TimeOfDay;
        if (inner == typeof(Guid)) return Guid.Parse(text);
        if (inner == typeof(byte[])) return Convert.FromBase64String(text);
        if (inner == typeof(bool)) return bool.Parse(text);
        if (inner == typeof(decimal)) return decimal.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(double)) return double.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(float)) return float.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(long)) return long.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(int)) return int.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(short)) return short.Parse(text, CultureInfo.InvariantCulture);
        if (inner == typeof(JsonDocument)) return JsonDocument.Parse(node.ToJsonString());
        if (inner == typeof(JsonElement)) return JsonDocument.Parse(node.ToJsonString()).RootElement;
        if (inner.IsEnum)
        {
            foreach (var field in inner.GetFields(BindingFlags.Public | BindingFlags.Static))
            {
                var names = field.GetCustomAttributes().Select(a => a.GetType().GetProperty("PgName")?.GetValue(a) as string ?? a.GetType().GetProperty("Value")?.GetValue(a) as string).Where(x => x != null).Append(field.Name);
                if (names.Contains(text)) return field.GetValue(null);
            }
            return Enum.Parse(inner, text, true);
        }
        var parse = inner.GetMethod("Parse", new[] { typeof(string) });
        if (parse != null) return parse.Invoke(null, new object[] { text });
        return JsonSerializer.Deserialize(node.ToJsonString(), type);
    }

    static JsonNode Normal(object v, IProperty p)
    {
        // An enum stored through a value converter (MySQL's ENUM, SQLite's text): what the column
        // holds is the converter's provider value.
        if (v is Enum && p.GetValueConverter() is { } converter) v = converter.ConvertToProvider(v);
        switch (v)
        {
            case null: return null;
            case DateTime d:
                var u = d.Kind == DateTimeKind.Local ? d.ToUniversalTime() : DateTime.SpecifyKind(d, DateTimeKind.Utc);
                return u.ToString("yyyy-MM-dd'T'HH:mm:ss.fff'Z'", CultureInfo.InvariantCulture);
            case DateTimeOffset o: return o.UtcDateTime.ToString("yyyy-MM-dd'T'HH:mm:ss.fff'Z'", CultureInfo.InvariantCulture);
            case DateOnly d0: return d0.ToString("yyyy-MM-dd", CultureInfo.InvariantCulture) + "T00:00:00.000Z";
            case TimeOnly t: return "1970-01-01T" + t.ToString("HH:mm:ss.fff", CultureInfo.InvariantCulture) + "Z";
            case TimeSpan ts: return "1970-01-01T" + new DateTime(ts.Ticks).ToString("HH:mm:ss.fff", CultureInfo.InvariantCulture) + "Z";
            case decimal m: return m.ToString(CultureInfo.InvariantCulture);
            case long l: return Math.Abs(l) > 9007199254740992L ? JsonValue.Create(l.ToString(CultureInfo.InvariantCulture)) : JsonValue.Create(l);
            case byte[] b: return Convert.ToBase64String(b);
            case Guid g: return g.ToString();
            case JsonDocument jd: return JsonNode.Parse(jd.RootElement.GetRawText());
            case JsonElement je: return JsonNode.Parse(je.GetRawText());
            case Enum e:
                var field = e.GetType().GetField(e.ToString());
                var name = field?.GetCustomAttributes().Select(a => a.GetType().GetProperty("PgName")?.GetValue(a) as string ?? a.GetType().GetProperty("Value")?.GetValue(a) as string).FirstOrDefault(x => x != null);
                return name ?? e.ToString();
            case string s:
                var colType = (p.GetColumnType() ?? "").ToLowerInvariant();
                // A Json column is a string property: its text is the JSON value.
                if ((colType.Contains("json") || (s.Length > 0 && (s[0] == '{' || s[0] == '['))) && IsJson(s)) return JsonNode.Parse(s);
                return s;
            case IEnumerable list:
                var arr = new JsonArray();
                foreach (var x in list) arr.Add(Normal(x, p));
                return arr;
            case bool or int or short or double or float: return JsonValue.Create(v);
            default: return JsonValue.Create(v.ToString());
        }
    }

    static bool IsJson(string s)
    {
        try { JsonNode.Parse(s); return true; } catch { return false; }
    }

    static JsonObject Error(Exception e) => new() { ["ok"] = false, ["error"] = Message(e) };

    static string Message(Exception e)
    {
        var inner = e;
        while (inner.InnerException != null) inner = inner.InnerException;
        var text = inner == e ? e.Message : $"{inner.Message} ({e.GetType().Name})";
        return text.Length > 600 ? text.Substring(0, 600) : text;
    }
}
