import { api } from "@/lib/axios";

export interface RevokeAllSessionsRequest {
  revokeCurrent: boolean;
}

export async function revokeAllSessions({
  revokeCurrent,
}: RevokeAllSessionsRequest) {
  await api.delete("/me/sessions", {
    data: {
      revoke_current: revokeCurrent,
    },
  });
}
