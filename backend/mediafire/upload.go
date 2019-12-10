// multpart upload for mediafire

package mediafire

import (
	"context"
	"io"
	"time"
	"github.com/pkg/errors"
	"github.com/rclone/rclone/backend/box/api"
	"github.com/rclone/rclone/fs/accounting"
)

// createUploadSession creates an upload session for the object
func (o *Object) createUploadSession(ctx context.Context, leaf, directoryID string, size int64) (response *api.UploadSessionResponse, err error) {
	return nil, errors.New("not implemented") // FIXME
}

// uploadPart uploads a part in an upload session
func (o *Object) uploadPart(ctx context.Context, SessionID string, offset, totalSize int64, chunk []byte, wrap accounting.WrapFn) (response *api.UploadPartResponse, err error) {
	return nil, errors.New("not implemented") // FIXME
}

// commitUpload finishes an upload session
func (o *Object) commitUpload(ctx context.Context, SessionID string, parts []api.Part, modTime time.Time, sha1sum []byte) (result *api.FolderItems, err error) {
	return nil, errors.New("not implemented") // FIXME
}

// abortUpload cancels an upload session
func (o *Object) abortUpload(ctx context.Context, SessionID string) (err error) {
	return errors.New("not implemented") // FIXME
}

// uploadMultipart uploads a file using multipart upload
func (o *Object) uploadMultipart(ctx context.Context, in io.Reader, leaf, directoryID string, size int64, modTime time.Time) (err error) {
	return errors.New("not implemented") // FIXME
}
