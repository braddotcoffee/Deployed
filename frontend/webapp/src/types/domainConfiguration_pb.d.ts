// package: datastore
// file: domainConfiguration.proto

import * as jspb from "google-protobuf";

export class DomainConfiguration extends jspb.Message {
  getDomain(): string;
  setDomain(value: string): void;

  getPort(): string;
  setPort(value: string): void;

  getForwardDirectory(): string;
  setForwardDirectory(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DomainConfiguration.AsObject;
  static toObject(includeInstance: boolean, msg: DomainConfiguration): DomainConfiguration.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DomainConfiguration, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DomainConfiguration;
  static deserializeBinaryFromReader(message: DomainConfiguration, reader: jspb.BinaryReader): DomainConfiguration;
}

export namespace DomainConfiguration {
  export type AsObject = {
    domain: string,
    port: string,
    forwardDirectory: string,
  }
}

