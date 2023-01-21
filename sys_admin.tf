resource "snowflake_role" "db_sys_admin" {
  name    = upper("${snowflake_database.db.name}_SYS_ADMIN")
  comment = upper(jsonencode(merge(var.tags_map, {})))

  lifecycle {
    replace_triggered_by = [
      snowflake_database.db.name
    ]
  }
}

resource "snowflake_role_ownership_grant" "db_sys_admin_ownership" {
  on_role_name = snowflake_role.db_sys_admin.name
  to_role_name = "USERADMIN"

  current_grants                = "COPY"
  revert_ownership_to_role_name = "USERADMIN"
}

resource "snowflake_role_grants" "db_sys_admin_granted_sysadmin" {
  role_name = snowflake_role.db_sys_admin.name

  roles = ["SYSADMIN"]

  enable_multiple_grants = true

  depends_on = [
    snowflake_role_ownership_grant.db_sys_admin_ownership
  ]

  lifecycle {
    replace_triggered_by = [
      snowflake_database.db.name
    ]
  }
}

resource "snowflake_database_grant" "db_sys_admin_usage" {
  database_name = snowflake_database.db.name

  privilege = "USAGE"
  roles     = [snowflake_role.db_sys_admin.name]

  enable_multiple_grants = true
  with_grant_option      = false

  depends_on = [
    snowflake_database_grant.db_ownership
  ]
}
