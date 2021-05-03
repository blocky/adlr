package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

var cerealList = []string{
	"fruitloops",
	"cheerios",
	"cocoapuffs",
}

var AuditLocks = []adlr.DependencyLock{
	adlr.DependencyLock{Name: "1", License: adlr.License{Kind: "fruitloops"}},
	adlr.DependencyLock{Name: "2", License: adlr.License{Kind: "cheerios"}},
	adlr.DependencyLock{Name: "3", License: adlr.License{Kind: "cocoapuffs"}},
}

func TestLicenseAuditorAuditLock(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(cerealList)
		a := adlr.MakeLicenseAuditorFromRaw(w)

		err := a.AuditLock(AuditLocks[0])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[1])
		assert.Nil(t, err)

		err = a.AuditLock(AuditLocks[2])
		assert.Nil(t, err)
	})
	t.Run("error on non-whitelist license", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw([]string{"unicorn"})
		a := adlr.MakeLicenseAuditorFromRaw(w)

		lock := AuditLocks[0]
		err := a.AuditLock(lock)
		assert.EqualError(t, err, lock.Name+": ["+lock.License.Kind+"]")

		lock = AuditLocks[1]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, lock.Name+": ["+lock.License.Kind+"]")

		lock = AuditLocks[2]
		err = a.AuditLock(lock)
		assert.EqualError(t, err, lock.Name+": ["+lock.License.Kind+"]")
	})
}

func TestLicenseAuditorAudit(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(cerealList)
		a := adlr.MakeLicenseAuditorFromRaw(w)

		err := a.Audit(AuditLocks)
		assert.Nil(t, err)
	})
	t.Run("error on non-whitelisted licenses", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw([]string{"unicorn"})
		a := adlr.MakeLicenseAuditorFromRaw(w)
		auditErr := "{\n " +
			"\"whitelist\": \"detected non-whitelisted licenses. Remove or Whitelist\",\n " +
			"\"licenses\": [\n  " +
			"\"1: [fruitloops]\",\n  " +
			"\"2: [cheerios]\",\n  " +
			"\"3: [cocoapuffs]\"\n ]\n}"

		err := a.Audit(AuditLocks)
		assert.EqualError(t, err, auditErr)
	})
}
