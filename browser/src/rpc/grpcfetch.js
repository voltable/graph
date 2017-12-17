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

export default function fetchGRPC(url, bytes) {
    let myHeaders = new Headers();
    myHeaders.append('Content-Type', 'application/grpc');

    return fetch(url, { method: 'POST', headers: myHeaders, body: grpcEncode(bytes) })
        .then((response) => {
            return response.arrayBuffer();
        }).then((bytes) => {
            return grpcDecode(bytes);
        });
}