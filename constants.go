package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

// virDomainState
const (
	VIR_DOMAIN_NOSTATE     = C.VIR_DOMAIN_NOSTATE
	VIR_DOMAIN_RUNNING     = C.VIR_DOMAIN_RUNNING
	VIR_DOMAIN_BLOCKED     = C.VIR_DOMAIN_BLOCKED
	VIR_DOMAIN_PAUSED      = C.VIR_DOMAIN_PAUSED
	VIR_DOMAIN_SHUTDOWN    = C.VIR_DOMAIN_SHUTDOWN
	VIR_DOMAIN_CRASHED     = C.VIR_DOMAIN_CRASHED
	VIR_DOMAIN_PMSUSPENDED = C.VIR_DOMAIN_PMSUSPENDED
	VIR_DOMAIN_SHUTOFF     = C.VIR_DOMAIN_SHUTOFF
)

//virConnectListAllDomainsFlags
const (
	VIR_CONNECT_LIST_DOMAINS_ACTIVE         = C.VIR_CONNECT_LIST_DOMAINS_ACTIVE
	VIR_CONNECT_LIST_DOMAINS_INACTIVE       = C.VIR_CONNECT_LIST_DOMAINS_INACTIVE
	VIR_CONNECT_LIST_DOMAINS_PERSISTENT     = C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT
	VIR_CONNECT_LIST_DOMAINS_TRANSIENT      = C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT
	VIR_CONNECT_LIST_DOMAINS_RUNNING        = C.VIR_CONNECT_LIST_DOMAINS_RUNNING
	VIR_CONNECT_LIST_DOMAINS_PAUSED         = C.VIR_CONNECT_LIST_DOMAINS_PAUSED
	VIR_CONNECT_LIST_DOMAINS_SHUTOFF        = C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF
	VIR_CONNECT_LIST_DOMAINS_OTHER          = C.VIR_CONNECT_LIST_DOMAINS_OTHER
	VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE    = C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE
	VIR_CONNECT_LIST_DOMAINS_AUTOSTART      = C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART   = C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART
	VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT
	VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT
)

// virDomainMetadataType
const (
	VIR_DOMAIN_METADATA_DESCRIPTION = C.VIR_DOMAIN_METADATA_DESCRIPTION
	VIR_DOMAIN_METADATA_TITLE       = C.VIR_DOMAIN_METADATA_TITLE
	VIR_DOMAIN_METADATA_ELEMENT     = C.VIR_DOMAIN_METADATA_ELEMENT
)

// virDomainVcpuFlags
const (
	VIR_DOMAIN_VCPU_CONFIG  = C.VIR_DOMAIN_VCPU_CONFIG
	VIR_DOMAIN_VCPU_CURRENT = C.VIR_DOMAIN_VCPU_CURRENT
	VIR_DOMAIN_VCPU_LIVE    = C.VIR_DOMAIN_VCPU_LIVE
	VIR_DOMAIN_VCPU_MAXIMUM = C.VIR_DOMAIN_VCPU_MAXIMUM
	VIR_DOMAIN_VCPU_GUEST   = C.VIR_DOMAIN_VCPU_GUEST
)

// virDomainModificationImpact
const (
	VIR_DOMAIN_AFFECT_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_AFFECT_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_AFFECT_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
)

// virDomainMemoryModFlags
const (
	VIR_DOMAIN_MEM_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_MEM_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_MEM_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_MEM_MAXIMUM = C.VIR_DOMAIN_MEM_MAXIMUM
)

// virStoragePoolState
const (
	VIR_STORAGE_POOL_INACTIVE     = C.VIR_STORAGE_POOL_INACTIVE     // Not running
	VIR_STORAGE_POOL_BUILDING     = C.VIR_STORAGE_POOL_BUILDING     // Initializing pool,not available
	VIR_STORAGE_POOL_RUNNING      = C.VIR_STORAGE_POOL_RUNNING      // Running normally
	VIR_STORAGE_POOL_DEGRADED     = C.VIR_STORAGE_POOL_DEGRADED     // Running degraded
	VIR_STORAGE_POOL_INACCESSIBLE = C.VIR_STORAGE_POOL_INACCESSIBLE // Running,but not accessible
)

