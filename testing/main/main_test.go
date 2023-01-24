package main_test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	g "github.com/tsanton/goflake-client/goflake"
	i "github.com/tsanton/goflake-client/goflake/integration"
	a "github.com/tsanton/goflake-client/goflake/models/assets"
	d "github.com/tsanton/goflake-client/goflake/models/describables"
	dg "github.com/tsanton/goflake-client/goflake/models/describables/grants"
	e "github.com/tsanton/goflake-client/goflake/models/entities"
	eg "github.com/tsanton/goflake-client/goflake/models/entities/grants"
	"github.com/tsanton/goflake-client/goflake/models/enums"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func TestMain(m *testing.M) {
	/* Setup */
	t := &testing.T{}
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./",
		Vars: map[string]interface{}{
			"snowflake_uid":     os.Getenv("SNOWFLAKE_UID"),
			"snowflake_pwd":     os.Getenv("SNOWFLAKE_PWD"),
			"snowflake_account": os.Getenv("SNOWFLAKE_ACCOUNT"),
			"snowflake_region":  os.Getenv("SNOWFLAKE_REGION"),
			"db_prefix":         "INTEGRATION",
			"db_identifier":     "TEST_DB",
		},
		EnvVars: map[string]string{},
		NoColor: true,
	})
	defer terraform.Destroy(t, terraformOptions)
	if _, err := terraform.InitAndApplyE(t, terraformOptions); err != nil {
		os.Exit(1)
	}
	/* Run */
	// exitCode := m.Run()
	m.Run()

	/* Teardown: if you os.exit() defer function are not run */
	// os.Exit(exitCode)

}

func Goflake() (*g.GoflakeClient, u.Stack[a.ISnowflakeAsset]) {
	stack := u.Stack[a.ISnowflakeAsset]{}
	cli := g.GoflakeClient{
		SnowflakeHost: os.Getenv("SNOWFLAKE_HOST"),
		SnowflakeUid:  os.Getenv("SNOWFLAKE_UID"),
		SnowflakePwd:  os.Getenv("SNOWFLAKE_PWD"),
		SnowflakeRole: os.Getenv("SNOWFLAKE_ROLE"),
		SnowflakeWh:   os.Getenv("SNOWFLAKE_WH"),
	}

	cli.Open()

	return &cli, stack
}

func Test_database_exists(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	db, err := g.Describe[e.Database](cli, &d.Database{Name: "INTEGRATION_TEST_DB"})
	i.ErrorFailNow(t, err)

	/* Assert */
	a.Nil(err)
	a.Equal(db.Name, "INTEGRATION_TEST_DB")
	a.Equal(db.Owner, "INTEGRATION_TEST_DB_SYS_ADMIN")
}

func Test_local_sys_admin_exists(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	role, err := g.Describe[e.Role](cli, &d.Role{Name: "INTEGRATION_TEST_DB_SYS_ADMIN"})

	/* Assert */
	a.Nil(err)
	a.Equal(role.Name, "INTEGRATION_TEST_DB_SYS_ADMIN")
	a.Equal(role.Owner, "USERADMIN")
}

func Test_local_user_admin_exists(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	role, err := g.Describe[e.Role](cli, &d.Role{Name: "INTEGRATION_TEST_DB_USER_ADMIN"})

	/* Assert */
	a.Nil(err)
	a.Equal(role.Name, "INTEGRATION_TEST_DB_USER_ADMIN")
	a.Equal(role.Owner, "USERADMIN")
}

// Test that role hierarchy is intact: db_sys_admin -> sysadmin -> accountadmin
func Test_local_sys_admin_hierarchy(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	hierarchy, err := g.Describe[e.RoleHierarchy](cli, &d.RoleHierarchy{RoleName: "INTEGRATION_TEST_DB_SYS_ADMIN"})

	/* Assert */
	a.Nil(err)
	sa, ok := lo.Find(hierarchy.InheritingRoles, func(i e.InheritedRole) bool { return i.ParentRoleName == "SYSADMIN" })
	a.True(ok)
	a.Equal(sa.DistanceFromSource, 0)

	aa, ok := lo.Find(hierarchy.InheritingRoles, func(i e.InheritedRole) bool { return i.ParentRoleName == "ACCOUNTADMIN" })
	a.True(ok)
	a.Equal(aa.DistanceFromSource, 1)
}

// Test that role hierarchy is intact: db_user_admin -> useradmin -> securityadmin -> accountadmin
func Test_local_user_admin_hierarchy(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	hierarchy, err := g.Describe[e.RoleHierarchy](cli, &d.RoleHierarchy{RoleName: "INTEGRATION_TEST_DB_USER_ADMIN"})

	/* Assert */
	a.Nil(err)
	ua, ok := lo.Find(hierarchy.InheritingRoles, func(i e.InheritedRole) bool { return i.ParentRoleName == "USERADMIN" })
	a.True(ok)
	a.Equal(ua.DistanceFromSource, 0)

	sa, ok := lo.Find(hierarchy.InheritingRoles, func(i e.InheritedRole) bool { return i.ParentRoleName == "SECURITYADMIN" })
	a.True(ok)
	a.Equal(sa.DistanceFromSource, 1)

	aa, ok := lo.Find(hierarchy.InheritingRoles, func(i e.InheritedRole) bool { return i.ParentRoleName == "ACCOUNTADMIN" })
	a.True(ok)
	a.Equal(aa.DistanceFromSource, 2)
}

func Test_local_sys_admin_privileges(t *testing.T) {
	/* Arrange */
	a := assert.New(t)
	cli, _ := Goflake()

	/* Act */
	res, err := g.Describe[*eg.RoleGrants](cli, &dg.RoleGrant{RoleName: "INTEGRATION_TEST_DB_SYS_ADMIN"})

	/* Assert */
	a.Nil(err)
	assert.Len(t, res.Grants, 3)

	dbUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool {
		return i.Privilege == enums.PrivilegeUsage && i.GrantedOn == enums.SnowflakeObjectDatabase
	})
	assert.True(t, ok)
	assert.Equal(t, "INTEGRATION_TEST_DB_SYS_ADMIN", dbUsage.GrantedBy)
	assert.Equal(t, "INTEGRATION_TEST_DB", dbUsage.GrantTargetName)

	dbOwnership, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool {
		return i.Privilege == enums.PrivilegeOwnership && i.GrantedOn == enums.SnowflakeObjectDatabase
	})
	assert.True(t, ok)
	assert.Equal(t, "INTEGRATION_TEST_DB_SYS_ADMIN", dbOwnership.GrantedBy)
	assert.Equal(t, "INTEGRATION_TEST_DB", dbOwnership.GrantTargetName)

	warehouseUsage, ok := lo.Find(res.Grants, func(i eg.RoleGrant) bool {
		return i.Privilege == enums.PrivilegeUsage && i.GrantedOn == enums.SnowflakeObjectWarehouse
	})
	assert.True(t, ok)
	assert.Equal(t, "SYSADMIN", warehouseUsage.GrantedBy)
	assert.Equal(t, "INTEGRATION_TEST_DB_SYS_ADMIN_WH", warehouseUsage.GrantTargetName)
}
