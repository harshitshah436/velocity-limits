# velocity-limits
 
This project provides an application to process banking transactions considering velocity limits.

## Overview

### Application functions

- High-level functionality is to accept or decline attempts to load funds into customers' accounts in real-time. More details can be found at [ProblemDetails](docs/ProblemDetails.md).
- Application reads the input.txt file which contains transactions to load funds and creates transactions array to store them. It assumes the input and output file path as the project root directory.
- Then, it attempts to load these transactions. If the transaction is a duplicate (determined by the same load ID and customer ID), it ignores all the following transactions. It validates each transaction and reset velocity limits if daily/weekly limits don't apply for the transaction date. Later, it processes the transaction and stores updated customer account details into the local storage.
- Once the transaction is processed (where it's approved/rejected), a response array will be created with accepted/rejected information. This array will be marshaled into the JSON object and written into the output.txt file.

### Technologies used

- [The Go Programming Language](https://golang.org)

## Getting Started

### Installation

#### Prerequisites
- `go`
    - [Install Go](https://golang.org/doc/install)
    - Verify using the command: `go version`
    - (Application tested on Mac OS)
        ```
        $ go version
            go version go1.16.3 darwin/amd64
        ```

#### Start application (Steps to Run)

In the project **root** directory, run:
```
go mod init velocity-limits
go mod tidy
cd cmd/velocity-limits/
go run .
```

### Installation using Docker (Cloud-native)

#### Prerequisites

- `docker` & `docker-compose`
    - Verify using the commands: `docker -v` & `docker-compose -v`
        ```
        $ docker -v
            Docker version 20.10.5, build 55c4c88
        $ docker-compose -v
            docker-compose version 1.29.0, build 07737305
        ```

#### Start application (Steps to Run)

In the project **root** directory, run:
```
docker-compose up
```

#### Stop application

```
docker-compose down
```

#### Restart containers with the new code

```
docker-compose up --build
```

## Unit tests

Using `testing` package, created unit tests for the application.

#### Run tests (Steps to Run)

In the project **root** directory, run:
```
$ cd test
$ go clean -testcache
$ go test -v ./...
```

#### Test Output (Verbose):
```
=== RUN   TestNewCustomerAccount
=== RUN   TestNewCustomerAccount/returns_a_new_customer_account
--- PASS: TestNewCustomerAccount (0.00s)
    --- PASS: TestNewCustomerAccount/returns_a_new_customer_account (0.00s)
=== RUN   TestNewDailyLimit
=== RUN   TestNewDailyLimit/returns_a_new_velocity_limits_per_day
--- PASS: TestNewDailyLimit (0.00s)
    --- PASS: TestNewDailyLimit/returns_a_new_velocity_limits_per_day (0.00s)
=== RUN   TestNewWeeklyLimit
=== RUN   TestNewWeeklyLimit/returns_a_new_velocity_limits_per_week
--- PASS: TestNewWeeklyLimit (0.00s)
    --- PASS: TestNewWeeklyLimit/returns_a_new_velocity_limits_per_week (0.00s)
=== RUN   TestValidateDailyLimit
=== RUN   TestValidateDailyLimit/returns_true_when_loading_below_max_load_limit_per_day
=== RUN   TestValidateDailyLimit/returns_true_when_loading_exactly_same_max_load_limit_per_day
=== RUN   TestValidateDailyLimit/returns_false_when_loading_more_than_allowed_max_load_limit_per_day
--- PASS: TestValidateDailyLimit (0.00s)
    --- PASS: TestValidateDailyLimit/returns_true_when_loading_below_max_load_limit_per_day (0.00s)
    --- PASS: TestValidateDailyLimit/returns_true_when_loading_exactly_same_max_load_limit_per_day (0.00s)
    --- PASS: TestValidateDailyLimit/returns_false_when_loading_more_than_allowed_max_load_limit_per_day (0.00s)
=== RUN   TestValidateWeeklyLimit
=== RUN   TestValidateWeeklyLimit/returns_true_when_loading_below_max_load_limit_per_week
=== RUN   TestValidateWeeklyLimit/returns_true_when_loading_exactly_same_max_load_limit_per_week
=== RUN   TestValidateWeeklyLimit/returns_false_when_loading_more_than_allowed_max_load_limit_per_week
--- PASS: TestValidateWeeklyLimit (0.00s)
    --- PASS: TestValidateWeeklyLimit/returns_true_when_loading_below_max_load_limit_per_week (0.00s)
    --- PASS: TestValidateWeeklyLimit/returns_true_when_loading_exactly_same_max_load_limit_per_week (0.00s)
    --- PASS: TestValidateWeeklyLimit/returns_false_when_loading_more_than_allowed_max_load_limit_per_week (0.00s)
=== RUN   TestDailyUpdateLimits
=== RUN   TestDailyUpdateLimits/updates_allocated_daily_limit_and_reflecting_amount_will_be_reduced
--- PASS: TestDailyUpdateLimits (0.00s)
    --- PASS: TestDailyUpdateLimits/updates_allocated_daily_limit_and_reflecting_amount_will_be_reduced (0.00s)
=== RUN   TestWeeklyUpdateLimits
=== RUN   TestWeeklyUpdateLimits/updates_allocated_weekly_limit_and_reflecting_amount_will_be_reduced
--- PASS: TestWeeklyUpdateLimits (0.00s)
    --- PASS: TestWeeklyUpdateLimits/updates_allocated_weekly_limit_and_reflecting_amount_will_be_reduced (0.00s)
=== RUN   TestResetLimits
=== RUN   TestResetLimits/should_not_reset_velocity_limits_if_load_time_is_within_daily/weekly_limits
=== RUN   TestResetLimits/should_reset_velocity_limits_if_load_time_is_before_current_day/week
--- PASS: TestResetLimits (0.00s)
    --- PASS: TestResetLimits/should_not_reset_velocity_limits_if_load_time_is_within_daily/weekly_limits (0.00s)
    --- PASS: TestResetLimits/should_reset_velocity_limits_if_load_time_is_before_current_day/week (0.00s)
=== RUN   TestLoadFunds
=== RUN   TestLoadFunds/should_return_true_when_max_load_per_day,_max_load_per_week_and_max_load_limits_are_not_reached
=== RUN   TestLoadFunds/should_return_false_when_max_load_per_day,_max_load_per_week_and_max_load_limits_are_reached
--- PASS: TestLoadFunds (0.00s)
    --- PASS: TestLoadFunds/should_return_true_when_max_load_per_day,_max_load_per_week_and_max_load_limits_are_not_reached (0.00s)
    --- PASS: TestLoadFunds/should_return_false_when_max_load_per_day,_max_load_per_week_and_max_load_limits_are_reached (0.00s)
=== RUN   TestNewResponse
=== RUN   TestNewResponse/should_return_a_new_response_struct
--- PASS: TestNewResponse (0.00s)
    --- PASS: TestNewResponse/should_return_a_new_response_struct (0.00s)
=== RUN   TestGetParsedAmount
=== RUN   TestGetParsedAmount/parses_input_load_amount_correcly
--- PASS: TestGetParsedAmount (0.00s)
    --- PASS: TestGetParsedAmount/parses_input_load_amount_correcly (0.00s)
PASS
ok  	velocity-limits/test/internal/models	0.221s
=== RUN   TestGetTransactionsFromInputFile
=== RUN   TestGetTransactionsFromInputFile/should_read_the_file_and_get_transactions
--- PASS: TestGetTransactionsFromInputFile (0.00s)
    --- PASS: TestGetTransactionsFromInputFile/should_read_the_file_and_get_transactions (0.00s)
=== RUN   TestValidateAndProcessTransaction
=== RUN   TestValidateAndProcessTransaction/should_validate_transactions_and_process_them_to_create_responses
2021/05/12 00:38:01 Ignoring a duplicate transaction: &{ID:6928 CustomerID:562 Amount:$3164.98 Time:2000-01-30 05:37:32 +0000 UTC}
--- PASS: TestValidateAndProcessTransaction (0.00s)
    --- PASS: TestValidateAndProcessTransaction/should_validate_transactions_and_process_them_to_create_responses (0.00s)
=== RUN   TestWriteResponsesToOutputFile
=== RUN   TestWriteResponsesToOutputFile/should_write_responses_to_the_output_file
--- PASS: TestWriteResponsesToOutputFile (0.00s)
    --- PASS: TestWriteResponsesToOutputFile/should_write_responses_to_the_output_file (0.00s)
PASS
ok  	velocity-limits/test/internal/service	0.450s
=== RUN   TestNewStorage
=== RUN   TestNewStorage/should_return_a_new_storage
--- PASS: TestNewStorage (0.00s)
    --- PASS: TestNewStorage/should_return_a_new_storage (0.00s)
=== RUN   TestStorageFunctions
=== RUN   TestStorageFunctions/should_get_a_nil_when_no_account_is_added_to_the_storage
=== RUN   TestStorageFunctions/should_add_an_account_to_the_storage
=== RUN   TestStorageFunctions/should_get_an_added_account_from_the_storage
=== RUN   TestStorageFunctions/should_add_a_tranasaction_to_the_storage_struct
=== RUN   TestStorageFunctions/should_return_true_when_adding_a_duplicate_transaction
=== RUN   TestStorageFunctions/should_return_false_when_adding_a_new_transaction
--- PASS: TestStorageFunctions (0.00s)
    --- PASS: TestStorageFunctions/should_get_a_nil_when_no_account_is_added_to_the_storage (0.00s)
    --- PASS: TestStorageFunctions/should_add_an_account_to_the_storage (0.00s)
    --- PASS: TestStorageFunctions/should_get_an_added_account_from_the_storage (0.00s)
    --- PASS: TestStorageFunctions/should_add_a_tranasaction_to_the_storage_struct (0.00s)
    --- PASS: TestStorageFunctions/should_return_true_when_adding_a_duplicate_transaction (0.00s)
    --- PASS: TestStorageFunctions/should_return_false_when_adding_a_new_transaction (0.00s)
PASS
ok  	velocity-limits/test/internal/storage	0.340s
```

## Resources

* [A tour of Go](https://tour.golang.org/list)
* [Go project layout standards](https://github.com/golang-standards/project-layout)
* [Go logging](https://www.honeybadger.io/blog/golang-logging/)
* [Go Testing](https://gobyexample.com/testing)