# timestamp:: @zod.custom(z.coerce.date())
class TaskLog < ApplicationRecord
  self.table_name = "task_log"

  def self.instance_method_already_implemented?(method_name)
    dangerous_attribute_method?(method_name) || super
  end
  private_class_method :instance_method_already_implemented?

  def read_attribute_for_serialization(name)
    if has_attribute?(name) && self.class.dangerous_attribute_method?(name)
      _read_attribute(name.to_s)
    else
      super
    end
  end

  attribute :id, default: -> { SecureRandom.uuid }
  alias_attribute :logger_value, :logger

  validates :level, presence: true
  validates :logger_value, presence: true
  validates :message, presence: true
  validates :sync_id, presence: true
  validates :task_id, presence: true
end
