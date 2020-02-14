## How to use this repo

1. Place the chart into `stable` folder
2. Add repo to your Helm repositories as:
```
helm repo add matic https://matic-insurance.github.io/helm-charts
```
3. Work with the charts
```
helm search matic/nginx
NAME        	CHART VERSION	APP VERSION	DESCRIPTION
matic/nginx	1.0.0        	1.0        	A Nginx Helm chart for Kubernetes
```
