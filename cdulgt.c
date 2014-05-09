#include <stdio.h>
#include <string.h>
#include <libscrypt.h>
//#include <b64.h>


int N = 16384;
int r = 8;
int p = 1;
const int keyLen = 8;

int main(void) {
	unsigned char digest[keyLen];
    const unsigned char password[] = "mypassword";
    const unsigned char salt[] = "mysalt";
    
	libscrypt_scrypt(password, sizeof(password), salt, sizeof(salt), N, r, p, digest, keyLen);

    for (int i = 0; i < keyLen; ++i) {
        printf("%x ", digest[i]);
    }
    printf("\n");

	return 0;
}