npm i && # installs the pem-jwk node package
openssl genrsa -out app.rsa 2048 &&
openssl rsa -in app.rsa -outform PEM -pubout -out app.rsa.pub &&
cat app.rsa | npm run pem > private.jwk &&
cat private.jwk | sed -e 's/>.*//g' | grep . > jwks.json &&
rm private.jwk &&
echo 'done!'