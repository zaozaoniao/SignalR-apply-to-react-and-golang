# SignalR-apply-to-react-and-golang
demo for  SignalR intergate  with  react  and  golang

> SingnalR is launched  by microsoft, is used for  realtime  communication, is also a  common  framework,  There have go  & node port version signalr.  
SignalR use  rpc to allow  server call web client javascript code in .NET code.  
SignalR privder a simple api to construct connection,connection managemnt, message push.



We  give a react-golang demo for the signalr realtime communication.  

#### Step:
 
 - react:  

 pay close attention to the  `clock.js`, is a react componment for the signalr client


-  golang

pay attention to the `signalrHub`, and init a hubserver with the specified  hub.

Then  server can push  message  at the right time

