// Code generated by go-bindata.
// sources:
// api/pkg/db/migrations/sqlite/000001_users_followers.down.sql
// api/pkg/db/migrations/sqlite/000001_users_followers.up.sql
// api/pkg/db/migrations/sqlite/000002_posts_comments.down.sql
// api/pkg/db/migrations/sqlite/000002_posts_comments.up.sql
// api/pkg/db/migrations/sqlite/000003_groups_events.down.sql
// api/pkg/db/migrations/sqlite/000003_groups_events.up.sql
// api/pkg/db/migrations/sqlite/000004_messages.down.sql
// api/pkg/db/migrations/sqlite/000004_messages.up.sql
// api/pkg/db/migrations/sqlite/000005_notifications.down.sql
// api/pkg/db/migrations/sqlite/000005_notifications.up.sql
// api/pkg/db/migrations/sqlite/000006_alter_posts_followers_groups.down.sql
// api/pkg/db/migrations/sqlite/000006_alter_posts_followers_groups.up.sql
// api/pkg/db/migrations/sqlite/000007_update_events.down.sql
// api/pkg/db/migrations/sqlite/000007_update_events.up.sql
// api/pkg/db/migrations/sqlite/000008_add_seed.down.sql
// api/pkg/db/migrations/sqlite/000008_add_seed.up.sql
// DO NOT EDIT!

package database

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_users_followersDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xcb\xcf\xc9\xc9\x2f\x4f\x2d\x2a\xb6\xe6\xc2\xae\xa0\xb4\x18\x2c\xc9\x05\x08\x00\x00\xff\xff\xd8\x66\x13\x19\x3f\x00\x00\x00")

func _000001_users_followersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_users_followersDownSql,
		"000001_users_followers.down.sql",
	)
}

func _000001_users_followersDownSql() (*asset, error) {
	bytes, err := _000001_users_followersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_users_followers.down.sql", size: 63, mode: os.FileMode(420), modTime: time.Unix(1681132613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_users_followersUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x51\xd1\x8e\xaa\x30\x10\x7d\x6e\xbf\x62\x1e\x35\xe1\x0f\x7c\x42\x1d\x4c\x73\x11\xee\xad\x25\xd1\x27\x52\xa0\x6a\x63\xa5\x84\x96\x6b\xf6\xef\x37\x10\xdd\x65\xb3\xa2\x8f\x73\xe6\x9c\x33\x73\x66\x56\x1c\x43\x81\x20\xc2\x65\x8c\xc0\x22\x48\x52\x01\xb8\x67\x3b\xb1\x83\xce\xa9\xd6\xc1\x8c\x12\x5d\x01\x4b\x04\x6e\x90\xc3\x5f\xce\xb6\x21\x3f\xc0\x1f\x3c\x04\x94\x1c\x6d\x5b\xcb\xab\x02\x81\x7b\x31\x28\x93\x2c\x8e\x03\x4a\x5c\xf7\x1c\x57\x57\xa9\xcd\x4f\x14\xb2\x84\xfd\xcb\x30\xa0\xa4\x91\xce\xdd\x6c\x5b\xfd\x52\x15\xba\xf5\xe7\x4a\x7e\xc0\x3a\x14\x28\xd8\xb6\x27\xd7\xba\xbc\x7c\x8d\x08\x28\x91\x85\xed\xfc\xa3\xd0\x57\x79\x52\x79\x23\xfd\xf9\x81\x94\xad\x92\x5e\x55\xb9\xf4\x63\x13\xed\xf2\xa6\x2b\x8c\x2e\x61\x99\xa6\xf1\xf7\x4a\x6b\x8c\xc2\x2c\x16\x70\x94\xc6\x29\x3a\x5f\x50\xfa\xe6\x4a\xb9\x53\xce\x69\x5b\xbb\x57\xc7\x1a\x88\xa3\xee\x28\xa1\xb7\x17\x55\x8f\x73\x1b\xf3\x7c\xe9\xb1\x88\x44\x29\x47\xb6\x49\x7a\x7b\x98\xdd\xdd\xe7\x94\x10\xc2\x31\x42\x8e\xc9\x0a\xef\x3f\x9c\xf5\xf8\x90\xe3\x55\x90\xa3\x35\xc6\xde\x7a\xfa\xcb\x8f\xf7\x24\x5d\x9f\x26\x92\x3c\x4c\x26\xda\xb2\x2c\x55\xe3\x55\x05\x85\xb5\x66\xa8\xbd\xfe\xaf\x86\x6a\x3a\xda\x78\xe6\x64\x3e\xf2\x5c\xf5\xe6\x28\xf3\x05\xfd\x0c\x00\x00\xff\xff\x49\x41\x46\x21\x01\x03\x00\x00")

