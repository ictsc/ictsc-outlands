// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=none"
// @generated from file admin/v1/announcement.proto (package admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteAnnouncementRequest, DeleteAnnouncementResponse, GetAnnouncementsRequest, GetAnnouncementsResponse, PatchAnnouncementRequest, PatchAnnouncementResponse, PostAnnouncementRequest, PostAnnouncementResponse } from "./announcement_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service admin.v1.AnnouncementService
 */
export const AnnouncementService = {
  typeName: "admin.v1.AnnouncementService",
  methods: {
    /**
     * @generated from rpc admin.v1.AnnouncementService.GetAnnouncements
     */
    getAnnouncements: {
      name: "GetAnnouncements",
      I: GetAnnouncementsRequest,
      O: GetAnnouncementsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AnnouncementService.PatchAnnouncement
     */
    patchAnnouncement: {
      name: "PatchAnnouncement",
      I: PatchAnnouncementRequest,
      O: PatchAnnouncementResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AnnouncementService.PostAnnouncement
     */
    postAnnouncement: {
      name: "PostAnnouncement",
      I: PostAnnouncementRequest,
      O: PostAnnouncementResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc admin.v1.AnnouncementService.DeleteAnnouncement
     */
    deleteAnnouncement: {
      name: "DeleteAnnouncement",
      I: DeleteAnnouncementRequest,
      O: DeleteAnnouncementResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

