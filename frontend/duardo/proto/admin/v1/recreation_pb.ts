// @generated by protoc-gen-es v1.7.2 with parameter "target=ts,import_extension=none"
// @generated from file admin/v1/recreation.proto (package admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum admin.v1.Status
 */
export enum Status {
  /**
   * @generated from enum value: STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: STATUS_IN_PROGRESS = 1;
   */
  IN_PROGRESS = 1,

  /**
   * @generated from enum value: STATUS_COMPLETED = 2;
   */
  COMPLETED = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Status)
proto3.util.setEnumType(Status, "admin.v1.Status", [
  { no: 0, name: "STATUS_UNSPECIFIED" },
  { no: 1, name: "STATUS_IN_PROGRESS" },
  { no: 2, name: "STATUS_COMPLETED" },
]);

/**
 * @generated from message admin.v1.Recreation
 */
export class Recreation extends Message<Recreation> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string team_id = 2;
   */
  teamId = "";

  /**
   * @generated from field: string problem_id = 3;
   */
  problemId = "";

  /**
   * @generated from field: admin.v1.Status status = 4;
   */
  status = Status.UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<Recreation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "admin.v1.Recreation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "team_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "status", kind: "enum", T: proto3.getEnumType(Status) },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Recreation {
    return new Recreation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Recreation {
    return new Recreation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Recreation {
    return new Recreation().fromJsonString(jsonString, options);
  }

  static equals(a: Recreation | PlainMessage<Recreation> | undefined, b: Recreation | PlainMessage<Recreation> | undefined): boolean {
    return proto3.util.equals(Recreation, a, b);
  }
}

/**
 * @generated from message admin.v1.GetRecreationsRequest
 */
export class GetRecreationsRequest extends Message<GetRecreationsRequest> {
  /**
   * @generated from field: string team_id = 1;
   */
  teamId = "";

  constructor(data?: PartialMessage<GetRecreationsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "admin.v1.GetRecreationsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "team_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecreationsRequest {
    return new GetRecreationsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecreationsRequest {
    return new GetRecreationsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecreationsRequest {
    return new GetRecreationsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecreationsRequest | PlainMessage<GetRecreationsRequest> | undefined, b: GetRecreationsRequest | PlainMessage<GetRecreationsRequest> | undefined): boolean {
    return proto3.util.equals(GetRecreationsRequest, a, b);
  }
}

/**
 * @generated from message admin.v1.GetRecreationsResponse
 */
export class GetRecreationsResponse extends Message<GetRecreationsResponse> {
  /**
   * @generated from field: repeated admin.v1.Recreation recreations = 1;
   */
  recreations: Recreation[] = [];

  constructor(data?: PartialMessage<GetRecreationsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "admin.v1.GetRecreationsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recreations", kind: "message", T: Recreation, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRecreationsResponse {
    return new GetRecreationsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRecreationsResponse {
    return new GetRecreationsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRecreationsResponse {
    return new GetRecreationsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetRecreationsResponse | PlainMessage<GetRecreationsResponse> | undefined, b: GetRecreationsResponse | PlainMessage<GetRecreationsResponse> | undefined): boolean {
    return proto3.util.equals(GetRecreationsResponse, a, b);
  }
}

/**
 * @generated from message admin.v1.PostAdminRecreationRequest
 */
export class PostAdminRecreationRequest extends Message<PostAdminRecreationRequest> {
  /**
   * @generated from field: string problem_id = 1;
   */
  problemId = "";

  constructor(data?: PartialMessage<PostAdminRecreationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "admin.v1.PostAdminRecreationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "problem_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostAdminRecreationRequest {
    return new PostAdminRecreationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostAdminRecreationRequest {
    return new PostAdminRecreationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostAdminRecreationRequest {
    return new PostAdminRecreationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: PostAdminRecreationRequest | PlainMessage<PostAdminRecreationRequest> | undefined, b: PostAdminRecreationRequest | PlainMessage<PostAdminRecreationRequest> | undefined): boolean {
    return proto3.util.equals(PostAdminRecreationRequest, a, b);
  }
}

/**
 * @generated from message admin.v1.PostAdminRecreationResponse
 */
export class PostAdminRecreationResponse extends Message<PostAdminRecreationResponse> {
  /**
   * @generated from field: admin.v1.Recreation recreation = 1;
   */
  recreation?: Recreation;

  constructor(data?: PartialMessage<PostAdminRecreationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "admin.v1.PostAdminRecreationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "recreation", kind: "message", T: Recreation },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PostAdminRecreationResponse {
    return new PostAdminRecreationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PostAdminRecreationResponse {
    return new PostAdminRecreationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PostAdminRecreationResponse {
    return new PostAdminRecreationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: PostAdminRecreationResponse | PlainMessage<PostAdminRecreationResponse> | undefined, b: PostAdminRecreationResponse | PlainMessage<PostAdminRecreationResponse> | undefined): boolean {
    return proto3.util.equals(PostAdminRecreationResponse, a, b);
  }
}

