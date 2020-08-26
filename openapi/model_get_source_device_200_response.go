/*
 * untitled API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 513
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetSourceDevice200Response struct for GetSourceDevice200Response
type GetSourceDevice200Response struct {
	Account GetSourceDevice200ResponseAccount `json:"account"`
	Config GetSourceDevice200ResponseConfig `json:"config"`
	Created string `json:"created"`
	Enabled bool `json:"enabled"`
	FcmToken string `json:"fcm_token"`
	Id string `json:"id"`
	InstallId string `json:"install_id"`
	Key string `json:"key"`
	Locale string `json:"locale"`
	Model string `json:"model"`
	Name string `json:"name"`
	Place float32 `json:"place"`
	Tos string `json:"tos"`
	Type string `json:"type"`
	Updated string `json:"updated"`
	WaitlistEnabled bool `json:"waitlist_enabled"`
	WarpEnabled bool `json:"warp_enabled"`
}
