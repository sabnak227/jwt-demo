npm i && # installs the pem-jwk node package
openssl genrsa -out app.rsa 2048 &&
openssl rsa -in app.rsa -outform PEM -pubout -out app.rsa.pub &&
cat app.rsa | npm run pem >private.jwk &&  # converts the RSA private key to JWK and directs the output to a file named private.jwk
echo 'done!'