func _000001_users_followersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_users_followersUpSql,
		"000001_users_followers.up.sql",
	)
}

func _000001_users_followersUpSql() (*asset, error) {
	bytes, err := _000001_users_followersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_users_followers.up.sql", size: 769, mode: os.FileMode(420), modTime: time.Unix(1695756018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_posts_commentsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\xcf\xcd\x4d\xcd\x2b\x29\xb6\xe6\xc2\x2e\x9f\x98\x93\x93\x5f\x9e\x9a\x12\x5f\x50\x94\x59\x96\x58\x92\x1a\x5f\x90\x5f\x8c\x5b\x31\x7e\x49\x90\x09\xc9\x95\xf1\x25\x95\x05\xa9\xd6\x5c\x80\x00\x00\x00\xff\xff\xf1\xfa\xe0\xd3\x8e\x00\x00\x00")

func _000002_posts_commentsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_posts_commentsDownSql,
		"000002_posts_comments.down.sql",
	)
}

func _000002_posts_commentsDownSql() (*asset, error) {
	bytes, err := _000002_posts_commentsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_posts_comments.down.sql", size: 142, mode: os.FileMode(420), modTime: time.Unix(1692185675, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_posts_commentsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x8f\x9b\x30\x10\x3d\xdb\xbf\x62\x94\x13\x48\xf4\xd0\xf6\xd8\x13\x4d\x27\x91\x55\x42\x5a\x63\xaa\xe4\x84\x08\x8c\x52\x4b\x7c\x09\x9c\x76\xf3\xef\x57\x66\xd1\x8a\x90\x25\xc9\x1e\xf6\xe8\x99\xf7\x98\x79\x6f\x1e\x4b\x89\xbe\x42\x50\xfe\xf7\x00\x41\xac\x20\xdc\x2a\xc0\x9d\x88\x54\x04\x4d\xdd\x99\x0e\x1c\xce\x74\x0e\x22\x54\xb8\x46\x09\xbf\xa4\xd8\xf8\x72\x0f\x3f\x71\xef\x71\x76\xea\xa8\x4d\x46\x5d\x4b\x0e\xe3\x20\xf0\x38\x6b\x5a\xfd\x2f\xcd\xce\x89\x39\x37\x34\x03\x31\xda\x14\x04\x0a\x77\x6a\x5c\xcd\xea\xca\x50\x65\xae\xeb\x2d\xa5\x86\xf2\x24\x35\xf0\xc3\x57\xa8\xc4\x06\xc7\x6d\x5d\xa6\x47\x4a\x9a\xd4\xfc\xed\x99\x1e\x67\x6c\xb5\x95\x28\xd6\xa1\xdd\x15\x9c\x61\x55\x17\x38\x63\x4c\xe2\x0a\x25\x86\x4b\x8c\xc0\xd6\x3b\x47\xe7\xee\x94\x30\x11\x70\x45\x1c\xf7\x7b\xbe\xfb\x8d\x73\x7e\xcb\xce\x11\x61\x70\x55\x57\x86\x8e\xd4\x8e\x5d\x1d\x6b\xaa\xd2\x92\xc0\xd0\x93\x81\x38\x14\xbf\xe3\x0b\xbd\x39\x75\x59\xab\x1b\xa3\xeb\xaa\x87\xf4\xe3\x45\x18\xa1\x54\xd6\xeb\xed\x64\x9c\xce\x3d\xb0\x9f\x73\xf9\x1f\x3f\x88\x31\x02\xee\x7c\xf6\x60\xd1\x9c\x0e\x85\xce\x16\xae\xc7\x9d\x2f\xf6\x69\x39\x86\xfa\xf7\x57\x0f\x16\xdd\xe9\xf0\xe9\xb5\x76\x47\x5e\x5a\x14\xf5\x7f\xca\x93\x01\x9f\xdc\x4f\x8f\x45\xcc\x44\xe3\x46\xb0\x66\xaf\xfa\xf6\x51\x2f\x6f\xfa\x32\x71\x0a\xef\x57\x7d\xe8\x86\x59\x5d\x96\x54\x7d\x88\xae\xb9\xdc\x5f\x07\xfb\xce\x9f\xf0\xb0\x41\xf0\x4e\x87\x60\xb0\xe8\x39\x00\x00\xff\xff\x0e\xf6\x62\xe4\x33\x04\x00\x00")

