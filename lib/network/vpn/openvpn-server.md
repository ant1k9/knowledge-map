[OpenVPN Server Configuration](https://www.pei.com/openvpn-server-configuration/)

*Description*: This document describes the process of building an OpenVPN server to facilitate secure remote access to systems.

*Labels*: #VPN #OpenVPN

*Docs*: https://dnaeon.github.io/firewall-rules-with-iptables-for-openvpn/

*Examples*:

```bash
$ echo 'mode server
local 10.128.0.16
port 1194
proto udp
dev tun
topology subnet
server 172.50.0.0 255.255.255.0
push "redirect-gateway def1"
push "dhcp-option DNS 8.8.8.8"
max-clients 100
keepalive 10 120
persist-key
persist-tun

user nobody
group nogroup
ca /etc/openvpn/server/ca.crt
cert /etc/openvpn/server/server.crt
key /etc/openvpn/server/server.key # This file should be kept secret
dh /etc/openvpn/server/dh.pem
tls-auth /etc/openvpn/server/ta.key 0
cipher AES-256-CBC
auth SHA512
tls-version-min 1.2
tls-cipher TLS-DHE-RSA-WITH-AES-256-GCM-SHA384:TLS-DHE-RSA-WITH-AES-128-GCM-SHA256:TLS-DHE-RSA-WITH-AES-256-CBC-SHA:TLS-DHE-RSA-WITH-CAMELLIA-256-CBC-SHA:TLS-DHE-RSA-WITH-AES-128-CBC-SHA:TLS-DHE-RSA-WITH-CAMELLIA-128-CBC-SHA

status openvpn-status.log
verb 4
mute 20' > /etc/openvpn/server/server.conf
$ nohup jopenvpn server/server.conf &
```
