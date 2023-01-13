/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Microvisor
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ListAppResponseMeta struct for ListAppResponseMeta
type ListAppResponseMeta struct {
	FirstPageUrl    string  `json:"first_page_url,omitempty"`
	NextPageUrl     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page,omitempty"`
	PageSize        int     `json:"page_size,omitempty"`
	PreviousPageUrl *string `json:"previous_page_url,omitempty"`
	Url             string  `json:"url,omitempty"`
	Key             string  `json:"key,omitempty"`
}