// virStoragePoolBuildFlags
const (
	VIR_STORAGE_POOL_BUILD_NEW          = C.VIR_STORAGE_POOL_BUILD_NEW          // Regular build from scratch
	VIR_STORAGE_POOL_BUILD_REPAIR       = C.VIR_STORAGE_POOL_BUILD_REPAIR       // Repair / reinitialize
	VIR_STORAGE_POOL_BUILD_RESIZE       = C.VIR_STORAGE_POOL_BUILD_RESIZE       // Extend existing pool
	VIR_STORAGE_POOL_BUILD_NO_OVERWRITE = C.VIR_STORAGE_POOL_BUILD_NO_OVERWRITE // Do not overwrite existing pool
	VIR_STORAGE_POOL_BUILD_OVERWRITE    = C.VIR_STORAGE_POOL_BUILD_OVERWRITE    // Overwrite data
)

// virDomainDestroyFlags
const (
	VIR_DOMAIN_DESTROY_DEFAULT  = C.VIR_DOMAIN_DESTROY_DEFAULT
	VIR_DOMAIN_DESTROY_GRACEFUL = C.VIR_DOMAIN_DESTROY_GRACEFUL
)

// virDomainShutdownFlags
const (
	VIR_DOMAIN_SHUTDOWN_DEFAULT        = C.VIR_DOMAIN_SHUTDOWN_DEFAULT
	VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN = C.VIR_DOMAIN_SHUTDOWN_ACPI_POWER_BTN
	VIR_DOMAIN_SHUTDOWN_GUEST_AGENT    = C.VIR_DOMAIN_SHUTDOWN_GUEST_AGENT
	VIR_DOMAIN_SHUTDOWN_INITCTL        = C.VIR_DOMAIN_SHUTDOWN_INITCTL
	VIR_DOMAIN_SHUTDOWN_SIGNAL         = C.VIR_DOMAIN_SHUTDOWN_SIGNAL
)

// virDomainAttachDeviceFlags
const (
	VIR_DOMAIN_DEVICE_MODIFY_CONFIG  = C.VIR_DOMAIN_AFFECT_CONFIG
	VIR_DOMAIN_DEVICE_MODIFY_CURRENT = C.VIR_DOMAIN_AFFECT_CURRENT
	VIR_DOMAIN_DEVICE_MODIFY_LIVE    = C.VIR_DOMAIN_AFFECT_LIVE
	VIR_DOMAIN_DEVICE_MODIFY_FORCE   = C.VIR_DOMAIN_DEVICE_MODIFY_FORCE
)

// virStorageVolCreateFlags
const (
	VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA = C.VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA
)

// virStorageVolDeleteFlags
const (
	VIR_STORAGE_VOL_DELETE_NORMAL = C.VIR_STORAGE_VOL_DELETE_NORMAL // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED = C.VIR_STORAGE_VOL_DELETE_ZEROED // Clear all data to zeros (slow)
)

// virStorageVolResizeFlags
const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = C.VIR_STORAGE_VOL_RESIZE_ALLOCATE // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = C.VIR_STORAGE_VOL_RESIZE_DELTA    // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = C.VIR_STORAGE_VOL_RESIZE_SHRINK   // allow decrease in capacity
)

// virStorageVolType
const (
	VIR_STORAGE_VOL_FILE    = C.VIR_STORAGE_VOL_FILE    // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK   = C.VIR_STORAGE_VOL_BLOCK   // Block based volumes
	VIR_STORAGE_VOL_DIR     = C.VIR_STORAGE_VOL_DIR     // Directory-passthrough based volume
	VIR_STORAGE_VOL_NETWORK = C.VIR_STORAGE_VOL_NETWORK //Network volumes like RBD (RADOS Block Device)
	VIR_STORAGE_VOL_NETDIR  = C.VIR_STORAGE_VOL_NETDIR  // Network accessible directory that can contain other network volumes
)

// virStorageVolWipeAlgorithm
const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = C.VIR_STORAGE_VOL_WIPE_ALG_ZERO       // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = C.VIR_STORAGE_VOL_WIPE_ALG_NNSA       // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = C.VIR_STORAGE_VOL_WIPE_ALG_DOD        // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = C.VIR_STORAGE_VOL_WIPE_ALG_BSI        // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = C.VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = C.VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = C.VIR_STORAGE_VOL_WIPE_ALG_RANDOM     // 1-pass random
)

