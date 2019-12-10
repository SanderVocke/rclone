// Test MediaFire filesystem interface
package mediafire_test

import (
	"testing"

	"github.com/rclone/rclone/backend/mediafire"
	"github.com/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestMediaFire:",
		NilObject:  (*mediafire.Object)(nil),
	})
}
