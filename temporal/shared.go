package temporal

type PresentParams struct{
	WorkflowID string
	PendingID string
}
type AuthorizeParams struct{
    ID int
    Amount int
}
type Account struct {
	ID int
	Available int
	Reserved int
}
type AuthorizeResponse struct {
	WorkflowID string
}