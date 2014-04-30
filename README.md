# ohai there

Chef's excellent system profiling tool, [ohai](http://docs.opscode.com/ohai.html), as a Dockerized web service.

## Getting and using it

Just run it from Docker and it'll pull from the index. It exposes port 8000.

	$ docker run -d -P progrium/ohai-there

Then use it!

	$ curl $(docker port `docker ps -lq` 8000)/memory
	{
	  "swap": {
	    "cached": "0kB",
	    "total": "3063804kB",
	    "free": "3063804kB"
	  },
	  "total": "1018608kB",
	  "free": "207756kB",
	  "cached": "551164kB"
	}

You can drill in:

	$ curl $(docker port `docker ps -lq` 8000)/memory/swap/total
	[
	  "3063804kB"
	]

Or you can see all attributes by not specifying:

	$ curl $(docker port `docker ps -lq` 8000)
	...

## Usefulness from inside a container

Certainly many of the attributes are not useful since they're attributes of the container environment. For example, it would be really nice if the hostname and IP attributes reflected that of the host. If somebody can come up with a clever way to make that work, that'd be great. But for basic system information and resource (CPU, memory), it's still very useful. 

One todo is to turn off builtin ohai plugins that provide totally useless information, or information we'll never get to accurately reflect the host from inside a container. 

So why even run it in a container? Well, ohai is a great command-line system profiling tool that works standalone. More people should use it. This version turns it into a web service that can quickly and easily run on any host with Docker. I made it to work with Flynn Layer 0 so that Flynn tooling can better programmatically know about the hosts they run on. 

## Sponsor

This project (as simple as it is) was made possible thanks to [DigitalOcean](http://digitalocean.com).

## License

BSD