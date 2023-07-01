// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Bedroom struct {
	ID             string     `json:"id"`
	PropertyUnitID string     `json:"propertyUnitId"`
	BedroomNumber  int        `json:"bedroomNumber"`
	EnSuite        bool       `json:"enSuite"`
	Master         bool       `json:"master"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type Caretaker struct {
	ID             string     `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Phone          string     `json:"phone"`
	IDVerification string     `json:"idVerification"`
	CountryCode    string     `json:"countryCode"`
	Verified       bool       `json:"verified"`
	ShootsInCharge []*Shoot   `json:"shootsInCharge"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type CaretakerInput struct {
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Phone          string      `json:"phone"`
	CountryCode    CountryCode `json:"countryCode"`
	IDVerification string      `json:"idVerification"`
}

type HandshakeInput struct {
	Phone string `json:"phone"`
}

type NewUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
}

type OnboardUserInput struct {
	Email      string `json:"email"`
	Onboarding bool   `json:"onboarding"`
}

type PropertyUnit struct {
	ID         string     `json:"id"`
	Bedrooms   []*Bedroom `json:"bedrooms"`
	PropertyID string     `json:"propertyId"`
	Price      string     `json:"price"`
	Bathrooms  int        `json:"bathrooms"`
	Type       string     `json:"type"`
	Tenancy    []*Tenant  `json:"tenancy"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type PropertyUnitInput struct {
	PropertyID string `json:"propertyId"`
	Bathrooms  int    `json:"bathrooms"`
}

type SetupPropertyInput struct {
	Name         string          `json:"name"`
	Town         string          `json:"town"`
	PostalCode   string          `json:"postalCode"`
	PropertyType string          `json:"propertyType"`
	Caretaker    *CaretakerInput `json:"caretaker"`
	Units        []*UnitInput    `json:"units"`
	Shoot        *ShootInput     `json:"shoot"`
	Creator      string          `json:"creator"`
}

type Shoot struct {
	ID         string     `json:"id"`
	PropertyID string     `json:"propertyId"`
	Date       time.Time  `json:"date"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	Contact    *Caretaker `json:"contact"`
	ContactID  string     `json:"contactId"`
}

type ShootInput struct {
	Date          time.Time `json:"date"`
	ContactPerson string    `json:"contactPerson"`
}

type Status struct {
	Success string `json:"success"`
}

type TenancyInput struct {
	StartDate      time.Time  `json:"startDate"`
	EndDate        *time.Time `json:"endDate"`
	PropertyUnitID string     `json:"propertyUnitId"`
}

type Tenant struct {
	ID             string     `json:"id"`
	StartDate      time.Time  `json:"startDate"`
	EndDate        *time.Time `json:"endDate"`
	PropertyUnitID string     `json:"propertyUnitId"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type Token struct {
	Token string `json:"token"`
}

type Town struct {
	ID         string `json:"id"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
}

type UnitAmenityInput struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type UnitBedroomInput struct {
	PropertyUnitID *string `json:"propertyUnitId"`
	BedroomNumber  int     `json:"bedroomNumber"`
	EnSuite        bool    `json:"enSuite"`
	Master         bool    `json:"master"`
}

type UnitInput struct {
	Name      string              `json:"name"`
	Price     string              `json:"price"`
	Type      string              `json:"type"`
	Amenities []*UnitAmenityInput `json:"amenities"`
	Bedrooms  []*UnitBedroomInput `json:"bedrooms"`
	Baths     int                 `json:"baths"`
}

type UpdateUserInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Onboarding bool   `json:"onboarding"`
}

type VerificationInput struct {
	Phone       string      `json:"phone"`
	Email       *string     `json:"email"`
	CountryCode CountryCode `json:"countryCode"`
	VerifyCode  *string     `json:"verifyCode"`
}

type CountryCode string

const (
	CountryCodeKe CountryCode = "KE"
)

var AllCountryCode = []CountryCode{
	CountryCodeKe,
}

func (e CountryCode) IsValid() bool {
	switch e {
	case CountryCodeKe:
		return true
	}
	return false
}

func (e CountryCode) String() string {
	return string(e)
}

func (e *CountryCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CountryCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CountryCode", str)
	}
	return nil
}

func (e CountryCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
