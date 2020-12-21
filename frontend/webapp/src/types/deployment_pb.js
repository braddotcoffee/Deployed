/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.exportSymbol('proto.datastore.Deployment', null, global);
goog.exportSymbol('proto.datastore.Deployment.DeployStatus', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.datastore.Deployment = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.datastore.Deployment, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.datastore.Deployment.displayName = 'proto.datastore.Deployment';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.datastore.Deployment.prototype.toObject = function(opt_includeInstance) {
  return proto.datastore.Deployment.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.datastore.Deployment} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.datastore.Deployment.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    repository: jspb.Message.getFieldWithDefault(msg, 2, ""),
    dockerfile: jspb.Message.getFieldWithDefault(msg, 3, ""),
    commit: jspb.Message.getFieldWithDefault(msg, 4, ""),
    lastDeploy: (f = msg.getLastDeploy()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    status: jspb.Message.getFieldWithDefault(msg, 6, 0),
    domain: jspb.Message.getFieldWithDefault(msg, 7, ""),
    buildCommand: jspb.Message.getFieldWithDefault(msg, 8, ""),
    outputDirectory: jspb.Message.getFieldWithDefault(msg, 9, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.datastore.Deployment}
 */
proto.datastore.Deployment.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.datastore.Deployment;
  return proto.datastore.Deployment.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.datastore.Deployment} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.datastore.Deployment}
 */
proto.datastore.Deployment.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRepository(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDockerfile(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCommit(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastDeploy(value);
      break;
    case 6:
      var value = /** @type {!proto.datastore.Deployment.DeployStatus} */ (reader.readEnum());
      msg.setStatus(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setDomain(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setBuildCommand(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setOutputDirectory(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.datastore.Deployment.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.datastore.Deployment.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.datastore.Deployment} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.datastore.Deployment.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRepository();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDockerfile();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCommit();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getLastDeploy();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      6,
      f
    );
  }
  f = message.getDomain();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getBuildCommand();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getOutputDirectory();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.datastore.Deployment.DeployStatus = {
  NOT_STARTED: 0,
  IN_PROGRESS: 1,
  COMPLETE: 2,
  ERROR: 3
};

/**
 * optional string name = 1;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string repository = 2;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getRepository = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setRepository = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string dockerfile = 3;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getDockerfile = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setDockerfile = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string commit = 4;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getCommit = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setCommit = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional google.protobuf.Timestamp last_deploy = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.datastore.Deployment.prototype.getLastDeploy = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/** @param {?proto.google.protobuf.Timestamp|undefined} value */
proto.datastore.Deployment.prototype.setLastDeploy = function(value) {
  jspb.Message.setWrapperField(this, 5, value);
};


proto.datastore.Deployment.prototype.clearLastDeploy = function() {
  this.setLastDeploy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.datastore.Deployment.prototype.hasLastDeploy = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional DeployStatus status = 6;
 * @return {!proto.datastore.Deployment.DeployStatus}
 */
proto.datastore.Deployment.prototype.getStatus = function() {
  return /** @type {!proto.datastore.Deployment.DeployStatus} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {!proto.datastore.Deployment.DeployStatus} value */
proto.datastore.Deployment.prototype.setStatus = function(value) {
  jspb.Message.setProto3EnumField(this, 6, value);
};


/**
 * optional string domain = 7;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getDomain = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setDomain = function(value) {
  jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string build_command = 8;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getBuildCommand = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setBuildCommand = function(value) {
  jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional string output_directory = 9;
 * @return {string}
 */
proto.datastore.Deployment.prototype.getOutputDirectory = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/** @param {string} value */
proto.datastore.Deployment.prototype.setOutputDirectory = function(value) {
  jspb.Message.setProto3StringField(this, 9, value);
};


goog.object.extend(exports, proto.datastore);
