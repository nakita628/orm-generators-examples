# Runs the plan through the Ecto schemas hekireki wrote, against the database Prisma created.
#
#   mix run harness.exs crud <plan.json> <url>
#
# Columns are found through each schema's own reflection (__schema__(:field_source, field)), so
# the plan's column lines up with whatever the field is called; a column no field maps is a
# finding. Values go in through Ecto.Type.cast/2, as a changeset would cast params.

defmodule Compat.Harness do
  def main(["crud", plan_path, url]) do
    plan = plan_path |> File.read!() |> Jason.decode!()
    repo = start_repo(plan["dialect"], url)

    {:ok, modules} = :application.get_key(:compat, :modules)

    schemas =
      for m <- modules,
          Code.ensure_loaded?(m),
          function_exported?(m, :__schema__, 1),
          m.__schema__(:source) != nil,
          into: %{},
          do: {m.__schema__(:source), m}

    {results, state} =
      Enum.reduce(plan["models"], {[], %{rows: %{}, keys: %{}}}, fn m, {acc, st} ->
        {r, st} = run_model(repo, schemas, m, st)
        {[r | acc], st}
      end)

    results = Enum.reverse(results)

    results =
      Enum.map(Enum.zip(plan["models"], results), fn {m, r} ->
        case st_key(state, m) do
          nil -> r
          {mod, key} -> Map.put(r, "relations", relations(repo, mod, key, m, state))
        end
      end)

    results =
      plan["models"]
      |> Enum.zip(results)
      |> Enum.reverse()
      |> Enum.map(fn {m, r} ->
        case st_key(state, m) do
          nil -> r
          {mod, key} -> Map.put(r, "delete", delete(repo, mod, key))
        end
      end)
      |> Enum.reverse()

    IO.write(Jason.encode!(%{"orm" => "ecto", "models" => results}))
  end

  defp st_key(state, m), do: Map.get(state.keys, m["name"])

  defp run_model(repo, schemas, m, st) do
    r = %{"model" => m["name"]}

    case Map.get(schemas, m["table"]) do
      nil ->
        {Map.put(r, "missing", true), st}

      mod ->
        columns = columns(mod)
        unmapped = for f <- m["fields"], not Map.has_key?(columns, f["column"]), do: f["column"]
        r = Map.put(r, "unmapped", unmapped)

        if not m["insertable"] do
          {Map.put(r, "skipped", "a required key closes a cycle"), st}
        else
          insert_model(repo, mod, columns, m, r, st)
        end
    end
  end

  defp insert_model(repo, mod, columns, m, r, st) do
    attempt =
      try do
        row =
          Enum.reduce(m["fields"], struct(mod), fn f, acc ->
            case {Map.get(columns, f["column"]), f["write"]} do
              {nil, _} ->
                acc

              {field, "value"} ->
                Map.put(acc, field, cast!(mod, field, f, f["value"]))

              {field, "parent"} ->
                parent = if f["fk"]["available"], do: Map.get(st.rows, f["fk"]["model"])
                value = if parent, do: cast!(mod, field, f, parent[f["fk"]["column"]])
                Map.put(acc, field, value)

              _ ->
                acc
            end
          end)

        repo.insert(row)
      rescue
        e -> {:error, Exception.message(e)}
      end

    case attempt do
      {:ok, inserted} ->
        pk = for field <- mod.__schema__(:primary_key), do: {field, Map.get(inserted, field)}
        r = Map.merge(r, %{"insert" => %{"ok" => true}, "pk" => Map.new(pk, fn {f, v} -> {source(mod, f), normal(mod, f, v)} end)})

        case safe(fn -> repo.get_by(mod, pk) end) do
          {:ok, nil} ->
            {Map.put(r, "read", %{"ok" => false, "error" => "no row came back for the key the insert returned"}), st}

          {:ok, loaded} ->
            values = Map.new(columns, fn {col, field} -> {col, normal(mod, field, Map.get(loaded, field))} end)
            r = Map.put(r, "read", %{"ok" => true, "values" => values})
            r = if m["update"], do: Map.put(r, "update", update(repo, mod, columns, pk, m)), else: r
            st = %{st | rows: Map.put(st.rows, m["name"], values), keys: Map.put(st.keys, m["name"], {mod, pk})}
            {r, st}

          {:error, e} ->
            {Map.put(r, "read", %{"ok" => false, "error" => e}), st}
        end

      {:error, %Ecto.Changeset{} = cs} ->
        {Map.put(r, "insert", %{"ok" => false, "error" => inspect(cs.errors)}), st}

      {:error, e} ->
        {Map.put(r, "insert", %{"ok" => false, "error" => to_message(e)}), st}
    end
  end

  defp update(repo, mod, columns, pk, m) do
    field = Map.get(columns, m["update"]["column"])

    safe(fn ->
      row = repo.get_by(mod, pk)
      {:ok, _} = row |> Ecto.Changeset.change(%{field => m["update"]["value"]}) |> repo.update()
      normal(mod, field, Map.get(repo.get_by(mod, pk), field))
    end)
    |> case do
      {:ok, v} -> %{"ok" => true, "value" => v}
      {:error, e} -> %{"ok" => false, "error" => e}
    end
  end

  defp relations(repo, mod, pk, m, state) do
    for rel <- m["relations"] do
      result = %{"name" => rel["name"], "kind" => rel["kind"]}

      if not Map.has_key?(state.keys, rel["model"]) do
        Map.put(result, "skipped", "no row on the other side")
      else
        case find_assoc(mod, rel["name"]) do
          nil ->
            Map.put(result, "found", false)

          assoc ->
            case safe(fn -> repo.get_by(mod, pk) |> repo.preload(assoc) |> Map.get(assoc) end) do
              {:ok, loaded} ->
                count =
                  cond do
                    is_list(loaded) -> length(loaded)
                    is_nil(loaded) -> 0
                    true -> 1
                  end

                Map.merge(result, %{"found" => true, "loaded" => count})

              {:error, e} ->
                Map.merge(result, %{"found" => true, "error" => e})
            end
        end
      end
    end
  end

  defp delete(repo, mod, pk) do
    case safe(fn ->
           repo.get_by(mod, pk) |> repo.delete!()
           repo.get_by(mod, pk)
         end) do
      {:ok, nil} -> %{"ok" => true}
      {:ok, _} -> %{"ok" => false, "error" => "the row is still there"}
      {:error, e} -> %{"ok" => false, "error" => e}
    end
  end

  defp columns(mod) do
    for field <- mod.__schema__(:fields), into: %{} do
      {source(mod, field), field}
    end
  end

  defp source(mod, field), do: to_string(mod.__schema__(:field_source, field))

  defp find_assoc(mod, prisma_name) do
    squash = fn s -> s |> to_string() |> String.downcase() |> String.replace("_", "") end
    Enum.find(mod.__schema__(:associations), fn a -> squash.(a) == squash.(prisma_name) end)
  end

  defp cast!(_mod, _field, _f, nil), do: nil

  defp cast!(mod, field, f, value) do
    type = mod.__schema__(:type, field)

    value =
      case f["type"] do
        "Bytes" when not is_nil(value) -> Base.decode64!(value)
        _ -> value
      end

    case Ecto.Type.cast(type, value) do
      {:ok, v} ->
        v

      _ ->
        # Ecto.Enum casts the atom's name; the plan holds the database value.
        case enum_key(type, value) do
          nil -> raise "cannot cast #{inspect(value)} to #{inspect(type)} (#{field})"
          key -> key
        end
    end
  end

  defp enum_key({:parameterized, {Ecto.Enum, %{mappings: mappings}}}, value) do
    Enum.find_value(mappings, fn {k, v} -> if v == value, do: k end)
  end

  defp enum_key({:parameterized, Ecto.Enum, %{mappings: mappings}}, value) do
    Enum.find_value(mappings, fn {k, v} -> if v == value, do: k end)
  end

  defp enum_key({:array, inner}, values) when is_list(values) do
    keys = Enum.map(values, &enum_key(inner, &1))
    if Enum.all?(keys), do: keys
  end

  defp enum_key(_, _), do: nil

  # A value in the plan's JSON form: what the database holds, not Elixir's term for it.
  defp normal(mod, field, value) do
    type = mod.__schema__(:type, field)

    case Ecto.Type.dump(type, value) do
      {:ok, dumped} -> json(dumped, type)
      _ -> json(value, type)
    end
  end

  defp json(nil, _), do: nil
  defp json(%DateTime{} = d, _), do: d |> DateTime.truncate(:millisecond) |> iso()
  defp json(%NaiveDateTime{} = d, _), do: d |> DateTime.from_naive!("Etc/UTC") |> DateTime.truncate(:millisecond) |> iso()
  defp json(%Date{} = d, _), do: Date.to_iso8601(d) <> "T00:00:00.000Z"
  defp json(%Time{} = t, _), do: "1970-01-01T" <> (t |> Time.truncate(:millisecond) |> ms_time()) <> "Z"
  defp json(%Decimal{} = d, _), do: Decimal.to_string(d, :normal)
  defp json(list, {:array, inner}) when is_list(list), do: Enum.map(list, &json(&1, inner))
  defp json(list, _) when is_list(list), do: Enum.map(list, &json(&1, :any))
  defp json(map, _) when is_map(map), do: map |> Jason.encode!() |> Jason.decode!()
  defp json(<<_::128>> = uuid, type) when type in [:binary_id, Ecto.UUID], do: Ecto.UUID.cast!(uuid)
  defp json(bin, :binary) when is_binary(bin), do: Base.encode64(bin)
  defp json(bin, _) when is_binary(bin), do: if(String.valid?(bin), do: bin, else: Base.encode64(bin))
  defp json(atom, _) when is_atom(atom) and atom not in [true, false], do: Atom.to_string(atom)
  defp json(n, _) when is_integer(n) and (n > 9_007_199_254_740_992 or n < -9_007_199_254_740_992), do: Integer.to_string(n)
  defp json(v, _), do: v

  defp iso(%DateTime{} = d) do
    ms = d.microsecond |> elem(0) |> div(1000) |> Integer.to_string() |> String.pad_leading(3, "0")
    Calendar.strftime(d, "%Y-%m-%dT%H:%M:%S.") <> ms <> "Z"
  end

  defp ms_time(t) do
    ms = t.microsecond |> elem(0) |> div(1000) |> Integer.to_string() |> String.pad_leading(3, "0")
    Calendar.strftime(t, "%H:%M:%S.") <> ms
  end

  defp safe(fun) do
    {:ok, fun.()}
  rescue
    e -> {:error, Exception.message(e) |> String.slice(0, 600)}
  end

  defp to_message(e) when is_exception(e), do: Exception.message(e) |> String.slice(0, 600)
  defp to_message(e), do: inspect(e) |> String.slice(0, 600)

  defp start_repo(dialect, url) do
    {repo, opts} =
      case dialect do
        "postgresql" -> {Compat.PgRepo, [url: url]}
        "mysql" -> {Compat.MyRepo, [url: url]}
        "sqlite" -> {Compat.LiteRepo, [database: String.replace_prefix(url, "file:", "")]}
      end

    Logger.configure(level: :warning)
    {:ok, _} = Application.ensure_all_started(:ecto_sql)
    {:ok, _} = repo.start_link(Keyword.merge(opts, pool_size: 1, log: false))
    repo
  end
end

Compat.Harness.main(System.argv())
