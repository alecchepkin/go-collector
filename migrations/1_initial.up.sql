-- set search_path to mytarget;

CREATE OR REPLACE FUNCTION upd_updated_at() RETURNS TRIGGER
  LANGUAGE plpgsql
AS
$$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$;

create or replace function add_time_fields(table_name text) returns void
  language plpgsql
as
$$
declare
  trigger_name text;
begin
  execute 'alter table ' || table_name || ' add column created_at timestamp with time zone default now() not null;';
  execute 'alter table ' || table_name || ' add column updated_at timestamp with time zone default now() not null;';
  trigger_name := 't_' || table_name || '_upt';

  execute 'create trigger ' || trigger_name || ' before update on ' || table_name ||
          ' for each row execute procedure upd_updated_at()';
end;
$$;


-- account
create table cabinet
(
  id            bigserial primary key,
  cabinet_id    int unique,
  account_id    int,
  username      text,
  name          text,
  email         text,
  balance       decimal(10, 2),
  currency      varchar(3),
  parent_id     int,
  is_agency     boolean default false,
  is_external   boolean default false,
  client_id     text,
  client_secret text,
  access_token  text,
  refresh_token text
);
select add_time_fields('cabinet');
create index on cabinet (is_agency);
create index on cabinet (parent_id);

insert into cabinet (cabinet_id, account_id, name, is_agency, client_id, client_secret)
values (2651700, 1886590, 'Mobile Marketing', true, 'NdxH34Mz6fmVN35P',
        'eaJ56C6YoSYqWPpsnNpMRh0F6wDOnZ8yxUz8V3fhI49ir20bCpHO3j6hQgDHVTFheEk16O3irg4WUMxXHi3y1OVTFHhbAGa3iWUSG868Py7KqY39iWDZB7Z6oXQJ1FxQsg9COwbS6jaLYsJOwQPmMAJh6x7PC3eNCUdgS38YbV2MoGStuafbfngQEOREAoNuyKeTRTHuFwJj3upvYkhEig6uVy7E8xercBbauPlNb58n464OrqoZUBe');

insert into cabinet (cabinet_id, account_id, name, is_agency, client_id, client_secret)
values (2651782, 1886702, 'Mobile Traffic', true, 'PH3IDLIGvAfwTD4a',
        'cYYf8qP4edmuUcCs5ngM1RZQmniotYoens0hCLuSIRpbr7Nh56bZBhSjNBpLo370ybXXFGuckW5sY90NHqQM6KfzpLPLEmPCu71rdYcd0z05NyDrx7vTR3107NPxaa3qeldVAf7R89GO8dZuHoH0uJsSaIGb9ArPHKHKM7kjwFeyc1leUlTYEGzXIIb9TcSMVrArfWaIPJMAMORbMMiLFIv');
--
insert into cabinet (cabinet_id, account_id, name, is_agency, client_id, client_secret)
values (2886399, 2066223, 'Герман Капнин', true, 'Fnf6gUaFmqHKzB0B',
        'BlAHSYquGdkpnmH3uv06FOKehB6UR42p7u3awzzly0cUsrKqyjcoDdAF9OlMVAUjoTJKb1FrT2mmGTXHX6aRLeghFQmIgePr0eWiQYh3PjohRh05GHbCUcHE6RuFKxiqNorLpqvJGj8LtrYru6Nr2s36G14c4KFDhhCVAESRK63bDljUUuniSvcSfI91qj0PfFNPdwCmVMMM9AUlYbApYyv8xlZFfyb0b');
--
-- campaign
create table campaign
(
  id               bigserial primary key,
  cabinet_id       int,
  campaign_id      int,
  name             text             not null,
  status           text             not null,
  created          text             not null,
  updated          text             not null,
  budget_limit     double precision not null,
  budget_limit_day double precision not null
);

select add_time_fields('campaign');

create unique index on campaign (cabinet_id, campaign_id);


-- region
create table region
(
  id        bigserial primary key,
  region_id int unique,
  name      text not null,
  parent_id int  not null,
  flags     text
);

select add_time_fields('region');
-- package
create table package
(
  id                   bigserial primary key,
  cabinet_id           int,
  package_id           int,
  name                 text,
  url_types            text,
  max_uniq_shows_limit int,
  priced_event_type    int,
  max_price_per_unit   text,
  description          text,
  created              text,
  price                text,
  updated              text,
  paid_event_type      int,
  status               text,
  url_type             text,
  pads_tree_id         int,
  objective            text,
  banner_format_id     int,
  related_package_ids  text,
  flags                text

);

select add_time_fields('package');
create unique index on package (cabinet_id, package_id);


-- banner
create table banner
(
  id                bigserial primary key,
  banner_id         bigint unique,
  campaign_id       int,
  moderation_status text,
  cabinet_id        int
);
select add_time_fields('banner');

-- stat_banner
create table stat_banner
(
  id         bigserial primary key,
  banner_id  int,
  cabinet_id int,
  date       date,
  shows      int,
  clicks     int,
  goals      int,
  spent      text,
  reach      int,
  total      int,
  increment  int
);
select add_time_fields('stat_banner');
create unique index on stat_banner (banner_id, date);
