import { Settings, Home, Search, User } from "lucide-react";
import { cn } from "@/lib/utils";
import { useAuth } from "@/hooks/use-auth";
import { Link, useLocation } from "react-router-dom";
import { Separator } from "./ui/separator";

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
    href: "/settings/account",
  },
];

export function MainNav() {
  const { user } = useAuth();
  const location = useLocation();

  return (
    <nav className="rounded-2xl border bg-card shadow-md p-3 space-y-2">
      <div className="text-lg font-semibold text-foreground px-1">Menu</div>
      <Separator className="my-4" />
      <ul className="space-y-1">
        {navItems.map((item, index) => {
          const Icon = item.icon;
          const isActive = location.pathname.startsWith(item.href);

          return (
            <li key={index}>
              <Link
                to={item.href}
                className={cn(
                  "flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-all",
                  isActive
                    ? "bg-primary/10 text-primary"
                    : "hover:bg-muted text-muted-foreground"
                )}
              >
                <Icon className="w-5 h-5 shrink-0" />
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
                "flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-all",
                location.pathname === `/${user.username}`
                  ? "bg-primary/10 text-primary"
                  : "hover:bg-muted text-muted-foreground"
              )}
            >
              <User className="w-5 h-5 shrink-0" />
              <span>Meu perfil</span>
            </Link>
          </li>
        )}
      </ul>
    </nav>
  );
}