func _000002_posts_commentsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_posts_commentsUpSql,
		"000002_posts_comments.up.sql",
	)
}

func _000002_posts_commentsUpSql() (*asset, error) {
	bytes, err := _000002_posts_commentsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_posts_comments.up.sql", size: 1075, mode: os.FileMode(420), modTime: time.Unix(1688552646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_groups_eventsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2f\xca\x2f\x2d\x88\x4f\x2d\x4b\xcd\x2b\x89\x4f\x2c\x29\x49\xcd\x4b\x49\xcc\x4b\x4e\xb5\xe6\x42\x52\x0d\x52\x0e\x55\x5d\x5a\x9c\x5a\x14\x0f\xd6\x52\x8c\x4b\x09\x92\x81\xf8\xd5\x14\x5b\x73\x71\x01\x02\x00\x00\xff\xff\x63\x95\xae\x25\x94\x00\x00\x00")

func _000003_groups_eventsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_groups_eventsDownSql,
		"000003_groups_events.down.sql",
	)
}

func _000003_groups_eventsDownSql() (*asset, error) {
	bytes, err := _000003_groups_eventsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_groups_events.down.sql", size: 148, mode: os.FileMode(420), modTime: time.Unix(1681132613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000003_groups_eventsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xc1\x6e\x83\x30\x10\x44\xcf\xf6\x57\xec\x31\x91\xf2\x07\x3d\x91\x74\x13\x59\x25\x50\x19\x57\x4a\x4e\xc8\x82\x55\xe4\xaa\x35\x08\x3b\xfd\xfe\x0a\x03\x2d\x69\x09\x8d\xda\x4b\xaf\x9e\xdd\x1d\xde\x30\x1b\x89\x91\x42\x50\xd1\x3a\x46\x10\x5b\x48\x52\x05\x78\x10\x99\xca\xe0\xd4\x54\xe7\xda\x2d\x38\x33\x25\x88\x44\xe1\x0e\x25\x3c\x4a\xb1\x8f\xe4\x11\x1e\xf0\xb8\xe2\xac\x68\x48\xfb\xaa\xc9\x47\x03\xed\x7e\xf2\x14\xc7\x2b\xce\xbc\xf1\x2f\x04\x0a\x0f\x6a\xfc\x5a\x92\x2b\x1a\x53\x7b\x53\xd9\xa0\x0d\x67\xa8\xcc\xb5\x87\xfb\x48\xa1\x12\x7b\x1c\x6f\xb0\x6d\x2a\x51\xec\x92\xd6\x14\x16\x9f\x9e\x4b\xce\x18\x93\xb8\x45\x89\xc9\x06\x33\x38\x3b\x6a\xdc\xa2\x7d\x5f\xde\x71\x3e\x03\xd6\x0e\xe6\x1d\x1d\xcc\xe1\x85\xb9\x69\xb6\xb0\x7d\x45\x7b\xae\x8c\xbd\x4e\xa3\x8b\x82\x6a\x4f\x25\xac\xd3\x34\x1e\x0b\x17\x94\xbd\xf5\x12\xf8\x77\x44\x08\x8c\x97\xf3\xc3\xf7\x7c\x5d\x18\x28\x6f\x48\xa5\x3b\x41\x6f\x64\xfd\x7c\x2c\x73\xec\x33\x91\xfd\xf0\x97\x83\x71\xee\xcd\x2b\x4d\xca\xad\xe0\x6a\x6d\x87\xcb\x37\xf7\xeb\x9f\x46\x9c\x6b\xef\xc9\x96\xda\x16\xf4\xdb\x0e\x76\x77\xa6\x35\xe3\x7a\x03\x63\x4f\xa1\x69\x7f\xa5\x1f\xcc\x26\xe9\x3f\x5a\xd3\x67\xf0\x1e\x00\x00\xff\xff\x39\x7b\xab\x95\x55\x04\x00\x00")

