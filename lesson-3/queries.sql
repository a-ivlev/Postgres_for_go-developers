
Сделаем запрос к БД и узнаем что сейчас оформлено в аренду и когда эту вещь должны вернуть.

select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at from rent_list
join client ON rent_list.client_id = client.id
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id;


 id | first_name | last_name | id |                   name                    |            name             | duration | rental_amount |       start_rent_at        |         expires_at         
----+------------+-----------+----+-------------------------------------------+-----------------------------+----------+---------------+----------------------------+----------------------------
  1 | Владимир   | Сидоров   |  8 | Горный велосипед GT 27.5 AGGRESSOR EXPERT | Аренда с оплатой за день.   |        2 |          1000 | 2021-07-12 08:02:43.726234 | 2021-07-14 08:02:43.726234
  2 | Иван       | Иванов    |  4 | Вентилятор напольный ZANUSSI ZFF-705      | Аренда с оплатой за месяц.  |        3 |         12000 | 2021-07-12 08:02:43.726234 | 2021-10-12 08:02:43.726234
  3 | Карина     | Белова    |  1 | Дрель-шуруповерт MAKITA DF0300            | Аренда с оплатой за час.    |        3 |           150 | 2021-07-12 08:02:43.726234 | 2021-07-12 11:02:43.726234
  4 | Семён      | Семёнов   |  2 | Дрель-шуруповерт HAMMER DRL420A           | Аренда с оплатой за неделю. |        1 |          1250 | 2021-07-12 08:02:43.726234 | 2021-07-19 08:02:43.726234
(4 rows)


Сделаем запрос по номеру телефона клиента и узнаем что взял в аренду данный клиент.

select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at from rent_list 
join client ON rent_list.client_id = client.id 
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id
WHERE client.phone = '+7 411 923 8377';

 id | first_name | last_name | id |              name              |          rental          | duration | rental_amount |       start_rent_at        |         expires_at         
----+------------+-----------+----+--------------------------------+--------------------------+----------+---------------+----------------------------+----------------------------
  3 | Карина     | Белова    |  1 | Дрель-шуруповерт MAKITA DF0300 | Аренда с оплатой за час. |        3 |           150 | 2021-07-12 08:02:43.726234 | 2021-07-12 11:02:43.726234
(1 row)


Осуществим поиск тех кто не вернул вовремя взятую в аренду вещь.

select rent_list.id, client.first_name, client.last_name, item.id, item.name, rent_price.name as rental, rent_list.duration,  rent_list.rental_amount, rent_list.start_rent_at, item.expires_at, now() as now from rent_list 
join client ON rent_list.client_id = client.id 
join item ON rent_list.item_id = item.id
join rent_price ON rent_list.rent_price_id = rent_price.id
WHERE item.expires_at < now();

 id | first_name | last_name | id |                   name                    |          rental           | duration | rental_amount |       start_rent_at        |         expires_at         |              now              
----+------------+-----------+----+-------------------------------------------+---------------------------+----------+---------------+----------------------------+----------------------------+-------------------------------
  1 | Владимир   | Сидоров   |  8 | Горный велосипед GT 27.5 AGGRESSOR EXPERT | Аренда с оплатой за день. |        2 |          1000 | 2021-07-12 08:02:43.726234 | 2021-07-14 08:02:43.726234 | 2021-07-14 12:48:34.982198+00
  3 | Карина     | Белова    |  1 | Дрель-шуруповерт MAKITA DF0300            | Аренда с оплатой за час.  |        3 |           150 | 2021-07-12 08:02:43.726234 | 2021-07-12 11:02:43.726234 | 2021-07-14 12:48:34.982198+00
(2 rows)

