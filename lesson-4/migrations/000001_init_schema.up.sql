CREATE TABLE client (
    id serial not null unique,
    first_name varchar(50) not null,
    middle_name varchar(50),
    last_name varchar(50) not null,  
    phone varchar(20) not null unique,
    registered_at timestamp default now()
);

CREATE TABLE item (
    id serial not null unique,
    name varchar(255) not null,
    description text,
    expires_at timestamp
);

CREATE TABLE rent_price (
    id serial not null unique,
    item_id int not null,
    name varchar(255) not null,
    price int not null

    CONSTRAINT price_check CHECK (price > 0),
    CONSTRAINT rent_price_item_id FOREIGN KEY (item_id) REFERENCES item (id)
);

CREATE TABLE rent_list (
    id serial not null unique,
    client_id int not null,
    item_id int not null,
    rent_price_id int not null,
    duration int not null,
    rental_amount int not null,
    start_rent_at timestamp not null default now(),
    end_rent_at timestamp,
    surcharge int,

    CONSTRAINT duration_check CHECK (duration > 0),
    CONSTRAINT rent_list_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    CONSTRAINT rent_list_item_id FOREIGN KEY (item_id) REFERENCES item (id),
    CONSTRAINT rent_list_rent_price_id FOREIGN KEY (rent_price_id) REFERENCES rent_price (id) 
);

