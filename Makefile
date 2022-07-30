newMigration:
	migrate create -ext sql -dir db/migrations -seq ${name}

migrateUpLocal:
	migrate -path db/migrations -database postgresql://tkkt:password@localhost:5432/tkkt?sslmode=disable -verbose up

migrateDownLocal:
	migrate -path db/migrations -database postgresql://tkkt:password@localhost:5432/tkkt?sslmode=disable -verbose down

migrateupdev:
	migrate -path db/migrations -database ${DEV_TICKET_SALES_PG_URL} -verbose up

migratedown1:
	 migrate -path db/migrations -database ${DEV_TICKET_SALES_PG_URL} -verbose down 1

migrateup1:
	 migrate -path db/migrations -database ${DEV_TICKET_SALES_PG_URL} -verbose up 1

migrateforceversion:
	 migrate -path db/migrations -database ${DEV_TICKET_SALES_PG_URL} force ${version}