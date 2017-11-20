#! /bin/sh                                                 

FILE=$1                                                    
gofmt -w $1                                                
go build -race -o ${FILE%%.*}_full "$1"                    
GOOS=linux go build -ldflags="-s -w" -o ${FILE%%.*} "$1"   
upx -f --brute ${FILE%%.*}                                 

ls -lh ${FILE%%.*}*  
