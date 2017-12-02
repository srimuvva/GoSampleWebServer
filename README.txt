###########LIMITATIONS###########
It seems http://api.icndb.com/jokes does not accept the unicode string as firstname or lastname,
but http://uinames.com/api/ can return a UTF8 unicode string eg: Εύρυτος which is may be greece or chinese?
So in this you get an emtpy response as the joke server not supporting UTF8. If you see an empty response when you
execute curl command, that means http://uinames.com/api/ has given us utf8 string and we did get bad reponse from 
http://api.icndb.com/jokes.
############################ 


##########INSTRUCTIONS to RUN the gowebserver within the HOST########
#######MAKESURE you have the connectivity repos..##########

# Go to you home dir
cd /home/smuvva/
# Checkout git
git clone https://github.com/srimuvva/GoSampleWebServer.git
cd GoSampleWebServer
##Install golang package, depending on your host, in this case Centos, so YUM, in ubuntu apt-get install golang-go$
[root@mapi09 GoSampleWebServer]# yum install go-lang
[root@mapi09 GoSampleWebServer]# go build webserv.go

# Run webserv using below cmd, we dont specify port, server uses 8080 port.
[root@mapi09 GoSampleWebServer]# ./webserv -port 5678
Now from any machine issue command curl<IP>:<Port> should return Joke. If you got empty respose that could be because of above said limitation that server does not support UTF8 unicode firstname and last names. Try more to get proper response.



##########INSTRUCTIONS to RUN the gowebserver as a container##########

#We have Dockerfile as well in GoSampleWebServer dir.

##Build docker image using command, conatiner uses the ubuntu as base image####
[root@mapi09 GoSampleWebServer]# sudo docker build -t gowebserver .

# Run the server mapping to any external port of your xhoice
[root@mapi09 GoSampleWebServer]# docker run -p <externalport>:8080 gowebserver:latest
Eg: docker run -p 6789:8080 gowebserver:latest

#test the server using from any extrenal host or within the host.
curl <IP>:6789

AS MENTIONED in LIMITATIONS, sometimes you see empty response because UTF8 firstname or lsatname no supported by Joke server, during the handling of this invalid response go http library issues a panic error log that can be seen on console, but continues to serve all the next requests.



##########ALTERNATE DESIGN  TODO###############
We can run the above web requests in parallel in separate go routines, as the only dependency is getting the firstname and lastname.
We can issue the http get http://api.icndb.com/jokes with dummy firstname and lastname, at the same time we issue http://uinames.com/api/
to get the firstname and lastname. Then we format the reponse from Joke server by replacing the dummy first and last names with the response from name server.
This way we can support the UTF8 string Jokes as well and also better performance as both the GET requests are run in parallel.
################################# 
