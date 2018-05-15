package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// TYPES FOR PAYMENT
type Payment struct {
	Type              string        `json:"type,omitempty"`
    ID                string        `json:"id,omitempty"`
    Version           int           `json:"version,omitempty"`
    OrganisationID    string        `json:"organisation_id,omitempty"`
    Attributes        *Attributes   `json:"attributes,omitempty"`
}

type Attributes struct {
    Amount               string                `json:"amount,omitempty"`
    BeneficiaryParty     *Beneficiary_party    `json:"beneficiary_party,omitempty"`
    ChargesInformation   *Charges_information  `json:"charges_information,omitempty"`
    Currency             string                `json:"currency,omitempty"`
    DebtorParty          *Debtor_party         `json:"debtor_party,omitempty"`
    EndToEndRef          string                `json:"end_to_end_reference,omitempty"`
    Fx                   *Fx                   `json:"fx,omitempty"`
    NumericRef           string                `json:"numeric_reference,omitempty"`
    PaymentId            string                `json:"payment_id,omitempty"`
    PaymentPurpose       string                `json:"payment_purpose,omitempty"`
    PaymentScheme        string                `json:"payment_scheme,omitempty"`
    PaymentType          string                `json:"payment_type,omitempty"`
    ProcessingDate       string                `json:"processing_date,omitempty"`
    Reference            string                `json:"reference,omitempty"`
    SchemePaymentSubType string                `json:"scheme_payment_sub_type,omitempty"`
    SchemePaymentType    string                `json:"scheme_payment_type,omitempty"`
    SponsorParty         *Sponsor_party        `json:"sponsor_party,omitempty"`
}

type Beneficiary_party struct {
    AccountName        string   `json:"account_name,omitempty"`
    AccountNumber      string   `json:"account_number,omitempty"`
    AccountNumberCode  string   `json:"account_number_code,omitempty"`
    AccountType        int      `json:"account_type,omitempty"`
    Address            string   `json:"address,omitempty"`
    BankId             string   `json:"bank_id,omitempty"`
    BankIdCode         string   `json:"bank_id_code,omitempty"`
    Name               string   `json:"name,omitempty"`
}

type Charges_information struct {
    BearerCode                string            `json:"bearer_code,omitempty"`
    SenderCharges             []Sender_charges  `json:"sender_charges,omitempty"`
    RecieverChargerAmount     string            `json:"receiver_charger_amount,omitempty"`
    ReceiverChargerCurrency   string            `json:"receiver_charger_currency,omitempty"`    
}

type Sender_charges struct {
    Amount    string  `json:"amount,omitempty"`
    Currency  string  `json:"currency,omitempty"`
}

type Debtor_party struct {
    AccountName       string  `json:"account_name,omitempty"`
    AccountNumber     string  `json:"account_number,omitempty"`
    AccountNumberCode string  `json:"account_number_code,omitempty"`
    Address           string  `json:"address,omitempty"`
    BankId            string  `json:"bank_id,omitempty"`
    BankIdCode        string  `json:"bank_id_code,omitempty"`
    Name              string  `json:"name,omitempty"`
}

type Fx struct {
    ContractReference  string  `json:"contract_reference,omitempty"`
    ExchangeRate       string  `json:"exchange_rate,omitempty"`
    OriginalAmount     string  `json:"original_amount,omitempty"`
    OriginalCurrency   string  `json:"original_currency,omitempty"`
}

type Sponsor_party struct {
    AccountNumber   string  `json:"account_number,omitempty"`
    BankId          string  `json:"bank_id,omitempty"`
    BankIdCode      string  `json:"bank_id_code,omitempty"`
}


var payments []Payment

// Display all the transactions
func GetPayments(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(payments)
}

// Display a single data
func GetPayment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range payments {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Payment{})
}

// create a new item
func CreatePayment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var payment Payment
    _ = json.NewDecoder(r.Body).Decode(&payment)
    //payments.Type = params["type"]
    payment.ID = params["id"]
    payments = append(payments, payment)
    json.NewEncoder(w).Encode(payments)
}

// Delete an item
func DeletePayment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range payments {
        if item.ID == params["id"] {
            payments = append(payments[:index], payments[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(payments)
    }
}

// Modify a payment
func ModifyPayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    for index, item := range payments {
        if item.ID == params["id"] {
        	//DELETE OLD VERSION
            payments = append(payments[:index], payments[index+1:]...)
            //CREATE "NEW" UPDATED
            var payment Payment
            _ = json.NewDecoder(r.Body).Decode(&payment)
            //payments.Type = params["type"]
            payment.ID = params["id"]
            payments = append(payments, payment)
            break
        }
        json.NewEncoder(w).Encode(payments)
    }
}


// main function to boot up everything
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/payments", GetPayments).Methods("GET")
    router.HandleFunc("/payment/{id}", GetPayment).Methods("GET")
    router.HandleFunc("/payment/create/{id}", CreatePayment).Methods("POST")
    router.HandleFunc("/payment/{id}", DeletePayment).Methods("DELETE")
    router.HandleFunc("/payment/edit/{id}", ModifyPayment).Methods("PATCH")
    log.Fatal(http.ListenAndServe(":8000", router))
}