// virSecretUsageType
const (
	VIR_SECRET_USAGE_TYPE_NONE   = C.VIR_SECRET_USAGE_TYPE_NONE
	VIR_SECRET_USAGE_TYPE_VOLUME = C.VIR_SECRET_USAGE_TYPE_VOLUME
	VIR_SECRET_USAGE_TYPE_CEPH   = C.VIR_SECRET_USAGE_TYPE_CEPH
	VIR_SECRET_USAGE_TYPE_ISCSI  = C.VIR_SECRET_USAGE_TYPE_ISCSI
)

// virConnectListAllNetworksFlags
const (
	VIR_CONNECT_LIST_NETWORKS_INACTIVE     = C.VIR_CONNECT_LIST_NETWORKS_INACTIVE
	VIR_CONNECT_LIST_NETWORKS_ACTIVE       = C.VIR_CONNECT_LIST_NETWORKS_ACTIVE
	VIR_CONNECT_LIST_NETWORKS_PERSISTENT   = C.VIR_CONNECT_LIST_NETWORKS_PERSISTENT
	VIR_CONNECT_LIST_NETWORKS_TRANSIENT    = C.VIR_CONNECT_LIST_NETWORKS_TRANSIENT
	VIR_CONNECT_LIST_NETWORKS_AUTOSTART    = C.VIR_CONNECT_LIST_NETWORKS_AUTOSTART
	VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART = C.VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART
)

// virConnectListAllStoragePoolsFlags
const (
	VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE     = C.VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE       = C.VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE
	VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT   = C.VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT
	VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT    = C.VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT
	VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART    = C.VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART
	VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART = C.VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART
	VIR_CONNECT_LIST_STORAGE_POOLS_DIR          = C.VIR_CONNECT_LIST_STORAGE_POOLS_DIR
	VIR_CONNECT_LIST_STORAGE_POOLS_FS           = C.VIR_CONNECT_LIST_STORAGE_POOLS_FS
	VIR_CONNECT_LIST_STORAGE_POOLS_NETFS        = C.VIR_CONNECT_LIST_STORAGE_POOLS_NETFS
	VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL      = C.VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL
	VIR_CONNECT_LIST_STORAGE_POOLS_DISK         = C.VIR_CONNECT_LIST_STORAGE_POOLS_DISK
	VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI        = C.VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_SCSI         = C.VIR_CONNECT_LIST_STORAGE_POOLS_SCSI
	VIR_CONNECT_LIST_STORAGE_POOLS_MPATH        = C.VIR_CONNECT_LIST_STORAGE_POOLS_MPATH
	VIR_CONNECT_LIST_STORAGE_POOLS_RBD          = C.VIR_CONNECT_LIST_STORAGE_POOLS_RBD
	VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG     = C.VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG
	VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER      = C.VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER
)

// virStreamFlags
const (
	VIR_STREAM_NONBLOCK = C.VIR_STREAM_NONBLOCK
)

// virKeycodeSet
const (
	VIR_KEYCODE_SET_LINUX  = C.VIR_KEYCODE_SET_LINUX
	VIR_KEYCODE_SET_XT     = C.VIR_KEYCODE_SET_XT
	VIR_KEYCODE_SET_ATSET1 = C.VIR_KEYCODE_SET_ATSET1
	VIR_KEYCODE_SET_ATSET2 = C.VIR_KEYCODE_SET_ATSET2
	VIR_KEYCODE_SET_ATSET3 = C.VIR_KEYCODE_SET_ATSET3
	VIR_KEYCODE_SET_OSX    = C.VIR_KEYCODE_SET_OSX
	VIR_KEYCODE_SET_XT_KBD = C.VIR_KEYCODE_SET_XT_KBD
	VIR_KEYCODE_SET_USB    = C.VIR_KEYCODE_SET_USB
	VIR_KEYCODE_SET_WIN32  = C.VIR_KEYCODE_SET_WIN32
	VIR_KEYCODE_SET_RFB    = C.VIR_KEYCODE_SET_RFB
)

// virDomainCreateFlags
const (
	VIR_DOMAIN_NONE               = C.VIR_DOMAIN_NONE
	VIR_DOMAIN_START_PAUSED       = C.VIR_DOMAIN_START_PAUSED
	VIR_DOMAIN_START_AUTODESTROY  = C.VIR_DOMAIN_START_AUTODESTROY
	VIR_DOMAIN_START_BYPASS_CACHE = C.VIR_DOMAIN_START_BYPASS_CACHE
	VIR_DOMAIN_START_FORCE_BOOT   = C.VIR_DOMAIN_START_FORCE_BOOT
)

