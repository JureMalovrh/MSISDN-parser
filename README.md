# MSISDN-parser

### Requirements: 
- [Vagrant](https://www.vagrantup.com/)


MSISDN number parser. It retruns MNO identifier, country dialing code, subscriber number and country identifier.

Note:
MNO identifiers are based on https://en.wikipedia.org/wiki/List_of_mobile_phone_number_series_by_country, so if there is none MNO provider listed for the country, Unknown provider is returned. 

Note 2:
North American countries are parsed as country code for america (+1) and state number (e.g. 202), not only coutry code.

For the RPC part, JSON-RPC was chosen. It was chosen because with JSON-RPC we can support more than just Go RPC calls. 

### Instalation:
- clone repo
- ```cd``` into it
- ```vagrant up``` note: this may take some time as this box is custom and because of this does not have consistent download rate
- ```vagrant ssh``` 
- after being succesfully ssh-ed, go to ```/home/vagrant/go/src/msisdn``` and run ```go run client.go```. This will start RPC server and run 10 different calls.
