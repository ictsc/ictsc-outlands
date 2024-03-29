// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file contestant/v1/rule.proto (package contestant.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetRuleRequest, GetRuleResponse } from "./rule_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service contestant.v1.RuleService
 */
export const RuleService = {
  typeName: "contestant.v1.RuleService",
  methods: {
    /**
     * @generated from rpc contestant.v1.RuleService.GetRule
     */
    getRule: {
      name: "GetRule",
      I: GetRuleRequest,
      O: GetRuleResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

