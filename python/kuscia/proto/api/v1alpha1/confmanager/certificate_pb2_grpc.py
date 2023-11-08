# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from kuscia.proto.api.v1alpha1.confmanager import certificate_pb2 as kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2


class CertificateServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GenerateKeyCerts = channel.unary_unary(
                '/kuscia.proto.api.v1alpha1.confmanager.CertificateService/GenerateKeyCerts',
                request_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsRequest.SerializeToString,
                response_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsResponse.FromString,
                )


class CertificateServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GenerateKeyCerts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CertificateServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GenerateKeyCerts': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateKeyCerts,
                    request_deserializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsRequest.FromString,
                    response_serializer=kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kuscia.proto.api.v1alpha1.confmanager.CertificateService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CertificateService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GenerateKeyCerts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kuscia.proto.api.v1alpha1.confmanager.CertificateService/GenerateKeyCerts',
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsRequest.SerializeToString,
            kuscia_dot_proto_dot_api_dot_v1alpha1_dot_confmanager_dot_certificate__pb2.GenerateKeyCertsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
