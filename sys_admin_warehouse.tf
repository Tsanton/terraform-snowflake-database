resource "snowflake_warehouse" "wh" {
  name                = upper("${snowflake_role.db_sys_admin.name}_WH")
  comment             = upper(jsonencode(merge(var.tags_map, {})))
  warehouse_size      = "x-small"
  initially_suspended = false
  auto_resume         = true
  auto_suspend        = 60
}

resource "snowflake_warehouse_grant" "wh_ownership" {
  warehouse_name = snowflake_warehouse.wh.name
  privilege      = "OWNERSHIP"

  roles = [
    "SYSADMIN"
  ]

  enable_multiple_grants = true
  with_grant_option      = false
}

resource "snowflake_warehouse_grant" "wh_usage" {
  warehouse_name = snowflake_warehouse.wh.name
  privilege      = "USAGE"

  roles = [
    snowflake_role.db_sys_admin.name
  ]

  enable_multiple_grants = true
  with_grant_option      = false

  depends_on = [
    snowflake_warehouse_grant.wh_ownership
  ]
}
