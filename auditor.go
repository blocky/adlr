package adlr

import "errors"

const NonWhitelistedLicenseErr = "non-whitelisted license: "

type LicenseAuditor struct {
	whitelist LicenseWhitelist
}

func MakeLicenseAuditor() LicenseAuditor {
	whitelist := MakeLicenseWhitelist()
	return MakeLicenseAuditorFromRaw(whitelist)
}

func MakeLicenseAuditorFromRaw(
	whitelist LicenseWhitelist,
) LicenseAuditor {
	return LicenseAuditor{whitelist}
}

func (auditor LicenseAuditor) Audit(
	locks []DependencyLock,
) error {
	var auditErrs []DependencyLock

	for _, lock := range locks {
		err := auditor.AuditLock(lock)
		if err != nil {
			lock.AddErrStr(err.Error())
			auditErrs = append(auditErrs, lock)
		}
	}
	if len(auditErrs) != 0 {
		return &LicenseAuditError{auditErrs}
	}
	return nil
}

func (auditor LicenseAuditor) AuditLock(
	lock DependencyLock,
) error {
	var license = lock.License.Kind

	if !auditor.whitelist.Find(license) {
		return errors.New(NonWhitelistedLicenseErr + license)
	}
	return nil
}