const VIR_DOMAIN_MEMORY_PARAM_UNLIMITED = C.VIR_DOMAIN_MEMORY_PARAM_UNLIMITED

// virDomainEventID
const (
	// event parameter in the callback is of type DomainLifecycleEvent
	VIR_DOMAIN_EVENT_ID_LIFECYCLE = C.VIR_DOMAIN_EVENT_ID_LIFECYCLE

	// event parameter in the callback is nil
	VIR_DOMAIN_EVENT_ID_REBOOT = C.VIR_DOMAIN_EVENT_ID_REBOOT

	// event parameter in the callback is of type DomainRTCChangeEvent
	VIR_DOMAIN_EVENT_ID_RTC_CHANGE = C.VIR_DOMAIN_EVENT_ID_RTC_CHANGE

	// event parameter in the callback is of type DomainWatchdogEvent
	VIR_DOMAIN_EVENT_ID_WATCHDOG = C.VIR_DOMAIN_EVENT_ID_WATCHDOG

	// event parameter in the callback is of type DomainIOErrorEvent
	VIR_DOMAIN_EVENT_ID_IO_ERROR = C.VIR_DOMAIN_EVENT_ID_IO_ERROR

	// event parameter in the callback is of type DomainGraphicsEvent
	VIR_DOMAIN_EVENT_ID_GRAPHICS = C.VIR_DOMAIN_EVENT_ID_GRAPHICS

	// virConnectDomainEventIOErrorReasonCallback
	VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON = C.VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON

	// event parameter in the callback is nil
	VIR_DOMAIN_EVENT_ID_CONTROL_ERROR = C.VIR_DOMAIN_EVENT_ID_CONTROL_ERROR

	// event parameter in the callback is of type DomainBlockJobEvent
	VIR_DOMAIN_EVENT_ID_BLOCK_JOB = C.VIR_DOMAIN_EVENT_ID_BLOCK_JOB

	// event parameter in the callback is of type DomainDiskChangeEvent
	VIR_DOMAIN_EVENT_ID_DISK_CHANGE = C.VIR_DOMAIN_EVENT_ID_DISK_CHANGE

	// event parameter in the callback is of type DomainTrayChangeEvent
	VIR_DOMAIN_EVENT_ID_TRAY_CHANGE = C.VIR_DOMAIN_EVENT_ID_TRAY_CHANGE

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMWAKEUP = C.VIR_DOMAIN_EVENT_ID_PMWAKEUP

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMSUSPEND = C.VIR_DOMAIN_EVENT_ID_PMSUSPEND

	// event parameter in the callback is of type DomainBalloonChangeEvent
	VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE = C.VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE

	// event parameter in the callback is of type DomainReasonEvent
	VIR_DOMAIN_EVENT_ID_PMSUSPEND_DISK = C.VIR_DOMAIN_EVENT_ID_PMSUSPEND_DISK

	// event parameter in the callback is of type DomainDeviceRemovedEvent
	VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED = C.VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED

	// TODO Post 1.2.4, enable later
	// event parameter in the callback is of type DomainBlockJobEvent
	// VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2 = C.VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2
)

// virDomainEventType
const (
	VIR_DOMAIN_EVENT_DEFINED     = C.VIR_DOMAIN_EVENT_DEFINED
	VIR_DOMAIN_EVENT_UNDEFINED   = C.VIR_DOMAIN_EVENT_UNDEFINED
	VIR_DOMAIN_EVENT_STARTED     = C.VIR_DOMAIN_EVENT_STARTED
	VIR_DOMAIN_EVENT_SUSPENDED   = C.VIR_DOMAIN_EVENT_SUSPENDED
	VIR_DOMAIN_EVENT_RESUMED     = C.VIR_DOMAIN_EVENT_RESUMED
	VIR_DOMAIN_EVENT_STOPPED     = C.VIR_DOMAIN_EVENT_STOPPED
	VIR_DOMAIN_EVENT_SHUTDOWN    = C.VIR_DOMAIN_EVENT_SHUTDOWN
	VIR_DOMAIN_EVENT_PMSUSPENDED = C.VIR_DOMAIN_EVENT_PMSUSPENDED
	VIR_DOMAIN_EVENT_CRASHED     = C.VIR_DOMAIN_EVENT_CRASHED
)

