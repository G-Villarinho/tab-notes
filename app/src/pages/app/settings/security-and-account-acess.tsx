import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  getUserSessions,
  GetUserSessionsResponse,
} from "@/api/get-user-sessions";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { AlertCircle } from "lucide-react";
import { format } from "date-fns";
import { ptBR } from "date-fns/locale";
import { Button } from "@/components/ui/button";

function formatDate(date?: string | null) {
  if (!date) return "-";
  return format(new Date(date), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR });
}

export function SecurityAndAccountAcessPage() {
  const queryClient = useQueryClient();

  const {
    data: sessions,
    isLoading,
    isError,
  } = useQuery<GetUserSessionsResponse[]>({
    queryKey: ["user-sessions"],
    queryFn: getUserSessions,
  });

  const { mutate: revokeSession, isPending: isRevoking } = useMutation({
    mutationFn: async (id: string) => {
      await fetch(`/api/me/sessions/${id}`, { method: "DELETE" });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["user-sessions"] });
    },
  });

  return (
    <div className="space-y-6">
      <div className="space-y-1">
        <h1 className="text-xl font-semibold">Segurança e acesso à conta</h1>
        <p className="text-muted-foreground text-sm">
          Gerencie como você acessa sua conta e mantenha sua conta protegida.
        </p>
      </div>

      <div className="space-y-4">
        <div className="rounded-xl border p-4 bg-background shadow-sm">
          <h2 className="font-medium text-base mb-1">Sessões ativas</h2>
          <p className="text-sm text-muted-foreground mb-4">
            Veja quais dispositivos estão conectados à sua conta.
          </p>

          {isLoading && (
            <div className="space-y-2">
              <Skeleton className="h-12 rounded-lg" />
              <Skeleton className="h-12 rounded-lg" />
              <Skeleton className="h-12 rounded-lg" />
            </div>
          )}

          {isError && (
            <div className="text-sm text-destructive flex items-center gap-2">
              <AlertCircle className="w-4 h-4" />
              Erro ao carregar sessões.
            </div>
          )}

          {sessions && sessions.length === 0 && (
            <p className="text-sm text-muted-foreground">
              Nenhuma sessão ativa encontrada.
            </p>
          )}

          {sessions && sessions.length > 0 && (
            <div className="space-y-2">
              {sessions.map((session) => {
                const isRevoked = Boolean(session.RevokedAt);
                const isCurrent = session.id === session.currentSessionId;

                return (
                  <div
                    key={session.id}
                    className="flex items-center justify-between rounded-lg border p-3"
                  >
                    <div className="flex flex-col text-sm">
                      <span className="font-medium">
                        Sessão iniciada em {formatDate(session.createdAt)}
                      </span>
                      <span className="text-muted-foreground">
                        Verificada em {formatDate(session.verifiedAt)}
                      </span>
                      {isRevoked && (
                        <span className="text-xs text-destructive">
                          Revogada em {formatDate(session.RevokedAt)}
                        </span>
                      )}
                    </div>

                    <div className="flex items-center gap-2">
                      {isRevoked ? (
                        <Badge variant="destructive">Revogada</Badge>
                      ) : isCurrent ? (
                        <Badge variant="secondary">Esta sessão</Badge>
                      ) : (
                        <>
                          <Badge variant="default">Ativa</Badge>
                          <Button
                            size="sm"
                            variant="ghost"
                            disabled={isRevoking}
                            onClick={() => revokeSession(session.id)}
                          >
                            Revogar
                          </Button>
                        </>
                      )}
                    </div>
                  </div>
                );
              })}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
