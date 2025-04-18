import { LoginForm } from "@/pages/auth/login/components/login-form";
import { Helmet } from "react-helmet-async";

export function LoginPage() {
  return (
    <>
      <Helmet title="login" />
      <LoginForm />
    </>
  );
}
