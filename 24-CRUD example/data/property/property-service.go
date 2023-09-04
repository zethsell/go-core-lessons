package property

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type property struct {
	ID                           uint32  `json:"id"`
	VendorId                     uint32  `json:"vendor_id"`
	Name                         *string `json:"name"`
	Active                       uint8   `json:"active"`
	CreatedAt                    *string `json:"created_at"`
	UpdatedAt                    *string `json:"updated_at"`
	DeletedAt                    *string `json:"deleted_at"`
	Reference                    *string `json:"reference"`
	Slug                         *string `json:"slug"`
	Tag                          *string `json:"tag"`
	PropertyStatusId             *uint32 `json:"property_status_id"`
	Published                    uint8   `json:"published"`
	PropertyStateId              *uint32 `json:"property_state_id"`
	GovernmentRegulations        *string `json:"government_regulations"`
	RentalDaysAllowed            *uint32 `json:"rental_days_allowed"`
	RentalDays                   *uint32 `json:"rental_days"`
	MaximumRentalDaysAllowed     *uint32 `json:"maximum_rental_days_allowed"`
	GovRegistrationRef           *string `json:"gov_registration_ref"`
	CurrencyCode                 *string `json:"currency_code"`
	CurrencyId                   *uint32 `json:"currency_id"`
	ReadyToPublish               *uint32 `json:"ready_to_publish"`
	MaximumRentalDaysCalculation *uint32 `json:"maximum_rental_days_calculation"`
	Integration                  *string `json:"integration"`
	IsSas                        uint8   `json:"is_sas"`
	IsBreakz                     uint8   `json:"is_breakz"`
	Ical                         *string `json:"ical"`
	PublishedAt                  *string `json:"published_at"`
	ReadyToPublishAt             *string `json:"ready_to_publish_at"`
	Ics                          *string `json:"ics"`
	StepValidation               *string `json:"step_validation"`
	AgeRestriction               uint32  `json:"age_restriction"`
	AgeRestrictionActive         uint8   `json:"age_restriction_active"`
	ValidateGovernment           uint8   `json:"validate_government"`
	UnpublishedAt                *string `json:"unpublished_at"`
}

func InsertProperty(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to get request"))
		return
	}

	var property property

	if err = json.Unmarshal(request, &property); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO properties (name, active, vendor_id) values (? , ?,  ?)")
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer statement.Close()

	result, err := statement.Exec(property.Name, property.Active, property.VendorId)
	if err != nil {
		w.Write([]byte("Failed to exec statment!"))
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to get ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Property stored successfully! id: %d", id)))
}

func ListProperties(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM properties")
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer rows.Close()

	var properties []property
	for rows.Next() {
		var property property

		if err := rows.Scan(
			&property.ID,
			&property.VendorId,
			&property.Name,
			&property.Active,
			&property.CreatedAt,
			&property.UpdatedAt,
			&property.DeletedAt,
			&property.Reference,
			&property.Slug,
			&property.Tag,
			&property.PropertyStatusId,
			&property.Published,
			&property.PropertyStateId,
			&property.GovernmentRegulations,
			&property.RentalDaysAllowed,
			&property.RentalDays,
			&property.MaximumRentalDaysAllowed,
			&property.GovRegistrationRef,
			&property.CurrencyCode,
			&property.CurrencyId,
			&property.ReadyToPublish,
			&property.MaximumRentalDaysCalculation,
			&property.Integration,
			&property.IsSas,
			&property.IsBreakz,
			&property.Ical,
			&property.PublishedAt,
			&property.ReadyToPublishAt,
			&property.Ics,
			&property.StepValidation,
			&property.AgeRestriction,
			&property.AgeRestrictionActive,
			&property.ValidateGovernment,
			&property.UnpublishedAt,
		); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		properties = append(properties, property)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(properties); err != nil {
		w.Write([]byte("Failed to mount JSON"))
		return
	}
}

func ShowProperty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to get params"))
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM properties WHERE id = ?", ID)
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}

	var property property
	if row.Next() {
		if err := row.Scan(
			&property.ID,
			&property.VendorId,
			&property.Name,
			&property.Active,
			&property.CreatedAt,
			&property.UpdatedAt,
			&property.DeletedAt,
			&property.Reference,
			&property.Slug,
			&property.Tag,
			&property.PropertyStatusId,
			&property.Published,
			&property.PropertyStateId,
			&property.GovernmentRegulations,
			&property.RentalDaysAllowed,
			&property.RentalDays,
			&property.MaximumRentalDaysAllowed,
			&property.GovRegistrationRef,
			&property.CurrencyCode,
			&property.CurrencyId,
			&property.ReadyToPublish,
			&property.MaximumRentalDaysCalculation,
			&property.Integration,
			&property.IsSas,
			&property.IsBreakz,
			&property.Ical,
			&property.PublishedAt,
			&property.ReadyToPublishAt,
			&property.Ics,
			&property.StepValidation,
			&property.AgeRestriction,
			&property.AgeRestrictionActive,
			&property.ValidateGovernment,
			&property.UnpublishedAt,
		); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(property); err != nil {
		w.Write([]byte("Failed to mount JSON"))
		return
	}
}

func UpdateProperty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to get params"))
	}

	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to get request"))
		return
	}

	var property property

	if err = json.Unmarshal(request, &property); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE properties SET name= ?, active =?, vendor_id=? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(property.Name, property.Active, property.VendorId, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar"))
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func DeleteProperty(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to get params"))
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM properties WHERE id = ?")
	if err != nil {
		w.Write([]byte("Failed to connect"))
		return
	}

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao excluir"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
