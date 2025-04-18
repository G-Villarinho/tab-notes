import { createBrowserRouter } from "react-router-dom";

// Layouts
import { AuthLayout } from "./pages/_layouts/auth";
import { AppLayout } from "./pages/_layouts/app";

// Auth pages
import { LoginPage } from "@/pages/auth/login";
import { LoginAccountNotFoundPage } from "@/pages/auth/login/login-account-not-found";
import { RegisterPage } from "@/pages/auth/register";

// App pages
import { FeedPage } from "@/pages/app/feed";
import { ProfilePage } from "./pages/app/profile";
import { SettingsLayout } from "./pages/app/settings/settings-layout";
import { SecurityAndAccountAcessPage } from "./pages/app/settings/security-and-account-acess";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <AuthLayout />,
    children: [
      {
        path: "login",
        element: <LoginPage />,
      },
      {
        path: "login/account-not-found",
        element: <LoginAccountNotFoundPage />,
      },
      {
        path: "register",
        element: <RegisterPage />,
      },
    ],
  },
  {
    path: "/",
    element: <AppLayout />,
    children: [
      {
        path: "home",
        element: <FeedPage />,
      },
      {
        path: ":username",
        element: <ProfilePage />,
      },
      {
        path: "settings",
        element: <SettingsLayout />,
        children: [
          {
            path: "account",
            element: <div>Account Settings</div>,
          },
          {
            path: "security-and-account-acess",
            element: <SecurityAndAccountAcessPage />,
          },
        ],
      },
    ],
  },
]);
