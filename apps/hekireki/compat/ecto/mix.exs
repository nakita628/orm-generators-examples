defmodule Compat.MixProject do
  use Mix.Project

  # Runs the compat plan through the Ecto schemas hekireki wrote. run.mjs copies the generated
  # lib/*.ex beside lib/compat.ex and runs `mix run harness.exs`.
  def project do
    [
      app: :compat,
      version: "0.0.0",
      elixir: "~> 1.14",
      deps: deps()
    ]
  end

  def application, do: [extra_applications: [:logger]]

  defp deps do
    [
      {:ecto, "== 3.14.1"},
      {:ecto_sql, "== 3.14.0"},
      {:ecto_ulid_next, "== 1.0.2"},
      {:jason, "~> 1.4"},
      {:postgrex, "~> 0.20"},
      {:myxql, "~> 0.7"},
      {:ecto_sqlite3, "~> 0.17"}
    ]
  end
end
