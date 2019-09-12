-- set search_path to mytarget;

drop table if exists region;
drop table if exists stat_banner;
drop table if exists banner;
drop table if exists package;
drop table if exists campaign;
drop table if exists cabinet;
drop function if exists add_time_fields(table_name text);
drop function if exists upd_updated_at();
