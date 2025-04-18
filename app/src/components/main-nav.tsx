import { Settings, Home, Search, User } from "lucide-react";
import { cn } from "@/lib/utils";
import { useAuth } from "@/hooks/use-auth";
import { Link, useLocation } from "react-router-dom";

const navItems = [
  {
    label: "Início",
    icon: Home,
    href: "/home",
  },
  {
    label: "Buscar",
    icon: Search,
    href: "/search",
  },
  {
    label: "Configurações",
    icon: Settings,
    href: "/configuracoes",
  },
];

export function MainNav() {
  const { user } = useAuth();
  const location = useLocation();

  return (
    <nav className="rounded-xl border p-4 shadow-sm bg-card space-y-2">
      <ul className="text-sm space-y-1">
        {navItems.map((item, index) => {
          const Icon = item.icon;
          const isActive = location.pathname === item.href;

          return (
            <li key={index}>
              <Link
                to={item.href}
                className={cn(
                  "flex items-center gap-2 rounded-md px-2 py-2 transition hover:bg-muted",
                  isActive && "bg-muted font-medium text-foreground"
                )}
              >
                <Icon className="w-4 h-4" />
                <span>{item.label}</span>
              </Link>
            </li>
          );
        })}

        {user && (
          <li>
            <Link
              to={`/${user.username}`}
              className={cn(
                "flex items-center gap-2 rounded-md px-2 py-2 transition hover:bg-muted",
                location.pathname === `/${user.username}` &&
                  "bg-muted font-medium text-foreground"
              )}
            >
              <User className="w-4 h-4" />
              <span>Meu perfil</span>
            </Link>
          </li>
        )}
      </ul>
    </nav>
  );
}
