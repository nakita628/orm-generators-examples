# @zod.import(["import { ZAdminTagTypeSchema } from '@documenso/lib/types/admin-tags';"])
#
# type:: [AdminTagType] @zod.custom.use(ZAdminTagTypeSchema)
class AdminTag < ApplicationRecord
  self.table_name = "AdminTag"
  self.inheritance_column = nil

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :type, presence: true
  validates :description, presence: true

  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :admin_tags
  belongs_to :organisation, foreign_key: "organisationId", optional: true, inverse_of: :admin_tags
  belongs_to :created_by_user, class_name: "User", foreign_key: "createdByUserId", optional: true,
                               inverse_of: :created_admin_tags
  belongs_to :admin_rule, foreign_key: "adminRuleId", optional: true, inverse_of: :admin_tags
end
