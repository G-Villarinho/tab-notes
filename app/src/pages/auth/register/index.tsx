import { Helmet } from "react-helmet-async";
import { RegisterForm } from "./register-form";
import { useLocation } from "react-router-dom";

interface LocationState {
  email?: string;
}

export function RegisterPage() {
  const location = useLocation();
  const state = location.state as LocationState | undefined;

  return (
    <>
      <Helmet title="Cadastre-se agora" />
      <RegisterForm email={state?.email} />
    </>
  );
}
