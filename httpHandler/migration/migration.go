package migration

import "github.com/rubenv/sql-migrate"

func Getmigration() migrate.MemoryMigrationSource {
	var Migration = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id: "123",
				Up: []string{`  
								create table user(
								name 	 varchar(50) NOT NULL,
								email_id  varchar(50) NOT NULL,
								password varchar(50) NOT NULL,
								age   	 int     	 NOT NULL,
								auth_id  varchar(50) NOT NULL,
								UNIQUE(auth_id),
								PRIMARY KEY(email_id)
								)`,
					`create table post(
								Name		varchar(50) 	NOT NULL,
								id			varchar(50)		NOT NULL,
								title		varchar(50) 	NOT NULL,
								discription	varchar(50) 	NOT NULL,
								FOREIGN KEY (id) REFERENCES user(email_id)
								)`},
				Down: []string{`
									drop table post;
									drop table user;
								`},
			},
		},
	}

	return *Migration
}
