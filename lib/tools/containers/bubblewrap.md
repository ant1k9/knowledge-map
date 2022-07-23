[bubblewrap](https://github.com/containers/bubblewrap)

*Description*: Unprivileged sandboxing tool

*Labels*: #C #Containers

*Links*:
  - https://jvns.ca/blog/2022/06/28/some-notes-on-bubblewrap/

*Examples*:

```bash
$ bwrap \
    --ro-bind / / \
    --bind /tmp /tmp \
    --proc /proc --dev /dev \
    --unshare-pid \
    --unshare-net \
    bash
```
