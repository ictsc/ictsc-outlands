// @generated by protoc-gen-connect-query v1.2.0 with parameter "target=ts,import_extension=none"
// @generated from file contestant/v1/ranking.proto (package contestant.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { GetRankingRequest, GetRankingResponse } from "./ranking_pb";

/**
 * @generated from rpc contestant.v1.RankingService.GetRanking
 */
export const getRanking = {
  localName: "getRanking",
  name: "GetRanking",
  kind: MethodKind.Unary,
  I: GetRankingRequest,
  O: GetRankingResponse,
  service: {
    typeName: "contestant.v1.RankingService"
  }
} as const;
