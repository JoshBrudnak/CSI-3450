drop table if exists budget_entry;
drop table if exists budget_category;
drop table if exists budget_user;

create table if not exists budget_user(id serial primary key, fname text, lname text, password text, email text, admin boolean, date date, t_money float);

create table if not exists budget_category(id serial primary key, category_name text, user_id int REFERENCES budget_user (id), b_value float);

create table if not exists budget_entry(id serial primary key, category_id int REFERENCES budget_category (id), date date, value float);

insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES('John', 'Doe', '$2a$10$VdPNYKHKomEf7aQhspANxOS4Y2vlAbgiOLcn5xWwaL2/NOOcETh5e$2a$10$nctfLFOnJLHkMLz1G3xYWead9Imsx4d8ji8iZNm4An7aWtqoRNiku', 'johndoe@oakland.edu', false, '2017-12-08', 15293);
insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES('Josh', 'Brudnak', '$2a$10$VdPNYKHKomEf7aQhspANxOS4Y2vlAbgiOLcn5xWwaL2/NOOcETh5e$2a$10$nctfLFOnJLHkMLz1G3xYWead9Imsx4d8ji8iZNm4An7aWtqoRNiku', 'joshuabrudnak@oakland.edu', true,  '2016-05-21', 459128);
insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES('Lilia', 'Suau', '$2a$10$VdPNYKHKomEf7aQhspANxOS4Y2vlAbgiOLcn5xWwaL2/NOOcETh5e$2a$10$nctfLFOnJLHkMLz1G3xYWead9Imsx4d8ji8iZNm4An7aWtqoRNiku', 'liliasuau@oakland.edu', true, '2013-05-30', 534928);
insert into budget_user(fname, lname, password, email, admin, date, t_money) VALUES('William', 'Carney', '$2a$10$VdPNYKHKomEf7aQhspANxOS4Y2vlAbgiOLcn5xWwaL2/NOOcETh5e$2a$10$nctfLFOnJLHkMLz1G3xYWead9Imsx4d8ji8iZNm4An7aWtqoRNiku', 'wcarney@oakland.edu', true, '2015-09-23', 52928);

insert into budget_category(user_id, category_name, b_value) VALUES(0, 'spending', 234);
insert into budget_category(user_id, category_name, b_value) VALUES(0, 'savings', 1304);
insert into budget_category(user_id, category_name, b_value) VALUES(0, 'investments', 1019);

insert into budget_category(user_id, category_name, b_value) VALUES(1, 'tithe', 450);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'vacation', 234);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'savings', 1304);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'investments', 1019);

insert into budget_category(user_id, category_name, b_value) VALUES(1, 'school', 1450);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'vacation', 12234);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'savings', 1304);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'fastfood', 2391);
insert into budget_category(user_id, category_name, b_value) VALUES(1, 'savings', 1019);

insert into budget_category(user_id, category_name, b_value) VALUES(2, 'charity', 150);
insert into budget_category(user_id, category_name, b_value) VALUES(2, 'vacation', 234);
insert into budget_category(user_id, category_name, b_value) VALUES(2, 'savings', 104);
insert into budget_category(user_id, category_name, b_value) VALUES(2, 'fastfood', 2391);

insert into budget_entry(category_id, date, value) VALUES(0, '2018-11-04', 34);
insert into budget_entry(category_id, date, value) VALUES(0, '2018-11-14', 54);
insert into budget_entry(category_id, date, value) VALUES(0, '2018-10-10', 92);
insert into budget_entry(category_id, date, value) VALUES(0, '2018-11-20', 89);
insert into budget_entry(category_id, date, value) VALUES(0, '2018-12-01', 5);

insert into budget_entry(category_id, date, value) VALUES(1, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(1, '2018-09-14', 4);
insert into budget_entry(category_id, date, value) VALUES(1, '2018-09-10', 552);
insert into budget_entry(category_id, date, value) VALUES(1, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(1, '2018-09-01', 9);

insert into budget_entry(category_id, date, value) VALUES(2, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(2, '2018-09-14', 194);
insert into budget_entry(category_id, date, value) VALUES(2, '2018-09-10', 552);
insert into budget_entry(category_id, date, value) VALUES(2, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(2, '2018-09-01', 9);

insert into budget_entry(category_id, date, value) VALUES(3, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(3, '2018-09-14', 94);
insert into budget_entry(category_id, date, value) VALUES(3, '2018-09-10', 1552);
insert into budget_entry(category_id, date, value) VALUES(3, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(3, '2018-09-01', 9);

insert into budget_entry(category_id, date, value) VALUES(4, '2018-12-06', 134);
insert into budget_entry(category_id, date, value) VALUES(4, '2018-09-14', 64);
insert into budget_entry(category_id, date, value) VALUES(4, '2018-09-10', 552);
insert into budget_entry(category_id, date, value) VALUES(4, '2018-4-20', 89);
insert into budget_entry(category_id, date, value) VALUES(4, '2018-09-01', 9);

insert into budget_entry(category_id, date, value) VALUES(5, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(5, '2018-12-14', 4);
insert into budget_entry(category_id, date, value) VALUES(5, '2018-09-10', 552);
insert into budget_entry(category_id, date, value) VALUES(5, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(5, '2018-09-01', 9);

insert into budget_entry(category_id, date, value) VALUES(6, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(6, '2018-09-14', 94);
insert into budget_entry(category_id, date, value) VALUES(6, '2018-09-10', 52);
insert into budget_entry(category_id, date, value) VALUES(6, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(6, '2018-09-01', 239);

insert into budget_entry(category_id, date, value) VALUES(7, '2018-09-04', 134);
insert into budget_entry(category_id, date, value) VALUES(7, '2018-09-14', 94);
insert into budget_entry(category_id, date, value) VALUES(7, '2018-09-10', 52);
insert into budget_entry(category_id, date, value) VALUES(7, '2018-09-20', 89);
insert into budget_entry(category_id, date, value) VALUES(7, '2018-09-01', 239);
