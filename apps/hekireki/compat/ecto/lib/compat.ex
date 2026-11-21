defmodule Compat.PgRepo do
  use Ecto.Repo, otp_app: :compat, adapter: Ecto.Adapters.Postgres
end

defmodule Compat.MyRepo do
  use Ecto.Repo, otp_app: :compat, adapter: Ecto.Adapters.MyXQL
end

defmodule Compat.LiteRepo do
  use Ecto.Repo, otp_app: :compat, adapter: Ecto.Adapters.SQLite3
end
