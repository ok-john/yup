
# yup

containers & privilage stuff

## Enter
```
make enter
```

## Run
```
make run
```

## How to
1. Be on on ubuntu, as an unprivelaged user
2. git clone this repository and enter its directory (ie; not in the sudo group)
3. Run `make end-to-end` to pull down a rootfs (ubuntu focal latest, you can remove after, it goes to isolated directory)

Then `make enter`, you'll see the following:
```bash
2021/06/22 15:29:14     +main(uid: 1002, gid: 1003, ppid:22473, pid: 22867)
2021/06/22 15:29:14         ^->+run(uid: 1002, gid: 1003, ppid:22473, pid: 22867)
2021/06/22 15:29:14     +main(uid: 0, gid: 0, ppid:0, pid: 1)
2021/06/22 15:29:14         ^->+child(uid: 0, gid: 0, ppid:0, pid: 1)
$ root@hostname
```
Check the id of your user 
```bash
$ id
> uid=0(root) gid=0(root) groups=0(root),65534(nogroup)
```

Look at your filesystem
```bash
$ ls -la

total 68
drwxr-xr-x  18 root   root    4096 Jun 22 22:35 .
drwxr-xr-x  18 root   root    4096 Jun 22 22:35 ..
lrwxrwxrwx   1 root   root       7 Jun 21 21:38 bin -> usr/bin
drwxr-xr-x   2 root   root    4096 Jun 21 21:44 boot
drwxr-xr-x   5 root   root    4096 Jun 21 21:40 dev
drwxr-xr-x  90 root   root    4096 Jun 21 21:44 etc
drwxr-xr-x   2 root   root    4096 Apr 15  2020 home
lrwxrwxrwx   1 root   root       7 Jun 21 21:38 lib -> usr/lib
lrwxrwxrwx   1 root   root       9 Jun 21 21:38 lib32 -> usr/lib32
lrwxrwxrwx   1 root   root       9 Jun 21 21:38 lib64 -> usr/lib64
lrwxrwxrwx   1 root   root      10 Jun 21 21:38 libx32 -> usr/libx32
drwxr-xr-x   2 root   root    4096 Jun 21 21:38 media
drwxr-xr-x   2 root   root    4096 Jun 21 21:38 mnt
drwxr-xr-x   2 root   root    4096 Jun 21 21:38 opt
dr-xr-xr-x 326 nobody nogroup    0 Jun 22 22:35 proc
drwx------   2 root   root    4096 Jun 21 21:40 root
drwxr-xr-x   3 root   root    4096 Jun 21 21:41 run
lrwxrwxrwx   1 root   root       8 Jun 21 21:38 sbin -> usr/sbin
drwxr-xr-x   6 root   root    4096 Jun 21 21:41 snap
drwxr-xr-x   2 root   root    4096 Jun 21 21:38 srv
drwxr-xr-x   2 root   root    4096 Apr 15  2020 sys
drwxr-xr-x   2 root   root    4096 Jun 21 21:41 tmp
drwxr-xr-x  15 root   root    4096 Jun 21 21:39 usr
drwxr-xr-x  13 root   root    4096 Jun 21 21:40 var
```

You're now root, or are you?

- - -

## Build
```
make end-to-end
```