func _000003_groups_eventsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000003_groups_eventsUpSql,
		"000003_groups_events.up.sql",
	)
}

func _000003_groups_eventsUpSql() (*asset, error) {
	bytes, err := _000003_groups_eventsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000003_groups_events.up.sql", size: 1109, mode: os.FileMode(420), modTime: time.Unix(1691005342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_messagesDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x4d\x2d\x2e\x4e\x4c\x4f\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xc3\xcb\x53\x7a\x20\x00\x00\x00")

func _000004_messagesDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_messagesDownSql,
		"000004_messages.down.sql",
	)
}

func _000004_messagesDownSql() (*asset, error) {
	bytes, err := _000004_messagesDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_messages.down.sql", size: 32, mode: os.FileMode(420), modTime: time.Unix(1681132613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000004_messagesUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x4d\x6e\xc2\x30\x10\x85\xd7\xf6\x29\x66\x99\x48\xb9\x41\x57\x2e\x9d\x20\xab\xc1\x54\xce\x54\x82\x55\x64\xe1\x51\xea\x05\x26\xb2\xcd\xfd\x2b\xe8\x8f\x4c\x56\x6c\xe7\xd3\x37\xef\xbd\x8d\x45\x45\x08\xa4\x5e\x07\x04\xdd\x83\xd9\x13\xe0\x41\x8f\x34\xc2\x99\x73\x76\x33\xe7\x46\x8a\xe0\x41\x1b\xc2\x2d\x5a\xf8\xb0\x7a\xa7\xec\x11\xde\xf1\xd8\x49\x91\x39\x7a\x4e\x53\xc5\x6f\x0f\xcc\xe7\x30\x74\x52\x24\x3e\x85\x25\x70\x2c\x15\xef\xa4\x98\xd3\xe5\xba\x3c\x9e\x4e\x97\x58\x38\x16\x20\x3c\x50\xfd\x21\x9c\xdd\xcc\xd3\xe2\xca\xd7\x1d\xfd\x04\x96\xc9\x15\x78\x53\x84\xa4\x77\xf8\x98\xe7\x7c\xcd\x3a\x29\xfa\xbd\x45\xbd\x35\xb7\xb6\xd0\xfc\x97\x6d\x41\x0a\x61\xb1\x47\x8b\x66\x83\x23\x5c\x33\xa7\x0c\x4d\xf0\xed\xca\xa8\x17\x3c\x2d\xfd\xed\x5b\x0b\xf7\xfb\xaf\xd1\xbe\x7c\x07\x00\x00\xff\xff\xea\x36\x03\x21\x79\x01\x00\x00")

func _000004_messagesUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000004_messagesUpSql,
		"000004_messages.up.sql",
	)
}

func _000004_messagesUpSql() (*asset, error) {
	bytes, err := _000004_messagesUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000004_messages.up.sql", size: 377, mode: os.FileMode(420), modTime: time.Unix(1691005342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_notificationsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x50\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x8b\x2f\xa9\x2c\x48\x2d\xb6\xe6\x22\x42\x65\x4a\x6a\x49\x62\x66\x0e\x51\x6a\x8b\xad\x01\x01\x00\x00\xff\xff\xd5\x93\xef\x6f\x7a\x00\x00\x00")

func _000005_notificationsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_notificationsDownSql,
		"000005_notifications.down.sql",
	)
}

