/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint. If the block_identifier is populated, a historical balance query should be performed.
type AccountBalanceRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	AccountIdentifier *AccountIdentifier `json:"account_identifier"`
	BlockIdentifier *PartialBlockIdentifier `json:"block_identifier,omitempty"`
	// In some cases, the caller may not want to retrieve all available balances for an AccountIdentifier. If the currencies field is populated, only balances for the specified currencies will be returned. If not populated, all available balances will be returned.
	Currencies []Currency `json:"currencies,omitempty"`
}
