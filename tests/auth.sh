#!/bin/bash

check_response () {
    echo "$2"
    # Extract the body and status code from the response
    body=$(echo "$1" | sed '$d')
    status_code=$(echo "$1" | tail -n1)

    # Check if the status code is 200
    if [ "$status_code" -ne 200 ]; then
        echo "Error: request failed with status code $status_code."
        echo "$body"
        exit 1
    else
        echo "$body"
        echo "$status_code"
        echo
    fi
}

check_token () {
    # Extract the token from the JSON response
    body=$(echo "$1" | sed '$d')
    token=$(echo "$body" | jq -r '.token')

    # Check if the token is not empty
    if [ -z "$token" ]; then
        echo "Error: token not found in the response."
        exit 1
    fi
}

udid="testudid"

# Register
response1=$(curl -s -w "\n%{http_code}" \
    -A "Buddy Tracker" \
    --header "Content-Type: application/json" \
    -X POST \
    --data "{\"udid\":\"$udid\",\"name\":\"Test\"}" \
    http://localhost:3001/api/v1/auth/register
)

check_response "$response1" "register"
check_token "$response1"

# Login
response2=$(curl -s -w "\n%{http_code}\n" \
    -A "Buddy Tracker" \
    --header "Content-Type: application/json" \
    -X POST \
    --data "{\"udid\":\"$udid\"}" \
    http://localhost:3001/api/v1/auth
)
check_response "$response2" "login"
check_token "$response2"

# Test Login on /buddies
response3=$(curl -s -w "\n%{http_code}\n" \
    -A "Buddy Tracker" \
    --header "Authorization: Bearer $token" \
    -X GET \
    http://localhost:3001/api/v1/buddies
)
check_response "$response3" "list buddies"
