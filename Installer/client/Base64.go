package main

import "os"
import "fmt"
import "encoding/base64"
import "net/url"
import "crypto/rand"

func PEMDecoder( PemFileBase64Str string ) []uint8 {
    PEMFileObj ,_ := base64.RawStdEncoding.DecodeString( PemFileBase64Str )
    return PEMFileObj
}

func PEMEncoder( PemFileName string ) string {
    PEMFileObj , _ := os.ReadFile( PemFileName )
    PEMFileObjBase64Str := base64.RawStdEncoding.EncodeToString( PEMFileObj )
    return PEMFileObjBase64Str
}

func URLCredEncoder( URLCred string ) string {
    return base64.RawURLEncoding.EncodeToString( []byte( URLCred ) )
}

// DoHResponse
func URLCreator( FQDN string , UserName string , Password string , QueryMap map[string]string ) {
    URLValues := url.Values{}
    for QueryMapKey , QueryMapValue := range QueryMap {
        URLValues.Add( QueryMapKey , QueryMapValue )
    }
    URLValuesStr = URLValues.Encode()
    URLCreds := url.UserPassword( URLCredEncoder( UserName ) , URLCredEncoder( Password ) )
    QueryURL := &url.URL {
        "Scheme" : "https" ,
        "User" :  URLCreds ,
        "Host" : FQDN ,
        "RawQuery" : URLValuesStr
    }
    return QueryURL.String()
}

func SessionID( ) string {
    RandomBytes := make( []byte , 32 )
    _ , _ = rand.Read( RandomBytes )
    ID := base64.RawURLEncoding.EncodeToString( RandomBytes )
    return ID
}

func main() {
    CertObjBase64 := PEMEncoder( "Client-TLS-ED25519.crt" )
    KeyObjBase64 := PEMEncoder( "Client-TLS-ED25519.key" )
    CertObj := PEMDecoder( CertObjBase64 )
    KeyObj := PEMDecoder( KeyObjBase64 )
    SecondRequestQueryMap := map[ string ]string { "Profile" : "US" , "ID" : SessionID() }
    SecondRequestURL := URLCreator( "domain.xyz" , "Sajjad" , "drama" , SecondRequestQueryMap )
}