// virDomainEventWatchdogAction:
// The action that is to be taken due to the watchdog device firing
const (
	// No action, watchdog ignored
	VIR_DOMAIN_EVENT_WATCHDOG_NONE = C.VIR_DOMAIN_EVENT_WATCHDOG_NONE

	// Guest CPUs are paused
	VIR_DOMAIN_EVENT_WATCHDOG_PAUSE = C.VIR_DOMAIN_EVENT_WATCHDOG_PAUSE

	// Guest CPUs are reset
	VIR_DOMAIN_EVENT_WATCHDOG_RESET = C.VIR_DOMAIN_EVENT_WATCHDOG_RESET

	// Guest is forcibly powered off
	VIR_DOMAIN_EVENT_WATCHDOG_POWEROFF = C.VIR_DOMAIN_EVENT_WATCHDOG_POWEROFF

	// Guest is requested to gracefully shutdown
	VIR_DOMAIN_EVENT_WATCHDOG_SHUTDOWN = C.VIR_DOMAIN_EVENT_WATCHDOG_SHUTDOWN

	// No action, a debug message logged
	VIR_DOMAIN_EVENT_WATCHDOG_DEBUG = C.VIR_DOMAIN_EVENT_WATCHDOG_DEBUG
)

// virDomainEventIOErrorAction
// The action that is to be taken due to an IO error occurring
const (
	// No action, IO error ignored
	VIR_DOMAIN_EVENT_IO_ERROR_NONE = C.VIR_DOMAIN_EVENT_IO_ERROR_NONE

	// Guest CPUs are paused
	VIR_DOMAIN_EVENT_IO_ERROR_PAUSE = C.VIR_DOMAIN_EVENT_IO_ERROR_PAUSE

	// IO error reported to guest OS
	VIR_DOMAIN_EVENT_IO_ERROR_REPORT = C.VIR_DOMAIN_EVENT_IO_ERROR_REPORT
)

// virDomainEventGraphicsPhase
// The phase of the graphics client connection
const (
	// Initial socket connection established
	VIR_DOMAIN_EVENT_GRAPHICS_CONNECT = C.VIR_DOMAIN_EVENT_GRAPHICS_CONNECT

	// Authentication & setup completed
	VIR_DOMAIN_EVENT_GRAPHICS_INITIALIZE = C.VIR_DOMAIN_EVENT_GRAPHICS_INITIALIZE

	// Final socket disconnection
	VIR_DOMAIN_EVENT_GRAPHICS_DISCONNECT = C.VIR_DOMAIN_EVENT_GRAPHICS_DISCONNECT
)

// virDomainEventGraphicsAddressType
const (
	// IPv4 address
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV4 = C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV4

	// IPv6 address
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV6 = C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_IPV6

	// UNIX socket path
	VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_UNIX = C.VIR_DOMAIN_EVENT_GRAPHICS_ADDRESS_UNIX
)

// virDomainBlockJobType
const (
	// Placeholder
	VIR_DOMAIN_BLOCK_JOB_TYPE_UNKNOWN = C.VIR_DOMAIN_BLOCK_JOB_TYPE_UNKNOWN

	// Block Pull (virDomainBlockPull, or virDomainBlockRebase without
	// flags), job ends on completion
	VIR_DOMAIN_BLOCK_JOB_TYPE_PULL = C.VIR_DOMAIN_BLOCK_JOB_TYPE_PULL

	// Block Copy (virDomainBlockCopy, or virDomainBlockRebase with
	// flags), job exists as long as mirroring is active
	VIR_DOMAIN_BLOCK_JOB_TYPE_COPY = C.VIR_DOMAIN_BLOCK_JOB_TYPE_COPY

	// Block Commit (virDomainBlockCommit without flags), job ends on
	// completion
	VIR_DOMAIN_BLOCK_JOB_TYPE_COMMIT = C.VIR_DOMAIN_BLOCK_JOB_TYPE_COMMIT

	// TODO Post 1.2.4, enable later
	// Active Block Commit (virDomainBlockCommit with flags), job
	// exists as long as sync is active
	// VIR_DOMAIN_BLOCK_JOB_TYPE_ACTIVE_COMMIT = C.VIR_DOMAIN_BLOCK_JOB_TYPE_ACTIVE_COMMIT
)

