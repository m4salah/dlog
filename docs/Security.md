dlog is designed to be accessed by trusted clients inside trusted environments. This means that usually it is not a good idea to expose the dlog instance directly to the internet or, in general, to an environment where untrusted clients can directly access the dlog TCP port.

If you want to expose it over unsecure HTTP (for development purposes or in LAN), please use `--serve-insecure true` flag.

# Listening on specific network interface

dlog accepts `--bind` flag that defines the interface which dlog should listen to. `--bind` is in the format `<ip.address.of.interface>:<port>`. 

- To listen on all interfaces on port 3000 pass `--bind 0.0.0.0:3000`
- To listen on specific interface pass the interface IP address `--bind 192.168.8.200:3000`

# Readonly mode

dlog accept a `--readonly` flag to signal all features not to write to the disk. Readonly mode is not a safe measure for exposing the server to the internet. additionally make sure you sandbox the process in a restricted environment such as docker, CGROUPS or another user that has readonly access to the disk. 

Extensions can ignore the readonly flag so make sure you use trusted extensions only in case you intend to expose dlog to the internet.


# Reporting Security Issues

Please report any issues to me on Keybase: https://keybase.io/emadelsaid
