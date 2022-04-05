/**
 * @fileoverview gRPC-Web generated client stub for user.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var validate_validate_pb = require('../../../../validate/validate_pb.js')

var google_api_annotations_pb = require('../../../../google/api/annotations_pb.js')
const proto = {};
proto.user = {};
proto.user.v1 = require('./user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.user.v1.UserSystemClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.user.v1.UserSystemPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.v1.RegisterRequest,
 *   !proto.user.v1.RegisterReply>}
 */
const methodDescriptor_UserSystem_Register = new grpc.web.MethodDescriptor(
  '/user.v1.UserSystem/Register',
  grpc.web.MethodType.UNARY,
  proto.user.v1.RegisterRequest,
  proto.user.v1.RegisterReply,
  /**
   * @param {!proto.user.v1.RegisterRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.v1.RegisterReply.deserializeBinary
);


/**
 * @param {!proto.user.v1.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.v1.RegisterReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.v1.RegisterReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.v1.UserSystemClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.v1.UserSystem/Register',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Register,
      callback);
};


/**
 * @param {!proto.user.v1.RegisterRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.v1.RegisterReply>}
 *     Promise that resolves to the response
 */
proto.user.v1.UserSystemPromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.v1.UserSystem/Register',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.v1.FollowersRequest,
 *   !proto.user.v1.FollowersReply>}
 */
const methodDescriptor_UserSystem_Followers = new grpc.web.MethodDescriptor(
  '/user.v1.UserSystem/Followers',
  grpc.web.MethodType.UNARY,
  proto.user.v1.FollowersRequest,
  proto.user.v1.FollowersReply,
  /**
   * @param {!proto.user.v1.FollowersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.v1.FollowersReply.deserializeBinary
);


/**
 * @param {!proto.user.v1.FollowersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.v1.FollowersReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.v1.FollowersReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.v1.UserSystemClient.prototype.followers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.v1.UserSystem/Followers',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Followers,
      callback);
};


/**
 * @param {!proto.user.v1.FollowersRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.v1.FollowersReply>}
 *     Promise that resolves to the response
 */
proto.user.v1.UserSystemPromiseClient.prototype.followers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.v1.UserSystem/Followers',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Followers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.v1.FollowingRequest,
 *   !proto.user.v1.FollowingReply>}
 */
const methodDescriptor_UserSystem_Following = new grpc.web.MethodDescriptor(
  '/user.v1.UserSystem/Following',
  grpc.web.MethodType.UNARY,
  proto.user.v1.FollowingRequest,
  proto.user.v1.FollowingReply,
  /**
   * @param {!proto.user.v1.FollowingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.v1.FollowingReply.deserializeBinary
);


/**
 * @param {!proto.user.v1.FollowingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.v1.FollowingReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.v1.FollowingReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.v1.UserSystemClient.prototype.following =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.v1.UserSystem/Following',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Following,
      callback);
};


/**
 * @param {!proto.user.v1.FollowingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.v1.FollowingReply>}
 *     Promise that resolves to the response
 */
proto.user.v1.UserSystemPromiseClient.prototype.following =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.v1.UserSystem/Following',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Following);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.user.v1.RelationRequest,
 *   !proto.user.v1.RelationReply>}
 */
const methodDescriptor_UserSystem_Relation = new grpc.web.MethodDescriptor(
  '/user.v1.UserSystem/Relation',
  grpc.web.MethodType.UNARY,
  proto.user.v1.RelationRequest,
  proto.user.v1.RelationReply,
  /**
   * @param {!proto.user.v1.RelationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.user.v1.RelationReply.deserializeBinary
);


/**
 * @param {!proto.user.v1.RelationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.user.v1.RelationReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.user.v1.RelationReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.user.v1.UserSystemClient.prototype.relation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/user.v1.UserSystem/Relation',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Relation,
      callback);
};


/**
 * @param {!proto.user.v1.RelationRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.user.v1.RelationReply>}
 *     Promise that resolves to the response
 */
proto.user.v1.UserSystemPromiseClient.prototype.relation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/user.v1.UserSystem/Relation',
      request,
      metadata || {},
      methodDescriptor_UserSystem_Relation);
};


module.exports = proto.user.v1;

