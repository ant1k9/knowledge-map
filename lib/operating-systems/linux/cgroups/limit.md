[Limit specific process memory on desktop linux with cgroups and earlyoom](https://raymii.org/s/articles/Limit_specific_process_memory_on_desktop_linux_with_cgroups.html)

*Description*: Article about custom configuration of cgroups

*Labels*: #Linux #CGroups

*Examples*:

```bash
$ apt install cgroup-tools
$ sudo cgcreate -t remy:remy -a remy:remy -g memory:/cgTeams
$ echo $(( 2048 * 1024 * 1024  )) | sudo tee /sys/fs/cgroup/memory/cgTeams/memory.limit_in_bytes #2 GB RAM
$ echo $(( 2049 * 1024 * 1024  )) | sudo tee /sys/fs/cgroup/memory/cgTeams/memory.memsw.limit_in_bytes #2GB swap, only works if you have swap
$ cgexec -g memory:cgTeams teams 
```
