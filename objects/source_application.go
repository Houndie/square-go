package objects

type Product string

const (
	ProductSquarePOS         Product = "SQUARE_POS"
	ProductExternalAPI       Product = "EXTERNAL_API"
	ProductBilling           Product = "BILLING"
	ProductAppointments      Product = "APPOINTMENTS"
	ProductInvoices          Product = "INVOICES"
	ProductOnlineStore       Product = "ONLINE_STORE"
	ProductPayroll           Product = "PAYROLL"
	ProductDashboard         Product = "DASHBOARD"
	ProductItemLibraryImport Product = "ITEM_LIBRARY_IMPORT"
	ProductOther             Product = "OTHER"
)

type SourceApplication struct {
	ApplicationID string  `json:"application_id,omitempty"i`
	Name          string  `json:"name,omitempty"`
	Product       Product `json:"product,omitempty"`
}