func _000005_notificationsDownSql() (*asset, error) {
	bytes, err := _000005_notificationsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_notifications.down.sql", size: 122, mode: os.FileMode(420), modTime: time.Unix(1681132613, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000005_notificationsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\x4f\x6f\xaa\x40\x14\xc5\xd7\xcc\xa7\xb8\x61\x05\x09\x8b\xf7\x67\xf9\x56\xe8\xbb\x1a\x52\x84\x66\x18\x1b\x5d\x11\x02\xd7\x66\x12\x3b\xd8\x99\xd1\xc6\x6f\xdf\x80\x7f\x8a\x88\xd5\x74\x7b\xcf\xb9\x67\x0e\xbf\xcb\x98\x63\x28\x10\x44\x38\x8a\x11\xa2\x09\x24\xa9\x00\x5c\x44\x99\xc8\x40\xd5\x56\xae\x64\x59\x58\x59\xab\xdc\xee\x37\x64\x3c\xe6\xc8\x0a\xa2\x44\xe0\x14\x39\x3c\xf3\x68\x16\xf2\x25\x3c\xe1\x32\x60\x00\x00\xaa\x78\x23\x10\xb8\x10\x6d\x4a\x32\x8f\xe3\xc3\x9c\x94\x95\x76\x7f\xa9\x30\xff\x1f\x63\x51\x92\x21\x17\x4d\x60\x3a\xf0\x1a\x78\xb2\x0a\xda\xd0\xe0\x18\xe1\xb3\x97\x30\x9e\x63\x06\xcc\xfb\x15\x80\xbb\xaa\xd7\xeb\xfa\x23\xd7\xf4\xbe\x25\x63\xdd\xf3\x84\xb4\x71\xfd\x80\x79\xbf\x03\x70\x5f\x75\xbd\xdd\x74\x2d\xed\xe0\xa0\xff\x39\xeb\x52\xed\xa4\xa5\x9e\xfc\x37\x00\x97\x76\xa4\x6c\x5f\xce\xdb\xa9\x71\x9b\x4f\x60\xec\x51\x82\x15\xd9\x42\xae\xef\x31\x34\xa4\x2a\xd2\x79\xc7\x72\xc9\xf2\x0a\xd3\x6d\xeb\x81\xd9\xb0\xee\x94\x9a\x0a\x4b\x55\x5e\x58\xf8\x1f\x0a\x14\xd1\x0c\xbb\xf2\x24\xe5\x18\x4d\x93\xa6\x17\x78\xe7\x4e\x3e\x30\xc7\xe1\x38\x41\x8e\xc9\x18\x33\xd8\x1a\xd2\xed\x99\xfc\xde\xc6\x50\xcb\xfe\xf2\xf0\xc1\xfd\xf6\xc7\x78\x14\xea\x3d\x9a\x9a\x4a\x92\xbb\x87\x79\x1e\x4f\x74\xdb\x6d\x88\x54\x17\xd9\xe9\x95\xa2\x6c\xf6\x61\x94\xa6\x57\xf0\x3a\x15\x7e\x86\xef\xab\xd4\xb7\x04\x8f\xb6\x13\xc3\xcf\x00\x00\x00\xff\xff\x27\x83\x99\x14\xd8\x03\x00\x00")

func _000005_notificationsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000005_notificationsUpSql,
		"000005_notifications.up.sql",
	)
}

func _000005_notificationsUpSql() (*asset, error) {
	bytes, err := _000005_notificationsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000005_notifications.up.sql", size: 984, mode: os.FileMode(420), modTime: time.Unix(1691005342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_alter_posts_followers_groupsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xc8\x2f\x2e\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc9\x2c\xc9\x49\xb5\xe6\xe2\xc2\x54\xe4\x12\xe4\x1f\x00\x53\x95\x5e\x94\x5f\x5a\x10\x9f\x99\x82\xa6\x10\x2c\x8c\xaa\x32\x33\x37\x31\x3d\x35\xbe\x20\xb1\x24\x03\x4d\x6d\x5a\x7e\x4e\x4e\x7e\x79\x6a\x11\x8a\xed\x89\xc9\x25\x99\x65\xa9\x0a\x4e\xfe\xfe\x3e\xd6\x80\x00\x00\x00\xff\xff\xd1\x71\x90\xd4\xa7\x00\x00\x00")

func _000006_alter_posts_followers_groupsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_alter_posts_followers_groupsDownSql,
		"000006_alter_posts_followers_groups.down.sql",
	)
}

