// source: api/user/service/v1/error_reason.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var errors_errors_pb = require('../../../../errors/errors_pb.js');
goog.object.extend(proto, errors_errors_pb);
goog.exportSymbol('proto.raiden.v1.ErrorReason', null, global);
/**
 * @enum {number}
 */
proto.raiden.v1.ErrorReason = {
  DATABASE: 0,
  VALIDATOR: 1
};

goog.object.extend(exports, proto.raiden.v1);
