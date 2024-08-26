# UDP 

UDP is User Datagram Protocol, Layer 4 OSI model, that uses a session-less
method of sending and receiving data inside datagrams structure.

People use UDP because it is a simple structure because it doesn't need to make
a formal connection-like session which a protocol like TCP will do. This allows
UDP to send data fast, data go brrr and all that. However, because of the lack of
formal connection, there are no guarantees that the order in which the data is sent
is the order in which the data will be received, or if all datagrams will be received
at all. It's a protocol where you just spray and pray. If you need some guarantees
on order or robustness of data, you either wrote your own logic or use TCP instead. 

Common applications include video data streaming, online video game connections,
and DNS lookups.

---

Here we will be using the `Go` programming language to try to understand how data
is streamed from a client to a server and back. 
