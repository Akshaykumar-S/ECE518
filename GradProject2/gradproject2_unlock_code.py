from hashlib import sha256 
cw_hex = 0x0138e177
start = 0x000009184E72A000
b = start+0x0000000000000001 
val = hex((cw_hex<<64)|b) 
d = val.replace('0x','0') 
#print(d) 
hash = sha256(bytes.fromhex(d)) 
hx = hash.hexdigest() 
hx = str(hx) 
#print(hx) 
#print(hx[0:8]) 
while b <= 0xffffffffffffffff:
    print(b)
    if (hx[0:8] == '0138e177'):
        print("Success:",b)
        exit()
    else:
        b = b + 0x0000000000000001
        d = hex((cw_hex << 64) | b)
        d = d.replace('0x', '0')
        hash = sha256(bytes.fromhex(d))
        hx = hash.hexdigest()
        hx = str(hx)