package asc

import (
	"fmt"
	"net/http"
)

// AppCategory defines model for AppCategory.
type AppCategory struct {
	Attributes *struct {
		Platforms *[]Platform `json:"platforms,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Parent *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"parent,omitempty"`
		Subcategories *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"subcategories,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppCategoriesResponse defines model for AppCategoriesResponse.
type AppCategoriesResponse struct {
	Data     []AppCategory      `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// AppCategoryResponse defines model for AppCategoryResponse.
type AppCategoryResponse struct {
	Data     AppCategory    `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// ListAppCategoriesQuery are query options for ListAppCategories
type ListAppCategoriesQuery struct {
	ExistsParent        []string `url:"exists[parent],omitempty"`
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	FilterPlatforms     []string `url:"filter[platforms],omitempty"`
	Include             []string `url:"include,omitempty"`
	Limit               int      `url:"limit,omitempty"`
	LimitSubcategories  []string `url:"limit[subcategories],omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// ListSubcategoriesForAppCategoryQuery are query options for ListSubcategoriesForAppCategory
type ListSubcategoriesForAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Limit               int      `url:"limit,omitempty"`
	Cursor              string   `url:"cursor,omitempty"`
}

// GetAppCategoryQuery are query options for GetAppCategory
type GetAppCategoryQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
	Include             []string `url:"include,omitempty"`
	LimitSubcategories  []string `url:"limit[subcategories],omitempty"`
}

// GetAppCategoryForAppInfoQuery are query options for GetAppCategoryForAppInfo
type GetAppCategoryForAppInfoQuery struct {
	FieldsAppCategories []string `url:"fields[appCategories],omitempty"`
}

// ListAppCategories lists all categories on the App Store, including the category and subcategory hierarchy.
func (s *AppsService) ListAppCategories(params *ListAppCategoriesQuery) (*AppCategoriesResponse, *http.Response, error) {
	res := new(AppCategoriesResponse)
	resp, err := s.client.get("appCategories", params, res)
	return res, resp, err
}

// ListSubcategoriesForAppCategory lists all App Store subcategories that belong to a specific category.
func (s *AppsService) ListSubcategoriesForAppCategory(id string, params *ListSubcategoriesForAppCategoryQuery) (*AppCategoriesResponse, *http.Response, error) {
	url := fmt.Sprintf("appCategories/%s/subcategories", id)
	res := new(AppCategoriesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppCategory gets a specific app category.
func (s *AppsService) GetAppCategory(id string, params *GetAppCategoryQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appCategories/%s", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetParentCategoryForAppCategory gets the App Store category to which a specific subcategory belongs.
func (s *AppsService) GetParentCategoryForAppCategory(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appCategories/%s/parent", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetPrimaryCategoryForAppInfo gets an app’s primary App Store category.
func (s *AppsService) GetPrimaryCategoryForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetSecondaryCategoryForAppInfo gets an app’s secondary App Store category.
func (s *AppsService) GetSecondaryCategoryForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondaryCategory", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s primary category.
func (s *AppsService) GetPrimarySubcategoryOneForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetPrimarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s primary category.
func (s *AppsService) GetPrimarySubcategoryTwoForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/primarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryOneForAppInfo gets the first App Store subcategory within an app’s secondary category.
func (s *AppsService) GetSecondarySubcategoryOneForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryOne", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetSecondarySubcategoryTwoForAppInfo gets the second App Store subcategory within an app’s secondary category.
func (s *AppsService) GetSecondarySubcategoryTwoForAppInfo(id string, params *GetAppCategoryForAppInfoQuery) (*AppCategoryResponse, *http.Response, error) {
	url := fmt.Sprintf("appInfos/%s/secondarySubcategoryTwo", id)
	res := new(AppCategoryResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}