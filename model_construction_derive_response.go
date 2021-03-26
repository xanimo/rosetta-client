/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ConstructionDeriveResponse is returned by the `/construction/derive` endpoint.
type ConstructionDeriveResponse struct {
	// [DEPRECATED by `account_identifier` in `v1.4.4`] Address in network-specific format.
	Address string `json:"address,omitempty"`
	AccountIdentifier *AccountIdentifier `json:"account_identifier,omitempty"`
	Metadata *interface{} `json:"metadata,omitempty"`
}
