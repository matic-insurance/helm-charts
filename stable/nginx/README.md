# nginx Helm chart

The Helm chart perform Nginx deployment to the Kubernetes cluster. With default parameters simple Nginx web-server will be installed and available via ClusterIP within Kubernetes CNI network.

## Requirements

* In order to overwrite default nginx config, custom configmap should be created in advance and passed to the chart as `configmap_name` parameter. Example of custom ConfigMap:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: configmap-nginx
data:
  default.conf:
  ...
```

* When SSL support is enabled, a chart will be looking for TLS secrets. Secret name is expected in following format: `<domain_name>-tls`. For example:

``` yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: example.com-tls
type: tls
data:
  ca.crt: 
  ...
  cat.key:
  ...
```

The TLS data will be mounted to container's  `/etc/nginx/certs/<domain_name>` directory.


## Usage

In order to install chart: `helm install -n nginx ./nginx --namespace mynamespace`
