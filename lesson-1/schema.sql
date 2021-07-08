-- Проект по прокату (краткосрочная аренда) вещей и интструмента. Клиен сервиса может прийти и оформить в аренду нужную ему вещь или инструмент.
-- Клиент выбирает нужную ему вещь, определяет вариант (с оплатой по часам, дням, неделям и т.д.) и срок на который он хочет взять её в аренду.
-- Можно оформить в аренду несколько вещей, каждая вещь взятая в аренду представляет собой отдельную запись в таблице (rent_list).
-- Представляет из себя 4 таблицы:
-- таблица аренованых вещей (rent_list) включающая в себя связь с таблицей клиенты (clients) и с таблицей (items) в которой содержится информация что и по какой цене сдаётся в аренду,
-- таблица (items) в свою очередь связана с таблицей (rent_price) в которой содержится информация по ценами на аренду (rent_price),
-- так как цена зависит от того что и на какой срок оформляется в аренду.

CREATE TABLE client (
    id serial not null unique,
    -- Фамилия
    first_name varchar(255) not null,
    -- Имя
    last_name varchar(255) not null,
    -- Отчество
    middle_name varchar(255),
    -- Также с помощью номера телефона делаем проверку на уникальность и избегаем дублирования записей.
    phone varchar(255) not null unique,
    -- Дата регистрации в сервисе.
    registered_at timestamp default now()
);

CREATE TABLE rent_price (
    id serial not null unique,
    -- Наименование варианта аренды. Например по часам, или по дням и т.д.
    name varchar(255) not null unique,
    -- Цена аренды за единицу (час, день, неделя и т.д.).
    price int not null
);

CREATE TABLE item (
    id serial not null unique,
    -- Наименование предмета аренды.
    name varchar(255) not null,
    -- Описание предмета аренды.
    description text,
    -- Из связанной таблицы (rent_price) выбираем вариант аренды.
    rent_price_id int not null,
    -- В это поле записывается информация о том когда (дата и время) должна закончиться аренда.
    expires_at timestamp,

    CONSTRAINT item_rent_price_id FOREIGN KEY (rent_price_id) REFERENCES rent_price (id) ON DELETE cascade
);


CREATE TABLE rent_list (
    id serial not null unique,
    -- Из связанной таблицы (client) выбираем клиета который берёт вещь в аренду.
    client_id int not null,
    -- Из связанной таблицы (item) выбираем вещь которую клиент берёт в аренду.
    item_id int not null,
    -- В это поле записывается информация о стоимости единицы аренды через связанную таблицу (item), а она в свою очередь из берёт информацию из
    -- связанной таблицы rent_price где выбран подходящий вариант аренды (с оплатой по часам, по дням и т.д.)
    rent_price int not null,
    -- В этом поле устанавливаем срок на который оформляется аренда. 
    duration int not null,
    -- В это поле расчитывается и записывается информация о стоимости аренды.
    rental_amount int not null,
    -- В это поле записывается дата и время начала аренды.
    start_rent_at timestamp not null,
    -- В это поле записывается информация о фактической дате и времени завершения аренды.
    -- И мы можем его сравнить с датой и временем из таблицы (item поле expires_at) когда должна была завершиться аренда.
    end_rent_at timestamp,
    -- В это поле рассчитывается и заноситься информация о необходимости совершения доплаты если время аренды было превышено.
    surcharge int,

    CONSTRAINT rent_list_client_id FOREIGN KEY (client_id) REFERENCES client (id),
    CONSTRAINT rent_list_item_id FOREIGN KEY (item_id) REFERENCES item (id)
    
);