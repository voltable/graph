import rpc from './vertex_pb';
import grpcFetch from './grpcfetch';


const payloadLen = 1;
const sizeLen = 4;
const compressionNone = 0;

function pad(n, width) {
    n = n + '';
    return n.length >= width ? n : new Array(width - n.length + 1).join('0') + n;
}

// grpcEncode adds the grpc header to the protocol buffer object
function grpcEncode(bytes) {
    let size = pad(bytes.length, 4);
    var bufHeader = new Int8Array(payloadLen + sizeLen + bytes.length);
    // Not supporting comprssion
    bufHeader[0] = compressionNone;
    bufHeader[1] = size[0];
    bufHeader[2] = size[1];
    bufHeader[3] = size[2];
    bufHeader[4] = size[3];

    for (let index = 5; index < bytes.length + 5; index++) {
        bufHeader[index] = bytes[index - 5];
    }

    return bufHeader
}

// grpcDecode returns the bytes for protocol buffer to deserialize
function grpcDecode(bytes) {
    // Not supporting comprssion so can ignore the header
    let length = bytes.byteLength - (payloadLen + sizeLen);
    return new Int8Array(bytes, payloadLen + sizeLen, length);
}

export default function () {
  let myHeaders = new Headers();
  myHeaders.append('Content-Type', 'application/grpc');
  let request = new rpc.QueryRequest({});
  request.setText("hi");

  // grpcFetch('/rpc.Graph/Query', request.serializeBinary()).then((bytes) => {
  //   console.log(bytes);
  //   var message2 = rpc.QueryReply.deserializeBinary(new ArrayBuffer(bytes));
  //   console.log('message2', message2);
  //   console.log('message2', message2.getPropertiesMap());
  //   console.log('message2', message2.getPropertiesMap().get('test'));
  //   console.log('message2', message2.getPropertiesMap().get('test').getId());
  // })

  fetch('/rpc.Graph/Query', {method: 'POST', headers: myHeaders, body: grpcEncode(request.serializeBinary())})
    .then((response) => {             
    return response.arrayBuffer() ;
  }).then((bytes) => {
    console.log('message2', bytes);
    let body = grpcDecode(bytes);
    console.log('message2', body);
    var message2 = rpc.QueryReply.deserializeBinary(new ArrayBuffer(body));
    console.log('message2', message2);
    console.log('message2', message2.getText());
    // console.log('message2', message2.getPropertiesMap().get('test'));
    // console.log('message2', message2.getPropertiesMap().get('test').getId());
  });
};