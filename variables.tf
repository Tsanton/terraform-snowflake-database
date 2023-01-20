variable "db_prefix" {
  description = "The prefix for the <db_name> that consists of <db_prefix>_<db_identifier>"
  type        = string

  validation {
    condition     = can(regex("^[A-Za-z]+$", var.db_prefix))
    error_message = "Only a-z and A-Z are valid characters for role prefixes."
  }
}

variable "db_identifier" {
  description = "The identifier for the <db_name> that consists of <db_prefix>_<db_identifier>"
  type        = string
}

variable "clone_from" {
  description = "The database name to create a clone from"
  type        = string
  default     = null
}

variable "tags_map" {
  type    = map(string)
  default = {}
}