// virConnectDomainEventBlockJobStatus
const (
	VIR_DOMAIN_BLOCK_JOB_COMPLETED = C.VIR_DOMAIN_BLOCK_JOB_COMPLETED
	VIR_DOMAIN_BLOCK_JOB_FAILED    = C.VIR_DOMAIN_BLOCK_JOB_FAILED
	VIR_DOMAIN_BLOCK_JOB_CANCELED  = C.VIR_DOMAIN_BLOCK_JOB_CANCELED
	VIR_DOMAIN_BLOCK_JOB_READY     = C.VIR_DOMAIN_BLOCK_JOB_READY
)

// virConnectDomainEventDiskChangeReason
const (
	// OldSrcPath is set
	VIR_DOMAIN_EVENT_DISK_CHANGE_MISSING_ON_START = C.VIR_DOMAIN_EVENT_DISK_CHANGE_MISSING_ON_START
	VIR_DOMAIN_EVENT_DISK_DROP_MISSING_ON_START   = C.VIR_DOMAIN_EVENT_DISK_DROP_MISSING_ON_START
)

// virConnectDomainEventTrayChangeReason
const (
	VIR_DOMAIN_EVENT_TRAY_CHANGE_OPEN  = C.VIR_DOMAIN_EVENT_TRAY_CHANGE_OPEN
	VIR_DOMAIN_EVENT_TRAY_CHANGE_CLOSE = C.VIR_DOMAIN_EVENT_TRAY_CHANGE_CLOSE
)

// virDomainRunningReason
const (
	VIR_DOMAIN_RUNNING_UNKNOWN            = C.VIR_DOMAIN_RUNNING_UNKNOWN
	VIR_DOMAIN_RUNNING_BOOTED             = C.VIR_DOMAIN_RUNNING_BOOTED             /* normal startup from boot */
	VIR_DOMAIN_RUNNING_MIGRATED           = C.VIR_DOMAIN_RUNNING_MIGRATED           /* migrated from another host */
	VIR_DOMAIN_RUNNING_RESTORED           = C.VIR_DOMAIN_RUNNING_RESTORED           /* restored from a state file */
	VIR_DOMAIN_RUNNING_FROM_SNAPSHOT      = C.VIR_DOMAIN_RUNNING_FROM_SNAPSHOT      /* restored from snapshot */
	VIR_DOMAIN_RUNNING_UNPAUSED           = C.VIR_DOMAIN_RUNNING_UNPAUSED           /* returned from paused state */
	VIR_DOMAIN_RUNNING_MIGRATION_CANCELED = C.VIR_DOMAIN_RUNNING_MIGRATION_CANCELED /* returned from migration */
	VIR_DOMAIN_RUNNING_SAVE_CANCELED      = C.VIR_DOMAIN_RUNNING_SAVE_CANCELED      /* returned from failed save process */
	VIR_DOMAIN_RUNNING_WAKEUP             = C.VIR_DOMAIN_RUNNING_WAKEUP             /* returned from pmsuspended due to wakeup event */
	VIR_DOMAIN_RUNNING_CRASHED            = C.VIR_DOMAIN_RUNNING_CRASHED            /* resumed from crashed */
)

// virDomainPausedReason
const (
	VIR_DOMAIN_PAUSED_UNKNOWN       = C.VIR_DOMAIN_PAUSED_UNKNOWN       /* the reason is unknown */
	VIR_DOMAIN_PAUSED_USER          = C.VIR_DOMAIN_PAUSED_USER          /* paused on user request */
	VIR_DOMAIN_PAUSED_MIGRATION     = C.VIR_DOMAIN_PAUSED_MIGRATION     /* paused for offline migration */
	VIR_DOMAIN_PAUSED_SAVE          = C.VIR_DOMAIN_PAUSED_SAVE          /* paused for save */
	VIR_DOMAIN_PAUSED_DUMP          = C.VIR_DOMAIN_PAUSED_DUMP          /* paused for offline core dump */
	VIR_DOMAIN_PAUSED_IOERROR       = C.VIR_DOMAIN_PAUSED_IOERROR       /* paused due to a disk I/O error */
	VIR_DOMAIN_PAUSED_WATCHDOG      = C.VIR_DOMAIN_PAUSED_WATCHDOG      /* paused due to a watchdog event */
	VIR_DOMAIN_PAUSED_FROM_SNAPSHOT = C.VIR_DOMAIN_PAUSED_FROM_SNAPSHOT /* paused after restoring from snapshot */
	VIR_DOMAIN_PAUSED_SHUTTING_DOWN = C.VIR_DOMAIN_PAUSED_SHUTTING_DOWN /* paused during shutdown process */
	VIR_DOMAIN_PAUSED_SNAPSHOT      = C.VIR_DOMAIN_PAUSED_SNAPSHOT      /* paused while creating a snapshot */
	VIR_DOMAIN_PAUSED_CRASHED       = C.VIR_DOMAIN_PAUSED_CRASHED       /* paused due to a guest crash */
)

