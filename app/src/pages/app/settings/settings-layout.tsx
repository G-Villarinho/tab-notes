import { Outlet, Link, useLocation } from "react-router-dom";
import { MainNav } from "@/components/main-nav";
import { cn } from "@/lib/utils";
import { Input } from "@/components/ui/input";
import { useState } from "react";

const settingsSections = [
  {
    label: "Sua conta",
    href: "/settings/account",
  },
  {
    label: "Segurança e acesso à conta",
    href: "/settings/security-and-account-acess",
  },
  {
    label: "Central de ajuda",
    href: "/settings/help-center",
  },
];

export function SettingsLayout() {
  const location = useLocation();
  const [search, setSearch] = useState("");

  const filteredSections = settingsSections.filter((section) =>
    section.label.toLowerCase().includes(search.toLowerCase())
  );

  return (
    <div className="flex flex-col gap-6 p-4 md:flex-row md:items-start">
      {/* Menu lateral de navegação geral */}
      <aside className="hidden md:flex flex-col w-[220px] shrink-0">
        <MainNav />
      </aside>

      {/* Conteúdo principal com menu de configurações e detalhes */}
      <div className="flex flex-1 gap-4">
        {/* Menu de configurações */}
        <div className="w-[260px] rounded-xl border bg-card/70 p-4 shadow-sm space-y-4">
          <h2 className="text-sm font-semibold text-muted-foreground">
            Configurações
          </h2>
          <div className="relative">
            <Input
              type="text"
              placeholder="Buscar configurações"
              className="pl-9 h-9 text-sm"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
            <span className="absolute left-2 top-1.5 text-muted-foreground">
              🔍
            </span>
          </div>
          <nav className="space-y-1 text-sm">
            {filteredSections.length > 0 ? (
              filteredSections.map((section, index) => {
                const isActive = location.pathname === section.href;

                return (
                  <Link
                    key={index}
                    to={section.href}
                    className={cn(
                      "block rounded-md px-3 py-2 transition",
                      isActive
                        ? "bg-primary text-primary-foreground font-semibold"
                        : "hover:bg-muted"
                    )}
                  >
                    {section.label}
                  </Link>
                );
              })
            ) : (
              <p className="text-sm text-muted-foreground px-2">
                Nenhuma configuração encontrada.
              </p>
            )}
          </nav>
        </div>

        {/* Conteúdo da configuração selecionada */}
        <section className="flex-1 min-h-[400px] rounded-xl border bg-card/70 p-6 shadow-sm">
          <div className="mb-4">
            <h2 className="text-sm font-semibold text-muted-foreground">
              Detalhes
            </h2>
          </div>
          <Outlet />
        </section>
      </div>
    </div>
  );
}
