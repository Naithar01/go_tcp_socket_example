# go_tcp_socket_example

## Server 
1. n 개의 클라이언트에게 accept를 받음
2. Read로 recv 받고 받은 recv를 다시 Write (send) 해줌 

## Client 
1. 열려있는 서버에 포트 입력하여 입장 가능
2. Write로 문자를 보내고 성공하면 다시 Read 받음 
