# test2
OpenFaas test
This is the main dashboard project that has the OpenFaaS Cloud webhooks running against it.
It also contains a simple Go handler which can be deployed as an OpenFaaS function. 
The function handles GET & POST methiods and in this case route the request to test4, which is an automatically
generated function (from cassuservice) that reads or writes to a Cassandra table depending upon whether a GET or POST.
The functional can be invoked for READ via the OpenFaaS GUI at https://system.sjfisher.com/dashboard/stevef1uk/go1?repoPath=stevef1uk/test2
or using curl for POST as follows:
<pre><code>
curl -d '{"id": 3, "message": "Sarah"}' -H "Content-Type: application/json" -v -X POST https://stevef1uk.sjfisher.com/go
</pre></code>
Alternatively the OpenFaaS Gateway GUI can be used: https://gw.sjfisher.com/ui/
To deploy the Gateway I followed the ofc_bootstrap approach to deploy OpenFaaS Private Cloud and then applied the following yaml:
<pre><code>
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
 name: openfaas-gw-ingress
 annotations:
  kubernetes.io/ingress.class: "nginx"
  cert-manager.io/issuer: letsencrypt-prod
spec:
 tls:
  - hosts:
    - gw.<<YOUR DOMAIN NAME>>
   secretName: openfaas-crt
 hosts:
  - host: gw.<<YOUR DOMAIN NAME>>
   serviceName: gateway
   servicePort: 8080
   path: /
 rules:
  - host: gw.<<YOUR DOMAIN NAME>>
   http:
    paths:
    - path: /
     backend:
      serviceName: gateway
      servicePort: 8080
</pre></code>
