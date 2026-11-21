package calcom.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

object DeploymentTable : IdTable<Int>("\"Deployment\"") {
    /** This is a single row table, so we use a fixed id */
    override val id = integer("id").default(1).entityId()
    val logo = text("logo").nullable()

    /** @zod.import(["import { DeploymentTheme } from '../../zod-utils'"]).custom.use(DeploymentTheme) */
    val theme = jsonb<String>("theme", { it }, { it }).nullable()
    val licenseKey = text("licenseKey").nullable()
    val signatureTokenEncrypted = text("signatureTokenEncrypted").nullable()
    val agreedLicenseAt = pgTimestamp("agreedLicenseAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Deployment_pkey\"")
}
