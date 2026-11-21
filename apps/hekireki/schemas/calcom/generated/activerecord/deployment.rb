# id:: This is a single row table, so we use a fixed id
# theme:: @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme)
class Deployment < ApplicationRecord
  self.table_name = "Deployment"

  attribute :id, default: 1
end
