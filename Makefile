build-amary:
	docker compose up amary redis_amary database_amary amary_client --build -d
build-auth:
	docker compose up auth database_auth redis_auth --build -d
build-amarolio:
	docker compose up gateway amarolio_client --build -d
start-amary:
	docker compose up amary redis_amary database_amary amary_client -d
start-auth:
	docker compose up auth database_auth redis_auth -d
start-amarolio:
	docker compose up gateway amarolio_client -d
stop-amary:
	docker compose down amary redis_amary database_amary amary_client
stop-auth:
	docker compose down auth database_auth redis_auth
stop-amarolio:
	docker compose down gateway amarolio_client
