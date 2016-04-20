//There are 3 levels health check:
//1.ping the server's address, check the response
//2.check the app's port, send a message to app's port. If the response is avalible, the web application's could be connected.
//3.send a http request to the specific page(API), if the http response's status eq 200, it means the wep application is
//in a good condition.
package endpoint


