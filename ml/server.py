from concurrent import futures
import grpc
import languagedetect_pb2
import languagedetect_pb2_grpc
from ml.languageDetection import DetectLanguageModel

class LanguageDetectServicer(languagedetect_pb2_grpc.LanguageDetectServicer):
    def GetLanguage(self, request, context):
        model = DetectLanguageModel()
        model.build_model()
        ret_val = model.detect_language(request.text)
        return languagedetect_pb2.LanReply(response=ret_val)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    languagedetect_pb2_grpc.add_LanguageDetectServicer_to_server(LanguageDetectServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()