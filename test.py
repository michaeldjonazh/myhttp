import socket

def getMountPointString():
    # {3} = HTTP/1.1 or HTTP/1.0
    mountPointString = "GET {0} {3}\r\nUser-Agent: {1}\r\nAuthorization: Basic {2}\r\n".format("/", "Safari/1.0", "MnZ0XcAxOjzyezVxd2By", "HTTP/1.0")
    #mountPointString += "Content-Length: 38\r\n"
    mountPointString += "\r\n"
    mountPointString += "\r\n" + "- I want to get this value over here -\r\n"
    
    #mountPointString = "GET / HTTP/1.0\r\nUser-Agent: Safari/1.0\r\nAuthorization: Basic MnZ0XcAxOjzyezVxd2By\r\n\r\n\r\n- I want to get this value over here -\r\n"
    
    
    return mountPointString

mesocket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
mesocket.connect_ex(("localhost", 8080))
mesocket.settimeout(10)
print("=== REQ ===")
print(getMountPointString().encode(encoding='utf-8'))
mesocket.sendall(getMountPointString().encode(encoding='utf-8'))
res = mesocket.recv(4096)
print("=== RES ===")
print(res)
mesocket.close()