// virDomainXMLFlags
const (
	VIR_DOMAIN_XML_SECURE     = C.VIR_DOMAIN_XML_SECURE     /* dump security sensitive information too */
	VIR_DOMAIN_XML_INACTIVE   = C.VIR_DOMAIN_XML_INACTIVE   /* dump inactive domain information */
	VIR_DOMAIN_XML_UPDATE_CPU = C.VIR_DOMAIN_XML_UPDATE_CPU /* update guest CPU requirements according to host CPU */
	VIR_DOMAIN_XML_MIGRATABLE = C.VIR_DOMAIN_XML_MIGRATABLE /* dump XML suitable for migration */
)

/*
 * QMP has two different kinds of ways to talk to QEMU. One is legacy (HMP,
 * or 'human' monitor protocol. The default is QMP, which is all-JSON.
 *
 * QMP json commands are of the format:
 * 	{"execute" : "query-cpus"}
 *
 * whereas the same command in 'HMP' would be:
 *	'info cpus'
 */
const (
	VIR_DOMAIN_QEMU_MONITOR_COMMAND_DEFAULT = 0
	VIR_DOMAIN_QEMU_MONITOR_COMMAND_HMP     = (1 << 0)
)

// virDomainEventDefinedDetailType
const (
	VIR_DOMAIN_EVENT_DEFINED_ADDED   = C.VIR_DOMAIN_EVENT_DEFINED_ADDED
	VIR_DOMAIN_EVENT_DEFINED_UPDATED = C.VIR_DOMAIN_EVENT_DEFINED_UPDATED
)

// virDomainEventUndefinedDetailType
const (
	VIR_DOMAIN_EVENT_UNDEFINED_REMOVED = C.VIR_DOMAIN_EVENT_UNDEFINED_REMOVED
)

// virDomainEventStartedDetailType
const (
	VIR_DOMAIN_EVENT_STARTED_BOOTED        = C.VIR_DOMAIN_EVENT_STARTED_BOOTED
	VIR_DOMAIN_EVENT_STARTED_MIGRATED      = C.VIR_DOMAIN_EVENT_STARTED_MIGRATED
	VIR_DOMAIN_EVENT_STARTED_RESTORED      = C.VIR_DOMAIN_EVENT_STARTED_RESTORED
	VIR_DOMAIN_EVENT_STARTED_FROM_SNAPSHOT = C.VIR_DOMAIN_EVENT_STARTED_FROM_SNAPSHOT
	VIR_DOMAIN_EVENT_STARTED_WAKEUP        = C.VIR_DOMAIN_EVENT_STARTED_WAKEUP
)

// virDomainEventSuspendedDetailType
const (
	VIR_DOMAIN_EVENT_SUSPENDED_PAUSED        = C.VIR_DOMAIN_EVENT_SUSPENDED_PAUSED
	VIR_DOMAIN_EVENT_SUSPENDED_MIGRATED      = C.VIR_DOMAIN_EVENT_SUSPENDED_MIGRATED
	VIR_DOMAIN_EVENT_SUSPENDED_IOERROR       = C.VIR_DOMAIN_EVENT_SUSPENDED_IOERROR
	VIR_DOMAIN_EVENT_SUSPENDED_WATCHDOG      = C.VIR_DOMAIN_EVENT_SUSPENDED_WATCHDOG
	VIR_DOMAIN_EVENT_SUSPENDED_RESTORED      = C.VIR_DOMAIN_EVENT_SUSPENDED_RESTORED
	VIR_DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT = C.VIR_DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT
	VIR_DOMAIN_EVENT_SUSPENDED_API_ERROR     = C.VIR_DOMAIN_EVENT_SUSPENDED_API_ERROR
)

