// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file admin/v1/answer.proto (package admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetAnswerRequest, GetAnswerResponse, GetAnswersRequest, GetAnswersResponse, PutPointRequest, PutPointResponse } from "./answer_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service admin.v1.AnswerService
 */
export const AnswerService = {
  typeName: "admin.v1.AnswerService",
  methods: {
    /**
     * @generated from rpc admin.v1.AnswerService.GetAnswer
     */
    getAnswer: {
      name: "GetAnswer",
      I: GetAnswerRequest,
      O: GetAnswerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AnswerService.GetAnswers
     */
    getAnswers: {
      name: "GetAnswers",
      I: GetAnswersRequest,
      O: GetAnswersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AnswerService.PutPoint
     */
    putPoint: {
      name: "PutPoint",
      I: PutPointRequest,
      O: PutPointResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

