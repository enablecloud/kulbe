/* 
 * Fusion Creator API
 *
 * .
 *
 * OpenAPI spec version: 0.1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ApplicationCreation struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name"`

	Services []ServiceName `json:"services,omitempty"`
}
