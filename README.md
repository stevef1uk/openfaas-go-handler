# test2
OpenFaas test
This is the main dashboard project that has the OpenFaaS Cloud webhooks running against it.
It also contains a simple Go handler which can be deployed as an OpenFaaS function. 
The function handles GET & POST methiods and in this case route the request to test4, which is an automatically
generated function (from cassuservice) that reads or writes to a Cassandra table depending upon whether a GET or POST.
