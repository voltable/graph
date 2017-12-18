import rpc from './vertex_pb';
import grpcFetch from './grpcfetch';


export default function () {
  let request = new rpc.QueryRequest({});
  request.setText("hi");

  grpcFetch('/rpc.Graph/Query', request.serializeBinary()).then((bytes) => {
    var message2 = rpc.QueryReply.deserializeBinary(bytes);
    console.log('message2', message2.getText());
    console.log('message2', message2.getPropertiesMap().get('test').getId());
  })
};