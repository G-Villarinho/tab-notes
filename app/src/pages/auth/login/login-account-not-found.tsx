import { Helmet } from "react-helmet-async";
import { Navigate, useLocation } from "react-router-dom";
import { LoginAccountNotFoundForm } from "./components/login-account-not-found-form";

interface LocationState {
  email?: string;
}

export function LoginAccountNotFoundPage() {
  const location = useLocation();
  const state = location.state as LocationState | undefined;

  if (!state?.email) {
    return <Navigate to="/login" />;
  }

  return (
    <>
      <Helmet title="Conta não encontrada" />
      <LoginAccountNotFoundForm latestEmail={state.email} />
    </>
  );
}
