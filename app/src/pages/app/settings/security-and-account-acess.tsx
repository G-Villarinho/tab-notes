export function SecurityAndAccountAcessPage() {
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
          <h2 className="font-medium text-base">Sessões ativas</h2>
          <p className="text-sm text-muted-foreground">
            Veja quais dispositivos estão conectados à sua conta.
          </p>
          {/* Aqui você pode futuramente listar sessões com botão de revogar */}
        </div>

        <div className="rounded-xl border p-4 bg-background shadow-sm">
          <h2 className="font-medium text-base">Trocar senha</h2>
          <p className="text-sm text-muted-foreground">
            Recomendado se você suspeitar de atividade suspeita.
          </p>
          <button className="mt-2 text-sm font-medium text-primary hover:underline">
            Alterar senha
          </button>
        </div>
      </div>
    </div>
  );
}
