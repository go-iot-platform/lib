package payment

type ServicePackage struct {
	// tableName   struct{} `pg:"service_package"`
	ServiceName string
	PackageID   int `pg:",pk"`
}
