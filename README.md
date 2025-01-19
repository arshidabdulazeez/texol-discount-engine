# texol-discount-engine
## Setup and Installation

### Prerequisites
Go (version 1.18 or later)

A terminal or command prompt

### Steps to Set Up

#### 1.Clone the repository:

git clone https://github.com/your-repo/texol-discount-engine.git
cd texol-discount-engine

#### 2.Initialize the Go module (if not already initialized):

go mod tidy

#### 3.Verify the rules.json file exists in the config/ directory with the required discount rules.

#### 4.Run the server:

go run ./server/main.go

The server will start on http://localhost:8080

#### 5. Open another Terminal and run

curl -X POST -H 'Content-Type: application/json' -d '{ "order_total": 600, "customer_type": "premium" }' http://localhost:8080/apply-discount

you can change the order_total and customer_type arguments and test.

