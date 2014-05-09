# Important Note
You can't compile the go with the cdulgt.c in the directory using `go build`. Either remove it before compiling it or specify the file to compile directly with:

    go build main.go

The C file depends on libscrypt found at: https://github.com/technion/libscrypt . It is included for testing purposes to see if the go code produces the same result. It doesn't but they might use slightly different approaches.

# Dulgt

Dulgt is a Norwegian word for hidden or veiled. It is a way of hiding your secrets such as your login passwords. So this is just another password manager software like 1Password, KeePass or LastPass? Not quite. When using usual password managers you need to store the passwords on a file. With Dulgt passwords are generated from the information you write into it, so you can recreate your passwords on any computer.