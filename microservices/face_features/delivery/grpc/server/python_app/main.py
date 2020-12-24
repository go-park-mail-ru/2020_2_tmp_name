from concurrent import futures
import server.protobuf.face_features_pb2_grpc as grpc_face
from server.server import Server
import grpc
import daemon


def serve():
    print("Starting server at: 8083")
    s = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    grpc_face.add_FaceGRPCHandlerServicer_to_server(Server(), s)
    s.add_insecure_port('face:8083')
    s.start()
    s.wait_for_termination()


# with daemon.DaemonContext():
serve()
