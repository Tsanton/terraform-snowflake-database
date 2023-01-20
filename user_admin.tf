resource "snowflake_role" "db_user_admin" {
  name    = upper("${snowflake_database.db.name}_USER_ADMIN")
  comment = upper(jsonencode(merge(var.tags_map, {})))
}

resource "snowflake_role_ownership_grant" "db_user_admin_ownership" {
  on_role_name = snowflake_role.db_user_admin.name
  to_role_name = "USERADMIN"

  current_grants                = "COPY"
  revert_ownership_to_role_name = "USERADMIN"
}

resource "snowflake_role_grants" "db_user_admin_granted_useradmin" {
  role_name = snowflake_role.db_user_admin.name

  roles = ["USERADMIN"]

  depends_on = [
    snowflake_role_ownership_grant.db_user_admin_ownership
  ]
}

resource "snowflake_account_grant" "db_user_admin_create_role_grant" {
  roles = [
    snowflake_role.db_user_admin.name
  ]
  privilege              = "CREATE ROLE"
  with_grant_option      = false
  enable_multiple_grants = true
}
