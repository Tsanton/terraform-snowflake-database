# **terraform-snowflake-schema**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Requirements](#requirements)
- [Providers](#providers)
- [Modules](#modules)
- [Resources](#resources)
- [Inputs](#inputs)
- [Outputs](#outputs)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_snowflake"></a> [snowflake](#requirement\_snowflake) | >=0.51.0, < 1.0.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_snowflake"></a> [snowflake](#provider\_snowflake) | >=0.51.0, < 1.0.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [snowflake_account_grant.db_user_admin_create_role_grant](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/account_grant) | resource |
| [snowflake_database.db](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/database) | resource |
| [snowflake_database_grant.db_ownership](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/database_grant) | resource |
| [snowflake_database_grant.db_sys_admin_usage](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/database_grant) | resource |
| [snowflake_role.db_sys_admin](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role) | resource |
| [snowflake_role.db_user_admin](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role) | resource |
| [snowflake_role_grants.db_sys_admin_granted_sysadmin](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role_grants) | resource |
| [snowflake_role_grants.db_user_admin_granted_useradmin](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role_grants) | resource |
| [snowflake_role_ownership_grant.db_sys_admin_ownership](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role_ownership_grant) | resource |
| [snowflake_role_ownership_grant.db_user_admin_ownership](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/role_ownership_grant) | resource |
| [snowflake_warehouse.wh](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/warehouse) | resource |
| [snowflake_warehouse_grant.wh_ownership](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/warehouse_grant) | resource |
| [snowflake_warehouse_grant.wh_usage](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/resources/warehouse_grant) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_clone_from"></a> [clone\_from](#input\_clone\_from) | The database name to create a clone from | `string` | `null` | no |
| <a name="input_db_identifier"></a> [db\_identifier](#input\_db\_identifier) | The identifier for the <db\_name> that consists of <db\_prefix>\_<db\_identifier> | `string` | n/a | yes |
| <a name="input_db_prefix"></a> [db\_prefix](#input\_db\_prefix) | The prefix for the <db\_name> that consists of <db\_prefix>\_<db\_identifier> | `string` | n/a | yes |
| <a name="input_tags_map"></a> [tags\_map](#input\_tags\_map) | n/a | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_db_name"></a> [db\_name](#output\_db\_name) | The name of the database |
| <a name="output_sys_admin_default_warehouse_name"></a> [sys\_admin\_default\_warehouse\_name](#output\_sys\_admin\_default\_warehouse\_name) | The name of the default warehouse for the db\_sys\_admin |
| <a name="output_sys_admin_name"></a> [sys\_admin\_name](#output\_sys\_admin\_name) | The name of the database sysadmin role |
| <a name="output_user_admin_name"></a> [user\_admin\_name](#output\_user\_admin\_name) | The name of the database useradmin role |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
