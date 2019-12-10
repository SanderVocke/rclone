// Package api has type definitions for mediafire
package api

import (
	"time"

	"github.com/pkg/errors"
)

// Time represents represents date and time information for the
// mediafire API, by using RFC3339
type Time time.Time

// MarshalJSON turns a Time into JSON (in UTC)
func (t *Time) MarshalJSON() (out []byte, err error) {
	return nil, errors.New("not implemented")
}

// UnmarshalJSON turns JSON into a Time
func (t *Time) UnmarshalJSON(data []byte) error {
	return errors.New("not implemented")
}

// Error is returned from box when things go wrong
type Error struct {
	// FIXME
}

// Error returns a string for the error and satisfies the error interface
func (e *Error) Error() string {
	return "Unknown error: not implemented" // FIXME
}

// Check Error satisfies the error interface
var _ error = (*Error)(nil)

// ItemFields are the fields needed for FileInfo
var ItemFields = "" // FIXME

// Types of things in Item
const (
	ItemTypeFolder    = "folder"
	ItemTypeFile      = "file"
	// FIXME more types
)

// Item describes a folder or a file as returned by Get Folder Items and others.
// FIXME: This is still as in Box, update for MediaFire
type Item struct {
	Type              string  `json:"type"`
	ID                string  `json:"id"`
	SequenceID        string  `json:"sequence_id"`
	Etag              string  `json:"etag"`
	SHA1              string  `json:"sha1"`
	Name              string  `json:"name"`
	Size              float64 `json:"size"` // box returns this in xEyy format for very large numbers - see #2261
	CreatedAt         Time    `json:"created_at"`
	ModifiedAt        Time    `json:"modified_at"`
	ContentCreatedAt  Time    `json:"content_created_at"`
	ContentModifiedAt Time    `json:"content_modified_at"`
	ItemStatus        string  `json:"item_status"` // active, trashed if the file has been moved to the trash, and deleted if the file has been permanently deleted
	SharedLink        struct {
		URL    string `json:"url,omitempty"`
		Access string `json:"access,omitempty"`
	} `json:"shared_link"`
}

// ModTime returns the modification time of the item
func (i *Item) ModTime() (t time.Time) {
	return time.Unix(0,0) // FIXME implement
}

// FolderItems is returned from the GetFolderItems call
type FolderItems struct {
	// FIXME implement
}

// Parent defined the ID of the parent directory
type Parent struct {
	// FIXME implement
}

// CreateFolder is the request for Create Folder
type CreateFolder struct {
	// FIXME implement
}

// UploadFile is the request for Upload File
type UploadFile struct {
	// FIXME implement
}

// UpdateFileModTime is used in Update File Info
type UpdateFileModTime struct {
	// FIXME implement
}

// UpdateFileMove is the request for Upload File to change name and parent
type UpdateFileMove struct {
	// FIXME implement
}

// CopyFile is the request for Copy File
type CopyFile struct {
	// FIXME implement
}

// CreateSharedLink is the request for Public Link
type CreateSharedLink struct {
	// FIXME implement
}

// UploadSessionRequest is uses in Create Upload Session
type UploadSessionRequest struct {
	// FIXME implement
}

// UploadSessionResponse is returned from Create Upload Session
type UploadSessionResponse struct {
	// FIXME implement
}

// Part defines the return from upload part call which are passed to commit upload also
type Part struct {
	// FIXME implement
}

// UploadPartResponse is returned from the upload part call
type UploadPartResponse struct {
	// FIXME implement
}

// CommitUpload is used in the Commit Upload call
type CommitUpload struct {
	// FIXME implement
}

// ConfigJSON defines the shape of a box config.json
type ConfigJSON struct {
	// FIXME implement
}

// AppSettings defines the shape of the boxAppSettings within box config.json
type AppSettings struct {
	// FIXME implement
}

// AppAuth defines the shape of the appAuth within boxAppSettings in config.json
type AppAuth struct {
	// FIXME implement
}
