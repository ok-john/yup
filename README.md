
# yup

containers & privilage stuff

use this library from an entirely unprivileged userspace

## Enter
```
make enter
```

## Run
```
make run
```

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
- - -
```
so you're root now, right?

## Build
```
make end-to-end
```
