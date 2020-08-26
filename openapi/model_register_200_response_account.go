/*
 * untitled API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 513
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Register200ResponseAccount struct for Register200ResponseAccount
type Register200ResponseAccount struct {
	AccountType string `json:"account_type"`
	Created string `json:"created"`
	Id string `json:"id"`
	License string `json:"license"`
	PremiumData float32 `json:"premium_data"`
	Quota float32 `json:"quota"`
	ReferralCount float32 `json:"referral_count"`
	ReferralRenewalCountdown float32 `json:"referral_renewal_countdown"`
	Role string `json:"role"`
	Updated string `json:"updated"`
	Usage float32 `json:"usage"`
	WarpPlus bool `json:"warp_plus"`
}
