A basic Blockchain written in Go.  This isnt automated or anything, more of a proof of concept to build on than anything.


Sample output:

boltik@Fezzan:~/code/basic-blockchain$ go run blockchain.go 
Block mined:  0000e617e12908367f288df9bacbf4e907519e025741618100e8571888eadb21
Block mined:  0000a881560f0ddcd92b2d82de2e2f26d56cfd88cc4c13e759bae1efb6b8d45a
Block mined:  0000c6ffab74a314549313e3f44dee2c9332580a5ca7eb8051637a3b381b2fd0
Index: 0
Timestamp: 2025-03-02 19:50:17.371765565 -0500 EST m=+0.000037585
Data: Genesis Block
Hash: 47ef4ce0ca5cd2fa2971c1ed9f1bd8591f75c3dee10006d924b40ec8233b64ad
PrevHash: 
--------------------------------
Index: 1
Timestamp: 2025-03-02 19:50:17.371827108 -0500 EST m=+0.000099108
Data: First Block
Hash: 0000e617e12908367f288df9bacbf4e907519e025741618100e8571888eadb21
PrevHash: 47ef4ce0ca5cd2fa2971c1ed9f1bd8591f75c3dee10006d924b40ec8233b64ad
--------------------------------
Index: 2
Timestamp: 2025-03-02 19:50:17.373022906 -0500 EST m=+0.001294916
Data: Second Block
Hash: 0000a881560f0ddcd92b2d82de2e2f26d56cfd88cc4c13e759bae1efb6b8d45a
PrevHash: 0000e617e12908367f288df9bacbf4e907519e025741618100e8571888eadb21
--------------------------------
Index: 3
Timestamp: 2025-03-02 19:50:17.388103716 -0500 EST m=+0.016375726
Data: Third Block
Hash: 0000c6ffab74a314549313e3f44dee2c9332580a5ca7eb8051637a3b381b2fd0
PrevHash: 0000a881560f0ddcd92b2d82de2e2f26d56cfd88cc4c13e759bae1efb6b8d45a
--------------------------------