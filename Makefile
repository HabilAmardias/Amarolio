build-amary:
	docker compose -f docker-compose.dev.yaml up amary redis_amary database_amary amary_client amary_redirect --build -d
build-auth:
	docker compose -f docker-compose.dev.yaml up auth database_auth redis_auth --build -d
build-auth-client:
	docker compose -f docker-compose.dev.yaml up auth_client --build -d
build-amarolio:
	docker compose -f docker-compose.dev.yaml up gateway amarolio_client --build -d
start-amary:
	docker compose up amary redis_amary database_amary amary_client -d
start-auth:
	docker compose up auth database_auth redis_auth -d
start-auth-client:
	docker compose up auth_client -d
start-amarolio:
	docker compose up gateway amarolio_client -d
stop-amary:
	docker compose down amary redis_amary database_amary amary_client
stop-auth:
	docker compose down auth database_auth redis_auth
stop-auth-client:
	docker compose down auth_client
stop-amarolio:
	docker compose down gateway amarolio_client