// virDomainEventResumedDetailType
const (
	VIR_DOMAIN_EVENT_RESUMED_UNPAUSED      = C.VIR_DOMAIN_EVENT_RESUMED_UNPAUSED
	VIR_DOMAIN_EVENT_RESUMED_MIGRATED      = C.VIR_DOMAIN_EVENT_RESUMED_MIGRATED
	VIR_DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT = C.VIR_DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT
)

// virDomainEventStoppedDetailType
const (
	VIR_DOMAIN_EVENT_STOPPED_SHUTDOWN      = C.VIR_DOMAIN_EVENT_STOPPED_SHUTDOWN
	VIR_DOMAIN_EVENT_STOPPED_DESTROYED     = C.VIR_DOMAIN_EVENT_STOPPED_DESTROYED
	VIR_DOMAIN_EVENT_STOPPED_CRASHED       = C.VIR_DOMAIN_EVENT_STOPPED_CRASHED
	VIR_DOMAIN_EVENT_STOPPED_MIGRATED      = C.VIR_DOMAIN_EVENT_STOPPED_MIGRATED
	VIR_DOMAIN_EVENT_STOPPED_SAVED         = C.VIR_DOMAIN_EVENT_STOPPED_SAVED
	VIR_DOMAIN_EVENT_STOPPED_FAILED        = C.VIR_DOMAIN_EVENT_STOPPED_FAILED
	VIR_DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT = C.VIR_DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT
)

// virDomainEventShutdownDetailType
const (
	VIR_DOMAIN_EVENT_SHUTDOWN_FINISHED = C.VIR_DOMAIN_EVENT_SHUTDOWN_FINISHED
)

// virDomainMemoryStatTags
const (
	VIR_DOMAIN_MEMORY_STAT_LAST           = C.VIR_DOMAIN_MEMORY_STAT_NR
	VIR_DOMAIN_MEMORY_STAT_SWAP_IN        = C.VIR_DOMAIN_MEMORY_STAT_SWAP_IN
	VIR_DOMAIN_MEMORY_STAT_SWAP_OUT       = C.VIR_DOMAIN_MEMORY_STAT_SWAP_OUT
	VIR_DOMAIN_MEMORY_STAT_MAJOR_FAULT    = C.VIR_DOMAIN_MEMORY_STAT_MAJOR_FAULT
	VIR_DOMAIN_MEMORY_STAT_MINOR_FAULT    = C.VIR_DOMAIN_MEMORY_STAT_MINOR_FAULT
	VIR_DOMAIN_MEMORY_STAT_UNUSED         = C.VIR_DOMAIN_MEMORY_STAT_UNUSED
	VIR_DOMAIN_MEMORY_STAT_AVAILABLE      = C.VIR_DOMAIN_MEMORY_STAT_AVAILABLE
	VIR_DOMAIN_MEMORY_STAT_ACTUAL_BALLOON = C.VIR_DOMAIN_MEMORY_STAT_ACTUAL_BALLOON
	VIR_DOMAIN_MEMORY_STAT_RSS            = C.VIR_DOMAIN_MEMORY_STAT_RSS
	VIR_DOMAIN_MEMORY_STAT_NR             = C.VIR_DOMAIN_MEMORY_STAT_NR
)

// virDomainCPUStatsTags
const (
	VIR_DOMAIN_CPU_STATS_CPUTIME    = C.VIR_DOMAIN_CPU_STATS_CPUTIME
	VIR_DOMAIN_CPU_STATS_SYSTEMTIME = C.VIR_DOMAIN_CPU_STATS_SYSTEMTIME
	VIR_DOMAIN_CPU_STATS_USERTIME   = C.VIR_DOMAIN_CPU_STATS_USERTIME
	VIR_DOMAIN_CPU_STATS_VCPUTIME   = C.VIR_DOMAIN_CPU_STATS_VCPUTIME
)

// virDomainInterfaceAddressesSource
const (
	VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE = 0
	VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT = 1
)

// virIPAddrType
const (
	VIR_IP_ADDR_TYPE_IPV4 = 0
	VIR_IP_ADDR_TYPE_IPV6 = 1
)
