create schema "strava-gate";

create table "strava-gate".users
(
	user_id int not null
		constraint users_pk
			primary key,
	access_token varchar(50) not null,
	refresh_token varchar(50) not null,
	expires_to integer not null
);

create sequence "strava-gate".users_user_id_seq;
alter table "strava-gate".users alter column user_id set default nextval('strava-gate.users_user_id_seq');
alter sequence "strava-gate".users_user_id_seq owned by "strava-gate".users.user_id;