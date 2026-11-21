defmodule Formbricks.ResponseQuotaLink do
  @moduledoc """
  Junction table linking responses to quotas.
  Tracks which responses counted towards specific quotas and their status.

  @property response - The response that counted towards the quota
  @property quota - The quota this response counted towards
  @property status - Whether the response was screened in or out
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          response_id: String.t(),
          quota_id: String.t(),
          status: atom(),
          response: Formbricks.Response.t() | nil,
          quota: Formbricks.SurveyQuota.t() | nil
        }

  schema "ResponseQuotaLink" do
    field(:status, Ecto.Enum, values: [:screenedIn, :screenedOut])
    field(:response_id, :string, primary_key: true, source: :responseId)
    field(:quota_id, :string, primary_key: true, source: :quotaId)
    belongs_to(:response, Formbricks.Response, foreign_key: :response_id, define_field: false, type: :string)
    belongs_to(:quota, Formbricks.SurveyQuota, foreign_key: :quota_id, define_field: false, type: :string)
  end
end