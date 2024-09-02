package ascertain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
)

var cerealList = []string{
	"fruitloops",
	"cheerios",
	"cocoapuffs",
}

var AuditLocks = []ascertain.DependencyLock{
	{
		Name:    "1",
		Version: "v1",
		License: ascertain.License{
			Kind: "fruitloops",
		},
	},
	{
		Name:    "2",
		Version: "v2",
		License: ascertain.License{
			Kind: "cheerios",
		},
	},
	{
		Name:    "3",
		Version: "v3",
		License: ascertain.License{
			Kind: "cocoapuffs",
		},
	},
}

func TestLicenseAuditorAuditLock(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := ascertain.MakeLicenseWhitelist(cerealList)
		a := ascertain.MakeLicenseAuditor(w)

		err := a.AuditLock(AuditLocks[0])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[1])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[2])
		assert.Nil(t, err)
	})
	t.Run("error on non-whitelist license", func(t *testing.T) {
		w := ascertain.MakeLicenseWhitelist([]string{"unicorn"})
		a := ascertain.MakeLicenseAuditor(w)

		lock := AuditLocks[0]
		err := a.AuditLock(lock)
		assert.EqualError(t, err, ascertain.NonWhitelistedLicenseErr+lock.License.Kind)

		lock = AuditLocks[1]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, ascertain.NonWhitelistedLicenseErr+lock.License.Kind)

		lock = AuditLocks[2]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, ascertain.NonWhitelistedLicenseErr+lock.License.Kind)
	})
}

func TestLicenseAuditorAudit(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := ascertain.MakeLicenseWhitelist(cerealList)
		a := ascertain.MakeLicenseAuditor(w)

		verified, err := a.Audit(AuditLocks...)
		assert.Nil(t, err)
		assert.Equal(t, verified, AuditLocks)
	})
	t.Run("error on non-whitelisted licenses", func(t *testing.T) {
		w := ascertain.MakeLicenseWhitelist([]string{"fruitloops", "cheerios"})
		a := ascertain.MakeLicenseAuditor(w)
		auditErr := "detected non-whitelisted licenses. Remove or Whitelist: [\n " +
			"{\n  " +
			"\"name\": \"3\",\n  " +
			"\"version\": \"v3\",\n  " +
			"\"err\": \"non-whitelisted license: cocoapuffs\",\n  " +
			"\"license\": {\n   " +
			"\"kind\": \"cocoapuffs\",\n   " +
			"\"text\": \"\"\n  " +
			"}\n }\n]"

		verified, err := a.Audit(AuditLocks...)
		assert.EqualError(t, err, auditErr)
		assert.Equal(t, verified, AuditLocks[:2])
	})
}
