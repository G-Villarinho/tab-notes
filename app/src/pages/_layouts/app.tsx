import { useEffect, useLayoutEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { isAxiosError } from "axios";

import { Header } from "@/components/header";
import { GridPattern } from "@/components/magicui/grid-pattern";
import { api } from "@/lib/axios";
import { cn } from "@/lib/utils";
import { useAuth } from "@/hooks/use-auth";
import { AppLoadingScreen } from "../../components/app-loading-screen";

export function AppLayout() {
  const navigate = useNavigate();

  useLayoutEffect(() => {
    const interceptorId = api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (isAxiosError(error)) {
          const status = error.response?.status;
          if (status === 401) {
            navigate("/login", { replace: true });
          }
        }

        return Promise.reject(error);
      }
    );

    return () => {
      api.interceptors.response.eject(interceptorId);
    };
  }, [navigate]);

  const { isAuthenticated, isLoading } = useAuth();

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      navigate("/login", { replace: true });
    }
  }, [isAuthenticated, isLoading, navigate]);

  if (isLoading) {
    return <AppLoadingScreen />;
  }

  return (
    <div className="relative min-h-screen">
      <GridPattern
        className={cn(
          "[mask-image:radial-gradient(circle at top left, white 30%, transparent 80%)] opacity-30 pointer-events-none absolute inset-0"
        )}
        width={32}
        height={32}
        x={-1}
        y={-1}
      />

      <div className="relative z-10">
        <Header />
        <main className="md:px-72 pt-10">
          <Outlet />
        </main>
      </div>
    </div>
  );
}
