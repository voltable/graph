import client from './vertex_pb';

export default function () { 

    var myHeaders = new Headers();
    myHeaders.append('Content-Type', 'application/grpc');
    var request = new client.QueryRequest({});
    request.setText("hi");  

    fetch('https://localhost:8080/client.Graph/Query', {method: 'POST', headers: myHeaders, body: request.serializeBinary()}).then(function(response) {
        console.log('hit', response);
      return response.blob();
    });
 };