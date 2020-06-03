# Mobile Wallet 2020
Mobile wallet for everyone

## Usage

### Install

    git clone https://github.com/gmbanaag/wallet2020
    make app 

### Run

    Before you run make sure you had updated the configuration at .env

    You may run the binary via:
    ./wallet2020

## API

    Documentation about the API may be found at https://gmbanaag.github.io/wallet2020/.

## Running with Docker

    You may also build a container to be upload to a docker repository:

    make build-docker

## Testing with dummy data

    I included dummy data to test the service. Please load db/dummy.sql to prepopulate the database

    For authentication, it uses bearer tokens but at the moment there's no specific service it connects to. You may indicate your own provider and please update the validation for the payload response of the token introspection endpoint. 

    The current validation consumes the following response payload:
    ```sh
    {
        "access_token": "",
        "client_id": "",
        "user_id": "",
        "scope": "",
        "expires_in": "" 
    }
    ```

    To test even without connecting to a valid authentication service you may use the following tokens, I had hardcoded.

    admin user: cec0482b1b77d46ab7f13b114e79ae3b3c01286d
    default user: ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b

    I had added sample requests below.

## Sample requests
    Transfer from source wallet to destination wallet:

    ```sh
    curl --location --request POST 'localhost:3000/v1/transfer' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```

    Get all current user's wallet:

    ```sh
    curl --location --request GET 'localhost:3000/v1/wallets' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```

    Get a specific wallet:

    ```sh
    curl --location --request GET 'localhost:3000/v1/wallets/ff7cc44a-b949-413c-9c75-6f34a5699915' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```

    Admin can get all available wallets:

    ```sh
    curl --location --request GET 'localhost:3000/v1/admin/wallets' \
    --header 'Authorization: Bearer cec0482b1b77d46ab7f13b114e79ae3b3c01286d' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```
    
    Admin can get all transactions:

    ```sh
    curl --location --request GET 'localhost:3000/v1/admin/transactions' \
    --header 'Authorization: Bearer cec0482b1b77d46ab7f13b114e79ae3b3c01286d' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```

    Non-admin trying to access an admin service:

    ```sh
    curl --location --request GET 'localhost:3000/v1/admin/transactions' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ``` 

    Getting current user's sent transfers:

    ```sh
    curl --location --request GET 'localhost:3000/v1/transactions/sent' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```

    Getting transfer received (This would result to HTTP 404 because user didnt received anything):
    
    ```sh
    curl --location --request GET 'localhost:3000/v1/transactions/received' \
    --header 'Authorization: Bearer ed405dcb8903bb7674dc7fbabebeeae8ebd8d30b' \
    --header 'Content-Type: application/json' \
    --data-raw '{"source_wallet_id":"ff7cc44a-b949-413c-9c75-6f34a5699915",
    "destination_wallet_id":"fa63af6e-3442-4f5f-9fb3-b6a33e3b9c9d",
    "amount": 10,
    "message":"here ya go"
    }'
    ```