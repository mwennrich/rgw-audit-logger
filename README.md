# rgw-audit-logger

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: rook-config-override
  namespace: rook-ceph
data:
  config: |
    [global]
    rgw ops log rados = false
    rgw enable ops log = true
    rgw ops log socket path = /var/log/ceph/ops-log.sock
```
