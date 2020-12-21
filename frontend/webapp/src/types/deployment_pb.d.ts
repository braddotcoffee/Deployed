// package: datastore
// file: deployment.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Deployment extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getRepository(): string;
  setRepository(value: string): void;

  getDockerfile(): string;
  setDockerfile(value: string): void;

  getCommit(): string;
  setCommit(value: string): void;

  hasLastDeploy(): boolean;
  clearLastDeploy(): void;
  getLastDeploy(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastDeploy(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getStatus(): Deployment.DeployStatusMap[keyof Deployment.DeployStatusMap];
  setStatus(value: Deployment.DeployStatusMap[keyof Deployment.DeployStatusMap]): void;

  getDomain(): string;
  setDomain(value: string): void;

  getBuildCommand(): string;
  setBuildCommand(value: string): void;

  getOutputDirectory(): string;
  setOutputDirectory(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Deployment.AsObject;
  static toObject(includeInstance: boolean, msg: Deployment): Deployment.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Deployment, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Deployment;
  static deserializeBinaryFromReader(message: Deployment, reader: jspb.BinaryReader): Deployment;
}

export namespace Deployment {
  export type AsObject = {
    name: string,
    repository: string,
    dockerfile: string,
    commit: string,
    lastDeploy?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    status: Deployment.DeployStatusMap[keyof Deployment.DeployStatusMap],
    domain: string,
    buildCommand: string,
    outputDirectory: string,
  }

  export interface DeployStatusMap {
    NOT_STARTED: 0;
    IN_PROGRESS: 1;
    COMPLETE: 2;
    ERROR: 3;
  }

  export const DeployStatus: DeployStatusMap;
}

