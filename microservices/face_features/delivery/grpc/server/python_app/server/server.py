import server.protobuf.face_features_pb2 as face
import server.protobuf.face_features_pb2_grpc as grpc_face
import server.api.api as api


class Server(grpc_face.FaceGRPCHandlerServicer):
    def HaveFace(self, request, context):
        faces = api.find_faces(request.path)

        return face.Face(have=faces != [])

    def AddMask(self, request, context):
        photo, mask = api.overlay_mask(request.path, request.mask)

        return face.Photo(path=photo, mask=mask)
