package packagemanager

import (
	"fmt"
	"time"

	lhgons "github.com/longhorn/go-common-libs/ns"
)

type PackageManagerType string

const (
	PackageManagerUnknown             = PackageManagerType("")
	PackageManagerApt                 = PackageManagerType("apt")
	PackageManagerYum                 = PackageManagerType("yum")
	PackageManagerZypper              = PackageManagerType("zypper")
	PackageManagerTransactionalUpdate = PackageManagerType("transactional-update")
	PackageManagerPacman              = PackageManagerType("pacman")
	// PackageManagerQlist            = PackageManagerType("qlist")
)

type PackageManager interface {
	UpdatePackageList() (string, error)
	InstallPackage(name string) (string, error)
	UninstallPackage(name string) (string, error)
	Modprobe(module string) (string, error)
	CheckModLoaded(module string) error
	StartService(name string) (string, error)
	GetServiceStatus(name string) (string, error)
	CheckPackageInstalled(name string) (string, error)
	Execute(envs []string, binary string, args []string, timeout time.Duration) (string, error)
}

func New(pkgMgrType PackageManagerType, executor *lhgons.Executor) (PackageManager, error) {
	switch pkgMgrType {
	case PackageManagerApt:
		return NewAptPackageManager(executor), nil
	case PackageManagerYum:
		return NewYumPackageManager(executor), nil
	case PackageManagerZypper:
		return NewZypperPackageManager(executor), nil
	case PackageManagerTransactionalUpdate:
		return NewTransactionalUpdatePackageManager(executor), nil
	case PackageManagerPacman:
		return NewPacmanPackageManager(executor), nil
	default:
		return nil, fmt.Errorf("unknown package manager type: %s", pkgMgrType)
	}
}
