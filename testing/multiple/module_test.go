package no_test_main_testing_test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	i "github.com/tsanton/goflake-client/goflake/integration"
)

func Test_module_multiple_dbs_one_db_name_update(t *testing.T) {
	/* Arrange */
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		PluginDir:    "/workspace/.providercache",
		TerraformDir: "./",
		Vars: map[string]interface{}{
			"snowflake_uid":     os.Getenv("SNOWFLAKE_UID"),
			"snowflake_pwd":     os.Getenv("SNOWFLAKE_PWD"),
			"snowflake_account": os.Getenv("SNOWFLAKE_ACCOUNT"),
			"snowflake_region":  os.Getenv("SNOWFLAKE_REGION"),
			"db_prefix":         "INTEGRATIONTEST",
			"db_identifier":     "UPDATE_DB",
		},
		EnvVars: map[string]string{},
		NoColor: true,
	})
	defer terraform.Destroy(t, terraformOptions)
	_, err := terraform.InitAndApplyE(t, terraformOptions)
	i.ErrorFailNow(t, err)

	/* Act */
	terraformOptions.Vars["db_identifier"] = "UPDATED_DB"
	_, err = terraform.ApplyE(t, terraformOptions)

	/* Assert */
	i.ErrorFailNow(t, err)
}
