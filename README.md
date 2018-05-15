# RESTFUL API for a payments system

This repository contains a small API REST for a payments system. It contains the following actions:

 - List all payments
 - Show one payment.
 - Create one payment.
 - Modify a payment.
 - Delete a payment.
 
## Prerrequisites

 - Golang installed.
 - Mux library.
 - Setup of the global variable GOPATH.
 - Postman or any other REST client for testing purposes.

## Struct of the JSON
The JSON is based on the documentation provided. A formatted example is provided to help run quick tests.

`{"type":"Payment",
 "id":"4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
 "version":0,
 "organisation_id":"743d5b63-8e6f-432e-a8fa-c5d8d2ee5fcb",
 "attributes":{
    "amount":"100.21",
    "beneficiary_party":{
       "account_name":"W Owens",
       "account_number":"31926819",
       "account_number_code":"BBAN",
       "account_type":0,
       "address":"1 The Beneficiary Localtown SE2",
       "bank_id":"403000",
       "bank_id_code":"GBDSC",
       "name":"Wilfred Jeremiah Owens"},
    "charges_information":{
       "bearer_code":"SHAR",
       "sender_charges":[
          {"amount":"5.00",
          "currency":"GBP"},
          {"amount":"10.00",
          "currency":"USD"}],
        "receiver_charges_amount":"1.00",
        "receiver_charges_currency":"USD"},
    "currency":"GBP",
    "debtor_party":{
       "account_name":"EJ Brown Black",
       "account_number":"GB29XABC10161234567801",
       "account_number_code":"IBAN",
       "address":"10 Debtor Crescent Sourcetown NE1",
       "bank_id":"203301",
       "bank_id_code":"GBDSC",
       "name":"Emelia Jane Brown"},
    "end_to_end_reference":"Wil piano Jan",
    "fx":{
       "contract_reference":"FX123",
       "exchange_rate":"2.00000",
       "original_amount":"200.42",
       "original_currency":"USD"},
    "numeric_reference":"1002001",
    "payment_id":"123456789012345678",
    "payment_purpose":"Paying for goods/services",
    "payment_scheme":"FPS",
    "payment_type":"Credit",
    "processing_date":"2017-01-18",
    "reference":"Payment for Em's piano lessons",
    "scheme_payment_sub_type":"InternetBanking",
    "scheme_payment_type":"ImmediatePayment",
    "sponsor_party":{
       "account_number":"56781234",
       "bank_id":"123123",
       "bank_id_code":"GBDSC"}
 }
}
`
## Available actions
A list of all the available actions and how they work.

### Show all payments
To get all the payments on the system. Should show all the populated payments on the system.

`GET localhost:8000/payments`

### Show one payment
To get one specific payment, must show the only corresponding payment with the ID specified. It should contain the same parameters as in the struct of the JSON showed before.

`GET localhost:8000/payment/{id}`

### Create payment
The way to create a payment is the following.

`POST localhost:8000/payment/create/{id}`

The body of the petition should include the struct of the example provided previously. The `create` tag is included to clarify the action.

### Modify payment
The way to modify a payment is the following.

`PATCH localhost:8000/payment/modify/{id}`

The body of the petition should include the struct of the example provided previously with the needed updates. The `modify` tag is included to clarify the action.

### Delete payment
The way to delete a payment is the following.

`DELETE localhost:8000/payment/{id}`

After the action, if a GET payments is performed the deleted payment should not appear.
## Tests
### Create payment
Run the code and test on Postman (or any other platform) the option of creating a payment. 

`POST localhost:8000/payment/create/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43`

On the body use the template provided previously. The output should return the new payment created.

### Show payment
On postman run 

`GET localhost:8000/payment/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43`

It should return the previous data created

### Modify payment
 On postman run 

`PATCH localhost:8000/payment/edit/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43`

Modify the previous body used to create a payment. When run again show payment it should return the modified payment.

### Show payments
Create more payments as showed in the test of create payment, after it run

`GET localhost:8000/payments`

It will show all the created payments.

### Delete payment
To delete a payment run on Postman the following.

`DELETE localhost:8000/payment/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43`

If later is run a get payment with the same Id or get payments, the deleted payment shouldn't appear.

## How to run
First is needed to add the Mux library 
`go get -u github.com/gorilla/mux`
Export the variable GOPATH pointing to the project.
Run the command  `go build && ./GoREST` by default it will work on port 8000