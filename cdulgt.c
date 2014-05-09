#include <stdio.h>
#include <string.h>
#include <libscrypt.h>
//#include <b64.h>

struct {
    uint8_t P[255];
    uint8_t S[255];
    int N, r, p, dklen;
} ts[] = {
    {"", "", 16, 1, 1, 64},
    {"password", "NaCl", 1024, 8, 16, 64},
    {"pleaseletmein", "SodiumChloride", 16384, 8, 1, 64},
    {"pleaseletmein", "SodiumChloride", 1048576, 8, 1, 64},
    {"mypassword", "mysalt", 16384, 8, 1, 8},    
};


int main(void) {
    uint8_t digest[1024];
    for(size_t i = 0; i < 5; ++i)
    {
    	libscrypt_scrypt(ts[i].P, strlen((const char *)ts[i].P), 
                         ts[i].S, strlen((const char *)ts[i].S), 
                         ts[i].N, ts[i].r, ts[i].p, digest, ts[i].dklen);

        for (int j = 0; j < ts[i].dklen; ++j) {
            printf("%x ", digest[j]);
        }
        printf("\n\n");
    }


	return 0;
}