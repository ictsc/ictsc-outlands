// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file admin/v1/auth.proto (package admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetCallbackRequest, GetCallbackResponse, PostCodeRequest, PostCodeResponse } from "./auth_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service admin.v1.AuthService
 */
export const AuthService = {
  typeName: "admin.v1.AuthService",
  methods: {
    /**
     * @generated from rpc admin.v1.AuthService.GetCallback
     */
    getCallback: {
      name: "GetCallback",
      I: GetCallbackRequest,
      O: GetCallbackResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AuthService.PostCode
     */
    postCode: {
      name: "PostCode",
      I: PostCodeRequest,
      O: PostCodeResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