func _000006_alter_posts_followers_groupsDownSql() (*asset, error) {
	bytes, err := _000006_alter_posts_followers_groupsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_alter_posts_followers_groups.down.sql", size: 167, mode: os.FileMode(420), modTime: time.Unix(1688552646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000006_alter_posts_followers_groupsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\x0a\xc2\x30\x10\x45\xf7\x39\xc5\x2c\xf5\x0c\x59\xc5\x66\x14\x21\xa6\x12\x23\xb8\x2b\xc1\xc6\x3a\x10\x49\x68\x46\xbd\xbe\x8b\x22\x58\xbb\x7e\xef\x3f\xbe\x32\x1e\x1d\x78\xb5\x31\x08\x25\x57\xae\xa0\x5d\x7b\x84\xa6\x35\xe7\x83\x05\x26\x4e\x51\x0a\xb1\xb4\x84\xd2\xfa\x6b\x0d\x63\x7e\x96\x8e\x7a\xd8\x5b\x8f\x3b\x74\x20\x1c\x6e\xd1\xa1\x6d\xf0\x34\xc1\xba\xa2\x7e\x2d\xc5\x3c\x34\x91\xdf\x10\x3d\xc2\x10\xbb\x12\xf8\x0e\x1e\x2f\xfe\x7f\x70\xcb\x29\xe5\x77\x1c\xe7\x1f\xc3\x95\xe9\x15\xe5\x27\x00\x00\xff\xff\xd2\x1d\x6f\x6b\xc9\x00\x00\x00")

func _000006_alter_posts_followers_groupsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000006_alter_posts_followers_groupsUpSql,
		"000006_alter_posts_followers_groups.up.sql",
	)
}

func _000006_alter_posts_followers_groupsUpSql() (*asset, error) {
	bytes, err := _000006_alter_posts_followers_groupsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000006_alter_posts_followers_groups.up.sql", size: 201, mode: os.FileMode(420), modTime: time.Unix(1688552646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_update_eventsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2f\xca\x2f\x2d\x88\x4f\x2d\x4b\xcd\x2b\x29\x56\xe0\x72\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc9\xcc\x4d\x2d\x2e\x48\xcc\x53\xf0\xf4\x0b\x71\x75\x77\x0d\xb2\xe6\xe2\xc2\xad\xcf\x25\xc8\x3f\x00\xa6\x51\x01\x2c\x18\x9f\x9a\x97\x12\x0f\x32\xc3\x1a\x10\x00\x00\xff\xff\xa2\x86\xb2\x2e\x6e\x00\x00\x00")

func _000007_update_eventsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_update_eventsDownSql,
		"000007_update_events.down.sql",
	)
}

func _000007_update_eventsDownSql() (*asset, error) {
	bytes, err := _000007_update_eventsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_update_events.down.sql", size: 110, mode: os.FileMode(420), modTime: time.Unix(1691005342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000007_update_eventsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2f\xca\x2f\x2d\x88\x4f\x2d\x4b\xcd\x2b\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\xc9\xcc\x4d\x2d\x2e\x48\xcc\xb3\xe6\xe2\xc2\xa9\x9e\xcb\xd1\xc5\x05\xa6\x1e\x2c\x14\x9f\x9a\x97\x12\x0f\xd2\xa9\xe0\xe2\x18\xe2\x1a\xe2\xe9\xeb\xaa\xe0\xe7\x1f\xa2\xe0\x17\xea\xe3\x63\x0d\x08\x00\x00\xff\xff\xdb\xa1\x51\x24\x76\x00\x00\x00")

func _000007_update_eventsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000007_update_eventsUpSql,
		"000007_update_events.up.sql",
	)
}

func _000007_update_eventsUpSql() (*asset, error) {
	bytes, err := _000007_update_eventsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000007_update_events.up.sql", size: 118, mode: os.FileMode(420), modTime: time.Unix(1691005342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000008_add_seedDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4e\x4d\x4d\xb1\xe6\x02\x04\x00\x00\xff\xff\xe8\xc1\xd9\x7a\x1c\x00\x00\x00")

func _000008_add_seedDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_add_seedDownSql,
		"000008_add_seed.down.sql",
	)
}

