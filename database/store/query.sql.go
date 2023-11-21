// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createAmenity = `-- name: CreateAmenity :one
INSERT INTO amenities (
  name, category, property_unit_id
) VALUES (
  $1, $2, $3
)
RETURNING id, name, provider, category, created_at, updated_at, property_unit_id
`

type CreateAmenityParams struct {
	Name           string        `json:"name"`
	Category       string        `json:"category"`
	PropertyUnitID uuid.NullUUID `json:"property_unit_id"`
}

func (q *Queries) CreateAmenity(ctx context.Context, arg CreateAmenityParams) (Amenity, error) {
	row := q.db.QueryRowContext(ctx, createAmenity, arg.Name, arg.Category, arg.PropertyUnitID)
	var i Amenity
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Provider,
		&i.Category,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
	)
	return i, err
}

const createCaretaker = `-- name: CreateCaretaker :one
INSERT INTO caretakers (
  first_name, last_name, phone, created_by
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, first_name, last_name, phone, verified, created_by, created_at, updated_at
`

type CreateCaretakerParams struct {
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Phone     string        `json:"phone"`
	CreatedBy uuid.NullUUID `json:"created_by"`
}

func (q *Queries) CreateCaretaker(ctx context.Context, arg CreateCaretakerParams) (Caretaker, error) {
	row := q.db.QueryRowContext(ctx, createCaretaker,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.CreatedBy,
	)
	var i Caretaker
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Verified,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createCaretakerAvatar = `-- name: CreateCaretakerAvatar :one
INSERT INTO uploads (
  upload, category, caretaker_id
) VALUES (
  $1, $2, $3
)
RETURNING id, upload, category, label, created_at, updated_at, property_unit_id, property_id, user_id, caretaker_id
`

type CreateCaretakerAvatarParams struct {
	Upload      string        `json:"upload"`
	Category    string        `json:"category"`
	CaretakerID uuid.NullUUID `json:"caretaker_id"`
}

func (q *Queries) CreateCaretakerAvatar(ctx context.Context, arg CreateCaretakerAvatarParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, createCaretakerAvatar, arg.Upload, arg.Category, arg.CaretakerID)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Upload,
		&i.Category,
		&i.Label,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.PropertyID,
		&i.UserID,
		&i.CaretakerID,
	)
	return i, err
}

const createInvoice = `-- name: CreateInvoice :one
INSERT INTO invoices (
  reference, phone, msid
) VALUES (
  $1, $2, $3
)
RETURNING id, msid, channel, currency, bank, auth_code, country_code, fees, amount, phone, status, reference, created_at, updated_at
`

type CreateInvoiceParams struct {
	Reference sql.NullString `json:"reference"`
	Phone     sql.NullString `json:"phone"`
	Msid      sql.NullString `json:"msid"`
}

func (q *Queries) CreateInvoice(ctx context.Context, arg CreateInvoiceParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, createInvoice, arg.Reference, arg.Phone, arg.Msid)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Msid,
		&i.Channel,
		&i.Currency,
		&i.Bank,
		&i.AuthCode,
		&i.CountryCode,
		&i.Fees,
		&i.Amount,
		&i.Phone,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProperty = `-- name: CreateProperty :one
INSERT INTO properties (
  name, type, created_by, caretaker_id, location
) VALUES (
  $1, $2, $3, $4, ST_GeomFromText($5::text)
)
RETURNING id, name, location, type, created_at, updated_at, created_by, caretaker_id
`

type CreatePropertyParams struct {
	Name        string        `json:"name"`
	Type        string        `json:"type"`
	CreatedBy   uuid.NullUUID `json:"created_by"`
	CaretakerID uuid.NullUUID `json:"caretaker_id"`
	Location    string        `json:"location"`
}

func (q *Queries) CreateProperty(ctx context.Context, arg CreatePropertyParams) (Property, error) {
	row := q.db.QueryRowContext(ctx, createProperty,
		arg.Name,
		arg.Type,
		arg.CreatedBy,
		arg.CaretakerID,
		arg.Location,
	)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.CaretakerID,
	)
	return i, err
}

const createPropertyThumbnail = `-- name: CreatePropertyThumbnail :one
INSERT INTO uploads (
  upload, category, property_id
) VALUES (
  $1, $2, $3
)
RETURNING id, upload, category, label, created_at, updated_at, property_unit_id, property_id, user_id, caretaker_id
`

type CreatePropertyThumbnailParams struct {
	Upload     string        `json:"upload"`
	Category   string        `json:"category"`
	PropertyID uuid.NullUUID `json:"property_id"`
}

func (q *Queries) CreatePropertyThumbnail(ctx context.Context, arg CreatePropertyThumbnailParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, createPropertyThumbnail, arg.Upload, arg.Category, arg.PropertyID)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Upload,
		&i.Category,
		&i.Label,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.PropertyID,
		&i.UserID,
		&i.CaretakerID,
	)
	return i, err
}

const createPropertyUnit = `-- name: CreatePropertyUnit :one
INSERT INTO property_units (
  property_id, bathrooms, name, type, price
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, type, state, price, bathrooms, created_at, updated_at, property_id
`

type CreatePropertyUnitParams struct {
	PropertyID uuid.NullUUID `json:"property_id"`
	Bathrooms  int32         `json:"bathrooms"`
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	Price      int32         `json:"price"`
}

func (q *Queries) CreatePropertyUnit(ctx context.Context, arg CreatePropertyUnitParams) (PropertyUnit, error) {
	row := q.db.QueryRowContext(ctx, createPropertyUnit,
		arg.PropertyID,
		arg.Bathrooms,
		arg.Name,
		arg.Type,
		arg.Price,
	)
	var i PropertyUnit
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.State,
		&i.Price,
		&i.Bathrooms,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyID,
	)
	return i, err
}

const createShootSchedule = `-- name: CreateShootSchedule :one
INSERT INTO shoots (
  shoot_date, property_id
) VALUES (
  $1, $2
)
RETURNING id, shoot_date, property_id, property_unit_id, status, caretaker_id, created_at, updated_at
`

type CreateShootScheduleParams struct {
	ShootDate  time.Time `json:"shoot_date"`
	PropertyID uuid.UUID `json:"property_id"`
}

func (q *Queries) CreateShootSchedule(ctx context.Context, arg CreateShootScheduleParams) (Shoot, error) {
	row := q.db.QueryRowContext(ctx, createShootSchedule, arg.ShootDate, arg.PropertyID)
	var i Shoot
	err := row.Scan(
		&i.ID,
		&i.ShootDate,
		&i.PropertyID,
		&i.PropertyUnitID,
		&i.Status,
		&i.CaretakerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTenant = `-- name: CreateTenant :one
INSERT INTO tenants (
  start_date, property_unit_id
) VALUES (
  $1, $2
)
RETURNING id, start_date, end_date, created_at, updated_at, property_unit_id, user_id
`

type CreateTenantParams struct {
	StartDate      time.Time     `json:"start_date"`
	PropertyUnitID uuid.NullUUID `json:"property_unit_id"`
}

func (q *Queries) CreateTenant(ctx context.Context, arg CreateTenantParams) (Tenant, error) {
	row := q.db.QueryRowContext(ctx, createTenant, arg.StartDate, arg.PropertyUnitID)
	var i Tenant
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.UserID,
	)
	return i, err
}

const createUnitBedroom = `-- name: CreateUnitBedroom :one
INSERT INTO bedrooms (
  property_unit_id, bedroom_number, en_suite, master
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, bedroom_number, en_suite, master, property_unit_id, created_at, updated_at
`

type CreateUnitBedroomParams struct {
	PropertyUnitID uuid.UUID `json:"property_unit_id"`
	BedroomNumber  int32     `json:"bedroom_number"`
	EnSuite        bool      `json:"en_suite"`
	Master         bool      `json:"master"`
}

func (q *Queries) CreateUnitBedroom(ctx context.Context, arg CreateUnitBedroomParams) (Bedroom, error) {
	row := q.db.QueryRowContext(ctx, createUnitBedroom,
		arg.PropertyUnitID,
		arg.BedroomNumber,
		arg.EnSuite,
		arg.Master,
	)
	var i Bedroom
	err := row.Scan(
		&i.ID,
		&i.BedroomNumber,
		&i.EnSuite,
		&i.Master,
		&i.PropertyUnitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUnitImage = `-- name: CreateUnitImage :one
INSERT INTO uploads (
  upload, category, label, property_unit_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, upload, category, label, created_at, updated_at, property_unit_id, property_id, user_id, caretaker_id
`

type CreateUnitImageParams struct {
	Upload         string         `json:"upload"`
	Category       string         `json:"category"`
	Label          sql.NullString `json:"label"`
	PropertyUnitID uuid.NullUUID  `json:"property_unit_id"`
}

func (q *Queries) CreateUnitImage(ctx context.Context, arg CreateUnitImageParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, createUnitImage,
		arg.Upload,
		arg.Category,
		arg.Label,
		arg.PropertyUnitID,
	)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Upload,
		&i.Category,
		&i.Label,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.PropertyID,
		&i.UserID,
		&i.CaretakerID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  phone
) VALUES (
  $1
)
RETURNING id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at
`

func (q *Queries) CreateUser(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUserAvatar = `-- name: CreateUserAvatar :one
INSERT INTO uploads (
  upload, category, user_id
) VALUES (
  $1, $2, $3
)
RETURNING id, upload, category, label, created_at, updated_at, property_unit_id, property_id, user_id, caretaker_id
`

type CreateUserAvatarParams struct {
	Upload   string        `json:"upload"`
	Category string        `json:"category"`
	UserID   uuid.NullUUID `json:"user_id"`
}

func (q *Queries) CreateUserAvatar(ctx context.Context, arg CreateUserAvatarParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, createUserAvatar, arg.Upload, arg.Category, arg.UserID)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Upload,
		&i.Category,
		&i.Label,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.PropertyID,
		&i.UserID,
		&i.CaretakerID,
	)
	return i, err
}

const findUserByPhone = `-- name: FindUserByPhone :one
SELECT id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at FROM users
WHERE phone = $1
`

func (q *Queries) FindUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCaretaker = `-- name: GetCaretaker :one
SELECT id, first_name, last_name, phone, verified, created_by, created_at, updated_at FROM caretakers
WHERE phone = $1
`

func (q *Queries) GetCaretaker(ctx context.Context, phone string) (Caretaker, error) {
	row := q.db.QueryRowContext(ctx, getCaretaker, phone)
	var i Caretaker
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Phone,
		&i.Verified,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCaretakerAvatar = `-- name: GetCaretakerAvatar :one
SELECT id, upload, category FROM uploads
WHERE caretaker_id = $1 AND category = $2
`

type GetCaretakerAvatarParams struct {
	CaretakerID uuid.NullUUID `json:"caretaker_id"`
	Category    string        `json:"category"`
}

type GetCaretakerAvatarRow struct {
	ID       uuid.UUID `json:"id"`
	Upload   string    `json:"upload"`
	Category string    `json:"category"`
}

func (q *Queries) GetCaretakerAvatar(ctx context.Context, arg GetCaretakerAvatarParams) (GetCaretakerAvatarRow, error) {
	row := q.db.QueryRowContext(ctx, getCaretakerAvatar, arg.CaretakerID, arg.Category)
	var i GetCaretakerAvatarRow
	err := row.Scan(&i.ID, &i.Upload, &i.Category)
	return i, err
}

const getProperty = `-- name: GetProperty :one
SELECT id, name, location, type, created_at, updated_at, created_by, caretaker_id FROM properties
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProperty(ctx context.Context, id uuid.UUID) (Property, error) {
	row := q.db.QueryRowContext(ctx, getProperty, id)
	var i Property
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Location,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.CaretakerID,
	)
	return i, err
}

const getPropertyThumbnail = `-- name: GetPropertyThumbnail :one
SELECT id, upload FROM uploads
WHERE property_id = $1 AND category = $2 LIMIT 1
`

type GetPropertyThumbnailParams struct {
	PropertyID uuid.NullUUID `json:"property_id"`
	Category   string        `json:"category"`
}

type GetPropertyThumbnailRow struct {
	ID     uuid.UUID `json:"id"`
	Upload string    `json:"upload"`
}

func (q *Queries) GetPropertyThumbnail(ctx context.Context, arg GetPropertyThumbnailParams) (GetPropertyThumbnailRow, error) {
	row := q.db.QueryRowContext(ctx, getPropertyThumbnail, arg.PropertyID, arg.Category)
	var i GetPropertyThumbnailRow
	err := row.Scan(&i.ID, &i.Upload)
	return i, err
}

const getPropertyUnits = `-- name: GetPropertyUnits :many
SELECT id, name, type, state, price, bathrooms, created_at, updated_at, property_id FROM property_units
WHERE property_id = $1
`

func (q *Queries) GetPropertyUnits(ctx context.Context, propertyID uuid.NullUUID) ([]PropertyUnit, error) {
	rows, err := q.db.QueryContext(ctx, getPropertyUnits, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PropertyUnit
	for rows.Next() {
		var i PropertyUnit
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.State,
			&i.Price,
			&i.Bathrooms,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitBedrooms = `-- name: GetUnitBedrooms :many
SELECT id, bedroom_number, en_suite, master, property_unit_id, created_at, updated_at FROM bedrooms
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitBedrooms(ctx context.Context, propertyUnitID uuid.UUID) ([]Bedroom, error) {
	rows, err := q.db.QueryContext(ctx, getUnitBedrooms, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bedroom
	for rows.Next() {
		var i Bedroom
		if err := rows.Scan(
			&i.ID,
			&i.BedroomNumber,
			&i.EnSuite,
			&i.Master,
			&i.PropertyUnitID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitImages = `-- name: GetUnitImages :many
SELECT id, upload, label FROM uploads
WHERE property_unit_id = $1 AND category = $2 LIMIT 1
`

type GetUnitImagesParams struct {
	PropertyUnitID uuid.NullUUID `json:"property_unit_id"`
	Category       string        `json:"category"`
}

type GetUnitImagesRow struct {
	ID     uuid.UUID      `json:"id"`
	Upload string         `json:"upload"`
	Label  sql.NullString `json:"label"`
}

func (q *Queries) GetUnitImages(ctx context.Context, arg GetUnitImagesParams) ([]GetUnitImagesRow, error) {
	rows, err := q.db.QueryContext(ctx, getUnitImages, arg.PropertyUnitID, arg.Category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUnitImagesRow
	for rows.Next() {
		var i GetUnitImagesRow
		if err := rows.Scan(&i.ID, &i.Upload, &i.Label); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnitTenancy = `-- name: GetUnitTenancy :many
SELECT id, start_date, end_date, created_at, updated_at, property_unit_id, user_id FROM tenants
WHERE property_unit_id = $1
`

func (q *Queries) GetUnitTenancy(ctx context.Context, propertyUnitID uuid.NullUUID) ([]Tenant, error) {
	rows, err := q.db.QueryContext(ctx, getUnitTenancy, propertyUnitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tenant
	for rows.Next() {
		var i Tenant
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PropertyUnitID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserAvatar = `-- name: GetUserAvatar :one
SELECT id, upload, category FROM uploads
WHERE user_id = $1 AND category = $2 LIMIT 1
`

type GetUserAvatarParams struct {
	UserID   uuid.NullUUID `json:"user_id"`
	Category string        `json:"category"`
}

type GetUserAvatarRow struct {
	ID       uuid.UUID `json:"id"`
	Upload   string    `json:"upload"`
	Category string    `json:"category"`
}

func (q *Queries) GetUserAvatar(ctx context.Context, arg GetUserAvatarParams) (GetUserAvatarRow, error) {
	row := q.db.QueryRowContext(ctx, getUserAvatar, arg.UserID, arg.Category)
	var i GetUserAvatarRow
	err := row.Scan(&i.ID, &i.Upload, &i.Category)
	return i, err
}

const mailingExists = `-- name: MailingExists :one
SELECT EXISTS(
  SELECT id, email, created_at, updated_at FROM mailings
  WHERE email = $1
)
`

func (q *Queries) MailingExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, mailingExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const occupiedUnitsCount = `-- name: OccupiedUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'occupied'
`

func (q *Queries) OccupiedUnitsCount(ctx context.Context, propertyID uuid.NullUUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, occupiedUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const propertiesCreatedBy = `-- name: PropertiesCreatedBy :many
SELECT id, name, location, type, created_at, updated_at, created_by, caretaker_id FROM properties
WHERE created_by = $1
`

func (q *Queries) PropertiesCreatedBy(ctx context.Context, createdBy uuid.NullUUID) ([]Property, error) {
	rows, err := q.db.QueryContext(ctx, propertiesCreatedBy, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Property
	for rows.Next() {
		var i Property
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Location,
			&i.Type,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.CaretakerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const propertyUnitsCount = `-- name: PropertyUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1
`

func (q *Queries) PropertyUnitsCount(ctx context.Context, propertyID uuid.NullUUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, propertyUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const saveMail = `-- name: SaveMail :one
INSERT INTO mailings (
  email
) VALUES (
  $1
)
RETURNING id, email, created_at, updated_at
`

func (q *Queries) SaveMail(ctx context.Context, email string) (Mailing, error) {
	row := q.db.QueryRowContext(ctx, saveMail, email)
	var i Mailing
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const trackSubscribeRetries = `-- name: TrackSubscribeRetries :one
UPDATE users SET subscribe_retries = $1
WHERE phone = $2
RETURNING id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at
`

type TrackSubscribeRetriesParams struct {
	SubscribeRetries int32  `json:"subscribe_retries"`
	Phone            string `json:"phone"`
}

func (q *Queries) TrackSubscribeRetries(ctx context.Context, arg TrackSubscribeRetriesParams) (User, error) {
	row := q.db.QueryRowContext(ctx, trackSubscribeRetries, arg.SubscribeRetries, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const unitAmenityCount = `-- name: UnitAmenityCount :one
SELECT COUNT(*) from amenities
WHERE property_unit_id = $1
`

func (q *Queries) UnitAmenityCount(ctx context.Context, propertyUnitID uuid.NullUUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, unitAmenityCount, propertyUnitID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateInvoiceForMpesa = `-- name: UpdateInvoiceForMpesa :one
UPDATE invoices
SET channel = $1, status = $2, amount = $3, currency = $4, bank = $5, auth_code = $6, country_code = $7, fees = $8, created_at = $9, updated_at = $10
WHERE reference = $11
RETURNING id, msid, channel, currency, bank, auth_code, country_code, fees, amount, phone, status, reference, created_at, updated_at
`

type UpdateInvoiceForMpesaParams struct {
	Channel     sql.NullString `json:"channel"`
	Status      string         `json:"status"`
	Amount      sql.NullString `json:"amount"`
	Currency    sql.NullString `json:"currency"`
	Bank        sql.NullString `json:"bank"`
	AuthCode    sql.NullString `json:"auth_code"`
	CountryCode sql.NullString `json:"country_code"`
	Fees        sql.NullString `json:"fees"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Reference   sql.NullString `json:"reference"`
}

func (q *Queries) UpdateInvoiceForMpesa(ctx context.Context, arg UpdateInvoiceForMpesaParams) (Invoice, error) {
	row := q.db.QueryRowContext(ctx, updateInvoiceForMpesa,
		arg.Channel,
		arg.Status,
		arg.Amount,
		arg.Currency,
		arg.Bank,
		arg.AuthCode,
		arg.CountryCode,
		arg.Fees,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Reference,
	)
	var i Invoice
	err := row.Scan(
		&i.ID,
		&i.Msid,
		&i.Channel,
		&i.Currency,
		&i.Bank,
		&i.AuthCode,
		&i.CountryCode,
		&i.Fees,
		&i.Amount,
		&i.Phone,
		&i.Status,
		&i.Reference,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateLandlord = `-- name: UpdateLandlord :one
UPDATE users
SET next_renewal = $1
WHERE phone = $2
RETURNING id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at
`

type UpdateLandlordParams struct {
	NextRenewal time.Time `json:"next_renewal"`
	Phone       string    `json:"phone"`
}

func (q *Queries) UpdateLandlord(ctx context.Context, arg UpdateLandlordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateLandlord, arg.NextRenewal, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUpload = `-- name: UpdateUpload :one
UPDATE uploads
SET upload = $1
WHERE id = $2
RETURNING id, upload, category, label, created_at, updated_at, property_unit_id, property_id, user_id, caretaker_id
`

type UpdateUploadParams struct {
	Upload string    `json:"upload"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) UpdateUpload(ctx context.Context, arg UpdateUploadParams) (Upload, error) {
	row := q.db.QueryRowContext(ctx, updateUpload, arg.Upload, arg.ID)
	var i Upload
	err := row.Scan(
		&i.ID,
		&i.Upload,
		&i.Category,
		&i.Label,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PropertyUnitID,
		&i.PropertyID,
		&i.UserID,
		&i.CaretakerID,
	)
	return i, err
}

const updateUserInfo = `-- name: UpdateUserInfo :one
UPDATE users
SET first_name = $1, last_name = $2
WHERE phone = $3
RETURNING id, first_name, last_name, subscribe_retries, next_renewal, phone, created_at, updated_at
`

type UpdateUserInfoParams struct {
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Phone     string         `json:"phone"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUserInfo, arg.FirstName, arg.LastName, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.SubscribeRetries,
		&i.NextRenewal,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const vacantUnitsCount = `-- name: VacantUnitsCount :one
SELECT COUNT(*) FROM property_units
WHERE property_id = $1 AND state = 'vacant'
`

func (q *Queries) VacantUnitsCount(ctx context.Context, propertyID uuid.NullUUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, vacantUnitsCount, propertyID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
