output "db_name" {
  description = "The name of the database"
  value       = snowflake_database.db.name
}

output "sys_admin_name" {
  description = "The name of the database sysadmin role"
  value       = snowflake_role.db_sys_admin.name
}

output "user_admin_name" {
  description = "The name of the database useradmin role"
  value       = snowflake_role.db_user_admin
}

output "sys_admin_default_warehouse_name" {
  description = "The name of the default warehouse for the db_sys_admin"
  value       = snowflake_warehouse.wh.name
}
