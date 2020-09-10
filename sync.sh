#!/usr/bin/env sh

cat hosts.txt | while read s
do
	rsync --exclude ".git" -av $1 ${s}:~/$2 &
	#rsync が入っていなかったら(install.shで入れるのでその後はrsyncでいい)
	#scp -r $1 ${s}:~/$2 &
done

wait
exit 0
