resource "snowflake_database" "db" {
  name                        = upper("${var.db_prefix}_${var.db_identifier}")
  comment                     = upper(jsonencode(merge(var.tags_map, {})))
  data_retention_time_in_days = 3

  from_database = var.clone_from == null ? "" : var.clone_from
}

resource "snowflake_database_grant" "db_ownership" {
  database_name = snowflake_database.db.name

  privilege = "OWNERSHIP"
  roles = [
    snowflake_role.db_sys_admin.name
  ]

  with_grant_option      = false
  enable_multiple_grants = true

  depends_on = [
    snowflake_role_ownership_grant.db_sys_admin_ownership
  ]
}
