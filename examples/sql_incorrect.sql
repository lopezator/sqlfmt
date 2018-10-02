crete table person (id int not null, name varchar(100) not null, display_name varchar(200) not null,
	metadata JSONB NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, PRIMARY KEY (id)
);