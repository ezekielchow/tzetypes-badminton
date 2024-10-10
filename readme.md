## Setup
1. Copy `.env.example` to `.env` and change the parameters if necessary
2. Run `docker compose up` to start
3. API endpoints will be available according to ports stated in the docker-compose.yml file

## Development
- Use `make create-migration` to create migrations
- Use `api` folder to add openapi schema. Then generate with `make openapi`. Remember to add folders required to hold the generated files
- In every module there is `handler, services, store`
  - `handler` is use to handle routing of your apis
  - `services` is use to hold your business logic
  - `store` is use to communicate with your persistent storage
- Use `make test` to run tests

## Important notes
- Valid reference code is `ABC1234`

## Flow
1. Use POST `api/portfolios` to add portfolio first
2. Use GET `api/portfolios` to list down all portfolios
3. Use POST `api/invest` to invest. Return is allocations towards each portfolio

## Improvements:
- Store deposits when calling `api/invest`
- Add swagger for dashboard for endpoints
- Find a way to validate portfolio_ids before doing allocation. Now validating using populatePortfolio()
- Validate request and responses against schema
- Add more tests for postgres code and services

## API documentation - POSTMAN
```
{
	"info": {
		"_postman_id": "c4bb5bdf-530f-4689-b56f-3274f30cc3f3",
		"name": "Stashaway Assignment",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json",
		"_exporter_id": "8535914"
	},
	"item": [
		{
			"name": "portfolios",
			"request": {
				"method": "GET",
				"header": [],
				"url": "http://localhost:3001/api/portfolios"
			},
			"response": []
		},
		{
			"name": "deposit-plans",
			"request": {
				"method": "GET",
				"header": [],
				"url": "http://localhost:3002/api/deposit-plans"
			},
			"response": []
		},
		{
			"name": "portfolios",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "[\n    {\n        \"name\": \"High Risk\",\n        \"description\": \"my short term xxx\"\n    },\n    {\n        \"name\": \"low risk\",\n        \"description\": \"do not take out\"\n    }\n]",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "http://localhost:3001/api/portfolios"
			},
			"response": []
		},
		{
			"name": "invest",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"deposit_plans\": [\n        {\n            \"name\": \"dp one time\",\n            \"type\": \"one-time\",\n            \"distributions\": [\n                {\n                    \"portfolio_id\": \"4aba68a5-dc88-499f-a837-2bc2ff6e4c9f\",\n                    \"amount\": 333.10\n                },\n                {\n                    \"portfolio_id\": \"15f1a0bb-85e3-481c-945c-b4371e9648ce\",\n                    \"amount\": 333.09\n                }\n            ]\n        },\n        {\n            \"name\": \"dp monthly\",\n            \"type\": \"monthly\",\n            \"distributions\": [\n                {\n                    \"portfolio_id\": \"4aba68a5-dc88-499f-a837-2bc2ff6e4c9f\",\n                    \"amount\": 0\n                },\n                {\n                    \"portfolio_id\": \"15f1a0bb-85e3-481c-945c-b4371e9648ce\",\n                    \"amount\": 100\n                }\n            ]\n        }\n    ],\n    \"deposits\": [\n        {\n            \"amount\": 660.65,\n            \"referenceCode\": \"ABC1234\"\n        }\n    ]\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": "http://localhost:3002/api/invest"
			},
			"response": []
		}
	]
}
```