func _000008_add_seedDownSql() (*asset, error) {
	bytes, err := _000008_add_seedDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_add_seed.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1695579204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000008_add_seedUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xcb\x31\x0a\x02\x41\x0c\x46\xe1\xda\x9c\xe2\x2f\x77\xcf\xb0\xd5\xac\x66\x61\x20\x24\xb0\xc9\x88\x9d\x08\x13\x41\xb0\x9b\xfb\x83\x38\xdd\x2b\xbe\x77\x3d\xb9\x04\x23\xca\x2e\x8c\x7a\x40\x2d\xc0\x8f\xea\xe1\x18\x99\x1d\x0b\x5d\x3e\xe3\xf9\xcf\xec\xd8\xcd\x64\x0a\x6d\x22\xb8\xf1\x51\x9a\x04\xde\xaf\xef\x48\x5a\x37\xa2\xaa\xce\x67\xa0\x6a\xd8\xbc\xe9\x5e\xa4\xb1\x63\x99\x64\xdd\x7e\x01\x00\x00\xff\xff\x39\x69\x5d\xa6\x6d\x00\x00\x00")

func _000008_add_seedUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000008_add_seedUpSql,
		"000008_add_seed.up.sql",
	)
}

func _000008_add_seedUpSql() (*asset, error) {
	bytes, err := _000008_add_seedUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000008_add_seed.up.sql", size: 109, mode: os.FileMode(420), modTime: time.Unix(1695579204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"000001_users_followers.down.sql": _000001_users_followersDownSql,
	"000001_users_followers.up.sql": _000001_users_followersUpSql,
	"000002_posts_comments.down.sql": _000002_posts_commentsDownSql,
	"000002_posts_comments.up.sql": _000002_posts_commentsUpSql,
	"000003_groups_events.down.sql": _000003_groups_eventsDownSql,
	"000003_groups_events.up.sql": _000003_groups_eventsUpSql,
	"000004_messages.down.sql": _000004_messagesDownSql,
	"000004_messages.up.sql": _000004_messagesUpSql,
	"000005_notifications.down.sql": _000005_notificationsDownSql,
	"000005_notifications.up.sql": _000005_notificationsUpSql,
	"000006_alter_posts_followers_groups.down.sql": _000006_alter_posts_followers_groupsDownSql,
	"000006_alter_posts_followers_groups.up.sql": _000006_alter_posts_followers_groupsUpSql,
	"000007_update_events.down.sql": _000007_update_eventsDownSql,
	"000007_update_events.up.sql": _000007_update_eventsUpSql,
	"000008_add_seed.down.sql": _000008_add_seedDownSql,
	"000008_add_seed.up.sql": _000008_add_seedUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"000001_users_followers.down.sql": &bintree{_000001_users_followersDownSql, map[string]*bintree{}},
	"000001_users_followers.up.sql": &bintree{_000001_users_followersUpSql, map[string]*bintree{}},
	"000002_posts_comments.down.sql": &bintree{_000002_posts_commentsDownSql, map[string]*bintree{}},
	"000002_posts_comments.up.sql": &bintree{_000002_posts_commentsUpSql, map[string]*bintree{}},
	"000003_groups_events.down.sql": &bintree{_000003_groups_eventsDownSql, map[string]*bintree{}},
	"000003_groups_events.up.sql": &bintree{_000003_groups_eventsUpSql, map[string]*bintree{}},
	"000004_messages.down.sql": &bintree{_000004_messagesDownSql, map[string]*bintree{}},
	"000004_messages.up.sql": &bintree{_000004_messagesUpSql, map[string]*bintree{}},
	"000005_notifications.down.sql": &bintree{_000005_notificationsDownSql, map[string]*bintree{}},
	"000005_notifications.up.sql": &bintree{_000005_notificationsUpSql, map[string]*bintree{}},
	"000006_alter_posts_followers_groups.down.sql": &bintree{_000006_alter_posts_followers_groupsDownSql, map[string]*bintree{}},
	"000006_alter_posts_followers_groups.up.sql": &bintree{_000006_alter_posts_followers_groupsUpSql, map[string]*bintree{}},
	"000007_update_events.down.sql": &bintree{_000007_update_eventsDownSql, map[string]*bintree{}},
	"000007_update_events.up.sql": &bintree{_000007_update_eventsUpSql, map[string]*bintree{}},
	"000008_add_seed.down.sql": &bintree{_000008_add_seedDownSql, map[string]*bintree{}},
	"000008_add_seed.up.sql": &bintree{_000008_add_seedUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
