# TCP

TCP is Transmission Control Protocol, a Layer 4 OSI Model network protocol, that
requires a formal connection being established first before it will allow data
to be transferred to and from the client and server connections.

Most of the data sent over IP (Internet Protocol, Layer 3 OSI) is over TCP.
That is because TCP sets up a reliable connection. 
Common use cases is for web browsing, email, text messages, file transfers, etc.
The extra steps needed to ensure reliable connection for (guaranteed order 
and sending of data packets over TCP) TCP makes it slower in comparison to UDP.
Therefore, if you need more speed and are willing to have some potential packet loss,
which might mean you will need to send more packets to get all the things you wanted
to send over to be received, go with UDP for live, real-time data transmission.


---

Here we will be using the `Go` programming language to try to understand how data
is streamed from a client to a server and back. 
