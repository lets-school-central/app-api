#!/bin/bash
set -e

if [ -f /data/data.db ]; then
	echo "Database already exists, skipping restore"
else
	echo "No database found, restoring from replica if exists"
	litestream restore -v -if-replica-exists /data/data.db
fi

exec litestream replicate -exec "/usr/local/bin/lets-school-central-backend serve --http=0.0.0.0:${PORT} --dir=/data"
