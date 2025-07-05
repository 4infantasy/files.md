sudo -u www-data HOME=/var/dropbox /var/dropbox.py status

sudo mount --bind -o nosymfollow  /var/dropbox/Dropbox/<dir> /app/storage/<id>

```go
	// Dropbox sometimes changes ctime for new files, for some reason.
	// So we have to query consequently.
	time.Sleep(100 * time.Millisecond)
	serverLastModified, err = userFS.Ctime(fs.DirRoot, path)
	logSync(fmt.Sprintf("Final server timestamp for '%s': %d", path, serverLastModified), r)
	serverLastModified, err = userFS.Ctime(fs.DirRoot, path)
	logSync(fmt.Sprintf("Final server timestamp for '%s': %d", path, serverLastModified), r)
```

Why Dropbox changes ctime on new files:
Metadata sync - Dropbox stores additional metadata (sync status, version info)
Extended attributes - Dropbox adds xattrs to track sync state
Permission changes - Dropbox might adjust permissions for sharing

Could be any of these.