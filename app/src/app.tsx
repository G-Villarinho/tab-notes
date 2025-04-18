import { Helmet, HelmetProvider } from "react-helmet-async";
import { QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider } from "react-router-dom";
import { Toaster } from "react-hot-toast";

import { router } from "@/routes";
import { queryClient } from "@/lib/react-query";
import { ThemeProvider } from "@/contexts/theme/provider";
import { AuthProvider } from "@/contexts/auth/provider";

import "@/index.css";

export function App() {
  return (
    <HelmetProvider>
      <Helmet titleTemplate="%s | tab.notes" />
      <ThemeProvider defaultTheme="light" storageKey="tab-notes-theme">
        <QueryClientProvider client={queryClient}>
          <AuthProvider>
            <Toaster />
            <RouterProvider router={router} />
          </AuthProvider>
        </QueryClientProvider>
      </ThemeProvider>
    </HelmetProvider>
  );
}
