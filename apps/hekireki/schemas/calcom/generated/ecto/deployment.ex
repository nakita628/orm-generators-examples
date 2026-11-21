defmodule Calcom.Deployment do
  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: integer(),
          logo: String.t() | nil,
          theme: map() | nil,
          license_key: String.t() | nil,
          signature_token_encrypted: String.t() | nil,
          agreed_license_at: DateTime.t() | nil
        }

  schema "Deployment" do
    # This is a single row table, so we use a fixed id
    field(:id, :integer, primary_key: true, default: 1)
    field(:logo, :string)
    # @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
    field(:theme, :map)
    field(:license_key, :string, source: :licenseKey)
    field(:signature_token_encrypted, :string, source: :signatureTokenEncrypted)
    field(:agreed_license_at, Calcom.PrismaDateTime, source: :agreedLicenseAt)
  end
end