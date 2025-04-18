import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

import { useAuth } from "@/hooks/use-auth";
import Icon from "@/assets/icon.svg";
import { ThemeToggle } from "@/components/theme-toggle";

export function AuthLayout() {
  const { isAuthenticated, isLoading } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isLoading && isAuthenticated) {
      navigate("/home", { replace: true });
    }
  }, [isAuthenticated, isLoading, navigate]);

  return (
    <div className="min-h-svh flex flex-col bg-background">
      <header className="w-full px-4 md:px-6 py-4 flex items-center justify-between">
        <div className="flex items-center gap-2">
          <img src={Icon} alt="Tab Notes" className="w-8 h-8" />
          <span className="text-lg font-bold text-foreground">Tab Notes</span>
        </div>

        <ThemeToggle />
      </header>

      <main className="flex-grow flex justify-center pt-12 md:px-10 py-6">
        <div className="w-full max-w-sm">
          <Outlet />
        </div>
      </main>

      <footer className="w-full border-t border-border px-4 md:px-6 py-6 text-center text-xs text-muted-foreground">
        <div className="flex flex-col items-center gap-2 sm:flex-row sm:justify-between">
          <p>
            © {new Date().getFullYear()}{" "}
            <span className="font-medium">Tab Notes</span>. Todos os direitos
            reservados.
          </p>
          <div className="flex gap-4">
            <a
              href="https://github.com/g-villarinho"
              target="_blank"
              rel="noreferrer"
              className="hover:underline hover:text-primary transition-colors"
            >
              GitHub
            </a>
            <a
              href="/termos"
              className="hover:underline hover:text-primary transition-colors"
            >
              Termos
            </a>
            <a
              href="/privacidade"
              className="hover:underline hover:text-primary transition-colors"
            >
              Privacidade
            </a>
          </div>
        </div>
      </footer>
    </div>
  );
}
