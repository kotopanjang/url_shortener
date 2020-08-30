#Questions

	You observed that 80% of the transactions are requests for 3 shortened identifier. 
	Please propose what will you change in the system that will optimise the performance

>What I will change on the system is,
	when the users request for an identifier, we will check on the database whether the data is already stored on the database or not. if yes, we will use that data insted of creating new record. It will minimize the database size.

#

	Assuming that the server is deployed in Singapore today, and a huge number of traffic comes from South Korea. 
	Korean users are complaining that the latency is too long. What would you do?

>1. Enable caching on the client browser
	Because it will be faster if we load from cache insted of accessing resource from server everytime. And yes it wil be slow at the first time open the app, but it will be faster after the browser get the cache.
>2. Clusterig server
	If its related to server, maaybe clustering will help
>3. Add some more servers in the biggest traffic region
	It will reduce the latency of that region.

#

	Imagine that you are developing a full website for this tinyURL application. 
	Draw a technical architecture diagram to show how the website, server, database, or any other 
	components interact with one another if you were to deploy it on AWS or GCP.
	
>Here is the diagram example how it works with cache
>https://drive.google.com/file/d/1tZmyc7pkgtLTrMtJm6r-X0rbeJxM7bd7/view?usp=sharing
>https://drive.google.com/file/d/1v4vfV96Rf26C1EPH_ApRLW6M04yMBrd9/view?usp=sharing
>https://drive.google.com/file/d/11-Kxob8TiWVTStP_ShSq_lCjojzcEvyh/view?usp=sharing
