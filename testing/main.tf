terraform {
  required_providers {
    snowflake = {
      source  = "Snowflake-Labs/snowflake"
      version = "0.53.0"
    }
  }
}
variable "snowflake_uid" { type = string }
variable "snowflake_pwd" { type = string }
variable "snowflake_account" { type = string }
variable "snowflake_region" { type = string }

provider "snowflake" {
  username = var.snowflake_uid
  password = var.snowflake_pwd
  account  = var.snowflake_account
  region   = var.snowflake_region
  role     = "ACCOUNTADMIN"
}

variable "db_prefix" { type = string }
variable "db_identifier" { type = string }

module "db" {
  source        = "../"
  db_prefix     = var.db_prefix
  db_identifier = var.db_identifier
}
